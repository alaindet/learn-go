package main

/*
The OldPrinter struct is forced to implement Fax() and Scan() even though only
because the Machine interface is not "segregated" enough and does too many things

The Printer and Scanner interface are better than Machine since their specialized

Smaller interfaces can be composed into bigger ones, like MultiFunctionMachine
*/

type Document struct {
	// ...
}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// // This is a "god" interface
// type Machine interface {
// 	Print(d Document)
// 	Fax(d Document)
// 	Scan(d Document)
// }

type MultiFunctionPrinter struct {
	// ...
}

func (p *MultiFunctionPrinter) Print(d Document) {
	panic("implement me")
}

func (p *MultiFunctionPrinter) Fax(d Document) {
	panic("implement me")
}

func (p *MultiFunctionPrinter) Scan(d Document) {
	panic("implement me")
}

type OldPrinter struct {
	// ...
}

func (p *OldPrinter) Print(d Document) {
	panic("implement me")
}

// // Due to segregation principle, this is not needed for OldPrinter
// func (p *OldPrinter) Fax(d Document) {
// 	panic("not supported")
// }

// // Due to segregation principle, this is not needed for OldPrinter
// func (p *OldPrinter) Scan(d Document) {
// 	panic("not supported")
// }

type Photocopier struct {
	// ...
}

type MultiFunctionMachine interface {
	Printer
	Scanner
}

func (p *Photocopier) Print() {
	panic("implement me")
}

func (p *Photocopier) Scan() {
	panic("implement me")
}

func main() {}
