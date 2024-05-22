// Code generated by "core generate"; DO NOT EDIT.

package events

import (
	"cogentcore.org/core/enums"
)

var _DropModsValues = []DropMods{0, 1, 2, 3, 4}

// DropModsN is the highest valid value for type DropMods, plus one.
const DropModsN DropMods = 5

var _DropModsValueMap = map[string]DropMods{`NoDropMod`: 0, `Copy`: 1, `Move`: 2, `Link`: 3, `Ignore`: 4}

var _DropModsDescMap = map[DropMods]string{0: ``, 1: `Copy is the default and implies data is just copied -- receiver can do with it as they please and source does not need to take any further action`, 2: `Move is signaled with a Shift or Meta key (by default) and implies that the source should delete itself when it receives the DropFromSource event action with this Mod value set -- receiver must update the Mod to reflect actual action taken, and be particularly careful with this one`, 3: `Link can be any other kind of alternative action -- link is applicable to files (symbolic link)`, 4: `Ignore means that the receiver chose to not process this drop`}

var _DropModsMap = map[DropMods]string{0: `NoDropMod`, 1: `Copy`, 2: `Move`, 3: `Link`, 4: `Ignore`}

// String returns the string representation of this DropMods value.
func (i DropMods) String() string { return enums.String(i, _DropModsMap) }

// SetString sets the DropMods value from its string representation,
// and returns an error if the string is invalid.
func (i *DropMods) SetString(s string) error {
	return enums.SetString(i, s, _DropModsValueMap, "DropMods")
}

// Int64 returns the DropMods value as an int64.
func (i DropMods) Int64() int64 { return int64(i) }

// SetInt64 sets the DropMods value from an int64.
func (i *DropMods) SetInt64(in int64) { *i = DropMods(in) }

// Desc returns the description of the DropMods value.
func (i DropMods) Desc() string { return enums.Desc(i, _DropModsDescMap) }

// DropModsValues returns all possible values for the type DropMods.
func DropModsValues() []DropMods { return _DropModsValues }

// Values returns all possible values for the type DropMods.
func (i DropMods) Values() []enums.Enum { return enums.Values(_DropModsValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i DropMods) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *DropMods) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "DropMods") }

var _ButtonsValues = []Buttons{0, 1, 2, 3}

// ButtonsN is the highest valid value for type Buttons, plus one.
const ButtonsN Buttons = 4

var _ButtonsValueMap = map[string]Buttons{`NoButton`: 0, `Left`: 1, `Middle`: 2, `Right`: 3}

var _ButtonsDescMap = map[Buttons]string{0: ``, 1: ``, 2: ``, 3: ``}

var _ButtonsMap = map[Buttons]string{0: `NoButton`, 1: `Left`, 2: `Middle`, 3: `Right`}

// String returns the string representation of this Buttons value.
func (i Buttons) String() string { return enums.String(i, _ButtonsMap) }

// SetString sets the Buttons value from its string representation,
// and returns an error if the string is invalid.
func (i *Buttons) SetString(s string) error {
	return enums.SetString(i, s, _ButtonsValueMap, "Buttons")
}

// Int64 returns the Buttons value as an int64.
func (i Buttons) Int64() int64 { return int64(i) }

// SetInt64 sets the Buttons value from an int64.
func (i *Buttons) SetInt64(in int64) { *i = Buttons(in) }

// Desc returns the description of the Buttons value.
func (i Buttons) Desc() string { return enums.Desc(i, _ButtonsDescMap) }

// ButtonsValues returns all possible values for the type Buttons.
func ButtonsValues() []Buttons { return _ButtonsValues }

// Values returns all possible values for the type Buttons.
func (i Buttons) Values() []enums.Enum { return enums.Values(_ButtonsValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Buttons) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Buttons) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Buttons") }

var _SelectModesValues = []SelectModes{0, 1, 2, 3, 4, 5, 6}

// SelectModesN is the highest valid value for type SelectModes, plus one.
const SelectModesN SelectModes = 7

var _SelectModesValueMap = map[string]SelectModes{`SelectOne`: 0, `ExtendContinuous`: 1, `ExtendOne`: 2, `NoSelect`: 3, `Unselect`: 4, `SelectQuiet`: 5, `UnselectQuiet`: 6}

var _SelectModesDescMap = map[SelectModes]string{0: `SelectOne selects a single item, and is the default when no modifier key is pressed`, 1: `ExtendContinuous, activated by Shift key, extends the selection to select a continuous region of selected items, with no gaps`, 2: `ExtendOne, activated by Control or Meta / Command, extends the selection by adding the one additional item just clicked on, creating a potentially discontinuous set of selected items`, 3: `NoSelect means do not update selection -- this is used programmatically and not available via modifier key`, 4: `Unselect means unselect items -- this is used programmatically and not available via modifier key -- typically ExtendOne will unselect if already selected`, 5: `SelectQuiet means select without doing other updates or signals -- for bulk updates with a final update at the end -- used programmatically`, 6: `UnselectQuiet means unselect without doing other updates or signals -- for bulk updates with a final update at the end -- used programmatically`}

var _SelectModesMap = map[SelectModes]string{0: `SelectOne`, 1: `ExtendContinuous`, 2: `ExtendOne`, 3: `NoSelect`, 4: `Unselect`, 5: `SelectQuiet`, 6: `UnselectQuiet`}

// String returns the string representation of this SelectModes value.
func (i SelectModes) String() string { return enums.String(i, _SelectModesMap) }

// SetString sets the SelectModes value from its string representation,
// and returns an error if the string is invalid.
func (i *SelectModes) SetString(s string) error {
	return enums.SetString(i, s, _SelectModesValueMap, "SelectModes")
}

// Int64 returns the SelectModes value as an int64.
func (i SelectModes) Int64() int64 { return int64(i) }

// SetInt64 sets the SelectModes value from an int64.
func (i *SelectModes) SetInt64(in int64) { *i = SelectModes(in) }

// Desc returns the description of the SelectModes value.
func (i SelectModes) Desc() string { return enums.Desc(i, _SelectModesDescMap) }

// SelectModesValues returns all possible values for the type SelectModes.
func SelectModesValues() []SelectModes { return _SelectModesValues }

// Values returns all possible values for the type SelectModes.
func (i SelectModes) Values() []enums.Enum { return enums.Values(_SelectModesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i SelectModes) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *SelectModes) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "SelectModes")
}

var _TypesValues = []Types{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45}

// TypesN is the highest valid value for type Types, plus one.
const TypesN Types = 46

var _TypesValueMap = map[string]Types{`UnknownType`: 0, `MouseDown`: 1, `MouseUp`: 2, `MouseMove`: 3, `MouseDrag`: 4, `Click`: 5, `DoubleClick`: 6, `TripleClick`: 7, `ContextMenu`: 8, `LongPressStart`: 9, `LongPressEnd`: 10, `MouseEnter`: 11, `MouseLeave`: 12, `LongHoverStart`: 13, `LongHoverEnd`: 14, `DragStart`: 15, `DragMove`: 16, `DragEnter`: 17, `DragLeave`: 18, `Drop`: 19, `DropDeleteSource`: 20, `SlideStart`: 21, `SlideMove`: 22, `SlideStop`: 23, `Scroll`: 24, `KeyDown`: 25, `KeyUp`: 26, `KeyChord`: 27, `TouchStart`: 28, `TouchEnd`: 29, `TouchMove`: 30, `Magnify`: 31, `Rotate`: 32, `Select`: 33, `Focus`: 34, `FocusLost`: 35, `Change`: 36, `Input`: 37, `Show`: 38, `Close`: 39, `Window`: 40, `WindowResize`: 41, `WindowPaint`: 42, `OS`: 43, `OSOpenFiles`: 44, `Custom`: 45}

var _TypesDescMap = map[Types]string{0: `zero value is an unknown type`, 1: `MouseDown happens when a mouse button is pressed down. See MouseButton() for which. See Click for a synthetic event representing a MouseDown followed by MouseUp on the same element with Left (primary) mouse button. Often that is the most useful.`, 2: `MouseUp happens when a mouse button is released. See MouseButton() for which.`, 3: `MouseMove is always sent when the mouse is moving but no button is down, even if there might be other higher-level events too. These can be numerous and thus it is typically more efficient to listen to other events derived from this. Not unique, and Prev position is updated during compression.`, 4: `MouseDrag is always sent when the mouse is moving and there is a button down, even if there might be other higher-level events too. The start pos indicates where (and when) button first was pressed. Not unique and Prev position is updated during compression.`, 5: `Click represents a MouseDown followed by MouseUp in sequence on the same element, with the Left (primary) button. This is the typical event for most basic user interaction.`, 6: `DoubleClick represents two Click events in a row in rapid succession.`, 7: `TripleClick represents three Click events in a row in rapid succession.`, 8: `ContextMenu represents a MouseDown/Up event with the Right mouse button (which is also activated by Control key + Left Click).`, 9: `LongPressStart is when the mouse has been relatively stable after MouseDown on an element for a minimum duration (500 msec default).`, 10: `LongPressEnd is sent after LongPressStart when the mouse has gone up, moved sufficiently, left the current element, or another input event has happened.`, 11: `MouseEnter is when the mouse enters the bounding box of a new element. It is used for setting the Hover state, and can trigger cursor changes. See DragEnter for alternative case during Drag events.`, 12: `MouseLeave is when the mouse leaves the bounding box of an element, that previously had a MouseEnter event. Given that elements can have overlapping bounding boxes (e.g., child elements within a container), it is not the case that a MouseEnter on a child triggers a MouseLeave on surrounding containers. See DragLeave for alternative case during Drag events.`, 13: `LongHoverStart is when the mouse has been relatively stable after MouseEnter on an element for a minimum duration (500 msec default). This triggers the LongHover state typically used for Tooltips.`, 14: `LongHoverEnd is after LongHoverStart when the mouse has moved sufficiently, left the current element, or another input event has happened, thereby terminating the LongHover state.`, 15: `DragStart is at the start of a drag-n-drop event sequence, when a Draggable element is Active and a sufficient distance of MouseDrag events has occurred to engage the DragStart event.`, 16: `DragMove is for a MouseDrag event during the drag-n-drop sequence. Usually don&#39;t need to listen to this one. MouseDrag is also sent.`, 17: `DragEnter is like MouseEnter but after a DragStart during a drag-n-drop sequence. MouseEnter is not sent in this case.`, 18: `DragLeave is like MouseLeave but after a DragStart during a drag-n-drop sequence. MouseLeave is not sent in this case.`, 19: `Drop is sent when an item being Dragged is dropped on top of a target element. The event struct should be DragDrop.`, 20: `DropDeleteSource is sent to the source Drag element if the Drag-n-Drop event is a Move type, which requires deleting the source element. The event struct should be DragDrop.`, 21: `SlideStart is for a Slideable element when Active and a sufficient distance of MouseDrag events has occurred to engage the SlideStart event. Sets the Sliding state.`, 22: `SlideMove is for a Slideable element after SlideStart is being dragged via MouseDrag events.`, 23: `SlideStop is when the mouse button is released on a Slideable element being dragged via MouseDrag events. This typically also accompanied by a Changed event for the new slider value.`, 24: `Scroll is for scroll wheel or other scrolling events (gestures). These are not unique and Delta is updated during compression.`, 25: `KeyDown is when a key is pressed down. This provides fine-grained data about each key as it happens. KeyChord is recommended for a more complete Key event.`, 26: `KeyUp is when a key is released. This provides fine-grained data about each key as it happens. KeyChord is recommended for a more complete Key event.`, 27: `KeyChord is only generated when a non-modifier key is released, and it also contains a string representation of the full chord, suitable for translation into keyboard commands, emacs-style etc. It can be somewhat delayed relative to the KeyUp.`, 28: `TouchStart is when a touch event starts, for the low-level touch event processing. TouchStart also activates MouseDown, Scroll, Magnify, or Rotate events depending on gesture recognition.`, 29: `TouchEnd is when a touch event ends, for the low-level touch event processing. TouchEnd also activates MouseUp events depending on gesture recognition.`, 30: `TouchMove is when a touch event moves, for the low-level touch event processing. TouchMove also activates MouseMove, Scroll, Magnify, or Rotate events depending on gesture recognition.`, 31: `Magnify is a touch-based magnify event (e.g., pinch)`, 32: `Rotate is a touch-based rotate event.`, 33: `Select is sent for any direction of selection change on (or within if relevant) a Selectable element. Typically need to query the element(s) to determine current selection state.`, 34: `Focus is sent when Focsable element receives Focus`, 35: `FocusLost is sent when Focsable element loses Focus`, 36: `Change is when a value represented by the element has been changed by the user and committed (for example, someone has typed text in a textfield and then pressed enter). This is *not* triggered when the value has not been committed; see [Input] for that. This is for Editable, Checkable, and Slidable items.`, 37: `Input is when a value represented by the element has changed, but has not necessarily been committed (for example, this triggers each time someone presses a key in a text field). This *is* triggered when the value has not been committed; see [Change] for a version that only occurs when the value is committed. This is for Editable, Checkable, and Slidable items.`, 38: `Show is sent to widgets when their Scene is first shown to the user in its final form. Listening to this event enables widgets to perform initial one-time activities on startup, in the context of a fully rendered display. This is guaranteed to only happen once per widget per Scene.`, 39: `Close is sent to widgets when their Scene is being closed. This is an opportunity to save unsaved edits, for example. This is guaranteed to only happen once per widget per Scene.`, 40: `Window reports on changes in the window position, visibility (iconify), focus changes, screen update, and closing. These are only sent once per event (Unique).`, 41: `WindowResize happens when the window has been resized, which can happen continuously during a user resizing episode. These are not Unique events, and are compressed to minimize lag.`, 42: `WindowPaint is sent continuously at FPS frequency (60 frames per second by default) to drive updating check on the window. It is not unique, will be compressed to keep pace with updating.`, 43: `OS is an operating system generated event (app level typically)`, 44: `OSOpenFiles is an event telling app to open given files`, 45: `Custom is a user-defined event with a data any field`}

var _TypesMap = map[Types]string{0: `UnknownType`, 1: `MouseDown`, 2: `MouseUp`, 3: `MouseMove`, 4: `MouseDrag`, 5: `Click`, 6: `DoubleClick`, 7: `TripleClick`, 8: `ContextMenu`, 9: `LongPressStart`, 10: `LongPressEnd`, 11: `MouseEnter`, 12: `MouseLeave`, 13: `LongHoverStart`, 14: `LongHoverEnd`, 15: `DragStart`, 16: `DragMove`, 17: `DragEnter`, 18: `DragLeave`, 19: `Drop`, 20: `DropDeleteSource`, 21: `SlideStart`, 22: `SlideMove`, 23: `SlideStop`, 24: `Scroll`, 25: `KeyDown`, 26: `KeyUp`, 27: `KeyChord`, 28: `TouchStart`, 29: `TouchEnd`, 30: `TouchMove`, 31: `Magnify`, 32: `Rotate`, 33: `Select`, 34: `Focus`, 35: `FocusLost`, 36: `Change`, 37: `Input`, 38: `Show`, 39: `Close`, 40: `Window`, 41: `WindowResize`, 42: `WindowPaint`, 43: `OS`, 44: `OSOpenFiles`, 45: `Custom`}

// String returns the string representation of this Types value.
func (i Types) String() string { return enums.String(i, _TypesMap) }

// SetString sets the Types value from its string representation,
// and returns an error if the string is invalid.
func (i *Types) SetString(s string) error { return enums.SetString(i, s, _TypesValueMap, "Types") }

// Int64 returns the Types value as an int64.
func (i Types) Int64() int64 { return int64(i) }

// SetInt64 sets the Types value from an int64.
func (i *Types) SetInt64(in int64) { *i = Types(in) }

// Desc returns the description of the Types value.
func (i Types) Desc() string { return enums.Desc(i, _TypesDescMap) }

// TypesValues returns all possible values for the type Types.
func TypesValues() []Types { return _TypesValues }

// Values returns all possible values for the type Types.
func (i Types) Values() []enums.Enum { return enums.Values(_TypesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Types) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Types) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Types") }

var _EventFlagsValues = []EventFlags{0, 1}

// EventFlagsN is the highest valid value for type EventFlags, plus one.
const EventFlagsN EventFlags = 2

var _EventFlagsValueMap = map[string]EventFlags{`Handled`: 0, `Unique`: 1}

var _EventFlagsDescMap = map[EventFlags]string{0: `Handled indicates that the event has been handled`, 1: `EventUnique indicates that the event is Unique and not to be compressed with like events.`}

var _EventFlagsMap = map[EventFlags]string{0: `Handled`, 1: `Unique`}

// String returns the string representation of this EventFlags value.
func (i EventFlags) String() string { return enums.BitFlagString(i, _EventFlagsValues) }

// BitIndexString returns the string representation of this EventFlags value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i EventFlags) BitIndexString() string { return enums.String(i, _EventFlagsMap) }

// SetString sets the EventFlags value from its string representation,
// and returns an error if the string is invalid.
func (i *EventFlags) SetString(s string) error { *i = 0; return i.SetStringOr(s) }

// SetStringOr sets the EventFlags value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *EventFlags) SetStringOr(s string) error {
	return enums.SetStringOr(i, s, _EventFlagsValueMap, "EventFlags")
}

// Int64 returns the EventFlags value as an int64.
func (i EventFlags) Int64() int64 { return int64(i) }

// SetInt64 sets the EventFlags value from an int64.
func (i *EventFlags) SetInt64(in int64) { *i = EventFlags(in) }

// Desc returns the description of the EventFlags value.
func (i EventFlags) Desc() string { return enums.Desc(i, _EventFlagsDescMap) }

// EventFlagsValues returns all possible values for the type EventFlags.
func EventFlagsValues() []EventFlags { return _EventFlagsValues }

// Values returns all possible values for the type EventFlags.
func (i EventFlags) Values() []enums.Enum { return enums.Values(_EventFlagsValues) }

// HasFlag returns whether these bit flags have the given bit flag set.
func (i EventFlags) HasFlag(f enums.BitFlag) bool { return enums.HasFlag((*int64)(&i), f) }

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *EventFlags) SetFlag(on bool, f ...enums.BitFlag) { enums.SetFlag((*int64)(i), on, f...) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i EventFlags) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *EventFlags) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "EventFlags")
}

var _WinActionsValues = []WinActions{0, 1, 2, 3, 4, 5, 6, 7}

// WinActionsN is the highest valid value for type WinActions, plus one.
const WinActionsN WinActions = 8

var _WinActionsValueMap = map[string]WinActions{`NoWinAction`: 0, `Close`: 1, `Minimize`: 2, `Move`: 3, `Focus`: 4, `FocusLost`: 5, `Show`: 6, `ScreenUpdate`: 7}

var _WinActionsDescMap = map[WinActions]string{0: `NoWinAction is the zero value for special types (Resize, Paint)`, 1: `WinClose means that the window is about to close, but has not yet closed.`, 2: `WinMinimize means that the window has been iconified / miniaturized / is no longer visible.`, 3: `WinMove means that the window was moved but NOT resized or changed in any other way -- does not require a redraw, but anything tracking positions will want to update.`, 4: `WinFocus indicates that the window has been activated for receiving user input.`, 5: `WinFocusLost indicates that the window is no longer activated for receiving input.`, 6: `WinShow is for the WindowShow event -- sent by the system shortly after the window has opened, to ensure that full rendering is completed with the proper size, and to trigger one-time actions such as configuring the main menu after the window has opened.`, 7: `ScreenUpdate occurs when any of the screen information is updated This event is sent to the first window on the list of active windows and it should then perform any necessary updating`}

var _WinActionsMap = map[WinActions]string{0: `NoWinAction`, 1: `Close`, 2: `Minimize`, 3: `Move`, 4: `Focus`, 5: `FocusLost`, 6: `Show`, 7: `ScreenUpdate`}

// String returns the string representation of this WinActions value.
func (i WinActions) String() string { return enums.String(i, _WinActionsMap) }

// SetString sets the WinActions value from its string representation,
// and returns an error if the string is invalid.
func (i *WinActions) SetString(s string) error {
	return enums.SetString(i, s, _WinActionsValueMap, "WinActions")
}

// Int64 returns the WinActions value as an int64.
func (i WinActions) Int64() int64 { return int64(i) }

// SetInt64 sets the WinActions value from an int64.
func (i *WinActions) SetInt64(in int64) { *i = WinActions(in) }

// Desc returns the description of the WinActions value.
func (i WinActions) Desc() string { return enums.Desc(i, _WinActionsDescMap) }

// WinActionsValues returns all possible values for the type WinActions.
func WinActionsValues() []WinActions { return _WinActionsValues }

// Values returns all possible values for the type WinActions.
func (i WinActions) Values() []enums.Enum { return enums.Values(_WinActionsValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i WinActions) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *WinActions) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "WinActions")
}
