package main

import (
	pb "Exercise/RPC/sensitiveWord/wordfilter"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"Exercise/RPC/sensitiveWord/server"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWordFilterServiceServer(s, &server.Router{})

	go func(){
		log.Printf("GRPC server listen %s ",":8080")
		if err := s.Serve(listen); err != nil { //启动服务器
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit:=make(chan os.Signal)
	signal.Reset(os.Interrupt,syscall.SIGTERM,syscall.SIGKILL)
	signal.Notify(quit,os.Interrupt,syscall.SIGTERM,syscall.SIGKILL)
	<-quit


}
