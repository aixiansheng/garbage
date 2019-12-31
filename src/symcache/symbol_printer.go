package symcache

import (
	"debug/macho"
	"fmt"
	"io"
)

type SymbolPrinter struct {
	SupportedArches []macho.Cpu
}

func (s *SymbolPrinter) SetSupportedArches(arches []macho.Cpu) {
	s.SupportedArches = arches
}

func (printer *SymbolPrinter) VisitMachFile(file MachFile) error {
	cpuType := file.CpuType()
	for _, arch := range(printer.SupportedArches) {
		if arch != cpuType {
			return nil
		}
	}

	fmt.Printf("Mach File: %s: cpu: %v\n", file.Path(), cpuType)
	fmt.Printf("Type: %s\n", MHType(file.Type()))
	fmt.Printf("Flags: %v\n", file.Flags())
	fmt.Printf("\n")

	for _, symbol := range(file.Symbols()) {
		if err := printer.printSymbol(symbol); err != nil {
			return err
		}
	}

	return nil
}

func (printer *SymbolPrinter) VisitFile(path string, r io.ReaderAt) error {
	if mh, err := macho.NewFile(r); err == nil {
		f := &MachFileCtx{
			r: r,
			path: path,
			machFile: mh,
		}
		return printer.VisitMachFile(f)
	} else if fh, err := macho.NewFatFile(r); err == nil {
		for _, fatarch := range(fh.Arches) {
			subreader := io.NewSectionReader(r, int64(fatarch.Offset), int64(fatarch.Size))
			f := &MachFileCtx{
				r: subreader,
				path: path,
				fatFile: fh,
				machFile: fatarch.File,
			}
			if err = printer.VisitMachFile(f); err != nil {
				return err;
			}
		}
	}

	return nil
}

//
// internals
//

func (printer *SymbolPrinter) printUndefSymbol(symbol UndefSymbol) error {
	fmt.Printf("type: undefined")
	if symbol.IsCommon() {
		fmt.Printf(" (common)")
	}
	fmt.Printf(" [%s]\n", symbol.SourceLibrary())
	fmt.Printf("desc: %v\n", symbol.DescFlags())
	return nil
}

func (printer *SymbolPrinter) printAbsoluteSymbol(symbol AbsoluteSymbol) error {
	fmt.Printf("type: absolute [%#x] isMHExecuteHeader: %v\n", symbol.Value(), symbol.IsMHExecuteHeader())
	return nil
}

func (printer *SymbolPrinter) printSectionSymbol(symbol SectionSymbol) error {
	fmt.Printf("type: defined_in_section [%v][%#x]\n", symbol.Sect(), symbol.Value())
	fmt.Printf("desc: %v\n", symbol.DescFlags())
	return nil
}

func (printer *SymbolPrinter) printPreboundSymbol(symbol PreboundSymbol) error {
	fmt.Printf("type: prebound\n")
	return nil
}

func (printer *SymbolPrinter) printIndirectSymbol(symbol IndirectSymbol) error {
	fmt.Printf("type: indirect [%s]\n", symbol.ReferencedSymbol())
	return nil
}

func (printer *SymbolPrinter) printStabSymbol(symbol StabSymbol) error {
	fmt.Printf("stab type: %s\n", symbol.Type())
	return nil
}

func (printer *SymbolPrinter) printSymbol(symbol Symbol) (err error) {
	fmt.Printf("name: %s\n", symbol.Name())

	fmt.Printf("flags:")
	if symbol.IsPext() {
		fmt.Printf(" private_extern")
	}

	if symbol.IsExt() {
		fmt.Printf(" extern")
	}
	fmt.Printf("\n")

	switch v := symbol.SubType().(type) {
	case StabSymbol:
		err = printer.printStabSymbol(v)
	case UndefSymbol:
		err = printer.printUndefSymbol(v)
	case AbsoluteSymbol:
		err = printer.printAbsoluteSymbol(v)
	case SectionSymbol:
		err = printer.printSectionSymbol(v)
	case IndirectSymbol:
		err = printer.printIndirectSymbol(v)
	case PreboundSymbol:
		err = printer.printPreboundSymbol(v)
	default:
		panic("Couldn't print symbol of unknown type")
	}

	fmt.Printf("\n\n")

	return
}

