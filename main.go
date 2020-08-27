package main

import (
	"github.com/joho/godotenv"
	"github.com/rivo/tview"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
