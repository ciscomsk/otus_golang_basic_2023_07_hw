package main

import "fmt"

const (
	BySize = "size"
	ByYear = "year"
	ByRate = "rate"
)

func main() {
	b1 := Book{year: 2010}
	b2 := Book{year: 2000}
	bc := NewBookComparer(ByYear)
	fmt.Println(bc.Compare(&b1, &b2))
}

type Book struct {
	id     string
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func (b *Book) ID() string {
	return b.id
}

func (b *Book) SetID(id string) {
	b.id = id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) Rate() float64 {
	return b.rate
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

type BookComparer struct {
	condition string
}

func NewBookComparer(condition string) *BookComparer {
	return &BookComparer{condition: condition}
}

func (bc *BookComparer) Compare(book1 *Book, book2 *Book) bool {
	var result bool

	switch bc.condition {
	case "year":
		if book1.year > book2.year {
			result = true
		} else {
			result = false
		}

	case "size":
		if book1.size > book2.size {
			result = true
		} else {
			result = false
		}

	case "rate":
		if book1.rate > book2.rate {
			result = true
		} else {
			result = false
		}
	}

	return result
}
