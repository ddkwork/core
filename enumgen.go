// Code generated by "enumgen"; DO NOT EDIT.

package goosi

import (
	"errors"
	"strconv"
	"strings"
	"sync/atomic"

	"goki.dev/enums"
)

var _PlatformsValues = []Platforms{0, 1, 2, 3, 4}

// PlatformsN is the highest valid value
// for type Platforms, plus one.
const PlatformsN Platforms = 5

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _PlatformsNoOp() {
	var x [1]struct{}
	_ = x[MacOS-(0)]
	_ = x[LinuxX11-(1)]
	_ = x[Windows-(2)]
	_ = x[IOS-(3)]
	_ = x[Android-(4)]
}

var _PlatformsNameToValueMap = map[string]Platforms{
	`MacOS`:    0,
	`macos`:    0,
	`LinuxX11`: 1,
	`linuxx11`: 1,
	`Windows`:  2,
	`windows`:  2,
	`IOS`:      3,
	`ios`:      3,
	`Android`:  4,
	`android`:  4,
}

var _PlatformsDescMap = map[Platforms]string{
	0: `MacOS is a mac desktop machine (aka Darwin)`,
	1: `LinuxX11 is a Linux OS machine running X11 window server`,
	2: `Windows is a Microsoft Windows machine`,
	3: `IOS is an Apple iOS or iPadOS mobile phone or iPad`,
	4: `Android is an Android mobile phone or tablet`,
}

var _PlatformsMap = map[Platforms]string{
	0: `MacOS`,
	1: `LinuxX11`,
	2: `Windows`,
	3: `IOS`,
	4: `Android`,
}

// String returns the string representation
// of this Platforms value.
func (i Platforms) String() string {
	if str, ok := _PlatformsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Platforms value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Platforms) SetString(s string) error {
	if val, ok := _PlatformsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _PlatformsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Platforms")
}

// Int64 returns the Platforms value as an int64.
func (i Platforms) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Platforms value from an int64.
func (i *Platforms) SetInt64(in int64) {
	*i = Platforms(in)
}

// Desc returns the description of the Platforms value.
func (i Platforms) Desc() string {
	if str, ok := _PlatformsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// PlatformsValues returns all possible values
// for the type Platforms.
func PlatformsValues() []Platforms {
	return _PlatformsValues
}

// Values returns all possible values
// for the type Platforms.
func (i Platforms) Values() []enums.Enum {
	res := make([]enums.Enum, len(_PlatformsValues))
	for i, d := range _PlatformsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Platforms.
func (i Platforms) IsValid() bool {
	_, ok := _PlatformsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Platforms) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Platforms) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

var _VirtualKeyboardTypesValues = []VirtualKeyboardTypes{0, 1, 2}

// VirtualKeyboardTypesN is the highest valid value
// for type VirtualKeyboardTypes, plus one.
const VirtualKeyboardTypesN VirtualKeyboardTypes = 3

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _VirtualKeyboardTypesNoOp() {
	var x [1]struct{}
	_ = x[DefaultKeyboard-(0)]
	_ = x[SingleLineKeyboard-(1)]
	_ = x[NumberKeyboard-(2)]
}

var _VirtualKeyboardTypesNameToValueMap = map[string]VirtualKeyboardTypes{
	`DefaultKeyboard`:    0,
	`defaultkeyboard`:    0,
	`SingleLineKeyboard`: 1,
	`singlelinekeyboard`: 1,
	`NumberKeyboard`:     2,
	`numberkeyboard`:     2,
}

var _VirtualKeyboardTypesDescMap = map[VirtualKeyboardTypes]string{
	0: `DefaultKeyboard is the keyboard with default input style and &#34;return&#34; return key`,
	1: `SingleLineKeyboard is the keyboard with default input style and &#34;Done&#34; return key`,
	2: `NumberKeyboard is the keyboard with number input style and &#34;Done&#34; return key`,
}

var _VirtualKeyboardTypesMap = map[VirtualKeyboardTypes]string{
	0: `DefaultKeyboard`,
	1: `SingleLineKeyboard`,
	2: `NumberKeyboard`,
}

// String returns the string representation
// of this VirtualKeyboardTypes value.
func (i VirtualKeyboardTypes) String() string {
	if str, ok := _VirtualKeyboardTypesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the VirtualKeyboardTypes value from its
// string representation, and returns an
// error if the string is invalid.
func (i *VirtualKeyboardTypes) SetString(s string) error {
	if val, ok := _VirtualKeyboardTypesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _VirtualKeyboardTypesNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type VirtualKeyboardTypes")
}

// Int64 returns the VirtualKeyboardTypes value as an int64.
func (i VirtualKeyboardTypes) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the VirtualKeyboardTypes value from an int64.
func (i *VirtualKeyboardTypes) SetInt64(in int64) {
	*i = VirtualKeyboardTypes(in)
}

// Desc returns the description of the VirtualKeyboardTypes value.
func (i VirtualKeyboardTypes) Desc() string {
	if str, ok := _VirtualKeyboardTypesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// VirtualKeyboardTypesValues returns all possible values
// for the type VirtualKeyboardTypes.
func VirtualKeyboardTypesValues() []VirtualKeyboardTypes {
	return _VirtualKeyboardTypesValues
}

// Values returns all possible values
// for the type VirtualKeyboardTypes.
func (i VirtualKeyboardTypes) Values() []enums.Enum {
	res := make([]enums.Enum, len(_VirtualKeyboardTypesValues))
	for i, d := range _VirtualKeyboardTypesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type VirtualKeyboardTypes.
func (i VirtualKeyboardTypes) IsValid() bool {
	_, ok := _VirtualKeyboardTypesMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i VirtualKeyboardTypes) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *VirtualKeyboardTypes) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

var _EventTypeValues = []EventType{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}

// EventTypeN is the highest valid value
// for type EventType, plus one.
const EventTypeN EventType = 22

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _EventTypeNoOp() {
	var x [1]struct{}
	_ = x[MouseEvent-(0)]
	_ = x[MouseMoveEvent-(1)]
	_ = x[MouseDragEvent-(2)]
	_ = x[MouseScrollEvent-(3)]
	_ = x[MouseFocusEvent-(4)]
	_ = x[MouseHoverEvent-(5)]
	_ = x[KeyEvent-(6)]
	_ = x[KeyChordEvent-(7)]
	_ = x[TouchEvent-(8)]
	_ = x[MagnifyEvent-(9)]
	_ = x[RotateEvent-(10)]
	_ = x[WindowEvent-(11)]
	_ = x[WindowResizeEvent-(12)]
	_ = x[WindowPaintEvent-(13)]
	_ = x[WindowShowEvent-(14)]
	_ = x[WindowFocusEvent-(15)]
	_ = x[DNDEvent-(16)]
	_ = x[DNDMoveEvent-(17)]
	_ = x[DNDFocusEvent-(18)]
	_ = x[OSEvent-(19)]
	_ = x[OSOpenFilesEvent-(20)]
	_ = x[CustomEventType-(21)]
}

var _EventTypeNameToValueMap = map[string]EventType{
	`MouseEvent`:        0,
	`mouseevent`:        0,
	`MouseMoveEvent`:    1,
	`mousemoveevent`:    1,
	`MouseDragEvent`:    2,
	`mousedragevent`:    2,
	`MouseScrollEvent`:  3,
	`mousescrollevent`:  3,
	`MouseFocusEvent`:   4,
	`mousefocusevent`:   4,
	`MouseHoverEvent`:   5,
	`mousehoverevent`:   5,
	`KeyEvent`:          6,
	`keyevent`:          6,
	`KeyChordEvent`:     7,
	`keychordevent`:     7,
	`TouchEvent`:        8,
	`touchevent`:        8,
	`MagnifyEvent`:      9,
	`magnifyevent`:      9,
	`RotateEvent`:       10,
	`rotateevent`:       10,
	`WindowEvent`:       11,
	`windowevent`:       11,
	`WindowResizeEvent`: 12,
	`windowresizeevent`: 12,
	`WindowPaintEvent`:  13,
	`windowpaintevent`:  13,
	`WindowShowEvent`:   14,
	`windowshowevent`:   14,
	`WindowFocusEvent`:  15,
	`windowfocusevent`:  15,
	`DNDEvent`:          16,
	`dndevent`:          16,
	`DNDMoveEvent`:      17,
	`dndmoveevent`:      17,
	`DNDFocusEvent`:     18,
	`dndfocusevent`:     18,
	`OSEvent`:           19,
	`osevent`:           19,
	`OSOpenFilesEvent`:  20,
	`osopenfilesevent`:  20,
	`CustomEventType`:   21,
	`customeventtype`:   21,
}

var _EventTypeDescMap = map[EventType]string{
	0:  `MouseEvent includes all mouse button actions, but not move or drag`,
	1:  `MouseMoveEvent is when the mouse is moving but no button is down`,
	2:  `MouseDragEvent is when the mouse is moving and there is a button down`,
	3:  `MouseScrollEvent is for mouse scroll wheel events`,
	4:  `MouseFocusEvent is for mouse focus (enter / exit of widget area) -- generated by gi.Window based on mouse move events`,
	5:  `MouseHoverEvent is for mouse hover -- generated by gi.Window based on mouse events`,
	6:  `KeyEvent for key pressed or released -- fine-grained data about each key as it happens`,
	7:  `KeyChordEvent is only generated when a non-modifier key is released, and it also contains a string representation of the full chord, suitable for translation into keyboard commands, emacs-style etc`,
	8:  `TouchEvent is a generic touch-based event`,
	9:  `MagnifyEvent is a touch-based magnify event (e.g., pinch)`,
	10: `RotateEvent is a touch-based rotate event`,
	11: `WindowEvent reports any changes in the window size, orientation, iconify, close, open, paint -- these are all &#34;internal&#34; events from OS to GUI system, and not sent to widgets`,
	12: `WindowResizeEvent is specifically for window resize events which need special treatment -- this is an internal event not sent to widgets`,
	13: `WindowPaintEvent is specifically for window paint events which need special treatment -- this is an internal event not sent to widgets, triggered right after window is opened for initial painting.`,
	14: `WindowShowEvent is a synthetic event sent to widget consumers, sent *only once* when window is shown for the very first time`,
	15: `WindowFocusEvent is a synthetic event sent to widget consumers, sent when window focus changes (action is Focus / DeFocus)`,
	16: `DNDEvent is for the Drag-n-Drop (DND) drop event`,
	17: `DNDMoveEvent is when the DND position has changed`,
	18: `DNDFocusEvent is for Enter / Exit events of the DND into / out of a given widget`,
	19: `OSEvent is an operating system generated event (app level typically)`,
	20: `OSOpenFilesEvent is an event telling app to open given files`,
	21: `CustomEventType is a user-defined event with a data any field`,
}

var _EventTypeMap = map[EventType]string{
	0:  `MouseEvent`,
	1:  `MouseMoveEvent`,
	2:  `MouseDragEvent`,
	3:  `MouseScrollEvent`,
	4:  `MouseFocusEvent`,
	5:  `MouseHoverEvent`,
	6:  `KeyEvent`,
	7:  `KeyChordEvent`,
	8:  `TouchEvent`,
	9:  `MagnifyEvent`,
	10: `RotateEvent`,
	11: `WindowEvent`,
	12: `WindowResizeEvent`,
	13: `WindowPaintEvent`,
	14: `WindowShowEvent`,
	15: `WindowFocusEvent`,
	16: `DNDEvent`,
	17: `DNDMoveEvent`,
	18: `DNDFocusEvent`,
	19: `OSEvent`,
	20: `OSOpenFilesEvent`,
	21: `CustomEventType`,
}

// String returns the string representation
// of this EventType value.
func (i EventType) String() string {
	if str, ok := _EventTypeMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the EventType value from its
// string representation, and returns an
// error if the string is invalid.
func (i *EventType) SetString(s string) error {
	if val, ok := _EventTypeNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _EventTypeNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type EventType")
}

// Int64 returns the EventType value as an int64.
func (i EventType) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the EventType value from an int64.
func (i *EventType) SetInt64(in int64) {
	*i = EventType(in)
}

// Desc returns the description of the EventType value.
func (i EventType) Desc() string {
	if str, ok := _EventTypeDescMap[i]; ok {
		return str
	}
	return i.String()
}

// EventTypeValues returns all possible values
// for the type EventType.
func EventTypeValues() []EventType {
	return _EventTypeValues
}

// Values returns all possible values
// for the type EventType.
func (i EventType) Values() []enums.Enum {
	res := make([]enums.Enum, len(_EventTypeValues))
	for i, d := range _EventTypeValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type EventType.
func (i EventType) IsValid() bool {
	_, ok := _EventTypeMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i EventType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *EventType) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

var _ScreenOrientationValues = []ScreenOrientation{0, 1, 2}

// ScreenOrientationN is the highest valid value
// for type ScreenOrientation, plus one.
const ScreenOrientationN ScreenOrientation = 3

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ScreenOrientationNoOp() {
	var x [1]struct{}
	_ = x[OrientationUnknown-(0)]
	_ = x[Portrait-(1)]
	_ = x[Landscape-(2)]
}

var _ScreenOrientationNameToValueMap = map[string]ScreenOrientation{
	`OrientationUnknown`: 0,
	`orientationunknown`: 0,
	`Portrait`:           1,
	`portrait`:           1,
	`Landscape`:          2,
	`landscape`:          2,
}

var _ScreenOrientationDescMap = map[ScreenOrientation]string{
	0: `OrientationUnknown means device orientation cannot be determined. Equivalent on Android to Configuration.ORIENTATION_UNKNOWN and on iOS to: UIDeviceOrientationUnknown UIDeviceOrientationFaceUp UIDeviceOrientationFaceDown`,
	1: `Portrait is a device oriented so it is tall and thin. Equivalent on Android to Configuration.ORIENTATION_PORTRAIT and on iOS to: UIDeviceOrientationPortrait UIDeviceOrientationPortraitUpsideDown`,
	2: `Landscape is a device oriented so it is short and wide. Equivalent on Android to Configuration.ORIENTATION_LANDSCAPE and on iOS to: UIDeviceOrientationLandscapeLeft UIDeviceOrientationLandscapeRight`,
}

var _ScreenOrientationMap = map[ScreenOrientation]string{
	0: `OrientationUnknown`,
	1: `Portrait`,
	2: `Landscape`,
}

// String returns the string representation
// of this ScreenOrientation value.
func (i ScreenOrientation) String() string {
	if str, ok := _ScreenOrientationMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the ScreenOrientation value from its
// string representation, and returns an
// error if the string is invalid.
func (i *ScreenOrientation) SetString(s string) error {
	if val, ok := _ScreenOrientationNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _ScreenOrientationNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type ScreenOrientation")
}

// Int64 returns the ScreenOrientation value as an int64.
func (i ScreenOrientation) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the ScreenOrientation value from an int64.
func (i *ScreenOrientation) SetInt64(in int64) {
	*i = ScreenOrientation(in)
}

// Desc returns the description of the ScreenOrientation value.
func (i ScreenOrientation) Desc() string {
	if str, ok := _ScreenOrientationDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ScreenOrientationValues returns all possible values
// for the type ScreenOrientation.
func ScreenOrientationValues() []ScreenOrientation {
	return _ScreenOrientationValues
}

// Values returns all possible values
// for the type ScreenOrientation.
func (i ScreenOrientation) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ScreenOrientationValues))
	for i, d := range _ScreenOrientationValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type ScreenOrientation.
func (i ScreenOrientation) IsValid() bool {
	_, ok := _ScreenOrientationMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i ScreenOrientation) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *ScreenOrientation) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

var _WindowFlagsValues = []WindowFlags{0, 1, 2, 3, 4, 5}

// WindowFlagsN is the highest valid value
// for type WindowFlags, plus one.
const WindowFlagsN WindowFlags = 6

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _WindowFlagsNoOp() {
	var x [1]struct{}
	_ = x[Dialog-(0)]
	_ = x[Modal-(1)]
	_ = x[Tool-(2)]
	_ = x[Fullscreen-(3)]
	_ = x[Minimized-(4)]
	_ = x[Focus-(5)]
}

var _WindowFlagsNameToValueMap = map[string]WindowFlags{
	`Dialog`:     0,
	`dialog`:     0,
	`Modal`:      1,
	`modal`:      1,
	`Tool`:       2,
	`tool`:       2,
	`Fullscreen`: 3,
	`fullscreen`: 3,
	`Minimized`:  4,
	`minimized`:  4,
	`Focus`:      5,
	`focus`:      5,
}

var _WindowFlagsDescMap = map[WindowFlags]string{
	0: `Dialog indicates that this is a temporary, pop-up window.`,
	1: `Modal indicates that this dialog window blocks events going to other windows until it is closed.`,
	2: `Tool indicates that this is a floating tool window that has minimized window decoration.`,
	3: `Fullscreen indicates a window that occupies the entire screen.`,
	4: `Minimized indicates a window reduced to an icon, or otherwise no longer visible or active. Otherwise, the window should be assumed to be visible.`,
	5: `Focus indicates that the window has the focus.`,
}

var _WindowFlagsMap = map[WindowFlags]string{
	0: `Dialog`,
	1: `Modal`,
	2: `Tool`,
	3: `Fullscreen`,
	4: `Minimized`,
	5: `Focus`,
}

// String returns the string representation
// of this WindowFlags value.
func (i WindowFlags) String() string {
	str := ""
	for _, ie := range _WindowFlagsValues {
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
// representation of this WindowFlags value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i WindowFlags) BitIndexString() string {
	if str, ok := _WindowFlagsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the WindowFlags value from its
// string representation, and returns an
// error if the string is invalid.
func (i *WindowFlags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the WindowFlags value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *WindowFlags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _WindowFlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _WindowFlagsNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else {
			return errors.New(flg + " is not a valid value for type WindowFlags")
		}
	}
	return nil
}

// Int64 returns the WindowFlags value as an int64.
func (i WindowFlags) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the WindowFlags value from an int64.
func (i *WindowFlags) SetInt64(in int64) {
	*i = WindowFlags(in)
}

// Desc returns the description of the WindowFlags value.
func (i WindowFlags) Desc() string {
	if str, ok := _WindowFlagsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// WindowFlagsValues returns all possible values
// for the type WindowFlags.
func WindowFlagsValues() []WindowFlags {
	return _WindowFlagsValues
}

// Values returns all possible values
// for the type WindowFlags.
func (i WindowFlags) Values() []enums.Enum {
	res := make([]enums.Enum, len(_WindowFlagsValues))
	for i, d := range _WindowFlagsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type WindowFlags.
func (i WindowFlags) IsValid() bool {
	_, ok := _WindowFlagsMap[i]
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i WindowFlags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *WindowFlags) SetFlag(on bool, f ...enums.BitFlag) {
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
func (i WindowFlags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *WindowFlags) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
