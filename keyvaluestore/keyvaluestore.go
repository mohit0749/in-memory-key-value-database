package keyvaluestore

import (
	"datastore/storage"
	"datastore/storage/inmemory"
	"datastore/types"
	"errors"
	"strings"
)

type keyValue struct {
	store storage.Store
	types map[string]types.Type
}

type Value struct {
	attr  string
	value interface{}
}

func NewKeyValueStore() *keyValue {
	return &keyValue{
		store: inmemory.NewStorage(),
		types: make(map[string]types.Type),
	}
}

func (k keyValue) Add(key string, attr string, value interface{}) error {
	key = strings.ToLower(key)
	typ, err := types.GetType(value)
	if err != nil {
		return err
	}
	prevTyp, ok := k.types[attr]
	if !ok {
		k.types[attr] = typ
	} else if prevTyp != typ {
		return errors.New("expecting type: " + prevTyp.Value())
	}
	err = k.store.Add(key, attr, value, typ)
	return err
}

func (k keyValue) Get(key string) ([]Value, error) {
	key = strings.ToLower(key)
	attrs, err := k.store.Get(key)
	if err != nil {
		return nil, err
	}
	keys := make([]Value, 0)
	for _, v := range attrs {
		keys = append(keys, Value{
			attr:  v.GetKey(),
			value: getValueByType(k.types[v.GetAttr()], v),
		})
	}
	return keys, nil
}

func (k keyValue) Remove(key string) error {
	key = strings.ToLower(key)
	return k.store.Remove(key)
}

func (k keyValue) Scan(attr string, value interface{}) ([]Value, error) {
	attr = strings.ToLower(attr)
	attrs, err := k.store.GetAttr(attr, value)
	if err != nil {
		return nil, err
	}
	keys := make([]Value, 0)
	for _, v := range attrs {
		keys = append(keys, Value{
			attr:  v.GetKey(),
			value: getValueByType(k.types[attr], v),
		})
	}
	return keys, nil
}

func getValueByType(typ types.Type, value storage.Attribute) (v interface{}) {
	switch typ {
	case types.Bool:
		v, _ = value.Bool()
	case types.Int:
		v, _ = value.Int()
	case types.Float:
		v, _ = value.Float()
	case types.String:
		v, _ = value.String()
	}
	return
}
