package main

import "fmt"

type hdmi struct{}

type ihdmi interface {
	InsertIntoHDMIPort()
	ReadData() string
}

func NewHDMI() ihdmi {
	return &hdmi{}
}

func (w *hdmi) InsertIntoHDMIPort() {
	fmt.Println("HDMI insert into  Port")
}

func (l *hdmi) ReadData() string {
	return "i'm HDMI"
}
