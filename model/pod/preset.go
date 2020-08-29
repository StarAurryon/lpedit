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

package pod

import "sort"
import "strings"
import "strconv"

const (
    PedalPosStart  uint8 = 0
    PedalPosAStart uint8 = 1
    PedalPosBStart uint8 = 2
    PedalPosAEnd   uint8 = 3
    PedalPosBEnd   uint8 = 4
    PedalPosEnd    uint8 = 5
    AmpAPos        uint8 = 7
    AmpBPos        uint8 = 8
)

const (
    PresetInput1Source uint32 = iota
    PresetInput2Source
    PresetGuitarInZ
    PresetTempo
)

type PedalBoardSplitChannel struct{
    aVol   float32
    bVol   float32
    aPan   float32
    bPan   float32
}

type Preset struct {
    id            uint8
    dts           [2]*DT
    items         [12]PedalBoardItem
    name          [16]byte
    parameters    [4]Parameter
    set           *Set
    split         PedalBoardSplitChannel
}

func newPreset(id uint8, s *Set) *Preset {
    p := &Preset{id: id, set: s}

    p.items[0] = newDisAmp(uint32(0), uint16(0), AmpAPos, p)
    p.items[1] = newNoCab(uint32(1), uint16(0), AmpBPos, p)
    p.items[2] = newDisAmp(uint32(2), uint16(0), AmpAPos, p)
    p.items[3] = newNoCab(uint32(3), uint16(0), AmpBPos, p)
    for id := uint32(4); id <= 11; id++ {
        p.items[id] = newNonePedal(id, uint16(id - 4), PedalPosStart, p)
    }

    p.dts[0] = newDT(0, uint32(0),p)
    p.dts[1] = newDT(1, uint32(2),p)

    p.parameters[0] = &ListParam{GenericParameter: GenericParameter{id: PresetInput1Source,
            name: "Input 1 Source", parent: p},
        binValueType: Int32Type, maxIDShift: 1,
        list: []string{"Guitar", "Mic", "Aux", "", "Guitar+Aux", "Guitar+Variax", "Guitar+Aux+Variax",
            "Variax", "Variax Mags"}}
    p.parameters[1] = &ListParam{GenericParameter: GenericParameter{id: PresetInput2Source,
            name: "Input 2 Source", parent: p}, binValueType: Int32Type,
        list: []string{"Same", "Guitar", "Mic", "Aux", "", "Guitar+Aux", "Guitar+Variax", "Guitar+Aux+Variax",
            "Variax", "Variax Mags"}}
    p.parameters[2] = &ListParam{GenericParameter: GenericParameter{id: PresetGuitarInZ,
            name: "Guitar In-Z", parent: p}, binValueType: Int32Type,
        list: []string{"Auto", "22K", "32K", "70K", "90K", "136K", "230K", "1M", "3.5M"}}
    p.parameters[3] = &RangeParam{GenericParameter: GenericParameter{id: PresetTempo,
            name: "Tempo", parent: p}, min: 30, max: 240}

    return p
}

func (p *Preset) GetAmp(id int) *Amp {
    if id < 0 || id > 2 {
        return nil
    }
    a, _ := p.GetItem(uint32(id*2)).(*Amp)
    return a
}

func (p *Preset) GetCab(id int) *Cab {
    if id < 0 || id > 2 {
        return nil
    }
    a, _ := p.GetItem(uint32((id*2) + 1)).(*Cab)
    return a
}

func (p *Preset) GetDT(ID int) *DT {
    for _, dt := range p.dts {
        if dt.GetID() == ID {
            return dt
        }
    }
    return nil
}

func (p *Preset) GetDT2(ampID uint32) *DT {
    for _, dt := range p.dts {
        if dt.GetAmpID() == ampID {
            return dt
        }
    }
    return nil
}

func (p *Preset) GetID() uint8 {
    return p.id
}

func (p *Preset) GetID2() string {
    id := p.GetID()
    section := (id / 4) + 1
    letter := string([]byte{uint8((id % 4)  + 65)})
    return strconv.Itoa(int(section)) + letter
}

func (p *Preset) GetItem(id uint32) PedalBoardItem {
    for _, pbi := range p.items {
        if pbi.GetID() == id {
            return pbi
        }
    }
    return nil
}

func (p *Preset) GetItems(posType uint8) []PedalBoardItem{
    ret := []PedalBoardItem{}
    if posType == AmpBPos && len(p.GetItems(AmpAPos)) == 0 { return ret }
    for _, pbi := range p.items {
        if _, _posType := pbi.GetPos(); _posType == posType {
            ret = append(ret, pbi)
        }
    }
    sort.Sort(SortablePosPBI(ret))
    return ret
}

func (p *Preset) GetName() string {
    ret := string(p.name[:])
    ret = strings.Trim(ret, " ")
    return ret
}

func (p *Preset) GetName2() [16]byte {
    return p.name
}

func (p *Preset) GetName3() []string {
    return []string{p.GetID2(), p.GetName()}
}

func (p *Preset) GetParam(id uint32) Parameter {
    for _, param := range p.parameters {
        if param.GetID() == id {
            return param
        }
    }
    return nil
}

func (p *Preset) GetParams() []Parameter {
    return p.parameters[:]
}

func (p *Preset) GetParamLen() uint16 {
    return uint16(len(p.parameters))
}

func (p *Preset) GetPedal(pos uint16) *Pedal {
    for _, pbi := range p.items {
        switch p := pbi.(type) {
        case *Pedal:
            if _pos, _ := pbi.GetPos(); _pos == pos {
                return p
            }
        }
    }
    return nil
}

func (p *Preset) GetPedal2(id int) *Pedal {
    if id < 0 || id > 7 {
        return nil
    }
    pedal, _ := p.GetItem(uint32(id + 4)).(*Pedal)
    return pedal
}

func (p *Preset) GetSet() *Set {
    return p.set
}

func (p *Preset) LockData() {
    p.set.LockData()
}

func (p *Preset) SetName(name [16]byte) {
    p.name = name
}

func (p *Preset) SetName2(name string) {
    src := []byte(name)
    copy(p.name[:], src)
    for i := len(src); i < len(p.name); i++ {
        p.name[i] = byte(' ')
    }
}

func (p *Preset) UnlockData() {
    p.set.UnlockData()
}
