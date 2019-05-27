package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	}, "")

	vbox := gtk.NewVBox(false, 16)


	window.SetSizeRequest(600, 600)

	btn := gtk.NewButtonWithLabel("Sort")
	btn.SetSizeRequest(60, 0)

	text := gtk.NewEntry()
	text.SetTooltipText("Write Name Here...")
	text.Widget.SetSizeRequest(200, 100)
	text.SetPosition(2)

	btn.Clicked(func() {
		content := text.GetText()
		values := strings.Split(content, ",")

		sort.Strings(values)

		str := ""
		for _, v := range values {
			str += v + "\n"
		}
		log.Println(str)

		err := ioutil.WriteFile("sorted.txt", []byte(str), os.ModePerm)
		if err != nil {
			log.Println(err)
		}

		l := gtk.NewLabel("Name sorted")
		vbox.Add(l)
	})

	vbox.Add(text)
	vbox.Add(btn)
	window.Add(vbox)
	window.ShowAll()
	gtk.Main()
}