package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/indes/telepush/internal/bot"
)

type PushRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"u,omitempty"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	userId := 0
	var msg string

	if u := params.Get("u"); u != "" {
		userId, _ = strconv.Atoi(u)
	}

	if r.Method == http.MethodPost {
		var pushReq PushRequest
		if err := json.NewDecoder(r.Body).Decode(&pushReq); err == nil {
			if userId == 0 {
				userId = pushReq.UserId
			}
			if pushReq.Title != "" && pushReq.Body != "" {
				msg = pushReq.Title + "\n" + pushReq.Body
			} else if pushReq.Title != "" {
				msg = pushReq.Title
			} else if pushReq.Body != "" {
				msg = pushReq.Body
			}
		}
	}

	if msg == "" {
		msg = params.Get("m")
	}
	if userId == 0 {
		userId, _ = strconv.Atoi(params.Get("u"))
	}

	if err := bot.NotifyTxtMessage(userId, msg); err != nil {
		fmt.Fprintf(w, "send message to %d failed!", userId)
		return
	}

	fmt.Fprintf(w, "send message to %d succes!", userId)
}
