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
import "github.com/therecipe/qt/gui"
import "github.com/therecipe/qt/widgets"

import "fmt"

import "lpedit/pedal"
import "lpedit/qtctrl"

type LPEdit struct {
    *LPEditUI
    about      *AboutUI
    pbSelector *PBSelector
    ctrl       *qtctrl.Controller
    pedals     []*Pedal
}

func NewLPEdit(c *qtctrl.Controller, p widgets.QWidget_ITF) *LPEdit {
    l := &LPEdit{LPEditUI: NewLPEditUI(p), ctrl: c}
    l.init()
    return l
}

func (l *LPEdit) init() {
    //UI Connections
    l.ConnectCloseEvent(l.windowClose)
    l.ActionAbout.ConnectTriggered(l.aboutClick)
    l.ActionSelect_Device.ConnectTriggered(l.pbSelectorClick)
    l.ActionQuit.ConnectTriggered(func(bool) {l.Close()})

    //PedalBoard Connections
    l.ctrl.ConnectParameterChange(l.updateParameter)
    l.ctrl.ConnectPedalBoardChange(l.updatePedalBoard)
    l.ctrl.ConnectTempoChange(l.updateTempo)
    l.ctrl.ConnectTypeChange(l.updateType)
    l.initPedals()
}

func (l *LPEdit) initPedals() {
    pedalType := l.ctrl.GetPedalType()

    for i := 0; i < 9; i++ {
        line := widgets.NewQFrame(l.ScrollPedalW, core.Qt__Widget)
        line.SetFrameShape(widgets.QFrame__HLine)
        line.SetFrameShadow(widgets.QFrame__Sunken)
        l.PedalsLayout.AddWidget(line, 0, 0)
        if i != 8 {
            p := NewPedal(l.ScrollPedalW, l.ctrl, pedalType, i)
            l.PedalsLayout.AddWidget(p, 0, 0)
            l.pedals = append(l.pedals, p)
        }
    }
}

func (l *LPEdit) aboutClick(vbo bool) {
    if l.about == nil {
        l.about = NewAboutUI(l)
        l.about.ConnectCloseEvent(func(event *gui.QCloseEvent) {
            l.about.DeleteLater()
            l.about = nil
        })
    }
    l.about.Show()
    l.about.Raise()
}

func (l *LPEdit) pbSelectorClick(vbo bool) {
    if l.pbSelector == nil {
        l.pbSelector = NewPBSelector(l.ctrl, l)
        l.pbSelector.ConnectCloseEvent(func(event *gui.QCloseEvent) {
            l.pbSelector.DeleteLater()
            l.pbSelector = nil
        })
    }
    l.pbSelector.Show()
    l.pbSelector.Raise()
}

func (l *LPEdit) updateParameter(param pedal.Parameter) {
    param.LockData()
    defer param.UnlockData()
    pbi := param.GetParent()
    switch p := pbi.(type) {
    case *pedal.Pedal:
        l.pedals[p.GetID()-4].updateParam(param)
    }
}

func (l *LPEdit) updatePedalBoard(pb *pedal.PedalBoard) {
    pb.LockData()
    defer pb.UnlockData()
    for i, p := range l.pedals {
        p.updatePedal(pb.GetPedal2(i))
    }
}

func (l *LPEdit) updateTempo(pb *pedal.PedalBoard) {
    pb.LockData()
    defer pb.UnlockData()
    l.Tempo.SetText(fmt.Sprintf("%.2f", pb.GetTempo()))
}

func (l *LPEdit) updateType(pbi pedal.PedalBoardItem) {
    pbi.LockData()
    defer pbi.UnlockData()
    switch p := pbi.(type) {
    case *pedal.Pedal:
        l.pedals[p.GetID()-4].updatePedal(p)
    }
}

func (l *LPEdit) windowClose(event *gui.QCloseEvent) {
    if l.ctrl.IsStarted() {
        l.ctrl.Stop()
    }
}
