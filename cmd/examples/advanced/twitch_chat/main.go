package main

import (
	"flag"

	"github.com/gravestench/director"
)

/*
	to get.go this set up properly you are going to need to do the following:
	* get.go an OATH token from https://twitchapps.com/tmi
	* set up an "application" for twitch development at https://dev.twitch.tv/console
*/

func main() {
	s := &testScene{}

	// HACK this should not need to be called, but the flag package uses a global singleton
	// which messes everything up if flag.Parse() is called multiple times.
	// at time of writing, Director calls it internally ;_:
	flag.StringVar(&s.twitch.userName, "user", "", "username")
	flag.StringVar(&s.twitch.channel, "channel", "", "channel")
	flag.StringVar(&s.twitch.oauthKey, "oauth", "", "oath key")
	flag.StringVar(&s.twitch.clientid, "clientid", "", "client id token")
	flag.StringVar(&s.twitch.clientSecret, "clientsecret", "", "client secret")

	flag.Parse()

	d := director.New()

	d.AddScene(s)

	if err := d.Run(); err != nil {
		panic(err)
	}
}
