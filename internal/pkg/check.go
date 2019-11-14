package pkg

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Website 网站
type Website struct {
	URL    string   `toml:"url"`
	Minute int64    `toml:"minute"`
	To     []string `toml:"to"`
}

// Websites 许多网站
type Websites []Website

// Check 检查http
func (w Website) Check() error {
	resp, err := http.Head(w.URL)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("check site %s fail", w.URL))
	}
	if resp.StatusCode == 200 {
		return nil
	}

	return fmt.Errorf("website: %s, ping fail, the http status code %d", w.URL, resp.StatusCode)
}
