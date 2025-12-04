package shared

import (
	"bufio"
	"log"
	"os"
)

func StreamProcess(path string, fn func(rune)) {
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
