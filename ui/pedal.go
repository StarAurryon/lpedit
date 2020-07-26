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
import "github.com/therecipe/qt/svg"
import "github.com/therecipe/qt/widgets"

import "sort"

import "lpedit/qtctrl"

type Pedal struct {
    *PedalUI
    id         int
    pedalType  map[string][]string
    ctrl       *qtctrl.Controller
}

func NewPedal(w widgets.QWidget_ITF, c *qtctrl.Controller,
        pt map[string][]string, id int) *Pedal {
    p := &Pedal{PedalUI: NewPedalUI(w), ctrl: c, pedalType: pt, id: id}
    p.init()
    p.initUI()
    return p
}

func (p *Pedal) initUI() {
    //Setting up knob
    svg := svg.NewQSvgRenderer2("ui/knob.svg", p)
    pix := gui.NewQPixmap2(p.Param1Knob.SizeHint())
    paint := gui.NewQPainter()
    paint.Scale(1, 1)

    pix.Fill(gui.NewQColor2(core.Qt__transparent))
    paint.Begin(pix)
    svg.Render(paint)
    paint.End()

    pal := gui.NewQPalette()
    pal.SetBrush(gui.QPalette__Background, gui.NewQBrush7(pix))
    p.Param1Knob.SetPalette(pal)

    //Setting up pedal type
    keys := make([]string, 0, len(p.pedalType))

    for k := range p.pedalType {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    p.FxType.ConnectCurrentTextChanged(p.fxTypeChanged)
    p.FxType.AddItems(keys)
}

func (p *Pedal) fxTypeChanged(fxType string) {
    p.FxModel.Clear()
    p.FxModel.AddItems(p.pedalType[fxType])
}

func (p *Pedal) hideParameter(id int) {
    switch id {
    case 0:
        p.Param0Lbl.Hide()
        p.Param0Mid.Hide()
        p.Param0Value.Hide()
    case 1:
        p.Param1Lbl.Hide()
        p.Param1Mid.Hide()
        p.Param1Value.Hide()
    case 2:
        p.Param2Lbl.Hide()
        p.Param2Mid.Hide()
        p.Param2Value.Hide()
    case 3:
        p.Param3Lbl.Hide()
        p.Param3Mid.Hide()
        p.Param3Value.Hide()
    case 4:
        p.Param4Lbl.Hide()
        p.Param4Mid.Hide()
        p.Param4Value.Hide()
    }
}

func (p *Pedal) setActive(status bool) {
    p.OnStatus.SetChecked(status)
}

func (p *Pedal) setParameterLabel(id int, s string) {
    switch id {
    case 0:
        p.Param0Lbl.SetText(s)
    case 1:
        p.Param1Lbl.SetText(s)
    case 2:
        p.Param2Lbl.SetText(s)
    case 3:
        p.Param3Lbl.SetText(s)
    case 4:
        p.Param4Lbl.SetText(s)
    }
}

func (p *Pedal) setParameterValueList(id int, s []string) {
    switch id {
    case 0:
        p.Param0Value.Clear()
        p.Param0Value.AddItems(s)
    case 1:
        p.Param1Value.Clear()
        p.Param1Value.AddItems(s)
    case 2:
        p.Param2Value.Clear()
        p.Param2Value.AddItems(s)
    case 3:
        p.Param4Value.Clear()
        p.Param3Value.AddItems(s)
    case 4:
        p.Param4Value.Clear()
        p.Param4Value.AddItems(s)
    }
}

func (p *Pedal) setParameterValue(id int, s string) {
    switch id {
    case 0:
        p.Param0Value.SetCurrentText(s)
    case 1:
        p.Param1Value.SetCurrentText(s)
    case 2:
        p.Param2Value.SetCurrentText(s)
    case 3:
        p.Param3Value.SetCurrentText(s)
    case 4:
        p.Param4Value.SetCurrentText(s)
    }
}

func (p *Pedal) showParameter(id int) {
    switch id {
    case 0:
        p.Param0Lbl.Show()
        p.Param0Mid.Show()
        p.Param0Value.Show()
    case 1:
        p.Param1Lbl.Show()
        p.Param1Mid.Show()
        p.Param1Value.Show()
    case 2:
        p.Param2Lbl.Show()
        p.Param2Mid.Show()
        p.Param2Value.Show()
    case 3:
        p.Param3Lbl.Show()
        p.Param3Mid.Show()
        p.Param3Value.Show()
    case 4:
        p.Param4Lbl.Show()
        p.Param4Mid.Show()
        p.Param4Value.Show()
    }
}

func (pUI *Pedal) updateModel() {
    p := pUI.ctrl.GetPedal(pUI.id)
    pUI.setActive(p.GetActive())
    pUI.FxType.SetCurrentText(p.GetPType())
    pUI.FxModel.SetCurrentText(p.GetName())
    nparam := int(p.GetParamLen())
    for j := 0; j < 5; j++ {
        if j < (nparam - 1) {
            param := p.GetParam(uint16(j+1))
            pname := param.GetName()
            if pname != "" {
                pUI.showParameter(j)
                pUI.setParameterLabel(j, pname)
                pUI.setParameterValueList(j, []string{param.GetValue()})
                pUI.setParameterValue(j, param.GetValue())
            } else {
                pUI.hideParameter(j)
            }
        } else {
            pUI.hideParameter(j)
        }
    }
}
