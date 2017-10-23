package main

import (
	"fmt"
	_ "github.com/lib/pq"
)

type Terrain struct {
	config Config
}
type Config struct {
	Name    string
	Country string
}

type Option interface {
	Apply(*Config)
}

type WithSplines struct {
	config Config
}

func WithReticulatedSplines() Option {
	return new(WithSplines)
}
func (w *WithSplines) Apply(c *Config) {
	w.config = *c
	fmt.Println(w)
}

type WithOutSpline struct {
	config Config
}

func WithOutReticulateSpline() Option {
	return new(WithOutSpline)
}
func (w WithOutSpline) Apply(c *Config) {
	w.config = *c
	fmt.Println(w)
}

type WithCity struct {
	config Config
}

func WithCities() Option {
	return new(WithCity)
}
func (w *WithCity) Apply(c *Config) {
	w.config = *c
	fmt.Println(w)
}

func NewTerrain(options ...Option) {
	var t Terrain
	t.config.Name = "London"
	t.config.Country = "UK"
	for _, opt := range options {
		opt.Apply(&t.config)
	}

}

func main() {
	NewTerrain(WithCities(), WithReticulatedSplines(), WithOutReticulateSpline())
}
