package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"time"

	"github.com/adityaadpandey/Redis/client"
)

const defaultListenAddr = ":5832"

type Config struct {
	ListenAddr string
}

type Server struct {
	Config
	peers     map[*Peer]bool
	ln        net.Listener
	addPeerCh chan *Peer
	quitCh    chan struct{}
	msgCh     chan []byte

	//
	kv *KV
}

func NewServer(cfg Config) *Server {
	if len(cfg.ListenAddr) == 0 {
		cfg.ListenAddr = defaultListenAddr
	}
	return &Server{
		Config:    cfg,
		peers:     make(map[*Peer]bool),
		addPeerCh: make(chan *Peer),
		quitCh:    make(chan struct{}),
		msgCh:     make(chan []byte),
		kv:        NewKV(),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.ln = ln
	go s.loop()
	slog.Info("server started", "listen_addr", s.ListenAddr)
	return s.acceptLoop()
}

func (s *Server) handleRawMessage(rawMsg []byte) error {
	cmd, err := parseCommand(string(rawMsg))
	if err != nil {
		slog.Error("parse command error", "error", err, "raw_msg", string(rawMsg))
		return err
	}
	switch v := cmd.(type) {
	case SetCommand:
		slog.Info("Received SET command", "key", v.key, "value", v.val)
		return s.kv.Set(v.key, v.val)

	}
	return nil
}

func (s *Server) loop() {
	for {
		select {
		case rawMsg := <-s.msgCh:
			if err := s.handleRawMessage(rawMsg); err != nil {
				slog.Error("handle raw message error", "error", err)
			}
		case <-s.quitCh:
			return
		case peer := <-s.addPeerCh:
			s.peers[peer] = true
			slog.Info("peer added", "peer", peer)

		}
	}
}

func (s *Server) acceptLoop() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			slog.Error("accept error", "error", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	peer := NewPeer(conn, s.msgCh)
	s.addPeerCh <- peer
	slog.Info("new connection", "peer", peer, "remote_addr", conn.RemoteAddr())
	if err := peer.readLoop(); err != nil {
		slog.Error("peer read error", "error", err, "remoteAdd", conn.RemoteAddr())
	}
}

func main() {
	server := NewServer(Config{})
	go func() {
		log.Fatal(server.Start())
	}()
	time.Sleep(time.Second * 2) // Wait for server to start

	for i := 0; i < 10; i++ {
		client := client.New("localhost:5832")
		key := "fooo" + fmt.Sprint(i)
		if err := client.Set(context.TODO(), key, "bar"); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(server.kv.data)
	fmt.Println("All SET commands sent successfully")

	select {}
}
