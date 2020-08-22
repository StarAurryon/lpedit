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

import "fmt"
import "sort"
import "sync"

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
    PedalBoardInput1Source uint32 = iota
    PedalBoardInput2Source
    PedalBoardGuitarInZ
)

type LPEObject interface {
    GetParam(uint32) Parameter
    GetParams() []Parameter
    GetParamLen() uint16
    GetName() string
    LockData()
    UnlockData()
}

type PedalBoardSplitChannel struct{
    aVol   float32
    bVol   float32
    aPan   float32
    bPan   float32
}

type PedalBoard struct {
    items         []PedalBoardItem
    dts           []*DT
    split         PedalBoardSplitChannel
    tempo         float32
    mux           sync.Mutex
    currentSet    *set
    currentPreset *preset
    setList       [NumberSet]*set
    parameters    []Parameter
}

func NewPedalBoard() *PedalBoard {
    pb := &PedalBoard{}

    for i := 0; i < NumberSet; i++ {
        name := fmt.Sprintf("Set %d", i + 1)
        pb.setList[i] = newSet(uint8(i), name)
    }

    pb.items = make([]PedalBoardItem, 12)

    pb.items[0] = newDisAmp(uint32(0), uint16(0), AmpAPos, pb)
    pb.items[1] = newNoCab(uint32(1), uint16(0), AmpAPos, pb)
    pb.items[2] = newDisAmp(uint32(2), uint16(0), AmpBPos, pb)
    pb.items[3] = newNoCab(uint32(3), uint16(0), AmpBPos, pb)
    for id := uint32(4); id <= 11; id++ {
        pb.items[id] = newNonePedal(id, uint16(id - 4), PedalPosStart, pb)
    }

    pb.dts = make([]*DT, 2)
    pb.dts[0] = newDT(0, uint32(0),pb)
    pb.dts[1] = newDT(1, uint32(2),pb)

    pb.parameters = make([]Parameter, 3)
    pb.parameters[0] = &ListParam{GenericParameter: GenericParameter{id: PedalBoardInput1Source,
            name: "Input 1 Source", parent: pb},
        binValueType: Int32Type, maxIDShift: 1,
        list: []string{"Guitar", "Mic", "Aux", "", "Guitar+Aux", "Guitar+Variax", "Guitar+Aux+Variax",
            "Variax", "Variax Mags"}}
    pb.parameters[1] = &ListParam{GenericParameter: GenericParameter{id: PedalBoardInput2Source,
            name: "Input 2 Source", parent: pb}, binValueType: Int32Type,
        list: []string{"Same", "Guitar", "Mic", "Aux", "", "Guitar+Aux", "Guitar+Variax", "Guitar+Aux+Variax",
            "Variax", "Variax Mags"}}
    pb.parameters[2] = &ListParam{GenericParameter: GenericParameter{id: PedalBoardGuitarInZ,
            name: "Guitar In-Z", parent: pb}, binValueType: Int32Type,
        list: []string{"Auto", "22K", "32K", "70K", "90K", "136K", "230K", "1M", "3.5M"}}

    return pb
}

func (pb *PedalBoard) GetAmp(id int) *Amp {
    if id < 0 || id > 2 {
        return nil
    }
    a, _ := pb.GetItem(uint32(id*2)).(*Amp)
    return a
}

func (pb *PedalBoard) GetCab(id int) *Cab {
    if id < 0 || id > 2 {
        return nil
    }
    a, _ := pb.GetItem(uint32((id*2) + 1)).(*Cab)
    return a
}

func (pb *PedalBoard) GetDT(ID int) *DT {
    for _, dt := range pb.dts {
        if dt.GetID() == ID {
            return dt
        }
    }
    return nil
}

func (pb *PedalBoard) GetDT2(ampID uint32) *DT {
    for _, dt := range pb.dts {
        if dt.GetAmpID() == ampID {
            return dt
        }
    }
    return nil
}

func (pb *PedalBoard) GetItem(id uint32) PedalBoardItem {
    for _, pbi := range pb.items {
        if pbi.GetID() == id {
            return pbi
        }
    }
    return nil
}

func (pb *PedalBoard) GetItems(posType uint8) []PedalBoardItem{
    ret := []PedalBoardItem{}
    if posType == AmpBPos && len(pb.GetItems(AmpAPos)) == 0 { return ret }
    for _, pbi := range pb.items {
        if _, _posType := pbi.GetPos(); _posType == posType {
            ret = append(ret, pbi)
        }
    }
    sort.Sort(SortablePosPBI(ret))
    return ret
}

func (pb *PedalBoard) GetName() string {
    return ""
}

func (pb *PedalBoard) GetParam(id uint32) Parameter {
    for _, p := range pb.parameters {
        if p.GetID() == id {
            return p
        }
    }
    return nil
}

func (pb *PedalBoard) GetParams() []Parameter {
    return pb.parameters
}

func (pb *PedalBoard) GetParamLen() uint16 {
    return uint16(len(pb.parameters))
}

func (pb *PedalBoard) GetPedal(pos uint16) PedalBoardItem {
    for _, pbi := range pb.items {
        if _pos, _ := pbi.GetPos(); _pos == pos {
            return pbi
        }
    }
    return nil
}

func (pb *PedalBoard) GetPedal2(id int) *Pedal {
    if id < 0 || id > 7 {
        return nil
    }
    p, _ := pb.GetItem(uint32(id + 4)).(*Pedal)
    return p
}

func (pb *PedalBoard) GetCurrentSet() (error, uint8) {
    if s := pb.currentSet; s != nil {
        return nil, s.GetID()
    }
    return fmt.Errorf("Current set is not defined"), 0
}

func (pb *PedalBoard) GetCurrentPreset() (error, uint8) {
    if p := pb.currentPreset; p != nil {
        return nil, p.GetID()
    }
    return fmt.Errorf("Current set is not defined"), 0
}

func (pb *PedalBoard) GetCurrentPresetName() []string {
    if p := pb.currentPreset; p != nil {
        return []string{p.GetID2(), p.GetName()}
    }
    return nil
}

func (pb *PedalBoard) GetTempo() float32 {
    return pb.tempo
}

func (pb *PedalBoard) GetPresetList(setID int) [][]string {
    ret := make([][]string, PresetPerSet)
    if setID < 0 || NumberSet <= setID { return ret }
    for i, preset := range pb.setList[setID].presetList {
        ret[i] = []string{preset.GetID2(), preset.GetName()}
    }
    return ret
}

func (pb *PedalBoard) GetSetList() []string {
    ret := make([]string, NumberSet)
    for i, set := range pb.setList {
        ret[i] = set.GetName()
    }
    return ret
}

func (c *PedalBoard) LockData(){
    c.mux.Lock()
}

func (pb *PedalBoard) SetCurrentPresetName(name string) {
    if p := pb.currentPreset; p != nil {
        p.SetName(name)
    }
}

func (pb *PedalBoard) SetCurrentPreset(id uint32) {
    if id < PresetPerSet && pb.currentSet != nil {
        pb.currentPreset = pb.currentSet.presetList[id]
    } else {
        pb.currentPreset = nil
    }
}

func (pb *PedalBoard) SetCurrentSet(id uint32) {
    if id < NumberSet {
        pb.currentSet = pb.setList[id]
    } else {
        pb.currentSet = nil
    }
}

func (pb *PedalBoard) SetCurrentSetName(name string) {
    if s := pb.currentSet; s != nil {
        s.SetName(name)
    }
}

func (pb *PedalBoard) SetTempo(t float32) {
    pb.tempo = t
}

func (c *PedalBoard) UnlockData(){
    c.mux.Unlock()
}
