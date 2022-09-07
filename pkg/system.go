package goreloaded

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(fileName string) []byte {
	// Open file for reading
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func WriteToFile(fileName string, text string) error {
	err := ioutil.WriteFile(fileName, []byte(text), 0666)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func checkError(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
