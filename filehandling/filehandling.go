package main

import (
	"fmt"
	"os"
)

func main() {
	// f,err := os.Open("filehandling/example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	//  fileinfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("File Name :-",fileinfo.Name())
	// fmt.Println("Folder :-",fileinfo.IsDir())
	// fmt.Println("File Size :-",fileinfo.Size())

	// buf := make([]byte,fileinfo.Size())
	//
	//  d,err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i:= 0 ; i < len(buf); i++ {
	// 	fmt.Println(d,string(buf[i]))
	// }

	// f, err := os.ReadFile("filehandling/example.txt")
	//
	// if err != nil {
	//    panic(err)
	// }
	//
	// fmt.Println(string(f))

	// dir, err := os.Open(".")
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// defer dir.Close()
	//
	// fileInfo, err := dir.ReadDir(-1)
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// for _,fi := range fileInfo {
	//    fmt.Println(fi.Name(),fi.IsDir())
	// }

	// f, err := os.Create("filehandling/example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	//
	// // f.WriteString("Hi Go")
	//
	// bytes := []byte("Hello Golang")
	//
	// f.Write(bytes)

	// sourceFile, err := os.Open("filehandling/example.txt")
	//
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()
	//
	// desFile, err := os.Create("filehandling/example2.txt")
	//
	// if err != nil {
	// 	panic(err)
	// }
	// defer desFile.Close()
	//
	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(desFile)
	//
	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	//
	// writer.Flush()
	//
	// fmt.Println("Written to new File Successfully")

	ee := os.Remove("filehandling/example2.txt")

	if ee != nil {
		panic(ee)
	}

	fmt.Println("File Deleted Successfully")

}
