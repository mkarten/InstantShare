package utils

import (
	"InstantShare/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Manager struct {
	Sessions map[string]*session
}

type session struct {
	token string
	user  *models.User
}

var SessionManager Manager

func generateToken() string {
	return uuid.New().String()
}

func init() {
	SessionManager.Sessions = make(map[string]*session)
}

func (m *Manager) CheckSession(c *gin.Context) bool {
	// get the session token
	token, err := c.Cookie("session_token")
	if err != nil {
		return false
	}
	// check if the session exists
	_, ok := m.Sessions[token]
	return ok
}

func (m *Manager) CreateSession(c *gin.Context, user *models.User) {
	// create a new session token
	token := generateToken()
	// create a new session
	newSession := &session{
		token: token,
		user:  user,
	}
	// add the session to the manager
	m.Sessions[token] = newSession
	// set the cookie
	c.SetCookie("session_token", token, 3600, "/", c.Request.Host, false, false)
}

func (m *Manager) GetSessionUser(c *gin.Context) *models.User {
	// get the session token
	token, err := c.Cookie("session_token")
	if err != nil {
		return nil
	}
	// get the session
	ses, ok := m.Sessions[token]
	if !ok {
		return nil
	}
	return ses.user
}

func (m *Manager) DeleteSession(c *gin.Context) {
	// get the session token
	token, err := c.Cookie("session_token")
	if err != nil {
		return
	}
	// delete the session
	delete(m.Sessions, token)
	// delete the cookie
	c.SetCookie("session_token", "", -1, "/", c.Request.Host, false, false)
}
