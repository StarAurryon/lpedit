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
import "log"
import "sort"
import "sync"

type ChangeType int

const (
    ActiveChange ChangeType = iota
    Error
    ErrorStop
    None
    NormalStop
    NormalStart
    ParameterChange
    PresetChange
    PresetLoad
    SetChange
    TempoChange
    TypeChange
    Warning
)

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

type PedalBoardSplitChannel struct{
    aVol   float32
    bVol   float32
    aPan   float32
    bPan   float32
}

type PedalBoard struct {
    items         []PedalBoardItem
    split         PedalBoardSplitChannel
    tempo         float32
    mux           sync.Mutex
    currentSet    *set
    currentPreset *preset
    setList       [numberSet]*set
}

func NewPedalBoard() *PedalBoard {
    pb := &PedalBoard{}

    for i := 0; i < numberSet; i++ {
        name := fmt.Sprintf("Set %d", i + 1)
        pb.setList[i] = newSet(uint8(i), name)
    }

    pb.items = make([]PedalBoardItem, 12)

    pb.items[0] = newDisAmp(uint32(0), uint16(0), AmpAPos, pb)
    pb.items[1] = newNoCab(uint32(1), pb)
    pb.items[2] = newDisAmp(uint32(2), uint16(0), AmpBPos, pb)
    pb.items[3] = newNoCab(uint32(3), pb)
    for id := uint32(4); id <= 11; id++ {
        pb.items[id] = newNonePedal(id, uint16(id - 4), PedalPosStart, pb)
    }
    return pb
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
    ret := make([][]string, presetPerSet)
    if setID < 0 || numberSet <= setID { return ret }
    for i, preset := range pb.setList[setID].presetList {
        ret[i] = []string{preset.GetID2(), preset.GetName()}
    }
    return ret
}

func (pb *PedalBoard) GetSetList() []string {
    ret := make([]string, numberSet)
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
    if id < presetPerSet && pb.currentSet != nil {
        pb.currentPreset = pb.currentSet.presetList[id]
    } else {
        pb.currentPreset = nil
    }
}

func (pb *PedalBoard) SetCurrentSet(id uint32) {
    if id < numberSet {
        pb.currentSet = pb.setList[id]
    } else {
        pb.currentSet = nil
    }
}

func (pb *PedalBoard) SetTempo(t float32) {
    pb.tempo = t
}

func (c *PedalBoard) UnlockData(){
    c.mux.Unlock()
}

func (pb PedalBoard) LogInfo() {
    var name string
    log.Printf("PedalBoard info\n")
    pname := pb.GetCurrentPresetName()
    if len(pname) >= 2 {
        name = fmt.Sprintf("%s: %s", pname[0], pname[1])
    }
    log.Printf("Preset name \"%s\"\n", name)
    log.Printf("PedalStart:\n")
    for _, pbi := range(pb.GetItems(PedalPosStart)) { pbi.LogInfo() }
    log.Printf("Channel A:\n")
    log.Printf("Volume %f, pan %f\n", pb.split.aVol, pb.split.aPan)
    for _, pbi := range(pb.GetItems(PedalPosAStart)) { pbi.LogInfo() }
    for _, pbi := range(pb.GetItems(AmpAPos)) { pbi.LogInfo() }
    for _, pbi := range(pb.GetItems(PedalPosAEnd)) { pbi.LogInfo() }
    log.Printf("Channel B:\n")
    log.Printf("Volume %f, pan %f\n", pb.split.bVol, pb.split.bPan)
    for _, pbi := range(pb.GetItems(PedalPosBStart)) { pbi.LogInfo() }
    for _, pbi := range(pb.GetItems(AmpBPos)) { pbi.LogInfo() }
    for _, pbi := range(pb.GetItems(PedalPosBEnd)) { pbi.LogInfo() }
    log.Printf("PedalEnd:\n")
    for _, pbi := range(pb.GetItems(PedalPosEnd)) { pbi.LogInfo() }
    log.Printf("\n")
}
