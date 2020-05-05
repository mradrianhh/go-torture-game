package models

// Level holds all the information regarding a level.
type Level struct {
	truth string
	enemy Enemy
}

// LevelList holds the global collection of levels in the game.
var LevelList []Level

func init() {
	LevelList = []Level{
		{truth: "True Truth 1", enemy: EnemyList[0]},
		{truth: "True Truth 2", enemy: EnemyList[1]},
		{truth: "True Truth 3", enemy: EnemyList[2]},
		{truth: "True Truth 4", enemy: EnemyList[3]},
		{truth: "True Truth 5", enemy: EnemyList[4]},
		{truth: "True Truth 6", enemy: EnemyList[5]},
		{truth: "True Truth 7", enemy: EnemyList[6]},
		{truth: "True Truth 8", enemy: EnemyList[7]},
		{truth: "True Truth 9", enemy: EnemyList[8]},
		{truth: "True Truth 10", enemy: EnemyList[9]},
	}
}

// GetTruth getter.
func (level *Level) GetTruth() string {
	return level.truth
}

// SetTruth setter.
func (level *Level) SetTruth(truth string) {
	level.truth = truth
}

// GetEnemy getter.
func (level *Level) GetEnemy() Enemy {
	return level.enemy
}

// SetEnemy setter.
func (level *Level) SetEnemy(enemy Enemy) {
	level.enemy = enemy
}
