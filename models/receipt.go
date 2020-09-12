package models

import (
	"fyne.io/fyne"
	"time"
)

type Receipt struct {
	id        string
	Label     string `default:"label"`
	PrintText string `default:"text"`
	EnteredAt time.Time
	Widget    fyne.Widget
}

func NewReceipt(id string, label string, content string) Receipt {
	return Receipt{
		id:        id,
		Label:     label,
		PrintText: content,
		EnteredAt: time.Now(),
	}
}

// func NewReceiptUI() receiptUI {
// 	UI := receiptUI{
// 		Container: fyne.NewContainer(),
// 		List:      []fyne.Widget{},
// 	}

// 	for _, r := range models.DB.Receipts {
// 		label := widget.NewLabel(r.Label)
// 		UI.List = append(UI.List, label)
// 		UI.Container.AddObject(label)
// 	}

// 	return UI
// }

// func (ui *receiptUI) AddReceipt(r models.Receipt) {

// 	label := widget.NewLabel(r.Label)
// 	ui.Container.AddObject(label)

// }
