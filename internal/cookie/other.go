package cookie

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gorilla/securecookie"
)

// Cookies is the securecookie instance that should be called by everyone
var Cookies *securecookie.SecureCookie

// GetSession reads and returns the data from the session cookie
func GetSession(r *http.Request) *Session {
	authToken := GetAuthToken(r)
	if authToken != "" { // use Auth tokens by default
		var ret Session
		Cookies.Decode("kn-sessionid", authToken, &ret)
		return &ret
	}
	cookie, err := r.Cookie("kn-sessionid")
	if err != nil {
		return nil
	}
	if cookie.Value == "" {
		return nil
	}
	var ret Session
	Cookies.Decode(cookie.Name, cookie.Value, &ret)
	return &ret
}

// SetSession sets the data to the session cookie
func SetSession(w http.ResponseWriter, sess Session) (string, error) {
	encoded, err := Cookies.Encode("kn-sessionid", sess)
	if err != nil {
		return "", err
	}
	cookie := &http.Cookie{
		Name:     "kn-sessionid",
		Value:    encoded,
		Path:     "/",
		HttpOnly: false,
		SameSite: http.SameSiteDefaultMode,
	}
	http.SetCookie(w, cookie)
	return encoded, nil
}

// GetAuthToken returns the authentication token from a request
func GetAuthToken(r *http.Request) string {
	header := r.Header.Get("Authorization")
	if strings.HasPrefix(header, "Bearer ") {
		return strings.TrimPrefix(header, "Bearer ")
	}
	return ""
}

// RemoveSessionCookie clears the session cookie, effectively revoking it. When setting MaxAge to 0, the browser will also clear it out
func RemoveSessionCookie(w http.ResponseWriter) {
	emptyCookie := &http.Cookie{
		Name:    "kn-sessionid",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}
	http.SetCookie(w, emptyCookie)
}

// Session represents the data storred in a session cookie
type Session struct {
	UserID int64
}

// Initialize must be called after setting the DataDir, but before everything else
func Initialize(dataDir string) {
	if err := os.MkdirAll(dataDir, 0775); err != nil {
		panic(err)
	}

	// generate secure cookie
	var secure []byte
	if _, err := os.Stat(path.Join(dataDir, "secretKey")); os.IsNotExist(err) {
		secure = securecookie.GenerateRandomKey(32)
		ioutil.WriteFile(path.Join(dataDir, "secretKey"), secure, 0777)
	} else {
		secure, err = ioutil.ReadFile(path.Join(dataDir, "secretKey"))
		if err != nil {
			panic(fmt.Sprintln("Could not read the secret key:", err))
		}
		secure = bytes.TrimSpace(secure)
	}
	if len(secure) != 32 {
		panic("Invalid secure key length")
	}
	Cookies = securecookie.New(secure, nil).SetSerializer(securecookie.JSONEncoder{})
}
