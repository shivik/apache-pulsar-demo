// consumer, err := client.Subscribe(pulsar.ConsumerOptions{
// 	Topic:            "topic-1",
// 	SubscriptionName: "my-sub",
// 	Type:             pulsar.Shared,
// })

// if err != nil {
// 	log.Fatal(err)
// }

// defer consumer.Close()

// for i := 0; i < 10; i++ {
// 	msg, err := consumer, Receive(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Received message with Id: %#v -- content: '%s'\n", msg.ID(), string(msg.Payload()))
// 	consumer.Ack(msg)
// }

// if err := consumer.Unsubscribe(); err != nil {
// 	log.Fatal(err)
// }