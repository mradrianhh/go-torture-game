package models

// Player represents the user controlling the game.
type Player struct {
	name string
	gold int
}

// User is the person controlling the game.
var User Player

func init() {
	User = Player{name: "User", gold: 0}
}

// GetName getter.
func (player *Player) GetName() string {
	return player.name
}

// SetName setter.
func (player *Player) SetName(name string) {
	player.name = name
}

// GetGold getter.
func (player *Player) GetGold() int {
	return player.gold
}

// SetGold setter.
func (player *Player) SetGold(gold int) {
	player.gold = gold
}
