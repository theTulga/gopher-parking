package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/thetulga/gopher-parking/models"
	"github.com/thetulga/gopher-parking/screens"
	"fmt"
)

const preferenceCurrentTab = "currentTab"

func main() {
	fmt.Print("HELLO")
	a := app.New()
	w := a.NewWindow("Box Layout")
	w.Resize(fyne.NewSize(1200, 800))
	models.DB = new(models.Database)
	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Home", theme.HomeIcon(), screens.HomeScreen()),
	)

	tabs.SetTabLocation(widget.TabLocationLeading)
	tabs.SelectTabIndex(a.Preferences().Int(preferenceCurrentTab))
	w.SetContent(tabs)
	w.ShowAndRun()
	a.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())
}
