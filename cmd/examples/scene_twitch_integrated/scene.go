package main

import (
	"flag"
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/gravestench/akara"
	"github.com/gravestench/mathlib"

	// and the other was easy to use at the time i initially wrote this example
	// using two twitch libraries because one of them provided a method for pulling emotes
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/nicklaw5/helix"

	. "github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/easing"
	"github.com/gravestench/director/pkg/systems/scene"
	"github.com/gravestench/director/pkg/systems/tween"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type testScene struct {
	scene.Scene
	ircClient   *twitch.Client
	helixClient *helix.Client
	twitch      struct {
		userName *string
		oauthKey *string
		channel *string
		clientid *string
		clientSecret *string
		userAccessToken *string
	}
	userColors map[string]color.Color
	emoteMap map[string]string
}

func (scene *testScene) Init(_ *akara.World) {
	scene.parseFlags() // the command line flags have all the twitch api stuff
	scene.setupClients() // we set up the two titch client instances
	scene.initEmotes() // set up a mapping of emote strings to URLs

	// users in chat will e
	scene.userColors = make(map[string]color.Color)

	go func() {
		scene.connect() // this is blocking, so we put in a goroutine
	}()
}

func (scene *testScene) initEmotes() {
	scene.emoteMap = make(map[string]string)

	allEmotes, err := scene.helixClient.GetGlobalEmotes()
	if err != nil {
		panic(err)
	}

	for _, emote := range allEmotes.Data.Emotes {
		scene.emoteMap[emote.Name] = emote.Images.Url4x
	}
}

func (scene *testScene) parseFlags() {
	scene.twitch.userName = flag.String("user", "", "username")
	scene.twitch.channel = flag.String("channel", "", "channel")
	scene.twitch.oauthKey = flag.String("oauth", "", "oath key")
	scene.twitch.clientid = flag.String("clientid", "", "client id token")
	scene.twitch.clientSecret = flag.String("clientsecret", "", "client secret")
	scene.twitch.userAccessToken = flag.String("useraccesstoken", "", "client id token")

	flag.Parse()
}

func (scene *testScene) setupClients() {
	if *scene.twitch.oauthKey == "" {
		panic("need an oauth key")
	}

	if *scene.twitch.userName == "" {
		panic("need a username")
	}

	if *scene.twitch.channel == "" {
		panic("need a channel name")
	}

	if *scene.twitch.clientid == "" {
		panic("need an client id")
	}

	if *scene.twitch.clientSecret == "" {
		panic("need a user access token see https://github.com/twitchdev/authentication-go-sample")
	}

	if *scene.twitch.userAccessToken == "" {
		panic("need an client id")
	}


	client, _ := helix.NewClient(&helix.Options{
		ClientID:        *scene.twitch.clientid,
		ClientSecret: *scene.twitch.clientSecret,
		UserAccessToken: *scene.twitch.userAccessToken,
		RedirectURI:     "http://localhost",
	})

	if isValid, _, err := client.ValidateToken(*scene.twitch.userAccessToken); !isValid {
		panic(err)
	}

	scene.helixClient = client

	// or client := twitch.NewAnonymousClient() for an anonymous user (no write capabilities)
	scene.ircClient = twitch.NewClient(*scene.twitch.userName, "oauth:" + *scene.twitch.oauthKey)

	scene.ircClient.OnPrivateMessage(func(msg twitch.PrivateMessage) {
		scene.newMessage(msg.User.Name, msg.Message)
	})

	scene.ircClient.OnUserJoinMessage(func(msg twitch.UserJoinMessage) {
		scene.newMessage(msg.User, fmt.Sprintf("%v has joined the chat!", msg.User))
	})

	scene.ircClient.Join(*scene.twitch.channel)
}

func (scene *testScene) getUserColor(name string) color.Color {
	c, found := scene.userColors[name]
	if !found {
		c = randColor()
		scene.userColors[name] = c
	}

	return c
}

func (scene *testScene) newMessage(name, msg string) {
	c := scene.getUserColor(name)

	x, y := scene.Window.Width/2, scene.Window.Height/2
	fontSize := scene.Window.Height / 20

	var entity Entity

	if emoteURL, found := scene.emoteMap[msg]; found {
		entity = scene.Add.Image(emoteURL, x, y)
	} else {
		entity = scene.Add.Label(msg, x, y, fontSize, "", c)

		text, found := scene.Components.Text.Get(entity)
		if !found {
			return
		}

		text.String = msg
	}

	trs, found := scene.Components.Transform.Get(entity)
	if !found {
		return
	}

	opacity, found := scene.Components.Opacity.Get(entity)
	if !found {
		opacity = scene.Components.Opacity.Add(entity)
	}

	startX, startY, endX, endY := scene.randomStartEnd()
	cx, cy := scene.Window.Width/2, scene.Window.Height/2

	rcx, rcy := cx/2 + rand.Intn(cx), cy/2 + rand.Intn(cy)

	tb := tween.NewBuilder()

	tb.Time(time.Second * 6)
	tb.Ease(easing.ElasticOut, []float64{0.5, 0.85, 0.5})
	tb.OnStart(func() {
		trs.Translation.Set(float64(startX), float64(startY), trs.Translation.Z)
		opacity.Value = 0
	})

	tb.OnUpdate(func(progress float64) {
		opacity.Value = progress
		x := float64(startX) + (progress * float64(rcx - startX))
		y := float64(startY) + (progress * float64(rcy - startY))

		trs.Translation.Set(x, y, trs.Translation.Z)
	})

	tb.OnComplete(func() {
		tb2 := tween.NewBuilder()

		tb2.Time(time.Second * 6)
		tb2.Ease(easing.ElasticOut, []float64{0.5, 0.85, 0.5})
		tb2.OnStart(func() {
			trs.Translation.Set(float64(startX), float64(startY), trs.Translation.Z)
			opacity.Value = 0
		})

		tb2.OnUpdate(func(progress float64) {
			opacity.Value = 1 - progress
			x := float64(rcx) + (progress * float64(endX - rcx))
			y := float64(rcy) + (progress * float64(endY - rcy))

			trs.Translation.Set(x, y, trs.Translation.Z)
		})

		tb2.OnComplete(func() {
			scene.RemoveEntity(entity)
		})

		scene.Tweens.New(tb2)
	})

	scene.Tweens.New(tb)
}

func (scene *testScene) resizeCameraWithWindow() {
	for _, e := range scene.Cameras {
		rt, found := scene.Components.RenderTexture2D.Get(e)
		if !found {
			continue
		}

		if int(rt.Texture.Width) != scene.Window.Width || int(rt.Texture.Height) != scene.Window.Height {
			t := rl.LoadRenderTexture(int32(scene.Window.Width), int32(scene.Window.Height))
			rt.RenderTexture2D = &t
		}
	}
}

func (scene *testScene) connect() {
	err := scene.ircClient.Connect()
	if err != nil {
		panic(err)
	}
}

func (scene *testScene) IsInitialized() bool {
	return scene.ircClient != nil
}

func (scene *testScene) Update() {
	scene.resizeCameraWithWindow()
}

func (scene *testScene) randomStartEnd() (x1, y1, x2, y2 int) {
	const (
		maxDegree = 360
	)

	dStart := rand.Intn(maxDegree)
	dEnd := (dStart + 180) % maxDegree

	distance := 1.5 * float64(scene.Window.Width)

	x1 = int(math.Sin(float64(dStart) * mathlib.DegreesToRadians) * distance)
	y1 = int(math.Cos(float64(dStart) * mathlib.DegreesToRadians) * distance)

	x2 = int(math.Sin(float64(dEnd) * mathlib.DegreesToRadians) * distance)
	y2 = int(math.Cos(float64(dEnd) * mathlib.DegreesToRadians) * distance)

	return x1, y1, x2, y2
}

func randColor() color.Color {
	return &color.RGBA{
		R: math.MaxUint8 - uint8(rand.Intn(math.MaxUint8/6)),
		G: math.MaxUint8 - uint8(rand.Intn(math.MaxUint8/6)),
		B: math.MaxUint8 - uint8(rand.Intn(math.MaxUint8/6)),
		A: math.MaxUint8,
	}
}
