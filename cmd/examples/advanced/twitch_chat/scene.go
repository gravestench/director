package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gravestench/akara"

	"github.com/faiface/mainthread"
	"github.com/gravestench/mathlib"

	. "github.com/gravestench/director"
	"github.com/gravestench/director/cmd/examples/advanced/twitch_chat/access_token"

	// using two twitch libraries because one of them provided a method for pulling emotes
	// and the other was easy to use at the time i initially wrote this example
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/nicklaw5/helix"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gravestench/director/pkg/easing"
)

const (
	fmtBTTVUserEmoteEndpoint = "https://api.betterttv.net/3/cached/users/twitch/%s"
	fmtBTTVEmoteURL          = "https://cdn.betterttv.net/emote/%s/3x"
)

type testScene struct {
	Scene
	ircClient   *twitch.Client
	helixClient *helix.Client
	twitch      struct {
		userName        string
		oauthKey        string
		channel         string
		clientid        string
		clientSecret    string
		userAccessToken string
		userid          string
	}
	userColors map[string]color.Color
	emoteMap   map[string]string
}

func (scene *testScene) Key() string {
	return "twitch integration test"
}

func (scene *testScene) Init(_ *World) {
	scene.parseFlags() // the command line flags have all the twitch api stuff

	scene.twitch.userAccessToken = access_token.Get(scene.twitch.clientid, scene.twitch.clientSecret)

	scene.setupClients() // we set up the two titch client instances
	scene.initEmotes()   // set up a mapping of emote strings to URLs

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
		log.Println(err)
		os.Exit(1)
	}

	for _, emote := range allEmotes.Data.Emotes {
		scene.emoteMap[emote.Name] = emote.Images.Url4x
	}

	scene.initBTTVEmotes()
}

func (scene *testScene) initBTTVEmotes() {
	res, err := http.Get(fmt.Sprintf(fmtBTTVUserEmoteEndpoint, scene.twitch.userid))
	if err != nil {
		return
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	userEmoteInfo := &BTTVUserEmotesDescriptor{}
	err = json.Unmarshal(data, userEmoteInfo)
	if err != nil {
		return
	}

	if userEmoteInfo.SharedEmotes == nil {
		return
	}

	for _, emote := range userEmoteInfo.SharedEmotes {
		emoteURL := fmt.Sprintf(fmtBTTVEmoteURL, emote.ID)
		scene.emoteMap[emote.Code] = emoteURL
	}
}

func (scene *testScene) parseFlags() {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	f.StringVar(&scene.twitch.userName, "user", "", "pick your username")
	f.StringVar(&scene.twitch.channel, "channel", "", "set the channel to monitor (your username if you want to use this with OBS)")
	f.StringVar(&scene.twitch.oauthKey, "oauth", "", "oath key (see https://twitchapps.com/tmi/)")
	f.StringVar(&scene.twitch.clientid, "clientid", "", "client id token (see https://dev.twitch.tv/console)")
	f.StringVar(&scene.twitch.clientSecret, "clientsecret", "", "client secret (see https://dev.twitch.tv/console)")

	if err := f.Parse(os.Args[1:]); err != nil {
		f.Usage()
		os.Exit(1)
	}

	if scene.twitch.oauthKey == "" ||
		scene.twitch.userName == "" ||
		scene.twitch.channel == "" ||
		scene.twitch.clientid == "" ||
		scene.twitch.clientSecret == "" {
		f.Usage()
		os.Exit(1)
	}

	if strings.HasPrefix(scene.twitch.oauthKey, "oauth:") {
		scene.twitch.oauthKey = strings.ReplaceAll(scene.twitch.oauthKey, "oauth:", "")
	}
}

func (scene *testScene) setupClients() {
	client, _ := helix.NewClient(&helix.Options{
		ClientID:        scene.twitch.clientid,
		ClientSecret:    scene.twitch.clientSecret,
		UserAccessToken: scene.twitch.userAccessToken,
		RedirectURI:     "http://localhost",
	})

	isValid, _, err := client.ValidateToken(scene.twitch.userAccessToken)
	if !isValid {
		log.Println("could not validate using user access token")
		os.Exit(1)
	}

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	scene.helixClient = client

	usersResponse, err := scene.helixClient.GetUsers(&helix.UsersParams{
		Logins: []string{scene.twitch.userName},
	})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if len(usersResponse.Data.Users) < 1 {
		log.Println("expecting to yield a user from twitch client")
		os.Exit(1)
	}

	scene.twitch.userid = usersResponse.Data.Users[0].ID

	// or client := twitch.NewAnonymousClient() for an anonymous user (no write capabilities)
	scene.ircClient = twitch.NewClient(scene.twitch.userName, "oauth:"+scene.twitch.oauthKey)

	scene.ircClient.OnPrivateMessage(func(msg twitch.PrivateMessage) {
		scene.newMessage(msg.User.Name, msg.Message)
	})

	scene.ircClient.OnUserJoinMessage(func(msg twitch.UserJoinMessage) {
		scene.newMessage(msg.User, fmt.Sprintf("%v has joined the chat!", msg.User))
	})

	scene.ircClient.Join(scene.twitch.channel)
}

func (scene *testScene) getUserColor(name string) color.Color {
	c, found := scene.userColors[name]
	if !found {
		c = newUserColor()
		scene.userColors[name] = c
	}

	return c
}

func (scene *testScene) newMessage(name, msg string) {
	fmt.Printf("%s: %s\n", name, msg)

	c := scene.getUserColor(name)

	rWidth := scene.Sys.Renderer.Window.Width
	rHeight := scene.Sys.Renderer.Window.Height

	x, y := rWidth/2, rHeight/2
	fontSize := rHeight / 20

	var entity akara.EID

	emoteURL, emoteFound := scene.emoteMap[msg]
	if emoteFound {
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

	if emoteFound {
		trs.Scale.Set(3, 3, 3)
	}

	opacity, found := scene.Components.Opacity.Get(entity)
	if !found {
		opacity = scene.Components.Opacity.Add(entity)
	}

	startX, startY, endX, endY := scene.randomStartEnd()
	cx, cy := rWidth/2, rHeight/2

	rcx, rcy := cx/2+rand.Intn(cx), cy/2+rand.Intn(cy)

	tb := scene.Sys.Tweens.New()

	tb.Time(time.Second * 6)
	tb.Ease(easing.ElasticOut, []float64{0.5, 0.85, 0.5})
	tb.OnStart(func() {
		trs.Translation.Set(float64(startX), float64(startY), trs.Translation.Z)
		opacity.Value = 0
	})

	tb.OnUpdate(func(progress float64) {
		opacity.Value = progress

		x := float64(startX) + (progress * float64(rcx-startX))
		y := float64(startY) + (progress * float64(rcy-startY))

		trs.Translation.Set(x, y, trs.Translation.Z)
	})

	tb.OnComplete(func() {
		tb2 := scene.Sys.Tweens.New()

		tb2.Time(time.Second * 6)
		tb2.Ease(easing.ElasticOut, []float64{0.5, 0.85, 0.5})
		tb2.OnStart(func() {
			trs.Translation.Set(float64(startX), float64(startY), trs.Translation.Z)
			opacity.Value = 1
		})

		tb2.OnUpdate(func(progress float64) {
			opacity.Value = 1 - progress
			x := float64(rcx) + (progress * float64(endX-rcx))
			y := float64(rcy) + (progress * float64(endY-rcy))

			trs.Translation.Set(x, y, trs.Translation.Z)
		})

		tb2.OnComplete(func() {
			scene.RemoveEntity(entity)
		})

		scene.Sys.Tweens.Add(tb2)
	})

	scene.Sys.Tweens.Add(tb)
}

func (scene *testScene) resizeCameraWithWindow() {
	rWidth := scene.Sys.Renderer.Window.Width
	rHeight := scene.Sys.Renderer.Window.Height

	for _, e := range scene.Viewports {
		vp, found := scene.Components.Viewport.Get(e)
		if !found {
			continue
		}

		vprt, found := scene.Components.RenderTexture2D.Get(e)
		if !found {
			continue
		}

		camrt, found := scene.Components.RenderTexture2D.Get(vp.CameraEntity)
		if !found {
			continue
		}

		if int(vprt.Texture.Width) != rWidth || int(vprt.Texture.Height) != rHeight {
			mainthread.Call(func() {
				t := rl.LoadRenderTexture(int32(rWidth), int32(rHeight))
				vprt.RenderTexture2D = &t
			})
		}

		if int(camrt.Texture.Width) != rWidth || int(camrt.Texture.Height) != rHeight {
			mainthread.Call(func() {
				t := rl.LoadRenderTexture(int32(rWidth), int32(rHeight))
				camrt.RenderTexture2D = &t
			})
		}
	}
}

func (scene *testScene) connect() {
	err := scene.ircClient.Connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
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

	rWidth := scene.Sys.Renderer.Window.Width
	distance := 1.5 * float64(rWidth)

	x1 = int(math.Sin(float64(dStart)*mathlib.DegreesToRadians) * distance)
	y1 = int(math.Cos(float64(dStart)*mathlib.DegreesToRadians) * distance)

	x2 = int(math.Sin(float64(dEnd)*mathlib.DegreesToRadians) * distance)
	y2 = int(math.Cos(float64(dEnd)*mathlib.DegreesToRadians) * distance)

	return x1, y1, x2, y2
}

func newUserColor() color.Color {
	c := &color.RGBA{
		R: math.MaxUint8,
		G: math.MaxUint8,
		B: math.MaxUint8,
		A: math.MaxUint8,
	}

	componentBudget := 256

	for componentBudget > 0 {
		componentBudget--

		which := rand.Intn(3)

		switch which {
		case 0:
			if c.R > 0 {
				c.R--
			}
		case 1:
			if c.G > 0 {
				c.G--
			}
		case 2:
			if c.B > 0 {
				c.B--
			}
		}
	}

	return c
}
