package app

import "github.com/rivo/tview"

const (
	blogsModalName = "blogsModal"
)

var blogsModal = tview.NewModal().
	SetText("Resize the window to see how the grid layout adapts").
	AddButtons([]string{"Ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
	pages.HidePage(blogsModalName)
})

var blogsModalF = tview.NewForm().
	AddInputField("First name:", "", 20, nil, nil).
	AddButton("Search", searchBlogs).
	AddButton("Clear", cancelSearchBlogs).
	SetBorder(true).
	SetTitle("Search")

	// 	AddButtons([]string{"Ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
	// 	pages.HidePage(blogsModalName)
	// })

	// f := tview.NewForm().

	// 		AddInputField("Last name:", "", 20, nil, nil).
	// 		AddDropDown("Role:", []string{"Engineer", "Manager", "Administration"}, 0, nil).
	// 		AddCheckbox("On vacation:", false, nil).
	// 		AddPasswordField("Password:", "", 10, '*', nil).
	// 		AddButton("Save", nextSlide).
	// 		AddButton("Cancel", nextSlide)
	// 	f.SetBorder(true).SetTitle("Employee Information")
