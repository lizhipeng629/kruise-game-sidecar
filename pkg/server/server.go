package server

import (
	"context"
	"log"
	"net"

	pb "github.com/magicsong/kidecar/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// server is used to implement pb.SDKServer.
type server struct {
	pb.UnimplementedSDKServer
}

// Ready implements pb.SDKServer
func (s *server) Ready(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	log.Println("Ready called")
	return &emptypb.Empty{}, nil
}

// Shutdown implements pb.SDKServer
func (s *server) Shutdown(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	log.Println("Shutdown called")
	return &emptypb.Empty{}, nil
}

// SetLabel implements pb.SDKServer
func (s *server) SetLabel(ctx context.Context, in *pb.SetLabelRequest) (*emptypb.Empty, error) {
	log.Println("SetLabel called")
	return &emptypb.Empty{}, nil
}

// GetLabel implements pb.SDKServer
func (s *server) GetLabel(ctx context.Context, in *pb.GetLabelRequest) (*pb.GetLabelResponse, error) {
	log.Println("GetLabel called")
	return &pb.GetLabelResponse{}, nil
}

// SetAnnotation implements pb.SDKServer
func (s *server) SetAnnotation(ctx context.Context, in *pb.SetAnnotationRequest) (*emptypb.Empty, error) {
	log.Println("SetAnnotation called")
	return &emptypb.Empty{}, nil
}

// GetAnnotation implements pb.SDKServer
func (s *server) GetAnnotation(ctx context.Context, in *pb.GetAnnotationRequest) (*pb.GetAnnotationResponse, error) {
	log.Println("GetAnnotation called")
	return &pb.GetAnnotationResponse{}, nil
}

// Allocate implements pb.SDKServer
func (s *server) Allocate(ctx context.Context, in *pb.AllocateRequest) (*emptypb.Empty, error) {
	log.Println("Allocate called")
	return &emptypb.Empty{}, nil
}

// Reserve implements pb.SDKServer
func (s *server) Reserve(ctx context.Context, in *pb.ReserveRequest) (*emptypb.Empty, error) {
	log.Println("Reserve called")
	return &emptypb.Empty{}, nil
}

// SetCapacity implements pb.SDKServer
func (s *server) SetCapacity(ctx context.Context, in *pb.SetCapacityRequest) (*emptypb.Empty, error) {
	log.Println("SetCapacity called")
	return &emptypb.Empty{}, nil
}

// StartGRPCServer ...
func StartGRPCServer(sig chan struct{}) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSDKServer(s, &server{})
	log.Println("Server is running on port 50051")

	// Serve in a goroutine
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for signal
	<-sig

	// Stop the server
	s.GracefulStop()
	log.Println("Server stopped")
}
