package models

import (
	"fyne.io/fyne"
	// "fyne.io/fyne/widget"
	"time"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	"log"
)

type Receipt struct {
	id        uuid.UUID
	Label     string `default:"label"`
	PrintText string `default:"text"`
	EnteredAt time.Time
	Widget    fyne.Widget
}

func NewReceipt(id uuid.UUID, label string, content string) Receipt {
	err := qrcode.WriteFile(id.String(), qrcode.Medium, 256, "./public/images/" + id.String() + ".png")
	if err != nil {
		log.Print(err)
	}
	return Receipt{
		id:        id,
		Label:     label,
		PrintText: content,
		EnteredAt: time.Now(),
	}
}
