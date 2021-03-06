package main

import (
	"github.com/fananchong/go-x/common"
)

var (
	xargs *Args          = &Args{}
	xlog  common.ILogger = common.NewGLogger()
)

type App struct {
	common.App
}

func NewApp() *App {
	this := &App{}
	this.Type = int(common.Room)
	this.Args = xargs
	this.Logger = xlog
	this.Derived = this
	return this
}

func (this *App) OnAppReady() {
}

func (this *App) OnAppShutDown() {
}
