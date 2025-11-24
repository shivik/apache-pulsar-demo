client, err :=pulsar.NewClient(pulsar.ClientOptions{
	URL:"pulsar://localhost:6605"
})

if err!=nil{
	log.Fatal(err)
}

defer client.Close()


go func(){
	prometheusPort:=2112
	log.Printf("Starting prometheus metric at http://localhost:%v/metrics\n", prometheusPort)
	http.Handle("/metrics", promhttp.Handler())
	err= http.ListenAndServe(":"+strconv.Itoa(prometheusPort), nil)

	if err!=nil {
		log.fatal(err)
	}
}()

producer, err:=client.CreateProducer(pulsar.ProducerOptions{
	Topic:"topic-1"
})

if err!=nil{
	log.Fatal(err)
}

defer producer.Close()

ctx:=context.Background()

webPort:=8082
http.HandleFunc("/produce", func(w http.ResposeWriter, r *http.Request){
	msgId, err := producer.Send(ctx, &pulsar.ProducerMessage {
		Payload: []byte(fmt.Sprintf("hello-world")),
	})

	if err!=nil {
		log.Fatal(err)
	}
	else{
		log.Printf("Published message: %v", msgId)
		fmt.Fprintf(w,"Message Published: %v", msgId)
	}
})