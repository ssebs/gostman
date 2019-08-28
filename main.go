package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ssebs/gostman/utils"
)

func parseFlags() *utils.Request {
	// Usage: gostman [options] <url>
	// e.g.: gostman -method POST -data '{"username": "test", "password": "foo"}' -headers 'Authorization: foobar' https://api.example.com/login/

	flag.Usage = func() {
		msg := "\nUsage: gostman [options] <url>\ne.g. gostman -method POST -data '{\"username\": \"test\", \"password\": \"foo\"}'\n\n"
		fmt.Fprintf(os.Stderr, msg)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
	}

	method := flag.String("method", "GET", "HTTP Method (GET, POST, etc).")
	data := flag.String("data", "{}", "Data for request in JSON format")
	headers := flag.String("headers", "Content-Type: application/json", "Headers for request.")

	flag.Parse()
	url := flag.Arg(0)

	if url == "" {
		log.Fatal("You must supply a URL")
	} else if !(strings.Contains(url, "http://") || strings.Contains(url, "https://")) {
		log.Fatal("URL must contain http:// or https://")
	}

	println(*method)
	println(*data)
	println(*headers)
	println(url)
	return utils.NewRequest(*method, *data, *headers, url)
}

func main() {
	// b := 5
	// println(utils.AddTwo(&b))
	// println("test")

	req := parseFlags()
	switch req.GetMethod() {
	case "GET":
		body, statusCode := utils.DoGET(req.GetURL(), req.GetHeaders())

		println("Response:" + string(statusCode) + "\n")
		println(body)
	case "POST":
		println("do post")
	}
}
