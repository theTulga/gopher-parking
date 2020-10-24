package util

import (
	"github.com/alexbrainman/printer"
	"github.com/skip2/go-qrcode"
	"log"
	"golang.org/x/sys/windows"
	"unsafe"
	"fmt"
)

var (
	mod = windows.NewLazyDLL("../SDK/sdk_64/POS_SDK.dll")

	procPortOpenW = mod.NewProc("POS_Port_OpenW")
	procTestPage = mod.NewProc("POS_Control_PrintTestpage") 
	procPrintStringW = mod.NewProc("POS_Output_PrintStringW") 
	procGetPrinterSTatusTSPL = mod.NewProc("GetPrinterStatusTSPL")
)

func LoadPrinterDLL() {
	 
	printerHandle, returnV, callErr := procPortOpenW.Call(
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("SP-USB1"))),
		uintptr(1002),
		uintptr(0))
	log.Print(printerHandle)
	log.Print(returnV)
	log.Print(callErr)

	ret, returnV, callErr := procTestPage.Call(printerHandle)
	log.Print(ret)
	log.Print(returnV)
	log.Print(callErr)
	
	ret, returnV, callErr = procGetPrinterSTatusTSPL.Call(printerHandle)
	log.Print(ret)
	log.Print(returnV)
	log.Print(callErr)
	
	log.Print(procTestPage)
	for _, m := range procTestPage {

    // m is a map[string]interface.
    // loop over keys and values in the map.
    for k, v := range m {
        fmt.Println(k, "value is", v)
    }
}

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