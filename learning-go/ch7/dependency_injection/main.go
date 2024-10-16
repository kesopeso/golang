package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	logger := LoggerAdapter(LogOutput)
	dataStore := SimpleDataStoreFactory()
	simpleLogic := SimpleLogicFactory(logger, dataStore)
	controller := ControllerFactory(logger, simpleLogic)
	http.HandleFunc("/hello", controller.SayHello)
	http.ListenAndServe(":8081", nil)
}

type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In Hello")
	userId := r.URL.Query().Get("user-id")
	message, err := c.logic.SayHello(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte(message))
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userId string) (string, error) {
	sl.l.Log("Saying welcome to " + userId)
	return greeting(sl, userId, "Hello")
}

func (sl SimpleLogic) SayGoodbye(userId string) (string, error) {
	sl.l.Log("Saying goodbye to " + userId)
	return greeting(sl, userId, "Goodbye")
}

func greeting(sl SimpleLogic, userId string, greeting string) (string, error) {
	username, err := sl.ds.UsernameForId(userId)
	if err != nil {
		return "", err
	}

	return greeting + " " + username, nil
}

type LoggerAdapter func(string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UsernameForId(userId string) (string, error) {
	username, ok := sds.userData[userId]
	if !ok {
		return "", errors.New("No such userId " + userId)
	}
	return username, nil
}

type DataStore interface {
	UsernameForId(userId string) (string, error)
}

type Logic interface {
	SayHello(userId string) (string, error)
	SayGoodbye(userId string) (string, error)
}

type Logger interface {
	Log(message string)
}

func LogOutput(message string) {
	fmt.Println(message)
}

func SimpleDataStoreFactory() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

func SimpleLogicFactory(logger Logger, dataStore DataStore) SimpleLogic {
	return SimpleLogic{
		l:  logger,
		ds: dataStore,
	}
}

func ControllerFactory(logger Logger, logic Logic) Controller {
	return Controller{
		l:     logger,
		logic: logic,
	}
}
