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
import "github.com/StarAurryon/qt/gui"
import "github.com/StarAurryon/qt/widgets"

import "fmt"

import "github.com/StarAurryon/lpedit/pedal"
import "github.com/StarAurryon/lpedit/qtctrl"

type LPEdit struct {
    *LPEditUI
    about      *AboutUI
    pbSelector *PBSelector
    ctrl       *qtctrl.Controller
    amps       []*Amp
    cabs       []*Cab
    pedals     []*Pedal
    parameters []Parameter
}

func NewLPEdit(c *qtctrl.Controller, p widgets.QWidget_ITF) *LPEdit {
    l := &LPEdit{LPEditUI: NewLPEditUI(p), ctrl: c}
    l.parameters = make([]Parameter, 3)
    l.parameters[0] = Parameter{label: l.InputSource1Lbl,
        value: l.InputSource1, vfunc: l.parameter0Changed}
    l.parameters[1] = Parameter{label: l.InputSource2Lbl,
        value: l.InputSource2, vfunc: l.parameter1Changed}
    l.parameters[2] = Parameter{label: l.GuitarInZLbl,
        value: l.GuitarInZ, vfunc: l.parameter2Changed}
    l.init()
    return l
}

func (l *LPEdit) connectSignal() {
    l.updateSets()
    l.updatePresets(0)

    //UI Connections
    l.SetList.ConnectCurrentIndexChanged(l.updatePresets)
    //PedalBoard Connections
    l.ctrl.ConnectActiveChange(l.updateActive)
    l.ctrl.ConnectParameterChange(l.updateParameter)
    l.ctrl.ConnectPresetLoad(l.updatePedalBoard)
    l.ctrl.ConnectSetChange(l.updateSet)
    l.ctrl.ConnectTempoChange(l.updateTempo)
    l.ctrl.ConnectTypeChange(l.updateType)

    for _, amp := range l.amps {
        amp.connectSignal()
    }
    for _, cab := range l.cabs {
        cab.connectSignal()
    }
    for _, pedal := range l.pedals {
        pedal.connectSignal()
    }

    for _, p := range l.parameters {
        p.value.ConnectActivated2(p.vfunc)
        p.value.SetEditable(true)
    }
}

func (l *LPEdit) disconnectSignal() {
    //UI Connections
    l.SetList.DisconnectCurrentIndexChanged()
    //PedalBoard Connections
    l.ctrl.DisconnectActiveChange()
    l.ctrl.DisconnectParameterChange()
    l.ctrl.DisconnectPresetLoad()
    l.ctrl.DisconnectSetChange()
    l.ctrl.DisconnectTempoChange()
    l.ctrl.DisconnectTypeChange()

    for _, amp := range l.amps {
        amp.disconnectSignal()
    }
    for _, cab := range l.cabs {
        cab.disconnectSignal()
    }
    for _, pedal := range l.pedals {
        pedal.disconnectSignal()
    }

    for _, p := range l.parameters {
        p.value.DisconnectActivated2()
        p.value.SetEditable(false)
    }
}

func (l *LPEdit) init() {
    //init
    l.initAmpsCabs()
    l.initPedals()

    //UI Connections
    l.ConnectCloseEvent(l.windowClose)
    l.ActionAbout.ConnectTriggered(l.aboutClick)
    l.ActionSelect_Device.ConnectTriggered(l.pbSelectorClick)
    l.ActionQuit.ConnectTriggered(func(bool) {l.Close()})
}

func (l *LPEdit) initAmpsCabs() {
    ampType := pedal.GetAmpType()
    cabType := pedal.GetCabType()

    names := []string{"Amp A", "Amp B"}

    for i := 0; i <= len(names); i++ {
        line := widgets.NewQFrame(l.ScrollPedalW, core.Qt__Widget)
        line.SetFrameShape(widgets.QFrame__HLine)
        line.SetFrameShadow(widgets.QFrame__Sunken)
        l.ScrollAmpW.Layout().AddWidget(line)
        if i < len(names) {
            a := NewAmp(l, l.ScrollAmpW, l.ctrl, ampType, i, names[i])
            l.ScrollAmpW.Layout().AddWidget(a)
            l.amps = append(l.amps, a)
            line := widgets.NewQFrame(l.ScrollPedalW, core.Qt__Widget)
            line.SetFrameShape(widgets.QFrame__HLine)
            line.SetFrameShadow(widgets.QFrame__Sunken)
            l.ScrollAmpW.Layout().AddWidget(line)
            c := NewCab(l, l.ScrollAmpW, l.ctrl, cabType, i)
            l.ScrollAmpW.Layout().AddWidget(c)
            l.cabs = append(l.cabs, c)
        }
    }
}

func (l *LPEdit) initPedals() {
    pedalType := pedal.GetPedalType()

    for i := 0; i < 9; i++ {
        line := widgets.NewQFrame(l.ScrollPedalW, core.Qt__Widget)
        line.SetFrameShape(widgets.QFrame__HLine)
        line.SetFrameShadow(widgets.QFrame__Sunken)
        l.PedalsLayout.AddWidget(line, 0, 0)
        if i != 8 {
            p := NewPedal(l, l.ScrollPedalW, l.ctrl, pedalType, i)
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

func (l *LPEdit) getParameter(id uint32) *Parameter{
    for i, param := range l.parameters {
        if param.id == id {
            return &l.parameters[i]
        }
    }
    return nil
}

func (l *LPEdit) parameter0Changed(val string) { l.parameterChanged(&l.parameters[0], val) }
func (l *LPEdit) parameter1Changed(val string) { l.parameterChanged(&l.parameters[1], val) }
func (l *LPEdit) parameter2Changed(val string) { l.parameterChanged(&l.parameters[2], val) }
func (l *LPEdit) parameter3Changed(val string) { l.parameterChanged(&l.parameters[3], val) }

func (l *LPEdit) parameterChanged(param *Parameter, val string) {
    l.ctrl.SetPedalBoardParameterValue(param.id, val)
}

func (l *LPEdit) pbSelectorClick(vbo bool) {
    if l.pbSelector == nil {
        l.pbSelector = NewPBSelector(l.ctrl, l, l)
        l.pbSelector.ConnectCloseEvent(func(event *gui.QCloseEvent) {
            l.pbSelector.DestroyQObject()
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
    case *pedal.Amp:
        l.amps[p.GetID()/2].setActive(p.GetActive())
    case *pedal.Pedal:
        l.pedals[p.GetID()-4].setActive(p.GetActive())
    }
}

func (l *LPEdit) updateParameter(param pedal.Parameter) {
    param.LockData()
    defer param.UnlockData()
    pbi := param.GetParent()
    switch p := pbi.(type) {
    case *pedal.Amp:
        l.amps[p.GetID()/2].updateParam(param)
    case *pedal.Cab:
        l.cabs[p.GetID()/2].updateParam(param)
    case *pedal.Pedal:
        l.pedals[p.GetID()-4].updateParam(param)
    case *pedal.PedalBoard:
        l.updateParam(param)
    }
}

func (l *LPEdit) updateParam(p pedal.Parameter) {
    param := l.getParameter(p.GetID())
    if param == nil { return }
    values := p.GetAllowedValues()

    valueIn := false
    for _, v := range values {
        if v == p.GetValueCurrent() {
            valueIn = true
            break
        }
    }

    if !valueIn {
        values = append([]string{p.GetValueCurrent()}, values...)
    }

    param.setValueList(values)
    if p.IsAllowingOtherValues() {
        param.setValueEditable(true)
    } else {
        param.setValueEditable(false)
    }
    param.setValue(p.GetValueCurrent())
}

func  (l *LPEdit) updateParameters(pb *pedal.PedalBoard) {
    for i, param := range pb.GetParams() {
        l.parameters[i].id = param.GetID()
        l.updateParam(param)
    }
}

func (l *LPEdit) updatePedalBoard(pb *pedal.PedalBoard) {
    pb.LockData()
    defer pb.UnlockData()
    l.updatePreset(pb)
    l.updatePedalBoardView(pb)
    l.updateParameters(pb)
    for i, a := range l.amps {
        a.updateAmp(pb.GetAmp(i))
    }
    for i, c := range l.cabs {
        c.updateCab(pb.GetCab(i))
    }
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

    ampA := pb.GetItems(pedal.AmpAPos)[0]
    posA, _ := ampA.GetPos()
    aStart := pb.GetItems(pedal.PedalPosAStart)
    aEnd := pb.GetItems(pedal.PedalPosAEnd)
    ampB := pb.GetItems(pedal.AmpBPos)[0]
    bStart := pb.GetItems(pedal.PedalPosBStart)
    bEnd := pb.GetItems(pedal.PedalPosBEnd)

    l.PedalBoardView.SetLayout(widgets.NewQHBoxLayout2(l.PedalBoardView))

    //Start
    startItems := pb.GetItems(pedal.PedalPosStart)
    if posA == 5 { startItems = append(startItems, ampA) }
    fillLayout(l.PedalBoardView.Layout(), true, startItems...)

    //Split
    var top, bot []pedal.PedalBoardItem
    if posA != 0 {
        top = append(aStart, aEnd...)
        bot = append(bStart, bEnd...)
    } else {
        top = append(append(aStart, ampA), aEnd...)
        bot = append(append(bStart, ampB), bEnd...)
    }

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

    pbiUI := widgets.NewQLabel(l.PedalBoardView, core.Qt__Widget)
    pbiUI.SetText("Main Mix")
    l.PedalBoardView.Layout().AddWidget(pbiUI)

    //End
    endItems := pb.GetItems(pedal.PedalPosEnd)
    if posA == 7 { endItems = append(endItems, ampA) }
    fillLayout(l.PedalBoardView.Layout(), true, endItems...)
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

    l.PresetList.Clear()

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

    l.SetList.Clear()
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
    case *pedal.Amp:
        l.amps[p.GetID()/2].updateAmp(p)
    case *pedal.Pedal:
        l.pedals[p.GetID()-4].updatePedal(p)
    }
    l.updatePedalBoardView(l.ctrl.GetPedalBoard())
}

func (l *LPEdit) windowClose(event *gui.QCloseEvent) {
    if l.ctrl.IsStarted() {
        l.ctrl.Stop()
    }
}
