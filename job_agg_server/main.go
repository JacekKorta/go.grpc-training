package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "jobs"
)


type Server struct {
	pb.UnimplementedJobDataSrvServer
}

func (s *Server) Count(ctx context.Context, in *pb.StandardRequest) (*pb.CountResponse, error) {
	log.Printf("Receive message body from client: %s and %s", in.Category, in.ExpLevel)
	var offers Offers = ParseJsonFile("./data/job_collection.txt")
	var totalCount int32 = int32(offers.Len(in.ExpLevel))
	return &pb.CountResponse{TotalOfferCount: totalCount}, nil
}

func (s *Server) GetSalaryDataBySeniority(ctx context.Context, in *pb.StandardRequest) (*pb.SalaryResponse, error) {
	log.Printf("Receive message body from client: %s and %s", in.Category, in.ExpLevel)
	var offers Offers = ParseJsonFile("./data/job_collection.txt")
	min, max, avg := offers.GetSalaryBySeniorityLvl(in.ExpLevel)
	return &pb.SalaryResponse{Min: min, Max: max, Avg: avg}, nil
}


func main() {

	log.Println("Starting server")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()
	pb.RegisterJobDataSrvServer(grpcServer, &s)

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}