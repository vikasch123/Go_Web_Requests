package main

import (
	"net/http"
	"log"
	"io"
	"github.com/gorilla/mux"
	"encoding/json"
)

type Student struct{
	Name string `json:"name"`
	ID string   `json:"id"`
	Course string `json:"course"`

}

var students =[]Student{
	{Name:"Vikas",ID:"123",Course:"CSE"},
	{Name:"Abhinav",ID:"312",Course:"CSE-IT"},

}
// allow cross-origin requests (especially from web clients like Hoppscotch):

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins (for dev only)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}


//Get /student  

func getStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")

	json.NewEncoder(w).Encode(students)

}

// Get /student/{id}

func getStudentByID(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-type","application/json")

	param :=mux.Vars(r)

	for _,stud := range students{
		if stud.ID==param["id"]{
			json.NewEncoder(w).Encode(stud)
			return
		}
		// error handle 
	}
		http.Error(w,"student with that id not found",http.StatusNotFound)
	
}

// Post /student 

func createStudent(w http.ResponseWriter, r *http.Request ){
	enableCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}


	w.Header().Set("Content-type","application/json")

	body,err := io.ReadAll(r.Body)

	if err!=nil{
		http.Error(w,"Invalid body",http.StatusBadRequest)
		return 
	}
	var newstudent Student

	json.Unmarshal(body,&newstudent)
	students=append(students,newstudent)
	log.Println("New Student added successfully!!!!")
	json.NewEncoder(w).Encode(newstudent)
}


func main(){
	router := mux.NewRouter()
	router.HandleFunc("/Student",getStudent).Methods("GET")
	router.HandleFunc("/Student/{id}",getStudentByID).Methods("GET")
	router.HandleFunc("/Student",createStudent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080",router))
}



