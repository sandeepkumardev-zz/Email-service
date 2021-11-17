package utils

import (
	"log"
	"os"
)

func Logger(str string) {
	f, err := os.OpenFile("logfile.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "prefix", log.LstdFlags)
	logger.Println(str)
}
