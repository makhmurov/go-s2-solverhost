package env

import "os"

type config struct {
	username  string
	dataURL   string
	verifyURL string
}

var cfg config

func (c *config) Init() {
	c.username = os.Getenv("USER_NAME")
	c.dataURL = os.Getenv("TASKS_URI")
	c.verifyURL = os.Getenv("VERIFY_URI")
}

func init() {
	cfg.Init()
}
