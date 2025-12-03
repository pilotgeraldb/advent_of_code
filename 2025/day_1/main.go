package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Dial struct {
	value  int32
	dir    string
	count  int32
	amount []rune
}

func NewDial() *Dial {
	return &Dial{
		value:  50,
		dir:    "",
		count:  0,
		amount: []rune{},
	}
}

func main() {
	d := NewDial()

	streamProcess("input.txt", func(r rune) {
		switch r {
		case 'L':
			d.dir = "L"
		case 'R':
			d.dir = "R"
		case '\r':
			return
		case '\n':
			d.Turn(d.getAmount())
			d.Reset()
		default:
			d.amount = append(d.amount, r)
		}
	})

	log.Printf("answer: %d", d.count)
}

func (d *Dial) Reset() {
	d.amount = []rune{}
}

func (d *Dial) getAmount() int32 {
	if len(d.amount) == 0 {
		return 0
	}
	v, err := strconv.Atoi(string(d.amount))
	if err != nil {
		log.Fatalf("failed to convert amount [%s] to int: %s", string(d.amount), err)
	}
	return int32(v)
}

func (d *Dial) Turn(steps int32) {
	if steps == 0 {
		// no turning required
		return
	}

	d.count += steps / 100 // how many times we cross 0

	rem := steps % 100 // remaining steps after full circles
	if rem == 0 {
		return
	}

	if d.dir == "R" {
		if d.value+rem >= 100 { // checks whether turning right (adding rem) will cross zero
			d.count++
		}
		d.value = (d.value + rem) % 100 // ensure non-negative
	} else {
		if d.value > 0 && d.value <= rem { // checks whether turning left (subtracting rem) will cross zero
			d.count++
		}
		d.value = (d.value - rem + 100) % 100 // ensure non-negative
	}
}

func streamProcess(path string, fn func(rune)) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		fn(r)
	}
}
