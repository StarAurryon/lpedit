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

import "github.com/StarAurryon/qt/widgets"

import "sort"

import "github.com/StarAurryon/lpedit/pedal"
import "github.com/StarAurryon/lpedit/qtctrl"

type Cab struct {
    *CabUI
    id         int
    cabType    []string
    ctrl       *qtctrl.Controller
    parameters [6]Parameter
    parent     *LPEdit
}

func NewCab(parent *LPEdit, w widgets.QWidget_ITF, ctrl *qtctrl.Controller,
        ct []string, id int) *Cab {
    c := &Cab{CabUI: NewCabUI(w), ctrl: ctrl, cabType: ct, id: id}
    c.parameters[0] = Parameter{label: c.Param0Lbl, mid: c.Param0Mid,
        value: c.Param0Value, vfunc: c.parameter0Changed}
    c.parameters[1] = Parameter{label: c.Param1Lbl, mid: c.Param1Mid,
        value: c.Param1Value, vfunc: c.parameter1Changed}
    c.parameters[2] = Parameter{label: c.Param2Lbl, mid: c.Param2Mid,
        value: c.Param2Value, vfunc: c.parameter2Changed}
    c.parameters[3] = Parameter{label: c.Param3Lbl, mid: c.Param3Mid,
        value: c.Param3Value, vfunc: c.parameter3Changed}
    c.parameters[4] = Parameter{label: c.Param4Lbl, mid: c.Param4Mid,
        value: c.Param4Value, vfunc: c.parameter4Changed}
    c.parameters[5] = Parameter{label: c.MicModelLbl, mid: nil,
        value: c.MicModel, vfunc: c.parameter5Changed}
    c.parent = parent
    c.init()
    c.initUI()
    return c
}

func (c *Cab) connectSignal() {
    c.CabModel.ConnectActivated2(c.cabModelUserChanged)
    for _, p := range c.parameters {
        p.value.ConnectActivated2(p.vfunc)
        p.value.SetEditable(true)
    }
}

func (c *Cab) disconnectSignal() {
    c.CabModel.DisconnectActivated2()
    for _, p := range c.parameters {
        p.value.DisconnectActivated2()
        p.value.SetEditable(false)
    }
}

func (c *Cab) initUI() {
    keys := make([]string, 0, len(c.cabType))

    for _, k := range c.cabType {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    c.CabModel.AddItems(keys)
}

func (c *Cab) cabModelUserChanged(value string) {}

func (c *Cab) getParameter(id uint32) *Parameter{
    for i, param := range c.parameters {
        if param.id == id {
            return &c.parameters[i]
        }
    }
    return nil
}

func (c *Cab) parameter0Changed(val string) { c.parameterChanged(&c.parameters[0], val) }
func (c *Cab) parameter1Changed(val string) { c.parameterChanged(&c.parameters[1], val) }
func (c *Cab) parameter2Changed(val string) { c.parameterChanged(&c.parameters[2], val) }
func (c *Cab) parameter3Changed(val string) { c.parameterChanged(&c.parameters[3], val) }
func (c *Cab) parameter4Changed(val string) { c.parameterChanged(&c.parameters[4], val) }
func (c *Cab) parameter5Changed(val string) { c.parameterChanged(&c.parameters[5], val) }

func (c *Cab) parameterChanged(param *Parameter, val string) {
    c.ctrl.SetCabParameterValue(uint32(c.id), param.id, val)
}

func (c *Cab) updateCab(cab *pedal.Cab) {
    c.CabModel.SetCurrentText(cab.GetName())
    if cab.GetHideParams() {
        for i := range c.parameters {
            c.parameters[i].hide()
        }
    } else {
        for i, param := range cab.GetParams() {
            c.parameters[i].id = param.GetID()
            c.updateParam(param)
        }
    }
}

func (c *Cab) updateParam(p pedal.Parameter) {
    param := c.getParameter(p.GetID())
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
    param.setLabel(p.GetName())
    param.show()
}
