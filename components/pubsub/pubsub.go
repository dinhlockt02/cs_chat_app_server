package pubsub

import "context"

type Topic string

type PubSub interface {
	Publish(ctx context.Context, topic Topic, data string) error
	Subscribe(ctx context.Context, topic Topic) <-chan string
}
