package main

type game struct {
	TerrorisPlayer []*Player
	PolicePlayer   []*Player
}

const (
	CT = 1
	TR = 2
)

func NewGame() *game {
	return &game{
		TerrorisPlayer: make([]*Player, 0, 15),
		PolicePlayer:   make([]*Player, 0, 15),
	}
}

func (c *game) addTerrorist() {
	player := NewPlayer(CT)
	c.TerrorisPlayer = append(c.TerrorisPlayer, player)
	return
}

func (c *game) addCounterTerrorist() {
	player := NewPlayer(TR)
	c.PolicePlayer = append(c.PolicePlayer, player)
	return
}
