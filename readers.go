package readers

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

type FixedLengthFileReader struct {
	StopSymbols  []int
	streamReader *bufio.Reader
	curPosition  int
}

func NewFixedLengthFileReader(SkipHeader bool, FieldWidths []int, reader *bufio.Reader) (*FixedLengthFileReader, error) {
	// fast forward to the next line
	if SkipHeader {
		if bytes, err := reader.ReadBytes('\n'); err != nil || len(bytes) == 0 {
			return nil, errors.New("Empty stream")
		}
	}

	return &FixedLengthFileReader{StopSymbols: FieldWidths, streamReader: reader, curPosition: 0}, nil
}

func (reader *FixedLengthFileReader) Scan() (string, error) {
	// check for last token on the line
	if reader.curPosition == len(reader.StopSymbols) {
		// we expect either or both /r/n are the last chars
		if token, err := reader.streamReader.ReadBytes('\n'); len(token) > 2 || err != nil {
			if err == nil {
				err = errors.New(fmt.Sprintf("broken file %v", token))
			}

			return string(token), err
		} else {
			reader.curPosition = 0
			return string(token), nil
		}
	}

	fieldSize := reader.StopSymbols[reader.curPosition]
	var tmpBuffer bytes.Buffer
	for i := 0; i < fieldSize; i++ {
		r, _, _ := reader.streamReader.ReadRune()
		tmpBuffer.WriteRune(r)
	}

	reader.curPosition = reader.curPosition + 1
	return tmpBuffer.String(), nil
}
