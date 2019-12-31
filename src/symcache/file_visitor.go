package symcache 

import (
	"io"
	"debug/macho"
)

type MachFile interface {
	Path() string
	IsFat() bool
	CpuType() macho.Cpu
	Type() MHType
	Flags() []MHFlag
	LoadedLib(idx int) string
	StrTab(idx int) string
	Symbols() []Symbol

	FatHeader() *macho.FatFile
	MachHeader() *macho.File
}

type MachFileVisitor interface {
	VisitFile(path string, reader io.ReaderAt) error
	VisitMachFile(MachFile) error
	SetSupportedArches([]macho.Cpu)
}

