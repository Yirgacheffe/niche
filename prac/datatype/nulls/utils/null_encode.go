package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type nullInt64 sql.NullInt64

// NullIntExample is the same, but uses a sql.NullInt64
type NullIntExample struct {
	Name string     `json:"name"`
	Age  *nullInt64 `json"age,omitempty"`
}

func (v *nullInt64) MarshalJSON() ([]byte, error) {

	if v.Valid {
		return json.Marshal(v.Int64)
	}

	return json.Marshal(nil)

}

func (v *nullInt64) UnmarshalJSON(b []byte) error {

	v.Valid = false

	if b != nil {
		v.Valid = true
		return json.Unmarshal(b, &v.Int64)
	}

	return nil // Re, return without error happened

}

// NullEncoding shows an alternative method for dealing
// with nil/omitted values
func NullEncoding() error {

	e := NullIntExample{}

	// no age
	err := json.Unmarshal([]byte(jsonBlob), &e)
	if err != nil {
		return err
	}
	fmt.Printf("nullint unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("nullint marshal, no age:", string(value))

	// age = 0
	err = json.Unmarshal([]byte(fullJSON), &e)
	if err != nil {
		return err
	}
	fmt.Printf("nullint unmarshal, age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("nullint marshal, age = 0:", string(value))

	return nil

}
