package main

import (
    "log"
	"os"
	"io"
	"bytes"
	"encoding/binary"

    "github.com/pkg/errors"

	"github.com/grpc/grpc-go/examples/rain/fsblkstorage"
	"github.com/grpc/grpc-go/examples/rain/common"
)


type blockFileWriter struct {
    filePath string
    file *os.File
}

func newBlockFileWriter(filePath string) (*blockFileWriter, error){
    writer := &blockFileWriter{filePath: filePath}
    return writer, writer.open()
}

func (w *blockFileWriter) open() error {
    file, err := os.OpenFile(w.filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
    if err != nil {
		return errors.Wrapf(err, "error opening block file writer for file %s", w.filePath)
	}
	w.file = file
	return nil
}

func (w *blockFileWriter) append(b []byte, sync bool) error {
    _, err := w.file.Write(b)
    if err != nil {
        return err
    }
    if sync {
        return w.file.Sync()
    }
    return nil
}

func (w *blockFileWriter) close() error{
	return errors.WithStack(w.file.Close())
}

type blockFileReader struct {
    file *os.File
}

func newBlockFileReader(filePath string) (*blockFileReader, error) {
    file, err := os.OpenFile(filePath, os.O_RDONLY, 0600)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening block file reader for file %s", filePath)
	}
	reader := &blockFileReader{file}
	return reader, nil
}

func (r *blockFileReader) read(offset int, length int) ([]byte, error) {
	b := make([]byte, length)
	// _, err := r.file.Read(b)
	_, err := r.file.ReadAt(b, int64(offset))
	if err != nil {
		return nil, errors.Wrapf(err, "error reading block file for offset %d and length %d", offset, length)
	}
	return b, nil
}

func (r *blockFileReader) readAll(length int) {
	b := make([]byte, length)
	var pos int64 = 0
	for {
		n, err := r.file.ReadAt(b, pos)
		if err != nil && err != io.EOF {
			log.Printf("error ReadAt")
		}
		if n == 0 {
			log.Printf("readAll: finish read")
			break
		}
		log.Printf("readAll: pos %v: %v and n=%v", pos, getBytesToInt(b), n)
		pos = pos + int64(8)
	}
}

func (r *blockFileReader) close() error {
	return errors.WithStack(r.file.Close())
}

// int-[]byte 解析
func getIntToBytes(n int) []byte {
    x := int64(n)
    bytesBuffer := bytes.NewBuffer([]byte{})
    binary.Write(bytesBuffer, binary.BigEndian, x)
    return bytesBuffer.Bytes()
}

func getBytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func main() {
	header := 123456
        
    filePath := "./ledger"
	fileWriter, err := newBlockFileWriter(filePath)
	if err != nil {
		errors.Wrapf(err, "write: error open file")
	}
    headerBytes := getIntToBytes(header)
	fileWriter.append(headerBytes, true)
	fileWriter.close()

	// log.Printf("%v", getBytesToInt(headerBytes))

	fileReader, err := newBlockFileReader(filePath)
	if err != nil {
		errors.Wrapf(err, "read: error open file")
	}

	// fileReader.readAll(6000)
	// TODO: bug!!!! 为什么当length为一个很大的数比如600时，会解析不出返回的结果，但是在read里可以解析
	b, err := fileReader.read(0, 8)
	if err != nil {
		errors.Wrapf(err, "read: error read file")
	}
	log.Printf("read header: %v", getBytesToInt(b))

	
	fileReader.close()    
}

