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

import "github.com/StarAurryon/lpedit/model/pod"

type presetPedalPos struct {
    pid   uint32
    ptype uint8
}

func (m Message) getPedalBoardItemID() uint32 {
    return binary.LittleEndian.Uint32(m.data[12:16])
}

func (m ActiveChange) Parse(pb *pod.PedalBoard) (error, int, interface{}) {
    id := m.getPedalBoardItemID()
    p := pb.GetItem(id)
    if p == nil {
        return fmt.Errorf("Item ID %d not found", id), ct.StatusWarning(), nil
    }
    var active bool
    if binary.LittleEndian.Uint32(m.data[16:]) > 0 {
        active = true
    } else {
        active = false
    }
    log.Printf("Active change on ID %d status %t\n", id, active)
    p.SetActive(active)
    return nil, ct.StatusActiveChange(), p
}

func (m *Message) parseParameterChange(paramFunc string, pb *pod.PedalBoard) (error, int, interface{}) {
    pid := m.getPedalBoardItemID()
    p := pb.GetItem(pid)
    if p == nil {
        return fmt.Errorf("Item ID %d not found", pid), ct.StatusWarning(), nil
    }
    id := binary.LittleEndian.Uint32(m.data[20:24])

    var v [4]byte
    copy(v[:], m.data[24:])

    param := p.GetParam(id)
    if param == nil {
        return fmt.Errorf("Parameter ID %d not found", id), ct.StatusWarning(), nil
    }
    if err := reflect.ValueOf(param).MethodByName(paramFunc).Interface().(func([4]byte) error)(v); err != nil {
        log.Printf("TODO: Fix the parameter type on pod.%s, parameter %s, func %s: %s \n", p.GetName(), param.GetName(), paramFunc, err)
    }
    return nil, ct.StatusNone(), param
}

func (m ParameterChange) Parse(pb *pod.PedalBoard) (error, int, interface{}) {
    err, pt, obj := m.parseParameterChange("SetBinValueCurrent", pb)
    if err != nil { return err, pt, obj }
    return err, ct.StatusParameterChange(), obj
}

func (m ParameterChangeMin) Parse(pb *pod.PedalBoard) (error, int, interface{}) {
    err, pt, obj := m.parseParameterChange("SetBinValueMin", pb)
    if err != nil { return err, pt, obj }
    return err, ct.StatusParameterChangeMin(), obj
}

func (m ParameterChangeMax) Parse(pb *pod.PedalBoard) (error, int, interface{}) {
    err, pt, obj := m.parseParameterChange("SetBinValueMax", pb)
    if err != nil { return err, pt, obj }
    return err, ct.StatusParameterChangeMax(), obj
}

func (m ParameterTempoChange) Parse(pb *pod.PedalBoard) (error, int, interface{}) {
    pid := m.getPedalBoardItemID()
    p := pb.GetItem(pid)
    if p == nil {
        return fmt.Errorf("Item ID %d not found", pid), ct.StatusWarning(), nil
    }
    param := p.GetParam(0)
    if param == nil {
        return fmt.Errorf("Parameter ID 0 not found"), ct.StatusWarning(), nil
    }
    value := float32(binary.LittleEndian.Uint32(m.data[16:]))
    binValue := [4]byte{}
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, value)
    copy(binValue[:], buf.Bytes())
    param.SetBinValueCurrent(binValue)
    return nil, ct.StatusParameterChange(), param
}

func (m ParameterTempoChange2) Parse(pb *pod.PedalBoard) (error, int, interface{}) {
    pid := m.getPedalBoardItemID()
    p := pb.GetItem(pid)
    if p == nil {
        return fmt.Errorf("Item ID %d not found", pid), ct.StatusWarning(), nil
    }
    param := p.GetParam(2)
    if param == nil {
        return fmt.Errorf("Parameter ID 2 not found"), ct.StatusWarning(), nil
    }
    value := float32(binary.LittleEndian.Uint32(m.data[16:]))
    binValue := [4]byte{}
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, value)
    copy(binValue[:], buf.Bytes())
    param.SetBinValueCurrent(binValue)
    return nil, ct.StatusParameterChange(), param
}

func (m PresetChange) Parse(pb *pod.PedalBoard) (error, int, interface{}) {
    pb.SetCurrentPreset(binary.LittleEndian.Uint32(m.data[8:]))
    return nil, ct.StatusPresetChange(), pb
}

func (m PresetChangeAlert) Parse(pb *pod.PedalBoard) (error, int, interface{}) {
    return nil, ct.StatusNone(), nil
}

func (m PresetLoad) Parse(pb *pod.PedalBoard) (error, int, interface{}) {
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
    return nil, ct.StatusPresetLoad(), pb
}

func (m PresetLoad) parseCabs(pb *pod.PedalBoard, data []byte) {
    cabs := []*pod.Cab {pb.GetCab(0), pb.GetCab(1)}
    offset := [2][2]int{[2]int{3412, 4096}, [2]int{3420, 4097}}
    parametersID := [2]uint32{pod.CabERID, pod.CabMicID}
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

func (m PresetLoad) parseDT(pb *pod.PedalBoard, data []byte) {
    dts := []*pod.DT{pb.GetDT(0), pb.GetDT(1)}
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

func (m PresetLoad) parseSetup(pb *pod.PedalBoard, data []byte) {
    params := []uint32{pod.PedalBoardGuitarInZ, pod.PedalBoardInput1Source,
        pod.PedalBoardInput2Source}
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

func (m PresetLoad) parsePedalBoardItem(pb *pod.PedalBoard, data [256]byte, pbiID uint32) {
    pbi := pb.GetItem(pbiID)

    pbiType := binary.LittleEndian.Uint32(data[0:4])
    pbi.SetType(pbiType)

    pos := binary.LittleEndian.Uint16(data[4:6])
    posType := uint8(data[6])
    pbi.SetPosWithoutCheck(pos, posType)

    active := false
    if data[8] == 1 { active = true }
    pbi.SetActive(active)

    tempos := []uint8{data[9], data[10]}

    const offset = 16
    var paramData [20]byte
    switch pbi.(type) {
    case *pod.Cab:
        for i, pType := range []uint32{pod.CabLowCutID, pod.CabResLevelID,
            pod.CabThumpID, pod.CabDecayID} {
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

func (m PresetLoad) parseParameterCab(pbi pod.PedalBoardItem, data [20]byte, paramID uint32) {
    param := pbi.GetParam(paramID)
    if param == nil {
        log.Printf("TODO: Parameter ID %d does not exist on item type %s\n",
            paramID, pbi.GetName())
        return
    }
    binValue := [4]byte{}
    copy(binValue[:], data[4:8])
    if err := param.SetBinValueCurrent(binValue); err != nil {
        log.Printf("TODO: Fix the parameter type on pod.%s, parameter %s current : %s \n", pbi.GetName(), param.GetName(), err)
    }
}

func (m PresetLoad) parseParameterNormal(pbi pod.PedalBoardItem, data [20]byte, tempos *[]uint8) {
    paramID := binary.LittleEndian.Uint32(data[0:4])
    param := pbi.GetParam(paramID)
    if param == nil {
        log.Printf("TODO: Parameter ID %d does not exist on pod.type %s\n",
            paramID, pbi.GetName())
        return
    }

    var v float32
    switch param.(type) {
    case *pod.TempoParam:
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
        log.Printf("TODO: Fix the parameter type on pod.%s, parameter %s current : %s \n", pbi.GetName(), param.GetName(), err)
    }
    copy(binValue[:], data[8:12])
    if err := param.SetBinValueMin(binValue); err != nil {
        log.Printf("TODO: Fix the parameter type on pod.%s, parameter %s min: %s \n", pbi.GetName(), param.GetName(), err)
    }
    copy(binValue[:], data[12:16])
    if err := param.SetBinValueMax(binValue); err != nil {
        log.Printf("TODO: Fix the parameter type on pod.%s, parameter %s max: %s \n", pbi.GetName(), param.GetName(), err)
    }
}

func (m SetChange) Parse(pb *pod.PedalBoard) (error, int, interface{}) {
    pb.SetCurrentSet(binary.LittleEndian.Uint32(m.data[8:]))
    return nil, ct.StatusSetChange(), pb
}

func (m SetLoad) Parse(pb *pod.PedalBoard) (error, int, interface{}) {
    pb.SetCurrentSetName(string(m.data[12:]))
    return nil, ct.StatusSetLoad(), nil
}

func (m SetupChange) Parse(pb *pod.PedalBoard) (error, int, interface{}) {
    setupType := binary.LittleEndian.Uint32(m.data[16:20])
    var value [4]byte
    copy(value[:], m.data[20:])

    switch setupType {
    case setupMessageTempo:
        return m.parseTempo(pb, value)
    case setupMessageCab0ER:
        return m.parseCab(pb, 0, pod.CabERID, value)
    case setupMessageCab1ER:
        return m.parseCab(pb, 1, pod.CabERID, value)
    case setupMessageCab0Mic:
        return m.parseCab(pb, 0, pod.CabMicID, value)
    case setupMessageCab1Mic:
        return m.parseCab(pb, 1, pod.CabMicID, value)
    case setupMessageCab0LoCut:
        return m.parseCab(pb, 0, pod.CabLowCutID, value)
    case setupMessageCab1LoCut:
        return m.parseCab(pb, 1, pod.CabLowCutID, value)
    case setupMessageCab0ResLvl:
        return m.parseCab(pb, 0, pod.CabResLevelID, value)
    case setupMessageCab1ResLvl:
        return m.parseCab(pb, 1, pod.CabResLevelID, value)
    case setupMessageCab0Thump:
        return m.parseCab(pb, 0, pod.CabThumpID, value)
    case setupMessageCab1Thump:
        return m.parseCab(pb, 1, pod.CabThumpID, value)
    case setupMessageCab0Decay:
        return m.parseCab(pb, 0, pod.CabDecayID, value)
    case setupMessageCab1Decay:
        return m.parseCab(pb, 1, pod.CabDecayID, value)
    case setupMessageInput1Source:
        return m.parsePedalBoard(pb, pod.PedalBoardInput1Source, value)
    case setupMessageInput2Source:
        return m.parsePedalBoard(pb, pod.PedalBoardInput2Source, value)
    case setupMessageGuitarInZ:
        return m.parsePedalBoard(pb, pod.PedalBoardGuitarInZ, value)
    }

    return nil, ct.StatusNone(), nil
}

func (m SetupChange) parseCab(pb *pod.PedalBoard, ID int, paramID uint32, value [4]byte) (error, int, interface{}) {
    c := pb.GetCab(ID)
    if c == nil {
        return fmt.Errorf("Can't find Cab %d", ID), ct.StatusWarning(), nil
    }
    p := c.GetParam(paramID)
    if p == nil {
        return fmt.Errorf("Can't get param %d, for Cab %d", paramID, ID), ct.StatusWarning(), nil
    }
    if err := p.SetBinValueCurrent(value); err != nil {
        return fmt.Errorf("Cant set Cab ID %d parameter ID %d value: %s", ID, paramID, err), ct.StatusWarning(), nil
    }
    return nil, ct.StatusParameterChange(), p
}

func (m SetupChange) parsePedalBoard(pb *pod.PedalBoard, parameterID uint32, value [4]byte) (error, int, interface{}) {
    p := pb.GetParam(parameterID)
    if p == nil {
        return fmt.Errorf("Can't get PedalBoard parameter ID %d", parameterID), ct.StatusWarning(), nil
    }
    if err := p.SetBinValueCurrent(value); err != nil {
        return fmt.Errorf("Cant set PedalBoard parameter ID %d value: %s", parameterID, err), ct.StatusWarning(), nil
    }
    return nil, ct.StatusParameterChange(), p
}

func (m SetupChange) parseTempo(pb *pod.PedalBoard, value [4]byte) (error, int, interface{}) {
    var v float32
    err := binary.Read(bytes.NewReader(value[:]), binary.LittleEndian, &v)
    if err != nil {
        return err, ct.StatusWarning(), nil
    }
    pb.SetTempo(v)
    return nil, ct.StatusTempoChange(), pb
}

func (m TypeChange) Parse(pb *pod.PedalBoard) (error, int, interface{}) {
    id := m.getPedalBoardItemID()
    p := pb.GetItem(id)
    if p == nil {
        return fmt.Errorf("Item ID %d not found", id), ct.StatusWarning(), nil
    }
    ptype := binary.LittleEndian.Uint32(m.data[16:])
    if err := p.SetType(ptype); err != nil {
        return err, ct.StatusWarning(), nil
    }
    return nil, ct.StatusTypeChange(), p
}
