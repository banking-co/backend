package utils

import (
	"fmt"
	"github.com/gobwas/ws"
	"log"
	"net/http"
	"time"
)

func SendError(w http.ResponseWriter, reason string, code uint16) {
	err := ws.RejectConnectionError(
		ws.RejectionReason(reason),
		ws.RejectionStatus(int(code)),
	)

	w.WriteHeader(int(code))
	if _, writeErr := w.Write([]byte(err.Error())); writeErr != nil {
		log.Printf("Error writing response: %v", writeErr)
	}

	fmt.Println(time.Now(), reason, code)
}
