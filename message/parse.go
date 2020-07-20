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
    // ppos contains the position as key
    ppos := map[uint16]presetPedalPos{}
    offset := 1072;

    pb.SetPresetName(string(m.data[8:40]))

    for i := uint32(4); i < 12; i++ {
        ptype := pedal.PedalType(binary.LittleEndian.Uint32(m.data[offset:offset+4]))
        pb.SetPedal(i, ptype)
        p := pb.GetPedal(i)

        // Position of the pedal
        pos := binary.LittleEndian.Uint16(m.data[offset+4:offset+6])
        ppos[pos] = presetPedalPos{pid: i,
            ptype: uint8(m.data[offset+6:offset+7][0])}

        // Value of pedal parameters
        vOffset := offset
        for j := uint32(0); j < p.GetParamLen(); j++ {
            vOffset += 20
            var v float32
            err := binary.Read(bytes.NewReader(m.data[vOffset:vOffset+4]), binary.LittleEndian, &v)
            if err != nil {
                return err
            }
            p.GetParam(j).SetValue(v)
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
        err := pb.GetPedal(ppos[k].pid).SetLastPos(k, ppos[k].ptype)
        if err != nil {
            return err
        }
    }
    return nil
}

func pedalActiveChange(m Message, pb *pedal.PedalBoard) error {
    id := m.getPedalID()
    p := pb.GetPedal(id)
    if p == nil {
        return fmt.Errorf("Pedal ID %d not found", id)
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

func pedalParameterChange (m Message, pb *pedal.PedalBoard) error {
    pid := m.getPedalID()
    p := pb.GetPedal(pid)
    if p == nil {
        return fmt.Errorf("Pedal ID %d not found", pid)
    }
    id := binary.LittleEndian.Uint32(m.data[20:24]) - 1058013185
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

func pedalTypeChange(m Message, pb *pedal.PedalBoard) error {
    id := m.getPedalID()
    ptype := pedal.PedalType(binary.LittleEndian.Uint32(m.data[16:]))
    return pb.SetPedal(id, ptype)
}
