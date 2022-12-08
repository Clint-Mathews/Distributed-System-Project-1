package helper

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func CreateRedisClient() redis.Conn {

	port := readConfig("REDIS_PORT")

	c, err := redis.Dial("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error making connection. Error: %+v", err)
	}
	return c

}

func PublishToQueue(redisClient redis.Conn, msg MsgType) {

	payload, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Error marshalling message. Error: %+v", err)
		return
	}

	if _, err := redisClient.Do("PUBLISH", "get-queue-data", payload); err != nil {
		log.Fatalf("Error publishing message. Error: %+v", err)
	}
}

func SubscribeToQueue(redisClient redis.Conn) {
	psc := redis.PubSubConn{Conn: redisClient}
	if err := psc.Subscribe("get-queue-data"); err != nil {
		log.Fatalf("Error subscribing to channel. Error: %+v", err)
	}

	for {
		responseMsg := MsgType{}
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			if err := json.Unmarshal([]byte(v.Data), &responseMsg); err != nil {
				log.Fatalf("Error in receiving message. Error: %+v", err)
			}
			responseMsg.saveMessage()
		case redis.Subscription:
			fmt.Printf("Channel - %s: Subscribed\n", v.Channel)
		case error:
			fmt.Println(v)
		}
	}
}
