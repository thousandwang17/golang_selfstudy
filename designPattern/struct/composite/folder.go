package main

import "fmt"

type folder struct {
	childerns []iComponent
	name      string
}

type iFolder interface {
	Add(iComponent)
	iComponent
}

func NewFolder(name string) iFolder {
	return &folder{
		name: name,
	}
}

func (f *folder) Add(i iComponent) {
	f.childerns = append(f.childerns, i)
}

func (f *folder) search(keyword string) {
	fmt.Printf("Searching for keyword %s in folder %s\n", keyword, f.name)

	for i := range f.childerns {
		f.childerns[i].search(keyword)
	}
}
