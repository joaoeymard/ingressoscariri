package session

import (
	"time"

	"github.com/JoaoEymard/ingressoscariri/api/utils/settings"
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris/sessions"
)

var (
	cookieName   = settings.GetSettings().CookieName
	hashKey      = []byte(settings.GetSettings().HashKey)
	blockKey     = []byte(settings.GetSettings().BlockKey)
	secureCookie = securecookie.New(hashKey, blockKey)

	session *sessions.Sessions
)

func init() {
	session = sessions.New(sessions.Config{
		Cookie:  cookieName,
		Encode:  secureCookie.Encode,
		Decode:  secureCookie.Decode,
		Expires: 5 * time.Minute,
	})
}

// GetSession Retorna a sessions
func GetSession() *sessions.Sessions {
	return session
}
