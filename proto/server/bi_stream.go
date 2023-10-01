package main

import (
	"io"
	pb "myeden.ai/m/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got req with name %v", req.Name)
		res := &pb.HelloResponse{
			Message: req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}