package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {
	bookShelf := &BookShelf{}

	asserts := []string{"A", "B", "C"}

	for _, assert := range asserts {
		bookShelf.AppendBook(&Book{assert})
	}

	i := 0
	for iterator := bookShelf.Iterator(); iterator.HasNext(); {
		if book := iterator.Next(); book.(*Book).name != asserts[i] {
			t.Errorf("Expect book.name is %s, but get %s", asserts[i], book.(*Book).name)
		}
		i++
	}

}
