package testpb

import (
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
)

// -------------------------------------------------------------------
// Enum1
// -------------------------------------------------------------------

func (e Enum1) IsValid() bool {
	_, ok := Enum1_name[int32(e)]
	return ok
}

// ------------------------- gqlgen ----------------------------------

var gql_Enum1_name = map[int32]string{
	0: "turkey",
	1: "BILLING",
	2: "blarg",
}

var gql_Enum1_value = map[string]int32{
	"turkey":  0,
	"BILLING": 1,
	"blarg":   2,
}

func (e Enum1) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(gql_Enum1_name[int32(e)]))
}

func (e *Enum1) UnmarshalGQL(v interface{}) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a valid Enum1", v)
	}
	res, ok := gql_Enum1_value[value]
	if !ok {
		return fmt.Errorf("%T is not a valid Enum1", v)
	}
	*e = Enum1(res)
	return nil
}

// --------------------------- db ------------------------------------

var db_Enum1_name = map[int32]string{
	0: "snake",
	1: "BILLING",
	2: "LEGAL",
}

var db_Enum1_value = map[string]int32{
	"snake":   0,
	"BILLING": 1,
	"LEGAL":   2,
}

func (e Enum1) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, fmt.Errorf("invalid Enum1 '%s'", e)
	}
	return db_Enum1_name[int32(e)], nil
}

func (e *Enum1) Scan(value interface{}) error {
	sv, err := driver.String.ConvertValue(value)
	if err != nil {
		return fmt.Errorf("failed to scan %#v into Enum1", value)
	}
	res, ok := int32(0), false
	switch v := sv.(type) {
	case string:
		res, ok = db_Enum1_value[v]
	case []byte:
		res, ok = db_Enum1_value[string(v)]
	default:
		panic("unexpected type from ConvertValue")
	}
	if !ok {
		panic(fmt.Errorf("invalid Enum1 '%s'", e))
	}
	*e = Enum1(res)
	return nil
}

// -------------------------------------------------------------------
// EchoResponse_Enum2
// -------------------------------------------------------------------

func (e EchoResponse_Enum2) IsValid() bool {
	_, ok := EchoResponse_Enum2_name[int32(e)]
	return ok
}

// ------------------------- gqlgen ----------------------------------

var gql_EchoResponse_Enum2_name = map[int32]string{
	0: "PROFILE",
	1: "BILLING",
	2: "LEGAL",
}

var gql_EchoResponse_Enum2_value = map[string]int32{
	"PROFILE": 0,
	"BILLING": 1,
	"LEGAL":   2,
}

func (e EchoResponse_Enum2) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(gql_EchoResponse_Enum2_name[int32(e)]))
}

func (e *EchoResponse_Enum2) UnmarshalGQL(v interface{}) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a valid EchoResponse_Enum2", v)
	}
	res, ok := gql_EchoResponse_Enum2_value[value]
	if !ok {
		return fmt.Errorf("%T is not a valid EchoResponse_Enum2", v)
	}
	*e = EchoResponse_Enum2(res)
	return nil
}

// --------------------------- db ------------------------------------

var db_EchoResponse_Enum2_name = map[int32]string{
	0: "PROFILE",
	1: "billing",
	2: "LEGAL",
}

var db_EchoResponse_Enum2_value = map[string]int32{
	"PROFILE": 0,
	"billing": 1,
	"LEGAL":   2,
}

func (e EchoResponse_Enum2) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, fmt.Errorf("invalid EchoResponse_Enum2 '%s'", e)
	}
	return db_EchoResponse_Enum2_name[int32(e)], nil
}

func (e *EchoResponse_Enum2) Scan(value interface{}) error {
	sv, err := driver.String.ConvertValue(value)
	if err != nil {
		return fmt.Errorf("failed to scan %#v into EchoResponse_Enum2", value)
	}
	res, ok := int32(0), false
	switch v := sv.(type) {
	case string:
		res, ok = db_EchoResponse_Enum2_value[v]
	case []byte:
		res, ok = db_EchoResponse_Enum2_value[string(v)]
	default:
		panic("unexpected type from ConvertValue")
	}
	if !ok {
		panic(fmt.Errorf("invalid Enum1 '%s'", e))
	}
	*e = EchoResponse_Enum2(res)
	return nil
}
