// Code generated by "core generate"; DO NOT EDIT.

package keymap

import (
	"cogentcore.org/core/enums"
)

var _FunctionsValues = []Functions{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65}

// FunctionsN is the highest valid value for type Functions, plus one.
const FunctionsN Functions = 66

var _FunctionsValueMap = map[string]Functions{`None`: 0, `MoveUp`: 1, `MoveDown`: 2, `MoveRight`: 3, `MoveLeft`: 4, `PageUp`: 5, `PageDown`: 6, `Home`: 7, `End`: 8, `DocHome`: 9, `DocEnd`: 10, `WordRight`: 11, `WordLeft`: 12, `FocusNext`: 13, `FocusPrev`: 14, `Enter`: 15, `Accept`: 16, `CancelSelect`: 17, `SelectMode`: 18, `SelectAll`: 19, `Abort`: 20, `Copy`: 21, `Cut`: 22, `Paste`: 23, `PasteHist`: 24, `Backspace`: 25, `BackspaceWord`: 26, `Delete`: 27, `DeleteWord`: 28, `Kill`: 29, `Duplicate`: 30, `Transpose`: 31, `TransposeWord`: 32, `Undo`: 33, `Redo`: 34, `Insert`: 35, `InsertAfter`: 36, `ZoomOut`: 37, `ZoomIn`: 38, `Refresh`: 39, `Recenter`: 40, `Complete`: 41, `Lookup`: 42, `Search`: 43, `Find`: 44, `Replace`: 45, `Jump`: 46, `HistPrev`: 47, `HistNext`: 48, `Menu`: 49, `WinFocusNext`: 50, `WinClose`: 51, `WinSnapshot`: 52, `New`: 53, `NewAlt1`: 54, `NewAlt2`: 55, `Open`: 56, `OpenAlt1`: 57, `OpenAlt2`: 58, `Save`: 59, `SaveAs`: 60, `SaveAlt`: 61, `CloseAlt1`: 62, `CloseAlt2`: 63, `MultiA`: 64, `MultiB`: 65}

var _FunctionsDescMap = map[Functions]string{0: ``, 1: ``, 2: ``, 3: ``, 4: ``, 5: ``, 6: ``, 7: `PageRight PageLeft`, 8: ``, 9: ``, 10: ``, 11: ``, 12: ``, 13: ``, 14: ``, 15: ``, 16: ``, 17: ``, 18: ``, 19: ``, 20: ``, 21: `EditItem`, 22: ``, 23: ``, 24: ``, 25: ``, 26: ``, 27: ``, 28: ``, 29: ``, 30: ``, 31: ``, 32: ``, 33: ``, 34: ``, 35: ``, 36: ``, 37: ``, 38: ``, 39: ``, 40: ``, 41: ``, 42: ``, 43: ``, 44: ``, 45: ``, 46: ``, 47: ``, 48: ``, 49: ``, 50: ``, 51: ``, 52: ``, 53: `Below are menu specific functions -- use these as shortcuts for menu buttons allows uniqueness of mapping and easy customization of all key buttons`, 54: ``, 55: ``, 56: ``, 57: ``, 58: ``, 59: ``, 60: ``, 61: ``, 62: ``, 63: ``, 64: ``, 65: ``}

var _FunctionsMap = map[Functions]string{0: `None`, 1: `MoveUp`, 2: `MoveDown`, 3: `MoveRight`, 4: `MoveLeft`, 5: `PageUp`, 6: `PageDown`, 7: `Home`, 8: `End`, 9: `DocHome`, 10: `DocEnd`, 11: `WordRight`, 12: `WordLeft`, 13: `FocusNext`, 14: `FocusPrev`, 15: `Enter`, 16: `Accept`, 17: `CancelSelect`, 18: `SelectMode`, 19: `SelectAll`, 20: `Abort`, 21: `Copy`, 22: `Cut`, 23: `Paste`, 24: `PasteHist`, 25: `Backspace`, 26: `BackspaceWord`, 27: `Delete`, 28: `DeleteWord`, 29: `Kill`, 30: `Duplicate`, 31: `Transpose`, 32: `TransposeWord`, 33: `Undo`, 34: `Redo`, 35: `Insert`, 36: `InsertAfter`, 37: `ZoomOut`, 38: `ZoomIn`, 39: `Refresh`, 40: `Recenter`, 41: `Complete`, 42: `Lookup`, 43: `Search`, 44: `Find`, 45: `Replace`, 46: `Jump`, 47: `HistPrev`, 48: `HistNext`, 49: `Menu`, 50: `WinFocusNext`, 51: `WinClose`, 52: `WinSnapshot`, 53: `New`, 54: `NewAlt1`, 55: `NewAlt2`, 56: `Open`, 57: `OpenAlt1`, 58: `OpenAlt2`, 59: `Save`, 60: `SaveAs`, 61: `SaveAlt`, 62: `CloseAlt1`, 63: `CloseAlt2`, 64: `MultiA`, 65: `MultiB`}

// String returns the string representation of this Functions value.
func (i Functions) String() string { return enums.String(i, _FunctionsMap) }

// SetString sets the Functions value from its string representation,
// and returns an error if the string is invalid.
func (i *Functions) SetString(s string) error {
	return enums.SetString(i, s, _FunctionsValueMap, "Functions")
}

// Int64 returns the Functions value as an int64.
func (i Functions) Int64() int64 { return int64(i) }

// SetInt64 sets the Functions value from an int64.
func (i *Functions) SetInt64(in int64) { *i = Functions(in) }

// Desc returns the description of the Functions value.
func (i Functions) Desc() string { return enums.Desc(i, _FunctionsDescMap) }

// FunctionsValues returns all possible values for the type Functions.
func FunctionsValues() []Functions { return _FunctionsValues }

// Values returns all possible values for the type Functions.
func (i Functions) Values() []enums.Enum { return enums.Values(_FunctionsValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Functions) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Functions) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "Functions")
}
