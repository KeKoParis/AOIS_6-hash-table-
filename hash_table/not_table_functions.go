package hash_table

import (
	"container/list"
	"fmt"
)

func findElement(checkingList *list.List, value string) (*list.Element, error) {
	for i := checkingList.Front(); i != nil; i.Next() {
		if i.Value == value {
			return i, nil
		}
	}

	return nil, fmt.Errorf("there's no %s in table", value)
}

func checkListReplica(checkingList *list.List, value string) bool {
	for i := checkingList.Front(); i != nil; i = i.Next() {
		if i.Value == value {
			return false
		}
	}

	return true
}
