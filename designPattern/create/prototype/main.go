package main

import "fmt"

func main() {
	folder := NewFolder("folder")
	subfolder := NewFolder("subfolder")
	folder.Add(NewFile("test1"))
	folder.Add(NewFile("test2"))
	folder.Add(NewFile("test3"))
	folder.Add(subfolder)

	subfolder.Add(NewFile("test4"))
	subfolder.Add(NewFile("test5"))

	folder.Print(" ")

	fmt.Println("")
	folder2 := folder.Clone()
	folder2.Print("-")

}
