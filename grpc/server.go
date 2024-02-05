package grpc_server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"go-profiler/gopsutil"
	"go-profiler/grpc/helloworld"
	grpc_process "go-profiler/grpc/process"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	helloworld.UnimplementedGreeterServer
	grpc_process.UnimplementedProcessServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

func (s *server) GetProcessInfo(ctx context.Context, in *grpc_process.ProcessRequest) (*grpc_process.ProcessReply, error) {
	processes, err := gopsutil.GetProcessesInfo()
	if err != nil {
		return nil, err
	}

	pid := uint32(in.GetPid()) // Convert pid to uint32
	for _, p := range processes {
		if p.ProcessId == pid {
			return &grpc_process.ProcessReply{
				Name:     p.Name,
				CpuUsage: float32(p.CPUUsage),
				MemUsage: p.Memory,
				Pid:      int32(p.ProcessId),
				Ctime:    p.CreateTime,
				Time:     p.Timestamp,
			}, nil
		}
	}
	return nil, fmt.Errorf("process with pid %d not found", pid)
}

type GrpcServer struct {
}

func NewGreeterServer() *GrpcServer {
	return &GrpcServer{}
}

func (gs *GrpcServer) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	grpc_process.RegisterProcessServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
