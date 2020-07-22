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

package message

import "lpedit/pedal"
import "encoding/binary"
import "fmt"
import "bytes"
import "sort"

type sUint16 []uint16

func (a sUint16) Len() int           { return len(a) }
func (a sUint16) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sUint16) Less(i, j int) bool { return a[i] < a[j] }

type presetPedalPos struct {
    pid   uint32
    ptype uint8
}

func (m Message) getPedalID() uint32 {
    return binary.LittleEndian.Uint32(m.data[12:16])
}

func loadPreset(m Message, pb *pedal.PedalBoard) error {
    pbiOrder := []uint32{0,2,1,3,4,5,6,7,8,9,10,11}
    // ppos contains the position as key
    ppos := map[uint16]presetPedalPos{}
    ampPos := uint16(0)
    ampPosType := uint8(0)
    pb.SetPresetName(string(m.data[8:40]))

    offset := 48

    for i := 0; i < len(pbiOrder); i++ {
        //Pedal Board Item Setup (Amp, Cab, Pedal)
        pbi := pb.GetItem(pbiOrder[i])
        itype := binary.LittleEndian.Uint32(m.data[offset:offset+4])
        pbi.SetType(itype)

        //Pedal Board Item order gathering
        switch pbi.(type){
        case *pedal.Pedal:
            pos := binary.LittleEndian.Uint16(m.data[offset+4:offset+6])
            ppos[pos] = presetPedalPos{pid: pbiOrder[i],
                ptype: uint8(m.data[offset+6:offset+7][0])}
        case *pedal.Amp:
            if i == 0 {
                ampPos = binary.LittleEndian.Uint16(m.data[offset+4:offset+6])
                ampPosType = uint8(m.data[offset+6:offset+7][0])
            }
        }

        //Pedal Board Parameter Setup
        poffset := offset
        for j := uint16(0); j < pbi.GetParamLen(); j++ {
            pidx := binary.LittleEndian.Uint16(m.data[poffset+16:poffset+18])
            var v float32
            binary.Read(bytes.NewReader(m.data[poffset+20:poffset+24]), binary.LittleEndian, &v)
            param := pbi.GetParam(pidx)
            if param != nil {
                param.SetValue(v)
            } else {
                fmt.Printf("TODO: Parameter ID %d does not exist on pedal type %s\n",
                    pidx, pbi.GetName())
            }

            poffset += 20
        }

        offset += 256
    }
    // Position of the pedal setup
    pos := make(sUint16, 0, len(ppos))
    for k := range ppos {
        pos = append(pos, k)
    }
    sort.Sort(pos)
    for k := uint16(0); k < uint16(len(pos)); k++ {
        err := pb.GetItem(ppos[k].pid).SetLastPos(k, ppos[k].ptype)
        if err != nil {
            return err
        }
    }
    pb.GetItem(0).SetLastPos(ampPos, ampPosType)
    return nil
}

func itemActiveChange(m Message, pb *pedal.PedalBoard) error {
    id := m.getPedalID()
    p := pb.GetItem(id)
    if p == nil {
        return fmt.Errorf("Item ID %d not found", id)
    }
    var active bool
    if binary.LittleEndian.Uint32(m.data[16:]) > 0 {
        active = true
    } else {
        active = false
    }
    fmt.Printf("Active change on ID %d status %t\n", id, active)
    p.SetActive(active)
    return nil
}

func itemParameterChange (m Message, pb *pedal.PedalBoard) error {
    pid := m.getPedalID()
    p := pb.GetItem(pid)
    if p == nil {
        return fmt.Errorf("Item ID %d not found", pid)
    }
    id := binary.LittleEndian.Uint16(m.data[20:22])
    var v float32
    err := binary.Read(bytes.NewReader(m.data[24:28]), binary.LittleEndian, &v)
    if err != nil {
        return err
    }

    param := p.GetParam(id)
    if param == nil {
        return fmt.Errorf("Parameter ID %d not found", id)
    }
    param.SetValue(v)
    return nil
}

func itemTypeChange(m Message, pb *pedal.PedalBoard) error {
    id := m.getPedalID()
    ptype := binary.LittleEndian.Uint32(m.data[16:])
    return pb.SetItem(id, ptype)
}
