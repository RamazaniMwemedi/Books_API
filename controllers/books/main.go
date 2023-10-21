package books

type Book struct {
	Title  string
	Author string
	ISBN   string
}

func GetBooks() []Book {
	books := []Book{
		Book{Title: "The Alchemist", Author: "Paulo Coelho", ISBN: "978-0062315007"},
		Book{Title: "The Catcher in the Rye", Author: "J.D. Salinger", ISBN: "978-0316769488"},
		Book{Title: "To Kill a Mockingbird", Author: "Harper Lee", ISBN: "978-0446310789"},
	}
	return books
}

