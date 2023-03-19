package robot

import (
	"fmt"
	"time"
)

type Action string

// Actions3 describes robot action in 3rd scenerio
type Action3 struct{ name, command string }

var N, E, S, W Dir = 0, 1, 2, 3

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	default:
		return
	}
}

func (d Dir) String() string {
	v := int(d)
	return fmt.Sprintf("%v", v)
}

func Right() {
	if Step1Robot.Dir != W {
		Step1Robot.Dir++
	} else {
		Step1Robot.Dir = N
	}
}

func Left() {
	if Step1Robot.Dir != N {
		Step1Robot.Dir--
	} else {
		Step1Robot.Dir = W
	}
}

func StartRobot(cmd chan Command, act chan Action) {
	for c := range cmd {
		act <- Action(c)
	}
}

func Room(rect Rect, step Step2Robot, act chan Action, req chan Step2Robot) {
	ticker := time.NewTicker(100 * time.Millisecond)
	position := step.Pos
	direction := step.Dir
	for {
		select {
		case a := <-act:
			applyMove(a, rect, &direction, &position)
		case <-ticker.C:
			req <- Step2Robot{direction, position}
			return
		}
	}
}

func applyMove(a Action, rect Rect, direction *Dir, position *Pos) bool {
	switch a {
	case "A":
		m := newMove(*direction, *position)
		if position.Easting+m.Easting <= rect.Max.Easting && position.Northing+m.Northing <= rect.Max.Northing {
			if position.Easting+m.Easting >= rect.Min.Easting && position.Northing+m.Northing >= rect.Min.Northing {
				position.Easting += m.Easting
				position.Northing += m.Northing
				return true
			}
		}
		return false
	case "R":
		if *direction != W {
			*direction++

		} else {
			*direction = N
		}
	case "L":
		if *direction != N {
			*direction--
		} else {
			*direction = W
		}
	}
	return true
}

func newMove(d Dir, pos Pos) Pos {
	switch d {
	case N:
		return Pos{0, 1}
	case E:
		return Pos{1, 0}
	case S:
		return Pos{0, -1}
	case W:
		return Pos{-1, 0}
	default:
		return Pos{}
	}
}

func inCommands(c rune) bool {
	switch string(c) {
	case "A", "R", "L":
		return true
	}
	return false
}

func StartRobot3(name, action string, act chan<- Action3, log chan<- string) {
	if name == "" {
		log <- "wrong name"
		return
	}
	for _, c := range action {
		if !inCommands(c) {
			log <- "wrong"
			return
		}
		act <- Action3{name: name, command: string(c)}
	}
}

func Room3(rect Rect, robots []Step3Robot, act <-chan Action3, rep chan<- []Step3Robot, log chan<- string) {
	validatedRobots := make(map[string]Step2Robot)
	occupiedPositions := make(map[RU]RU)
	for _, r := range robots {
		if north, ok := occupiedPositions[r.Step2Robot.Pos.Easting]; ok && north == r.Step2Robot.Pos.Northing {
			log <- "conflict"
			continue
		}
		if _, ok := validatedRobots[r.Name]; ok {
			log <- "duplicated"
			continue
		}
		if !rect.inside(r.Pos) {
			log <- "outside"
			continue
		}
		if r.Name == "" {
			continue
		}
		occupiedPositions[r.Step2Robot.Pos.Easting] = r.Step2Robot.Pos.Northing
		validatedRobots[r.Name] = r.Step2Robot
	}

	ticker := time.NewTicker(100 * time.Millisecond)
OUTER:
	for {
		select {
		case action := <-act:
			step2R, ok := validatedRobots[action.name]
			if !ok {
				log <- "unknown"
				continue OUTER
			}
			if ok := applyMove(Action(action.command), rect, &step2R.Dir, &step2R.Pos); !ok {
				log <- "wall"
				continue OUTER
			}
		INNER:
			for name, step := range validatedRobots {
				if name == action.name {
					continue INNER
				}
				if step.Easting == step2R.Easting && step.Northing == step2R.Northing {
					log <- "colision"
					continue OUTER
				}
			}
			validatedRobots[action.name] = step2R
		case <-ticker.C:
			robotsResponse := make([]Step3Robot, len(validatedRobots))
			i := 0
			for name, step2R := range validatedRobots {
				robotsResponse[i] = Step3Robot{name, step2R}
				i++
			}
			rep <- robotsResponse
		}

	}
}

func (rect Rect) inside(position Pos) bool {
	if position.Easting > rect.Max.Easting || position.Northing > rect.Max.Northing {
		return false
	}
	if position.Easting < rect.Min.Easting || position.Northing < rect.Min.Northing {
		return false
	}
	return true
}
