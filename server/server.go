package server

import (
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"net"
	"net/http"
	"rabotyaga-go-backend/structures"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/utils"
	"strconv"
	"time"
)

type OnCallbackFunc = func(conn net.Conn, op ws.OpCode, sign *vkapps.Params, data json.RawMessage)

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
		fmt.Printf("Connection at [%s] \r\n", time.Now())
		validate, err := vkapps.ParamsVerify(r.URL.String(), "8PogTmDn5uru9WPdXuup")
		if !validate {
			http.Error(w, fmt.Sprintf("Invalid request: %v", err), http.StatusUnauthorized)
		}

		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid request: %v", err), http.StatusUnauthorized)
			return
		}

		vkParams, err := vkapps.NewParams(r.URL)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid request: %v", err), http.StatusUnauthorized)
			return
		}

		vkTs, err := strconv.ParseInt(vkParams.VkTs, 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid request: %v", err), http.StatusUnauthorized)
			return
		}

		if time.Now().Sub(time.Unix(vkTs, 0)) <= 30*time.Minute {
			http.Error(w, fmt.Sprintf("Invalid request: token is rotten"), http.StatusUnauthorized)
			return
		}

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
						fc(conn, op, vkParams, message.Data)
					}
				}
			}
		}()
	}))

	if err != nil {
		log.Panicln("Server error starting")
	}
}

func (s *Server) OnSocket(e types.EventType, cb OnCallbackFunc) {
	if _, ok := s.events[e]; !ok {
		s.events[e] = []OnCallbackFunc{}
	}

	s.events[e] = append(s.events[e], cb)
}
