package structure

import "testing"

func TestBridge(t *testing.T) {
	hp := new(hp_printer)
	epson := new(epson_printer)
	mac := new(macos)
	mac.setPrinter(hp)
	mac.os_print()
	windows := new(windowsos)
	windows.setPrinter(hp)
	windows.os_print()
	windows.setPrinter(epson)
	windows.os_print()
	mac.setPrinter(epson)
	mac.os_print()
}
