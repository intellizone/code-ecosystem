package main

type Product struct {
	Id          int
	Name        string
	Description string
	Stock       int
}

type Inventory []Product

var inventory = Inventory{
	{Id: 1, Name: "Hamam Soap", Description: "Can wash you well.", Stock: 50},
	{Id: 2, Name: "Steel Bottle", Description: "Can be used to store water or other liquids", Stock: 3},
}

func (p Inventory) Exists(id int) bool {
	for _, i := range p {
		if i.Id == id {
			return true
		}
	}
	return false
}

func (p Inventory) NewId() int {
	id := len(p)
	for {
		if !p.Exists(id) {
			return id
		}
		id++
	}
}
