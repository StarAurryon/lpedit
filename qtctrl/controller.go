/*
 * Copyright (C) 2020 Nicolas SCHWARTZ
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301, USA
 */

package qtctrl

import "github.com/therecipe/qt/core"

import "github.com/StarAurryon/lpedit-lib/controller"
import "github.com/StarAurryon/lpedit-lib/model/pod"

var sg controller.Signal

type Controller struct {
    core.QObject
    controller.Controller
    _ func() `constructor:"init"`
    _ func(pod.PedalBoardItem) `signal:ActiveChange`
    _ func() `signal:"Loop"`
    _ func(string) `signal:"LoopError"`
    _ func() `signal:InitDone`
    _ func(pod.Parameter) `signal:ParameterChange`
    _ func(*pod.Preset) `signal:PresetChange`
    _ func(*pod.Preset) `signal:PresetLoad`
    _ func(int) `signal:Progress`
    _ func(*pod.Set) `signal:SetChange`
    _ func(pod.PedalBoardItem) `signal:TypeChange`
}

func (c *Controller) init() {
    c.Controller = *controller.NewController()
    c.SetNotify(c.notif)
}

func (c *Controller) notif(err error, n int, obj interface{}) {
    switch n {
    case sg.StatusActiveChange():
        c.ActiveChange(obj.(pod.PedalBoardItem))
    case sg.StatusNormalStart():
        c.Loop()
    case sg.StatusNormalStop():
        c.Loop()
    case sg.StatusErrorStop():
        c.LoopError(err.Error())
    case sg.StatusInitDone():
        c.InitDone()
    case sg.StatusParameterChange():
        c.ParameterChange(obj.(pod.Parameter))
    case sg.StatusPresetChange():
        c.PresetChange(obj.(*pod.Preset))
    case sg.StatusPresetLoad():
        c.PresetLoad(obj.(*pod.Preset))
    case sg.StatusProgress():
        c.Progress(obj.(int))
    case sg.StatusSetChange():
        c.SetChange(obj.(*pod.Set))
    case sg.StatusTypeChange():
        c.TypeChange(obj.(pod.PedalBoardItem))
    }
}
