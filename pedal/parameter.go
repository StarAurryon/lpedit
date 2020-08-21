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

import "bytes"
import "encoding/binary"
import "fmt"
import "math"
import "strings"
import "strconv"

const (
    Int32Type  uint32 = 0
    float32Type uint32 = 1
)

type Parameter interface {
    Copy() Parameter
    GetAllowedValues() []string
    GetBinValueCurrent() [4]byte
    GetBinValueMin() [4]byte
    GetBinValueMax() [4]byte
    GetBinValueType() uint32
    GetID() uint32
    GetName() string
    GetParent() LPEObject
    GetValueCurrent() string
    GetValueMin() string
    GetValueMax() string
    IsAllowingOtherValues() bool
    LockData()
    SetBinValueCurrent([4]byte) error
    SetBinValueMin([4]byte) error
    SetBinValueMax([4]byte) error
    SetParent(PedalBoardItem)
    SetValueCurrent(string) error
    SetValueMin(string) error
    SetValueMax(string) error
    UnlockData()
}

type GenericParameter struct {
    id        uint32
    name      string
    parent    LPEObject
}

func (p *GenericParameter) GetID() uint32 { return p.id }
func (p *GenericParameter) GetName() string { return p.name }
func (p *GenericParameter) GetParent() LPEObject { return p.parent }
func (p *GenericParameter) LockData() { p.parent.LockData() }
func (p *GenericParameter) SetParent(parent PedalBoardItem) { p.parent = parent }
func (p *GenericParameter) UnlockData() { p.parent.UnlockData() }

func to4Bytes(obj interface{}) [4]byte {
    ret := [4]byte{}
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, obj)
    copy(ret[:], buf.Bytes())
    return ret
}

func from4BytesToFloat32(v [4]byte) float32 {
    var ret float32
    err := binary.Read(bytes.NewReader(v[:]), binary.LittleEndian, &ret)
    if err != nil {
        return 0
    }
    return ret
}

func from4BytesToInt32(v [4]byte) int32 {
    var ret int32
    err := binary.Read(bytes.NewReader(v[:]), binary.LittleEndian, &ret)
    if err != nil {
        return 0
    }
    return ret
}

type FreqParam struct {
    GenericParameter
    max          float32
    min          float32
    valueCurrent float32
    valueMin     float32
    valueMax     float32
}

func (p *FreqParam) Copy() Parameter {
    _p := new(FreqParam)
    *_p = *p
    return _p
}

func (p *FreqParam) IsAllowingOtherValues() bool { return true }
func (p *FreqParam) GetAllowedValues() []string { return nil }
func (p *FreqParam) GetBinValueCurrent() [4]byte { return to4Bytes(p.valueCurrent) }
func (p *FreqParam) GetBinValueMin() [4]byte { return to4Bytes(p.valueMin) }
func (p *FreqParam) GetBinValueMax() [4]byte { return to4Bytes(p.valueMax) }
func (p *FreqParam) GetBinValueType() uint32 { return float32Type }

func (p *FreqParam) getValue(v float32) string {
    return fmt.Sprintf("%dHz", int(math.Round(float64((v * (p.max - p.min)) + p.min))))
}

func (p *FreqParam) GetValueCurrent() string { return p.getValue(p.valueCurrent) }
func (p *FreqParam) GetValueMin() string { return p.getValue(p.valueMin) }
func (p *FreqParam) GetValueMax() string { return p.getValue(p.valueMax) }

func (p *FreqParam) setBinValue(dst *float32, value [4]byte) error {
    v := from4BytesToFloat32(value)
    if v > 1 || v < 0 {
        return fmt.Errorf("The binary value must be comprised between 0 and 1")
    }
    *dst = v
    return nil
}

func (p *FreqParam) SetBinValueCurrent(value [4]byte) error { return p.setBinValue(&p.valueCurrent, value) }
func (p *FreqParam) SetBinValueMin(value [4]byte) error { return p.setBinValue(&p.valueMin, value) }
func (p *FreqParam) SetBinValueMax(value [4]byte) error { return p.setBinValue(&p.valueMax, value) }

func (p *FreqParam) setValue(dst *float32, s string) error {
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
    *dst = (v - p.min) / (p.max - p.min)
    return nil
}

func (p *FreqParam) SetValueCurrent(s string) error { return p.setValue(&p.valueCurrent, s) }
func (p *FreqParam) SetValueMin(s string) error { return p.setValue(&p.valueMin, s) }
func (p *FreqParam) SetValueMax(s string) error { return p.setValue(&p.valueMax, s) }

type FreqKParam struct {
    GenericParameter
    max          float32
    min          float32
    valueCurrent float32
    valueMin     float32
    valueMax     float32
}

func (p *FreqKParam) Copy() Parameter {
    _p := new(FreqKParam)
    *_p = *p
    return _p
}

func (p *FreqKParam) IsAllowingOtherValues() bool { return true }
func (p *FreqKParam) GetAllowedValues() []string { return nil }
func (p *FreqKParam) GetBinValueCurrent() [4]byte { return to4Bytes(p.valueCurrent) }
func (p *FreqKParam) GetBinValueMin() [4]byte { return to4Bytes(p.valueMin) }
func (p *FreqKParam) GetBinValueMax() [4]byte { return to4Bytes(p.valueMax) }
func (p *FreqKParam) GetBinValueType() uint32 { return float32Type }

func (p *FreqKParam) getValue(v float32) string {
    return fmt.Sprintf("%.1fKHz", float64((v * (p.max - p.min)) + p.min))
}

func (p *FreqKParam) GetValueCurrent() string { return p.getValue(p.valueCurrent) }
func (p *FreqKParam) GetValueMin() string { return p.getValue(p.valueMin) }
func (p *FreqKParam) GetValueMax() string { return p.getValue(p.valueMax) }

func (p *FreqKParam) setBinValue(dst *float32, value [4]byte) error {
    v := from4BytesToFloat32(value)
    if v > 1 || v < 0 {
        return fmt.Errorf("The binary value must be comprised between 0 and 1")
    }
    *dst = v
    return nil
}

func (p *FreqKParam) SetBinValueCurrent(value [4]byte) error { return p.setBinValue(&p.valueCurrent, value) }
func (p *FreqKParam) SetBinValueMin(value [4]byte) error { return p.setBinValue(&p.valueMin, value) }
func (p *FreqKParam) SetBinValueMax(value [4]byte) error { return p.setBinValue(&p.valueMax, value) }

func (p *FreqKParam) setValue(dst *float32, s string) error {
    s = strings.Replace(s, " ", "", -1)
    s = strings.Replace(s, "KHz", "", -1)
    s = strings.Replace(s, "khz", "", -1)
    vi, err := strconv.ParseFloat(s, 32)
    if err != nil {
        return err
    }
    v := float32(vi)
    if  v > p.max || p.min > v {
        return fmt.Errorf("The value must be comprised between %.1f and %.1f", p.min, p.max)
    }
    *dst = (v - p.min) / (p.max - p.min)
    return nil
}

func (p *FreqKParam) SetValueCurrent(s string) error { return p.setValue(&p.valueCurrent, s) }
func (p *FreqKParam) SetValueMin(s string) error { return p.setValue(&p.valueMin, s) }
func (p *FreqKParam) SetValueMax(s string) error { return p.setValue(&p.valueMax, s) }

type ListParam struct {
    GenericParameter
    list         []string
    valueCurrent interface{}
    valueMin     interface{}
    valueMax     interface{}
    binValueType uint32
    maxIDShift   int
}

func (p *ListParam) Copy() Parameter {
    _p := new(ListParam)
    *_p = *p
    return _p
}

func (p *ListParam) IsAllowingOtherValues() bool { return false }

func (p *ListParam) GetAllowedValues() []string {
    ret := []string{}
    for _, s := range p.list {
        if s != "" {
            ret = append(ret, s)
        }
    }
    return ret
}

func (p *ListParam) GetBinValueCurrent() [4]byte { return to4Bytes(p.valueCurrent) }
func (p *ListParam) GetBinValueMin() [4]byte { return to4Bytes(p.valueMin) }
func (p *ListParam) GetBinValueMax() [4]byte { return to4Bytes(p.valueMax) }
func (p *ListParam) GetBinValueType() uint32 { return p.binValueType }

func (p *ListParam) getValue(v interface{}) string {
    if v == nil {
        return p.list[0]
    }
    switch p.binValueType {
    case Int32Type:
        value := v.(int32)
        return p.list[value - int32(p.maxIDShift)]
    case float32Type:
        value := v.(float32)
        return p.list[int(math.Round(float64(value) * float64((len(p.list) - 1 + p.maxIDShift))))]
    default:
        return ""
    }
}

func (p *ListParam) GetValueCurrent() string { return p.getValue(p.valueCurrent) }
func (p *ListParam) GetValueMin() string { return p.getValue(p.valueMin) }
func (p *ListParam) GetValueMax() string { return p.getValue(p.valueMax) }

func (p *ListParam) setBinValue(dst *interface{}, value [4]byte) error {
    switch p.binValueType {
    case Int32Type:
        v := from4BytesToInt32(value)
        max := (len(p.list) + p.maxIDShift) - 1
        if v > int32(max) || v < int32(p.maxIDShift) {
            return fmt.Errorf("The binary value must be comprised between %d and %d", p.maxIDShift,  max)
        }
        if p.list[v - int32(p.maxIDShift)] == "" {
            return fmt.Errorf("The binary value is not valid %d", v)
        }
        *dst = v
    case float32Type:
        v := from4BytesToFloat32(value)
        if v > 1 || v < 0 {
            return fmt.Errorf("The binary value must be comprised between 0 and 1")
        }
        *dst = v
    }
    return nil
}

func (p *ListParam) SetBinValueCurrent(value [4]byte) error { return p.setBinValue(&p.valueCurrent, value) }
func (p *ListParam) SetBinValueMin(value [4]byte) error { return p.setBinValue(&p.valueMin, value) }
func (p *ListParam) SetBinValueMax(value [4]byte) error { return p.setBinValue(&p.valueMax, value) }

func (p *ListParam) setValue(dst *interface{}, s string) error {
    if s == "" {
        return fmt.Errorf("The value must not be empty")
    }

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
    switch p.binValueType {
    case Int32Type:
        *dst = int32(i + p.maxIDShift)
    case float32Type:
        *dst = float32(i) / float32((len(p.list) - 1 + p.maxIDShift))
    }
    return nil
}

func (p *ListParam) SetValueCurrent(s string) error { return p.setValue(&p.valueCurrent, s) }
func (p *ListParam) SetValueMin(s string) error { return p.setValue(&p.valueMin, s) }
func (p *ListParam) SetValueMax(s string) error { return p.setValue(&p.valueMax, s) }

type PerCentParam struct {
    GenericParameter
    valueCurrent  float32
    valueMin  float32
    valueMax  float32
}

func (p *PerCentParam) Copy() Parameter {
    _p := new(PerCentParam)
    *_p = *p
    return _p
}

func (p *PerCentParam) IsAllowingOtherValues() bool { return true }
func (p *PerCentParam) GetAllowedValues() []string { return nil }
func (p *PerCentParam) GetBinValueCurrent() [4]byte { return to4Bytes(p.valueCurrent) }
func (p *PerCentParam) GetBinValueMin() [4]byte { return to4Bytes(p.valueMin) }
func (p *PerCentParam) GetBinValueMax() [4]byte { return to4Bytes(p.valueMax) }
func (p *PerCentParam) GetBinValueType() uint32 { return float32Type }

func (p *PerCentParam) getValue(v float32) string {
    return fmt.Sprintf("%d%%", int(math.Round(float64(v*100))))
}

func (p *PerCentParam) GetValueCurrent() string { return p.getValue(p.valueCurrent) }
func (p *PerCentParam) GetValueMin() string { return p.getValue(p.valueMin) }
func (p *PerCentParam) GetValueMax() string { return p.getValue(p.valueMax) }

func (p *PerCentParam) setBinValue(dst *float32, value [4]byte) error {
    v := from4BytesToFloat32(value)
    if v > 1 || v < 0 {
        return fmt.Errorf("The binary value must be comprised between 0 and 1")
    }
    *dst = v
    return nil
}

func (p *PerCentParam) SetBinValueCurrent(value [4]byte) error { return p.setBinValue(&p.valueCurrent, value) }
func (p *PerCentParam) SetBinValueMin(value [4]byte) error { return p.setBinValue(&p.valueMin, value) }
func (p *PerCentParam) SetBinValueMax(value [4]byte) error { return p.setBinValue(&p.valueMax, value) }

func (p *PerCentParam) setValue(dst *float32, s string) error {
    s = strings.Replace(s, " ", "", -1)
    s = strings.Replace(s, "%", "", -1)
    vi, err := strconv.Atoi(s)
    if err != nil {
        return err
    }
    if vi > 100 || vi < 0 {
        return fmt.Errorf("The value must be comprised between 0 and 100")
    }
    *dst = float32(vi)/100
    return nil
}

func (p *PerCentParam) SetValueCurrent(s string) error { return p.setValue(&p.valueCurrent, s) }
func (p *PerCentParam) SetValueMin(s string) error { return p.setValue(&p.valueMin, s) }
func (p *PerCentParam) SetValueMax(s string) error { return p.setValue(&p.valueMax, s) }

type RangeParam struct {
    GenericParameter
    max          float32
    min          float32
    valueCurrent float32
    valueMin     float32
    valueMax     float32
}

func (p *RangeParam) Copy() Parameter {
    _p := new(RangeParam)
    *_p = *p
    return _p
}

func (p *RangeParam) IsAllowingOtherValues() bool { return true }
func (p *RangeParam) GetAllowedValues() []string { return nil }
func (p *RangeParam) GetBinValueCurrent() [4]byte { return to4Bytes(p.valueCurrent) }
func (p *RangeParam) GetBinValueMin() [4]byte { return to4Bytes(p.valueMin) }
func (p *RangeParam) GetBinValueMax() [4]byte { return to4Bytes(p.valueMax) }
func (p *RangeParam) GetBinValueType() uint32 { return float32Type }

func (p *RangeParam) getValue(v float32) string {
    return fmt.Sprintf("%.1f", (v * (p.max - p.min)) + p.min)
}

func (p *RangeParam) GetValueCurrent() string { return p.getValue(p.valueCurrent) }
func (p *RangeParam) GetValueMin() string { return p.getValue(p.valueMin) }
func (p *RangeParam) GetValueMax() string { return p.getValue(p.valueMax) }

func (p *RangeParam) setBinValue(dst *float32, value [4]byte) error {
    v := from4BytesToFloat32(value)
    if v > 1 || v < 0 {
        return fmt.Errorf("The binary value must be comprised between 0 and 1")
    }
    *dst = v
    return nil
}

func (p *RangeParam) SetBinValueCurrent(value [4]byte) error { return p.setBinValue(&p.valueCurrent, value) }
func (p *RangeParam) SetBinValueMin(value [4]byte) error { return p.setBinValue(&p.valueMin, value) }
func (p *RangeParam) SetBinValueMax(value [4]byte) error { return p.setBinValue(&p.valueMax, value) }

func (p *RangeParam) setValue(dst *float32, s string) error {
    s = strings.Replace(s, " ", "", -1)
    vi, err := strconv.ParseFloat(s, 32)
    if err != nil {
        return err
    }
    v := float32(vi)
    if  v > p.max || p.min > v {
        return fmt.Errorf("The value must be comprised between %.1f and %.1f", p.min, p.max)
    }
    *dst = (v - p.min)/(p.max - p.min)
    return nil
}

func (p *RangeParam) SetValueCurrent(s string) error { return p.setValue(&p.valueCurrent, s) }
func (p *RangeParam) SetValueMin(s string) error { return p.setValue(&p.valueMin, s) }
func (p *RangeParam) SetValueMax(s string) error { return p.setValue(&p.valueMax, s) }

type TempoParam struct {
    GenericParameter
    max          float32
    min          float32
    valueCurrent float32
    valueMin     float32
    valueMax     float32
}

func (p *TempoParam) Copy() Parameter {
    _p := new(TempoParam)
    *_p = *p
    return _p
}

func (p *TempoParam) IsAllowingOtherValues() bool { return true }

func (p *TempoParam) GetAllowedValues() []string {
    return []string{"Whole", "1/2 (dot)", "1/2", "1/2 (3)", "1/4 (dot)", "1/4",
        "1/4 (3)", "8th (dot)", "8th", "8th (3)", "16 (dot)", "16", "16 (3)",
        "32 (dot)", "32", "32 (3)", "64 (dot)", "64", "64 (3)"}
}

func (p *TempoParam) GetBinValueCurrent() [4]byte { return to4Bytes(p.valueCurrent) }
func (p *TempoParam) GetBinValueMin() [4]byte { return to4Bytes(p.valueMin) }
func (p *TempoParam) GetBinValueMax() [4]byte { return to4Bytes(p.valueMax) }
func (p *TempoParam) GetBinValueType() uint32 { return float32Type }

func (p *TempoParam) getValue(v float32) string {
    if v > 1 {
        return p.GetAllowedValues()[int(v) - 2]
    }
    return fmt.Sprintf("%.2fHz", (v * (p.max - p.min)) + p.min)
}

func (p *TempoParam) GetValueCurrent() string { return p.getValue(p.valueCurrent) }
func (p *TempoParam) GetValueMin() string { return p.getValue(p.valueMin) }
func (p *TempoParam) GetValueMax() string { return p.getValue(p.valueMax) }

func (p *TempoParam) setBinValue(dst *float32, value [4]byte) error {
    v := from4BytesToFloat32(value)
    maxV := float32(len(p.GetAllowedValues()))
    if v > maxV || v < 0 {
        return fmt.Errorf("The binary value must be comprised between 0 and %.1f", maxV)
    }
    *dst = v
    return nil
}

func (p *TempoParam) SetBinValueCurrent(value [4]byte) error { return p.setBinValue(&p.valueCurrent, value) }
func (p *TempoParam) SetBinValueMin(value [4]byte) error { return p.setBinValue(&p.valueMin, value) }
func (p *TempoParam) SetBinValueMax(value [4]byte) error { return p.setBinValue(&p.valueMax, value) }

func (p *TempoParam) setValue(dst *float32, s string) error {
    for i, _s := range p.GetAllowedValues() {
        if s == _s {
            *dst = float32(i+2)
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
    if  v > p.max || p.min > v {
        return fmt.Errorf("The value must be comprised between 0.10 and 15 or be in the list")
    }
    *dst = (v - p.min)/(p.max - p.min)
    return nil
}

func (p * TempoParam) SetValueCurrent(s string) error { return p.setValue(&p.valueCurrent, s) }
func (p * TempoParam) SetValueMin(s string) error { return p.setValue(&p.valueMin, s) }
func (p * TempoParam) SetValueMax(s string) error { return p.setValue(&p.valueMax, s) }

type TimeParam struct {
    GenericParameter
    maxMs        int
    valueCurrent float32
    valueMin     float32
    valueMax     float32
}

func (p *TimeParam) Copy() Parameter {
    _p := new(TimeParam)
    *_p = *p
    return _p
}

func (p *TimeParam) IsAllowingOtherValues() bool { return true }
func (p *TimeParam) GetAllowedValues() []string { return nil }
func (p *TimeParam) GetBinValueCurrent() [4]byte { return to4Bytes(p.valueCurrent) }
func (p *TimeParam) GetBinValueMin() [4]byte { return to4Bytes(p.valueMin) }
func (p *TimeParam) GetBinValueMax() [4]byte { return to4Bytes(p.valueMax) }
func (p *TimeParam) GetBinValueType() uint32 { return float32Type }

func (p *TimeParam) getValue(v float32) string {
    return fmt.Sprintf("%dms", int(v*float32(p.maxMs)))
}

func (p *TimeParam) GetValueCurrent() string { return p.getValue(p.valueCurrent) }
func (p *TimeParam) GetValueMin() string { return p.getValue(p.valueMin) }
func (p *TimeParam) GetValueMax() string { return p.getValue(p.valueMax) }

func (p *TimeParam) setBinValue(dst *float32, value [4]byte) error {
    v := from4BytesToFloat32(value)
    if v > 1 || v < 0 {
        return fmt.Errorf("The binary value must be comprised between 0 and 1")
    }
    *dst = v
    return nil
}

func (p *TimeParam) SetBinValueCurrent(value [4]byte) error { return p.setBinValue(&p.valueCurrent, value) }
func (p *TimeParam) SetBinValueMin(value [4]byte) error { return p.setBinValue(&p.valueMin, value) }
func (p *TimeParam) SetBinValueMax(value [4]byte) error { return p.setBinValue(&p.valueMax, value) }

func (p *TimeParam) setValue(dst *float32, s string) error {
    s = strings.Replace(s, " ", "", -1)
    s = strings.Replace(s, "ms", "", -1)
    vi, err := strconv.Atoi(s)
    if err != nil {
        return err
    }
    if  vi > p.maxMs || vi < 0 {
        return fmt.Errorf("The value must be comprised between 0 and %d", p.maxMs)
    }
    *dst = float32(vi)/float32(p.maxMs)
    return nil
}

func (p *TimeParam) SetValueCurrent(s string) error { return p.setValue(&p.valueCurrent, s) }
func (p *TimeParam) SetValueMin(s string) error { return p.setValue(&p.valueMin, s) }
func (p *TimeParam) SetValueMax(s string) error { return p.setValue(&p.valueMax, s) }
