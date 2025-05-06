package main

import (
	"context"
	"fmt"
	"github.com/ascii8/nakama-go"
	"github.com/rs/zerolog"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/nakama-test-client/internal/clients"
	"log"
	"os"
	"time"
)

type GameTest struct {
	cl     *clients.Client
	logger zerolog.Logger
	ctx    context.Context
}

func NewGameTest(logger zerolog.Logger) *GameTest {
	ctx := context.Background()

	g := &GameTest{ctx: ctx, logger: logger}

	url := os.Getenv("NAKAMA_URL")
	if url == "" {
		url = "http://127.0.0.1:7350"
	}

	g.cl = clients.NewClient(
		clients.WithDebug(),
		clients.WithPersist(),
		clients.WithLogf(func(s string, v ...interface{}) {
			g.logger.Debug().CallerSkipFrame(1).Msgf(s, v...)
		}),
		clients.WithHandler(g),
		clients.WithURL(url),
	)

	return g
}

func (g *GameTest) DisconnectHandler(ctx context.Context, err error) {
	go func() {
		<-time.After(1 * time.Second)
		for i := 0; !g.cl.Connected(); i++ {
			//g.connectedLabel = strings.Repeat(".", 1+i%5)
			<-time.After(1 * time.Second)
			log.Print("disconnecting...")
			//g.window.Invalidate()
		}
	}()
}

func (g *GameTest) StateHandler(ctx context.Context) {
	log.Println("Handle state...")
}

func (g *GameTest) ConnectHandler(ctx context.Context) {
	log.Println("ConnectHandler...")

	//time.Sleep(1 * time.Second)
	//log.Println("creating match...")
	//g.cl(g.ctx, func(err error) {

	//time.Sleep(1 * time.Second)
	// log.Println("Creating match...")
	// g.cl.CreateMatch(g.ctx)

	time.Sleep(1 * time.Second)
	log.Println("Add match...")
	g.cl.AddToMatch(g.ctx)

	time.Sleep(1 * time.Second)
	log.Println("joining match...")

	err := g.cl.Join(g.ctx)
	if err != nil {
		log.Println(err)
	}
}

func (g *GameTest) Shutdown() {
	log.Println("ConnectHandler...")
	err := g.cl.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (g *GameTest) Run() error {
	if err := g.cl.Open(g.ctx); err != nil {
		return err
	}
	for {
		//log.Println("Running...")
		time.Sleep(1 * time.Second)
	}
}

func (g *GameTest) notificationsHandler(ctx context.Context, msg *nakama.MatchmakerMatchedMsg) {
	log.Println("notificationsHandler...")
	log.Println(fmt.Sprintf("Matched: %v", msg))
}

func main() {

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	w := zerolog.NewConsoleWriter(func(cw *zerolog.ConsoleWriter) {
		cw.Out = os.Stdout
		cw.TimeFormat = "2006-01-02 15:04:05"
		cw.PartsOrder = []string{zerolog.TimestampFieldName, zerolog.LevelFieldName, zerolog.CallerFieldName, zerolog.MessageFieldName}
		cw.FieldsExclude = cw.PartsOrder
	})

	// creating logger
	logger := zerolog.New(w).With().Timestamp().Logger()

	test := NewGameTest(logger)

	err := test.Run()

	if err != nil {
		log.Fatal(err)
	}
}
