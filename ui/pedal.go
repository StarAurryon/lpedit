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

/*import "github.com/therecipe/qt/core"
import "github.com/therecipe/qt/gui"
import "github.com/therecipe/qt/svg"*/
import "github.com/therecipe/qt/widgets"

import "sort"

import "lpedit/pedal"
import "lpedit/qtctrl"

type Pedal struct {
    *PedalUI
    id         int
    pedalType  map[string][]string
    ctrl       *qtctrl.Controller
    labels     []*widgets.QLabel
    mids       []*widgets.QWidget
    values     []*widgets.QComboBox
}

func NewPedal(w widgets.QWidget_ITF, c *qtctrl.Controller,
        pt map[string][]string, id int) *Pedal {
    p := &Pedal{PedalUI: NewPedalUI(w), ctrl: c, pedalType: pt, id: id}
    p.labels = []*widgets.QLabel{
        p.Param0Lbl, p.Param1Lbl , p.Param2Lbl, p.Param3Lbl, p.Param4Lbl,
    }
    p.mids = []*widgets.QWidget{
        p.Param0Mid, p.Param1Mid , p.Param2Mid, p.Param3Mid, p.Param4Mid,
    }
    p.values = []*widgets.QComboBox{
        p.Param0Value, p.Param1Value , p.Param2Value, p.Param3Value, p.Param4Value,
    }
    p.init()
    p.initUI()
    return p
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

    p.FxType.ConnectCurrentTextChanged(p.fxTypeChanged)
    p.FxType.AddItems(keys)
}

func (p *Pedal) fxTypeChanged(fxType string) {
    p.FxModel.Clear()
    p.FxModel.AddItems(p.pedalType[fxType])
}

func (p *Pedal) hideParameter(id int) {
    if id < 0 || id >= len(p.labels) { return }
    p.labels[id].Hide()
    p.mids[id].Hide()
    p.values[id].Hide()
}

func (p *Pedal) setActive(status bool) {
    p.OnStatus.SetChecked(status)
}

func (p *Pedal) setParameterLabel(id int, s string) {
    if id < 0 || id >= len(p.labels) { return }
    p.labels[id].SetText(s)
}

func (p *Pedal) setParameterValueList(id int, s []string) {
    if id < 0 || id >= len(p.labels) { return }
    p.values[id].Clear()
    p.values[id].AddItems(s)
}

func (p *Pedal) setParameterValue(id int, s string) {
    if id < 0 || id >= len(p.labels) { return }
    p.values[id].SetCurrentText(s)
}

func (p *Pedal) showParameter(id int) {
    if id < 0 || id >= len(p.labels) { return }
    p.labels[id].Show()
    p.mids[id].Show()
    p.values[id].Show()
}

func (pUI *Pedal) updatePedal(p *pedal.Pedal) {
    pUI.setActive(p.GetActive())
    pUI.FxType.SetCurrentText(p.GetSType())
    pUI.FxModel.SetCurrentText(p.GetName())
    for i := range pUI.labels {
        pUI.hideParameter(i)
    }
    for _, param := range p.GetParams() {
        pUI.updateParam(param)
    }
}

func (pUI * Pedal) updateParam(p pedal.Parameter) {
    id := int(p.GetID())
    if id == 0 || id > 5 { return }
    id--
    if !p.IsNull() {
        pUI.showParameter(id)
        pUI.setParameterLabel(id, p.GetName())
        pUI.setParameterValueList(id, []string{p.GetValue()})
        pUI.setParameterValue(id, p.GetValue())
    } else {
        pUI.hideParameter(id)
    }
}
