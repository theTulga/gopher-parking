package util

import (
	"github.com/alexbrainman/printer"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"log"
	"os"
)

func GenerateQRCode(str string) {
	qrCode, _ := qr.Encode(str, qr.M, qr.Unicode)

	qrCode, _ = barcode.Scale(qrCode, 200, 200)
	log.Print(qrCode)
	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	png.Encode(file, qrCode)
}

func Print(str string) {
	name, err := prt.Default() // returns name of Default Printer as string
	p, err := prt.Open(name)   // Opens the named printer and returns a *Printer
	if err != nil {
		log.Print("NO printer named", name)
		log.fatal(err)
		return
	}
	err = p.StartDocument("test", "text") // test: doc name, text: doc type
	if err != nil {
		log.Print("Cant start new document")
		log.fatal(err)
		return
	}
	err = p.StartPage() // begin a new page
	if err != nil {
		log.Print("Cant start new page")
		log.fatal(err)
	}
	n, err := p.Write([]byte(str)) // Send some text to the printer
	if err != nil {
		log.Print("Cant start write", str)
		log.fatal(err)
	}
	err = p.PageEnd() // end of page
	if err != nil {
		log.Print("Cant end page")
		log.fatal(err)
	}
	err = p.DocumentEnd() // end of document
	if err != nil {
		log.Print("Cant end document")
		log.fatal(err)
	}
	err = p.Close() // close the resource
	if err != nil {
		log.fatal(err)
	}
}
