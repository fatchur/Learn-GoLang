package main

import (
	"context"
	"encoding/json"

	//"gcp-pubsub-mock/models"
	"testing"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/pubsub/pstest"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

const (
	topicID = "test-fruit-topic"
)

type FruitInterface struct {
	Page   int
	Fruits []string
}

func PublishMessage(ctx context.Context, topic *pubsub.Topic, fruit *FruitInterface) error {
	data, err := json.Marshal(fruit)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = topic.Publish(ctx, &pubsub.Message{Data: data}).Get(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func Test_PublishMessage(t *testing.T) {
	srv := pstest.NewServer()
	defer srv.Close()

	conn, err := grpc.Dial(srv.Addr, grpc.WithInsecure())
	if !assert.NoError(t, err) {
		return
	}

	defer conn.Close()

	ctx := context.Background()
	projectID := uuid.New().String()
	client, err := pubsub.NewClient(ctx, projectID, option.WithGRPCConn(conn))
	if !assert.NoError(t, err) {
		return
	}

	_, err = client.CreateTopic(ctx, topicID)
	if !assert.NoError(t, err) {
		return
	}

	fruit := &FruitInterface{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	topic := client.Topic(topicID)
	err = PublishMessage(ctx, topic, fruit)

	assert.NoError(t, err)
}
