package search

import (
	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/linkedList"
	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/revlog"
)

func LinearSearch(list linkedList.LinkedList, id int) *linkedList.LinkedListNode {
	current := list.First()
	for current != nil {
		currentData := current.Data.(revlog.Revision)
		if currentData.ID() == id {
			return current
		}
		current = current.Next
	}
	return nil
}
