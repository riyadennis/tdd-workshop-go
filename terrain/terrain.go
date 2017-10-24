package main

import (
	"fmt"
)

type Terrain struct {
	config TerrainConfig
}
type TerrainConfig struct {
	Name    string
	Country string
}

type Option interface {
	Apply(*TerrainConfig)
}

type WithSplines struct {
	config TerrainConfig
}

func WithReticulatedSplines() Option {
	return new(WithSplines)
}
func (w *WithSplines) Apply(c *TerrainConfig) {
	w.config = *c
	fmt.Println(w)
}

type WithOutSpline struct {
	config TerrainConfig
}

func WithOutReticulateSpline() Option {
	return new(WithOutSpline)
}
func (w WithOutSpline) Apply(c *TerrainConfig) {
	w.config = *c
	fmt.Println(w)
}

type WithCity struct {
	config TerrainConfig
	Cities int
}

func WithCities(n int) Option {
	return &WithCity{
		Cities:n,
	}
}
func (w *WithCity) Apply(c *TerrainConfig) {
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
	NewTerrain(WithCities(7), WithReticulatedSplines(), WithOutReticulateSpline())
}
