package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	service "microservice/core"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//createFindHandler prepares the handler for find requests
func createFindHandler(svc service.MentionService) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		nameParam := ps.ByName("name")
		log.Printf("FindHandler(name = %s)\n", nameParam)

		bandMentions, err := svc.Find(nameParam)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNoContent)
			log.Println(err)
		}
		fmt.Println(bandMentions)

		byteResponse, err := json.Marshal(bandMentions)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(byteResponse)
		log.Println("OK")
	}
}

//createCreateHandler prepares the handler for create requests
func createCreateHandler(svc service.MentionService) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		nameParam := ps.ByName("name")
		log.Printf("CreateHandler(name = %s)\n", nameParam)

		err := svc.Create(nameParam)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		log.Println("OK")
	}
}

//createIncreaseHandler prepares the handler for increase requests
func createIncreaseHandler(svc service.MentionService) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		nameParam := ps.ByName("name")
		log.Printf("IncreaseHandler(name = %s)\n", nameParam)

		//finds which entity to have its counter increased
		bm, err := svc.Find(nameParam)
		if len(bm) == 0 {
			log.Printf("%v not found!\n", nameParam)
			http.Error(w, "band was not found!", 500)
			return
		}

		bmResp, err := svc.Increase(&bm[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		byteResponse, err := json.Marshal(bmResp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		log.Println("OK")
		w.WriteHeader(http.StatusOK)
		w.Write(byteResponse)
	}
}

//CreateHTTPServer prepares the instance of a new httprouter
func CreateHTTPServer(svc service.MentionService) (*httprouter.Router, error) {

	router := httprouter.New()
	router.GET("/find/:name", createFindHandler(svc))
	router.GET("/create/:name", createCreateHandler(svc))
	router.GET("/increase/:name", createIncreaseHandler(svc))

	log.Println("HTTPServer created...")
	return router, nil
}
