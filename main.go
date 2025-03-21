package main

import (
	"fmt"
	"net/http"
)

func main(){

	
	http.HandleFunc("/hello", greetUser)

	http.ListenAndServe(":8000",nil)

	readNumber()
}
func greetUser (writer http.ResponseWriter, request *http.Request) {
	var greet = "hello ishara aaaaaaaaaaaaaaa"
	fmt.Fprintf(writer, greet) 
} 

func readNumber(){
	println("isharaaaaaaaaaaaaaaa")
}


func printTask(taskArray[]string){
	for index,taskItem := range(taskArray){
		fmt.Printf("%v. %v",index+1,taskItem)
		println()
	}
 }

 func addTask(taskArray[] string,newTask string ) [] string{
	 var updatedTask = append(taskArray,newTask)
	 printTask(taskArray)
	 return updatedTask

}





