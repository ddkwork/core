// Code generated by "core generate -add-types"; DO NOT EDIT.

package xyz

import (
	"cogentcore.org/core/enums"
)

var _LightColorsValues = []LightColors{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

// LightColorsN is the highest valid value for type LightColors, plus one.
const LightColorsN LightColors = 15

var _LightColorsValueMap = map[string]LightColors{`DirectSun`: 0, `CarbonArc`: 1, `Halogen`: 2, `Tungsten100W`: 3, `Tungsten40W`: 4, `Candle`: 5, `Overcast`: 6, `FluorWarm`: 7, `FluorStd`: 8, `FluorCool`: 9, `FluorFull`: 10, `FluorGrow`: 11, `MercuryVapor`: 12, `SodiumVapor`: 13, `MetalHalide`: 14}

var _LightColorsDescMap = map[LightColors]string{0: ``, 1: ``, 2: ``, 3: ``, 4: ``, 5: ``, 6: ``, 7: ``, 8: ``, 9: ``, 10: ``, 11: ``, 12: ``, 13: ``, 14: ``}

var _LightColorsMap = map[LightColors]string{0: `DirectSun`, 1: `CarbonArc`, 2: `Halogen`, 3: `Tungsten100W`, 4: `Tungsten40W`, 5: `Candle`, 6: `Overcast`, 7: `FluorWarm`, 8: `FluorStd`, 9: `FluorCool`, 10: `FluorFull`, 11: `FluorGrow`, 12: `MercuryVapor`, 13: `SodiumVapor`, 14: `MetalHalide`}

// String returns the string representation of this LightColors value.
func (i LightColors) String() string { return enums.String(i, _LightColorsMap) }

// SetString sets the LightColors value from its string representation,
// and returns an error if the string is invalid.
func (i *LightColors) SetString(s string) error {
	return enums.SetString(i, s, _LightColorsValueMap, "LightColors")
}

// Int64 returns the LightColors value as an int64.
func (i LightColors) Int64() int64 { return int64(i) }

// SetInt64 sets the LightColors value from an int64.
func (i *LightColors) SetInt64(in int64) { *i = LightColors(in) }

// Desc returns the description of the LightColors value.
func (i LightColors) Desc() string { return enums.Desc(i, _LightColorsDescMap) }

// LightColorsValues returns all possible values for the type LightColors.
func LightColorsValues() []LightColors { return _LightColorsValues }

// Values returns all possible values for the type LightColors.
func (i LightColors) Values() []enums.Enum { return enums.Values(_LightColorsValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i LightColors) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *LightColors) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "LightColors")
}

var _RenderClassesValues = []RenderClasses{0, 1, 2, 3, 4, 5, 6}

// RenderClassesN is the highest valid value for type RenderClasses, plus one.
const RenderClassesN RenderClasses = 7

var _RenderClassesValueMap = map[string]RenderClasses{`None`: 0, `OpaqueTexture`: 1, `OpaqueUniform`: 2, `OpaqueVertex`: 3, `TransTexture`: 4, `TransUniform`: 5, `TransVertex`: 6}

var _RenderClassesDescMap = map[RenderClasses]string{0: ``, 1: ``, 2: ``, 3: ``, 4: ``, 5: ``, 6: ``}

var _RenderClassesMap = map[RenderClasses]string{0: `None`, 1: `OpaqueTexture`, 2: `OpaqueUniform`, 3: `OpaqueVertex`, 4: `TransTexture`, 5: `TransUniform`, 6: `TransVertex`}

// String returns the string representation of this RenderClasses value.
func (i RenderClasses) String() string { return enums.String(i, _RenderClassesMap) }

// SetString sets the RenderClasses value from its string representation,
// and returns an error if the string is invalid.
func (i *RenderClasses) SetString(s string) error {
	return enums.SetString(i, s, _RenderClassesValueMap, "RenderClasses")
}

// Int64 returns the RenderClasses value as an int64.
func (i RenderClasses) Int64() int64 { return int64(i) }

// SetInt64 sets the RenderClasses value from an int64.
func (i *RenderClasses) SetInt64(in int64) { *i = RenderClasses(in) }

// Desc returns the description of the RenderClasses value.
func (i RenderClasses) Desc() string { return enums.Desc(i, _RenderClassesDescMap) }

// RenderClassesValues returns all possible values for the type RenderClasses.
func RenderClassesValues() []RenderClasses { return _RenderClassesValues }

// Values returns all possible values for the type RenderClasses.
func (i RenderClasses) Values() []enums.Enum { return enums.Values(_RenderClassesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i RenderClasses) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *RenderClasses) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "RenderClasses")
}
