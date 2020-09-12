package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/thetulga/gopher-parking/models"
	"log"
)

func makeLabelGroup() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		widget.NewLabel("Орж ирсэн цаг:"),
		widget.NewLabel("Зогссон хугацаа:"),
		widget.NewLabel("Төлөх дүн:"),
	)
}

func newTicketButton(label string, printText string) fyne.Widget {
	btn := widget.NewButton(label, func() {
		r := models.NewReceipt("1234", label, printText)
		log.Print("ADDING ", label)
		models.DB.AddReceipt(r)
	})
	return btn
}

func entryContainer() fyne.Widget {
	truckButton := newTicketButton("Том тээврийн хэрэгсэл", "Том тээврийн хэрэгслийн талон хэвлэгдлээ.")
	suvButton := newTicketButton("Дунд тээврийн хэрэгсэл", "Дунд тээврийн хэрэгслийн талон хэвлэгдлээ.")
	sedanButton := newTicketButton("Жижиг тээврийн хэрэгсэл", "Жижиг тээврийн хэрэгслийн талон хэвлэгдлээ.")
	group := widget.NewGroup("Талон хэвлэх", truckButton, suvButton, sedanButton)
	return group
}

func infoContainer() fyne.Widget {
	return widget.NewGroup("Сүүлийн талоны мэдээлэл",
		fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(1),
			makeLabelGroup(),
		),
	)
}

func TicketScreen(Rcontainer fyne.CanvasObject) fyne.CanvasObject {

	return fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(2),
		entryContainer(),
		// infoContainer(),
		Rcontainer,
	)
}
