// Code generated by "goki generate"; DO NOT EDIT.

package svg

import (
	"errors"
	"strconv"
	"strings"
	"sync/atomic"

	"goki.dev/enums"
	"goki.dev/ki/v2"
)

var _NodeFlagsValues = []NodeFlags{7}

// NodeFlagsN is the highest valid value
// for type NodeFlags, plus one.
const NodeFlagsN NodeFlags = 8

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _NodeFlagsNoOp() {
	var x [1]struct{}
	_ = x[IsDef-(7)]
}

var _NodeFlagsNameToValueMap = map[string]NodeFlags{
	`IsDef`: 7,
	`isdef`: 7,
}

var _NodeFlagsDescMap = map[NodeFlags]string{
	7: `Rendering means that the SVG is currently redrawing Can be useful to check for animations etc to decide whether to drive another update`,
}

var _NodeFlagsMap = map[NodeFlags]string{
	7: `IsDef`,
}

// String returns the string representation
// of this NodeFlags value.
func (i NodeFlags) String() string {
	str := ""
	for _, ie := range ki.FlagsValues() {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	for _, ie := range _NodeFlagsValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string
// representation of this NodeFlags value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i NodeFlags) BitIndexString() string {
	if str, ok := _NodeFlagsMap[i]; ok {
		return str
	}
	return ki.Flags(i).BitIndexString()
}

// SetString sets the NodeFlags value from its
// string representation, and returns an
// error if the string is invalid.
func (i *NodeFlags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the NodeFlags value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *NodeFlags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _NodeFlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _NodeFlagsNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else {
			err := (*ki.Flags)(i).SetStringOr(flg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Int64 returns the NodeFlags value as an int64.
func (i NodeFlags) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the NodeFlags value from an int64.
func (i *NodeFlags) SetInt64(in int64) {
	*i = NodeFlags(in)
}

// Desc returns the description of the NodeFlags value.
func (i NodeFlags) Desc() string {
	if str, ok := _NodeFlagsDescMap[i]; ok {
		return str
	}
	return ki.Flags(i).Desc()
}

// NodeFlagsValues returns all possible values
// for the type NodeFlags.
func NodeFlagsValues() []NodeFlags {
	es := ki.FlagsValues()
	res := make([]NodeFlags, len(es))
	for i, e := range es {
		res[i] = NodeFlags(e)
	}
	res = append(res, _NodeFlagsValues...)
	return res
}

// Values returns all possible values
// for the type NodeFlags.
func (i NodeFlags) Values() []enums.Enum {
	es := ki.FlagsValues()
	les := len(es)
	res := make([]enums.Enum, les+len(_NodeFlagsValues))
	for i, d := range es {
		res[i] = d
	}
	for i, d := range _NodeFlagsValues {
		res[i+les] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type NodeFlags.
func (i NodeFlags) IsValid() bool {
	_, ok := _NodeFlagsMap[i]
	if !ok {
		return ki.Flags(i).IsValid()
	}
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i NodeFlags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *NodeFlags) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i NodeFlags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *NodeFlags) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

var _ViewBoxMeetOrSliceValues = []ViewBoxMeetOrSlice{0, 1}

// ViewBoxMeetOrSliceN is the highest valid value
// for type ViewBoxMeetOrSlice, plus one.
const ViewBoxMeetOrSliceN ViewBoxMeetOrSlice = 2

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ViewBoxMeetOrSliceNoOp() {
	var x [1]struct{}
	_ = x[Meet-(0)]
	_ = x[Slice-(1)]
}

var _ViewBoxMeetOrSliceNameToValueMap = map[string]ViewBoxMeetOrSlice{
	`Meet`:  0,
	`meet`:  0,
	`Slice`: 1,
	`slice`: 1,
}

var _ViewBoxMeetOrSliceDescMap = map[ViewBoxMeetOrSlice]string{
	0: `Meet means the entire ViewBox is visible within Viewport, and it is scaled up as much as possible to meet the align constraints`,
	1: `Slice means the entire ViewBox is covered by the ViewBox, and the ViewBox is scaled down as much as possible, while still meeting the align constraints`,
}

var _ViewBoxMeetOrSliceMap = map[ViewBoxMeetOrSlice]string{
	0: `Meet`,
	1: `Slice`,
}

// String returns the string representation
// of this ViewBoxMeetOrSlice value.
func (i ViewBoxMeetOrSlice) String() string {
	if str, ok := _ViewBoxMeetOrSliceMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the ViewBoxMeetOrSlice value from its
// string representation, and returns an
// error if the string is invalid.
func (i *ViewBoxMeetOrSlice) SetString(s string) error {
	if val, ok := _ViewBoxMeetOrSliceNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _ViewBoxMeetOrSliceNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type ViewBoxMeetOrSlice")
}

// Int64 returns the ViewBoxMeetOrSlice value as an int64.
func (i ViewBoxMeetOrSlice) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the ViewBoxMeetOrSlice value from an int64.
func (i *ViewBoxMeetOrSlice) SetInt64(in int64) {
	*i = ViewBoxMeetOrSlice(in)
}

// Desc returns the description of the ViewBoxMeetOrSlice value.
func (i ViewBoxMeetOrSlice) Desc() string {
	if str, ok := _ViewBoxMeetOrSliceDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ViewBoxMeetOrSliceValues returns all possible values
// for the type ViewBoxMeetOrSlice.
func ViewBoxMeetOrSliceValues() []ViewBoxMeetOrSlice {
	return _ViewBoxMeetOrSliceValues
}

// Values returns all possible values
// for the type ViewBoxMeetOrSlice.
func (i ViewBoxMeetOrSlice) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ViewBoxMeetOrSliceValues))
	for i, d := range _ViewBoxMeetOrSliceValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type ViewBoxMeetOrSlice.
func (i ViewBoxMeetOrSlice) IsValid() bool {
	_, ok := _ViewBoxMeetOrSliceMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i ViewBoxMeetOrSlice) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *ViewBoxMeetOrSlice) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
