package repository

import (
	"apiMongo/model"
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("error while connecting mongodb", err)
	}
	log.Println("mongodb successfuly connection", err)

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

	log.Println("ping success")
	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	emp1 := uuid.New().String()
	emp2 := uuid.New().String()

	//connect to collection
	coll := mongoTestClient.Database("companydb").Collection("employee_test")

	empRepo := EmployeeRepo{MongoCollection: coll}

	// InsertEmployee 1 data
	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Name Lastname",
			Department: "Policeman",
			EmployeeID: emp1,
		}
		result, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal("insert 1 operation failed", err)
		}
		t.Log("Insert 1 successful", result)
	})

	// InsertEmployee 2 data
	t.Run("Insert Employee 2", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Name Surname",
			Department: "Managers",
			EmployeeID: emp2,
		}
		result, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal("insert 2 operaion failed", err)
		}
		t.Log("Insert 2 successful", result)
	})

	//Get Employee 1 Data
	t.Run("Get Employee 1", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)

		if err != nil {
			t.Fatal("get operation failed", err)
		}
		t.Log("get operation successful", result)
	})

	//Get All Employees
	t.Run("Get All Employees", func(t *testing.T) {
		result, err := empRepo.FindAllEmployee()

		if err != nil {
			t.Fatal("Get Operation Faled", err)
		}
		t.Log("employees", result)
	})

	//Update Employee 2 data
	t.Run("Update Employee 2 dapertment", func(t *testing.T) {
		emp := model.Employee{
			Department: "Work",
		}

		result, err := empRepo.UpdateEmployeeByID(emp2, &emp)
		if err != nil {
			log.Fatal("update operation is failed", err)
		}
		t.Log("update count", result)
	})

	//Delete employee 1 data
	t.Run("Delete employee 1 data", func(t *testing.T) {
		result, err := empRepo.DeleteEmployeeByID(emp1)
		if err != nil {
			log.Fatal("delete operation is failed")
		}
		t.Log("delete count", result)
	})

	//Delete all employees
	t.Run("Delete All Employess fon cleanup", func(t *testing.T) {
		result, err := empRepo.DeleteAllEmployee()
		if err != nil {
			log.Fatal("delete All operation is failed")
		}
		t.Log("delete count", result)
	})
}
