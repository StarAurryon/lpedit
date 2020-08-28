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

import "bytes"
import "encoding/binary"
import "math"
import "reflect"

import "github.com/StarAurryon/lpedit/model/pod"

const (
    CurrentSet uint16 = 0xFFFF
    CurrentPreset uint16 = 0xFFFF
)

func genHeader(m IMessage) *bytes.Buffer {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, m.GetType())
    binary.Write(buf, binary.LittleEndian, messageWrite)
    binary.Write(buf, binary.LittleEndian, m.GetSubType())
    return buf
}

func genSetupChange(paramID uint32, vtype uint32, value [4]byte) IMessage {
    var m *SetupChange
    m = newMessage2(reflect.TypeOf(m)).(*SetupChange)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, [4]byte{})
    binary.Write(buf, binary.LittleEndian, vtype)
    binary.Write(buf, binary.LittleEndian, paramID)
    binary.Write(buf, binary.LittleEndian, value)
    m.data = buf.Bytes()

    return m
}

func GenDTClassChange(dt *pod.DT) IMessage {
    var paramID uint32 = 0x28 + (uint32(dt.GetID()) * 3)
    value := [4]byte{dt.GetBinClass()}
    return genSetupChange(paramID, pod.Int32Type, value)
}

func GenDTModeChange(dt *pod.DT) IMessage {
    var paramID uint32 = 0x27 + (uint32(dt.GetID()) * 3)
    value := [4]byte{dt.GetBinMode()}
    return genSetupChange(paramID, pod.Int32Type, value)
}

func GenDTTopologyChange(dt *pod.DT) IMessage {
    var paramID uint32 = 0x26 + (uint32(dt.GetID()) * 3)
    value := [4]byte{dt.GetBinTopology()}
    return genSetupChange(paramID, pod.Int32Type, value)
}

func GenActiveChange(pbi pod.PedalBoardItem) IMessage {
    var m *ActiveChange
    m = newMessage2(reflect.TypeOf(m)).(*ActiveChange)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, uint32(0))
    binary.Write(buf, binary.LittleEndian, pbi.GetID())
    binary.Write(buf, binary.LittleEndian, pbi.GetActive2())
    m.data = buf.Bytes()

    return m
}

func genParameterChange(m IMessage, v [4]byte, p pod.Parameter) IMessage {
    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, uint32(0))
    binary.Write(buf, binary.LittleEndian, p.GetParent().(pod.PedalBoardItem).GetID())
    binary.Write(buf, binary.LittleEndian, p.GetBinValueType())
    id := p.GetID()

    binary.Write(buf, binary.LittleEndian, id)
    binary.Write(buf, binary.LittleEndian, v)
    m.SetData(buf.Bytes())

    return m
}

func GenParameterChange(p pod.Parameter) IMessage {
    var m *ParameterChange
    m = newMessage2(reflect.TypeOf(m)).(*ParameterChange)
    return genParameterChange(m, p.GetBinValueCurrent(), p)
}

func GenParameterCabChange(p pod.Parameter) IMessage {
    var paramID uint32
    cabID, pID := p.GetParent().(pod.PedalBoardItem).GetID()/2, p.GetID()

    switch  {
    case cabID == 0 && pID == pod.CabERID:
        paramID = setupMessageCab0ER
    case cabID == 1 && pID == pod.CabERID:
        paramID = setupMessageCab1ER
    case cabID == 0 && pID == pod.CabMicID:
        paramID = setupMessageCab0Mic
    case cabID == 1 && pID == pod.CabMicID:
        paramID = setupMessageCab1Mic
    case cabID == 0 && pID == pod.CabLowCutID:
        paramID = setupMessageCab0LoCut
    case cabID == 1 && pID == pod.CabLowCutID:
        paramID = setupMessageCab1LoCut
    case cabID == 0 && pID == pod.CabResLevelID:
        paramID = setupMessageCab0ResLvl
    case cabID == 1 && pID == pod.CabResLevelID:
        paramID = setupMessageCab1ResLvl
    case cabID == 0 && pID == pod.CabThumpID:
        paramID = setupMessageCab0Thump
    case cabID == 1 && pID == pod.CabThumpID:
        paramID = setupMessageCab1Thump
    case cabID == 0 && pID == pod.CabDecayID:
        paramID = setupMessageCab0Decay
    case cabID == 1 && pID == pod.CabDecayID:
        paramID = setupMessageCab1Decay
    }
    return genSetupChange(paramID, p.GetBinValueType(), p.GetBinValueCurrent())
}

func GenParameterPedalBoardChange(p pod.Parameter) IMessage {
    var paramID uint32

    switch p.GetID() {
    case pod.PedalBoardInput1Source:
        paramID = setupMessageInput1Source
    case pod.PedalBoardInput2Source:
        paramID = setupMessageInput2Source
    case pod.PedalBoardGuitarInZ:
        paramID = setupMessageGuitarInZ
    }
    return genSetupChange(paramID, p.GetBinValueType(), p.GetBinValueCurrent())
}

func GenParameterChangeMin(p pod.Parameter) IMessage {
    var m *ParameterChangeMin
    m = newMessage2(reflect.TypeOf(m)).(*ParameterChangeMin)
    return genParameterChange(m, p.GetBinValueMin(), p)
}

func GenParameterChangeMax(p pod.Parameter) IMessage {
    var m *ParameterChangeMax
    m = newMessage2(reflect.TypeOf(m)).(*ParameterChangeMax)
    return genParameterChange(m, p.GetBinValueMax(), p)
}

func GenParameterTempoChange(p pod.Parameter) IMessage {
    var m *ParameterTempoChange
    m = newMessage2(reflect.TypeOf(m)).(*ParameterTempoChange)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, uint32(0))
    binary.Write(buf, binary.LittleEndian, p.GetParent().(pod.PedalBoardItem).GetID())
    tmpValue := p.GetBinValueCurrent()
    var binValue float32
    binary.Read(bytes.NewReader(tmpValue[:]), binary.LittleEndian, &binValue)
    if binValue > 1 {
        binary.Write(buf, binary.LittleEndian, uint32(math.Round(float64(binValue))))
    } else {
        binary.Write(buf, binary.LittleEndian, uint32(0))
    }
    m.data = buf.Bytes()

    return m
}

func GenParameterTempoChange2(p pod.Parameter) IMessage {
    var m *ParameterTempoChange2
    m = newMessage2(reflect.TypeOf(m)).(*ParameterTempoChange2)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, uint32(0))
    binary.Write(buf, binary.LittleEndian, p.GetParent().(pod.PedalBoardItem).GetID())
    tmpValue := p.GetBinValueCurrent()
    var binValue float32
    binary.Read(bytes.NewReader(tmpValue[:]), binary.LittleEndian, &binValue)
    if binValue > 1 {
        binary.Write(buf, binary.LittleEndian, uint32(math.Round(float64(binValue))))
    } else {
        binary.Write(buf, binary.LittleEndian, uint32(0))
    }
    m.data = buf.Bytes()

    return m
}

func GenPresetChange(presetID uint8) IMessage {
    var m *PresetChange
    m = newMessage2(reflect.TypeOf(m)).(*PresetChange)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, uint32(presetID))
    m.data = buf.Bytes()

    return m
}

func GenPresetChangeAlert() IMessage {
    var m *PresetChangeAlert
    m = newMessage2(reflect.TypeOf(m)).(*PresetChangeAlert)
    return m
}

func GenPresetLoad() IMessage {
    var m *PresetLoad
    m = newMessage2(reflect.TypeOf(m)).(*PresetLoad)

    return m
}

func GenPresetQuery(presetID uint16, setID uint16) IMessage {
    var m *PresetQuery
    m = newMessage2(reflect.TypeOf(m)).(*PresetQuery)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, presetID)
    binary.Write(buf, binary.LittleEndian, setID)
    m.data = buf.Bytes()

    return m
}

//DirtyHack TODO: Cleanup
func GenPresetSet(pb *pod.PedalBoard, oldMsg *PresetLoad, presetID uint16, setID uint16) IMessage {
    var m *PresetSet
    m = newMessage2(reflect.TypeOf(m)).(*PresetSet)

    data := oldMsg.data

    pbiOrder := []uint32{0,2,1,3,4,5,6,7,8,9,10,11}
    offset := 48

    for _, id := range pbiOrder {
        pbi := pb.GetItem(id)
        pos, posType := pbi.GetPos()
        bPos := make([]byte, 2)
        binary.LittleEndian.PutUint16(bPos, pos)

        data[offset+4] = bPos[0]
        data[offset+5] = bPos[1]
        data[offset+6] = posType

        offset += 256
    }

    //TODO: FIX
    name := pb.GetCurrentPresetName()

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, presetID)
    binary.Write(buf, binary.LittleEndian, setID)
    binary.Write(buf, binary.LittleEndian, name)
    binary.Write(buf, binary.LittleEndian, data[24:])
    m.data = buf.Bytes()

    return m
}

func GenSetChange(setID uint8) IMessage {
    var m *SetChange
    m = newMessage2(reflect.TypeOf(m)).(*SetChange)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, uint32(setID))
    m.data = buf.Bytes()

    return m
}

func GenSetQuery(id uint32) IMessage {
    var m *SetQuery
    m = newMessage2(reflect.TypeOf(m)).(*SetQuery)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, id)
    m.data = buf.Bytes()

    return m
}

func GenSetupChange() IMessage {
    var m *SetupChange
    m = newMessage2(reflect.TypeOf(m)).(*SetupChange)
    return m
}

func GenTypeChange(pbi pod.PedalBoardItem) IMessage {
    var m *TypeChange
    m = newMessage2(reflect.TypeOf(m)).(*TypeChange)

    pbiType := pbi.GetType()
    //Fix for disabled Amp/Cab/Pedal
    if pbiType & 0xFFFF == 0xFFFF {
        pbiType = (pbiType & 0xFFFF) + 0x7FFF0000
    }

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, uint32(0))
    binary.Write(buf, binary.LittleEndian, pbi.GetID())
    binary.Write(buf, binary.LittleEndian, pbiType)
    m.data = buf.Bytes()

    return m
}

func genStatusQuery(statusID uint32) IMessage {
    var m *StatusQuery
    m = newMessage2(reflect.TypeOf(m)).(*StatusQuery)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, statusID)
    m.data = buf.Bytes()

    return m
}

func GenStatusQueryPresetID() IMessage {
    return genStatusQuery(statusIDPreset)
}

func GenStatusQuerySave() IMessage {
    return genStatusQuery(statusSave)
}

func GenStatusQuerySetID() IMessage {
    return genStatusQuery(statusIDSet)
}
