package main

import (
	"flag"
	"fmt"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/easing"
	"github.com/gravestench/director/pkg/systems/scene"
	"github.com/gravestench/director/pkg/systems/tween"
	"github.com/gravestench/mathlib"
	"image/color"
	"math"
	"math/rand"
	"time"
)

type TestScene struct {
	scene.Scene
	*twitch.Client
	twitch struct {
		userName *string
		oauthKey *string
		channel *string
	}
	userColors map[string]color.Color
}

func (s *TestScene) Init(_ *akara.World) {
	s.parseFlags()
	s.setupClient()

	s.userColors = make(map[string]color.Color)

	go func() {
		s.connect() // this is blocking, so we put in a goroutine
	}()
}

func (s *TestScene) parseFlags() {
	s.twitch.userName = flag.String("u", "", "username")
	s.twitch.oauthKey = flag.String("o", "", "oath key, like \"oauth:kjahsdkjahsdkjahsd\"")
	s.twitch.channel = flag.String("c", "", "channel")

	flag.Parse()
}

func (s *TestScene) setupClient() {
	if *s.twitch.oauthKey == "" {
		panic("need an oauth key")
	}

	if *s.twitch.userName == "" {
		panic("need a username")
	}

	if *s.twitch.channel == "" {
		panic("need a channel name")
	}

	// or client := twitch.NewAnonymousClient() for an anonymous user (no write capabilities)
	s.Client = twitch.NewClient(*s.twitch.userName, *s.twitch.oauthKey)

	s.OnPrivateMessage(func(msg twitch.PrivateMessage) {
		s.newMessage(msg.User.Name, fmt.Sprintf("%v", msg.Message))
	})

	s.OnUserJoinMessage(func(msg twitch.UserJoinMessage) {
		s.newMessage(msg.User, fmt.Sprintf("%v has joined the chat!", msg.User))
	})

	s.Client.Join(*s.twitch.channel)
}

func (s *TestScene) getUserColor(name string) color.Color {
	c, found := s.userColors[name]
	if !found {
		c = randColor()
		s.userColors[name] = c
	}

	return c
}

func (s *TestScene) newMessage(name, msg string) {
	c := s.getUserColor(name)

	x, y := s.Window.Width/2, s.Window.Height/2
	fontSize := s.Window.Height / 20

	label := s.Add.Label(msg, x, y, fontSize, "", c)

	trs, found := s.Components.Transform.Get(label)
	if !found {
		return
	}

	opacity, found := s.Components.Opacity.Get(label)
	if !found {
		opacity = s.Components.Opacity.Add(label)
	}

	text, found := s.Components.Text.Get(label)
	if !found {
		return
	}

	startX, startY, endX, endY := s.randomStartEnd()
	cx, cy := s.Window.Width/2, s.Window.Height/2

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
			s.RemoveEntity(label)
		})

		s.Tweens.New(tb2)
	})

	s.Tweens.New(tb)

	text.String = msg
}

func (s *TestScene) connect() {
	err := s.Connect()
	if err != nil {
		panic(err)
	}
}

func (s *TestScene) IsInitialized() bool {
	return s.Client != nil
}

func (s *TestScene) Update() {
	// no op
}

func (s *TestScene) randomStartEnd() (x1, y1, x2, y2 int) {
	const (
		maxDegree = 360
	)

	dStart := rand.Intn(maxDegree)
	dEnd := (dStart + 180) % maxDegree

	distance := 1.5 * float64(s.Window.Width)

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
