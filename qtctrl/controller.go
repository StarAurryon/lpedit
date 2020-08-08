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

import "github.com/StarAurryon/qt/core"

import "github.com/StarAurryon/lpedit/controller"
import "github.com/StarAurryon/lpedit/pedal"

type Controller struct {
    core.QObject
    controller.Controller
    _ func() `constructor:"init"`
    _ func(pedal.PedalBoardItem) `signal:ActiveChange`
    _ func() `signal:"Loop"`
    _ func(string) `signal:"LoopError"`
    _ func(pedal.Parameter) `signal:ParameterChange`
    _ func(*pedal.PedalBoard) `signal:PresetLoad`
    _ func(*pedal.PedalBoard) `signal:SetChange`
    _ func(*pedal.PedalBoard) `signal:TempoChange`
    _ func(pedal.PedalBoardItem) `signal:TypeChange`
}

func (c *Controller) init() {
    c.Controller = *controller.NewController()
    c.SetNotify(c.notif)
}

func (c *Controller) notif(err error, n pedal.ChangeType, obj interface{}) {
    switch n {
    case pedal.ActiveChange:
        c.ActiveChange(obj.(pedal.PedalBoardItem))
    case pedal.NormalStart:
        c.Loop()
    case pedal.NormalStop:
        c.Loop()
    case pedal.ErrorStop:
        c.LoopError(err.Error())
    case pedal.ParameterChange:
        c.ParameterChange(obj.(pedal.Parameter))
    case pedal.PresetLoad:
        c.PresetLoad(obj.(*pedal.PedalBoard))
    case pedal.SetChange:
        c.SetChange(obj.(*pedal.PedalBoard))
    case pedal.TypeChange:
        c.TypeChange(obj.(pedal.PedalBoardItem))
    case pedal.TempoChange:
        c.TempoChange(obj.(*pedal.PedalBoard))
    }
}
