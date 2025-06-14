package main

import(
	"net/http"
	"fmt"
	//"os"
	"io"
	
)
const url = "https://chatgpt.com/"


func main(){
	response,err:=http.Get(url)

	if err!=nil{
		panic(err)
	}

	defer response.Body.Close()

	databytes,err:=io.ReadAll(response.Body)

	if err!=nil{
		panic(err)
	}

	content := string(databytes)

	fmt.Printf("The body of the response from %v is %v\n ",url,content)


}
