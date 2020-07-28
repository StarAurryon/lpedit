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
import "strings"
import "strconv"

type PedalBoardItem interface {
    GetActive() bool
    GetID() uint32
    GetParam(id uint16) Parameter
    GetParamLen() uint16
    GetName() string
    SetActive(bool)
    SetLastPos(uint16, uint8) error
    SetType(uint32) error
    LogInfo()
    remove()
}

type Parameter interface {
    Copy() Parameter
    IsNull() bool
    GetName() string
    GetValue() string
    SetValue(string) error
    GetBinValue() float32
    SetBinValue(float32) error
}

type NullParam struct {}

func (p *NullParam) Copy() Parameter {
    _p := new(NullParam)
    *_p = *p
    return _p
}

func (p *NullParam) IsNull() bool { return true }
func (p *NullParam) GetName() string { return "Null" }
func (p *NullParam) GetValue() string { return "" }
func (p *NullParam) SetValue(string) error { return fmt.Errorf("Null parameter") }
func (p *NullParam) GetBinValue() float32 { return 0 }
func (p *NullParam) SetBinValue(float32) error { return fmt.Errorf("Null parameter") }

type PerCentParam struct {
    name  string
    value float32
}

func (p *PerCentParam) Copy() Parameter {
    _p := new(PerCentParam)
    *_p = *p
    return _p
}

func (p *PerCentParam) IsNull() bool { return false }
func (p *PerCentParam) GetName() string { return p.name }

func (p *PerCentParam) GetValue() string {
    return fmt.Sprintf("%d%%", int(p.value*100))
}

func (p *PerCentParam) SetValue(s string) error {
    s = strings.Replace(s, " ", "", -1)
    s = strings.Replace(s, "%", "", -1)
    vi, err := strconv.Atoi(s)
    if err != nil {
        return err
    }
    if vi > 100 || vi < 0 {
        return fmt.Errorf("The value must be comprised between 0 and 100")
    }
    p.value = float32(vi)/100
    return nil
}

func (p *PerCentParam) GetBinValue() float32 {
    return p.value
}

func (p *PerCentParam) SetBinValue(v float32) error {
    if v > 1 || v < 0 {
        return fmt.Errorf("The binary value must be comprised between 0 and 1")
    }
    p.value = v
    return nil
}

type TimeParam struct {
    name  string
    maxMs int
    value float32
}

func (p *TimeParam) Copy() Parameter {
    _p := new(TimeParam)
    *_p = *p
    return _p
}

func (p *TimeParam) IsNull() bool { return false }
func (p *TimeParam) GetName() string { return p.name }

func (p *TimeParam) GetValue() string {
    return fmt.Sprintf("%dms", int(p.value*float32(p.maxMs)))
}

func (p *TimeParam) SetValue(s string) error {
    s = strings.Replace(s, " ", "", -1)
    s = strings.Replace(s, "ms", "", -1)
    vi, err := strconv.Atoi(s)
    if err != nil {
        return err
    }
    if  vi > p.maxMs || vi < 0 {
        return fmt.Errorf("The value must be comprised between 0 and %d", p.maxMs)
    }
    p.value = float32(vi)/float32(p.maxMs)
    return nil
}

func (p *TimeParam) GetBinValue() float32 {
    return p.value
}

func (p *TimeParam) SetBinValue(v float32) error {
    if v > 1 || v < 0 {
        return fmt.Errorf("The binary value must be comprised between 0 and 1")
    }
    p.value = v
    return nil
}
