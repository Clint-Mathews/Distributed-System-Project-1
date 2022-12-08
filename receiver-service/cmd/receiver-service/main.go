package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Clint-Mathews/Distributed-System-Project-1/receiver-service/helper"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

type RestAPI struct {
	subRedisClient redis.Conn
	pubRedisClient redis.Conn
}

func main() {
	api := RestAPI{
		subRedisClient: helper.CreateRedisClient(),
		pubRedisClient: helper.CreateRedisClient(),
	}

	helper.CreateFile()

	// Creates Rest Endpoints
	go api.createServer()

	// Subscibes to redis queue
	helper.SubscribeToQueue(api.subRedisClient)

	defer helper.DeleteFile()

}

func (api *RestAPI) createServer() {
	port := 8000
	url := fmt.Sprintf(":%d", port)

	r := mux.NewRouter()
	r.HandleFunc("/", api.GetQueueRecords).Methods("GET")
	r.HandleFunc("/", api.PostQueueRecords).Methods("POST")
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         url,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Reciver service started on port: %d", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Unable to Listen And Serve on port: %+v", err)
		os.Exit(1)
	}
}

func (*RestAPI) GetQueueRecords(w http.ResponseWriter, r *http.Request) {
	data := helper.GetMessage()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (api *RestAPI) PostQueueRecords(w http.ResponseWriter, r *http.Request) {
	var msg helper.MsgType
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	go helper.PublishToQueue(api.pubRedisClient, msg)
	w.WriteHeader(http.StatusOK)
}
