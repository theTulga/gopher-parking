package util

import (
	"github.com/alexbrainman/printer"
	"github.com/skip2/go-qrcode"
	"log"
	"golang.org/x/sys/windows"
	"unsafe"
)

var (
	mod = windows.NewLazyDLL("../SDK/sdk_64/POS_SDK.dll")

	procPortOpenW = mod.NewProc("POS_Port_OpenW")
	procTestPage = mod.NewProc("POS_Control_PrintTestpage") 
	procGetPrinterSTatusTSPL = mod.NewProc("GetPrinterStatusTSPL")
	procPrintOneDimensionalBarCodeW = mod.NewProc("POS_Output_PrintOneDimensionalBarcodeW")
	procPrintTwoDimensionalBarCodeW = mod.NewProc("POS_Output_PrintTwoDimensionalBarcodeW")
	procPrintStringW = mod.NewProc("POS_Output_PrintStringW") 
)

func LoadPrinterDLL() {
	 
	printerHandle, returnV, callErr := procPortOpenW.Call(
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("SP-USB1"))),
		uintptr(1002),
		uintptr(0))
	// log.Print(printerHandle)
	// log.Print(returnV)
	// log.Print(callErr)
	// qr, err := qrcode.New("Tsogtgerel Suvd-Erdene Battulga Bat-Orgil Nandin-Erdene", qrcode.Medium)
	// ret, returnV, callErr := procTestPage.Call(printerHandle)
	// ret, returnV, callErr := procPrintStringW.Call(printerHandle, uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("Nandin-Erdene\n\nNANDIAA"))))
	// ret, returnV, callErr := procPrintStringW.Call(printerHandle, uintptr("Nandin-Erdene"))
	// log.Print(ret)
	// log.Print(returnV)
	// log.Print(callErr)

	// ret, returnV, callErr := procPrintTwoDimensionalBarCodeW.Call(
	// 	printerHandle,
	// 	4102,
	// 	2,
	// 	77,
	// 	4,
	// 	// "Test.com",
	// 	uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("Tsogtgerel Suvd-Erdene Battulga Bat-Orgil Nandin-Erdene"))),
	// )
	// log.Print(ret)
	// log.Print(returnV)
	// log.Print(callErr)

	

	ret, returnV, callErr := procPrintOneDimensionalBarCodeW.Call(
		printerHandle,
		4074,
		2,
		50,
		4013,
		// 3123123,
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("Tsogtgerel Suvd-Erdene Battulga Bat-Orgil Nandin-Erdene"))),
	)
	log.Print(ret)
	log.Print(returnV)
	log.Print(callErr)

}


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
	p, err := printer.Open("Microsoft Print to PDF") // Opens the named printer and returns a *Printer
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