package structure

import "fmt"

type Os interface {
	os_print()
	setPrinter(printer)
}

type printer interface {
	print()
}

type macos struct {
	printer printer
}

func (m *macos) os_print() {
	fmt.Print("macos print ")
	m.printer.print()
}

func (m *macos) setPrinter(p printer) {
	m.printer = p
}

type windowsos struct {
	printer printer
}

func (w *windowsos) os_print() {
	fmt.Print("windows print")
	w.printer.print()
}

func (w *windowsos) setPrinter(p printer) {
	w.printer = p
}

type epson_printer struct {
	printer printer
}

func (e *epson_printer) print() {
	fmt.Println("on epson printer, you need own one")
}

type hp_printer struct {
}

func (h *hp_printer) print() {
	fmt.Println("on hp printer, you need own one")
}
