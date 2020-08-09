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

import "sort"

import "github.com/StarAurryon/lpedit/pedal"
import "github.com/StarAurryon/lpedit/qtctrl"

type Amp struct {
    *AmpUI
    id         int
    ampType    []string
    ctrl       *qtctrl.Controller
    labels     []*widgets.QLabel
    mids       []*widgets.QWidget
    values     []*widgets.QComboBox
    valuesFunc []func(string)
    parent     *LPEdit
}

func NewAmp(parent *LPEdit, w widgets.QWidget_ITF, c *qtctrl.Controller,
        at []string, id int, name string) *Amp {
    a := &Amp{AmpUI: NewAmpUI(w), ctrl: c, ampType: at, id: id}
    a.labels = []*widgets.QLabel{
        a.Param0Lbl, a.Param1Lbl , a.Param2Lbl, a.Param3Lbl, a.Param4Lbl,
        a.Param5Lbl, a.Param6Lbl , a.Param7Lbl, a.Param8Lbl, a.Param9Lbl,
        a.Param10Lbl,
    }
    a.mids = []*widgets.QWidget{
        a.Param0Mid, a.Param1Mid , a.Param2Mid, a.Param3Mid, a.Param4Mid,
        a.Param5Mid, a.Param6Mid , a.Param7Mid, a.Param8Mid, a.Param9Mid,
        a.Param10Mid,
    }
    a.values = []*widgets.QComboBox{
        a.Param0Value, a.Param1Value , a.Param2Value, a.Param3Value,
        a.Param4Value, a.Param5Value, a.Param6Value, a.Param7Value,
        a.Param8Value, a.Param9Value, a.Param10Value,
    }
    a.valuesFunc = []func(string) {
        a.parameter0Changed, a.parameter1Changed, a.parameter2Changed,
        a.parameter3Changed, a.parameter4Changed, a.parameter5Changed,
        a.parameter6Changed, a.parameter7Changed, a.parameter8Changed,
        a.parameter9Changed, a.parameter10Changed,
    }
    a.AmpName.SetText(name)
    a.parent = parent
    a.init()
    a.initUI()
    return a
}

func (a *Amp) initUI() {
    keys := make([]string, 0, len(a.ampType))

    for _, k := range a.ampType {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    a.AmpModel.ConnectActivated2(a.ampModelUserChanged)
    a.AmpModel.AddItems(keys)
    for i := range a.values {
        a.values[i].ConnectActivated2(a.valuesFunc[i])
        a.values[i].SetEditable(true)
    }
}

func (a *Amp) ampModelUserChanged(fxType string) {
    a.ctrl.SetAmpType(uint32(a.id), a.AmpModel.CurrentText())
}

func (a *Amp) hideParameter(id int) {
    if id < 0 || id >= len(a.labels) { return }
    a.labels[id].Hide()
    a.mids[id].Hide()
    a.values[id].Hide()
}

func (a *Amp) parameter0Changed(val string) { a.parameterChanged(0, val) }
func (a *Amp) parameter1Changed(val string) { a.parameterChanged(1, val) }
func (a *Amp) parameter2Changed(val string) { a.parameterChanged(2, val) }
func (a *Amp) parameter3Changed(val string) { a.parameterChanged(3, val) }
func (a *Amp) parameter4Changed(val string) { a.parameterChanged(4, val) }
func (a *Amp) parameter5Changed(val string) { a.parameterChanged(5, val) }
func (a *Amp) parameter6Changed(val string) { a.parameterChanged(6, val) }
func (a *Amp) parameter7Changed(val string) { a.parameterChanged(7, val) }
func (a *Amp) parameter8Changed(val string) { a.parameterChanged(8, val) }
func (a *Amp) parameter9Changed(val string) { a.parameterChanged(9, val) }
func (a *Amp) parameter10Changed(val string) { a.parameterChanged(10, val) }

func (a *Amp) parameterChanged(id int, val string) {
    err := a.ctrl.SetAmpParameterValue(uint32(a.id), uint16(id), val)
    if err != nil {
        mb := widgets.NewQMessageBox(a)
        mb.Critical(a, "An error occured", err.Error(), widgets.QMessageBox__Ok, 0)
    }
    pb := a.ctrl.GetPedalBoard()
    pb.LockData()
    param := pb.GetAmp(a.id).GetParam(uint16(id))
    a.updateParam(param)
    pb.UnlockData()
}

func (a *Amp) updateAmp(amp *pedal.Amp) {
    if a.id == 0 {
        if pos, _ := amp.GetPos(); pos != 0 {
            a.parent.amps[1].Hide()
        } else {
            a.parent.amps[1].Show()
        }
    }
    a.setActive(amp.GetActive())
    a.AmpModel.SetCurrentText(amp.GetName())
    for i := range a.labels {
        a.hideParameter(i)
    }
    for _, param := range amp.GetParams() {
        a.updateParam(param)
    }
}

func (a *Amp) updateParam(p pedal.Parameter) {
    id := int(p.GetID())
    if id > len(a.labels) { return }
    a.setParameterValueList(id, []string{p.GetValue()})
    a.setParameterValue(id, p.GetValue())
    a.setParameterLabel(id, p.GetName())
    a.showParameter(id)
}

func (a *Amp) setActive(status bool) {
    a.OnStatus.SetChecked(status)
}

func (a *Amp) setParameterLabel(id int, s string) {
    if id < 0 || id >= len(a.labels) { return }
    a.labels[id].SetText(s)
}

func (a *Amp) setParameterValueList(id int, s []string) {
    if id < 0 || id >= len(a.labels) { return }
    a.values[id].Clear()
    a.values[id].AddItems(s)
}

func (a *Amp) setParameterValue(id int, s string) {
    if id < 0 || id >= len(a.labels) { return }
    a.values[id].SetCurrentText(s)
}

func (a *Amp) showParameter(id int) {
    if id < 0 || id >= len(a.labels) { return }
    a.labels[id].Show()
    a.mids[id].Show()
    a.values[id].Show()
}

func (l *LPEdit) initAmpsCabs() {
    pedalType := l.ctrl.GetAmpType()

    names := []string{"Amp A", "Amp B"}

    for i := 0; i <= len(names); i++ {
        line := widgets.NewQFrame(l.ScrollPedalW, core.Qt__Widget)
        line.SetFrameShape(widgets.QFrame__HLine)
        line.SetFrameShadow(widgets.QFrame__Sunken)
        l.ScrollAmpW.Layout().AddWidget(line)
        if i < len(names) {
            a := NewAmp(l, l.ScrollPedalW, l.ctrl, pedalType, i, names[i])
            l.ScrollAmpW.Layout().AddWidget(a)
            l.amps = append(l.amps, a)
        }
    }
}
