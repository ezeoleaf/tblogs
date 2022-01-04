package main

import (
	"github.com/ezeoleaf/tblogs/app"
	"github.com/ezeoleaf/tblogs/data"
)

func main() {
	ds := data.NewService()

	a := app.NewApp(ds)

	a.Start()
}
