package symcache

import (
	"fmt"
)

//
// Mach Header flags
//
type MHFlag uint32
const (
	MH_NOUNDEFS MHFlag = 0x1
	MH_INCRLINK = 0x2
	MH_DYLDLINK = 0x4
	MH_BINDATLOAD = 0x8
	MH_PREBOUND = 0x10
	MH_SPLIT_SEGS = 0x20
	MH_LAZY_INIT = 0x40
	MH_TWOLEVEL = 0x80
	MH_FORCE_FLAT = 0x100
	MH_NOMULTIDEFS = 0x200
	MH_NOFIXPREBINDING = 0x400
	MH_PREBINDABLE  = 0x800
	MH_ALLMODSBOUND = 0x1000
	MH_SUBSECTIONS_VIA_SYMBOLS = 0x2000
	MH_CANONICAL = 0x4000
	MH_WEAK_DEFINES = 0x8000
	MH_BINDS_TO_WEAK = 0x10000
	MH_ALLOW_STACK_EXECUTION = 0x20000
	MH_ROOT_SAFE = 0x40000
	MH_SETUID_SAFE = 0x80000
	MH_NO_REEXPORTED_DYLIBS = 0x100000
	MH_PIE = 0x200000
	MH_DEAD_STRIPPABLE_DYLIB = 0x400000
	MH_HAS_TLV_DESCRIPTORS = 0x800000
	MH_NO_HEAP_EXECUTION = 0x1000000
	MH_APP_EXTENSION_SAFE = 0x02000000
	MH_NLIST_OUTOFSYNC_WITH_DYLDINFO = 0x04000000
)

var MHFlags = map[MHFlag]string {
	MH_NOUNDEFS : "MH_NOUNDEFS",
	MH_INCRLINK : "MH_INCRLINK",
	MH_DYLDLINK : "MH_DYLDLINK",
	MH_BINDATLOAD : "MH_BINDATLOAD",
	MH_PREBOUND : "MH_PREBOUND",
	MH_SPLIT_SEGS : "MH_SPLIT_SEGS",
	MH_LAZY_INIT : "MH_LAZY_INIT",
	MH_TWOLEVEL : "MH_TWOLEVEL",
	MH_FORCE_FLAT : "MH_FORCE_FLAT",
	MH_NOMULTIDEFS : "MH_NOMULTIDEFS",
	MH_NOFIXPREBINDING : "MH_NOFIXPREBINDING",
	MH_PREBINDABLE : "MH_PREBINDABLE",
	MH_ALLMODSBOUND : "MH_ALLMODSBOUND",
	MH_SUBSECTIONS_VIA_SYMBOLS : "MH_SUBSECTIONS_VIA_SYMBOLS",
	MH_CANONICAL : "MH_CANONICAL",
	MH_WEAK_DEFINES : "MH_WEAK_DEFINES",
	MH_BINDS_TO_WEAK : "MH_BINDS_TO_WEAK",
	MH_ALLOW_STACK_EXECUTION : "MH_ALLOW_STACK_EXECUTION",
	MH_ROOT_SAFE : "MH_ROOT_SAFE",
	MH_SETUID_SAFE : "MH_SETUID_SAFE",
	MH_NO_REEXPORTED_DYLIBS : "MH_NO_REEXPORTED_DYLIBS",
	MH_PIE : "MH_PIE",
	MH_DEAD_STRIPPABLE_DYLIB : "MH_DEAD_STRIPPABLE_DYLIB",
	MH_HAS_TLV_DESCRIPTORS : "MH_HAS_TLV_DESCRIPTORS",
	MH_NO_HEAP_EXECUTION : "MH_NO_HEAP_EXECUTION",
	MH_APP_EXTENSION_SAFE : "MH_APP_EXTENSION_SAFE",
	MH_NLIST_OUTOFSYNC_WITH_DYLDINFO : "MH_NLIST_OUTOFSYNC_WITH_DYLDINFO",
}

func (f MHFlag) String() string {
	return MHFlags[f]
}

func FieldHasMHFlag(field uint32, flag MHFlag) bool {
	if (uint32(flag) & field) == uint32(flag) {
		return true
	}
	return false
}

func FieldMHFlags(field uint32) (ret []MHFlag) {
	for flag, _ := range(MHFlags) {
		if (uint32(flag) & field) == uint32(flag) {
			ret = append(ret, flag)
		}
	}
	return
}

//
// Mach file types
//
type MHType uint32
const (
	MH_OBJECT MHType = 0x1
	MH_EXECUTE = 0x2
	MH_FVMLIB = 0x3
	MH_CORE = 0x4
	MH_PRELOAD = 0x5
	MH_DYLIB = 0x6
	MH_DYLINKER = 0x7
	MH_BUNDLE = 0x8
	MH_DYLIB_STUB = 0x9
	MH_DSYM = 0xa
	MH_KEXT_BUNDLE = 0xb
)

var MHTypes = map[MHType]string {
	MH_OBJECT : "MH_OBJECT",
	MH_EXECUTE : "MH_EXECUTE",
	MH_FVMLIB : "MH_FVMLIB",
	MH_CORE : "MH_CORE",
	MH_PRELOAD : "MH_PRELOAD",
	MH_DYLIB : "MH_DYLIB",
	MH_DYLINKER : "MH_DYLINKER",
	MH_BUNDLE : "MH_BUNDLE",
	MH_DYLIB_STUB : "MH_DYLIB_STUB",
	MH_DSYM : "MH_DSYM",
	MH_KEXT_BUNDLE : "MH_KEXT_BUNDLE",
}

func (t MHType) String() string {
	return MHTypes[t]
}

func FieldMHType(field uint32) MHType {
	for mhtype, _ := range(MHTypes) {
		if uint32(mhtype) == field {
			return mhtype
		}
	}
	panic("Bad MHType")
}

//
// Load Command types
//
const (
	LC_REQ_DYLD uint32 = 0x80000000
)

type LCType uint32
const (
	LC_SEGMENT LCType = 0x1
	LC_SYMTAB = 0x2
	LC_SYMSEG = 0x3
	LC_THREAD = 0x4
	LC_UNIXTHREAD = 0x5
	LC_LOADFVMLIB = 0x6
	LC_IDFVMLIB = 0x7
	LC_IDENT = 0x8
	LC_FVMFILE = 0x9
	LC_PREPAGE = 0xa
	LC_DYSYMTAB = 0xb
	LC_LOAD_DYLIB = 0xc
	LC_ID_DYLIB = 0xd
	LC_LOAD_DYLINKER = 0xe
	LC_ID_DYLINKER = 0xf
	LC_PREBOUND_DYLIB = 0x10
	LC_ROUTINES = 0x11
	LC_SUB_FRAMEWORK = 0x12
	LC_SUB_UMBRELLA = 0x13
	LC_SUB_CLIENT = 0x14
	LC_SUB_LIBRARY = 0x15
	LC_TWOLEVEL_HINTS = 0x16
	LC_PREBIND_CKSUM = 0x17
	LC_LOAD_WEAK_DYLIB = LCType(0x18 | LC_REQ_DYLD)
	LC_SEGMENT_64 = 0x19
	LC_ROUTINES_64 = 0x1a
	LC_UUID = 0x1b
	LC_RPATH = LCType(0x1c | LC_REQ_DYLD)
	LC_CODE_SIGNATURE = 0x1d
	LC_SEGMENT_SPLIT_INFO = 0x1e
	LC_REEXPORT_DYLIB = LCType(0x1f | LC_REQ_DYLD)
	LC_LAZY_LOAD_DYLIB = 0x20
	LC_ENCRYPTION_INFO = 0x21
	LC_DYLD_INFO = 0x22
	LC_DYLD_INFO_ONLY = LCType(0x22|LC_REQ_DYLD)
	LC_LOAD_UPWARD_DYLIB = LCType(0x23 | LC_REQ_DYLD)
	LC_VERSION_MIN_MACOSX = 0x24
	LC_VERSION_MIN_IPHONEOS = 0x25
	LC_FUNCTION_STARTS = 0x26
	LC_DYLD_ENVIRONMENT = 0x27
	LC_MAIN = LCType(0x28 | LC_REQ_DYLD)
	LC_DATA_IN_CODE = 0x29
	LC_SOURCE_VERSION = 0x2A
	LC_DYLIB_CODE_SIGN_DRS = 0x2B
	LC_ENCRYPTION_INFO_64 = 0x2C
	LC_LINKER_OPTION = 0x2D
	LC_LINKER_OPTIMIZATION_HINT = 0x2E
	LC_VERSION_MIN_TVOS = 0x2F
	LC_VERSION_MIN_WATCHOS = 0x30
	LC_NOTE = 0x31
	LC_BUILD_VERSION = 0x32
)

var LCTypes = map[LCType]string {
	LC_SEGMENT : "LC_SEGMENT",
	LC_SYMTAB : "LC_SYMTAB",
	LC_SYMSEG : "LC_SYMSEG",
	LC_THREAD : "LC_THREAD",
	LC_UNIXTHREAD : "LC_UNIXTHREAD",
	LC_LOADFVMLIB : "LC_LOADFVMLIB",
	LC_IDFVMLIB : "LC_IDFVMLIB",
	LC_IDENT : "LC_IDENT",
	LC_FVMFILE : "LC_FVMFILE",
	LC_PREPAGE : "LC_PREPAGE",
	LC_DYSYMTAB : "LC_DYSYMTAB",
	LC_LOAD_DYLIB : "LC_LOAD_DYLIB",
	LC_ID_DYLIB : "LC_ID_DYLIB",
	LC_LOAD_DYLINKER : "LC_LOAD_DYLINKER",
	LC_ID_DYLINKER : "LC_ID_DYLINKER",
	LC_PREBOUND_DYLIB : "LC_PREBOUND_DYLIB",
	LC_ROUTINES : "LC_ROUTINES",
	LC_SUB_FRAMEWORK : "LC_SUB_FRAMEWORK",
	LC_SUB_UMBRELLA : "LC_SUB_UMBRELLA",
	LC_SUB_CLIENT : "LC_SUB_CLIENT",
	LC_SUB_LIBRARY : "LC_SUB_LIBRARY",
	LC_TWOLEVEL_HINTS : "LC_TWOLEVEL_HINTS",
	LC_PREBIND_CKSUM : "LC_PREBIND_CKSUM",
	LC_LOAD_WEAK_DYLIB : "LC_LOAD_WEAK_DYLIB",
	LC_SEGMENT_64 : "LC_SEGMENT_64",
	LC_ROUTINES_64 : "LC_ROUTINES_64",
	LC_UUID : "LC_UUID",
	LC_RPATH : "LC_RPATH",
	LC_CODE_SIGNATURE : "LC_CODE_SIGNATURE",
	LC_SEGMENT_SPLIT_INFO : "LC_SEGMENT_SPLIT_INFO",
	LC_REEXPORT_DYLIB : "LC_REEXPORT_DYLIB",
	LC_LAZY_LOAD_DYLIB : "LC_LAZY_LOAD_DYLIB",
	LC_ENCRYPTION_INFO : "LC_ENCRYPTION_INFO",
	LC_DYLD_INFO : "LC_DYLD_INFO",
	LC_DYLD_INFO_ONLY : "LC_DYLD_INFO_ONLY",
	LC_LOAD_UPWARD_DYLIB : "LC_LOAD_UPWARD_DYLIB",
	LC_VERSION_MIN_MACOSX : "LC_VERSION_MIN_MACOSX",
	LC_VERSION_MIN_IPHONEOS : "LC_VERSION_MIN_IPHONEOS",
	LC_FUNCTION_STARTS : "LC_FUNCTION_STARTS",
	LC_DYLD_ENVIRONMENT : "LC_DYLD_ENVIRONMENT",
	LC_MAIN : "LC_MAIN",
	LC_DATA_IN_CODE : "LC_DATA_IN_CODE",
	LC_SOURCE_VERSION : "LC_SOURCE_VERSION",
	LC_DYLIB_CODE_SIGN_DRS : "LC_DYLIB_CODE_SIGN_DRS",
	LC_ENCRYPTION_INFO_64 : "LC_ENCRYPTION_INFO_64",
	LC_LINKER_OPTION : "LC_LINKER_OPTION",
	LC_LINKER_OPTIMIZATION_HINT : "LC_LINKER_OPTIMIZATION_HINT",
	LC_VERSION_MIN_TVOS : "LC_VERSION_MIN_TVOS",
	LC_VERSION_MIN_WATCHOS : "LC_VERSION_MIN_WATCHOS",
	LC_NOTE : "LC_NOTE",
	LC_BUILD_VERSION : "LC_BUILD_VERSION",
}

func (t LCType) String() string {
	return LCTypes[t]
}

func FieldLCType(field uint32) LCType {
	for lctype, _ := range(LCTypes) {
		if uint32(lctype) == field {
			return lctype
		}
	}
	panic("Bad LCType")
}

//
// STAB symbol types
//
type StabType uint8
const (
	N_GSYM StabType = 0x20
	N_FNAME = 0x22
	N_FUN = 0x24
	N_STSYM = 0x26
	N_LCSYM = 0x28
	N_BNSYM = 0x2e
	N_AST = 0x32
	N_OPT = 0x3c
	N_RSYM = 0x40
	N_SLINE = 0x44
	N_ENSYM = 0x4e
	N_SSYM = 0x60
	N_SO = 0x64
	N_OSO = 0x66
	N_LSYM = 0x80
	N_BINCL = 0x82
	N_SOL = 0x84
	N_PARAMS = 0x86
	N_VERSION = 0x88
	N_OLEVEL = 0x8A
	N_PSYM = 0xa0
	N_EINCL = 0xa2
	N_ENTRY = 0xa4
	N_LBRAC = 0xc0
	N_EXCL = 0xc2
	N_RBRAC = 0xe0
	N_BCOMM = 0xe2
	N_ECOMM = 0xe4
	N_ECOML = 0xe8
	N_LENG = 0xfe
	N_PC = 0x30
)

var StabTypes = map[StabType]string {
	N_GSYM : "N_GSYM",
	N_FNAME : "N_FNAME",
	N_FUN : "N_FUN",
	N_STSYM : "N_STSYM",
	N_LCSYM : "N_LCSYM",
	N_BNSYM : "N_BNSYM",
	N_AST : "N_AST",
	N_OPT : "N_OPT",
	N_RSYM : "N_RSYM",
	N_SLINE : "N_SLINE",
	N_ENSYM : "N_ENSYM",
	N_SSYM : "N_SSYM",
	N_SO : "N_SO",
	N_OSO : "N_OSO",
	N_LSYM : "N_LSYM",
	N_BINCL : "N_BINCL",
	N_SOL : "N_SOL",
	N_PARAMS : "N_PARAMS",
	N_VERSION : "N_VERSION",
	N_OLEVEL : "N_OLEVEL",
	N_PSYM : "N_PSYM",
	N_EINCL : "N_EINCL",
	N_ENTRY : "N_ENTRY",
	N_LBRAC : "N_LBRAC",
	N_EXCL : "N_EXCL",
	N_RBRAC : "N_RBRAC",
	N_BCOMM : "N_BCOMM",
	N_ECOMM : "N_ECOMM",
	N_ECOML : "N_ECOML",
	N_LENG : "N_LENG",
	N_PC : "N_PC",
}

func (t StabType) String() string {
	return StabTypes[t]
}

func FieldStabType(field uint8) StabType {
	for stabtype, _ := range(StabTypes) {
		if uint8(stabtype) == field {
			return stabtype
		}
	}
	panic("Bad StabType")
}

//
// symbol type mask (n_type)
//
type SymbolTypeMask uint8
const (
	N_STAB SymbolTypeMask = 0xe0
	N_PEXT = 0x10
	N_TYPE = 0x0e
	N_EXT = 0x01
)

var SymbolTypeMasks = map[SymbolTypeMask]string {
	N_STAB : "N_STAB",
	N_PEXT : "N_PEXT",
	N_TYPE : "N_TYPE",
	N_EXT : "N_EXT",
}

func (m SymbolTypeMask) String() string {
	return SymbolTypeMasks[m]
}

func IsStab(field uint8) bool {
	if (uint8(N_STAB) & field) != 0 {
		return true
	}
	return false
}

func IsExt(field uint8) bool {
	if (uint8(N_EXT) & field) != 0 {
		return true
	}
	return false
}

func IsPext(field uint8) bool {
	if (uint8(N_PEXT) & field) != 0 {
		return true
	}
	return false
}

func GetSymbolType(n_type uint8) SymbolType {
	return SymbolType(n_type & uint8(N_TYPE))
}

//
// symbol type (non-stab from n_type)
//
type SymbolType uint8
const (
	N_UNDF SymbolType = 0x0
	N_ABS = 0x2
	N_SECT = 0xe
	N_PBUD = 0xc
	N_INDR = 0xa
)

var SymbolTypes = map[SymbolType]string {
	N_UNDF : "N_UNDF",
	N_ABS : "N_ABS",
	N_SECT : "N_SECT",
	N_PBUD : "N_PBUD",
	N_INDR : "N_INDR",
}

func (t SymbolType) String() string {
	return SymbolTypes[t]
}

func FieldSymbolType(field uint8) SymbolType {
	for symtype, _ := range(SymbolTypes) {
		if uint8(symtype) == field {
			return symtype
		}
	}
	panic("Bad SymbolType");
}

const (
	REFERENCE_TYPE uint16 = 0x7
)

//
// Symbol description masks (n_desc) (non-unique/context sensitive)
//
type SymbolDesc uint16
const (
	// the REFERENCE_TYPE ones
	REFERENCE_FLAG_UNDEFINED_NON_LAZY SymbolDesc = iota + 1
	REFERENCE_FLAG_UNDEFINED_LAZY
	REFERENCE_FLAG_DEFINED
	REFERENCE_FLAG_PRIVATE_DEFINED
	REFERENCE_FLAG_PRIVATE_UNDEFINED_NON_LAZY
	REFERENCE_FLAG_PRIVATE_UNDEFINED_LAZY

	// other n_desc flags
	REFERENCED_DYNAMICALLY
	N_DESC_DISCARDED
	N_NO_DEAD_STRIP
	N_WEAK_REF
	N_WEAK_DEF
	N_REF_TO_WEAK
	N_ARM_THUMB_DEF
	N_SYMBOL_RESOLVER
	N_ALT_ENTRY
)

type ValueAndName struct {
	value uint16
	name string
}

var SymbolDescs = map[SymbolDesc]ValueAndName {
	REFERENCE_FLAG_UNDEFINED_NON_LAZY : { 0x0, "REFERENCE_FLAG_UNDEFINED_NON_LAZY" },
	REFERENCE_FLAG_UNDEFINED_LAZY : { 0x1, "REFERENCE_FLAG_UNDEFINED_LAZY" },
	REFERENCE_FLAG_DEFINED : { 0x2, "REFERENCE_FLAG_DEFINED" },
	REFERENCE_FLAG_PRIVATE_DEFINED : { 0x3, "REFERENCE_FLAG_PRIVATE_DEFINED" },
	REFERENCE_FLAG_PRIVATE_UNDEFINED_NON_LAZY : { 0x4, "REFERENCE_FLAG_PRIVATE_UNDEFINED_NON_LAZY" },
	REFERENCE_FLAG_PRIVATE_UNDEFINED_LAZY : { 0x5, "REFERENCE_FLAG_PRIVATE_UNDEFINED_LAZY" },
	REFERENCED_DYNAMICALLY : { 0x10, "REFERENCED_DYNAMICALLY" },
	N_DESC_DISCARDED : { 0x20, "N_DESC_DISCARDED" },
	N_NO_DEAD_STRIP : { 0x20, "N_NO_DEAD_STRIP" },
	N_WEAK_REF : { 0x40, "N_WEAK_REF" },
	N_WEAK_DEF : { 0x80, "N_WEAK_DEF" },
	N_REF_TO_WEAK : { 0x80, "N_REF_TO_WEAK" },
	N_ARM_THUMB_DEF : { 0x8, "N_ARM_THUMB_DEF" },
	N_SYMBOL_RESOLVER : { 0x100, "N_SYMBOL_RESOLVER" },
	N_ALT_ENTRY : { 0x200, "N_ALT_ENTRY" },
}

func (d SymbolDesc) String() string {
	for desc, valName := range(SymbolDescs) {
		if desc == d {
			return valName.name
		}
	}
	panic("Bad SymbolDesc")
}

//
// This is no fun.  we should probably look through the indirect symbol table and then check the DESC
// or dysymtab... this might be better for object files, see how NM does object files
//

// performs some validation and returns the list
func FieldSymbolDescs(field uint16, mhtype MHType, ntype SymbolType) (ret []SymbolDesc) {
	appendIfMatch := func(val uint16, dsc SymbolDesc, innerCheck func(uint16, SymbolDesc) bool) {
		if (val & field) == val {
			if innerCheck(val, dsc) {
				ret = append(ret, dsc)
			}
		}
	}

	all := func(val uint16, dsc SymbolDesc) bool {
		return true
	}

	undef := func(val uint16, dsc SymbolDesc) bool {
		if ntype != N_UNDF {
			fmt.Printf("%s\n", dsc)
			panic("non-undefined symbol used bad SymbolDesc")
		}
		return true
	}
	nonUndef := func(val uint16, dsc SymbolDesc) bool {
		if ntype == N_UNDF {
			fmt.Printf("%s", dsc)
			panic("undefined symbol used bad SymbolDesc")
		}
		return true
	}

	for desc, valName := range(SymbolDescs) {
		switch desc {
		// these two are for undefined symbols only
		case REFERENCE_FLAG_UNDEFINED_NON_LAZY:
			// data
			// this flag's value is 0, so we can't just mask and check
			if (field & 0xff) == 0 && ntype == N_UNDF {
				ret = append(ret, desc)
			}
		case REFERENCE_FLAG_UNDEFINED_LAZY:
			// text
			appendIfMatch(valName.value, desc, undef)

		// the next two are for defined symbols
		case REFERENCE_FLAG_DEFINED:
			appendIfMatch(valName.value, desc, undef)
		case REFERENCE_FLAG_PRIVATE_DEFINED:
			appendIfMatch(valName.value, desc, nonUndef)

		// althought the next two say UNDEFINED, they are only for defined symbols, non-exported
		case REFERENCE_FLAG_PRIVATE_UNDEFINED_NON_LAZY:
			// data
			appendIfMatch(valName.value, desc, nonUndef)
		case REFERENCE_FLAG_PRIVATE_UNDEFINED_LAZY:
			// text
			appendIfMatch(valName.value, desc, nonUndef)

		// referenced by dyld
		case REFERENCED_DYNAMICALLY:
			appendIfMatch(valName.value, desc, nonUndef)


		case N_DESC_DISCARDED:
			// do nothing, only in-memory
		case N_NO_DEAD_STRIP:
			if mhtype == MH_OBJECT {
				appendIfMatch(valName.value, desc, all)
			}
		case N_WEAK_REF:
			appendIfMatch(valName.value, desc, undef)
		case N_WEAK_DEF:
			// could add additional logic to detect coalesced section, which is the only place this symbol can reside
			if ntype != N_UNDF {
				appendIfMatch(valName.value, desc, all)
			}
		case N_REF_TO_WEAK:
			if ntype == N_UNDF {
				appendIfMatch(valName.value, desc, all)
			}
		case N_ARM_THUMB_DEF:
			// not sure if this is only N_UNDF && N_SECT...
			appendIfMatch(valName.value, desc, all)
		// these two are only for MH_OBJECTs, in addition N_ALT_ENTRY is only for defined symbols
		case N_SYMBOL_RESOLVER:
			if mhtype == MH_OBJECT {
				appendIfMatch(valName.value, desc, all)
			}
		case N_ALT_ENTRY:
			if mhtype == MH_OBJECT {
				appendIfMatch(valName.value, desc, nonUndef)
			}
		default:
			panic("Unknown Desc Type")
		}
	}
	return
}

const (
	SELF_LIBRARY_ORDINAL uint8 = 0x0
	MAX_LIBRARY_ORDINAL = 0xfd
	DYNAMIC_LOOKUP_ORDINAL = 0xfe
	EXECUTABLE_ORDINAL = 0xff
)

const (
	NO_SECT uint8 = 0
	MAX_SECT = 255
)

