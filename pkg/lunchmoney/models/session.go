package models

type Session struct {
	apiKey string
	userID string
}

func (s *Session) GetAPIKey() string {
	return s.apiKey
}

func (s *Session) SetAPIKey(apiKey string) {
	s.apiKey = apiKey
}

func (s *Session) GetUserID() string {
	return s.userID
}

func (s *Session) SetUserID(userID string) {
	s.userID = userID
}
