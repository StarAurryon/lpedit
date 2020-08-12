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
    parameters [11]Parameter
    parent     *LPEdit
}

func NewAmp(parent *LPEdit, w widgets.QWidget_ITF, c *qtctrl.Controller,
        at []string, id int, name string) *Amp {
    a := &Amp{AmpUI: NewAmpUI(w), ctrl: c, ampType: at, id: id}
    a.parameters[0] = Parameter{label: a.Param0Lbl, mid: a.Param0Mid,
        value: a.Param0Value, vfunc: a.parameter0Changed}
    a.parameters[1] = Parameter{label: a.Param1Lbl, mid: a.Param1Mid,
        value: a.Param1Value, vfunc: a.parameter1Changed}
    a.parameters[2] = Parameter{label: a.Param2Lbl, mid: a.Param2Mid,
        value: a.Param2Value, vfunc: a.parameter2Changed}
    a.parameters[3] = Parameter{label: a.Param3Lbl, mid: a.Param3Mid,
        value: a.Param3Value, vfunc: a.parameter3Changed}
    a.parameters[4] = Parameter{label: a.Param4Lbl, mid: a.Param4Mid,
        value: a.Param4Value, vfunc: a.parameter4Changed}
    a.parameters[5] = Parameter{label: a.Param5Lbl, mid: a.Param5Mid,
        value: a.Param5Value, vfunc: a.parameter5Changed}
    a.parameters[6] = Parameter{label: a.Param6Lbl, mid: a.Param6Mid,
        value: a.Param6Value, vfunc: a.parameter6Changed}
    a.parameters[7] = Parameter{label: a.Param7Lbl, mid: a.Param7Mid,
        value: a.Param7Value, vfunc: a.parameter7Changed}
    a.parameters[8] = Parameter{label: a.Param8Lbl, mid: a.Param8Mid,
        value: a.Param8Value, vfunc: a.parameter8Changed}
    a.parameters[9] = Parameter{label: a.Param9Lbl, mid: a.Param9Mid,
        value: a.Param9Value, vfunc: a.parameter9Changed}
    a.parameters[10] = Parameter{label: a.Param10Lbl, mid: a.Param10Mid,
        value: a.Param10Value, vfunc: a.parameter10Changed}
    a.AmpName.SetText(name)
    a.parent = parent
    a.init()
    a.initUI()
    return a
}

func (a *Amp) connectSignal() {
    a.AmpModel.ConnectActivated2(a.ampModelUserChanged)
    for _, p := range a.parameters {
        p.value.ConnectActivated2(p.vfunc)
        p.value.SetEditable(true)
    }
}

func (a *Amp) disconnectSignal() {
    a.AmpModel.DisconnectActivated2()
    for _, p := range a.parameters {
        p.value.DisconnectActivated2()
        p.value.SetEditable(false)
    }
}

func (a *Amp) initUI() {
    keys := make([]string, 0, len(a.ampType))

    for _, k := range a.ampType {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    a.AmpModel.AddItems(keys)
}

func (a *Amp) ampModelUserChanged(fxType string) {
    a.ctrl.SetAmpType(uint32(a.id), a.AmpModel.CurrentText())
}

func (a *Amp) getParameter(id uint32) *Parameter{
    for i, param := range a.parameters {
        if param.id == id {
            return &a.parameters[i]
        }
    }
    return nil
}

func (a *Amp) hideParameter(p *Parameter) {
    p.label.Hide()
    p.mid.Hide()
    p.value.Hide()
}

func (a *Amp) parameter0Changed(val string) { a.parameterChanged(&a.parameters[0], val) }
func (a *Amp) parameter1Changed(val string) { a.parameterChanged(&a.parameters[1], val) }
func (a *Amp) parameter2Changed(val string) { a.parameterChanged(&a.parameters[2], val) }
func (a *Amp) parameter3Changed(val string) { a.parameterChanged(&a.parameters[3], val) }
func (a *Amp) parameter4Changed(val string) { a.parameterChanged(&a.parameters[4], val) }
func (a *Amp) parameter5Changed(val string) { a.parameterChanged(&a.parameters[5], val) }
func (a *Amp) parameter6Changed(val string) { a.parameterChanged(&a.parameters[6], val) }
func (a *Amp) parameter7Changed(val string) { a.parameterChanged(&a.parameters[7], val) }
func (a *Amp) parameter8Changed(val string) { a.parameterChanged(&a.parameters[8], val) }
func (a *Amp) parameter9Changed(val string) { a.parameterChanged(&a.parameters[9], val) }
func (a *Amp) parameter10Changed(val string) { a.parameterChanged(&a.parameters[10], val) }

func (a *Amp) parameterChanged(paramUI *Parameter, val string) {
    err := a.ctrl.SetAmpParameterValue(uint32(a.id), paramUI.id, val)
    if err != nil {
        mb := widgets.NewQMessageBox(a)
        mb.Critical(a, "An error occured", err.Error(), widgets.QMessageBox__Ok, 0)
    }
    pb := a.ctrl.GetPedalBoard()
    pb.LockData()
    param := pb.GetAmp(a.id).GetParam(paramUI.id)
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
    for i := range a.parameters {
        a.parameters[i].id = 0
        a.hideParameter(&a.parameters[i])
    }
    for i, param := range amp.GetParams() {
        a.parameters[i].id = param.GetID()
        a.updateParam(param)
    }
}

func (a *Amp) updateParam(p pedal.Parameter) {
    param := a.getParameter(p.GetID())
    if param == nil { return }
    a.setParameterValueList(param, []string{p.GetValueCurrent()})
    a.setParameterValue(param, p.GetValueCurrent())
    a.setParameterLabel(param, p.GetName())
    a.showParameter(param)
}

func (a *Amp) setActive(status bool) {
    a.OnStatus.SetChecked(status)
}

func (a *Amp) setParameterLabel(p *Parameter, s string) {
    p.label.SetText(s)
}

func (a *Amp) setParameterValueList(p *Parameter, s []string) {
    p.value.Clear()
    p.value.AddItems(s)
}

func (a *Amp) setParameterValue(p *Parameter, s string) {
    p.value.SetCurrentText(s)
}

func (a *Amp) showParameter(p *Parameter) {
    p.label.Show()
    p.mid.Show()
    p.value.Show()
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
