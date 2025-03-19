
package main

import "fmt"

func main(){
	var EventName = "Earth conference"
	const TotalTicket = 100
	var availableTicket int = -45

	fmt.Println("Hi,welcome to",EventName )
	fmt.Println("you have total number of tickets",TotalTicket,"still available",availableTicket)
	fmt.Printf("hii,welcome to %v\n",EventName)

	var userName string
	fmt.Printf("username is %v",userName)
	fmt.Printf("the type of name variable is %T",userName)

	userEmail:="ishrapeiris2001@gmail.com"
	fmt.Print(userEmail)
var firstName string
var lastName string
var emailAddress string
var age int
fmt.Print("enter your first name ")
	fmt.Scan(&firstName)
 fmt.Print("enter your lastname")
 fmt.Scan(&lastName)
 fmt.Print("enter your email addres")
 fmt.Scan(&emailAddress)
 fmt.Print("enter your age")
 fmt.Scan(&age)

 fmt.Printf("hi, %v %v your %v received and your age is %v",firstName,lastName,emailAddress,age)

}
