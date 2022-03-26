package conrtoller

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/test/engine"
	"example.com/test/view"
)

func CommandListener(w *engine.World) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		//for windows :
		text = strings.Replace(text, "\r", "", -1)

		if text == "exit" {
			fmt.Println("You exit the game")
			return
		}
		message := ParceCommand(w, text)
		// if strings.Compare("hi", text) == 0 {
		// 	fmt.Println("hello, Yourself")
		fmt.Println(message)
		// }

	}
}

type Commander interface {
	DoCommand() string
}

type Command func(w *engine.World, param []string) (*engine.World, string)

type ComandAlias int

const (
	cmd_north ComandAlias = iota
	cmd_south
	cmd_west
	cmd_east
	cmd_enter
	cmd_out

	cmd_look
	cmd_help
	cmd_inventory
	cmd_take
	cmd_drop
	cmd_talk
	cmd_exit
)

var CommandNames = map[ComandAlias]string{
	cmd_north: "north",
	cmd_south: "south",
	cmd_west:  "west",
	cmd_east:  "east",
	cmd_enter: "enter",
	cmd_out:   "out",
	cmd_look:  "look",
	cmd_exit:  "exit",
	cmd_help:  "help",
}

func ParceCommand(w *engine.World, commandText string) string {
	//TODO implement dificult commands from 2-3 words
	CommandsMaper := make(map[string]Command)

	CommandsMaper[CommandNames[cmd_north]] = MoveNorth
	CommandsMaper[CommandNames[cmd_look]] = LookAround
	CommandsMaper[CommandNames[cmd_south]] = MoveSouth
	CommandsMaper[CommandNames[cmd_west]] = MoveWest
	CommandsMaper[CommandNames[cmd_east]] = MoveEast
	CommandsMaper[CommandNames[cmd_enter]] = MoveIn
	CommandsMaper[CommandNames[cmd_out]] = MoveOut
	CommandsMaper[CommandNames[cmd_exit]] = nil
	CommandsMaper[CommandNames[cmd_help]] = Help

	action := CommandsMaper[commandText]

	if action != nil {
		_, text := action(w, []string{commandText})
		return text
	}

	return "Command not found"

}

// type Commander interface{
// 	Validate(W)
// }

func MoveNorth(w *engine.World, param []string) (*engine.World, string) {
	//validate(w,)
	if w.Player() != nil {
		Hero := w.Player()
		if engine.IsNil(Hero.Mobs().Location().North()) {
			return w, "Command Move North : north way is blocked" // error = TODO made comand . name , etc
		}
		Hero.Move(engine.North)
		return w, "Command Move North : Moved"
	}
	return w, "Command Move North : Invalid Player"
}

func MoveSouth(w *engine.World, param []string) (*engine.World, string) {
	//validate(w,)
	comandName := "Command Move South"
	if w.Player() != nil {
		Hero := w.Player()
		if engine.IsNil(Hero.Mobs().Location().South()) {
			return w, comandName + " : South way is blocked" // error = TODO made comand . name , etc
		}
		Hero.Move(engine.South)
		return w, comandName + " : Moved"
	}
	return w, comandName + " : Invalid Player"
}

func MoveWest(w *engine.World, param []string) (*engine.World, string) {
	//validate(w,)
	comandName := "Command Move West"
	if w.Player() != nil {
		Hero := w.Player()
		if engine.IsNil(Hero.Mobs().Location().West()) {
			return w, comandName + " : West way is blocked" // error = TODO made comand . name , etc
		}
		Hero.Move(engine.West)
		return w, comandName + " : Moved"
	}
	return w, comandName + " : Invalid Player"
}

func MoveEast(w *engine.World, param []string) (*engine.World, string) {
	//validate(w,)
	comandName := "Command Move East"
	if w.Player() != nil {
		Hero := w.Player()
		if engine.IsNil(Hero.Mobs().Location().East()) {
			return w, comandName + " : East way is blocked" // error = TODO made comand . name , etc
		}
		Hero.Move(engine.East)
		return w, comandName + " : Moved"
	}
	return w, comandName + " : Invalid Player"
}

func MoveIn(w *engine.World, param []string) (*engine.World, string) {
	//validate(w,)
	comandName := "Command Move In"
	if w.Player() != nil {
		Hero := w.Player()
		if engine.IsNil(Hero.Mobs().Location().Enter()) {
			return w, comandName + " : Can not enter. Way is blocked" // error = TODO made comand . name , etc
		}
		Hero.Move(engine.Enter)
		return w, comandName + " : Moved"
	}
	return w, comandName + " : Invalid Player"
}

func MoveOut(w *engine.World, param []string) (*engine.World, string) {
	//validate(w,)
	comandName := "Command Move Out"
	if w.Player() != nil {
		Hero := w.Player()
		if engine.IsNil(Hero.Mobs().Location().Exit()) {
			return w, comandName + " : Can not exit. Way is blocked" // error = TODO made comand . name , etc
		}
		Hero.Move(engine.Out)
		return w, comandName + " : Moved"
	}
	return w, comandName + " : Invalid Player"
}

func LookAround(w *engine.World, param []string) (*engine.World, string) {
	//validate(w,)
	if w.Player() != nil {
		Hero := w.Player()

		look := view.Mob{*Hero}

		// if engine.IsNil(Hero.Mobs().Location.().North()) {
		// 	return w, "Command Look Around : north way is blocked" // error = TODO made comand . name , etc
		// }
		// Hero.Move(engine.North)
		return w, "Command Look Around : " + look.View()
	}
	return w, "Command Look Around : Invalid Player"
}

func Help(w *engine.World, param []string) (*engine.World, string) {
	//TODO implement it
	comandsDescription := map[ComandAlias]string{
		cmd_north: "Move Hero to North.",
		cmd_south: "Move Hero to South.",
		cmd_west:  "Move Hero to West.",
		cmd_east:  "Move Hero to East.",
		cmd_enter: "Hero will enter into cell or building",
		cmd_out:   "Hero will go out from cell or building",
		cmd_look:  "Hero will look around. output with explanation of current position, buildings, exits, directions, ect.",
		cmd_help:  "This help",
		cmd_exit:  "Exit game with no save progress",
	}

	//TODO : add detailed help
	ret := "help on comands \r\n"
	for k, v := range comandsDescription {
		ret += CommandNames[k] + " : " + v + "\r\n"
	}
	return w, ret + "\r\n"

	//validate(w,)
	// if w.Player() != nil {
	// 	Hero := w.Player()

	// 	look := view.Mob{*Hero}

	// 	// if engine.IsNil(Hero.Mobs().Location.().North()) {
	// 	// 	return w, "Command Look Around : north way is blocked" // error = TODO made comand . name , etc
	// 	// }
	// 	// Hero.Move(engine.North)
	// 	return w, "Command Look Around : " + look.View()
	// }
	return w, "Command Look Around : Invalid Player"
}

func main() {
	//TODO implement commands
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

	}

}
