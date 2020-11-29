package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("example/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	raw := string(dat)
	r := csv.NewReader(strings.NewReader(raw))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}
