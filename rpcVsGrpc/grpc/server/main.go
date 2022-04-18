package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "grpc/square"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 50052, "the server port")

type server struct {
	pb.UnimplementedCalculatorServer
}

type Result struct {
	Num, Ans int
}

func (c *server) Square(ctx context.Context, in *pb.SquareNumber) (*pb.SquareReply, error) {
	nums := in.GetNumber()
	nums = nums * nums
	return &pb.SquareReply{Number: nums}, nil
}

func (c *server) ListFeatures(in *pb.SquareNumber, stream pb.Calculator_ListFeaturesServer) error {
	out := []*pb.SquareReply{
		&pb.SquareReply{Number: 1},
		&pb.SquareReply{Number: 2},
		&pb.SquareReply{Number: 3},
	}

	for _, v := range out {
		err := stream.Send(v)

		if err != nil {
			return err
		}

		<-time.After(1 * time.Second)
	}

	return nil
}

func (c *server) Chat(stream pb.Calculator_ChatServer) error {
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.CheetRespond{
				Respond: true,
			})
		}

		if err != nil {
			return err
		}

		log.Printf(" userID : %v, time : %v , mag : %v \n", msg.GetUserID(), msg.GetSendTime(), msg.GetMsg())

	}
}

func (c *server) MapLog(stream pb.Calculator_MapLogServer) error {
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		data.Latitude += 1
		data.Longitude -= 100

		if err_send := stream.Send(data); err_send != nil {
			return err
		}
	}
}

func main() {
	flag.Parse()
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
