package readers

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestFixedLengthFileReaderReturnsErrorIfStreamIsEmpty(t *testing.T) {
	stream := bytes.NewReader([]byte(""))
	_, err := NewFixedLengthFileReader(true, []int{0, 1, 2}, bufio.NewReader(stream))
	if err == nil {
		fmt.Println(err.Error())
		t.Fatal(err.Error())
	}
}

func TestFixedLengthFileReaderSkipsHeaderIfSkipHeaderIsTrue(t *testing.T) {
	stream := bytes.NewReader([]byte("1234\n111"))
	r, _ := NewFixedLengthFileReader(true, []int{1, 1, 1}, bufio.NewReader(stream))
	var fields = make([]string, 3)
	for _, f := range fields {
		f, _ = r.Scan()
		if f != "1" {
			t.FailNow()
		}
	}
}

func TestFixedLengthFileReaderPrintHeaderIfSkipHeaderIsFalse(t *testing.T) {
	t.Log("Not supported at this time")
}

func TestFixedLengthFileReaderPrintHeaderReturnsErrorIfFileIsBroken(t *testing.T) {
	stream := bytes.NewReader([]byte("1234\n11111"))
	r, _ := NewFixedLengthFileReader(true, []int{1, 1, 1}, bufio.NewReader(stream))
	var fields = make([]string, 3)
	for _, f := range fields {
		f, _ = r.Scan()
		if f != "1" {
			t.FailNow()
		}
	}

	//expecting eof
	_, err := r.Scan()
	if err == nil {
		t.Fatal(err)
	}
}

func TestFixedLengthFileReaderPrintHeaderReturnsErrorIfEOF(t *testing.T) {
	stream := bytes.NewReader([]byte("1234\n111"))
	r, _ := NewFixedLengthFileReader(true, []int{1, 1, 1}, bufio.NewReader(stream))
	var fields = make([]string, 3)
	for range fields {
		r.Scan()
	}

	token, err := r.Scan()
	if err != io.EOF {
		t.Fatal(fmt.Sprintf("%s", token))
	}
}
