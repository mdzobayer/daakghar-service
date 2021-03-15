package api

import (
	"net/http"

	"github.com/daakghar-service/jwork"
)

// JHandler handles json api
type JHandler func(http.ResponseWriter, *http.Request) jwork.Worker
