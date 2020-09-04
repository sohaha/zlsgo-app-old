package redis

import (
	"context"

	"app/global"
)

// Subscribe 订阅
func Subscribe(handle func(channel, payload string), channels ...string) error {
	rdb := global.Redis
	pubsub := rdb.Subscribe(context.Background(), channels...)
	ch := pubsub.Channel()
	go func() {
		for m := range ch {
			handle(m.Channel, m.Payload)
		}
	}()
	return nil
}

// Publish 发布
func Publish(channel string, message interface{}) error {
	return global.Redis.Publish(context.Background(), channel, message).Err()
}
