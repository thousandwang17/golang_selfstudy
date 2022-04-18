package main

import (
	"context"
	"flag"
	pb "grpc/square"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("addr", "localhost:50052", "The server address in the format of host:port")
)

//  simple
func simple_grpc(client pb.CalculatorClient, number *pb.SquareNumber) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err2 := client.Square(ctx, number)

	if err2 != nil {
		log.Fatalf("could not greet: %v", err2)
	}

	log.Printf("%v", r)
}

// serve stream
func server_stream(client pb.CalculatorClient, number *pb.SquareNumber) {

	ctx_lsit, cancel_list := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel_list()

	stream, errList := client.ListFeatures(ctx_lsit, number)

	if errList != nil {
		log.Fatalf("could not greet: %v", errList)
	}

	for {
		data, err := stream.Recv()

		if err != nil {
			log.Printf("%v", err)
			break
		}

		num := data.GetNumber()
		log.Printf("%v", num)
	}
}

// client stream
func chat_stream(client pb.CalculatorClient, msg []*pb.CheetMsg) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.Chat(ctx)

	if err != nil {
		log.Fatalf("chat connect err %v", err)
	}

	for _, v := range msg {
		if err := stream.Send(v); err != nil {
			log.Printf(" chat stream err :%v", err)
			break
		}
		<-time.After(time.Second)
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("chat stream close err : %v", err)
		return
	}

	log.Printf("%v", reply)
}

func Map_stream(client pb.CalculatorClient, mapData []*pb.MapLogDate) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.MapLog(ctx)

	if err != nil {
		log.Fatalf("Map connect err %v", err)
	}

	done := make(chan struct{})
	go func() {
		for {
			data, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}

			if err != nil {
				log.Printf("Map stream  err : %v", err)
				continue
			}
			log.Printf("Data : %v", data)
		}
	}()

	for _, v := range mapData {
		if err := stream.Send(v); err != nil {
			log.Printf("%v", err)
		}
	}

	stream.CloseSend()
	<-done

}

func main() {
	// var opts []grpc.DialOption
	// opts = append(opts, grpc.WithInsecure())
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	simple_grpc(client, &pb.SquareNumber{Number: 10})

	server_stream(client, &pb.SquareNumber{Number: 11})

	chat_stream(client, []*pb.CheetMsg{
		&pb.CheetMsg{UserID: 1, SendTime: 1648700757, Msg: "123"},
		&pb.CheetMsg{UserID: 2, SendTime: 1648700857, Msg: "456"},
		&pb.CheetMsg{UserID: 3, SendTime: 1648700957, Msg: "789"},
	})

	Map_stream(client, []*pb.MapLogDate{
		&pb.MapLogDate{Latitude: 1000, Longitude: 2000},
		&pb.MapLogDate{Latitude: 1100, Longitude: 2400},
		&pb.MapLogDate{Latitude: 1200, Longitude: 2300},
	})
}
