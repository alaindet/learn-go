// Code generated by "enumer -type=Month -output enumer_enum.go -json -yaml"; DO NOT EDIT.

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _MonthName = "MonthNoneMonthJanuaryMonthFebruaryMonthMarchMonthAprilMonthMayMonthJuneMonthJulyMonthAugustMonthSeptemberMonthOctoberMonthNovemberMonthDecember"

var _MonthIndex = [...]uint8{0, 9, 21, 34, 44, 54, 62, 71, 80, 91, 105, 117, 130, 143}

const _MonthLowerName = "monthnonemonthjanuarymonthfebruarymonthmarchmonthaprilmonthmaymonthjunemonthjulymonthaugustmonthseptembermonthoctobermonthnovembermonthdecember"

func (i Month) String() string {
	if i < 0 || i >= Month(len(_MonthIndex)-1) {
		return fmt.Sprintf("Month(%d)", i)
	}
	return _MonthName[_MonthIndex[i]:_MonthIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _MonthNoOp() {
	var x [1]struct{}
	_ = x[MonthNone-(0)]
	_ = x[MonthJanuary-(1)]
	_ = x[MonthFebruary-(2)]
	_ = x[MonthMarch-(3)]
	_ = x[MonthApril-(4)]
	_ = x[MonthMay-(5)]
	_ = x[MonthJune-(6)]
	_ = x[MonthJuly-(7)]
	_ = x[MonthAugust-(8)]
	_ = x[MonthSeptember-(9)]
	_ = x[MonthOctober-(10)]
	_ = x[MonthNovember-(11)]
	_ = x[MonthDecember-(12)]
}

var _MonthValues = []Month{MonthNone, MonthJanuary, MonthFebruary, MonthMarch, MonthApril, MonthMay, MonthJune, MonthJuly, MonthAugust, MonthSeptember, MonthOctober, MonthNovember, MonthDecember}

var _MonthNameToValueMap = map[string]Month{
	_MonthName[0:9]:          MonthNone,
	_MonthLowerName[0:9]:     MonthNone,
	_MonthName[9:21]:         MonthJanuary,
	_MonthLowerName[9:21]:    MonthJanuary,
	_MonthName[21:34]:        MonthFebruary,
	_MonthLowerName[21:34]:   MonthFebruary,
	_MonthName[34:44]:        MonthMarch,
	_MonthLowerName[34:44]:   MonthMarch,
	_MonthName[44:54]:        MonthApril,
	_MonthLowerName[44:54]:   MonthApril,
	_MonthName[54:62]:        MonthMay,
	_MonthLowerName[54:62]:   MonthMay,
	_MonthName[62:71]:        MonthJune,
	_MonthLowerName[62:71]:   MonthJune,
	_MonthName[71:80]:        MonthJuly,
	_MonthLowerName[71:80]:   MonthJuly,
	_MonthName[80:91]:        MonthAugust,
	_MonthLowerName[80:91]:   MonthAugust,
	_MonthName[91:105]:       MonthSeptember,
	_MonthLowerName[91:105]:  MonthSeptember,
	_MonthName[105:117]:      MonthOctober,
	_MonthLowerName[105:117]: MonthOctober,
	_MonthName[117:130]:      MonthNovember,
	_MonthLowerName[117:130]: MonthNovember,
	_MonthName[130:143]:      MonthDecember,
	_MonthLowerName[130:143]: MonthDecember,
}

var _MonthNames = []string{
	_MonthName[0:9],
	_MonthName[9:21],
	_MonthName[21:34],
	_MonthName[34:44],
	_MonthName[44:54],
	_MonthName[54:62],
	_MonthName[62:71],
	_MonthName[71:80],
	_MonthName[80:91],
	_MonthName[91:105],
	_MonthName[105:117],
	_MonthName[117:130],
	_MonthName[130:143],
}

// MonthString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func MonthString(s string) (Month, error) {
	if val, ok := _MonthNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _MonthNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Month values", s)
}

// MonthValues returns all values of the enum
func MonthValues() []Month {
	return _MonthValues
}

// MonthStrings returns a slice of all String values of the enum
func MonthStrings() []string {
	strs := make([]string, len(_MonthNames))
	copy(strs, _MonthNames)
	return strs
}

// IsAMonth returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Month) IsAMonth() bool {
	for _, v := range _MonthValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Month
func (i Month) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Month
func (i *Month) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Month should be a string, got %s", data)
	}

	var err error
	*i, err = MonthString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Month
func (i Month) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Month
func (i *Month) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = MonthString(s)
	return err
}
