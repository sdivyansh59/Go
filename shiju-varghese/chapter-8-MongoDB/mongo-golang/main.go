package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)


type Person struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string	`json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname string		`json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

// const url = "mongodb+srv://divyansh:divyansh@sandbox.fma8t.mongodb.net/?retryWrites=true&w=majority"

func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type","application/json")
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	collection := client.Database("sandbox").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx,person)
	json.NewEncoder(w).Encode(result)

}


func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type","application/json")
	var people [] Person
	collection := client.Database("sandbox").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx,bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message: "}`+ err.Error() + `"}`))
		return
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err() ; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message: "}`+ err.Error() + `"}`))
		return
	}

	json.NewEncoder(w).Encode(people)

}

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request){

}

func main() {
	// creating context
	r := mux.NewRouter()
	r.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	r.HandleFunc("/people",  GetPeopleEndpoint).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}



func init() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
    	ApplyURI("mongodb+srv://divyansh:test123@sandbox.fma8t.mongodb.net/?retryWrites=true&w=majority").
    	SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
	    log.Fatal(err)
	}

}