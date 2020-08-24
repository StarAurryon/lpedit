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
import "log"
import "reflect"

import "github.com/StarAurryon/lpedit/model/pod"

const (
    messageRead uint32  = 1073743882
    messageWrite uint32 = 134299658
)

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

type IMessage interface {
    Copy() IMessage
    GetData() []byte
    GetType() uint16
    GetSubType() uint16
    IsOk() bool
    LogInfo()
    Parse(*pod.PedalBoard) (error, int, interface{})
    SetData([]byte)
}

type Message struct {
    data   []byte
    mname  string
    msize  int
    mtype  uint16
    smtype uint16
}

func (m *Message) Copy() IMessage {
    _m := new(Message)
    *_m = *m
    return _m
}

func (m *Message) GetData() []byte {
    return m.data
}

func (m *Message) GetType() uint16 { return m.mtype }
func (m *Message) GetSubType() uint16 { return m.smtype }
func (m *Message) IsOk() bool { return m.msize <= len(m.data) }
func (m *Message) SetData(data []byte) { m.data = data }

func (m *Message) Parse(*pod.PedalBoard) (error, int, interface{}) {
    info := fmt.Sprintf("No defined pase fuction for %s message, mtype: %d, smtype %d",
        m.mname, m.mtype, m.smtype)
    return fmt.Errorf(info), ct.StatusWarning(), nil
}


type ActiveChange struct {
    Message
}

func (m *ActiveChange) Copy() IMessage {
    _m := new(ActiveChange)
    *_m = *m
    return _m
}

type TypeChange struct {
    Message
}

func (m *TypeChange) Copy() IMessage {
    _m := new(TypeChange)
    *_m = *m
    return _m
}

type ParameterChange struct {
    Message
}

func (m *ParameterChange) Copy() IMessage {
    _m := new(ParameterChange)
    *_m = *m
    return _m
}

type ParameterChangeMin struct {
    Message
}

func (m *ParameterChangeMin) Copy() IMessage {
    _m := new(ParameterChangeMin)
    *_m = *m
    return _m
}

type ParameterChangeMax struct {
    Message
}

func (m *ParameterChangeMax) Copy() IMessage {
    _m := new(ParameterChangeMax)
    *_m = *m
    return _m
}

type ParameterTempoChange struct {
    Message
}

func (m *ParameterTempoChange) Copy() IMessage {
    _m := new(ParameterTempoChange)
    *_m = *m
    return _m
}

type ParameterTempoChange2 struct {
    Message
}

func (m *ParameterTempoChange2) Copy() IMessage {
    _m := new(ParameterTempoChange2)
    *_m = *m
    return _m
}

type PresetChange struct {
    Message
}

func (m *PresetChange) Copy() IMessage {
    _m := new(PresetChange)
    *_m = *m
    return _m
}

type PresetChangeAlert struct {
    Message
}

func (m *PresetChangeAlert) Copy() IMessage {
    _m := new(PresetChangeAlert)
    *_m = *m
    return _m
}

type PresetLoad struct {
    Message
}

func (m *PresetLoad) Copy() IMessage {
    _m := new(PresetLoad)
    *_m = *m
    return _m
}

type PresetSet struct {
    Message
}

func (m *PresetSet) Copy() IMessage {
    _m := new(PresetSet)
    *_m = *m
    return _m
}

type PresetQuery struct {
    Message
}

func (m *PresetQuery) Copy() IMessage {
    _m := new(PresetQuery)
    *_m = *m
    return _m
}


type SetChange struct {
    Message
}

func (m *SetChange) Copy() IMessage {
    _m := new(SetChange)
    *_m = *m
    return _m
}

type SetLoad struct {
    Message
}

func (m *SetLoad) Copy() IMessage {
    _m := new(SetLoad)
    *_m = *m
    return _m
}

type SetQuery struct {
    Message
}

func (m *SetQuery) Copy() IMessage {
    _m := new(SetQuery)
    *_m = *m
    return _m
}

type SetupChange struct {
    Message
}

func (m *SetupChange) Copy() IMessage {
    _m := new(SetupChange)
    *_m = *m
    return _m
}

var messages = []IMessage{
    &ActiveChange{Message: Message{mtype: 4, smtype: 4864, msize: 20, mname: "Item Active Change"}},
    &TypeChange{Message: Message{mtype: 4, smtype: 4352, msize: 20, mname: "Item Type Change"}},
    &PresetChange{Message: Message{mtype: 2, smtype: 9984, msize: 12, mname: "Preset change"}},
    &PresetChangeAlert{Message: Message{mtype: 1, smtype: 8960, msize: 8, mname: "Alert Preset Change"}},
    &PresetLoad{Message: Message{mtype: 1025, smtype: 256, msize: 4104, mname: "Preset Load"}},
    &PresetSet{Message: Message{mtype: 1026, smtype: 512, msize: 4108, mname: "Preset Set"}},
    &PresetQuery{Message: Message{mtype: 2, smtype: 0, msize: 12, mname: "Preset Query"}},
    &ParameterChange{Message: Message{mtype: 6, smtype: 11520, msize: 28, mname: "Item Parameter Change"}},
    &ParameterChangeMin{Message: Message{mtype: 6, smtype: 11776, msize: 28, mname: "Item Parameter Change Min"}},
    &ParameterChangeMax{Message: Message{mtype: 6, smtype: 12032, msize: 28, mname: "Item Parameter Change Max"}},
    &ParameterTempoChange{Message: Message{mtype: 4, smtype: 5120, msize: 20, mname: "Item Parameter Tempo Change"}},
    &ParameterTempoChange2{Message: Message{mtype: 4, smtype: 12544, msize: 20, mname: "Item Parameter Tempo Change"}},
    &SetChange{Message: Message{mtype: 2, smtype: 11264, msize: 12, mname: "Set Change"}},
    &SetLoad{Message: Message{mtype: 6, smtype: 10496, msize: 12, mname: "Set Query"}},
    &SetQuery{Message: Message{mtype: 2, smtype: 10240, msize: 12, mname: "Set Query"}},
    &SetupChange{Message: Message{mtype: 5, smtype: 5632, msize: 24, mname: "Setup Change"}},
}

func newMessage(mtype uint16, smtype uint16) IMessage {
    for _, m := range messages {
        if m.GetType() == mtype && m.GetSubType() == smtype {
            return m.Copy()
        }
    }
    return nil
}

func newMessage2(mtype reflect.Type) IMessage {
    for _, m := range messages {
        if mtype == reflect.TypeOf(m) {
            return m.Copy()
        }
    }
    return nil
}

func NewMessage(rm RawMessage) (error, IMessage) {
    if rm.mtype != RawMessageBegin {
        return fmt.Errorf("You can't init a message with a non Begin Type"), nil
    }
    if len(rm.data) < 8 {
        return fmt.Errorf("The size of the RawMessage is too small to get the messageType"), nil
    }

    mtype := binary.LittleEndian.Uint16(rm.data[0:2])
    smtype := binary.LittleEndian.Uint16(rm.data[6:8])
    m := newMessage(mtype, smtype)
    if m == nil {
        return nil, &Message{mname: "Unknown", data: rm.data,
            mtype: mtype, smtype: smtype}
    }

    m.SetData(rm.data)
    if !m.IsOk() {
        return fmt.Errorf("The size of the RawMessage is too small\n"), nil
    }
    return nil, m
}

func (m Message) LogInfo() {
    log.Printf("Message info\n")
    log.Printf("Name %s\n", m.mname)
    log.Printf("Data size %d, effective size, %d\n", m.msize, len(m.data))
    log.Printf("Content: %x\n\n", m.data)
}
