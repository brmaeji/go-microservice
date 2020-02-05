package main

import (
	"fmt"
	"log"
	data "microservice/adapters/data"
	endpoints "microservice/adapters/endpoints"
	service "microservice/core"
	"net/http"
	"os"
	"os/signal"
)

func cleanExit() {
	fmt.Println("Clean exit")
	os.Exit(0)
}

func main() {
	log.Println("Starting microservice...")

	//variables definitions
	var port string

	//setup a CTRL+C listener to execute clean exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			cleanExit()
		}
	}()

	//get runtime arguments
	argsWithoutProg := os.Args[1:]

	//parses dataAdapterType
	var dataType data.AdapterType
	if len(argsWithoutProg) == 0 {
		dataType = data.Memory
	} else {
		dataTypeArg := argsWithoutProg[0]
		log.Printf("Chosen data adapter type: %v\n", dataTypeArg)
		switch dataTypeArg {
		case "memory":
			dataType = data.Memory
			break
		}

		if len(argsWithoutProg) == 2 {
			port = fmt.Sprintf(":%v", argsWithoutProg[1])
		} else {
			port = ":8080"
		}
	}

	//create new service
	srv, err := service.NewMentionsService(dataType)
	if err != nil {
		log.Println("Could not create MentionsService:")
		log.Fatalln(err)
	}

	//creates new http endpoint
	router, err := endpoints.CreateHTTPServer(srv)
	if err != nil {
		log.Println("Could not create HTTPServer:")
		log.Fatalln(err)
	}

	log.Printf("Listening on port %v!\n", port)
	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Println("Could not listen for connections on HTTP Server:")
		log.Fatalln(err)
	}

	// Maybe disposes were not ran!
	log.Println("Clean exit was not called!")
	os.Exit(1)
}
