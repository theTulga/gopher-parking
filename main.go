package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	// "fyne.io/fyne/canvas"
	// "fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/thetulga/p/screens"
	// "image/color"
	// "log"
)

const preferenceCurrentTab = "currentTab"

func main() {
	a := app.New()
	w := a.NewWindow("Box Layout")
	w.Resize(fyne.NewSize(1200, 800))
	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Талон", theme.SearchIcon(), screens.HomeScreen()),
		widget.NewTabItemWithIcon("Graphics", theme.DocumentCreateIcon(), screens.GraphicsScreen()),
		widget.NewTabItemWithIcon("Widgets", theme.CheckButtonCheckedIcon(), screens.WidgetScreen()),
		widget.NewTabItemWithIcon("Containers", theme.ViewRestoreIcon(), screens.ContainerScreen()),
		widget.NewTabItemWithIcon("WindowsSS", theme.ViewFullScreenIcon(), screens.DialogScreen(w)))
	tabs.SetTabLocation(widget.TabLocationLeading)
	tabs.SelectTabIndex(a.Preferences().Int(preferenceCurrentTab))
	w.SetContent(tabs)
	w.ShowAndRun()
	a.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())
}
