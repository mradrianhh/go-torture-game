package models

import "fmt"

// Enemy - in singleplayer mode - represents the target that the player is torturing.
//
// If the enemy's health reaches 0 before you acquire the truth, you lose.
// If the enemy's mentality reaches 0, you'll get the truth. Between the stages of full mentality and zero mentality, the enemy might reveal the "truth", but it won't be correct.
// Only if the enemy's mentality reaches 0 will it be correct, but you will never know.
type Enemy struct {
	name      string
	health    int
	mentality int
}

// EnemyList holds the global collection of all the enemies in the game.
var EnemyList []Enemy

var fakeTruths []string

func init() {
	fakeTruths = []string{
		"Fake Truth 1",
		"Fake Truth 2",
	}

	EnemyList = []Enemy{
		{name: "Level 1 Enemy", health: 100, mentality: 10},
		{name: "Level 2 Enemy", health: 100, mentality: 20},
		{name: "Level 3 Enemy", health: 110, mentality: 20},
		{name: "Level 4 Enemy", health: 120, mentality: 20},
		{name: "Level 5 Enemy", health: 150, mentality: 50},
		{name: "Level 6 Enemy", health: 150, mentality: 60},
		{name: "Level 7 Enemy", health: 150, mentality: 70},
		{name: "Level 8 Enemy", health: 160, mentality: 70},
		{name: "Level 9 Enemy", health: 170, mentality: 70},
		{name: "Level 10 Enemy", health: 200, mentality: 100},
	}
}

// GetName getter.
func (enemy *Enemy) GetName() string {
	return enemy.name
}

// SetName setter.
func (enemy *Enemy) SetName(name string) {
	enemy.name = name
}

// GetHealth getter.
func (enemy *Enemy) GetHealth() int {
	return enemy.health
}

// SetHealth setter.
func (enemy *Enemy) SetHealth(health int) {
	enemy.health = health
}

// GetMentality getter.
func (enemy *Enemy) GetMentality() int {
	return enemy.mentality
}

// SetMentality setter.
func (enemy *Enemy) SetMentality(mentality int) {
	enemy.mentality = mentality
}

// RevealFakeTruth reveals a single, random fake-truth to the player.
func (enemy *Enemy) RevealFakeTruth() {
	fmt.Println(fakeTruths[0])
}
