package main

import (
	"golu360/sql-parser/parser"
)

func main() {
	var query string = "Select * from table"
	parser.Parse(query)

}
