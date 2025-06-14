package main

import (
	"encoding/json"
	"net/http"
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


func getEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")

	json.NewEncoder(w).Encode(employees)

}

func main(){
	http.HandleFunc("/employees",getEmployee)
	http.ListenAndServe(":8080",nil)


}
