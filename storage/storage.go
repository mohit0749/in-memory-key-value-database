package storage

import (
	"datastore/types"
	"errors"
)

var (
	ErrKeyNotExists     = errors.New("key does not exists")
	ErrKeyAttrNotExists = errors.New("key does not exists")
)

type Store interface {
	Add(key, attr string, value interface{}, typeInfo types.Type) error
	Get(key string) ([]Attribute, error)
	GetAttr(attr string, value interface{}) ([]Attribute, error)
	Remove(key string) error
}

type Values interface {
	GetKey() string
	GetAllValue() []Attribute
	GetValue(attr string)
}

type Attribute interface {
	GetKey() string
	GetAttr() string
	Bool() (bool, error)
	Int() (int, error)
	String() (string, error)
	Float() (float64, error)
}
