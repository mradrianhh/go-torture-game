// Package main represents the entry point of the game.
//
// Details:
// 	- As a singleplayer game, the "target *Player" can be excluded as there will ever only be one global target in the game realm.
//		Upon extension to a multiplayer game, the attack method requires a target, not just an executioner.
//
// Decisions:
// 	1. Should the "game" make the attacks, or the "player"'s. As in: upon attack, should the game perform the attack method, or should it stand
// 		as an extension method on the player.
//	1.a Upon further consideration, the development tends me towards letting a global entity execute the actions.
//		By using an event-system, I can notify the global manager from the model-classes.
//
// Objectives:
// 	- The goal of the game is to get the enemy to reveal the truth.
//
// Challenges:
// 	- The difficulty of the game is set through your lack of knowledge regarding the enemy's mentality or health. If the mentality reaches 0, you get the truth,
// 		but how will you ever know? And if you push too hard, their health will reach 0, and you will lose.
package main
