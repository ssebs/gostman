package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

func doRequest(httpType string, url string, data ...string) string {
	var outstr string
	log.Println(data)

	switch httpType {
	case "get":
		log.Println("GET" + url)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal("HTTP Error", err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		outstr = string(body)
	case "post":
		log.Println("POST" + url)
	}
	return outstr
}

func addWidgets(grid *gtk.Grid) {

	sw, _ := gtk.ScrolledWindowNew(nil, nil)
	sw.SetHExpand(true)
	sw.SetVExpand(true)

	space, _ := gtk.LabelNew("")
	title, _ := gtk.LabelNew("gostman")
	reqTypeLabel, _ := gtk.LabelNew("Request Method:")
	urlLabel, _ := gtk.LabelNew("URL:")
	jsonDataLabel, _ := gtk.LabelNew("JSON Data:")

	urlEntry, _ := gtk.EntryNew()
	respTxt, _ := gtk.TextViewNew()
	jsonDataEntry, _ := gtk.TextViewNew()
	submit, _ := gtk.ButtonNewWithLabel("Submit Request")

	// respTxt.SetEditable(false)
	respTxt.SetWrapMode(gtk.WRAP_WORD)

	SubmitReq := func() {
		urlTxt, _ := urlEntry.GetText()
		buff, _ := respTxt.GetBuffer()
		// https://reqres.in/api/users/
		buff.SetText(jsonPrettyPrint(doRequest("get", urlTxt)))
	}

	submit.Connect("clicked", SubmitReq)
	urlEntry.Connect("activate", SubmitReq)

	/*
		-------------
		1|	title	| => gostman
		2| labl	inp	| => ReqType input
		3| labl	inp	| => url input
		4| labl	inp	| => jsondata input
		5|	Btn		| => input
		6|	Divider	| =>
		7|	Output	| => outputtxt
		-------------
	*/

	// Attach(widget, XPos, YPos, Width, Height)
	grid.Attach(title, 0, 0, 3, 1)
	grid.Attach(reqTypeLabel, 0, 1, 1, 1)

	grid.Attach(urlLabel, 0, 2, 1, 1)
	grid.Attach(urlEntry, 1, 2, 2, 1)
	grid.Attach(jsonDataLabel, 0, 3, 1, 1)
	grid.Attach(jsonDataEntry, 1, 3, 2, 4)
	grid.Attach(space, 0, 4, 3, 3)
	grid.Attach(submit, 0, 7, 3, 1)

	sw.Add(respTxt)
	grid.Attach(sw, 0, 8, 3, 1)
	// grid.Attach(respTxt, 0, 7, 3, 1)
}

func main() {
	const appID = "com.ssebs.gostman"
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal("Could not create application.", err)
	}

	app.Connect("activate", func() {
		window, err := gtk.ApplicationWindowNew(app)
		if err != nil {
			log.Fatal("Unable to create window:", err)
		}
		window.SetTitle("gostman")
		window.SetDefaultSize(800, 600)

		grid, err := gtk.GridNew()
		if err != nil {
			log.Fatal("Unable to create grid:", err)
		}

		grid.SetOrientation(gtk.ORIENTATION_VERTICAL)
		addWidgets(grid)
		window.Add(grid)
		window.ShowAll()

	})

	app.Connect("destroy", func() {
		gtk.MainQuit()
	})

	app.Run(os.Args)
}
