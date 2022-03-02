package main

import (
	"log"
	"os"
	"strconv"
)

type table struct {
	col                int
	row                int
	highlightPositions []int
}

var err error

func main() {
	table := table{}
	table.col, err = strconv.Atoi(os.Args[1])
	table.row, err = strconv.Atoi(os.Args[2])
	if err != nil {
		log.Println("Args is not number")
		os.Exit(0)
	}

	for _, value := range os.Args[3:] {
		position, err := strconv.Atoi(value)
		if position > table.col || position > table.row {
			log.Println("Highlight positions is incorrect")
			os.Exit(0)
		}
		if err != nil {
			log.Println("Args is not number")
			os.Exit(0)
		}
		table.highlightPositions = append(table.highlightPositions, position)
	}

	log.Printf("total non-highlighted cells is %d", table.countNonHighlighted())
}

func (table table) countNonHighlighted() int {
	totalCell := table.col * table.row
	highlightCell := (len(table.highlightPositions) * table.col) + (len(table.highlightPositions) * table.row) - (len(table.highlightPositions) * len(table.highlightPositions))

	return totalCell - highlightCell
}
