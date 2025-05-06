package clients

import (
	"context"
	"fmt"
	"github.com/ascii8/nakama-go"
	"github.com/google/uuid"
	"github.com/rs/xid"
	"net/http"
	"sync"
)

type MatchState struct{}
type Client struct {
	cl         *nakama.Client
	conn       *nakama.Conn
	debug      bool
	userId     string
	username   string
	createUser bool
	logf       func(string, ...interface{})
	persist    bool

	ticketId string
	matchId  string
	state    *MatchState
	waiting  bool

	rw sync.RWMutex

	connectHandler              func(context.Context)
	disconnectHandler           func(context.Context, error)
	errorHandler                func(context.Context, *nakama.ErrorMsg)
	channelMessageHandler       func(context.Context, *nakama.ChannelMessageMsg)
	channelPresenceEventHandler func(context.Context, *nakama.ChannelPresenceEventMsg)
	matchDataHandler            func(context.Context, *nakama.MatchDataMsg)
	matchPresenceEventHandler   func(context.Context, *nakama.MatchPresenceEventMsg)
	matchmakerMatchedHandler    func(context.Context, *nakama.MatchmakerMatchedMsg)
	notificationsHandler        func(context.Context, *nakama.NotificationsMsg)
	statusPresenceEventHandler  func(context.Context, *nakama.StatusPresenceEventMsg)
	streamDataHandler           func(context.Context, *nakama.StreamDataMsg)
	streamPresenceEventHandler  func(context.Context, *nakama.StreamPresenceEventMsg)
	stateHandler                func(context.Context)
}

func NewClient(opts ...Option) *Client {
	cl := &Client{
		logf:    func(string, ...interface{}) {},
		waiting: true,
	}
	cl.cl = nakama.New(
		nakama.WithURL("http://127.0.0.1:7350"),
		nakama.WithServerKey("defaultkey"),
		nakama.WithAuthHandler(cl),
	)
	for _, o := range opts {
		o(cl)
	}
	if cl.debug {
		nakama.WithTransport(&http.Transport{
			DisableCompression: true,
		})(cl.cl)
		nakama.WithLogger(cl.logf)(cl.cl)
	}
	if cl.userId == "" {

		cl.userId = uuid.New().String()
		cl.logf("userID: %s", cl.userId)
	}
	if cl.username == "" {
		cl.username = xid.New().String()
		cl.logf("userName: %s", cl.userId)
	}

	cl.createUser = true
	return cl
}

func (cl *Client) SetUser(userId, username string, createUser bool) {
	cl.userId, cl.username = userId, username
	cl.createUser = createUser
}

func Dial(ctx context.Context, opts ...Option) (*Client, error) {
	cl := NewClient(opts...)
	if err := cl.Open(ctx); err != nil {
		return nil, err
	}
	return cl, nil
}

func (cl *Client) Open(ctx context.Context) error {
	conn := cl.conn
	if conn == nil {
		opts := []nakama.ConnOption{
			nakama.WithConnHandler(cl),
			nakama.WithConnPersist(cl.persist),
		}
		if cl.debug {
			opts = append(opts, nakama.WithConnFormat("json"))
		}
		var err error
		if conn, err = cl.cl.NewConn(ctx, opts...); err != nil {
			return err
		}
		cl.conn = conn
	}
	return nil
}

func (cl *Client) Close() error {
	_ = cl.Leave(context.Background())
	cl.rw.Lock()
	defer cl.rw.Unlock()
	if cl.conn != nil {
		_ = cl.conn.CloseWithStopErr(true, true, nil)
	}
	cl.state = nil
	return nil
}

func (cl *Client) Connected() bool {
	conn := cl.conn
	return conn != nil && conn.Connected()
}

func (cl *Client) State() *MatchState {
	return cl.state
}

func (cl *Client) MatchId() string {
	return cl.matchId
}

func (cl *Client) Ready(ctx context.Context) bool {
	return true
	//ch := make(chan bool, 1)
	//go func() {
	//	defer close(ch)
	//	for {
	//		if state := cl.state; state != nil && state.State.RematchCountdown == 0 {
	//			ch <- true
	//			return
	//		}
	//		select {
	//		case <-ctx.Done():
	//			return
	//		case <-time.After(50 * time.Millisecond):
	//		}
	//	}
	//}()
	//select {
	//case <-ctx.Done():
	//	return false
	//case res := <-ch:
	//	return res
	//}
}

func (cl *Client) Next(ctx context.Context) bool {
	return true
	//ch := make(chan bool, 1)
	//go func() {
	//	defer close(ch)
	//	for {
	//		cl.rw.RLock()
	//		waiting, state := cl.waiting, cl.state
	//		cl.rw.RUnlock()
	//		switch {
	//		case waiting || state == nil:
	//		case state.State.Winner != 0,
	//			state.State.Draw,
	//			state.State.RematchCountdown != 0:
	//			return
	//		case state.YourTurn:
	//			ch <- true
	//			return
	//		}
	//		select {
	//		case <-ctx.Done():
	//			return
	//		case <-time.After(50 * time.Millisecond):
	//		}
	//	}
	//}()
	//select {
	//case <-ctx.Done():
	//	return false
	//case res := <-ch:
	//	return res
	//}
}

func (cl *Client) AuthHandler(ctx context.Context, nakamaClient *nakama.Client) error {
	return nakamaClient.AuthenticateDevice(ctx, cl.userId, cl.createUser, cl.username)
}

func (cl *Client) ConnectHandler(ctx context.Context) {
	cl.logf("Connect!")
	if cl.connectHandler != nil {
		cl.connectHandler(ctx)
	}
}

func (cl *Client) DisconnectHandler(ctx context.Context, err error) {
	cl.logf("Disconnect: %v", err)
	if cl.disconnectHandler != nil {
		cl.disconnectHandler(ctx, err)
	}
}

func (cl *Client) ErrorHandler(ctx context.Context, msg *nakama.ErrorMsg) {
	cl.logf("ErrorHandler: %+v", msg)
	if cl.errorHandler != nil {
		cl.errorHandler(ctx, msg)
	}
}

func (cl *Client) ChannelMessageHandler(ctx context.Context, msg *nakama.ChannelMessageMsg) {
	cl.logf("ChannelMessage: %+v", msg)
	if cl.channelMessageHandler != nil {
		cl.channelMessageHandler(ctx, msg)
	}
}

func (cl *Client) ChannelPresenceEventHandler(ctx context.Context, msg *nakama.ChannelPresenceEventMsg) {
	cl.logf("ChannelPresenceEvent: %+v", msg)
	if cl.channelPresenceEventHandler != nil {
		cl.channelPresenceEventHandler(ctx, msg)
	}
}

func (cl *Client) MatchDataHandler(ctx context.Context, msg *nakama.MatchDataMsg) {
	cl.logf("MatchData: %+v", msg)
	//state := new(MatchState)
	//if err := state.Unmarshal(msg.Data); err != nil {
	//	cl.logf("unable to unmarshal MatchData: %v", err)
	//	state = nil
	//}
	//cl.rw.Lock()
	//defer cl.rw.Unlock()
	//prev := cl.state
	//cl.waiting, cl.state = state == nil, state
	//if cl.matchDataHandler != nil {
	//	cl.matchDataHandler(ctx, msg)
	//}
	//if cl.stateHandler == nil {
	//	return
	//}
	//switch {
	//case prev == nil && state != nil,
	//	prev != nil && state == nil,
	//	prev.YourTurn != state.YourTurn,
	//	prev.State.RematchCountdown != state.State.RematchCountdown,
	//	state.State.Winner != 0,
	//	state.State.Draw:
	//	cl.stateHandler(ctx)
	//}
}

func (cl *Client) MatchPresenceEventHandler(ctx context.Context, msg *nakama.MatchPresenceEventMsg) {
	cl.logf("MatchPresenceEvent: %+v", msg)
	if len(msg.Leaves) != 0 {
		cl.rw.Lock()
		cl.state = nil
		cl.rw.Unlock()
		if cl.stateHandler != nil {
			cl.stateHandler(ctx)
		}
	}
	if cl.matchPresenceEventHandler != nil {
		cl.matchPresenceEventHandler(ctx, msg)
	}
}

type UpdateInstanceInfoRequest struct {
	Id          string         `json:"id"`
	PlayerCount int            `json:"player_count"`
	Metadata    map[string]any `json:"metadata"`
}

func (cl *Client) CreateMatch(ctx context.Context) {
	cl.conn.MatchCreateAsync(ctx, "test", func(matchMsg *nakama.MatchMsg, err error) {
		if err != nil {
			cl.logf("CreateMatch err: %+v", err)
			return
		}

		cl.logf("CreateMatch: %+v", matchMsg)
		cl.matchId = matchMsg.MatchId
	})
}

func (cl *Client) AddToMatch(ctx context.Context) {
	msg := nakama.MatchmakerAddMsg{
		MinCount: 2,
		MaxCount: 2,
		Query:    "*",
	}

	cl.conn.MatchmakerAddAsync(ctx, &msg, func(ticket *nakama.MatchmakerTicketMsg, err error) {
		if err != nil {
			cl.logf("CreateMatch err: %+v", err)
			return
		}

		cl.logf("Matchmaker Ticket Msg %+v", ticket)
	})
}

func (cl *Client) MatchmakerMatchedHandler(ctx context.Context, msg *nakama.MatchmakerMatchedMsg) {
	cl.logf("MatchmakerMatched: %+v", msg)
}

func (cl *Client) NotificationsHandler(ctx context.Context, msg *nakama.NotificationsMsg) {
	cl.logf("Notifications: %+v", msg)
	if cl.notificationsHandler != nil {
		cl.notificationsHandler(ctx, msg)
	}
}

func (cl *Client) StatusPresenceEventHandler(ctx context.Context, msg *nakama.StatusPresenceEventMsg) {
	cl.logf("StatusPresenceEvent: %+v", msg)
	if cl.statusPresenceEventHandler != nil {
		cl.statusPresenceEventHandler(ctx, msg)
	}
}

func (cl *Client) StreamDataHandler(ctx context.Context, msg *nakama.StreamDataMsg) {
	cl.logf("StreamData: %+v", msg)
	if cl.streamDataHandler != nil {
		cl.streamDataHandler(ctx, msg)
	}
}

func (cl *Client) StreamPresenceEventHandler(ctx context.Context, msg *nakama.StreamPresenceEventMsg) {
	cl.logf("StreamPresenceEvent: %+v", msg)
	if cl.streamPresenceEventHandler != nil {
		cl.streamPresenceEventHandler(ctx, msg)
	}
}

func (cl *Client) Join(ctx context.Context) error {
	cl.logf("Join: joining match")
	if cl.ticketId != "" {
		return fmt.Errorf("waiting matchmaker %s", cl.ticketId)
	}
	cl.conn.MatchmakerAddAsync(ctx, nakama.MatchmakerAdd("*", 2, 2), func(msg *nakama.MatchmakerTicketMsg, err error) {
		switch {
		case err != nil:
			cl.logf("Join: unable to join match: %v", err)
		default:
			cl.rw.Lock()
			defer cl.rw.Unlock()
			ticketId := msg.GetTicket()

			cl.logf("Join: added matchmaker ticket %q", ticketId)
			cl.ticketId = ticketId
		}
	})
	return nil
}

func (cl *Client) JoinAsync(ctx context.Context, f func(error)) {
	go func() {
		if err := cl.Join(ctx); f != nil {
			f(err)
		}
	}()
}

func (cl *Client) Leave(ctx context.Context) error {
	cl.logf("Leave: leaving match")
	cl.rw.Lock()
	defer cl.rw.Unlock()
	if cl.ticketId != "" {
		cl.conn.MatchmakerRemoveAsync(ctx, cl.ticketId, nil)
	}
	if cl.matchId != "" {
		cl.conn.MatchLeaveAsync(ctx, cl.matchId, nil)
	}
	cl.ticketId, cl.matchId, cl.waiting, cl.state = "", "", true, nil
	return nil
}

func (cl *Client) LeaveAsync(ctx context.Context, f func(error)) {
	go func() {
		if err := cl.Leave(ctx); f != nil {
			f(err)
		}
	}()
}

func (cl *Client) Move(ctx context.Context, row, col int) error {
	cl.logf("Move: moving %d, %d", row, col)
	//cl.rw.RLock()
	//matchId, state := cl.matchId, cl.state
	//cl.rw.RUnlock()
	//if matchId == "" || state == nil {
	//	return fmt.Errorf("no active match")
	//}
	//data, err := NewMove(row, col).Marshal()
	//if err != nil {
	//	return fmt.Errorf("unable to marshal move: %w", err)
	//}
	//cl.rw.Lock()
	//defer cl.rw.Unlock()
	//cl.waiting = true
	//return cl.conn.MatchDataSend(ctx, matchId, OpCodeMove, data, true, nil)

	return nil
}

func (cl *Client) MoveAsync(ctx context.Context, row, col int, f func(error)) {
	go func() {
		if err := cl.Move(ctx, row, col); f != nil {
			f(err)
		}
	}()
}
