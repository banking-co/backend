package server

import (
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"net"
	"net/http"
	"rabotyaga-go-backend/structures"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/utils"
	"time"
)

type OnCallbackFunc = func(conn net.Conn, op ws.OpCode, data json.RawMessage)

type Server struct {
	events map[types.EventType][]OnCallbackFunc
}

func Init() *Server {
	s := Server{
		events: make(map[types.EventType][]OnCallbackFunc),
	}

	fmt.Println("Server started!")
	return &s
}

func (s *Server) Listen() {
	err := http.ListenAndServe(":3001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Connection")
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Panicln("Upgrade HTTP error")
			return
		}

		go func() {
			defer func() {
				if err := conn.Close(); err != nil {
					logInfo := fmt.Sprintf("[ %d ]: connection closed, address: %s", time.Now().Unix(), conn.RemoteAddr())
					fmt.Println(logInfo)
				}
			}()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					break
				}

				message, err := utils.UnmarshalData[structures.EventParams](msg)
				if err != nil {
					break
				}

				if cbs, ok := s.events[message.Event]; ok {
					for _, fc := range cbs {
						fc(conn, op, message.Data)
					}
				}
			}
		}()
	}))

	if err != nil {
		log.Panicln("Server error starting")
	}
}

func (s *Server) On(e types.EventType, cb OnCallbackFunc) {
	if _, ok := s.events[e]; !ok {
		s.events[e] = []OnCallbackFunc{}
	}

	s.events[e] = append(s.events[e], cb)
}
