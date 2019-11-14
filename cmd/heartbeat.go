package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/deliangyang/heartbeat/internal/pkg"
)

var (
	conf string
)

func init() {
	flag.StringVar(&conf, "conf", "", "mail config filename")
	flag.Parse()
}

func main() {
	if err := pkg.LoadFile(conf); err != nil {
		log.Fatal(err)
	}
	conf := pkg.GetConfig()

	wg := sync.WaitGroup{}

	for _, website := range conf.Websites {
		wg.Add(1)
		go func(w pkg.Website) {
			for {
				log.Println("start", w.URL)
				if err := w.Check(); err != nil {
					log.Println(err)
					message := fmt.Sprintf("website: %s has been downtime,\nerror: %s", w.URL, err.Error())
					if err := pkg.SendMail(conf.Mail, w.To, message); err != nil {
						wg.Done()
						log.Fatal(err)
					}
				}
				log.Println("end", w.URL)
				time.Sleep(time.Duration(w.Minute) * time.Minute)
			}
		}(website)
	}

	wg.Wait()
}
