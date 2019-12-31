package symcache

import (
	"bytes"
	"debug/macho"
	"encoding/binary"
	"io"
)

type MachFileCtx struct {
	path string
	r io.ReaderAt
	fatFile *macho.FatFile
	machFile *macho.File

	strtabLoaded bool
	strtab []byte

	loadedDylibsLoaded bool
	loadedDylibs []string
}

func (file *MachFileCtx) Path() string {
	return file.path
}

func (file *MachFileCtx) IsFat() bool {
	return file.fatFile != nil
}

func (file *MachFileCtx) CpuType() macho.Cpu {
	return file.machFile.Cpu
}

func (file *MachFileCtx) Type() MHType {
	return FieldMHType(uint32(file.machFile.Type))
}

func (file *MachFileCtx) Flags() []MHFlag {
	return FieldMHFlags(file.machFile.Flags)
}

func (file *MachFileCtx) LoadedLib(idx int) string {
	if file.loadedDylibsLoaded == false {
		for _, lcmd_ := range(file.machFile.Loads) {
			lcmd := lcmd_.Raw()
			cmd := file.machFile.ByteOrder.Uint32(lcmd[0:4])

			switch LCType(cmd) {
			case LC_LAZY_LOAD_DYLIB:
				fallthrough
			case LC_LOAD_DYLIB:
				fallthrough
			case LC_LOAD_WEAK_DYLIB:
				fallthrough
			case LC_REEXPORT_DYLIB:
				fallthrough
			case LC_LOAD_UPWARD_DYLIB:
				var hdr macho.DylibCmd
				b := bytes.NewReader(lcmd)
				if err := binary.Read(b, file.machFile.ByteOrder, &hdr); err != nil {
					panic("Bad dylib load cmd")
				}
				file.loadedDylibs = append(file.loadedDylibs, cstring(lcmd[hdr.Name:]))
			}
		}

		file.loadedDylibsLoaded = true
	}

	if !FieldHasMHFlag(file.machFile.Flags, MH_TWOLEVEL) {
		return "FLAT"
	} else {
		switch uint8(idx) {
		case SELF_LIBRARY_ORDINAL:
			return "SELF_LIBRARY_ORDINAL"
		case DYNAMIC_LOOKUP_ORDINAL:
			return "DYNAMIC_LOOKUP_ORDINAL"
		case EXECUTABLE_ORDINAL:
			return "EXECUTABLE_ORDINAL"
		default:
			return file.loadedDylibs[idx-1]
		}
	}
}

func (file *MachFileCtx) StrTab(idx int) string {
	if file.strtabLoaded == false {
		var symtabCmd macho.SymtabCmd
		r := bytes.NewReader(file.machFile.Symtab.LoadBytes)
		if binary.Read(r, file.machFile.ByteOrder, &symtabCmd) != nil {
			panic("Couldn't read symtab cmd")
		}

		file.strtab = make([]byte, symtabCmd.Strsize)
		if _, err := file.r.ReadAt(file.strtab, int64(symtabCmd.Stroff)); err != nil {
			panic("Couldn't read strtab")
		}
		file.strtabLoaded = true
	}

	return strTabStr(&file.strtab, idx)
}

func (file *MachFileCtx) Symbols() (ret []Symbol) {
	if file.machFile.Symtab != nil {
		for _, sym := range(file.machFile.Symtab.Syms) {
			ret = append(ret, NewSymbolCtx(&sym, file))
		}
	}

	return
}

func (file *MachFileCtx) FatHeader() *macho.FatFile {
	return file.fatFile
}

func (file *MachFileCtx) MachHeader() *macho.File {
	return file.machFile
}

// from debug/macho

func cstring(b []byte) string {
	var i int
	for i = 0; i < len(b) && b[i] != 0; i++ {
	}
	return string(b[0:i])
}

func strTabStr(stab *[]byte, start int) string {
	return cstring((*stab)[start:])
}
