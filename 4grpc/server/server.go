package main

import (
    "log"
    "net"

    "golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/grpc/grpc-go/examples/rain/block"
)

const (
    port = "41005"
)

type server struct{
    pb.UnimplementedBlockServer
}

func (t *server) SaveBlockHeader(ctx context.Context, request *pb.BlockHeaderSaveReq) (response *pb.BlockHeaderSaveRes, err error) {
    log.Printf("Received: %v", request.Header)
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

