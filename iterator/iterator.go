package iterator

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Book struct {
	name string
}

type BookShelf struct {
	Books []*Book
}

func (b *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{BookShelf: b}

}

func (b *BookShelf) AppendBook(book *Book) {
	b.Books = append(b.Books, book)
}

type BookShelfIterator struct {
	BookShelf *BookShelf
	Index     int
}

func (bi *BookShelfIterator) HasNext() bool {
	if bi.Index < len(bi.BookShelf.Books) {
		return true
	} else {
		return false
	}
}

func (bi *BookShelfIterator) Next() interface{} {
	book := bi.BookShelf.Books[bi.Index]
	bi.Index++
	return book
}
