package screens

// import (
// 	"fyne.io/fyne"
// 	"fyne.io/fyne/layout"
// 	"fyne.io/fyne/widget"
// 	"github.com/thetulga/gopher-parking/models"
// 	"log"
// )

// func makeLabelGroup() fyne.CanvasObject {
// 	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
// 		widget.NewLabel("Орж ирсэн цаг:"),
// 		widget.NewLabel("Зогссон хугацаа:"),
// 		widget.NewLabel("Төлөх дүн:"),
// 	)
// }

// func nnewTicketButton(label string, printText string) fyne.Widget {
// 	btn := widget.NewButton(label, func() {
// 		r := models.NewReceipt("1d34123", label, printText)
// 		log.Print("ADDING ", label)
// 		models.DB.AddReceipt(r)
// 	})
// 	return btn
// }

// func eentryContainer() fyne.Widget {
// 	truckButton := newTicketButton("Truck", "Truck entered parking")
// 	suvButton := newTicketButton("SUV", "SUV entered parking")
// 	sedanButton := newTicketButton("Sedan", "Sedan entered parking")
// 	group := widget.NewGroup("Print Receipt", truckButton, suvButton, sedanButton)
// 	return group
// }

// func infoContainer() fyne.Widget {
// 	return widget.NewGroup("Сүүлийн талоны мэдээлэл",
// 		fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(1),
// 			makeLabelGroup(),
// 		),
// 	)
// }

// func TicketScreen(Rcontainer fyne.CanvasObject) fyne.CanvasObject {

// 	return fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(2),
// 		eentryContainer(),
// 		// infoContainer(),
// 		Rcontainer,
// 	)
// }
