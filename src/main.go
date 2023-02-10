package main

import (
	"fmt"
	"github.com/klauspost/compress/zstd"
	"io"
	"os"
)

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
		folderName: "testFolders",
		folderPath: "E:\\GoWorkspace\\BackUp\\",
	}
	folders := []Folder{folder}
	compressFolder(folders)
	//fmt.Printf("folder: %v\n", folder)

}

func createFolder(folderName string, excluded []string) Folder {
	folder := Folder{
		folderName: folderName,
		excluded:   excluded}

	return folder
}

/*
*
We use zstd as compression algorithm
*/
func compressFolder(folders []Folder) {

	folder, _ := os.OpenFile("E:\\GoWorkspace\\BackUp\\encoded\\test.zst", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 2)
	enc, _ := zstd.NewWriter(folder)
	/*
		ReadDir -> Inhalte des Order
		Wenn Inhalt Ordner -> Muss nochmal ReadDir machen
		für Jedes File in den Ordnern -> Compress
	*/
	for i := range folders {
		//hier sollte eine neue Funktion aufgerufen werden für rekursion
		dir, err := os.ReadDir(folders[i].folderPath + folders[i].folderName)
		if err != nil {
			fmt.Printf(err.Error())
		}
		for j := range dir {
			fmt.Printf(dir[j].Name() + "\n")
			if dir[j].IsDir() == false {
				bla := folders[i].folderPath + folders[i].folderName + "\\" + dir[j].Name()
				file, _ := os.Open(bla)
				_, err = io.Copy(enc, file)
				if err != nil {
					enc.Close()
					fmt.Printf(err.Error())
					file.Close()
				}
			}
		}
		erro := enc.Close()
		if erro != nil {
			enc.Close()
			fmt.Printf(erro.Error())
		}
	}

	//os.Create()
	//os.Open("/hello.txt")
	//testing := []byte("hello")
	//writer := io.Writer(testing)
	//zstd.NewWriter(writer)
	/**
	algrotihm = getAlgo(compressAlgo) /
	foreach(folder in folders){
	algorithm(folder)

	*/

}
func Compress(in io.Reader, out io.Writer) error {
	enc, err := zstd.NewWriter(out)
	if err != nil {
		return err
	}
	_, err = io.Copy(enc, in)
	if err != nil {
		enc.Close()
		return err
	}
	return enc.Close()
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
