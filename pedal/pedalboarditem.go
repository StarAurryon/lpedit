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

package pedal

import "fmt"

type PedalBoardItem interface {
    GetActive() bool
    GetID() uint32
    GetParam(id uint16) *Parameter
    GetParamLen() uint16
    GetName() string
    SetActive(bool)
    SetLastPos(uint16, uint8) error
    SetType(uint32) error
    PrintInfo()
    remove()
}

type Parameter struct {
    name  string
    value float32
}

func (p *Parameter) GetName() string {
    return p.name
}

func (p *Parameter) GetValue() string {
    if len(p.name) != 0 {
        return fmt.Sprintf("%f", p.value)
    }
    return ""
}

func (p *Parameter) SetValue(v float32) {
    if len(p.name) != 0 {
        p.value = v
    }
}
