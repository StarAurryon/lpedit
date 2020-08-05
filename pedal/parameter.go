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
import "math"
import "strings"
import "strconv"

type Parameter interface {
    Copy() Parameter
    GetAllowedValues() []string
    GetBinValue() float32
    GetID() uint16
    GetName() string
    GetParent() PedalBoardItem
    GetValue() string
    IsAllowingOtherValues() bool
    IsNull() bool
    LockData()
    SetBinValue(float32) error
    SetParent(PedalBoardItem)
    SetValue(string) error
    UnlockData()
}

type FreqParam struct {
    name      string
    max       float32
    min       float32
    parent    PedalBoardItem
    value     float32
}

func (p *FreqParam) Copy() Parameter {
    _p := new(FreqParam)
    *_p = *p
    return _p
}

func (p *FreqParam) IsAllowingOtherValues() bool { return true }
func (p *FreqParam) IsNull() bool { return false }
func (p *FreqParam) GetAllowedValues() []string { return nil }
func (p *FreqParam) GetBinValue() float32 { return p.value }

func (p *FreqParam) GetID() (uint16) {
    _, id := p.GetParent().GetParamID(p)
    return id
}

func (p *FreqParam) GetName() string { return p.name }
func (p *FreqParam) GetParent() PedalBoardItem { return p.parent }

func (p *FreqParam) GetValue() string {
    return fmt.Sprintf("%dHz", int(math.Round(float64((p.value * (p.max - p.min)) + p.min))))
}

func (p *FreqParam) LockData() { p.parent.LockData() }

func (p *FreqParam) SetValue(s string) error {
    s = strings.Replace(s, " ", "", -1)
    s = strings.Replace(s, "Hz", "", -1)
    s = strings.Replace(s, "hz", "", -1)
    vi, err := strconv.Atoi(s)
    if err != nil {
        return err
    }
    v := float32(vi)
    if  v > p.max || p.min > v {
        return fmt.Errorf("The value must be comprised between %.1f and %.1f", p.min, p.max)
    }
    p.value = (v/(p.max - p.min)) - p.min
    return nil
}

func (p *FreqParam) SetBinValue(v float32) error {
    if v > 1 || v < 0 {
        return fmt.Errorf("The binary value must be comprised between 0 and 1")
    }
    p.value = v
    return nil
}

func (p *FreqParam) SetParent(parent PedalBoardItem) { p.parent = parent }
func (p *FreqParam) UnlockData() { p.parent.UnlockData() }

type ListParam struct {
    name  string
    list  []string
    parent PedalBoardItem
    value float32
}

func (p *ListParam) Copy() Parameter {
    _p := new(ListParam)
    *_p = *p
    return _p
}


func (p *ListParam) IsAllowingOtherValues() bool { return false }
func (p *ListParam) IsNull() bool { return false }
func (p *ListParam) GetAllowedValues() []string { return p.list }
func (p *ListParam) GetBinValue() float32 { return p.value }

func (p *ListParam) GetID() (uint16) {
    _, id := p.GetParent().GetParamID(p)
    return id
}

func (p *ListParam) GetName() string { return p.name }
func (p *ListParam) GetParent() PedalBoardItem { return p.parent }

func (p *ListParam) GetValue() string {
    return p.list[int(math.Round(float64(p.value) * float64((len(p.list) - 1))))]
}

func (p *ListParam) LockData() { p.parent.LockData() }

func (p *ListParam) SetValue(s string) error {
    found := false
    i := 0
    v := ""
    for i, v = range p.list {
        if v == s {
            found = true
            break
        }
    }
    if !found {
        return fmt.Errorf("The value must be in the list")
    }
    p.value = float32(i) / float32((len(p.list) - 1))
    return nil
}

func (p *ListParam) SetBinValue(v float32) error {
    if v > 1 || v < 0 {
        return fmt.Errorf("The binary value must be comprised between 0 and 1")
    }
    p.value = v
    return nil
}

func (p *ListParam) SetParent(parent PedalBoardItem) { p.parent = parent }
func (p *ListParam) UnlockData() { p.parent.UnlockData() }

type NullParam struct {
    parent PedalBoardItem
}

func (p *NullParam) Copy() Parameter {
    _p := new(NullParam)
    *_p = *p
    return _p
}

func (p *NullParam) IsAllowingOtherValues() bool { return false }
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

func (p *PerCentParam) IsAllowingOtherValues() bool { return true }
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
    return fmt.Sprintf("%d%%", int(math.Round(float64(p.value*100))))
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

type RangeParam struct {
    name      string
    max       float32
    min       float32
    increment float32
    parent    PedalBoardItem
    value     float32
}

func (p *RangeParam) Copy() Parameter {
    _p := new(RangeParam)
    *_p = *p
    return _p
}

func (p *RangeParam) IsAllowingOtherValues() bool { return true }
func (p *RangeParam) IsNull() bool { return false }
func (p *RangeParam) GetAllowedValues() []string { return nil }
func (p *RangeParam) GetBinValue() float32 { return p.value }

func (p *RangeParam) GetID() (uint16) {
    _, id := p.GetParent().GetParamID(p)
    return id
}

func (p *RangeParam) GetName() string { return p.name }
func (p *RangeParam) GetParent() PedalBoardItem { return p.parent }

func (p *RangeParam) GetValue() string {
    return fmt.Sprintf("%.1f", (p.value * (p.max - p.min)) + p.min)
}

func (p *RangeParam) LockData() { p.parent.LockData() }

func (p *RangeParam) SetValue(s string) error {
    s = strings.Replace(s, " ", "", -1)
    vi, err := strconv.ParseFloat(s, 32)
    if err != nil {
        return err
    }
    v := float32(vi)
    if  v > p.max || p.min > v {
        return fmt.Errorf("The value must be comprised between %.1f and %.1f", p.min, p.max)
    }
    p.value = (v/(p.max - p.min)) - p.min
    return nil
}

func (p *RangeParam) SetBinValue(v float32) error {
    if v > 1 || v < 0 {
        return fmt.Errorf("The binary value must be comprised between 0 and 1")
    }
    p.value = v
    return nil
}

func (p *RangeParam) SetParent(parent PedalBoardItem) { p.parent = parent }
func (p *RangeParam) UnlockData() { p.parent.UnlockData() }

type TempoParam struct {
    name  string
    parent PedalBoardItem
    value float32
}

func (p *TempoParam) Copy() Parameter {
    _p := new(TempoParam)
    *_p = *p
    return _p
}

func (p *TempoParam) IsAllowingOtherValues() bool { return true }
func (p *TempoParam) IsNull() bool { return false }

func (p *TempoParam) GetAllowedValues() []string {
    return []string{"Whole", "1/2 (dot)", "1/2", "1/2 (3)", "1/4 (dot)", "1/4",
        "1/4 (3)", "8th (dot)", "8th", "8th (3)", "16 (dot)", "16", "16 (3)",
        "32 (dot)", "32", "32 (3)", "64 (dot)", "64", "64 (3)"}
}

func (p *TempoParam) GetBinValue() float32 { return p.value }

func (p *TempoParam) GetID() (uint16) {
    _, id := p.GetParent().GetParamID(p)
    return id
}

func (p *TempoParam) GetName() string { return p.name }
func (p *TempoParam) GetParent() PedalBoardItem { return p.parent }

func (p *TempoParam) GetValue() string {
    if p.value > 1 {
        return p.GetAllowedValues()[int(p.value) - 2]
    }
    return fmt.Sprintf("%.2fHz", (p.value * (15 - 0.10)) + 0.10)
}

func (p *TempoParam) LockData() { p.parent.LockData() }

func (p *TempoParam) SetValue(s string) error {
    for i, _s := range p.GetAllowedValues() {
        if s == _s {
            p.value = float32(i+2)
            return nil
        }
    }

    s = strings.Replace(s, " ", "", -1)
    s = strings.Replace(s, "Hz", "", -1)
    s = strings.Replace(s, "hz", "", -1)

    vi, err := strconv.ParseFloat(s, 32)
    if err != nil {
        return err
    }
    v := float32(vi)
    if  v > 15 || 0.10 > v {
        return fmt.Errorf("The value must be comprised between 0.10 and 15 or be in the list")
    }
    p.value = (v/(15 - 0.10)) - 0.10
    return nil
}

func (p *TempoParam) SetBinValue(v float32) error {
    p.value = v
    return nil
}

func (p *TempoParam) SetParent(parent PedalBoardItem) { p.parent = parent }
func (p *TempoParam) UnlockData() { p.parent.UnlockData() }

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

func (p *TimeParam) IsAllowingOtherValues() bool { return true }
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
