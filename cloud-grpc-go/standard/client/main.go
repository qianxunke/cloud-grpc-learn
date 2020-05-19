package main

import (
	pdf "cloud_grpc_go/pdf"
	"context"
	"google.golang.org/grpc"
	"time"

	"log"
)

func main() {

	con, err := grpc.Dial("localhost:8488", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer con.Close()
	c := pdf.NewHelloClient(con)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pdf.HelloRequest{Name: "Go grpc client"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}
