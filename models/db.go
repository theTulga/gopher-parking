package models

import (
	"fyne.io/fyne/widget"
	"github.com/thetulga/gopher-parking/util"
)

var DB *Database

type Database struct {
	Entries           []Entry
	ReceiptsContainer *widget.Box
	Receipts          []Receipt
}

func (db Database) AddReceipt(r Receipt) {

	db.Receipts = append(db.Receipts, r)
	// label := widget.NewLabel(r.Label)
	// db.ReceiptsContainer.Append(label)
	util.Print(r.id.String())
}

func (db Database) AddElement(e *Entry) {
	db.Entries = append(db.Entries, *e)
}
