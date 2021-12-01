package types

import (
	"errors"
)

type Type string

func (t Type) Value() string {
	return string(t)
}

var (
	Bool   Type = "bool"
	Int    Type = "int"
	Float  Type = "float"
	String Type = "string"

	ErrInvalidType = errors.New("invalid type")
)

func GetType(value interface{}) (Type, error) {
	switch value.(type) {
	case bool:
		return Bool, nil
	case int:
		return Int, nil
	case float64:
		return Float, nil
	case string:
		return String, nil
	}
	return Type(""), ErrInvalidType
}
