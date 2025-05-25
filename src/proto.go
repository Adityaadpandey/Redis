package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/tidwall/resp"
)

const (
	CommandSet = "SET"
)

type Command interface {
}

type SetCommand struct {
	key, val []byte
}

func parseCommand(raw string) (Command, error) {
	rd := resp.NewReader(bytes.NewBufferString(raw))
	for {
		v, _, err := rd.ReadValue()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Read %s\n", v.Type())
		if v.Type() == resp.Array {
			for i, value := range v.Array() {
				switch value.String() {
				case CommandSet:
					if len(v.Array()) != 3 {
						return nil, fmt.Errorf("invalid SET command: expected 3 arguments, got %d", len(v.Array()))
					}
					cmd := SetCommand{
						key: v.Array()[1].Bytes(),
						val: v.Array()[2].Bytes(),
					}
					return cmd, nil
				default:
				}
				fmt.Printf("  #%d %s, value: '%s'\n", i, v.Type(), v)
			}
		}
	}

	return nil, fmt.Errorf("invalid or unknow command Received: %s", raw)
}
