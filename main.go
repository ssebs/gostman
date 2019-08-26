package main

import "github.com/ssebs/gostman/utils"

func parseFlags() *utils.Request {
	// Usage: gostman [options] <url>
	// e.g.: gostman -M POST -D '{"username": "test", "password": "foo"} -H 'Authorization: foobar' https://api.example.com/login/

	m := "GET"
	d := ""
	h := "Authorization foo bar"
	u := "https://example.com/api/v2"

	return utils.NewRequest(m, d, h, u)
}

func main() {
	// b := 5
	// println(utils.AddTwo(&b))
	// println("test")

	req := parseFlags()
	println(req.GetMethod())
}
