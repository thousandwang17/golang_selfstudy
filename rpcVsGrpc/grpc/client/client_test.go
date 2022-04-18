package main

import (
	"context"
	"flag"
	pb "grpc/square"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"testing"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (c *server) Square(ctx context.Context, in *pb.SquareNumber) (*pb.SquareReply, error) {
	nums := in.GetNumber()
	nums = nums * nums
	return &pb.SquareReply{Number: nums}, nil
}

func BenchmarkGrpc(t *testing.B) {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	t.ResetTimer()
	t.RunParallel(func(p *testing.PB) {
		for p.Next() {
			if _, err := client.Square(ctx, &pb.SquareNumber{Number: 12}); err != nil {
				log.Fatalf("Failed to call Cal.Square. %v", err)
			}
		}
	})
}

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestSquare(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)
	log.Printf("%v", 123)
	r, err2 := client.Square(ctx, &pb.SquareNumber{Number: 10})
	log.Printf("%v", 456)

	if err2 != nil {
		log.Fatalf("could not greet: %v", err2)
	}

	log.Printf("%v", r)
	// Test for output here.
}
