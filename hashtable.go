// Package hashtable provides a hash table data structure
// with STRING type keys
// using singly linked lists for collision resolution.
//
// A hashtable is a common data structure that allows
// efficient lookup, insertion and deletion of elements in the
// average case wherein it takes constant time. Note that the
// worst case time complexity of all these operations are O(N)
// however.
//
// This implementation assigns the keys to singly linked lists 
// which will in general have size 1, when a collision
// occurs the value which was already assigned to the key
// generated will now have a next-pointer to the new
// value and the new value will have a prev-pointer to 
// the old one, therefore letting us avoid collision.
// The hash table has initial size 8, following the reasoning
// of python's dictionary implementation, providing the best performance in 
// the most common use-cases, and will dynamically increase.
// The average time complexity for lookup and deletion will in fact
// be O(k) where k is the average length of the linked lists.
// Insertion is always constant since the lists contains a last-pointer.

// Code produced by Teo Jansson Minne, May 11 - 2020.

package hashtable


import (
	"bytes"

	"hash/fnv"
)

// HashTable is the table itself containing linked lists.
type HashTable struct {
	size       uint64
	sizeValues int
	lists      []*linkedList
	loadfactor float32
}

// data is what is stored in the hash table.
type data struct {
	key   []byte
	value interface{}
}

// New constructs a new HashTable.
func New() *HashTable {
	lists := make([]*linkedList, 8)
	for i := 0; i < 8; i++ {
		lists[i] = newList()
	}
	return &HashTable{
		size:       8,
		sizeValues: 0,
		lists:      lists,
		loadfactor: 0.75,
	}
}

// NewCustom constructs a new HashTable with custom parameters
// loadfactor and size to achieve best performance for
// a specific use-case. (ADVANCED USERS)
func NewCustom(loadfactor float32, size uint64) *HashTable {
	lists := make([]*linkedList, 8)
	for i := 0; i < int(size); i++ {
		lists[i] = newList()
	}
	return &HashTable{
		size:       size,
		sizeValues: 0,
		lists:      lists,
		loadfactor: loadfactor,
	}
}

// hasher return a hash for a given key.
// hashing function used is the Fowler Noll Vo
// algorithm.
func hasher(k []byte) uint64 {
	hash := fnv.New64a()
	hash.Write(k)
	return hash.Sum64()
}

// Find returns the value pointed at by the input key.
// If not in table returns nil.
func (tab *HashTable) Find(key string) interface{} {
	k := []byte(key)
	list := tab.lists[hasher(k)%tab.size]
	for current := list.first; current != nil; current = current.next {
		if bytes.Equal(current.element.key, k) {
			return current.element.value
		}
	}
	return nil
}

// Remove removes the input key and its value from the table.
func (tab *HashTable) Remove(key string) {
	k := []byte(key)
	list := tab.lists[hasher(k)%tab.size]
	prev := &listelement{}
	for current := list.first; current != nil; current = current.next {
		if bytes.Equal(current.element.key, k) {
			if current == list.first {
				list.first = current.next
				current = &listelement{}
				tab.sizeValues--
				return
			}
			prev.next = current.next
			current = &listelement{}
			tab.sizeValues--
			return
		}
		prev = current
	}
	return
}

// Add takes a key and a corresponding value as input at puts it in the hash table.
// also increases the size of the hash table if needed.
func (tab *HashTable) Add(key string, val interface{}) {
	if (tab.sizeValues + 1) / tab.Size() > tab.loadfactor{
		tab.sizeUp()
	}
	k := []byte(key)
	datapoint := data{
		key:   k,
		value: val,
	}
	list := tab.lists[hasher(k)%tab.size]
	for current := list.first; current != nil; current = current.next {
		if current.element.key == k{
			tab.Remove(k)
		}
	list.add(datapoint)
	tab.sizeValues++
	return
}

// sizeAdd is used in resizing methods.
func (tab *HashTable) sizeAdd(datapoint data, tempsize uint64) {
	k := datapoint.key
	newlist := tab.lists[hasher(k)%tempsize]
	newlist.add(datapoint)
	tab.sizeValues++
	return
}

// Size returns the current size(amount of lists) of the HashTable.
func (tab *HashTable) Size() int {
	return int(tab.size)
}

// SizeLists returns the amount of linked lists in use of the HashTable.
// Primary use is in conjunction with Size() call to judge efficiency.
func (tab *HashTable) SizeLists() int {
	counter := 0
	for i := 0; i > tab.Size(); i++ {
		if tab.lists[i].first != nil {
			counter++
		}
	}
	return counter
}

// SizeValues returns the total amount of values in the HashTable.
func (tab *HashTable) SizeValues() int {
	return tab.sizeValues
}

// Doubles the amount of lists in hashtable
func (tab *HashTable) sizeUp() {
	for i := 0; i < tab.Size(); i++ {
		tab.lists = append(tab.lists, newList())
	}
	tempsize := tab.size * uint64(2)
	var list *linkedList
	var key string
	for j := 0; j < tab.Size(); j++ {
		list = tab.lists[j]
		for current := list.first; current != nil; current = current.next {
			tab.sizeAdd(current.element, tempsize)
			key = string(current.elemet.key[:])
			tab.Remove(key)
		}
	}
	tab.size = tempsize
	return
}

// Below is the linked list and none of this is exported //

// linkedList is a singly linked list of elements of type "data".
type linkedList struct {
	first *listelement // first element in list
	last  *listelement // last element in list
}

type listelement struct {
	element data
	next    *listelement
}

// newList creates an empty list
func newList() *linkedList {
	return &linkedList{}
}

// add inserts the given element at the end of this list.
func (L *linkedList) add(element data) {
	elem := &listelement{
		element: element,
	}
	if L.first != nil {
		L.last.next = elem
		L.last = elem
	} else {
		L.first = elem
		L.last = elem
	}
}
