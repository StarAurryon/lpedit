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

import "github.com/therecipe/qt/widgets"

import "sort"
import "strconv"

import "github.com/StarAurryon/lpedit-lib/model/pod"
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
        value: a.Param0Value, knob: a.Param0Knob, vfunc: a.parameter0Changed,
        kfunc: a.parameter0Changed2}
    a.parameters[1] = Parameter{label: a.Param1Lbl, mid: a.Param1Mid,
        value: a.Param1Value, knob: a.Param1Knob, vfunc: a.parameter1Changed,
        kfunc: a.parameter1Changed2}
    a.parameters[2] = Parameter{label: a.Param2Lbl, mid: a.Param2Mid,
        value: a.Param2Value, knob: a.Param2Knob, vfunc: a.parameter2Changed,
        kfunc: a.parameter2Changed2}
    a.parameters[3] = Parameter{label: a.Param3Lbl, mid: a.Param3Mid,
        value: a.Param3Value, knob: a.Param3Knob, vfunc: a.parameter3Changed,
        kfunc: a.parameter3Changed2}
    a.parameters[4] = Parameter{label: a.Param4Lbl, mid: a.Param4Mid,
        value: a.Param4Value, knob: a.Param4Knob, vfunc: a.parameter4Changed,
        kfunc: a.parameter4Changed2}
    a.parameters[5] = Parameter{label: a.Param5Lbl, mid: a.Param5Mid,
        value: a.Param5Value, knob: a.Param5Knob, vfunc: a.parameter5Changed,
        kfunc: a.parameter5Changed2}
    a.parameters[6] = Parameter{label: a.Param6Lbl, mid: a.Param6Mid,
        value: a.Param6Value, knob: a.Param6Knob, vfunc: a.parameter6Changed,
        kfunc: a.parameter6Changed2}
    a.parameters[7] = Parameter{label: a.Param7Lbl, mid: a.Param7Mid,
        value: a.Param7Value, knob: a.Param7Knob, vfunc: a.parameter7Changed,
        kfunc: a.parameter7Changed2}
    a.parameters[8] = Parameter{label: a.Param8Lbl, mid: a.Param8Mid,
        value: a.Param8Value, knob: a.Param8Knob, vfunc: a.parameter8Changed,
        kfunc: a.parameter8Changed2}
    a.parameters[9] = Parameter{label: a.Param9Lbl, mid: a.Param9Mid,
        value: a.Param9Value, knob: a.Param9Knob, vfunc: a.parameter9Changed,
        kfunc: a.parameter9Changed2}
    a.parameters[10] = Parameter{label: a.Param10Lbl, mid: a.Param10Mid,
        value: a.Param10Value, knob: a.Param10Knob, vfunc: a.parameter10Changed,
        kfunc: a.parameter10Changed2}
    a.AmpName.SetText(name)
    a.parent = parent
    a.init()
    return a
}

func (a *Amp) init() {
    keys := make([]string, 0, len(a.ampType))

    for _, k := range a.ampType {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    a.AmpModel.AddItems(keys)
    a.Topology.AddItems(pod.GetAllowedTopology())
}

func (a *Amp) connectSignal() {
    a.AmpModel.ConnectActivated2(a.ampModelUserChanged)
    for _, p := range a.parameters {
        p.value.ConnectActivated2(p.vfunc)
        p.value.SetEditable(true)
        p.knob.ConnectValueChanged(p.kfunc)
    }
    a.OnStatus.ConnectClicked(a.onStatusUserChanged)
    a.ClassAAB.ConnectClicked(a.dtClassAABUserChanged)
    a.ModeTriPent.ConnectClicked(a.dtModeTriPentUserChanged)
    a.Topology.ConnectActivated2(a.dtTopologyUserChanged)
}

func (a *Amp) disconnectSignal() {
    a.AmpModel.DisconnectActivated2()
    for _, p := range a.parameters {
        p.value.DisconnectActivated2()
        p.value.SetEditable(false)
        p.knob.DisconnectValueChanged()
    }
    a.OnStatus.DisconnectClicked()
    a.ClassAAB.DisconnectClicked()
    a.ModeTriPent.DisconnectClicked()
    a.Topology.DisconnectActivated2()
}

func (a *Amp) ampModelUserChanged(fxType string) {
    a.ctrl.SetAmpType(uint32(a.id), a.AmpModel.CurrentText())
}

func (a *Amp) dtClassAABUserChanged(state bool) {
    var value string
    if state {
        value = "A/B"
    } else {
        value = "A"
    }
    if err := a.ctrl.SetDTClass(a.id, value); err != nil {
        mb := widgets.NewQMessageBox(a)
        mb.Critical(a, "An error occured", err.Error(), widgets.QMessageBox__Ok, 0)
    }
}

func (a *Amp) dtModeTriPentUserChanged(state bool) {
    var value string
    if state {
        value = "Pent"
    } else {
        value = "Tri"
    }

    if err := a.ctrl.SetDTMode(a.id, value); err != nil {
        mb := widgets.NewQMessageBox(a)
        mb.Critical(a, "An error occured", err.Error(), widgets.QMessageBox__Ok, 0)
    }
}

func (a *Amp) dtTopologyUserChanged(value string) {
    if err := a.ctrl.SetDTTopology(a.id, value); err != nil {
        mb := widgets.NewQMessageBox(a)
        mb.Critical(a, "An error occured", err.Error(), widgets.QMessageBox__Ok, 0)
    }
}

func (a *Amp) getParameter(id uint32) *Parameter{
    for i, param := range a.parameters {
        if param.id == id {
            return &a.parameters[i]
        }
    }
    return nil
}

func (a *Amp) hideDt() {
    a.DtLabel.Hide()
    a.ClassAAB.Hide()
    a.ModeTriPent.Hide()
    a.Topology.Hide()
    a.TopologyLbl.Hide()
}

func (a *Amp) onStatusUserChanged(state bool) {
    a.ctrl.SetAmpActive(uint32(a.id), state)
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

func (a *Amp) parameter0Changed2(val int) { a.parameterChanged(&a.parameters[0], strconv.Itoa(val)) }
func (a *Amp) parameter1Changed2(val int) { a.parameterChanged(&a.parameters[1], strconv.Itoa(val)) }
func (a *Amp) parameter2Changed2(val int) { a.parameterChanged(&a.parameters[2], strconv.Itoa(val)) }
func (a *Amp) parameter3Changed2(val int) { a.parameterChanged(&a.parameters[3], strconv.Itoa(val)) }
func (a *Amp) parameter4Changed2(val int) { a.parameterChanged(&a.parameters[4], strconv.Itoa(val)) }
func (a *Amp) parameter5Changed2(val int) { a.parameterChanged(&a.parameters[5], strconv.Itoa(val)) }
func (a *Amp) parameter6Changed2(val int) { a.parameterChanged(&a.parameters[6], strconv.Itoa(val)) }
func (a *Amp) parameter7Changed2(val int) { a.parameterChanged(&a.parameters[7], strconv.Itoa(val)) }
func (a *Amp) parameter8Changed2(val int) { a.parameterChanged(&a.parameters[8], strconv.Itoa(val)) }
func (a *Amp) parameter9Changed2(val int) { a.parameterChanged(&a.parameters[9], strconv.Itoa(val)) }
func (a *Amp) parameter10Changed2(val int) { a.parameterChanged(&a.parameters[10], strconv.Itoa(val)) }

func (a *Amp) parameterChanged(paramUI *Parameter, val string) {
    err := a.ctrl.SetAmpParameterValue(uint32(a.id), paramUI.id, val)
    if err != nil {
        mb := widgets.NewQMessageBox(a)
        mb.Critical(a, "An error occured", err.Error(), widgets.QMessageBox__Ok, 0)
    }
}

func (a *Amp) updateAmp(amp *pod.Amp) {
    if a.id == 0 {
        if pos, _ := amp.GetPos(); pos != 0 {
            a.parent.amps[1].Hide()
            a.parent.cabs[1].Hide()
        } else {
            a.parent.amps[1].Show()
            a.parent.cabs[1].Show()
        }
    }
    if amp.HasDt() {
        a.showDt()
        dt := amp.GetDT()
        a.ClassAAB.SetChecked(dt.GetClass() == "A/B")
        a.ModeTriPent.SetChecked(dt.GetMode() == "Pent")
        a.Topology.SetCurrentText(dt.GetTopology())
    } else {
        a.hideDt()
    }
    a.setActive(amp.GetActive())
    a.AmpModel.SetCurrentText(amp.GetName())
    for i := range a.parameters {
        a.parameters[i].id = 0
        a.parameters[i].hide()
    }
    for i, param := range amp.GetParams() {
        a.parameters[i].id = param.GetID()
        a.updateParam(param)
    }
}

func (a *Amp) updateParam(p pod.Parameter) {
    param := a.getParameter(p.GetID())
    param.setValueList([]string{p.GetValueCurrent()})
    param.setValue(p.GetValueCurrent())
    param.setLabel(p.GetName())
    param.show()

    switch p.(type) {
    case *pod.ListParam:
        param.hideKnob()
    default:
        min, max := p.GetValueRange()
        param.setValueKnob(int(p.GetValueCurrent2()), min, max)
        param.showKnob()
    }
}

func (a *Amp) setActive(status bool) {
    a.OnStatus.SetChecked(status)
}

func (a *Amp) showDt() {
    a.DtLabel.Show()
    a.ClassAAB.Show()
    a.ModeTriPent.Show()
    a.Topology.Show()
    a.TopologyLbl.Show()
}
