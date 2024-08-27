package main

import (
	"context"
	"crypto/tls"
	"flag"
	"log/slog"
	"time"

	"github.com/yaninyzwitty/grpc-starter-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	serverAddress := flag.String("server", "localhost:8080", "The server address in the format of host:port")
	flag.Parse()
	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})

	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, *serverAddress, opts...)
	if err != nil {
		slog.Error("Failed to dial server", "error", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	res, err := client.Sum(ctx, &pb.NumbersRequest{
		Numbers: []int64{10, 10, 10, 10, 10},
	})
	if err != nil {
		slog.Error("error sending request", "error", err)
	}

	slog.Info("Response from server", "result", res.Result)
	// Output: Response from server: result: 50"

}
