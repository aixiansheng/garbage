package symcache

import (
	"debug/macho"
)


type SymbolCtx struct {
	Sym macho.Symbol
	File MachFile
}

func NewSymbolCtx(s *macho.Symbol, f MachFile) *SymbolCtx {
	return &SymbolCtx{
		Sym: *s,
		File: f,
	}
}

func (s *SymbolCtx) IsExt() bool {
	return IsExt(s.Sym.Type)
}

func (s *SymbolCtx) IsPext() bool {
	return IsPext(s.Sym.Type)
}

func (s *SymbolCtx) IsStab() bool {
	return IsStab(s.Sym.Type)
}

func (s *SymbolCtx) Name() string {
	return s.Sym.Name
}

func (s *SymbolCtx) SubType() interface{} {
	if (s.IsStab()) {
		return NewStabSymbol(s)
	} else {
		switch GetSymbolType(s.Sym.Type) {
		case N_UNDF:
			return NewNTypeUndefined(s)
		case N_ABS:
			return NewNTypeAbsolute(s)
		case N_SECT:
			return NewNTypeSection(s)
		case N_PBUD:
			return NewNTypePrebound(s)
		case N_INDR:
			return NewNTypeIndirect(s)
		default:
			panic("Bad Symbol Type")
		}
	}
}

//
// Stab
//
type NStabSymbol struct {
	*SymbolCtx
}

func NewStabSymbol(s *SymbolCtx) *NStabSymbol {
	return &NStabSymbol{s}
}

func (s *NStabSymbol) Type() StabType {
	return StabType(s.Sym.Type)
}


//
// NType
//
type NTypeSymbol struct {
	*SymbolCtx
}

func (s *NTypeSymbol) Type() SymbolType {
	return GetSymbolType(s.Sym.Type)
}

func (s *NTypeSymbol) DescFlags() []SymbolDesc {
	return FieldSymbolDescs(s.Sym.Desc, s.File.Type(), s.Type())
}

//
// Undef
//

type NTypeUndefined struct {
	*NTypeSymbol
}

func NewNTypeUndefined(s *SymbolCtx) *NTypeUndefined {
	if s.Sym.Sect != NO_SECT {
		panic("Undefined symbol, but n_sect != NO_SECT")
	}
	return &NTypeUndefined{&NTypeSymbol{s}}
}

func (s *NTypeUndefined) IsCommon() bool {
	if s.IsExt() && s.Sym.Value != 0 && s.File.MachHeader().Type == macho.TypeObj {
		return true
	}
	return false
}

func (s *NTypeUndefined) SourceLibrary() string {
        return s.File.LoadedLib(int((s.Sym.Desc >> 8) & 0xff))
}

//
// Absolute
//

type NTypeAbsolute struct {
	*NTypeSymbol
}

func NewNTypeAbsolute(s *SymbolCtx) *NTypeAbsolute {
	if s.Sym.Sect != NO_SECT && s.Sym.Name != "__mh_execute_header" {
		panic("Absolute symbol, but n_sect != NO_SECT (and it wasn't __mh_execute_header)")
	}

	return &NTypeAbsolute{&NTypeSymbol{s}}
}

func (s *NTypeAbsolute) Value() uint64 {
	return s.Sym.Value
}

func (s *NTypeAbsolute) IsMHExecuteHeader() bool {
	if s.Sym.Name == "__mh_execute_header" {
		return true
	}
	return false
}

//
// Section
//

type NTypeSection struct {
	*NTypeSymbol
}

func NewNTypeSection(s *SymbolCtx) *NTypeSection {
	return &NTypeSection{&NTypeSymbol{s}}
}

func (s *NTypeSection) Sect() uint8 {
	return s.Sym.Sect
}

func (s *NTypeSection) Value() uint64 {
	return s.Sym.Value
}

//
// Prebound
//

type NTypePrebound struct {
	*NTypeSymbol
}

func NewNTypePrebound(s *SymbolCtx) *NTypePrebound {
	if s.Sym.Sect != NO_SECT {
		panic("Prebound symbol, but n_sect != NO_SECT")
	}

	return &NTypePrebound{&NTypeSymbol{s}}
}

//
// Indirect
//

type NTypeIndirect struct {
	*NTypeSymbol
}

func NewNTypeIndirect(s *SymbolCtx) *NTypeIndirect {
	return &NTypeIndirect{&NTypeSymbol{s}}
}

func (s *NTypeIndirect) ReferencedSymbol() string {
	return s.File.StrTab(int(s.Sym.Value))
}
