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

import "log"
import "io/ioutil"
import "os"

import "github.com/StarAurryon/lpedit-lib/model/pod"
import "github.com/StarAurryon/lpedit/qtctrl"

type LPEdit struct {
    *LPEditUI
    about       *AboutUI
    pbSelector  *PBSelector
    ctrl        *qtctrl.Controller
    amps        []*Amp
    cabs        []*Cab
    pedals      []*Pedal
    parameters  [4]Parameter
    progressWnd *widgets.QProgressDialog
}

func NewLPEdit(c *qtctrl.Controller, p widgets.QWidget_ITF) *LPEdit {
    l := &LPEdit{LPEditUI: NewLPEditUI(p), ctrl: c}
    l.parameters[0] = Parameter{label: l.InputSource1Lbl,
        value: l.InputSource1, vfunc: l.parameter0Changed}
    l.parameters[1] = Parameter{label: l.InputSource2Lbl,
        value: l.InputSource2, vfunc: l.parameter1Changed}
    l.parameters[2] = Parameter{label: l.GuitarInZLbl,
        value: l.GuitarInZ, vfunc: l.parameter2Changed}
    l.parameters[3] = Parameter{label: l.TempoLbl,
        value: l.Tempo, vfunc: l.parameter3Changed}
    log.SetOutput(ioutil.Discard)
    l.init()
    return l
}

func (l *LPEdit) init() {
    //Icon
    ps := string(os.PathSeparator)
    iconPath := "ui" + ps + "knob.png"
    icon := gui.NewQIcon5(iconPath)
    l.SetWindowIcon(icon)

    //init
    l.initAmpsCabs()
    l.initPedals()

    //UI Connections
    l.ConnectCloseEvent(l.windowClose)
    l.ActionAbout.ConnectTriggered(l.aboutClick)
    l.ActionDebug.ConnectTriggered(l.debugChange)
    l.ActionSelect_Device.ConnectTriggered(l.pbSelectorClick)
    l.ActionQuit.ConnectTriggered(func(bool) {l.Close()})
    l.ctrl.ConnectLoopError(l.loopError)
    l.ctrl.ConnectProgress(l.progress)

    //PresetList
    l.PresetList.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
    l.PresetList.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
}

func (l *LPEdit) connectSignal() {
    s := l.ctrl.GetPod().GetCurrentSet()
    p := l.ctrl.GetPod().GetCurrentPreset()
    l.updateSets()
    l.updateSet(s)
    l.updatePresets(l.SetList.CurrentIndex())
    l.updatePreset(p)
    l.updatePedalBoard(p)

    //UI Connections
    l.DiscardChanges.ConnectClicked(l.discardPresetChanges)
    l.PresetList.ConnectItemChanged(l.updateCurrentPresetName)
    l.PresetList.ConnectClicked(l.changePreset)
    l.SetList.ConnectCurrentIndexChanged(l.updatePresets)
    l.Save.ConnectClicked(l.savePreset)
    //PedalBoard Connections
    l.ctrl.ConnectActiveChange(l.updateActive)
    l.ctrl.ConnectParameterChange(l.updateParameter)
    l.ctrl.ConnectPresetChange(l.updatePreset)
    l.ctrl.ConnectPresetLoad(l.updatePedalBoard)
    l.ctrl.ConnectSetChange(l.updateSet)
    l.ctrl.ConnectTypeChange(l.updateType)

    for _, amp := range l.amps {
        amp.connectSignal()
    }
    for _, cab := range l.cabs {
        cab.connectSignal()
    }
    for _, pedal:= range l.pedals {
        pedal.connectSignal()
    }

    for _, p := range l.parameters {
        p.value.ConnectActivated2(p.vfunc)
    }
}

func (l *LPEdit) disconnectSignal() {
    //UI Connections
    l.DiscardChanges.DisconnectClicked()
    l.PresetList.DisconnectItemChanged()
    l.PresetList.DisconnectClicked()
    l.SetList.DisconnectCurrentIndexChanged()
    l.Save.DisconnectClicked()
    //PedalBoard Connections
    l.ctrl.DisconnectActiveChange()
    l.ctrl.DisconnectParameterChange()
    l.ctrl.DisconnectPresetChange()
    l.ctrl.DisconnectPresetLoad()
    l.ctrl.DisconnectSetChange()
    l.ctrl.DisconnectTypeChange()

    for _, amp := range l.amps {
        amp.disconnectSignal()
    }
    for _, cab := range l.cabs {
        cab.disconnectSignal()
    }
    for _, pedal:= range l.pedals {
        pedal.disconnectSignal()
    }

    for _, p := range l.parameters {
        p.value.DisconnectActivated2()
        p.value.SetEditable(false)
    }
}

func (l *LPEdit) initAmpsCabs() {
    ampType := pod.GetAmpType()
    cabType := pod.GetCabType()

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
    pedalype := pod.GetPedalType()

    for i := 0; i < 9; i++ {
        line := widgets.NewQFrame(l.ScrollPedalW, core.Qt__Widget)
        line.SetFrameShape(widgets.QFrame__HLine)
        line.SetFrameShadow(widgets.QFrame__Sunken)
        l.PedalsLayout.AddWidget(line, 0, 0)
        if i != 8 {
            p := NewPedal(l, l.ScrollPedalW, l.ctrl, pedalype, i)
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

func (l *LPEdit) debugChange(status bool) {
    if status {
        log.SetOutput(os.Stdout)
    } else {
        log.SetOutput(ioutil.Discard)
    }
}

func (l *LPEdit) getParameter(id uint32) *Parameter{
    for i, param := range l.parameters {
        if param.id == id {
            return &l.parameters[i]
        }
    }
    return nil
}

func (l *LPEdit) changePreset(model *core.QModelIndex) {
    pod := l.ctrl.GetPod()
    pod.LockData()
    currentPreset := pod.GetCurrentPreset()
    pod.UnlockData()

    if currentPreset == nil {
        return
    }

    if model.Row() != int(currentPreset.GetID()) || l.SetList.CurrentIndex() != int(currentPreset.GetSet().GetID()) {
        l.ctrl.SetPreset(uint8(model.Row()), uint8(l.SetList.CurrentIndex()))
    }
}

func (l *LPEdit) discardPresetChanges(bool) {
    l.ctrl.ReloadPreset()
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

func (l *LPEdit) savePreset(bool) {
    l.ctrl.SavePreset()
}

func (l *LPEdit) updateActive(pbi pod.PedalBoardItem) {
    pbi.LockData()
    defer pbi.UnlockData()
    switch p := pbi.(type) {
    case *pod.Amp:
        l.amps[p.GetID()/2].setActive(p.GetActive())
    case *pod.Pedal:
        l.pedals[p.GetID()-4].setActive(p.GetActive())
    }
}

func (l *LPEdit) updateCurrentPresetName(item *widgets.QTableWidgetItem) {
    l.ctrl.SetCurrentPresetName(item.Text())
}

func (l *LPEdit) updateParameter(param pod.Parameter) {
    param.LockData()
    defer param.UnlockData()
    pbi := param.GetParent()
    switch p := pbi.(type) {
    case *pod.Amp:
        l.amps[p.GetID()/2].updateParam(param)
    case *pod.Cab:
        l.cabs[p.GetID()/2].updateParam(param)
    case *pod.Pedal:
        l.pedals[p.GetID()-4].updateParam(param)
    case *pod.Preset:
        l.updateParam(param)
    }
}

func (l *LPEdit) updateParam(p pod.Parameter) {
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

func  (l *LPEdit) updateParameters(pb *pod.Preset) {
    for i, param := range pb.GetParams() {
        l.parameters[i].id = param.GetID()
        l.updateParam(param)
    }
}

func (l *LPEdit) updatePedalBoard(p *pod.Preset) {
    p.LockData()
    defer p.UnlockData()
    l.updatePedalBoardView(p)
    l.updateParameters(p)
    for i, a := range l.amps {
        a.updateAmp(p.GetAmp(i))
    }
    for i, c := range l.cabs {
        c.updateCab(p.GetCab(i))
    }
    for i, pedal := range l.pedals {
        pedal.updatePedal(p.GetPedal2(i))
    }
}

func (l *LPEdit) updatePreset(p *pod.Preset) {
    id := p.GetID()
    pname := p.GetName3()

    for i, s := range pname {
        l.PresetList.Item(int(id), i).SetText(s)
    }
    l.PresetList.SelectRow(int(id))
}

func (l *LPEdit) updatePresets(index int) {
    pod := l.ctrl.GetPod()
    pod.LockData()
    presets := pod.GetSet(uint8(index)).GetPresetList()
    pod.UnlockData()

    l.PresetList.Clear()

    l.PresetList.HorizontalHeader().SetVisible(false)
    l.PresetList.VerticalHeader().SetVisible(false)
    l.PresetList.SetRowCount(len(presets))
    l.PresetList.SetColumnCount(2)

    for i, preset := range presets {
        id := widgets.NewQTableWidgetItem2(preset[0], 0)
        id.SetFlags(core.Qt__ItemIsSelectable|core.Qt__ItemIsEnabled)
        l.PresetList.SetItem(i, 0, id)
        l.PresetList.SetItem(i, 1, widgets.NewQTableWidgetItem2(preset[1], 0))
    }
    l.PresetList.ResizeColumnsToContents()
}

func (l *LPEdit) updateSet(p *pod.Set) {
    p.LockData()
    id := p.GetID()
    p.UnlockData()

    l.SetList.SetCurrentIndex(int(id))
}

func (l *LPEdit) updateSets() {
    pod := l.ctrl.GetPod()
    pod.LockData()
    setList := pod.GetSetList()
    pod.UnlockData()

    l.SetList.Clear()
    l.SetList.AddItems(setList)
}

func (l *LPEdit) updateType(pbi pod.PedalBoardItem) {
    pbi.LockData()
    defer pbi.UnlockData()
    switch p := pbi.(type) {
    case *pod.Amp:
        l.amps[p.GetID()/2].updateAmp(p)
    case *pod.Pedal:
        l.pedals[p.GetID()-4].updatePedal(p)
    }
    l.updatePedalBoardView(pbi.GetPreset())
}

func (l *LPEdit) windowClose(event *gui.QCloseEvent) {
    if l.ctrl.IsStarted() {
        l.ctrl.Stop()
    }
}
