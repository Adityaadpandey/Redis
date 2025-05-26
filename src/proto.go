package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/tidwall/resp"
)

const (
	CommandSet   = "SET"
	CommandGet   = "GET"
	CommandHello = "HELLO"
)

type Command interface {
}

type SetCommand struct {
	key, val []byte
}

type GetCommand struct {
	key []byte
}
type HelloCommand struct {
	value string
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
		// fmt.Printf("Read %s\n", v.Type())
		if v.Type() == resp.Array {
			for _, value := range v.Array() {
				switch value.String() {
				case CommandGet:
					if len(v.Array()) != 2 {
						return nil, fmt.Errorf("invalid GET command: expected 2 arguments, got %d", len(v.Array()))
					}
					cmd := GetCommand{
						key: v.Array()[1].Bytes(),
					}
					return cmd, nil
				case CommandSet:
					if len(v.Array()) != 3 {
						return nil, fmt.Errorf("invalid SET command: expected 3 arguments, got %d", len(v.Array()))
					}
					cmd := SetCommand{
						key: v.Array()[1].Bytes(),
						val: v.Array()[2].Bytes(),
					}
					return cmd, nil
				case CommandHello:
					cmd := HelloCommand{
						value: v.Array()[1].String(),
					}
					return cmd, nil
				default:
				}
				// fmt.Printf("  #%d %s, value: '%s'\n", i, v.Type(), v)
			}
		}
	}

	return nil, fmt.Errorf("invalid or unknow command Received: %s", raw)
}

func respWriteMap(m map[string]string) string {
	buf := bytes.Buffer{}
	buf.WriteString("%" + fmt.Sprintf("%d\r\n", len(m)))
	for k, v := range m {
		buf.WriteString(fmt.Sprintf("+%s\r\n", k))
		buf.WriteString(fmt.Sprintf("+%s\r\n", v))
	}
	return buf.String()
}
