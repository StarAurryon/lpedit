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

import "github.com/StarAurryon/lpedit/pedal"

func genHeader(m IMessage) *bytes.Buffer {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, m.GetType())
    binary.Write(buf, binary.LittleEndian, messageWrite)
    binary.Write(buf, binary.LittleEndian, m.GetSubType())
    return buf
}

func GenActiveChange(pbi pedal.PedalBoardItem) IMessage {
    var m *ActiveChange
    m = newMessage2(reflect.TypeOf(m)).(*ActiveChange)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, uint32(0))
    binary.Write(buf, binary.LittleEndian, pbi.GetID())
    binary.Write(buf, binary.LittleEndian, pbi.GetActive2())
    m.data = buf.Bytes()

    return m
}

func GenParameterChange(p pedal.Parameter) IMessage {
    var m *ParameterChange
    m = newMessage2(reflect.TypeOf(m)).(*ParameterChange)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, uint32(0))
    binary.Write(buf, binary.LittleEndian, p.GetParent().GetID())
    binary.Write(buf, binary.LittleEndian, p.GetBinValueType())
    id := p.GetID()
    idAdd := uint16(0)
    p2, ok := p.(*pedal.PerCentParam)
    if ok {
        idAdd = p2.GetIDAdd()
        id -= idAdd
    }
    //FIXME: Dirty
    if p.GetBinValueType() == pedal.Int32Type {
        id += 6142
    }

    binary.Write(buf, binary.LittleEndian, id)
    if idAdd == 1 {
        binary.Write(buf, binary.LittleEndian, uint8(1))
    } else if id > 6142 {
        binary.Write(buf, binary.LittleEndian, uint8(0))
    } else {
        binary.Write(buf, binary.LittleEndian, uint8(16))
    }
    binary.Write(buf, binary.LittleEndian, uint8(63))
    binary.Write(buf, binary.LittleEndian, p.GetBinValue())
    m.data = buf.Bytes()

    return m
}

func GenParameterTempoChange(p pedal.Parameter) IMessage {
    var m *ParameterTempoChange
    m = newMessage2(reflect.TypeOf(m)).(*ParameterTempoChange)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, uint32(0))
    binary.Write(buf, binary.LittleEndian, p.GetParent().GetID())
    tmpValue := p.GetBinValue()
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

func GenParameterTempoChange2(p pedal.Parameter) IMessage {
    var m *ParameterTempoChange2
    m = newMessage2(reflect.TypeOf(m)).(*ParameterTempoChange2)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, uint32(0))
    binary.Write(buf, binary.LittleEndian, p.GetParent().GetID())
    tmpValue := p.GetBinValue()
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

func GenPresetChange() IMessage {
    var m *PresetChange
    m = newMessage2(reflect.TypeOf(m)).(*PresetChange)
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

func GenSetChange() IMessage {
    var m *SetChange
    m = newMessage2(reflect.TypeOf(m)).(*SetChange)
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

func GenTypeChange(pbi pedal.PedalBoardItem) IMessage {
    var m *TypeChange
    m = newMessage2(reflect.TypeOf(m)).(*TypeChange)

    buf := genHeader(m)
    binary.Write(buf, binary.LittleEndian, uint32(0))
    binary.Write(buf, binary.LittleEndian, pbi.GetID())
    binary.Write(buf, binary.LittleEndian, pbi.GetType())
    m.data = buf.Bytes()

    return m
}
