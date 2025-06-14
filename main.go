package main

import (
	"encoding/json"
	"net/http"
)

type Student struct{
	Name string `json:"name"`
	ID string `json:"id"`
	Course string `json:"course"`
}

var student= []Student{
	{Name:"Vikas",ID:"210",Course:"CSE"},
	{Name:"Jayesh",ID:"213",Course:"ECE"},
} 


func GetStuds(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(student)

}

func main(){
	http.HandleFunc("/students",GetStuds)
	http.ListenAndServe(":8080",nil)


}
