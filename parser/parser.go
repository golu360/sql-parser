package parser

import (
	"errors"
	"fmt"
	"golu360/sql-parser/states"
	"strings"
)

type sqlparser struct {
	sql       string
	tokens    []string
	pos       int
	state     states.State
	queryType string
}

func (p *sqlparser) tokenize() (bool, error) {
	var tokens []string = strings.Split(p.sql, " ")
	if len(tokens) == 1 {
		return len(tokens) == 0, errors.New("Empty query")
	}
	p.tokens = tokens
	p.state = states.TOKENIZE_COMPLETE
	return len(tokens) == 1, nil
}
func (p *sqlparser) peek() string {
	if len(p.tokens) > 0 {
		return p.tokens[p.pos]
	}
	return ""
}

func Parse(sql string) {
	var queryParser sqlparser = sqlparser{sql, nil, 0, states.TOKENIZE, ""}
	_, err := queryParser.tokenize()
	if err != nil {
		fmt.Println(err)
	}
	queryParser.setQueryType()
	fmt.Println(queryParser)
}

func (p *sqlparser) setQueryType() {
	p.state = states.QUERY_TYPE_PARSE
	switch strings.ToUpper(p.peek()) {
	case "SELECT":
		p.queryType = "SELECT"
	case "UPDATE":
		p.queryType = "UPDATE"
	case "INSERT":
		p.queryType = "INSERT"
	case "DELETE":
		p.queryType = "DELETE"
	}
	p.state = states.QUERY_TYPE_PARSE_COMPLETE

}
