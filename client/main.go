package main

import (
	"context"
	// "crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jicodes/go-grpc/pb"
)

func main() {
	serverAddr := flag.String(
		"server_addr", "localhost:8080", 
		"The server address in the format of host:port",
	)
	flag.Parse()

	// Use tls for production
	// creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})
	// opts := []grpc.DialOption{
	// 	grpc.WithTransportCredentials(creds),
	// }

	// Disable tls for local testing
	opts := []grpc.DialOption{
    grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *serverAddr, opts...)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	res, err := client.Add(ctx, &pb.CalculationRequest{A: 3, B: 2})

	if err != nil {
		log.Fatalln("Error sending request:", err)
	}

	fmt.Println("Add result:", res.Result)

}