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

func (m Message) getPedalID() uint32 {
    return binary.LittleEndian.Uint32(m.data[12:16])
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
    return p.SetParameter(id, v)
}

func pedalTypeChange(m Message, pb *pedal.PedalBoard) error {
    id := m.getPedalID()
    ptype := pedal.PedalType(binary.LittleEndian.Uint32(m.data[16:]))
    return pb.SetPedal(id, ptype)
}
