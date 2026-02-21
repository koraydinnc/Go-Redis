package server

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

type TCPServer struct {
	addr string
}

func NewTCPServer(addr string) *TCPServer {
	return &TCPServer{addr: addr}
}

func (s *TCPServer) Start() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	log.Println("MiniRedis sunucusu başlatıldı")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Client bağlanırken hata oluştu:", err)
			continue
		}

		connectedClient := conn.RemoteAddr().String()
		log.Printf("Client bağlandı: %s", connectedClient)

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("Client bağlantısı kapatıldı")
				_, werr := conn.Write([]byte("+KAPATILDI\r\n"))
				if werr != nil {
					log.Println("Client'e yazarken hata oluştu:", werr)
				}
				return
			}

			log.Println("Client okunurken hata oluştu:", err)
			return
		}

		msg := strings.TrimSpace(line)
		log.Printf("Client gonderdi: %q | from=%s", msg, conn.RemoteAddr().String())

		_, werr := conn.Write([]byte("BAĞLANTI KURULDU\r\n"))
		if werr != nil {
			log.Println("Client'e yazarken hata oluştu:", werr)
			return
		}

		log.Println("Client'e +OK gönderildi")
	}
}
