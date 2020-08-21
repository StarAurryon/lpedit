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

const (
    setupMessageTempo uint32 = 0x17
    setupMessageCab0ER uint32 = 0x32
    setupMessageCab1ER uint32 = 0x33
    setupMessageCab0Mic uint32 = 0x34
    setupMessageCab1Mic uint32 = 0x35
    setupMessageInput1Source uint32 = 0x36
    setupMessageInput2Source uint32 = 0x37
    setupMessageGuitarInZ uint32 = 0x55
    setupMessageCab0LoCut uint32 = 0x57
    setupMessageCab1LoCut uint32 = 0x58
    setupMessageCab0ResLvl uint32 = 0x59
    setupMessageCab1ResLvl uint32 = 0x5a
    setupMessageCab0Thump uint32 = 0x5b
    setupMessageCab1Thump uint32 = 0x5c
    setupMessageCab0Decay uint32 = 0x5d
    setupMessageCab1Decay uint32 = 0x5e
)

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
    m.parseDT(pb, m.data)
    m.parseCabs(pb, m.data)
    m.parseSetup(pb, m.data)
    return nil, pedal.PresetLoad, pb
}

func (m PresetLoad) parseCabs(pb *pedal.PedalBoard, data []byte) {
    cabs := []*pedal.Cab {pb.GetCab(0), pb.GetCab(1)}
    offset := [2][2]int{[2]int{3412, 4096}, [2]int{3420, 4097}}
    parametersID := [2]uint32{pedal.CabERID, pedal.CabMicID}
    parametersSize := [2]int{4, 1}
    for i, cab := range cabs {
        if cab == nil {
            log.Printf("Can't find Cab ID %d\n", i)
            continue
        }
        for j, pType := range parametersID {
            p := cab.GetParam(pType)
            if p == nil {
                log.Printf("Can't find Cab ID %d, parameter %d\n", i, pType)
                continue
            }
            value := [4]byte{}
            copy(value[:], data[offset[i][j]:offset[i][j]+parametersSize[j]])
            if err := p.SetBinValueCurrent(value); err != nil {
                log.Printf("Can't set value Cab ID %d, parameter %d: %s\n", i, pType, err)
            }
        }
    }
}

func (m PresetLoad) parseDT(pb *pedal.PedalBoard, data []byte) {
    dts := []*pedal.DT{pb.GetDT(0), pb.GetDT(1)}
    offset := [2][3]int{[3]int{3124,3125,3126}, [3]int{3132, 3133, 3134}}
    for i, dt := range dts {
        if dt == nil {
            log.Printf("Can't find DT ID %d\n", i)
        } else {
            if err := dt.SetBinTopology(data[offset[i][0]]); err != nil {
                log.Printf("Error while setting DT ID %d Topology: %s\n", i, err)
            }
            if err := dt.SetBinClass(data[offset[i][1]]); err != nil {
                log.Printf("Error while setting DT ID %d Class: %s\n", i, err)
            }
            if err := dt.SetBinMode(data[offset[i][2]]); err != nil {
                log.Printf("Error while setting DT ID %d Mode: %s\n", i, err)
            }
        }
    }
}

func (m PresetLoad) parseSetup(pb *pedal.PedalBoard, data []byte) {
    params := []uint32{pedal.PedalBoardGuitarInZ, pedal.PedalBoardInput1Source,
        pedal.PedalBoardInput2Source}
    offset := []int{3546, 4102, 4103}
    for i, pType := range params {
        p := pb.GetParam(pType)
        if p == nil {
            log.Printf("Can't find PedalBoard Parameter ID %d\n", pType)
            continue
        }
        value := [4]byte{}
        value[0] = data[offset[i]]
        if err := p.SetBinValueCurrent(value); err != nil {
            log.Printf("Error while setting PedalBoard Parameter ID %d: %s\n", pType, err)
        }
    }
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
    switch pbi.(type) {
    case *pedal.Cab:
        for i, pType := range []uint32{pedal.CabLowCutID, pedal.CabResLevelID,
            pedal.CabThumpID, pedal.CabDecayID} {
            start := offset + (i * 20)
            end := start + 20
            copy(paramData[:], data[start:end])
            m.parseParameterCab(pbi, paramData, pType)
        }
    default:
        for i := uint16(0); i < pbi.GetParamLen(); i++ {
            start := offset + (i * 20)
            end := start + 20
            copy(paramData[:], data[start:end])
            m.parseParameterNormal(pbi, paramData, &tempos)
        }
    }
}

func (m PresetLoad) parseParameterCab(pbi pedal.PedalBoardItem, data [20]byte, paramID uint32) {
    param := pbi.GetParam(paramID)
    if param == nil {
        log.Printf("TODO: Parameter ID %d does not exist on item type %s\n",
            paramID, pbi.GetName())
        return
    }
    binValue := [4]byte{}
    copy(binValue[:], data[4:8])
    if err := param.SetBinValueCurrent(binValue); err != nil {
        log.Printf("TODO: Fix the parameter type on pedal %s, parameter %s current : %s \n", pbi.GetName(), param.GetName(), err)
    }
}

func (m PresetLoad) parseParameterNormal(pbi pedal.PedalBoardItem, data [20]byte, tempos *[]uint8) {
    paramID := binary.LittleEndian.Uint32(data[0:4])
    param := pbi.GetParam(paramID)
    if param == nil {
        log.Printf("TODO: Parameter ID %d does not exist on pedal type %s\n",
            paramID, pbi.GetName())
        return
    }

    var v float32
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

    binValue := [4]byte{}
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, v)
    copy(binValue[:], buf.Bytes())
    if err := param.SetBinValueCurrent(binValue); err != nil {
        log.Printf("TODO: Fix the parameter type on pedal %s, parameter %s current : %s \n", pbi.GetName(), param.GetName(), err)
    }
    copy(binValue[:], data[8:12])
    if err := param.SetBinValueMin(binValue); err != nil {
        log.Printf("TODO: Fix the parameter type on pedal %s, parameter %s min: %s \n", pbi.GetName(), param.GetName(), err)
    }
    copy(binValue[:], data[12:16])
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
    setupType := binary.LittleEndian.Uint32(m.data[16:20])
    var value [4]byte
    copy(value[:], m.data[20:])

    switch setupType {
    case setupMessageTempo:
        return m.parseTempo(pb, value)
    case setupMessageCab0ER:
        return m.parseCab(pb, 0, pedal.CabERID, value)
    case setupMessageCab1ER:
        return m.parseCab(pb, 1, pedal.CabERID, value)
    case setupMessageCab0Mic:
        return m.parseCab(pb, 0, pedal.CabMicID, value)
    case setupMessageCab1Mic:
        return m.parseCab(pb, 1, pedal.CabMicID, value)
    case setupMessageCab0LoCut:
        return m.parseCab(pb, 0, pedal.CabLowCutID, value)
    case setupMessageCab1LoCut:
        return m.parseCab(pb, 1, pedal.CabLowCutID, value)
    case setupMessageCab0ResLvl:
        return m.parseCab(pb, 0, pedal.CabResLevelID, value)
    case setupMessageCab1ResLvl:
        return m.parseCab(pb, 1, pedal.CabResLevelID, value)
    case setupMessageCab0Thump:
        return m.parseCab(pb, 0, pedal.CabThumpID, value)
    case setupMessageCab1Thump:
        return m.parseCab(pb, 1, pedal.CabThumpID, value)
    case setupMessageCab0Decay:
        return m.parseCab(pb, 0, pedal.CabDecayID, value)
    case setupMessageCab1Decay:
        return m.parseCab(pb, 1, pedal.CabDecayID, value)
    case setupMessageInput1Source:
        return m.parsePedalBoard(pb, pedal.PedalBoardInput1Source, value)
    case setupMessageInput2Source:
        return m.parsePedalBoard(pb, pedal.PedalBoardInput2Source, value)
    case setupMessageGuitarInZ:
        return m.parsePedalBoard(pb, pedal.PedalBoardGuitarInZ, value)
    }

    return nil, pedal.None, nil
}

func (m SetupChange) parseCab(pb *pedal.PedalBoard, ID int, paramID uint32, value [4]byte) (error, pedal.ChangeType, interface{}) {
    c := pb.GetCab(ID)
    if c == nil {
        return fmt.Errorf("Can't find Cab %d", ID), pedal.Warning, nil
    }
    p := c.GetParam(paramID)
    if p == nil {
        return fmt.Errorf("Can't get param %d, for Cab %d", paramID, ID), pedal.Warning, nil
    }
    if err := p.SetBinValueCurrent(value); err != nil {
        return fmt.Errorf("Cant set Cab ID %d parameter ID %d value: %s", ID, paramID, err), pedal.Warning, nil
    }
    return nil, pedal.ParameterChange, p
}

func (m SetupChange) parsePedalBoard(pb *pedal.PedalBoard, parameterID uint32, value [4]byte) (error, pedal.ChangeType, interface{}) {
    p := pb.GetParam(parameterID)
    if p == nil {
        return fmt.Errorf("Can't get PedalBoard parameter ID %d", parameterID), pedal.Warning, nil
    }
    if err := p.SetBinValueCurrent(value); err != nil {
        return fmt.Errorf("Cant set PedalBoard parameter ID %d value: %s", parameterID, err), pedal.Warning, nil
    }
    return nil, pedal.ParameterChange, p
}

func (m SetupChange) parseTempo(pb *pedal.PedalBoard, value [4]byte) (error, pedal.ChangeType, interface{}) {
    var v float32
    err := binary.Read(bytes.NewReader(value[:]), binary.LittleEndian, &v)
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
