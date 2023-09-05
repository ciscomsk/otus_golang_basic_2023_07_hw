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
	bc := newBookComparer(ByYear)
	fmt.Println(bc.compare(&b1, &b2))
}

type Book struct {
	id     string
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func (b Book) getID() string {
	return b.id
}

func (b *Book) setID(id string) {
	b.id = id
}

func (b Book) getTitle() string {
	return b.title
}

func (b *Book) setTitle(title string) {
	b.title = title
}

func (b Book) getAuthor() string {
	return b.author
}

func (b *Book) setAuthor(author string) {
	b.author = author
}

func (b Book) getYear() int {
	return b.year
}

func (b *Book) setYear(year int) {
	b.year = year
}

func (b Book) getSize() int {
	return b.size
}

func (b *Book) setSize(size int) {
	b.size = size
}

func (b Book) getRate() float64 {
	return b.rate
}

func (b *Book) setRate(rate float64) {
	b.rate = rate
}

type BookComparer struct {
	condition string
}

func newBookComparer(condition string) *BookComparer {
	return &BookComparer{condition: condition}
}

func (bc BookComparer) compare(book1 *Book, book2 *Book) bool {
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
