package search

import (
	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/linkedlist"
	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/revlog"
)

func LinearSearch(list linkedlist.LinkedList, id int) *linkedlist.LinkedListNode {
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
