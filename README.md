# go-fcm
(in beta!)

[![GoDoc](https://godoc.org/github.com/riftbit/gofcm?status.svg)](https://godoc.org/github.com/riftbit/gofcm)
[![Go Report Card](https://goreportcard.com/badge/github.com/riftbit/gofcm)](https://goreportcard.com/report/github.com/riftbit/gofcm)

This project basicly was forked from [github.com/edganiukov/fcm](https://github.com/edganiukov/fcm) and [github.com/appleboy/go-fcm](https://github.com/appleboy/go-fcm).

## Difference with appleboy package

* [x] Go modules with semantic versioning
* [x] valyala/fasthttp client instead of net/http
* [x] mailru/easyjson client instead of encoding/json
* [x] Send() returns original body ([]byte) too (if FCM answer changed you can parse by yourself and not wait for package update)
* [x] Some optimizations 

Golang client library for Firebase Cloud Messaging. Implemented only [HTTP client](https://firebase.google.com/docs/cloud-messaging/http-server-ref#downstream).

More information on [Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging/)

## Feature

* [x] Send messages to a topic
* [x] Send messages to a device list
* [x] Supports condition attribute (fcm only)

## Getting Started

To install gofcm, use `go get`:

```bash
go get github.com/riftbit/gofcm
```

## Sample Usage

Here is a simple example illustrating how to use FCM library:

```go
package main

import (
	"log"

	"github.com/riftbit/gofcm"
)

func main() {
	// Create the message to be sent.
	msg := &fcm.Message{
		To: "sample_device_token",
		Data: map[string]interface{}{
			"foo": "bar",
		},
	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient("sample_api_key")
	if err != nil {
		log.Fatalln(err)
	}

	// Send the message and receive the response without retries.
	response, body, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", response)
	log.Printf("%#v\n", body)
}
```
