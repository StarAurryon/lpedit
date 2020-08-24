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

import "github.com/StarAurryon/qt/gui"
import "github.com/StarAurryon/qt/widgets"

import "os"

type Parameter struct {
    id    uint32
    label *widgets.QLabel
    mid   *widgets.QWidget
    value *widgets.QComboBox
    knob  *widgets.QLabel
    vfunc func(string)
}

func (param *Parameter) setLabel(s string) {
    if param.label != nil {
        param.label.SetText(s)
    }
}

func (param *Parameter) setupKnob() {
    if param.knob != nil {
        ps := string(os.PathSeparator)
        iconPath := "ui" + ps + "knob.png"
        pixmap := gui.NewQPixmap3(iconPath, "", 0)
        param.knob.SetScaledContents(true)
        param.knob.SetPixmap(pixmap)
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

func (param *Parameter) setValue( s string) {
    if param.value != nil {
        param.value.SetCurrentText(s)
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
