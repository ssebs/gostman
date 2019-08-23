package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func doRequest(httpType string, url string, data ...string) string {
	switch httpType {
	case "get":
		log.Println("GET" + url)
	case "post":
		log.Println("POST" + url)
	}
	return "blah"
}

func addWidgets(grid *gtk.Grid) {

	label, err := gtk.LabelNew("gostman")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create entry", err)
	}

	btn, err := gtk.ButtonNewWithLabel("Submit Request")
	if err != nil {
		log.Fatal("Unable to create btn", err)
	}
	txtArea, err := gtk.TextViewNew()
	if err != nil {
		log.Fatal("Unable to create txtview")
	}
	txtArea.SetEditable(false)

	btn.Connect("clicked", func() {
		txt, _ := entry.GetText()
		buff, _ := txtArea.GetBuffer()
		buff.SetText(doRequest(txt, "http://httpbin.org/get"))
	})
	entry.Connect("activate", func() {
		txt, _ := entry.GetText()
		buff, _ := txtArea.GetBuffer()
		buff.SetText(doRequest(txt, "http://httpbin.org/get"))
	})

	grid.Attach(label, 0, 0, 1, 1)
	grid.Attach(entry, 0, 1, 2, 1)
	grid.Attach(btn, 0, 2, 1, 1)
	grid.Attach(txtArea, 0, 3, 2, 1)
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
