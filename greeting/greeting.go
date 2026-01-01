package greeting

import "fmt"

func SayGreeting(name string) string{
	message := fmt.Sprintf("Hi, %v. Welcome !",name)
	return message
}
