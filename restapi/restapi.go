package restapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Action struct {
	method   string
	resource string
}

type RestApi struct {
	actions map[Action]func(int64, []byte) (int, interface{})
}

func (rApi *RestApi) Run() {
	http.HandleFunc("/", rApi.routerHandler)
	http.ListenAndServe(":8080", nil)
}

func (rApi *RestApi) AddAction(method string, resource string, resolver func(int64, []byte) (int, interface{})) {
	action := Action{
		method:   method,
		resource: resource,
	}
	rApi.actions[action] = resolver
}

func (rApi *RestApi) routerHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("ROUTE: " + req.Method + " " + req.URL.Path)
	id := rApi.getId(req)
	controller := rApi.getController(req)
	if controller != nil {
		status, message := controller(id, getPayload(req))
		sendResponse(&w, status, message)
	} else {
		message := make(map[string]string)
		message["message"] = "404 - Not Found"
		sendResponse(&w, http.StatusNotFound, message)
	}
}

func (rApi *RestApi) getController(req *http.Request) func(int64, []byte) (int, interface{}) {
	for action, controller := range rApi.actions {
		if req.Method == action.method && strings.HasPrefix(req.URL.Path, action.resource) {
			return controller
		}
	}
	return nil
}

func (rApi *RestApi) getId(req *http.Request) int64 {
	urlParts := strings.Split(req.URL.Path, "/")
	id := int64(0)
	if len(urlParts) == 3 {
		idTemp, _ := strconv.Atoi(urlParts[2])
		id = int64(idTemp)
	}
	return id
}

func NewRestApi() *RestApi {
	rApi := &RestApi{
		actions: make(map[Action]func(int64, []byte) (int, interface{})),
	}
	return rApi
}

func sendResponse(w *http.ResponseWriter, statusResponse int, bodyResponse interface{}) {
	(*w).WriteHeader(statusResponse)
	tasksJsonResponse, err := json.Marshal(bodyResponse)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	(*w).Write(tasksJsonResponse)
}

func getPayload(r *http.Request) []byte {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	return reqBody
}
