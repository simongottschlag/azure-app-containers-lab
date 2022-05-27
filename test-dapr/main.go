package main

import (
	"context"
	"fmt"
	"os"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "application returned an error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	client, err := newDaprClient(50001)
	if err != nil {
		return err
	}

	ctx := context.Background()
	data := []byte("hello")
	store := "blob"

	if err := client.SaveState(ctx, store, "key1", data, map[string]string{}); err != nil {
		return err
	}

	item, err := client.GetState(ctx, store, "key1", map[string]string{})
	if err != nil {
		return err
	}

	fmt.Printf("data [key:%s etag:%s]: %s\n", item.Key, item.Etag, string(item.Value))

	if err := client.DeleteState(ctx, store, "key1", map[string]string{}); err != nil {
		return err
	}

	return nil
}

func newDaprClient(port int) (dapr.Client, error) {
	var client dapr.Client
	var err error
	for i := 0; i < 10; i++ {
		client, err = dapr.NewClientWithPort(fmt.Sprintf("%d", port))
		if err == nil {
			return client, nil
		}
		fmt.Fprintf(os.Stderr, "unable to initialize Dapr client, sleeping for 1 second\n")
		time.Sleep(1 * time.Second)
	}

	return nil, err
}
