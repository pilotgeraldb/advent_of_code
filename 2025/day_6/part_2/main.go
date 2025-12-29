package main

import (
	"log"
	shared "shared_aoc_2025"
	"strconv"
	"strings"
)

type value struct {
	strVal string
}

// Int converts the value to an integer.
func (v value) Int() int {
	strings.TrimSpace(v.strVal)
	intVal, err := strconv.Atoi(v.strVal)
	if err != nil {
		panic(err)
	}
	return intVal
}

// IsEmpty checks if the string value is empty.
func (v value) IsEmpty() bool {
	return len(strings.Trim(v.strVal, " ")) == 0
}

type column struct {
	values    []value
	operation Operator
}

var (
	columns []column
)

func main() {
	cv := value{}
	colIdx := 0
	shared.StreamProcess("test.txt", func(si shared.StreamInfo) {
		switch si.R {
		case '+', '-', '*', '/':
		case '\n', '\r':
			cv = value{}
			colIdx++
			if colIdx >= len(columns) {
				colIdx = 0
			}
		case ' ':
			//cv.strVal += si.R.String()
			if si.LineNumber == 1 {
				//if !cv.IsEmpty() {
				columns = append(columns, column{
					values: []value{cv},
				})
				//}
			}
			cv = value{}
		default:
			cv.strVal += si.R.String()
		}
	}, shared.WithEOFCallback(func() {
		total := 0
		log.Printf("total: %d", total)
	}), shared.WithNewlineCallback(func(si shared.StreamInfo) {

	}), shared.WithCRCallback(func(si shared.StreamInfo) {

	}))
}

// 7098065460541 -- correct
