package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mskKandula/GolangPython/WineVarieties/pb"

	"google.golang.org/grpc"
)

var client pb.WineClassifierServiceClient

func main() {
	addr := "serverservice:9000"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	client = pb.NewWineClassifierServiceClient(conn)

	http.HandleFunc("/predict", predictHandler)
	fmt.Println("listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

//Handler returns http respone in string format.
func predictHandler(w http.ResponseWriter, r *http.Request) {
	req := pb.WineReviewRequest{
		Review: "Dry with a fruity aroma",
	}

	resp, err := client.GetWineVariety(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Predicted wine variety: %v", resp.Variety)
	log.Printf("Predicted wine variety: %v", resp.Variety)
}
