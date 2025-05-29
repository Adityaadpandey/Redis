package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/adityaadpandey/Redis/client"
)

func TestFooBar(t *testing.T) {
	in := map[string]string{
		"first":  "1",
		"second": "2",
	}
	out := respWriteMap(in)
	fmt.Println(out)
}

func TestNewClients(t *testing.T) {
	server := NewServer(Config{})
	go func() {
		log.Fatal(server.Start())
	}()
	time.Sleep(1 * time.Second) // Give the server some time to start

	nClients := 10
	wg := sync.WaitGroup{}
	wg.Add(nClients)

	for i := 0; i < nClients; i++ {
		go func(it int) {
			c, err := client.New("localhost:5832")
			if err != nil {
				log.Fatal(err)
			}
			defer c.Close()
			key := fmt.Sprintf("client__%d", it)
			// value := fmt.Sprintf("client_x_%d", it)
			if err := c.Set(context.TODO(), key, "123"); err != nil {
				log.Fatal(err)
			}
			val, err := c.Get(context.TODO(), key)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Client %d got this val => %s \n", it, val)
			wg.Done()
		}(i)

	}
	wg.Wait()
	time.Sleep(1 * time.Second) // Wait for all clients to finish
	if len(server.peers) != 0 {
		t.Fatalf("Expected 0 peers, but got %d", len(server.peers))
	}

}
