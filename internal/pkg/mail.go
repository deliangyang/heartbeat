package pkg

import (
	"crypto/tls"

	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
)

var (
	dialer *gomail.Dialer
)

// MailConfig 邮件配置
type MailConfig struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	From     string `toml:"from"`
	Port     int    `toml:"port"`
}

// LoadDialer 加载gomail dialer
func (c MailConfig) LoadDialer() *gomail.Dialer {
	if dialer == nil {
		dialer = gomail.NewDialer(c.Host, c.Port, c.Username, c.Password)
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return dialer
}

// SendMail 发送邮件
func SendMail(c MailConfig, to []string, message string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", c.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "网站挂了！！！")
	m.SetBody("text/html", message)
	if err := c.LoadDialer().DialAndSend(m); err != nil {
		return errors.Wrap(err, "send mail fail")
	}
	return nil
}
