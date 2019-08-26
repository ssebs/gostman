package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ssebs/gostman/utils"
)

func parseFlags() *utils.Request {
	// Usage: gostman [options] <url>
	// e.g.: gostman -M POST -D '{"username": "test", "password": "foo"}' -H 'Authorization: foobar' https://api.example.com/login/

	flag.Usage = func() {
		msg := "\nUsage: gostman [options] <url>\ne.g. gostman -method POST -data '{\"username\": \"test\", \"password\": \"foo\"}'\n\n"
		fmt.Fprintf(os.Stderr, msg)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
	}

	method := flag.String("method", "GET", "HTTP Method (GET, POST, etc).")
	data := flag.String("data", "{}", "Data for request in JSON format")
	headers := flag.String("headers", "{}", "Headers for request.")

	flag.Parse()
	url := flag.Arg(0)

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
	req.GetData()
	// println(req.GetMethod())
}
