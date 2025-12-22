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
	LineRuneArr []rune
}

func NewEmptyStreamInfo() StreamInfo {
	return StreamInfo{}
}

func NewStreamInfo(r rune, line, lineIndex, globalIndex int) StreamInfo {
	return StreamInfo{
		R:           r,
		Line:        line,
		LineIndex:   lineIndex,
		GlobalIndex: globalIndex,
	}
}

func (s StreamInfo) GetLine() string {
	return string(s.LineRuneArr)
}

type StreamProcessConfig struct {
	EOFCallback func() // end of file callback
}

func NewStreamProcessConfig(opts ...StreamProcessOption) StreamProcessConfig {
	cfg := StreamProcessConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

type StreamProcessOption func(*StreamProcessConfig)

func WithEOFCallback(eofCallback func()) StreamProcessOption {
	return func(o *StreamProcessConfig) {
		o.EOFCallback = eofCallback
	}
}

func StreamProcess(path string, fn func(sctx StreamInfo), opts ...StreamProcessOption) {
	if fn == nil {
		log.Fatalf("fn cannot be nil")
	}

	cfg := NewStreamProcessConfig(opts...)

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	c := 0
	g := 0
	l := 1

	sctx := NewEmptyStreamInfo()

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if cfg.EOFCallback != nil {
				cfg.EOFCallback()
			}
			break
		}
		sctx.R = r
		sctx.Line = l
		sctx.LineIndex = c
		sctx.GlobalIndex = g
		if r != '\n' && r != '\r' {
			sctx.LineRuneArr = append(sctx.LineRuneArr, r)
		}
		fn(sctx)
		if r == '\n' {
			l++
			g++
			c = 0
			sctx.LineRuneArr = []rune{}
			continue
		} else {
			c++
			g++
		}
	}
}
