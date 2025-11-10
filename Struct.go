package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func struct1() {
	books := Books{"Go 语言", "ljx", "Go 语言教程", 101}
	fmt.Printf("Book title : %s\n", books.title)
	fmt.Printf("Book author : %s\n", books.author)
	fmt.Printf("Book subject : %s\n", books.subject)
	fmt.Printf("Book book_id : %d\n", books.bookId)
	books.bookId = 102
	fmt.Println(books)
}
