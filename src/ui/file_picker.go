package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"os"
)

// CreateNewFile creates a new file through a save dialog
// initializes it with a counter value of 0, and handles errors.
func CreateNewFile(window fyne.Window) {
	//rChan := make(chan string, 1)
	fileSaveDialog := dialog.NewFileSave(
		func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if writer == nil {
				return
			}

			// Get the file path and create the file with an initial counter
			filePath := writer.URI().Path()
			fmt.Println("New file created:", filePath)

			// Initialize the new file with a counter value of 0
			err = os.WriteFile(filePath, []byte("0"), 0644)
			if err != nil {
				dialog.ShowError(err, window)
			} else {
				fmt.Println("New file initialized with a counter.")
			}
		}, window)

	// Set a filter to save only as a .txt file
	fileSaveDialog.SetFileName("newfile.txt")
	fileSaveDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
	fileSaveDialog.Show()
}
