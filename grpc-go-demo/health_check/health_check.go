package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
	defer conn.Close()
	client := healthpb.NewHealthClient(conn)
	resp, _ := client.Check(context.Background(), &healthpb.HealthCheckRequest{}, grpc.FailFast(false))
	if resp.Status != healthpb.HealthCheckResponse_SERVING {
		print("Failed to check health!")
	} else {
		println("Health ok!")
	}
}
