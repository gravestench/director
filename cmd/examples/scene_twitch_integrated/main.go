package main

import (
	director "github.com/gravestench/director/pkg"
)

/*
	to get this set up properly you are going to need to do the following:
	* get an OATH token from https://twitchapps.com/tmi
	* set up an "application" for twitch development at https://dev.twitch.tv/console
*/

func main() {
	d := director.New()

	d.Window.Width = 1024
	d.Window.Height = 768
	d.TargetFPS = 60

	d.AddScene(&testScene{})

	if err := d.Run(); err != nil {
		panic(err)
	}
}

