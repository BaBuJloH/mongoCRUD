package usecase

import (
	"apiMongo/model"
	"apiMongo/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (srv *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}
	//assign new employee id
	emp.EmployeeID = uuid.NewString()

	repo := repository.EmployeeRepo{MongoCollection: srv.MongoCollection}
	//insert employee
	insertID, err := repo.InsertEmployee(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("insert error", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusOK)
	log.Println("employee inserted with id", insertID, emp)
}

func (srv *EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	//get employee id
	empID := mux.Vars(r)["id"]
	log.Println("employee id", empID)

	repo := repository.EmployeeRepo{MongoCollection: srv.MongoCollection}
	emp, err := repo.FindEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (srv *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: srv.MongoCollection}

	emp, err := repo.FindAllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (srv *EmployeeService) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	// get employee id

	empID := mux.Vars(r)["id"]
	log.Println("employee id", empID)

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid employee id")
		res.Error = "invalid employee id"
		return
	}
	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}
	emp.EmployeeID = empID
	repo := repository.EmployeeRepo{MongoCollection: srv.MongoCollection}
	count, err := repo.UpdateEmployeeByID(empID, &emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error: ", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}

func (srv *EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	//get employee id
	empID := mux.Vars(r)["id"]
	log.Println("employee id", empID)

	repo := repository.EmployeeRepo{MongoCollection: srv.MongoCollection}

	count, err := repo.DeleteEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error: ", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}

func (srv *EmployeeService) DeleteAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: srv.MongoCollection}

	count, err := repo.DeleteAllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error: ", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}
