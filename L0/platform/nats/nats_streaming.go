package nats_s

import (
    "context"
    "fmt"
    "task-l0/pkg/configs"
    "github.com/nats-io/stan.go"
)

type Client interface {
    Publish(subject string, data []byte) error
    Subscribe(subject string, cb stan.MsgHandler, opts ...stan.SubscriptionOption) (stan.Subscription, error)
}

type natsClient struct {
    conn stan.Conn
}

func (nc *natsClient) Publish(subject string, data []byte) error {
    return nc.conn.Publish(subject, data)
}

func (nc *natsClient) Subscribe(subject string, cb stan.MsgHandler, opts ...stan.SubscriptionOption) (stan.Subscription, error) {
    return nc.conn.Subscribe(subject, cb, opts...)
}

func NewClient(ctx context.Context, cfg *configs.Config) (Client, error) {
    sc, err := stan.Connect(
        cfg.Nats.ClusterID, cfg.Nats.ClientID, stan.Pings(1, 5), stan.NatsURL(cfg.Nats.Url),
    )
    if err != nil {
        fmt.Printf("Nats stream: %v\n", err)
        return nil, err
    }

    return &natsClient{conn: sc}, nil
}