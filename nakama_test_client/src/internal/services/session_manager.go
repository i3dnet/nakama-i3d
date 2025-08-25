package services

import pb "gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/nakama-test-client/pkg/grpc_client"

type session struct {
	displayName  string
	sessionToken string
}

type SessionManager struct {
	session *session
}

func NewSessionManager(displayName string, sessionToken string) *SessionManager {
	return &SessionManager{
		session: &session{
			displayName:  displayName,
			sessionToken: sessionToken,
		},
	}
}

func (m *SessionManager) SetSessionToken(sessionToken string) {
	m.session.sessionToken = sessionToken
}

func (m *SessionManager) GetSessionToken() string {
	return m.session.sessionToken
}

func (m *SessionManager) SetDisplayName(displayName string) {
	m.session.displayName = displayName
}

func (m *SessionManager) GetDisplayName() string {
	return m.session.displayName
}

func (m *SessionManager) GetPlayer() *pb.Player {
	return &pb.Player{
		DisplayName: m.session.displayName,
		SessionId:   m.session.sessionToken,
	}
}
