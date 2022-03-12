package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	pb "github.com/ahmetsoykan/grpc-tutorials/text-to-speech/api"
	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	output := flag.String("o", "output.wav", "wav file where the output will be written")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("please pass an argument")
		os.Exit(1)
	}

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connecto to %s: %v", *backend, err)
	}

	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)

	flag.Args()
	text := pb.Text{Text: strings.Join(flag.Args(), " ")}
	res, err := client.Say(context.Background(), &text)
	if err != nil {
		log.Fatalf("could not say %s: %v", text.Text, err)
	}

	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("could not write to %s: %v", *output, err)
	}

}
