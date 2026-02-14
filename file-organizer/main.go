package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Which folder do you want to organize?:")
	if scanner.Scan(){
		targetFolder := scanner.Text()
    folderExits, err := checkFolderIsExitOrNot(targetFolder)
  
		if err != nil {
		  fmt.Printf("Error checking folder: %v\n", err)
		  return
	  }

		fmt.Println("folder is exited : " , folderExits)
	} 
}


func checkFolderIsExitOrNot(path string) (bool, error){
	info, err := os.Stat(path)
	if err == nil{
		fmt.Println("folder name is : ", info.Name())
		return info.IsDir() , nil
	}
	if errors.Is(err, os.ErrNotExist){
     return false, err
	}
	return false, err
}
