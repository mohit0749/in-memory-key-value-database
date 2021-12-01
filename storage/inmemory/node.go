package inmemory

import "container/list"

type node struct {
	dll        *list.List
	mp         map[interface{}]*list.Element
	keyValueMp map[string]*list.Element
}

func newNode() *node {
	return &node{list.New(), make(map[interface{}]*list.Element), make(map[string]*list.Element)}
}
