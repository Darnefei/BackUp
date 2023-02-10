package main

import "fmt"

type Folder struct {
	folderName string
	folderPath string
	excluded   []string
}

func main() {
	/*
		bla := Folder{
			folderName: "bla",
			folderPath: "hi",
			excluded:   []string{"10", "20", "30", "40", "50"}}
	*/
	folder := Folder{
		folderName: "data",
	}

	fmt.Printf("folder: %v\n", folder)

}

func createFolder(folderName string, folderPath string, excluded []string) Folder {
	folder := Folder{
		folderName: folderName,
		folderPath: folderPath,
		excluded:   excluded}

	return folder
}

func compressFolder(folders []Folder, compressAlgo string) {

}

func saveFolderToJson(folders []Folder) {

}

func removeFolderToJson(folders []Folder) {

}

func addFolderToJson(folders Folder) {

}

func getFolderFromJson() []Folder {
	return nil
}

func validatePath(folder Folder) {

}
