package booksModule

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func GetBooks() []Book {
	books := []Book{
		{ID: 1, Title: "The Alchemist", Author: "Paulo Coelho", ISBN: "978-0062315007"},
		{ID: 2, Title: "The Catcher in the Rye", Author: "J.D. Salinger", ISBN: "978-0316769488"},
		{ID: 3, Title: "To Kill a Mockingbird", Author: "Harper Lee", ISBN: "978-0446310789"},
	}
	return books
}
