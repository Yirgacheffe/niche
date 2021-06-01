package main

import (
	"encoding/json"
	"grpc-json/internal"
	"net/http"

	pb "grpc-json/keyvalue"

	"github.com/apex/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Controller struct {
	*internal.KeyValue
}

func (c *Controller) SetHandler(w http.ResponseWriter, r *http.Request) {
	var kv pb.SetKeyValueRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&kv); err != nil {
		log.Errorf("failed to decode: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	gresp, err := c.Set(r.Context(), &kv)
	if err != nil {
		log.Errorf("failed to set: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(gresp)
	if err != nil {
		log.Errorf("failed to marshal: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (c *Controller) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	kv := pb.GetKeyValueRequest{Key: key}

	gresp, err := c.Get(r.Context(), &kv)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		log.Errorf("failed to get: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(gresp)
	if err != nil {
		log.Errorf("failed to marshal: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
