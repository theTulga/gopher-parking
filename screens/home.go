package screens

import (
	"fyne.io/fyne"
	// "fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	// "image/color"
	"log"
)

type ParkingCard struct {
	entryTime *widget.Label
	duration  *widget.Label
	dueAmount *widget.Label
}

func loadInfoContainer(e *ParkingCard) fyne.Widget {
	return widget.NewGroup("Сүүлийн Талоны мэдээлэл",
		fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(2),
			makeLabelGroup(),
			makeInfoGroup(e),
		),
	)
}

func makeLabelGroup() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		widget.NewLabel("Орж ирсэн цаг:"),
		widget.NewLabel("Зогссон хугацаа:"),
		widget.NewLabel("Төлөх дүн:"),
	)
}

func makeInfoGroup(e *ParkingCard) fyne.CanvasObject {
	ticketGroup := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		e.entryTime,
		e.duration,
		e.dueAmount,
	)
	return ticketGroup
}

func loadTestingGroup(e *ParkingCard) fyne.Widget {
	group := widget.NewGroup("Тестлэх хэсэг",
		widget.NewButton("Талон уншуулах тест", func() {
			e.duration.SetText("1h 30min")
			e.entryTime.SetText("9:48AM")
			e.dueAmount.SetText("$1.50")
			log.Print("Changing value of duration to 1hour 30min")
		}),
	)
	return group
}

// HomeScreen loads a graphics example panel for the demo app
func HomeScreen() fyne.CanvasObject {
	e := &ParkingCard{
		entryTime: widget.NewLabel(""),
		duration:  widget.NewLabel(""),
		dueAmount: widget.NewLabel(""),
	}

	return fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(2),
		loadTestingGroup(e),
		loadInfoContainer(e),
	)
}
