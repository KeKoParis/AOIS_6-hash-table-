package hash_table

import (
	"container/list"
)

const sizeTable = 20

type HashTable struct {
	pointers   [sizeTable]*list.List
	collisions [sizeTable]bool
}

func (ht *HashTable) NewHashTable() {
	for i := range ht.pointers {
		ht.pointers[i] = list.New()
	}
}

func (ht *HashTable) Add(key string, value string) {
	var tableIndex, hashValue int32

	hashValue = Hash(key)
	tableIndex = hashValue % sizeTable

	if (*ht.pointers[tableIndex]).Back() == nil {
		(*ht.pointers[tableIndex]).PushBack(value)
	} else {
		if checkListReplica(&(*ht.pointers[tableIndex]), value) {
			(*ht.pointers[tableIndex]).PushBack(value)
			ht.collisions[tableIndex] = true
		}
	}

}

func (ht *HashTable) Delete(key string, value string) {
	var tableIndex, hashValue int32

	hashValue = Hash(key)
	tableIndex = hashValue % sizeTable

	if ht.collisions[tableIndex] == true {
		foundElement, isError := findElement(&(*ht.pointers[tableIndex]), value)
		if foundElement == nil {
			panic(isError)
		}
		(*ht.pointers[tableIndex]).Remove(foundElement)
	} else {
		(*ht.pointers[tableIndex]).Remove((*ht.pointers[tableIndex]).Back())
	}
}

func (ht *HashTable) Get(key string) list.List {
	var tableIndex, hashValue int32

	hashValue = Hash(key)
	tableIndex = hashValue % sizeTable

	return *ht.pointers[tableIndex]
}
