package compressor

import (
	"bytes"
	"compress/zlib"
	"errors"
	"io"
	"io/ioutil"
)

type ZlibCompressor struct {
}

func (_ ZlibCompressor) Zip(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
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
func (_ ZlibCompressor) Unzip(data []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(r)
	if err != nil && err != io.EOF && !errors.Is(err, io.ErrUnexpectedEOF) {
		return nil, err
	}
	return data, nil
}
