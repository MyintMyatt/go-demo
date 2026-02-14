package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
  "path/filepath"
	"time"
	"github.com/schollz/progressbar/v3"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Which folder do you want to organize?:")
	if scanner.Scan(){
		targetFolder := scanner.Text()
    _, err := checkFolderIsExitOrNot(targetFolder)
  
		if err != nil {
		  fmt.Printf("Error checking folder: %v\n", err)
		  return
	  }

		createDefaultFolder(targetFolder)
    organizeFolder(targetFolder)
	} 
}

// organize folders 
func organizeFolder(targetFolder string){
	filesNFolders,err := os.ReadDir(targetFolder)
	if err != nil{
		fmt.Println("reading folder error: " , err)
		os.Exit(1)
	}
  var files []os.DirEntry
	for _, entry := range filesNFolders {
		if !entry.IsDir() {
			files = append(files, entry)
		}
	}
	noOfFiles := 0
	start := time.Now()
	// total files
	bar := progressbar.Default(int64(len(files)))
	for _, filesNFolder := range filesNFolders{
		// check folder or file , if folder , we don't organize
	   if !filesNFolder.IsDir(){
				fileInfo, err := filesNFolder.Info()
				if err != nil{
					fmt.Println("reading file error: " , err)
					os.Exit(1)
				}

				oldPath := filepath.Join(targetFolder, fileInfo.Name())
				fileExt := filepath.Ext(oldPath)

				switch fileExt{
				case ".png", ".jpg", ".jpeg", ".svg":
					newPath := filepath.Join(targetFolder, "Images", fileInfo.Name())
					err := os.Rename(oldPath, newPath)
					check(err)
          noOfFiles++
				case ".mp4", ".mov", ".avi", ".amv":
					newPath := filepath.Join(targetFolder, "Videos", fileInfo.Name())
					err := os.Rename(oldPath, newPath)
					check(err)
					noOfFiles++
				case ".pdf", ".docx", ".csv", ".xlsx":
				  newPath := filepath.Join(targetFolder, "Docs", fileInfo.Name())
				  err = os.Rename(oldPath, newPath)
				  check(err)
				  noOfFiles++
			  case ".mp3", ".wav", ".aac":
				  newPath := filepath.Join(targetFolder, "Music", fileInfo.Name())
			   	err = os.Rename(oldPath, newPath)
				  check(err)
				  noOfFiles++
			  default:
				  newPath := filepath.Join(targetFolder, "Others", fileInfo.Name())
				  err = os.Rename(oldPath, newPath)
			  	check(err)
				  noOfFiles++
				}
				bar.Add(1)
		 }
 	}

	if noOfFiles > 0{
		fmt.Println("#################################")
		fmt.Println(noOfFiles,"number of file moved.")
		elasped := time.Since(start)
		fmt.Printf("Processing time : %v\nms" , elasped.Milliseconds())
	}
}


// create default folder 
func createDefaultFolder(targetFolder string){
	defaultFolders := []string{"Music", "Videos", "Docs", "Images", "Others"}
	for _, folder := range defaultFolders{
		_, err := os.Stat(folder)
		if os.IsNotExist(err){
			os.Mkdir(filepath.Join(targetFolder, folder),0755)
		}
	}
}

// check folder is exited or not
func checkFolderIsExitOrNot(path string) (bool, error){
	info, err := os.Stat(path)
	if err == nil{
		fmt.Println("folder name is : ", info.Name(), info.Sys())
		return info.IsDir() , nil
	}
	if errors.Is(err, os.ErrNotExist){
     return false, err
	}
	return false, err
}

func check(err error) {
	if err != nil {
		fmt.Printf("Error Happened %s \n", err)
		os.Exit(1)
	}
}

