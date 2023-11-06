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

const (
    minKnobAngle float64 = 45
    maxKnobAngle float64 = 335
)

type Parameter struct {
    id      uint32
    label   *widgets.QLabel
    mid     *widgets.QWidget
    value   *widgets.QComboBox
    knob    *widgets.QDial
    vfunc   func(string)
    kfunc   func(int)
}

func (param *Parameter) setLabel(s string) {
    if param.label != nil {
        param.label.SetText(s)
    }
}

func (param *Parameter) setValueEditable(editable bool) {
    if param.value != nil {
        param.value.SetEditable(editable)
    }
}

func (param *Parameter) setValueList(s []string) {
    if param.value != nil {
        param.value.Clear()
        param.value.AddItems(s)
    }
}

func (param *Parameter) setValue(s string) {
    if param.value != nil {
        param.value.SetCurrentText(s)
    }
}

func (param *Parameter) setValueKnob(value int, min int, max int) {
    if param.knob != nil {
        param.knob.BlockSignals(true)
        param.knob.SetMinimum(min)
        param.knob.SetMaximum(max)
        param.knob.SetValue(value)
        param.knob.BlockSignals(false)
    }
}

func (param *Parameter) hide() {
    if param.label != nil {
        param.label.Hide()
    }
    if param.mid != nil {
        param.mid.Hide()
    }
    if param.value != nil {
        param.value.Hide()
    }
}

func (param *Parameter) hideKnob() {
    if param.knob != nil {
        param.knob.Show()
    }
}

func (param *Parameter) show() {
    if param.label != nil {
        param.label.Show()
    }
    if param.mid != nil {
        param.mid.Show()
    }
    if param.value != nil {
        param.value.Show()
    }
}

func (param *Parameter) showKnob() {
    if param.knob != nil {
        param.knob.Show()
    }
}
