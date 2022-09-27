package main

type Player struct {
	dress    IDress
	playType int
	lang     int
	long     int
}

func NewPlayer(playerType int) *Player {
	var dressType int
	switch playerType {
	case CT:
		dressType = PoliceDressType
	case TR:
		dressType = TerroristDressType
	}

	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playType: playerType,
		dress:    dress,
	}
}
