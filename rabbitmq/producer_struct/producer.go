package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ProduceRabbitmqMsg(w http.ResponseWriter, r *http.Request) {
	qName := r.FormValue("queue_name")

	rq, err := NewRabbitMqClient()
	if err != nil {
		log.Fatalf("RabbitMQ Connection err : %s", err.Error())
	}
	defer rq.Close()

	if err := rq.getChannel(); err != nil {
		log.Printf("RabbitMQ Get Channel Error : %s\n", err.Error())
		return
	}

	if err := rq.queueDeclare(qName, false, false, false, false, nil); err != nil {
		log.Printf("RabbitMQ Queue Declare Error : %s\n", err.Error())
		return
	}

	payload, err := makePayloadMsg()
	if err != nil {
		log.Printf("make payload Error : %s\n", err.Error())
		return
	}

	if err := rq.publishMq(qName, "", false, false, payload); err != nil {
		log.Printf("Failed to publish a queue, err : %s\n", err.Error())
		return
	}

	log.Printf("Success Sent : %s\n", string(payload))
}

func makePayloadMsg() ([]byte, error) {

	callbackData := CallbackDataMsg{
		AccountId:        "jake",
		ContentId:        "jake_overaly_test_1",
		Error:            4882,
		GroupNm:          "kollus",
		MediaOutputs:     "",
		SnapshotOutput:   "",
		State:            "complete",
		ThumbnailOutputs: "",
		WaveformOutput:   "",
	}

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded;charset=UTF-8;",
	}

	callback := Callback{
		RequestID:          "RR",
		ContentProviderKey: "jungin-kim",
		CallbackMethod:     "post",
		CallbackURL:        "http://localhost:3020/hello",
		CallbackHeader:     header,
		CallbackData:       callbackData,
		MediaContentID:     5,
	}

	Req := CallbackRequestQueue{
		Callback: callback,
	}

	payload, err := json.MarshalIndent(Req, "", "")
	if err != nil {
		return nil, err
	}

	return payload, nil
}
