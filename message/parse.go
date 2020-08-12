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
import "reflect"

import "github.com/StarAurryon/lpedit/pedal"

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

func (m *Message) parseParameterChange(paramFunc string, pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    pid := m.getPedalBoardItemID()
    p := pb.GetItem(pid)
    if p == nil {
        return fmt.Errorf("Item ID %d not found", pid), pedal.Warning, nil
    }
    id := binary.LittleEndian.Uint32(m.data[20:24])

    var v [4]byte
    copy(v[:], m.data[24:])

    param := p.GetParam(id)
    if param == nil {
        return fmt.Errorf("Parameter ID %d not found", id), pedal.Warning, nil
    }
    if err := reflect.ValueOf(param).MethodByName(paramFunc).Interface().(func([4]byte) error)(v); err != nil {
        log.Printf("TODO: Fix the parameter type on pedal %s, parameter %s, func %s: %s \n", p.GetName(), param.GetName(), paramFunc, err)
    }
    return nil, pedal.None, param
}

func (m ParameterChange) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    err, pt, obj := m.parseParameterChange("SetBinValueCurrent", pb)
    if err != nil { return err, pt, obj }
    return err, pedal.ParameterChange, obj
}

func (m ParameterChangeMin) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    err, pt, obj := m.parseParameterChange("SetBinValueMin", pb)
    if err != nil { return err, pt, obj }
    return err, pedal.ParameterChangeMin, obj
}

func (m ParameterChangeMax) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    err, pt, obj := m.parseParameterChange("SetBinValueMax", pb)
    if err != nil { return err, pt, obj }
    return err, pedal.ParameterChangeMax, obj
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
    param.SetBinValueCurrent(binValue)
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
    param.SetBinValueCurrent(binValue)
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

    const offset = 48
    var data [256]byte
    for i, id := range pbiOrder {
        start := offset + (i * 256)
        end := start + 256
        copy(data[:], m.data[start:end])
        m.parsePedalBoardItem(pb, data, id)
    }
    return nil, pedal.PresetLoad, pb
}

func (m PresetLoad) parsePedalBoardItem(pb *pedal.PedalBoard, data [256]byte, pbiID uint32) {
    pbi := pb.GetItem(pbiID)

    pbiType := binary.LittleEndian.Uint32(data[0:4])
    pbi.SetType(pbiType)

    pos := binary.LittleEndian.Uint16(data[4:6])
    posType := uint8(data[6])
    pbi.SetPos(pos, posType)

    active := false
    if data[8] == 1 { active = true }
    pbi.SetActive(active)

    tempos := []uint8{data[9], data[10]}

    const offset = 16
    var paramData [20]byte
    for i := uint16(0); i < pbi.GetParamLen(); i++ {
        start := offset + (i * 20)
        end := start + 20
        copy(paramData[:], data[start:end])
        m.parseParameter(pbi, paramData, &tempos)
    }
}

func (m PresetLoad) parseParameter(pbi pedal.PedalBoardItem, data [20]byte, tempos *[]uint8) {
    paramID := binary.LittleEndian.Uint32(data[0:4])
    param := pbi.GetParam(paramID)
    if param == nil {
        log.Printf("TODO: Parameter ID %d does not exist on pedal type %s\n",
            paramID, pbi.GetName())
        return
    }

    var v, vMin, vMax float32
    switch param.(type) {
    case *pedal.TempoParam:
        var tempo uint8
        tempo, *tempos = (*tempos)[0], (*tempos)[1:]
        if tempo > 1 {
            v = float32(tempo)
            break
        }
        binary.Read(bytes.NewReader(data[4:8]), binary.LittleEndian, &v)
    default:
        binary.Read(bytes.NewReader(data[4:8]), binary.LittleEndian, &v)
    }

    binary.Read(bytes.NewReader(data[8:12]), binary.LittleEndian, &vMin)
    binary.Read(bytes.NewReader(data[12:16]), binary.LittleEndian, &vMax)
    binValue := [4]byte{}
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, v)
    copy(binValue[:], buf.Bytes())
    if err := param.SetBinValueCurrent(binValue); err != nil {
        log.Printf("TODO: Fix the parameter type on pedal %s, parameter %s current : %s \n", pbi.GetName(), param.GetName(), err)
    }
    if err := param.SetBinValueMin(binValue); err != nil {
        log.Printf("TODO: Fix the parameter type on pedal %s, parameter %s min: %s \n", pbi.GetName(), param.GetName(), err)
    }
    if err := param.SetBinValueMax(binValue); err != nil {
        log.Printf("TODO: Fix the parameter type on pedal %s, parameter %s max: %s \n", pbi.GetName(), param.GetName(), err)
    }
}

func (m SetChange) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    pb.SetCurrentSet(binary.LittleEndian.Uint32(m.data[8:]))
    return nil, pedal.SetChange, pb
}

func (m SetLoad) Parse(pb *pedal.PedalBoard) (error, pedal.ChangeType, interface{}) {
    pb.SetCurrentSetName(string(m.data[12:]))
    return nil, pedal.SetLoad, nil
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
