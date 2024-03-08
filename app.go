package main

import (
	"context"

	"github.com/StarAurryon/lpedit-lib/controller"
	"github.com/StarAurryon/lpedit-lib/model/pod"
	"github.com/StarAurryon/lpedit-lib/status"
	"github.com/StarAurryon/lpedit/model"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
	controller.Controller
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		Controller: controller.New(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.SetNotify(a.notif)
}

func (a *App) notif(n status.St, obj interface{}, err error) {
	switch n {
	case status.ActiveChange:
		runtime.EventsEmit(a.ctx, "activeChange")
	case status.NormalStart:
		runtime.EventsEmit(a.ctx, "start")
	case status.NormalStop:
		runtime.EventsEmit(a.ctx, "stop")
	case status.ErrorStop:
		runtime.EventsEmit(a.ctx, "stop", err.Error())
	case status.InitDone:
		runtime.EventsEmit(a.ctx, "initDone")
	case status.ParameterChange:
		runtime.EventsEmit(a.ctx, "parameterChange", func() interface{} {
			obj := obj.(pod.Parameter)
			obj.LockData()
			defer obj.UnlockData()

			return model.ToParameter(obj)
		}())
	case status.PresetLoad:
		fallthrough
	case status.PresetChange:
		runtime.EventsEmit(a.ctx, "presetChange", func() interface{} {
			obj := obj.(*pod.Preset)
			obj.LockData()
			defer obj.UnlockData()

			return model.ToPreset(obj)
		}())
	case status.Progress:
		runtime.EventsEmit(a.ctx, "statusProgress", obj.(int))
	case status.SetLoad:
		fallthrough
	case status.SetChange:
		runtime.EventsEmit(a.ctx, "setChange", func() interface{} {
			obj := obj.(*pod.Set)
			obj.LockData()
			defer obj.UnlockData()

			return model.ToSet(obj)
		}())
		// case status.TypeChange:
		// 	runtime.EventsEmit(a.ctx, "typeChange", obj.(pod.PedalBoardItem))
	}
}
