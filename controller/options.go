package controller

import (
	"html/template"
)

type Options struct {
	GenerateCCode   bool
	CompileCCode    bool
	GenerateCtlCode bool
	ConfPath        string
	Templates       []string
	FuncMap         template.FuncMap
	OutputStrategy  strategy
	CHeaders        string
	OutputDir       string
}
type strategy string

const (
	file   strategy = "file"
	stdout strategy = "stdout"
)

type Option func(*Options)

func WithGenerateCCode() Option {
	return func(o *Options) {
		o.GenerateCCode = true
	}
}

func WithCompileCCode() Option {
	return func(o *Options) {
		o.CompileCCode = true
	}
}

func WithGenerateCtlCode() Option {
	return func(o *Options) {
		o.GenerateCtlCode = true
	}
}

func WithConfig(configPath string) Option {
	return func(o *Options) {
		o.ConfPath = configPath
	}
}

func WithTemplates(temp ...string) Option {
	return func(o *Options) {
		o.Templates = append(o.Templates, temp...)
	}
}

func WithFunc(name string, f any) Option {
	return func(o *Options) {
		o.FuncMap[name] = f
	}
}

func WithOutputFile(path string) Option {
	return func(o *Options) {
		o.OutputStrategy = file
		o.OutputDir = path
	}
}

func WithStdOut() Option {
	return func(o *Options) {
		o.OutputStrategy = stdout
	}
}

func WithCHeaders(path string) Option {
	return func(o *Options) {
		o.CHeaders = path
	}
}
