package main

import "errors"

const (
	PoliceDressType    = 1
	TerroristDressType = 2
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[int]IDress),
	}
)

type dressFactory struct {
	dressMap map[int]IDress
}

func (d *dressFactory) getDressByType(dressType int) (IDress, error) {
	if _, exist := d.dressMap[dressType]; exist {
		return d.dressMap[dressType], nil
	}

	var dress IDress
	switch dressType {
	case PoliceDressType:
		dress = NewPoliceDress()
	case TerroristDressType:
		dress = NewPoliceDress()
	}

	if dress != nil {
		d.dressMap[dressType] = dress
		return dress, nil
	}

	return nil, errors.New("not support this dress type")
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}
