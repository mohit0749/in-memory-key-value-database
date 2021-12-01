package inmemory

import (
	"container/list"
	"datastore/storage"
	"datastore/types"
	"sync"
)

type store struct {
	attr       map[string]*node
	keyAttrMap map[string]map[string]struct{}
	mutex      *sync.Mutex
}

func NewStorage() *store {
	return &store{make(map[string]*node),
		make(map[string]map[string]struct{}),
		&sync.Mutex{}}
}

func (s store) Add(key, attr string, value interface{}, typeInfo types.Type) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	n, ok := s.attr[attr]
	if !ok {
		n = newNode()
		n.keyValueMp = make(map[string]*list.Element)
		s.attr[attr] = n
	}
	item := &element{
		key:      key,
		value:    value,
		typeInfo: typeInfo,
		attr:     attr,
	}
	var listNodePtr *list.Element
	if _, ok := n.mp[value]; !ok {
		listNodePtr = n.dll.PushBack(item)
		n.mp[value] = listNodePtr
	} else {
		listNodePtr = n.dll.InsertAfter(item, n.mp[value])
	}
	if _, ok := s.keyAttrMap[key]; !ok {
		s.keyAttrMap[key] = make(map[string]struct{})
	}
	s.keyAttrMap[key][attr] = struct{}{}
	if _, ok := n.keyValueMp[key]; !ok {
		n.keyValueMp[key] = listNodePtr
	}
	return nil
}

func (s store) Get(key string) ([]storage.Attribute, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.keyAttrMap[key]; !ok {
		return nil, storage.ErrKeyNotExists
	}
	values := make([]storage.Attribute, 0)
	for k, _ := range s.keyAttrMap[key] {
		n := s.attr[k]
		item := n.keyValueMp[key].Value.(*element)
		values = append(values, item)
	}
	return values, nil
}

func (s store) GetAttr(attr string, value interface{}) ([]storage.Attribute, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	n, ok := s.attr[attr]
	if !ok {
		return nil, storage.ErrKeyAttrNotExists
	}
	start := n.mp[value]
	results := make([]storage.Attribute, 0)
	for start != nil {
		item := start.Value.(*element)
		if item.value != value {
			break
		}
		results = append(results, item)
		start = start.Next()
	}
	return results, nil
}

func (s store) Remove(key string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.keyAttrMap[key]; !ok {
		return storage.ErrKeyNotExists
	}
	for attr, _ := range s.keyAttrMap[key] {
		s.removeAttr(key, attr)
	}
	delete(s.keyAttrMap, key)
	return nil
}

func (s store) removeAttr(key, attr string) {
	n := s.attr[attr]
	elem := n.keyValueMp[key].Value
	item := elem.(*element)
	delete(n.keyValueMp, key)
	if elem == n.mp[item.value] && n.mp[item.value].Next() == nil {
		delete(n.mp, item.value)
	} else {
		n.mp[item.value] = n.mp[item.value].Next()
	}
}
