package net

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
}

func (s *Server) Run() {
	listen, err := net.Listen("tcp", ":1122")
	if err != nil {
		panic(fmt.Sprint("Listen: ", err))
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			log.Println("new conn from ", conn.RemoteAddr().String())
			p := Package{}
			_, err := p.ReadFrom(conn)
			if err != nil {
				log.Println("ReadFrom: ", err)
				return
			} else {
				switch p.Cmd {
				case One:
					log.Println("server cmd: One")
					p := NewPackage(p.Cmd, User{Id: 1, Name: "zing"})
					_, err := p.WriteTo(conn)
					if err != nil {
						log.Println("One WriteTo: ", err)
					}
				case List:
					users := make([]User, 100000)
					for i := 0; i < 100000; i++ {
						users[i] = User{Id: i + 1, Name: fmt.Sprintf("name-%d", i+1)}
					}
					p := NewPackage(List, users)
					_, err := p.WriteTo(conn)
					if err != nil {
						log.Println("List WriteTo: ", err)
					}
				}
			}
		}(conn)
	}
}
