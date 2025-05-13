package main

type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	books := []Book{
		{"The Great Gatsby", "F. Scott Fitzgerald", 1925},
		{"To Kill a Mockingbird", "Harper Lee", 1960},
		{"1984", "George Orwell", 1949},
		{"Pride and Prejudice", "Jane Austen", 1813},
		{"The Catcher in the Rye", "J.D. Salinger", 1951},
		{"The Hobbit", "J.R.R. Tolkien", 1937},
		{"Fahrenheit 451", "Ray Bradbury", 1953},
	}
	// Print the list of books
	for _, book := range books {
		println("Title: ", book.Title)
		println("Author:", book.Author)
		println("Year:  ", book.Year)
		println()
	}
	// Print the total number of books
	println("Total number of books:", len(books))
}
