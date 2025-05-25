package main

import (
	"testing"
)

func TestProtocol(t *testing.T) {
	raw := "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"

	cmd, err := parseCommand(raw)
	if err != nil {
		t.Fatalf("failed to parse command: %v", err)
	}
	if setCmd, ok := cmd.(SetCommand); ok {
		if string(setCmd.key) != "foo" || string(setCmd.val) != "bar" {
			t.Errorf("expected key=foo, val=bar; got key=%s, val=%s", string(setCmd.key), string(setCmd.val))
		}
	} else {
		t.Fatalf("expected SetCommand, got %T", cmd)
	}
}
