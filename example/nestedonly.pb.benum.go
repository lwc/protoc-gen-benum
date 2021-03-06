// Code generated by protoc-gen-benum. DO NOT EDIT.
package example

import (
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
)

// -------------------------------------------------------------------
// OutResponse_Enum2
// -------------------------------------------------------------------

func (e OutResponse_Enum2) IsValid() bool {
	_, ok := OutResponse_Enum2_name[int32(e)]
	return ok
}

// ------------------------- gqlgen ----------------------------------

var gql_OutResponse_Enum2_name = map[int32]string{
	0: "PROFILE",
	1: "BILLING",
	2: "LEGAL",
}

var gql_OutResponse_Enum2_value = map[string]int32{
	"PROFILE": 0,
	"BILLING": 1,
	"LEGAL":   2,
}

func (e OutResponse_Enum2) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(gql_OutResponse_Enum2_name[int32(e)]))
}

func (e *OutResponse_Enum2) UnmarshalGQL(v interface{}) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a valid OutResponse_Enum2", v)
	}
	res, ok := gql_OutResponse_Enum2_value[value]
	if !ok {
		return fmt.Errorf("%T is not a valid OutResponse_Enum2", v)
	}
	*e = OutResponse_Enum2(res)
	return nil
}

// --------------------------- db ------------------------------------

var db_OutResponse_Enum2_name = map[int32]string{
	0: "PROFILE",
	1: "billing",
	2: "LEGAL",
}

var db_OutResponse_Enum2_value = map[string]int32{
	"PROFILE": 0,
	"billing": 1,
	"LEGAL":   2,
}

func (e OutResponse_Enum2) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, fmt.Errorf("invalid OutResponse_Enum2 '%s'", e)
	}
	return db_OutResponse_Enum2_name[int32(e)], nil
}

func (e *OutResponse_Enum2) Scan(value interface{}) error {
	sv, err := driver.String.ConvertValue(value)
	if err != nil {
		return fmt.Errorf("failed to scan %#v into OutResponse_Enum2", value)
	}
	res, ok := int32(0), false
	switch v := sv.(type) {
	case string:
		res, ok = db_OutResponse_Enum2_value[v]
	case []byte:
		res, ok = db_OutResponse_Enum2_value[string(v)]
	default:
		panic("unexpected type from ConvertValue")
	}
	if !ok {
		return fmt.Errorf("invalid database value for OutResponse_Enum2: '%s'", sv)
	}
	*e = OutResponse_Enum2(res)
	return nil
}
