package game

import (
	"fmt"
	"os"

	"github.com/mradrianhh/go-torture-game/events"
	"github.com/mradrianhh/go-torture-game/models"
)

var response int

// Start ...
func Start() {
	showMainMenu()
}

func showMainMenu() {
	for {
		fmt.Println("1 - Start Game | 2 - Exit Game")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
				fmt.Println("Start Game")
			case 2:
				os.Exit(0)
			default:
				fmt.Println("Couldn't understand. Please try again.")
			}
		}
	}
}

// Manager ...
var Manager events.EventListener

// CurrentLevel is the current level the player is on.
var CurrentLevel models.Level

// CurrentPlayer is the player currently playing.
var CurrentPlayer models.Player

func init() {
	setCurrentLevel(1)
	Manager.Events = make(map[string]func())
	Manager.Events["torture"] = torture
	go Manager.Listen(events.GlobalEventStream)
}

// Sets the current level.
//
// Example - Set to level 1:
// 	setCurrentLevel(1) => CurrentLevel = models.LevelList[0].
func setCurrentLevel(level int) {
	CurrentLevel = models.LevelList[level-1]
}

func torture() {
	enemy := CurrentLevel.GetEnemy()
	fmt.Printf("Torturing %s...\n", enemy.GetName())
}
