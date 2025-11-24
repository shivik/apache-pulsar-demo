package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/apache/pulsar-client-go/pulsar"
)

// "github.com/apache/pulsar-client-go/pulsar"

func main() {
	fmt.Println("Starting http server")
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		msg := r.URL.Query().Get("query")
		err := publishMessage(msg)
		if err != nil {
			w.Write([]byte("msg failed to published"))
		} else {
			w.Write([]byte("msg successfully published"))
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func publishMessage(msg string) error {
	producer, err:=client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic"
	})

if err!=nil {
	log.Fatal(err)
}

_,err = producer.Send(context.Background(),&pulsar.ProducerMessage {
	payload: [] byte("Hello"),
})

defer producer.Close()

if err!=nil {
	fmt.Println("Falied to Publish the Message", err)
}

fmt.Println("Message Published")
}
