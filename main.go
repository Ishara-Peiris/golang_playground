
package main

import "fmt"

func main(){
	var EventName = "Earth conference"
	const TotalTicket = 100
	var availableTicket int = 45

	fmt.Println("Hi,welcome to",EventName )


	fmt.Println("you have total number of tickets",TotalTicket,"still available",availableTicket)
	fmt.Printf("hii,welcome to %v\n",EventName)

	var userName string
	fmt.Printf("username is %v",userName)
	fmt.Printf("the type of name variable is %T",userName)

	userEmail:="ishrapeiris2001@gmail.com"
	fmt.Print(userEmail)
var firstName string
var emailAddress string
var userTicket int
fmt.Print("enter your first name ")
	fmt.Scan(&firstName)

 
 
 fmt.Scan(&emailAddress)
 fmt.Print("enter how many tickets you order")
 fmt.Scan(&userTicket)
availableTicket=availableTicket-userTicket

 fmt.Printf("hi, %v  your %v received and your email:",firstName,emailAddress)
 fmt.Print("available ticket amout is",availableTicket)

 var userArray[50]string
 userArray[1]=firstName
 fmt.Print(userArray[1])

 var userslice[]string
 userslice=append(userslice, userEmail)
 fmt.Print(userslice)
 fmt.Print(len(userArray))

var taskOne = "getup"
taskTwo := "meditate"
taskThree := "reading"
 var todoArray = []string{taskOne,taskTwo,taskThree}

 for _,todoItem := range(todoArray){
	fmt.Printf("%v\n",todoItem)
 }


 printTask(todoArray)

}
func printTask(taskArray[]string){
	for index,taskItem := range(taskArray){
		fmt.Printf("%v. %v",index+1,taskItem)
	}
 }



