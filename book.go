package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title: \t\t%q\n"+
			"Author: \t\t%q\n"+
			"Year Published: \t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{
		ID: 1,
		Title: "The Hitchhiker's Guide to the Galaxy",
		Author: "Douglas Adams",
		YearPublished: 1979,
	},
	Book{
		ID: 2,
		Title: "The Restaurant at the End of the Universe",
		Author: "Douglas Adams",
		YearPublished: 1980,
	},
	Book{
		ID: 3,
		Title: "Life, the Universe and Everything",
		Author: "Douglas Adams",
		YearPublished: 1982,
	},
	Book{
		ID: 4,
		Title: "The Hobbit",
		Author: "J.R.R. Tolkien",
		YearPublished: 1937,
	},
	Book{
		ID: 5,
		Title: "The Fellowship of the Ring",
		Author: "J.R.R. Tolkien",
		YearPublished: 1954,
	},
	Book{
		ID: 6,
		Title: "I, Robot",
		Author: "Isaac Asimov",
		YearPublished: 1950,
	},
}
