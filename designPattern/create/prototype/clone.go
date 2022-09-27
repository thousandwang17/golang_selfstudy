package main

type Inode interface {
	Clone() Inode // deep copy
	Print(indentation string)
}
