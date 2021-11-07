package env

import "os"

type config struct {
	username  string
	dataURL   string
	verifyURL string
}

var cfg config = config{
	username:  os.Getenv("USER_NAME"),
	dataURL:   os.Getenv("DATASET_URL"),
	verifyURL: os.Getenv("VERIFY_URL"),
}

func (c *config) Init() {
	c.username = os.Getenv("USER_NAME")
	c.dataURL = os.Getenv("DATASET_URL")
	c.verifyURL = os.Getenv("VERIFY_URI")
}

func init() {
	cfg.Init()
}

func Username() string {
	return cfg.username
}

func DataURL() string {
	return cfg.dataURL
}

func VerifyURL() string {
	return cfg.verifyURL
}
