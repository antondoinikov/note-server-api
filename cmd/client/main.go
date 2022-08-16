package main

import (
	"context"
	desc "github.com/anton7191/Grpc_Note/pkg/note_v1"
	"google.golang.org/grpc"
	"log"
)

const address = "localhost:2406"

func main() {

	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()
	client := desc.NewNoteV1Client(con)
	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "First",
		Text:   "Help me!",
		Author: "Anton",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println("Id: ", res.Id)
}
