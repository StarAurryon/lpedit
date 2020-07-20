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
    a    []*Pedal
    b    []*Pedal
    aVol float32
    bVol float32
    aPan float32
    bPan float32
}

type pedalPos uint16

const (
    pedalPosStart = 0
    pedalPosA     = 1
    pedalPosB     = 2
    pedalPosEnd   = 5
)

type PedalBoard struct {
    start  []*Pedal
    end    []*Pedal
    pchan  PedalBoardChannel
    pname  string
}

func NewPedalBoard() *PedalBoard {
    p := &PedalBoard{}
    for id := uint32(4); id <= 11; id++ {
        NewNonePedal(id, p, &p.start)
    }
    return p
}

func (pb PedalBoard) PrintInfo() {
    fmt.Printf("PedalBoard info\n")
    fmt.Printf("Preset name \"%s\"\n", pb.pname)
    fmt.Printf("PedalStart:\n")
    for _, p := range(pb.start) {
        p.PrintInfo()
    }
    fmt.Printf("Channel A:\n")
    fmt.Printf("Volume %f, pan %f\n", pb.pchan.aVol, pb.pchan.aPan)
    for _, p := range(pb.pchan.a) {
        p.PrintInfo()
    }
    fmt.Printf("Channel B:\n")
    fmt.Printf("Volume %f, pan %f\n", pb.pchan.bVol, pb.pchan.bPan)
    for _, p := range(pb.pchan.b) {
        p.PrintInfo()
    }
    fmt.Printf("PedalEnd:\n")
    for _, p := range(pb.end) {
        p.PrintInfo()
    }
    fmt.Printf("\n")
}

func (pb *PedalBoard) GetPedal(id uint32) *Pedal{
    var p *Pedal
    ps := []*[]*Pedal{&pb.start, &pb.pchan.a, &pb.pchan.b, &pb.end}
    for i, _ := range(ps) {
        for j, _ := range(*ps[i]) {
            p = (*ps[i])[j]
            if p.id == id {
                return p
            }
        }
    }
    return nil
}

func (pb *PedalBoard) SetPedal(id uint32, ptype PedalType) error{
    p := pb.GetPedal(id)
    if p == nil {
        return fmt.Errorf("Pedal ID %d not found", id)
    }
    pt, found := pedals[ptype]
    if !found {
        return fmt.Errorf("Pedal type not found, code: %d", ptype)
    }
    plist := p.plist
    *p = pt
    p.id = id
    p.pb = pb
    p.plist = plist
    return nil
}

func (pb *PedalBoard) SetPresetName(pname string) {
    pb.pname = pname
}
