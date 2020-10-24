package screens

import (
	"github.com/google/uuid"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"log"
	"github.com/thetulga/gopher-parking/models"
	"github.com/thetulga/gopher-parking/util"
)

func newTicketButton(label string, printText string) fyne.Widget {
	btn := widget.NewButton(label, func() {
		r := models.NewReceipt(uuid.New(), label, printText)
		log.Print("ADDING ", label)
		models.DB.AddReceipt(r)
	})
	return btn
}


func entryContainer() fyne.Widget {
	truckButton := newTicketButton("1.5-аас дээш даацтай том оврийн ачааны машин.", "ТТТТТТТТТТТТТТ.")
	suvButton := newTicketButton("1.5 тонн даацтай жижиг оврын ачааны машин", "Дунд тээврийн хэрэгслийн талон хэвлэгдлээ.")
	sedanButton := newTicketButton("SEDAN", "Жижиг тээврийн хэрэгслийн талон хэвлэгдлээ.")
	group := widget.NewGroup("Print Receipt", truckButton, suvButton, sedanButton)
	return group
}
func HomeScreen() fyne.CanvasObject {
	util.PrintPrinterNames()
	util.LoadPrinterDLL()
	return fyne.NewContainerWithLayout(
		layout.NewAdaptiveGridLayout(2),
		entryContainer(),
	)
}