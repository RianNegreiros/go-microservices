package main

import (
	"context"
	"log"
	"time"

	"github.com/RianNegreiros/go-microservices/logger-service/data"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logger").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error wrinting to mmongo: ", err)
		return err
	}

	*resp = "Processed payload: " + payload.Name + " - " + payload.Data
	return nil
}
