package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"

	"github.com/tidwall/resp"
)

type Client struct {
	addr string
	conn net.Conn
}

func New(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

func (c *Client) Set(ctx context.Context, key, val string) error {
	if c.conn == nil {
		conn, err := net.Dial("tcp", c.addr)
		if err != nil {
			return err
		}
		c.conn = conn
	}

	buf := &bytes.Buffer{}
	wr := resp.NewWriter(buf)
	wr.WriteArray([]resp.Value{
		resp.StringValue("SET"),
		resp.StringValue((key)),
		resp.StringValue(val),
	})

	fmt.Printf("%s", buf.String())
	// conn.Write(buf.Bytes())
	_, err := io.Copy(c.conn, buf)
	return err
}
