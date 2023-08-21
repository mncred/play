// package config provides static shared configuration for the whole app (front and back)
package config

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var config []byte

type Provider struct {
	config Config
}

func (p *Provider) LoadDefault() error {
	if err := yaml.Unmarshal(config, &p.config); err != nil {
		return err
	}
	return nil
}

func (p *Provider) Get() Config {
	return p.config
}

// Config provides static app configuration
type Config struct {
	Id     string `json:"id" yaml:"id"`
	Debug  Debug  `json:"debug" yaml:"debug"`
	Window Window `json:"window" yaml:"window"`
}

type Window struct {
	Header     WindowHeader     `json:"header" yaml:"header"`
	Width      WindowSize       `json:"width" yaml:"width"`
	Height     WindowSize       `json:"height" yaml:"height"`
	Resizable  bool             `json:"resizable" yaml:"resizable"`
	Appearance WindowAppearance `json:"appearance" yaml:"appearance"`
}

type WindowHeader struct {
	Native     bool                   `json:"native" yaml:"native"`
	Title      string                 `json:"title" yaml:"title"`
	Maximise   bool                   `json:"maximise" yaml:"maximise"`
	Background WindowHeaderBackground `json:"background" yaml:"background"`
	Text       WindowHeaderText       `json:"text" yaml:"text"`
}

type WindowSize struct {
	Default int `json:"default" yaml:"default"`
	Min     int `json:"min" yaml:"min"`
	Max     int `json:"max" yaml:"max"`
}

type WindowAppearance struct {
	Roundness  float64                    `json:"roundness" yaml:"roundness"`
	Background WindowAppearanceBackground `json:"background" yaml:"background"`
}

type WindowAppearanceBackground struct {
	Color string `json:"color" yaml:"color"`
}

type WindowHeaderBackground struct {
	Color string `json:"color" yaml:"color"`
}
type WindowHeaderText struct {
	Color string `json:"color" yaml:"color"`
}

type Debug struct {
	Inspector bool `json:"inspector" yaml:"inspector"`
	Log       Log  `json:"log" yaml:"log"`
}

type Log struct {
	Level  string `json:"level" yaml:"level"`
	Output string `json:"output" yaml:"output"`
}
