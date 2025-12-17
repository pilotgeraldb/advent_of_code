package shared

import (
	"bufio"
	"log"
	"os"
)

func StreamProcess(path string, fn func(r rune, line int, idx int)) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	c := 0
	l := 1
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		fn(r, l, c)
		if r == '\n' {
			l++
			c = 0
			continue
		}
		c++
	}
}
