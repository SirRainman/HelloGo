package main

import (
    "log"
    "runtime"
    "sync"
	"time"
	"math/rand"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/grpc/grpc-go/examples/rain/block"
)

var (
    wg sync.WaitGroup   
)

const (
    networkType = "tcp"
    address     = "localhost:41005"
    parallel    = 1  //连接并行度
    times       = 3  //每连接请求次数
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    currTime := time.Now()

    //并行请求
    for i := 0; i < int(parallel); i++ {
        wg.Add(1)
        go func() {
			defer wg.Done()
			// TODO: 为什么这样不行
			// client := getClient()

			defer func() {log.Printf("routine %v", i)}()

			conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			client := pb.NewBlockClient(conn)
			
			for i := 0; i < int(times); i++ {
				SaveBlockHeader(client)
			}
        }()
    }
    wg.Wait()
	
    log.Printf("time taken: %.2f ", time.Now().Sub(currTime).Seconds())
}

// TODO:bug 为什么这样不行？
func getClient() pb.BlockClient {
    //建立连接
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	return pb.NewBlockClient(conn)
}

func SaveBlockHeader(client pb.BlockClient) {
    var request pb.BlockHeaderSaveReq
    r := rand.Intn(123456)
    request.Header = int32(r);

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.SaveBlockHeader(ctx, &request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

    //判断返回结果是否正确
    log.Printf("response :%#v", response.Res)
}