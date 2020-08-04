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
    GetActive2() uint32
    GetID() uint32
    GetParam(id uint16) Parameter
    GetParams() []Parameter
    GetParamID(Parameter) (error, uint16)
    GetParamLen() uint16
    GetName() string
    GetType() uint32
    GetPos() (uint16, uint8)
    LockData()
    SetActive(bool)
    SetPos(uint16, uint8)
    SetType(uint32) error
    SetType2(string, string)
    UnlockData()
    LogInfo()
}

type SortablePosPBI []PedalBoardItem

func (s SortablePosPBI) Len() int           { return len(s) }

func (s SortablePosPBI) Less(i, j int) bool {
    posI, _ := s[i].GetPos()
    posJ, _ := s[j].GetPos()
    return posI < posJ
}

func (s SortablePosPBI) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type Parameter interface {
    Copy() Parameter
    GetAllowedValues() []string
    GetBinValue() float32
    GetID() uint16
    GetName() string
    GetParent() PedalBoardItem
    GetValue() string
    IsNull() bool
    LockData()
    SetBinValue(float32) error
    SetParent(PedalBoardItem)
    SetValue(string) error
    UnlockData()
}

type NullParam struct {
    parent PedalBoardItem
}

func (p *NullParam) Copy() Parameter {
    _p := new(NullParam)
    *_p = *p
    return _p
}

func (p *NullParam) IsNull() bool { return true }
func (p *NullParam) GetAllowedValues() []string { return nil }
func (p *NullParam) GetBinValue() float32 { return 0 }

func (p *NullParam) GetID() uint16 {
    _, id := p.GetParent().GetParamID(p)
    return id
}

func (p *NullParam) GetName() string { return "Null" }
func (p *NullParam) GetParent() PedalBoardItem { return p.parent }
func (p *NullParam) GetValue() string { return "" }
func (p *NullParam) LockData() { p.parent.LockData() }
func (p *NullParam) SetValue(string) error { return fmt.Errorf("Null parameter") }
func (p *NullParam) SetBinValue(float32) error { return fmt.Errorf("Null parameter") }
func (p *NullParam) SetParent(parent PedalBoardItem) { p.parent = parent }
func (p *NullParam) UnlockData() { p.parent.UnlockData() }

type PerCentParam struct {
    name  string
    parent PedalBoardItem
    value float32
}

func (p *PerCentParam) Copy() Parameter {
    _p := new(PerCentParam)
    *_p = *p
    return _p
}

func (p *PerCentParam) IsNull() bool { return false }
func (p *PerCentParam) GetAllowedValues() []string { return nil }
func (p *PerCentParam) GetBinValue() float32 { return p.value }

func (p *PerCentParam) GetID() uint16 {
    _, id := p.GetParent().GetParamID(p)
    return id
}

func (p *PerCentParam) GetName() string { return p.name }
func (p *PerCentParam) GetParent() PedalBoardItem { return p.parent }

func (p *PerCentParam) GetValue() string {
    return fmt.Sprintf("%d%%", int(p.value*100))
}

func (p *PerCentParam) LockData() { p.parent.LockData() }

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

func (p *PerCentParam) SetBinValue(v float32) error {
    if v > 1 || v < 0 {
        return fmt.Errorf("The binary value must be comprised between 0 and 1")
    }
    p.value = v
    return nil
}

func (p *PerCentParam) SetParent(parent PedalBoardItem) { p.parent = parent }
func (p *PerCentParam) UnlockData() { p.parent.UnlockData() }

type TimeParam struct {
    name  string
    maxMs int
    parent PedalBoardItem
    value float32
}

func (p *TimeParam) Copy() Parameter {
    _p := new(TimeParam)
    *_p = *p
    return _p
}

func (p *TimeParam) IsNull() bool { return false }
func (p *TimeParam) GetAllowedValues() []string { return nil }
func (p *TimeParam) GetBinValue() float32 { return p.value }

func (p *TimeParam) GetID() (uint16) {
    _, id := p.GetParent().GetParamID(p)
    return id
}

func (p *TimeParam) GetName() string { return p.name }
func (p *TimeParam) GetParent() PedalBoardItem { return p.parent }

func (p *TimeParam) GetValue() string {
    return fmt.Sprintf("%dms", int(p.value*float32(p.maxMs)))
}

func (p *TimeParam) LockData() { p.parent.LockData() }

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

func (p *TimeParam) SetBinValue(v float32) error {
    if v > 1 || v < 0 {
        return fmt.Errorf("The binary value must be comprised between 0 and 1")
    }
    p.value = v
    return nil
}

func (p *TimeParam) SetParent(parent PedalBoardItem) { p.parent = parent }
func (p *TimeParam) UnlockData() { p.parent.UnlockData() }
