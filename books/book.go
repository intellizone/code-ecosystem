package main

import "fmt"

type Book struct {
	Id            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{
		Id:            1,
		Title:         "The Power of Comics: History, Form and Culture",
		Author:        "Randy Duncan",
		YearPublished: 2009,
	},
	{
		Id:            2,
		Title:         "To Kill a Mockingbird",
		Author:        "Harper Lee",
		YearPublished: 1960,
	},
	{
		Id:            3,
		Title:         "A Little Book for Little Children",
		Author:        "Thomas White",
		YearPublished: 1702,
	},
	{
		Id:            4,
		Title:         "A Token for Children",
		Author:        "James Janeway",
		YearPublished: 1709,
	},
	{
		Id:            5,
		Title:         "Moral Tales for Young People",
		Author:        "Maria Edgeworth",
		YearPublished: 1801,
	},
	{
		Id:            6,
		Title:         "The South Sea Whaler",
		Author:        "W. H. G. Kingston",
		YearPublished: 1859,
	},
	{
		Id:            7,
		Title:         "The Red Fighter Pilot",
		Author:        "Manfred von Richthofen",
		YearPublished: 1917,
	},
	{
		Id:            8,
		Title:         "The C Programming Language",
		Author:        "Dennis Ritchie",
		YearPublished: 1978,
	},
	{
		Id:            9,
		Title:         "Winged Warfare",
		Author:        "Billy Bishop",
		YearPublished: 1918,
	},
	{
		Id:            10,
		Title:         "Five Years in the Royal Flying Corps",
		Author:        "James McCudden",
		YearPublished: 1919,
	},
	{
		Id:            11,
		Title:         "Sagittarius Rising",
		Author:        "Cecil Lewis",
		YearPublished: 1936,
	},
	{
		Id:            12,
		Title:         "Wind, Sand and Stars",
		Author:        "Antoine de Saint-Exupery",
		YearPublished: 1939,
	},
	{
		Id:            13,
		Title:         "Reach for the Sky",
		Author:        "Paul Brickhill",
		YearPublished: 1954,
	},
	{
		Id:            14,
		Title:         "The Right Stuff",
		Author:        "Tom Wolfe",
		YearPublished: 1979,
	},
	{
		Id:            15,
		Title:         "Fate Is the Hunter",
		Author:        "Ernest K. Gann",
		YearPublished: 1961,
	},
}
