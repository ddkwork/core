// Code generated by "core generate"; DO NOT EDIT.

package filetree

import (
	"cogentcore.org/core/enums"
)

var _DirFlagsValues = []DirFlags{0, 1, 2}

// DirFlagsN is the highest valid value for type DirFlags, plus one.
const DirFlagsN DirFlags = 3

var _DirFlagsValueMap = map[string]DirFlags{`IsOpen`: 0, `SortByName`: 1, `SortByModTime`: 2}

var _DirFlagsDescMap = map[DirFlags]string{0: `DirIsOpen means directory is open -- else closed`, 1: `DirSortByName means sort the directory entries by name. this is mutex with other sorts -- keeping option open for non-binary sort choices.`, 2: `DirSortByModTime means sort the directory entries by modification time`}

var _DirFlagsMap = map[DirFlags]string{0: `IsOpen`, 1: `SortByName`, 2: `SortByModTime`}

// String returns the string representation of this DirFlags value.
func (i DirFlags) String() string { return enums.BitFlagString(i, _DirFlagsValues) }

// BitIndexString returns the string representation of this DirFlags value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i DirFlags) BitIndexString() string { return enums.String(i, _DirFlagsMap) }

// SetString sets the DirFlags value from its string representation,
// and returns an error if the string is invalid.
func (i *DirFlags) SetString(s string) error { *i = 0; return i.SetStringOr(s) }

// SetStringOr sets the DirFlags value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *DirFlags) SetStringOr(s string) error {
	return enums.SetStringOr(i, s, _DirFlagsValueMap, "DirFlags")
}

// Int64 returns the DirFlags value as an int64.
func (i DirFlags) Int64() int64 { return int64(i) }

// SetInt64 sets the DirFlags value from an int64.
func (i *DirFlags) SetInt64(in int64) { *i = DirFlags(in) }

// Desc returns the description of the DirFlags value.
func (i DirFlags) Desc() string { return enums.Desc(i, _DirFlagsDescMap) }

// DirFlagsValues returns all possible values for the type DirFlags.
func DirFlagsValues() []DirFlags { return _DirFlagsValues }

// Values returns all possible values for the type DirFlags.
func (i DirFlags) Values() []enums.Enum { return enums.Values(_DirFlagsValues) }

// HasFlag returns whether these bit flags have the given bit flag set.
func (i DirFlags) HasFlag(f enums.BitFlag) bool { return enums.HasFlag((*int64)(&i), f) }

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *DirFlags) SetFlag(on bool, f ...enums.BitFlag) { enums.SetFlag((*int64)(i), on, f...) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i DirFlags) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *DirFlags) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "DirFlags") }

var _FindLocValues = []FindLoc{0, 1, 2, 3, 4}

// FindLocN is the highest valid value for type FindLoc, plus one.
const FindLocN FindLoc = 5

var _FindLocValueMap = map[string]FindLoc{`Open`: 0, `All`: 1, `File`: 2, `Dir`: 3, `NotTop`: 4}

var _FindLocDescMap = map[FindLoc]string{0: `FindOpen finds in all open folders in the left file browser`, 1: `FindLocAll finds in all directories under the root path. can be slow for large file trees`, 2: `FindLocFile only finds in the current active file`, 3: `FindLocDir only finds in the directory of the current active file`, 4: `FindLocNotTop finds in all open folders *except* the top-level folder`}

var _FindLocMap = map[FindLoc]string{0: `Open`, 1: `All`, 2: `File`, 3: `Dir`, 4: `NotTop`}

// String returns the string representation of this FindLoc value.
func (i FindLoc) String() string { return enums.String(i, _FindLocMap) }

// SetString sets the FindLoc value from its string representation,
// and returns an error if the string is invalid.
func (i *FindLoc) SetString(s string) error {
	return enums.SetString(i, s, _FindLocValueMap, "FindLoc")
}

// Int64 returns the FindLoc value as an int64.
func (i FindLoc) Int64() int64 { return int64(i) }

// SetInt64 sets the FindLoc value from an int64.
func (i *FindLoc) SetInt64(in int64) { *i = FindLoc(in) }

// Desc returns the description of the FindLoc value.
func (i FindLoc) Desc() string { return enums.Desc(i, _FindLocDescMap) }

// FindLocValues returns all possible values for the type FindLoc.
func FindLocValues() []FindLoc { return _FindLocValues }

// Values returns all possible values for the type FindLoc.
func (i FindLoc) Values() []enums.Enum { return enums.Values(_FindLocValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i FindLoc) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *FindLoc) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "FindLoc") }
