package main

type GameEvent = string

const (
	EventReset    GameEvent = "reset"
	EventStart              = "begin"
	EventEnd                = "end"
	EventMakeMove           = "move"
	EventPrompt             = "prompt"
)
