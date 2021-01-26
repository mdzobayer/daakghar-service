package config

import (
	"github.com/spf13/viper"
)

// Basic holds application basic configuration
type Basic struct {
	Version string
	Name    string
	Port    string
}

// Load loads basic configuration
func (b *Basic) Load() {
	b.Version = viper.GetString("basic.version")
	b.Name = viper.GetString("basic.name")
	b.Port = viper.GetString("basic.port")
}

// Mgo holds mongodb configuration
type Mgo struct {
	URI string
}

// Load loads mongo configuration
func (m *Mgo) Load() {
	m.URI = viper.GetString("mgo.uri")
}

// Token retunrs token configuration
type Token struct {
	Key string
}

// Load loads token configuration
func (t *Token) Load() {
	t.Key = viper.GetString("token.key")
}

// App provides app config
type App struct {
	basic Basic
	mgo   Mgo
	token Token
}

// Load loads application configuration
func (a *App) Load() {
	a.basic.Load()
	a.mgo.Load()
	a.token.Load()
}

// Basic returns Basic configuration
func (a App) Basic() Basic {
	return a.basic
}

// Token returns Token configuration
func (a App) Token() Token {
	return a.token
}

// Mgo returns mongodb configuration
func (a App) Mgo() Mgo {
	return a.mgo
}

var app App

// Get retuns Application configuration
func Get() *App {
	return &app
}
