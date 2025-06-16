package main

import (
	"encoding/json"
	"net/http"
	"io"
	"log"
	"github.com/gorilla/mux"
)


type Employee struct{
	Name string `json:"name"`
	ID string   `json:"id"`
	Salary string `json:"salary"`

}

var employees = []Employee{
	{Name:"Vikas",ID:"312",Salary:"1300000"},
	{Name:"Jai",ID:"341",Salary:"1200000"},
}

// Get /employees 
func getEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")

	json.NewEncoder(w).Encode(employees)

}

// Get /employees/{id}

func getEmployeeByID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	param := mux.Vars(r)

	for _,emp := range employees{
		if emp.ID==param["id"]{
			json.NewEncoder(w).Encode(emp)
			return
		}
	}
	http.Error(w,"Employee not found",http.StatusNotFound)

}

// Post /employees 

func createEmployee(w http.ResponseWriter ,  r * http.Request){
	w.Header().Set("Content-type","application/json")
	body,err := io.ReadAll(r.Body)
	if err!=nil{
		http.Error(w,"Invalid body",http.StatusBadRequest)
		return 
	}

	var newemp Employee
	json.Unmarshal(body,&newemp)
	employees=append(employees,newemp)
	json.NewEncoder(w).Encode(newemp)
}

// Put /employees/{id}

func updateEmployeeByID(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	body,_:=io.ReadAll(r.Body)

	var updatedEmp Employee
	json.Unmarshal(body,&updatedEmp)

	for i,emp:= range employees{
		if emp.ID==params["id"]{
employees[i].Name = updatedEmp.Name
employees[i].Salary = updatedEmp.Salary
	json.NewEncoder(w).Encode(employees[i])		
			return 
		}

	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}



func main(){
	router:=mux.NewRouter()
	router.HandleFunc("/employees",getEmployee).Methods("GET")
	router.HandleFunc("/employees/{id}",getEmployeeByID).Methods("GET")
	router.HandleFunc("/employees",createEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}",updateEmployeeByID).Methods("PUT")
	log.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080",router))


}
