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
import "sync"

type PedalBoardChannel struct{
    aStart []PedalBoardItem
    aAmp   []PedalBoardItem
    aEnd   []PedalBoardItem
    bStart []PedalBoardItem
    bAmp   []PedalBoardItem
    bEnd   []PedalBoardItem
    aVol   float32
    bVol   float32
    aPan   float32
    bPan   float32
}

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

type PedalBoard struct {
    start         []PedalBoardItem
    startAmp      []PedalBoardItem
    end           []PedalBoardItem
    endAmp        []PedalBoardItem
    pchan         PedalBoardChannel
    bAmp          []PedalBoardItem //Only for backup of channel B
    cabs          []PedalBoardItem //Pod still sending infos even if not present
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
        pb.setList[i] = newSet(uint32(i), name)
    }

    ca := newNoCab(uint32(1), pb, &pb.cabs)
    newDisAmp(uint32(0), pb, &pb.pchan.aAmp, ca)
    cb := newNoCab(uint32(3), pb, &pb.cabs)
    newDisAmp(uint32(2), pb, &pb.pchan.bAmp, cb)
    for id := uint32(4); id <= 11; id++ {
        newNonePedal(id, pb, &pb.start)
    }
    return pb
}

func (pb *PedalBoard) GetItem(id uint32) PedalBoardItem {
    var p PedalBoardItem
    ps := []*[]PedalBoardItem{&pb.start, &pb.startAmp,
         &pb.pchan.aStart, &pb.pchan.aAmp, &pb.pchan.aEnd,
         &pb.pchan.bStart, &pb.pchan.bAmp, &pb.pchan.bEnd,
         &pb.endAmp, &pb.end,
         &pb.bAmp, &pb.cabs}
    for i, _ := range(ps) {
        for j, _ := range(*ps[i]) {
            p = (*ps[i])[j]
            if p.GetID() == id {
                return p
            }
        }
    }
    return nil
}

func (pb *PedalBoard) GetItems(posType uint8) []PedalBoardItem{
    switch posType {
    case PedalPosStart:
        return pb.start
    case PedalPosAStart:
        return pb.pchan.aStart
    case PedalPosBStart:
        return pb.pchan.bStart
    case PedalPosAEnd:
        return pb.pchan.aEnd
    case PedalPosBEnd:
        return pb.pchan.bEnd
    case PedalPosEnd:
        return pb.end
    case AmpAPos:
        return pb.pchan.aAmp
    case AmpBPos:
        return pb.pchan.bAmp
    }
    return nil
}

func (pb *PedalBoard) GetPedal(pos uint16) PedalBoardItem {
    list := append(pb.start, pb.pchan.aStart...)
    list = append(list, pb.pchan.aEnd...)
    list = append(list, pb.pchan.bStart...)
    list = append(list, pb.pchan.bEnd...)
    list = append(list, pb.end...)
    if pos < uint16(len(list)) {
        return list[pos]
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

func (pb *PedalBoard) GetCurrentSet() (error, uint32) {
    if s := pb.currentSet; s != nil {
        return nil, s.GetID()
    }
    return fmt.Errorf("Current set is not defined"), 0
}

func (pb *PedalBoard) GetCurrentPreset() (error, uint32) {
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
    for _, pbi := range(pb.start) { pbi.LogInfo() }
    for _, pbi := range(pb.startAmp) { pbi.LogInfo() }
    log.Printf("Channel A:\n")
    log.Printf("Volume %f, pan %f\n", pb.pchan.aVol, pb.pchan.aPan)
    for _, pbi := range(pb.pchan.aStart) { pbi.LogInfo() }
    for _, pbi := range(pb.pchan.aAmp) { pbi.LogInfo() }
    for _, pbi := range(pb.pchan.aEnd) { pbi.LogInfo() }
    log.Printf("Channel B:\n")
    log.Printf("Volume %f, pan %f\n", pb.pchan.bVol, pb.pchan.bPan)
    for _, pbi := range(pb.pchan.bStart) { pbi.LogInfo() }
    for _, pbi := range(pb.pchan.bAmp) { pbi.LogInfo() }
    for _, pbi := range(pb.pchan.bEnd) { pbi.LogInfo() }
    log.Printf("PedalEnd:\n")
    for _, pbi := range(pb.endAmp) { pbi.LogInfo() }
    for _, pbi := range(pb.end) { pbi.LogInfo() }
    log.Printf("\n")
}
