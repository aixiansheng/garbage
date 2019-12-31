package symcache 

type SymbolVisitor interface {
	VisitSymbol(Symbol) error
}

type Symbol interface {
	IsExt() bool
	IsPext() bool
	IsStab() bool

	Name() string

	SubType() interface{}
}

type StabSymbol interface {
	Type() StabType
}

type NType interface {
	Type() SymbolType
	DescFlags() []SymbolDesc
}

type UndefSymbol interface {
	NType
	IsCommon() bool
	SourceLibrary() string
}

type AbsoluteSymbol interface {
	NType
	Value() uint64
	IsMHExecuteHeader() bool
}

type SectionSymbol interface {
	NType
	Sect() uint8
	Value() uint64
}

type PreboundSymbol interface {
	NType
}

type IndirectSymbol interface {
	NType
	ReferencedSymbol() string
}
