package ui

//func DisplayGUI() {
//	myWindow := fyne.CurrentApp().NewWindow("Key Config Example")
//
//	changeIncrementKeyButton := widget.NewButton("Set Increment Key", func() {
//		entry := widget.NewEntry()
//		entry.SetPlaceHolder("Enter key combination (e.g., ctrl+alt+up)")
//
//		// Call the package function to listen for the hotkey input.
//		pkg.ListenForHotkey() // Assuming this captures the key combination from user.
//
//		// Create a dialog to capture key combination input
//		dialog.ShowForm("Set Increment Key", "Save", "Cancel", []*widget.FormItem{
//			widget.NewFormItem("Key Combination", entry),
//		}, func(confirm bool) {
//			// If the user confirms and the input isn't empty
//			if confirm && entry.Text != "" {
//				// Update the config with the new increment key
//				config.IncrementKey = entry.Text
//
//				// Save the updated configuration
//				saveConfig(config)
//
//				// Update the label with the new increment key
//				incrementLabel.SetText("Current Increment Key: " + config.IncrementKey)
//
//				// Print a confirmation message
//				fmt.Println("Increment key updated:", config.IncrementKey)
//			}
//		}, myWindow)
//	})
//
//	// Layout code (add button and label to window)
//	myWindow.SetContent(widget.NewVBox(
//		changeIncrementKeyButton,
//		incrementLabel,
//	))
//
//	myWindow.ShowAndRun()
//	//
//	//
//	//
//	//
//	//
//	//
//	//// Load the config file to get the key combinations
//	//config, err := loadConfig()
//	//if err != nil {
//	//	fmt.Println("Error loading config:", err)
//	//}
//	//
//	//// Set up the Fyne app
//	//myApp := app.New()
//	//myWindow := myApp.NewWindow("Counter File Manager")
//	//
//	//// Label to display the current key combinations
//	//incrementLabel := widget.NewLabel("Current Increment Key: " + config.IncrementKey)
//	//decrementLabel := widget.NewLabel("Current Decrement Key: " + config.DecrementKey)
//	//
//	//// Button to create a new file
//	//createFileButton := widget.NewButton("Create New Counter File", func() {
//	//	pkg.CreateNewFile(myWindow)
//	//})
//	//
//	//// Button to change the increment key combination
//	//changeIncrementKeyButton := widget.NewButton("Set Increment Key", func() {
//	//	//rChan := make(chan string, 1)
//	//	entry := widget.NewEntry()
//	//	entry.SetPlaceHolder("Enter key combination (e.g., ctrl+alt+up)")
//	//
//	//	pkg.ListenForHotkey()
//	//	//
//	//	//// Create a dialog to set the increment key
//	//	//dialog.ShowForm("Set Increment Key", "Save", "Cancel", []*widget.FormItem{
//	//	//	widget.NewFormItem("Key Combination", entry),
//	//	//}, func(confirm bool) {
//	//	//	if confirm && entry.Text != "" {
//	//	//		config.IncrementKey = entry.Text
//	//	//		saveConfig(config)
//	//	//		incrementLabel.SetText("Current Increment Key: " + config.IncrementKey)
//	//	//		fmt.Println("Increment key updated:", config.IncrementKey)
//	//	//	}
//	//	//}, myWindow)
//	//})
//	//
//	//// Button to change the decrement key combination
//	//changeDecrementKeyButton := widget.NewButton("Set Decrement Key", func() {
//	//	entry := widget.NewEntry()
//	//	entry.SetPlaceHolder("Enter key combination (e.g., ctrl+alt+down)")
//	//
//	//	// Create a dialog to set the decrement key
//	//	dialog.ShowForm("Set Decrement Key", "Save", "Cancel", []*widget.FormItem{
//	//		widget.NewFormItem("Key Combination", entry),
//	//	}, func(confirm bool) {
//	//		if confirm && entry.Text != "" {
//	//			config.DecrementKey = entry.Text
//	//			saveConfig(config)
//	//			decrementLabel.SetText("Current Decrement Key: " + config.DecrementKey)
//	//			fmt.Println("Decrement key updated:", config.DecrementKey)
//	//		}
//	//	}, myWindow)
//	//})
//	//
//	//// Main layout
//	//myWindow.SetContent(container.NewVBox(
//	//	createFileButton,
//	//	incrementLabel,
//	//	changeIncrementKeyButton,
//	//	decrementLabel,
//	//	changeDecrementKeyButton,
//	//))
//	//
//	//// Register global hotkeys for increment and decrement
//	//registerGlobalHotkey(config.IncrementKey, func() {
//	//	modifyCounter(1)
//	//})
//	//registerGlobalHotkey(config.DecrementKey, func() {
//	//	modifyCounter(-1)
//	//})
//	//
//	//// Show the window
//	//myWindow.ShowAndRun()
//}
