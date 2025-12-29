package shared

import (
	"bufio"
	"log"
	"os"
)

type StreamInfo struct {
	R           Character
	LineNumber  int // 1-based line number
	LineIndex   int // 0-based column in line (excluding newline)
	GlobalIndex int // 0-based position in file
	LineRuneArr []rune
}

// GetLine returns the current lines text, as it has been read so-far, as a string.
func (s StreamInfo) GetLine() string {
	return string(s.LineRuneArr)
}

// StreamProcessConfig holds configuration options for StreamProcess.
type StreamProcessConfig struct {
	EOFCallback     func()
	NewlineCallback func(si StreamInfo)
	CRCallback      func(sctx StreamInfo)
}

// StreamProcessOption defines a functional option for configuring StreamProcess.
type StreamProcessOption func(*StreamProcessConfig)

// WithEOFCallback sets a callback function to be called when EOF is reached.
func WithEOFCallback(f func()) StreamProcessOption {
	return func(c *StreamProcessConfig) { c.EOFCallback = f }
}

// WithNewlineCallback sets a callback function to be called when a newline is encountered.
func WithNewlineCallback(f func(si StreamInfo)) StreamProcessOption {
	return func(c *StreamProcessConfig) { c.NewlineCallback = f }
}

// WithCRCallback sets a callback function to be called when a carriage return is encountered.
func WithCRCallback(f func(sctx StreamInfo)) StreamProcessOption {
	return func(c *StreamProcessConfig) { c.CRCallback = f }
}

// StreamProcess reads a file at the given path and processes it character by character,
func StreamProcess(path string, fn func(StreamInfo), opts ...StreamProcessOption) error {
	if fn == nil {
		log.Fatal("fn cannot be nil")
	}

	cfg := StreamProcessConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	line := 1
	col := 0
	globalIdx := 0
	var lineRunes []rune

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == bufio.ErrBufferFull {
				continue
			}
			if cfg.EOFCallback != nil {
				cfg.EOFCallback()
			}
			if err.Error() == "EOF" {
				return nil
			}
			return err
		}

		sctx := StreamInfo{
			R:           Character(r),
			LineNumber:  line,
			LineIndex:   col,
			GlobalIndex: globalIdx,
			LineRuneArr: append([]rune(nil), lineRunes...),
		}

		if r != CarriageReturn && r != Newline {
			fn(sctx)
		}

		if r == CarriageReturn {
			if cfg.CRCallback != nil {
				cfg.CRCallback(sctx)
			}
			if next, err := reader.Peek(1); err == nil && len(next) == 1 && next[0] == byte(Newline) {
				reader.ReadRune()
				globalIdx++
			}
			if cfg.NewlineCallback != nil {
				cfg.NewlineCallback(sctx)
			}
			line++
			col = 0
			lineRunes = lineRunes[:0]
			globalIdx++
			continue
		}

		if r == Newline {
			if cfg.NewlineCallback != nil {
				cfg.NewlineCallback(sctx)
			}
			line++
			col = 0
			lineRunes = lineRunes[:0]
			globalIdx++
			continue
		}

		lineRunes = append(lineRunes, r)
		col++
		globalIdx++
	}
}
