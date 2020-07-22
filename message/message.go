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

type messageType uint32

type messageInfo struct {
    size  uint16
    stype string
    psinfo *map[subMessageType]subMessageInfo
}

type Message struct {
    data []byte
    info messageInfo
    sinfo subMessageInfo
}

type subMessageType uint32

type subMessageInfo struct {
    stype string
    parse func (Message, *pedal.PedalBoard) error
}

var messages = map[messageType]messageInfo {
    134873092: messageInfo{size: 20, stype: "PedalBoardItem",
        psinfo: &map[subMessageType]subMessageInfo {
            285229056: subMessageInfo{stype: "Type change", parse: itemTypeChange},
            318783488: subMessageInfo{stype: "Active change", parse: itemActiveChange},
        }},
    134874113: messageInfo{size: 4104, stype: "Preset",
        psinfo: &map[subMessageType]subMessageInfo {
            16793600:  subMessageInfo{stype:"Load", parse: loadPreset},
        }},
    134873089: messageInfo{size: 8, stype: "Prepare preset change?",
        psinfo: &map[subMessageType]subMessageInfo {
            587218944: subMessageInfo{stype:"Event", parse: nil},
        }},
    134873090: messageInfo{size: 12, stype: "Prepare preset change?",
        psinfo: &map[subMessageType]subMessageInfo {
            738213888: subMessageInfo{stype:"Parameter 1", parse: nil},
            654327808: subMessageInfo{stype:"Parameter 3", parse: nil},
        }},
    134873094: messageInfo{size: 28, stype: "Item parameter",
        psinfo: &map[subMessageType]subMessageInfo {
            754991104: subMessageInfo{stype: "Parameter change", parse: itemParameterChange},
            771768320: subMessageInfo{stype: "Parameter change 2 ", parse: nil},
            788545536: subMessageInfo{stype: "Parameter change 3 ", parse: nil},
        }},
    134873093: messageInfo{size: 24, stype: "Setup",
        psinfo: &map[subMessageType]subMessageInfo {
        }},
}

func NewMessage(rm RawMessage) (error, *Message) {
    if rm.mtype != rawMessageTypeBegin {
        return fmt.Errorf("You can't init a message with a non Begin Type"), nil
    }
    v, found := messages[rm.getMessageType()]
    if !found {
        return fmt.Errorf("Message type is unknown, code: %d",
             rm.getMessageType()), nil
    }
    m := &Message{data: rm.data, info: v}
    if !m.Ready() {
        return nil, m
    }
    return m.fillSubMessageInfo(), m
}

func (m *Message) Extend(rm RawMessage) error{
    if rm.mtype != rawMessageTypeExt {
        return fmt.Errorf("You can't extend a message with non Ext type")
    }
    m.data = append(m.data, rm.data...)

    if !m.Ready() {
        return nil
    }
    return m.fillSubMessageInfo()
}

func (m *Message) fillSubMessageInfo() error {
    if !m.Ready() {
        return nil
    }
    smtype := subMessageType(binary.LittleEndian.Uint32(m.data[4:8]))
    v, found := (*m.info.psinfo)[smtype]
    if found {
        m.sinfo = v
        return nil
    } else {
        m.sinfo = subMessageInfo{stype: "Unknown"}
        return fmt.Errorf("Type \"%s\": Message subtype is unknwon, code: %d",
             m.info.stype, smtype)
    }
}

func (m Message) Parse(pb *pedal.PedalBoard) error{
    if m.sinfo.parse == nil {
        return fmt.Errorf("No parse function defined to handle the message\n")
    }
    return m.sinfo.parse(m, pb)
}

func (m Message) PrintInfo() {
    fmt.Printf("Message info\n")
    fmt.Printf("Type %s\n", m.info.stype)
    fmt.Printf("SubType %s\n", m.sinfo.stype)
    fmt.Printf("Data size %d, effective size, %d\n", m.info.size, len(m.data))
    fmt.Printf("Content: %x\n\n", m.data)
}

func (m *Message) Ready() bool {
    return len(m.data) >= int(m.info.size)
}
