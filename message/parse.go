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

import "encoding/binary"
import "fmt"
import "bytes"
import "log"

import "lpedit/pedal"

type presetPedalPos struct {
    pid   uint32
    ptype uint8
}

func (m Message) getPedalBoardItemID() uint32 {
    return binary.LittleEndian.Uint32(m.data[12:16])
}

func (m ActiveChange) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    id := m.getPedalBoardItemID()
    p := pb.GetItem(id)
    if p == nil {
        return fmt.Errorf("Item ID %d not found", id), pedal.Warning, nil
    }
    var active bool
    if binary.LittleEndian.Uint32(m.data[16:]) > 0 {
        active = true
    } else {
        active = false
    }
    log.Printf("Active change on ID %d status %t\n", id, active)
    p.SetActive(active)
    return nil, pedal.ActiveChange, p
}

func (m ParameterChange) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    pid := m.getPedalBoardItemID()
    p := pb.GetItem(pid)
    if p == nil {
        return fmt.Errorf("Item ID %d not found", pid), pedal.Warning, nil
    }
    id := binary.LittleEndian.Uint16(m.data[20:22])
    //Dirty fix, need to understand more the protocol
    if m.data[22] == 1 {
        id++
    }
    var v [4]byte
    copy(v[:], m.data[24:])

    //Dirty fix, need to understand more the protocol
    id %= 6142

    fmt.Println(id)
    param := p.GetParam(id)
    if param == nil {
        return fmt.Errorf("Parameter ID %d not found", id), pedal.Warning, nil
    }
    if err := param.SetBinValue(v); err != nil {
        log.Printf("TODO: Fix the parameter type on pedal %s: %s \n", p.GetName(), err)
    }
    return nil, pedal.ParameterChange, param
}

func (m ParameterTempoChange) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    pid := m.getPedalBoardItemID()
    p := pb.GetItem(pid)
    if p == nil {
        return fmt.Errorf("Item ID %d not found", pid), pedal.Warning, nil
    }
    param := p.GetParam(0)
    if param == nil {
        return fmt.Errorf("Parameter ID 0 not found"), pedal.Warning, nil
    }
    value := float32(binary.LittleEndian.Uint32(m.data[16:]))
    binValue := [4]byte{}
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, value)
    copy(binValue[:], buf.Bytes())
    param.SetBinValue(binValue)
    return nil, pedal.ParameterChange, param
}

func (m ParameterTempoChange2) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    pid := m.getPedalBoardItemID()
    p := pb.GetItem(pid)
    if p == nil {
        return fmt.Errorf("Item ID %d not found", pid), pedal.Warning, nil
    }
    param := p.GetParam(2)
    if param == nil {
        return fmt.Errorf("Parameter ID 2 not found"), pedal.Warning, nil
    }
    value := float32(binary.LittleEndian.Uint32(m.data[16:]))
    binValue := [4]byte{}
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, value)
    copy(binValue[:], buf.Bytes())
    param.SetBinValue(binValue)
    return nil, pedal.ParameterChange, param
}

func (m PresetChange) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    pb.SetCurrentPreset(binary.LittleEndian.Uint32(m.data[8:]))
    return nil, pedal.PresetChange, pb
}

func (m PresetChangeAlert) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    return nil, pedal.None, nil
}

func (m PresetLoad) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    pbiOrder := []uint32{0,2,1,3,4,5,6,7,8,9,10,11}
    pb.SetCurrentPresetName(string(m.data[8:40]))

    offset := 48

    for i := 0; i < len(pbiOrder); i++ {
        //Pedal Board Item Setup (Amp, Cab, Pedal)
        pbi := pb.GetItem(pbiOrder[i])
        itype := binary.LittleEndian.Uint32(m.data[offset:offset+4])
        pbi.SetType(itype)

        //Pedal Board Item order gathering
        pos := binary.LittleEndian.Uint16(m.data[offset+4:offset+6])
        posType := uint8(m.data[offset+6:offset+7][0])
        pbi.SetPos(pos, posType)

        //Pedal Board Parameter Setup
        poffset := offset
        for j := uint16(0); j < pbi.GetParamLen(); j++ {
            pidx := binary.LittleEndian.Uint16(m.data[poffset+16:poffset+18])
            pidx %= 6142
            param := pbi.GetParam(pidx)
            if param != nil {
                tempoValue := binary.LittleEndian.Uint16(m.data[poffset+9:poffset+11])
                var v float32
                if j == 0 && tempoValue > 1{
                    v = float32(tempoValue)
                } else {
                    binary.Read(bytes.NewReader(m.data[poffset+20:poffset+24]), binary.LittleEndian, &v)
                }
                binValue := [4]byte{}
                buf := new(bytes.Buffer)
                binary.Write(buf, binary.LittleEndian, v)
                copy(binValue[:], buf.Bytes())
                if err := param.SetBinValue(binValue); err != nil {
                    log.Printf("TODO: Fix the parameter type on pedal %s: %s \n", pbi.GetName(), err)
                }
            } else {
                log.Printf("TODO: Parameter ID %d does not exist on pedal type %s\n",
                    pidx, pbi.GetName())
            }

            poffset += 20
        }

        offset += 256
    }
    return nil, pedal.PresetLoad, pb
}

func (m SetChange) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    pb.SetCurrentSet(binary.LittleEndian.Uint32(m.data[8:]))
    return nil, pedal.SetChange, pb
}

func (m SetupChange) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    var v float32
    err := binary.Read(bytes.NewReader(m.data[20:]), binary.LittleEndian, &v)
    if err != nil {
        return err, pedal.Warning, nil
    }
    pb.SetTempo(v)
    return nil, pedal.TempoChange, pb
}

func (m TypeChange) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    id := m.getPedalBoardItemID()
    p := pb.GetItem(id)
    if p == nil {
        return fmt.Errorf("Item ID %d not found", id), pedal.Warning, nil
    }
    ptype := binary.LittleEndian.Uint32(m.data[16:])
    if err := p.SetType(ptype); err != nil {
        return err, pedal.Warning, nil
    }
    return nil, pedal.TypeChange, p
}
