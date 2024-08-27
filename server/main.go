package main

import (
	"context"
	"log/slog"
	"net"

	"github.com/yaninyzwitty/grpc-starter-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedCalculatorServer
}

func (s *Server) Add(ctx context.Context, in *pb.CalculationRequest) (*pb.CalcuationResponse, error) {
	return &pb.CalcuationResponse{Result: in.A + in.B}, nil
}

func (s *Server) Divide(ctx context.Context, in *pb.CalculationRequest) (*pb.CalcuationResponse, error) {
	if in.B == 0 {
		return nil, status.Error(codes.InvalidArgument, "cannot divide by zero")

	}
	return &pb.CalcuationResponse{Result: in.A / in.B}, nil
}

func (s *Server) Sum(ctx context.Context, in *pb.NumbersRequest) (*pb.CalcuationResponse, error) {
	var sum int64
	for _, n := range in.Numbers {
		sum += n
	}

	return &pb.CalcuationResponse{Result: sum}, nil

}

func (s *Server) Multiply(ctx context.Context, in *pb.MultiplicationRequest) (*pb.MultiplicationResponse, error) {
	if in.B == 0 {
		return nil, status.Error(codes.InvalidArgument, "cannot multiply by zero")

	}

	return &pb.MultiplicationResponse{Result: in.A * in.B}, nil

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		slog.Error("Failed to listen", "error", err)
	}
	s := grpc.NewServer()

	reflection.Register(s) //good for testing
	pb.RegisterCalculatorServer(s, &Server{})

	if err := s.Serve(listener); err != nil {
		slog.Error("Failed to serve", "error", err)
	}
}

// grpcurl -plaintext localhost:8080 list
// grpcurl -plaintext -d '{"a": 50, "b": 40}' localhost:8080 calculator.Calculator/Add
// grpcurl -plaintext -d '{"a": 10, "b": 5}' localhost:8080 calculator.Calculator/Add
