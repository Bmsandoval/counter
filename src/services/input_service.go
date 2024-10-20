package services

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/bmsandoval/counter/src/ui"
)

type InputService interface {
	PromptUserForFileLoc() (string, error)
}

type inputServiceImpl struct{}

func NewInputService() InputService {
	return &inputServiceImpl{}
}

func (s inputServiceImpl) PromptUserForFileLoc() (string, error) {
	a := app.New()
	w := a.NewWindow("TODO App")
	ui.CreateNewFile(w)
	w.SetContent(widget.NewLabel("TODOs will go here"))
	w.ShowAndRun()
	return "", nil
}
