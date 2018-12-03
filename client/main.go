package main

import (
	pb "Exercise/RPC/sensitiveWord/wordfilter"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func main() {
	connection, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()

	c := pb.NewWordFilterServiceClient(connection)

	text := "淫荡"

	if len(os.Args) > 1 {
		log.Println(os.Args[0])
		text = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.ContainsSensitiveWord(ctx, &pb.ContainSensitiveWordRequest{Text: text})
	if err != nil {
		log.Fatalf("could not test: %v", err)
	}
	log.Printf("ContainsSentisiveWord : %v ", r.IsSensitiveWord)
}
