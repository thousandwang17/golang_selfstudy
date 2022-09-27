package main

import "errors"

type lightingAdapter struct {
	joint interface{}
}

func NewLightingAdapter(i interface{}) (ILighting, error) {
	switch i.(type) {
	case iusb:
	case ihdmi:
	case lighting:
	default:
		return nil, errors.New("not support this type")
	}

	return &lightingAdapter{
		joint: i,
	}, nil
}

func (l *lightingAdapter) InsertIntoLightningPort() {
	switch v := l.joint.(type) {
	case iusb:
		v.InsertIntoUSBPort()
	case ihdmi:
		v.InsertIntoHDMIPort()
	case ILighting:
		v.InsertIntoLightningPort()
	}
}

func (l *lightingAdapter) ReadData() string {
	res := ""
	switch v := l.joint.(type) {
	case iusb:
		res = v.ReadData()
	case ihdmi:
		res = v.ReadData()
	case ILighting:
		res = v.ReadData()
	}
	return res
}
