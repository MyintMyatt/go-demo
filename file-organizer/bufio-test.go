package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	fmt.Println("File Organizer ########")
  
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("enter text: ")

	if scanner.Scan(){
     input := scanner.Text()
		 inByte := scanner.Bytes()
		 fmt.Printf("You entered : %s\n", input)
		 fmt.Printf("You entered as numbers: %v\n", inByte) 
	}

	if err := scanner.Err(); err != nil{
		fmt.Fprint(os.Stderr,"error reading input:",err)
	}


	for i := 0; i <=100 ; i++{
	}

}

func checkErr(err error){
   if err != nil{
		 fmt.Printf("Error happened %s \n", err)
	 }
}

