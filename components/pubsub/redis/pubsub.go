package redispubsub

import (
	"context"
	"cs_chat_app_server/components/pubsub"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type RedisPubSub struct {
	client *redis.Client
}

func NewRedisPubSub(client *redis.Client) *RedisPubSub {
	return &RedisPubSub{client: client}
}

func (ps *RedisPubSub) Publish(ctx context.Context, topic pubsub.Topic, data string) error {
	marshaled, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ps.client.Publish(ctx, string(topic), marshaled).Err()
}

func (ps *RedisPubSub) Subscribe(ctx context.Context, topic pubsub.Topic) <-chan string {

	c := make(chan string)

	_pubsub := ps.client.Subscribe(ctx, string(topic))
	ch := _pubsub.Channel()
	go func() {
		for msg := range ch {
			var data string
			err := json.Unmarshal([]byte(msg.Payload), &data)
			if err != nil {
				continue
			}
			c <- data
		}
	}()

	return c
}
