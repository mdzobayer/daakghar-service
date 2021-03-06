package api

import (
	"encoding/json"
	"net/http"

	"github.com/daakghar-service/data/api/conversation"
	"github.com/daakghar-service/data/api/filters"
	"github.com/daakghar-service/jwork"
	jconversation "github.com/daakghar-service/jwork/conversation"
	"github.com/pkg/errors"
)

// ConversationCreate handles conversation create
func ConversationCreate(user string) JHandler {
	return func(w http.ResponseWriter, r *http.Request) jwork.Worker {

		conversation := conversation.Create{}
		err := json.NewDecoder(r.Body).Decode(&conversation)
		if err != nil {
			return jwork.NewErr(errors.Wrap(err, "api.ConversationCreate, could not decode conversation create request"))
		}

		return jconversation.NewCreate(user, conversation)
	}
}

// ConversationGet handles get conversation
func ConversationGet(user string) JHandler {

	return func(w http.ResponseWriter, r *http.Request) jwork.Worker {
		id := filters.ID{}
		if err := id.Parse(r); err != nil {
			return jwork.NewErr(errors.Wrap(err, "api.ConversationGet, parse id param"))
		}

		return jconversation.NewPrivateRead(id)
	}
}
