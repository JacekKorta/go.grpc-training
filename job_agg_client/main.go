package main

import (
	"log"

	pb "jobs"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewJobDataSrvClient(conn)

	response, err := c.GetSalaryDataBySeniority(context.Background(), &pb.StandardRequest{Category: "python", ExpLevel: "mid"})
	if err != nil {
		log.Fatalf("Error when calling Count: %s", err)
	}
	log.Printf("Response from server: min %d | max: %d | avg %v", response.Min, response.Max, response.Avg)

}