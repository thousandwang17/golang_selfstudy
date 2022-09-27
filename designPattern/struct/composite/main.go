package main

func main() {
	folder1 := NewFolder("folder1")
	folder1.Add(NewFile("test1"))
	folder1.Add(NewFile("test2"))
	folder1.Add(NewFile("test3"))

	folder2 := NewFolder("folder2")
	folder2.Add(NewFile("test4"))
	folder2.Add(NewFile("test5"))
	folder2.Add(NewFile("test6"))
	folder1.Add(folder2)

	folder1.search("abc")
	folder2.search("456")
}
