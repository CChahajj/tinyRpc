package compressor

import (
	"bytes"
	"errors"
	"github.com/golang/snappy"
	"io"
	"io/ioutil"
	"log"
)

type SnappyCompressor struct {
}

func (_ SnappyCompressor) Zip(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := snappy.NewBufferedWriter(&b)
	defer func() {
		err := w.Close()
		if err != nil {
			log.Println("close snappy:", err)
		}
	}()
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	err = w.Flush()
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (_ SnappyCompressor) Unzip(data []byte) ([]byte, error) {
	r := snappy.NewReader(bytes.NewReader(data))
	data, err := ioutil.ReadAll(r)
	if err != nil && err != io.EOF && errors.Is(err, io.ErrUnexpectedEOF) {
		return nil, err
	}
	return data, nil
}
