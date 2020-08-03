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
    //init
    l.updateSets()
    l.updatePresets(0)
    l.initPedals()

    //UI Connections
    l.ConnectCloseEvent(l.windowClose)
    l.ActionAbout.ConnectTriggered(l.aboutClick)
    l.ActionSelect_Device.ConnectTriggered(l.pbSelectorClick)
    l.ActionQuit.ConnectTriggered(func(bool) {l.Close()})
    l.SetList.ConnectCurrentIndexChanged(l.updatePresets)

    //PedalBoard Connections
    l.ctrl.ConnectActiveChange(l.updateActive)
    l.ctrl.ConnectParameterChange(l.updateParameter)
    l.ctrl.ConnectPresetLoad(l.updatePedalBoard)
    l.ctrl.ConnectSetChange(l.updateSet)
    l.ctrl.ConnectTempoChange(l.updateTempo)
    l.ctrl.ConnectTypeChange(l.updateType)
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

func (l *LPEdit) updateActive(pbi pedal.PedalBoardItem) {
    pbi.LockData()
    defer pbi.UnlockData()
    switch p := pbi.(type) {
    case *pedal.Pedal:
        l.pedals[p.GetID()-4].setActive(p.GetActive())
    }
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
    l.updatePreset(pb)
    l.updatePedalBoardView(pb)
    for i, p := range l.pedals {
        p.updatePedal(pb.GetPedal2(i))
    }
}

func (l *LPEdit) updatePedalBoardView(pb *pedal.PedalBoard) {
    fillLayout := func (lay widgets.QLayout_ITF, line bool, pbis ...pedal.PedalBoardItem) {
        max := len(pbis)
        for i := 0; i < (max + 1); i ++ {
            if line {
                AddLine(lay, l.PedalBoardView, widgets.QFrame__HLine)
            }
            if i < max {
                pbi := pbis[i]
                pbiUI := widgets.NewQLabel(l.PedalBoardView, core.Qt__Widget)
                pbiUI.SetText(pbi.GetName())
                AddWidget(lay, pbiUI)
            }
        }
    }

    for _, item := range l.PedalBoardView.Children(){
        item.DestroyQObject()
    }

    l.PedalBoardView.SetLayout(widgets.NewQHBoxLayout2(l.PedalBoardView))

    fillLayout(l.PedalBoardView.Layout(), true, pb.GetItems(pedal.PedalPosStart)...)

    ampA := pb.GetItems(pedal.AmpAPos)
    aStart := pb.GetItems(pedal.PedalPosAStart)
    aEnd := pb.GetItems(pedal.PedalPosAEnd)
    ampB := pb.GetItems(pedal.AmpBPos)
    bStart := pb.GetItems(pedal.PedalPosBStart)
    bEnd := pb.GetItems(pedal.PedalPosBEnd)

    if (len(ampB) + len(bStart) + len(bEnd) + len(ampA) + len(aStart) + len(aEnd)) > 0 {
        top := append(append(aStart, ampA...), aEnd...)
        bot := append(append(bStart, ampB...), bEnd...)

        max := Max(len(top), len(bot))
        for i := 0; i < (max + 1); i++ {
            split := widgets.NewQWidget(l.PedalBoardView, core.Qt__Widget)
            split.SetLayout(widgets.NewQVBoxLayout2(split))
            l.PedalBoardView.Layout().AddWidget(split)
            AddLine(split.Layout(), split, widgets.QFrame__HLine)
            AddLine(split.Layout(), split, widgets.QFrame__HLine)

            if i < max {
                split = widgets.NewQWidget(l.PedalBoardView, core.Qt__Widget)
                split.SetLayout(widgets.NewQVBoxLayout2(split))
                l.PedalBoardView.Layout().AddWidget(split)
                if i < len(top) {
                    fillLayout(split.Layout(), false, top[i])
                } else {
                    AddLine(split.Layout(), split, widgets.QFrame__HLine)
                }
                if i < len(bot) {
                    fillLayout(split.Layout(), false, bot[i])
                } else {
                    AddLine(split.Layout(), split, widgets.QFrame__HLine)
                }
            }
        }
    }

    pbiUI := widgets.NewQLabel(l.PedalBoardView, core.Qt__Widget)
    pbiUI.SetText("Main Mix")
    l.PedalBoardView.Layout().AddWidget(pbiUI)

    fillLayout(l.PedalBoardView.Layout(), true, pb.GetItems(pedal.PedalPosEnd)...)
}

func (l *LPEdit) updatePreset(pb *pedal.PedalBoard) {
    err, id := pb.GetCurrentPreset()
    pname := pb.GetCurrentPresetName()

    if err != nil { return }
    for i, s := range pname {
        l.PresetList.Item(int(id), i).SetText(s)
        l.PresetList.Item(int(id), i).SetText(s)
    }
    l.PresetList.SelectRow(int(id))
}

func (l *LPEdit) updatePresets(index int) {
    pb := l.ctrl.GetPedalBoard()
    pb.LockData()
    presets := pb.GetPresetList(index)
    pb.UnlockData()

    /*for _, c := range l.PresetList.Children() {
        c.DestroyQObject()
    }*/
    l.PresetList.HorizontalHeader().SetVisible(false)
    l.PresetList.VerticalHeader().SetVisible(false)
    l.PresetList.SetRowCount(len(presets))
    l.PresetList.SetColumnCount(2)
    l.PresetList.SetColumnWidth(0, 10)
    l.PresetList.SetColumnWidth(1, 200)
    for i, preset := range presets {
        id := widgets.NewQTableWidgetItem2(preset[0], 0)
        id.SetFlags(core.Qt__ItemIsSelectable|core.Qt__ItemIsEnabled)
        l.PresetList.SetItem(i, 0, id)
        l.PresetList.SetItem(i, 1, widgets.NewQTableWidgetItem2(preset[1], 0))
    }
}

func (l *LPEdit) updateSet(pb *pedal.PedalBoard) {
    pb.LockData()
    err, id := pb.GetCurrentSet()
    pb.UnlockData()

    if err != nil { return }
    l.SetList.SetCurrentIndex(int(id))
}

func (l *LPEdit) updateSets() {
    pb := l.ctrl.GetPedalBoard()
    pb.LockData()
    setList := pb.GetSetList()
    pb.UnlockData()

    l.SetList.AddItems(setList)
}

func (l *LPEdit) updateTempo(pb *pedal.PedalBoard) {
    pb.LockData()
    text := fmt.Sprintf("%.2f", pb.GetTempo())
    pb.UnlockData()

    l.Tempo.SetText(text)
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
