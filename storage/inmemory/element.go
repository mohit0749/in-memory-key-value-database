package inmemory

import "datastore/types"

type element struct {
	key      string
	attr     string
	value    interface{}
	typeInfo types.Type
}

func (e element) GetKey() string {
	return e.key
}

func (e element) GetAttr() string {
	return e.attr
}

func (e element) Bool() (bool, error) {
	if e.typeInfo == types.Bool {
		v, _ := e.value.(bool)
		return v, nil
	}
	return false, types.ErrInvalidType
}

func (e element) Int() (int, error) {
	if e.typeInfo == types.Int {
		v, _ := e.value.(int)
		return v, nil
	}
	return 0, types.ErrInvalidType
}

func (e element) String() (string, error) {
	if e.typeInfo == types.String {
		v, _ := e.value.(string)
		return v, nil
	}
	return "", types.ErrInvalidType
}

func (e element) Float() (float64, error) {
	if e.typeInfo == types.Float {
		v, _ := e.value.(float64)
		return v, nil
	}
	return 0.0, types.ErrInvalidType
}
