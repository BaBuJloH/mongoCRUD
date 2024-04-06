package main

import (
	"apiMongo/usecase"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env loadind error", err)
	}
	log.Println("env file loaded")

	//create mongo Client
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("connection error", err)
	}
	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

	log.Println("mongo connected")
}

func main() {
	//close the mongo onnection
	defer mongoClient.Disconnect(context.Background())

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	//create employee service
	empService := usecase.EmployeeService{MongoCollection: coll}

	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)

	r.HandleFunc("/employee", empService.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}", empService.GetEmployeeByID).Methods(http.MethodGet)
	r.HandleFunc("/employee", empService.GetAllEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", empService.UpdateEmployeeByID).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", empService.DeleteEmployeeByID).Methods(http.MethodDelete)
	r.HandleFunc("/employee", empService.DeleteAllEmployee).Methods(http.MethodDelete)

	log.Println("server is running on port 4444")
	http.ListenAndServe(":4444", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running..."))
}
