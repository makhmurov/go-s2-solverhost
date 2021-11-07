package verify

import (
	"fmt"
	"os"
)

type config struct {
	username  string
	dataURL   string
	verifyURL string
}

const (
	evn_username string = "USER_NAME"
	evn_dataset  string = "DATASET_URL"
	evn_verify   string = "VERIFY_URL"
)

var cfg config = config{
	username:  os.Getenv(evn_username),
	dataURL:   os.Getenv(evn_dataset),
	verifyURL: os.Getenv(evn_verify),
}

func init() {
	var e error = nil
	var errorTpl = "Error: %s environment variable is unset."

	if cfg.dataURL == "" {
		e = fmt.Errorf(errorTpl, evn_dataset)
		fmt.Println(e)
	}

	if cfg.verifyURL == "" {
		e = fmt.Errorf(errorTpl, evn_verify)
		fmt.Println(e)
	}

	if cfg.username == "" {
		e = fmt.Errorf(errorTpl, evn_username)
		fmt.Println(e)
	}

	if e != nil {
		fmt.Println("Exit. Set required environment variables for correct operation.")
		os.Exit(1)
	}
}
