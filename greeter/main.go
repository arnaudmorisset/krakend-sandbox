package main

import (
	"context"
	"log"
	"net"

	pb "github.com/arnaudmorisset/krakend-sandbox/greeter/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type greeterService struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("received: %v", in.Name)
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

type userService struct {
	pb.UnimplementedUserServer
}

func (s *userService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if in.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id wasn't given")
	}

	log.Printf("user_service: received request for id: %v", in.GetId())
	return &pb.GetUserResponse{Name: "Arnaud"}, nil
}

type profileService struct {
	pb.UnimplementedProfileServer
}

func (s *profileService) GetProfile(ctx context.Context, in *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	if in.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id wasn't given")
	}

	log.Printf("profile_service: received request for id: %v", in.GetId())
	return &pb.GetProfileResponse{JobTitle: "Software Engineer", Age: 29}, nil
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := grpc.MaxRecvMsgSize(577659245)
	s := grpc.NewServer(opts)
	pb.RegisterGreeterServer(s, &greeterService{})
	pb.RegisterUserServer(s, &userService{})
	pb.RegisterProfileServer(s, &profileService{})
	log.Printf("server listening at %v", l.Addr())
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
