package client

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestNewClients(t *testing.T) {

	c, err := New("localhost:5832")
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Set(context.TODO(), "key", "45678"); err != nil {
		log.Fatal(err)
	}
	val, err := c.Get(context.TODO(), "key")
	if err != nil {
		log.Fatal(err)
	}
	n, _ := strconv.Atoi(val)
	fmt.Print(n)
	fmt.Printf("Client got this val => %s \n", val)

}

func TestNewClient(t *testing.T) {
	c, err := New("localhost:5832")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		key := "fooo" + fmt.Sprint(i)
		fmt.Println("start SET\n ", key)
		if err := c.Set(context.TODO(), key, "134"); err != nil {
			log.Fatal(err)
		}
		fmt.Println("start GET\n ")

		val, err := c.Get(context.TODO(), key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("GET", key, "=>", val)
	}
}
