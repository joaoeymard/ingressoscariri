package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
)

type Settings struct {
	Database struct {
		ConnectionRw Connection `json:"connectionRw"`
		ConnectionRo Connection `json:"connectionRo"`
	} `json:"database"`
	Listen     string `json:"listen"`
	CookieName string `json:"cookieName"`
	HashKey    string `json:"hashKey"`
	BlockKey   string `json:"blockKey"`
}

type Connection struct {
	Host        string `json:"host"`
	User        string `json:"user"`
	Pass        string `json:"pass"`
	Name        string `json:"name"`
	MaxOpenConn int    `json:"maxOpenConn"`
	MaxIdleConn int    `json:"maxIdleConn"`
}

var (
	Regex = map[string]string{
		"integer": "[0-9]+",
		"string":  "",
	}
)

var (
	environments = map[string]string{"production": "api/utils/prod.json", "development": "api/utils/dev.json"}
	settings     Settings
	session      *sessions.Sessions
	ctx          context.Context
	env          string
)

func init() {
	env = os.Getenv("GO_UTILS")
	if env == "" {
		if GoDetails, _ := strconv.ParseBool(os.Getenv("GO_DETAILS")); GoDetails {
			fmt.Println("[Warning] Setting development environment due to lack of GO_UTILS value")
		} else {
			fmt.Println("[Warning] GO_UTILS")
		}
		env = "development"
	}
	loadSettingsByEnv(env)
	loadSession()
}

// loadSettingsByEnv Receber as configurações do json, correspondente ao env, e setar no struct
func loadSettingsByEnv(env string) {

	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		if GoDetails, _ := strconv.ParseBool(os.Getenv("GO_DETAILS")); GoDetails {
			fmt.Println("[Error] While reading config file", err)
		} else {
			fmt.Println("[Error] ReadFile environments[env]")
		}
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		if GoDetails, _ := strconv.ParseBool(os.Getenv("GO_DETAILS")); GoDetails {
			fmt.Println("[Error] While parsing config file", jsonErr)
		} else {
			fmt.Println("[Error] Unmarshal settings")
		}
	}
}

func loadSession() {
	cookieName := settings.CookieName
	hashKey := []byte(settings.HashKey)
	blockKey := []byte(settings.BlockKey)
	secureCookie := securecookie.New(hashKey, blockKey)

	session = sessions.New(sessions.Config{
		Cookie:  cookieName,
		Encode:  secureCookie.Encode,
		Decode:  secureCookie.Decode,
		Expires: 5 * time.Minute,
	})
}

// GetSettings Retorna as configurações
func GetSettings() Settings {
	return settings
}

func GetSession() *sessions.Sessions {
	return session
}

// IsTestEnvironment Teste
func IsTestEnvironment() bool {
	return env == "tests"
}
