package util

import (
	"github.com/alexbrainman/printer"
	"github.com/skip2/go-qrcode"
	// "image/png"
	"log"
	// "os"
)

func GenerateQRCode(str string) {
	err := qrcode.WriteFile(str, qrcode.Medium, 256, "./public/images/" + str + ".png")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Created Image")
}

func PrintPrinterNames() {
	printers, err := printer.ReadNames()
	if err != nil {
		log.Fatal("error in reading printer names: ",err)
	}
	log.Print("available printers:")
	for _, printer := range printers {
		log.Print(printer)
	}
}

func Print(str string) {
	log.Print("Printer name: ", "AnyDesk Printer")
	p, err := printer.Open("80mm Series Printer") // Opens the named printer and returns a *Printer
	if err != nil {
	    log.Fatal(err)
	}
	err = p.StartDocument("test", "text") // test: doc name, text: doc type
	if err != nil {
	    log.Fatal(err)
	}
	err = p.StartPage() // begin a new page
	if err != nil {
	    log.Fatal(err)
	}
	n, err := p.Write([]byte("Suudliin Mashin orloo\nAraas n tom mashin orloo")) // Send some text to the printer
	if err != nil {
	    log.Fatal(err)
	}
	log.Print("Num of bytes written to printer:", n)
	err = p.EndPage() // end of page
	if err != nil {
	    log.Fatal(err)
	}
	err = p.EndDocument() // end of document
	if err != nil {
	    log.Fatal(err)
	}
	err = p.Close() // close the resource
	if err != nil {
	    log.Fatal(err)
	}
}