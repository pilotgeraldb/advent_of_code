package shared

import (
	"bufio"
	"log"
	"os"
)

type StreamInfo struct {
	R           rune
	Line        int
	LineIndex   int
	GlobalIndex int
}

func NewStreamInfo(r rune, line, lineIndex, globalIndex int) StreamInfo {
	return StreamInfo{
		R:           r,
		Line:        line,
		LineIndex:   lineIndex,
		GlobalIndex: globalIndex,
	}
}

func StreamProcess(path string, fn func(sctx StreamInfo)) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	c := 0
	g := 0
	l := 1

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		sctx := NewStreamInfo(r, l, c, g)
		fn(sctx)
		if r == '\n' {
			l++
			c = 0
			continue
		}
		c++
		g++
	}
}
