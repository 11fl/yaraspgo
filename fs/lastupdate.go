package fs

import (
	"log"
	"os"
)

func WriteUpdate(d string) {
	data := []byte(d)
	err := os.WriteFile("lastupdate", data, 0644)

	if err != nil {
		log.Fatal("Can't write to file :", err)
	}
}

func ReadUpdate() string {
	dat, err := os.ReadFile("lastupdate")
	if err != nil {
		log.Fatal("Can't readfile :", err)
	}
	return string(dat)
}
