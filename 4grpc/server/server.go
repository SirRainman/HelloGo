package main

import (
    "log"
    "net"
    "os"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "github.com/pkg/errors"
	pb "github.com/grpc/grpc-go/examples/rain/block"
)

const (
    port = "41005"
)

type server struct{
    pb.UnimplementedBlockServer
}

type blockFileWriter struct {
    filePath string
    file *os.File
}

func newBlockFileWriter(filePath string) (*blockFileWriter, error){
    writer := &blockFileWriter{filePath: filePath}
    return writer, writer.open()
}

func (w *blockFileWriter) open() error {
    file, err = os.OpenFile(w.filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
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

func (w *blockfileWriter) close() error {
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
	reader := &blockfileReader{file}
	return reader, nil
}

func (r *blockfileReader) read(offset int, length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := r.file.ReadAt(b, int64(offset))
	if err != nil {
		return nil, errors.Wrapf(err, "error reading block file for offset %d and length %d", offset, length)
	}
	return b, nil
}

func (r *blockfileReader) close() error {
	return errors.WithStack(r.file.Close())
}






func getHeaderBytes(header int) []byte {
    x := int32(header)

    bytesBuffer := bytes.NewBuffer([]byte{})
    binary.Write(bytesBuffer, binary.BigEndian, x)
    return bytesBuffer.Bytes()
}







func (t *server) SaveBlockHeader(ctx context.Context, request *pb.BlockHeaderSaveReq) (response *pb.BlockHeaderSaveRes, err error) {
    // 存到账本里
    log.Printf("Received: %v", request.Header)
    
    filePath := "./ledger.txt"
    fileWriter := newBlockFileWriter(filePath)
    headerBytes = getHeaderBytes(request.Header)
    fileWriter.append(headerBytes)
    fileWriter.close()

    // 返回执行结果
    response = &pb.BlockHeaderSaveRes{
        Res: "success",
    }
    return response, err
}

func main() {
    lis, err := net.Listen("tcp", ":"+port) 
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()   
    pb.RegisterBlockServer(s, &server{})
    if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
    log.Println("grpc server in: %s", port)
}

