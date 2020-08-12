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

/*import "github.com/StarAurryon/qt/core"
import "github.com/StarAurryon/qt/gui"
import "github.com/StarAurryon/qt/svg"*/
import "github.com/StarAurryon/qt/widgets"

import "fmt"
import "sort"

import "github.com/StarAurryon/lpedit/pedal"
import "github.com/StarAurryon/lpedit/qtctrl"

type Pedal struct {
    *PedalUI
    id         int
    pedalType  map[string][]string
    ctrl       *qtctrl.Controller
    parameters [5]Parameter
    parent     *LPEdit
}

func NewPedal(parent *LPEdit, w widgets.QWidget_ITF, c *qtctrl.Controller,
        pt map[string][]string, id int) *Pedal {
    p := &Pedal{PedalUI: NewPedalUI(w), ctrl: c, pedalType: pt, id: id}
    p.parameters[0] = Parameter{label: p.Param0Lbl, mid: p.Param0Mid,
        value: p.Param0Value, vfunc: p.parameter0Changed}
    p.parameters[1] = Parameter{label: p.Param1Lbl, mid: p.Param1Mid,
        value: p.Param1Value, vfunc: p.parameter1Changed}
    p.parameters[2] = Parameter{label: p.Param2Lbl, mid: p.Param2Mid,
        value: p.Param2Value, vfunc: p.parameter2Changed}
    p.parameters[3] = Parameter{label: p.Param3Lbl, mid: p.Param3Mid,
        value: p.Param3Value, vfunc: p.parameter3Changed}
    p.parameters[4] = Parameter{label: p.Param4Lbl, mid: p.Param4Mid,
        value: p.Param4Value, vfunc: p.parameter4Changed}
    p.parent = parent
    p.init()
    p.initUI()
    return p
}

func (p *Pedal) connectSignal() {
    p.OnStatus.ConnectClicked(p.onStatusChanged)
    p.FxModel.ConnectActivated2(p.fxModelUserChanged)
    p.FxType.ConnectActivated2(p.fxTypeUserChanged)
    p.FxType.ConnectCurrentTextChanged(p.fxTypeChanged)
    for _, param := range p.parameters {
        param.value.ConnectActivated2(param.vfunc)
    }
}

func (p *Pedal) disconnectSignal() {
    p.OnStatus.DisconnectClicked()
    p.FxModel.DisconnectActivated2()
    p.FxType.DisconnectActivated2()
    p.FxType.DisconnectCurrentTextChanged()
    for _, param := range p.parameters {
        param.value.DisconnectActivated2()
    }
}

func (p *Pedal) initUI() {
    //Setting up knob
    /*svg := svg.NewQSvgRenderer2("ui/knob.svg", p)
    pix := gui.NewQPixmap2(p.Param1Knob.SizeHint())
    paint := gui.NewQPainter()
    paint.Scale(1, 1)

    pix.Fill(gui.NewQColor2(core.Qt__transparent))
    paint.Begin(pix)
    svg.Render(paint)
    paint.End()

    pal := gui.NewQPalette()
    pal.SetBrush(gui.QPalette__Background, gui.NewQBrush7(pix))
    p.Param1Knob.SetPalette(pal)*/

    //Setting up pedal type
    keys := make([]string, 0, len(p.pedalType))

    for k := range p.pedalType {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    p.FxType.AddItems(keys)
}

func (p *Pedal) fxModelUserChanged(fxModel string) {
    p.ctrl.SetPedalType(uint32(p.id), p.FxType.CurrentText(), fxModel)
    pb := p.ctrl.GetPedalBoard()
    pb.LockData()
    p.parent.updatePedalBoardView(pb)
    pb.UnlockData()
}

func (p *Pedal) fxTypeUserChanged(fxType string) {
    p.ctrl.SetPedalType(uint32(p.id), fxType, p.FxModel.CurrentText())
    pb := p.ctrl.GetPedalBoard()
    pb.LockData()
    p.parent.updatePedalBoardView(pb)
    pb.UnlockData()
}

func (p *Pedal) fxTypeChanged(fxType string) {
    p.FxModel.Clear()
    p.FxModel.AddItems(p.pedalType[fxType])
}

func (p *Pedal) getParameter(id uint32) *Parameter{
    for i, param := range p.parameters {
        if param.id == id {
            return &p.parameters[i]
        }
    }
    return nil
}

func (p *Pedal) hideParameter(param *Parameter) {
    param.label.Hide()
    param.mid.Hide()
    param.value.Hide()
}

func (p *Pedal) parameter0Changed(v string){ p.parameterChanged(&p.parameters[0], v) }
func (p *Pedal) parameter1Changed(v string){ p.parameterChanged(&p.parameters[1], v) }
func (p *Pedal) parameter2Changed(v string){ p.parameterChanged(&p.parameters[2], v) }
func (p *Pedal) parameter3Changed(v string){ p.parameterChanged(&p.parameters[3], v) }
func (p *Pedal) parameter4Changed(v string){ p.parameterChanged(&p.parameters[4], v) }

func (p *Pedal) parameterChanged(paramUI *Parameter, v string) {
    fmt.Println(paramUI.id)
    err := p.ctrl.SetPedalParameterValue(uint32(p.id), paramUI.id, v)
    if err != nil {
        mb := widgets.NewQMessageBox(p)
        mb.Critical(p, "An error occured", err.Error(), widgets.QMessageBox__Ok, 0)
    }
    param := p.ctrl.GetPedalBoard().GetPedal2(p.id).GetParam(paramUI.id)
    param.LockData()
    p.updateParam(param)
    param.UnlockData()
}

func (p *Pedal) onStatusChanged(checked bool) {
    p.ctrl.SetPedalActive(uint32(p.id), checked)
}

func (p *Pedal) setActive(status bool) {
    p.OnStatus.SetChecked(status)
}

func (p *Pedal) setParameterLabel(param *Parameter, s string) {
    param.label.SetText(s)
}

func (p *Pedal) setParameterValueEditable(param *Parameter, editable bool) {
    param.value.SetEditable(editable)
}

func (p *Pedal) setParameterValueList(param *Parameter, s []string) {
    param.value.Clear()
    param.value.AddItems(s)
}

func (p *Pedal) setParameterValue(param *Parameter, s string) {
    param.value.SetCurrentText(s)
}

func (p *Pedal) showParameter(param *Parameter) {
    param.label.Show()
    param.mid.Show()
    param.value.Show()
}

func (pUI *Pedal) updatePedal(p *pedal.Pedal) {
    pUI.setActive(p.GetActive())
    pUI.FxType.SetCurrentText(p.GetSType())
    pUI.FxModel.SetCurrentText(p.GetName())
    for i := range pUI.parameters {
        pUI.parameters[i].id = 0
        pUI.hideParameter(&pUI.parameters[i])
    }
    for i, param := range p.GetParams() {
        pUI.parameters[i].id = param.GetID()
        pUI.updateParam(param)
    }
}

func (pUI * Pedal) updateParam(p pedal.Parameter) {
    param := pUI.getParameter(p.GetID())
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

    pUI.setParameterValueList(param, values)
    if p.IsAllowingOtherValues() {
        pUI.setParameterValueEditable(param, true)
    } else {
        pUI.setParameterValueEditable(param, false)
    }
    pUI.setParameterValue(param, p.GetValueCurrent())
    pUI.setParameterLabel(param, p.GetName())
    pUI.showParameter(param)
}
