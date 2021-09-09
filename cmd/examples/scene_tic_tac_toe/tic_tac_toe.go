package main

import (
	"fmt"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/systems/scene"
	"math/rand"
)

type TicTacToe struct {
	scene.Scene
	components struct {
		gameState GameStateFactory
	}
	gameStates *akara.Subscription
}

func (scene *TicTacToe) Update() {
	scene.eval()
}

func (scene *TicTacToe) Init(_ *akara.World) {
	scene.InjectComponent(&GameState{}, &scene.components.gameState.ComponentFactory)
	scene.bindEvents()
	scene.components.gameState.Add(scene.Director.NewEntity())
	scene.subscribe()
	scene.reset()
}

func (scene *TicTacToe) bindEvents() {
	scene.Events.On(EventReset, func(args ...interface{}) {
		scene.reset()
	})

	scene.Events.On(EventStart, func(args ...interface{}) {
		scene.start()
	})

	scene.Events.On(EventEnd, func(args ...interface{}) {
		scene.end()
	})

	scene.Events.On(EventMakeMove, func(args ...interface{}) {
		if len(args) < 1 {
			return
		}

		where, ok := args[0].(int)
		if !ok {
			return
		}

		scene.movePlayer(where)
	})

	scene.Events.On(EventPrompt, func(args ...interface{}) {
		if len(args) < 1 {
			return
		}

		message, ok := args[0].(string)
		if !ok {
			return
		}

		scene.prompt(message)
	})
}

func (scene *TicTacToe) subscribe() {
	filter := scene.Director.NewComponentFilter()
	filter.Require(&GameState{})

	scene.gameStates = scene.Director.AddSubscription(filter.Build())
}

func (scene *TicTacToe) IsInitialized() bool {
	return scene.gameStates != nil
}

func (scene *TicTacToe) getState() *GameState {
	for _, eid := range scene.gameStates.GetEntities()[:1] {
		s, found := scene.components.gameState.Get(eid)
		if !found {
			break
		}

		return s
	}

	return nil
}

func (scene *TicTacToe) reset() {
	s := scene.getState()

	for idx := range s.state.cells {
		s.cells[idx] = empty
	}

	s.turn = Player(rand.Intn(int(numPlayers)))

	scene.Events.Emit(EventStart)
}

func (scene *TicTacToe) start() {
	s := scene.getState()

	msg := fmt.Sprintf("start, player %s!", s.turn)

	scene.Events.Emit(EventPrompt, msg)
}

func (scene *TicTacToe) end() {}

func (scene *TicTacToe) movePlayer(where int) {}

func (scene *TicTacToe) prompt(text string) {}

var winConditions = [][]int{
	{},
}

func (scene *TicTacToe) eval() Player {
	s := scene.getState()

	for idx := range s.cells {
		if s.cells[idx] == empty {
			return empty
		}
	}

	// if we get here, all cells have been filled
	return scene.checkForWinner()
}

func (scene *TicTacToe) checkForWinner() Player {
	s := scene.getState()

	for _, c := range winConditions {
		if scene.winConditionMet(s, c, PlayerO) {
			return PlayerO
		}

		if scene.winConditionMet(s, c, PlayerX) {
			return PlayerX
		}
	}

	return empty
}

func (scene *TicTacToe) winConditionMet(s *GameState, condition []int, p Player) bool {
	isWinning := false

	for idx := range condition {
		isWinning = isWinning || s.cells[condition[idx]] != p
	}

	return isWinning
}
