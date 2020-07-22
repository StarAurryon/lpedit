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

type pedalPos uint16

const (
    pedalPosStart  uint8 = 0
    pedalPosAStart uint8 = 1
    pedalPosBStart uint8 = 2
    pedalPosAEnd   uint8 = 3
    pedalPosBEnd   uint8 = 4
    pedalPosEnd    uint8 = 5
    ampAPos        uint8 = 7
    ampBPos        uint8 = 8
)

type PedalBoard struct {
    start    []PedalBoardItem
    startAmp []PedalBoardItem
    end      []PedalBoardItem
    endAmp   []PedalBoardItem
    pchan    PedalBoardChannel
    pname    string
    bAmp     []PedalBoardItem //Only for backup of channel B
    cabs     []PedalBoardItem //Pod still sending infos even if not present
}

func NewPedalBoard() *PedalBoard {
    pb := &PedalBoard{}
    ca := newNoneCab(uint32(1), pb, &pb.cabs)
    newDisAmp(uint32(0), pb, &pb.pchan.aAmp, ca)
    cb := newNoneCab(uint32(3), pb, &pb.cabs)
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

func (pb *PedalBoard) SetItem(id uint32, ptype uint32) error {
    p := pb.GetItem(id)
    if p == nil {
        return fmt.Errorf("Pedal ID %d not found", id)
    }
    return p.SetType(ptype)
}

func (pb *PedalBoard) SetPresetName(pname string) {
    pb.pname = pname
}

func (pb PedalBoard) PrintInfo() {
    fmt.Printf("PedalBoard info\n")
    fmt.Printf("Preset name \"%s\"\n", pb.pname)
    fmt.Printf("PedalStart:\n")
    for _, pbi := range(pb.start) { pbi.PrintInfo() }
    for _, pbi := range(pb.startAmp) { pbi.PrintInfo() }
    fmt.Printf("Channel A:\n")
    fmt.Printf("Volume %f, pan %f\n", pb.pchan.aVol, pb.pchan.aPan)
    for _, pbi := range(pb.pchan.aStart) { pbi.PrintInfo() }
    for _, pbi := range(pb.pchan.aAmp) { pbi.PrintInfo() }
    for _, pbi := range(pb.pchan.aEnd) { pbi.PrintInfo() }
    fmt.Printf("Channel B:\n")
    fmt.Printf("Volume %f, pan %f\n", pb.pchan.bVol, pb.pchan.bPan)
    for _, pbi := range(pb.pchan.bStart) { pbi.PrintInfo() }
    for _, pbi := range(pb.pchan.bAmp) { pbi.PrintInfo() }
    for _, pbi := range(pb.pchan.bEnd) { pbi.PrintInfo() }
    fmt.Printf("PedalEnd:\n")
    for _, pbi := range(pb.endAmp) { pbi.PrintInfo() }
    for _, pbi := range(pb.end) { pbi.PrintInfo() }
    fmt.Printf("\n")
}
