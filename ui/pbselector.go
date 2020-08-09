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

import "github.com/StarAurryon/qt/core"
import "github.com/StarAurryon/qt/widgets"

import "github.com/StarAurryon/lpedit/qtctrl"

type PBSelector struct {
    *PBSelectorUI
    ctrl *qtctrl.Controller
    progress *widgets.QProgressDialog
    parent *LPEdit
}

func NewPBSelector(c *qtctrl.Controller, p widgets.QWidget_ITF, parent *LPEdit) *PBSelector {
    pb := &PBSelector{PBSelectorUI: NewPBSelectorUI(p), ctrl: c, parent: parent}
    pb.init()
    pb.update()
    return pb
}

func (pb *PBSelector) init() {
    pb.ctrl.ConnectLoop(pb.loop)
    pb.ctrl.ConnectLoopError(pb.loopError)
    pb.ctrl.ConnectPresetLoadProgress(pb.presetLoadProgress)
    pb.ctrl.ConnectSetLoadProgress(pb.setLoadProgress)
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
    pb.parent.disconnectSignal()
}

func(pb *PBSelector) loop() {
    pb.updateButtons()
    if pb.ctrl.IsStarted() {
        pb.ctrl.QueryAllSets()
    }
}

func (pb *PBSelector) loopError(err string) {
    pb.updateButtons()
    mb := widgets.NewQMessageBox(pb)
    mb.Critical(pb, "An error occured", err, widgets.QMessageBox__Ok, 0)
    pb.parent.disconnectSignal()
}

func (pb *PBSelector) setLoadProgress(progress int) {
    if progress != 100 {
        if pb.progress == nil {
            pb.progress = widgets.NewQProgressDialog2("Progress", "", 0, 100, pb, 0)
            pb.progress.SetWindowTitle("Progress")
        }
        pb.progress.SetValue(progress)
    } else {
        if pb.progress != nil {
            pb.progress.Close()
            pb.progress.DestroyQObject()
            pb.progress = nil
        }
        pb.ctrl.QueryAllPresets()
    }
}

func (pb *PBSelector) presetLoadProgress(progress int) {
    if progress != 100 {
        if pb.progress == nil {
            pb.progress = widgets.NewQProgressDialog2("Progress", "", 0, 100, pb, 0)
            pb.progress.SetWindowTitle("Progress")
        }
        pb.progress.SetValue(progress)
    } else {
        if pb.progress != nil {
            pb.progress.Close()
            pb.progress.DestroyQObject()
            pb.progress = nil
            pb.parent.connectSignal()
        }
    }
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
