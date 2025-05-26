package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
)

func TestNewClients(t *testing.T) {
	nClients := 100000
	wg := sync.WaitGroup{}
	wg.Add(nClients)

	for i := 0; i < nClients; i++ {
		go func(it int) {
			c, err := New("localhost:5832")
			if err != nil {
				log.Fatal(err)
			}
			key := fmt.Sprintf("client_%d", it)
			value := fmt.Sprintf("client_x_%d", it)
			if err := c.Set(context.TODO(), key, value); err != nil {
				log.Fatal(err)
			}
			val, err := c.Get(context.TODO(), key)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Client %d got this val %s => \n", it, val)
			wg.Done()
		}(i)
	}
	wg.Wait()

}

// func TestNewClient(t *testing.T) {
// 	c, err := New("localhost:5832")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	time.Sleep(time.Second)
// 	for i := 0; i < 10; i++ {
// 		key := "fooo" + fmt.Sprint(i)
// 		fmt.Println("start SET\n ", key)
// 		if err := c.Set(context.TODO(), key, "bar"+fmt.Sprint(i)); err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("start GET\n ")

// 		val, err := c.Get(context.TODO(), key)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("GET", key, "=>", val)
// 	}
// }
