package main

import "fmt"

type folder struct {
	childFile []Inode
	name      string
}

type Folder interface {
	GetFolderName() string
	Add(Inode)
	Inode
}

func NewFolder(name string) Folder {
	return &folder{
		childFile: []Inode{},
		name:      name,
	}
}

func (f *folder) GetFolderName() string {
	return f.name
}

func (f *folder) Print(prefix string) {
	fmt.Println(prefix + "\x1B[44m\x1B[1m" + f.name + "\x1b[0m")
	for i := range f.childFile {
		f.childFile[i].Print(prefix + prefix)
	}
}

func (f *folder) Add(i Inode) {
	f.childFile = append(f.childFile, i)
}

// deep copy
func (f *folder) Clone() Inode {
	clone := &folder{
		name:      f.name + "_copy",
		childFile: []Inode{},
	}

	for i := range f.childFile {
		copy := f.childFile[i].Clone()
		clone.childFile = append(clone.childFile, copy)
	}

	return clone
}
