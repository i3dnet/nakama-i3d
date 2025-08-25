package clients

import (
	"context"
	"github.com/ascii8/nakama-go"
)

type Option func(*Client)

func WithServerKey(serverKey string) Option {
	return func(cl *Client) {
		nakama.WithServerKey(serverKey)(cl.cl)
	}
}

func WithURL(urlstr string) Option {
	return func(cl *Client) {
		nakama.WithURL(urlstr)(cl.cl)
	}
}

func WithLogf(logf func(string, ...interface{})) Option {
	return func(cl *Client) {
		cl.logf = logf
	}
}

func WithUserId(userId string) Option {
	return func(cl *Client) {
		cl.userId = userId
	}
}

func WithUsername(username string) Option {
	return func(cl *Client) {
		cl.username = username
	}
}

func WithDebug() Option {
	return func(cl *Client) {
		cl.debug = true
	}
}

func WithPersist() Option {
	return func(cl *Client) {
		cl.persist = true
	}
}

func WithHandler(handler nakama.ConnHandler) Option {
	return func(cl *Client) {
		if x, ok := handler.(interface {
			ConnectHandler(context.Context)
		}); ok {
			cl.connectHandler = x.ConnectHandler
		}
		if x, ok := handler.(interface {
			DisconnectHandler(context.Context, error)
		}); ok {
			cl.disconnectHandler = x.DisconnectHandler
		}
		if x, ok := handler.(interface {
			ErrorHandler(context.Context, *nakama.ErrorMsg)
		}); ok {
			cl.errorHandler = x.ErrorHandler
		}
		if x, ok := handler.(interface {
			ChannelMessageHandler(context.Context, *nakama.ChannelMessageMsg)
		}); ok {
			cl.channelMessageHandler = x.ChannelMessageHandler
		}
		if x, ok := handler.(interface {
			ChannelPresenceEventHandler(context.Context, *nakama.ChannelPresenceEventMsg)
		}); ok {
			cl.channelPresenceEventHandler = x.ChannelPresenceEventHandler
		}
		if x, ok := handler.(interface {
			MatchDataHandler(context.Context, *nakama.MatchDataMsg)
		}); ok {
			cl.matchDataHandler = x.MatchDataHandler
		}
		if x, ok := handler.(interface {
			MatchPresenceEventHandler(context.Context, *nakama.MatchPresenceEventMsg)
		}); ok {
			cl.matchPresenceEventHandler = x.MatchPresenceEventHandler
		}
		if x, ok := handler.(interface {
			MatchmakerMatchedHandler(context.Context, *nakama.MatchmakerMatchedMsg)
		}); ok {
			cl.matchmakerMatchedHandler = x.MatchmakerMatchedHandler
		}
		if x, ok := handler.(interface {
			NotificationsHandler(context.Context, *nakama.NotificationsMsg)
		}); ok {
			cl.notificationsHandler = x.NotificationsHandler
		}
		if x, ok := handler.(interface {
			StatusPresenceEventHandler(context.Context, *nakama.StatusPresenceEventMsg)
		}); ok {
			cl.statusPresenceEventHandler = x.StatusPresenceEventHandler
		}
		if x, ok := handler.(interface {
			StreamDataHandler(context.Context, *nakama.StreamDataMsg)
		}); ok {
			cl.streamDataHandler = x.StreamDataHandler
		}
		if x, ok := handler.(interface {
			StreamPresenceEventHandler(context.Context, *nakama.StreamPresenceEventMsg)
		}); ok {
			cl.streamPresenceEventHandler = x.StreamPresenceEventHandler
		}
		if x, ok := handler.(interface {
			StateHandler(context.Context)
		}); ok {
			cl.stateHandler = x.StateHandler
		}
	}
}
