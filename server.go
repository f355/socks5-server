package main

import (
	"encoding/json"
	"github.com/armon/go-socks5"
	"log"
	"os"
)

type LoggingCredentials struct {
	UserPass map[string]string
	Logger   *log.Logger
}

func (c LoggingCredentials) Valid(user, password string) bool {
	pass, ok := c.UserPass[user]
	if !ok {
		c.Logger.Printf("[WRN] unknown user: %v", user)
		return false
	}

	if password == pass {
		c.Logger.Printf("[INF] authenticated: %v", user)
		return true
	}

	c.Logger.Printf("[WRN] wrong password for user: %v", user)
	return false
}

func main() {
	var user_pass map[string]string
	json.Unmarshal([]byte(os.Getenv("PROXY_AUTH")), &user_pass)

	logger := log.New(os.Stdout, "", log.LstdFlags)

	creds := LoggingCredentials{
		UserPass: user_pass,
		Logger:   logger,
	}
	cator := socks5.UserPassAuthenticator{Credentials: creds}

	// Create a SOCKS5 server
	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{cator},
		Logger:      logger,
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", "0.0.0.0:1080"); err != nil {
		panic(err)
	}
}
