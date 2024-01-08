package codec

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
)

func recvFrame(r io.Reader) ([]byte, error) {
	size, err := binary.ReadUvarint(r.(io.ByteReader))
	if err != nil {
		return nil, err
	}
	if size != 0 {
		data := make([]byte, size)
		if err = read(r, data); err != nil {
			return nil, err
		}
	}
	return data, nil
}

func sendFrame(w io.Writer, data []byte) error {

}

func write(w io.Writer, data []byte) {

}
func read(r io.Reader, data []byte) error {
	for index := 0; index < len(data); {
		n, err := r.Read(data[index:])
		if err != nil {
			var netErr net.Error
			if !errors.As(err, &netErr) {
				return err
			}
		}
		index += n
	}
	return nil
}
