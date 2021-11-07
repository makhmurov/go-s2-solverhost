package verify

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
