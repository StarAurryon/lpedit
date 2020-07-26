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

package ui

import "github.com/therecipe/qt/core"
import "github.com/therecipe/qt/widgets"

import "lpedit/qtctrl"

type PBSelector struct {
    *PBSelectorUI
    ctrl *qtctrl.Controller
}

func NewPBSelector(c *qtctrl.Controller, p widgets.QWidget_ITF) *PBSelector {
    pb := &PBSelector{PBSelectorUI: NewPBSelectorUI(p), ctrl: c}
    pb.init()
    pb.update()
    return pb
}

func (pb *PBSelector) init() {
    pb.ctrl.ConnectLoop(pb.updateButtons)
    pb.ctrl.ConnectLoopError(pb.loopError)
    pb.CloseButton.ConnectClicked(pb.clickClose)
    pb.StartButton.ConnectClicked(pb.clickStart)
    pb.StopButton.ConnectClicked(pb.clickStop)
}

func (pb *PBSelector) clickClose(bool) {
    pb.Close()
}

func (pb *PBSelector) clickStart(bool) {
    dev := pb.ListDev.CurrentText()
    pb.StartButton.SetEnabled(false)
    pb.ctrl.Start(dev)
}

func (pb *PBSelector) clickStop(bool) {
    pb.ctrl.Stop()
}

func (pb *PBSelector) loopError(err string) {
    pb.updateButtons()
    mb := widgets.NewQMessageBox(pb)
    mb.Critical(pb, "An error occured", err, widgets.QMessageBox__Ok, 0)
}

func (pb *PBSelector) update() {
    pb.updateButtons()

    //Populate dev list
    pb.ListDev.Clear()
    devs := pb.ctrl.ListDevices()
    for _, dev := range devs {
        pb.ListDev.AddItem(dev[0], core.NewQVariant15(dev[1]))
    }
}

func(pb *PBSelector) updateButtons() {
    started := pb.ctrl.IsStarted()
    pb.StartButton.SetEnabled(!started)
    pb.StopButton.SetEnabled(started)
    pb.ListDev.SetEnabled(!started)
}

func (pb *PBSelector) notify() {
    core.NewQEvent(core.QEvent__UpdateRequest)
}
