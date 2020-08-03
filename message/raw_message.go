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
import "fmt"
import "log"

const (
    RawMessageBegin uint16 = 1
    RawMessageExt   uint16 = 4
    maxDataSize       int = 60
)

type RawMessage struct {
    size  uint16
    mtype uint16
    data  []byte
}

func NewRawMessage(data []byte) *RawMessage {
    var m RawMessage
    m.size = binary.LittleEndian.Uint16(data[:2])
    m.mtype = binary.LittleEndian.Uint16(data[2:4])
    if len(data) > 4 {
        m.data = data[4:len(data)]
    }
    return &m
}

func NewRawMessages(m IMessage) []*RawMessage {
    var buf []byte
    data := m.getData()
    size := len(data) / maxDataSize
    if len(data) % maxDataSize != 0 {
        size ++
    }
    ret := make([]*RawMessage, size)
    mtype := RawMessageBegin
    for i := range ret {
        offset := i * maxDataSize
        if (offset + maxDataSize) < len(data) {
            buf = data[offset:offset+maxDataSize]
        } else {
            buf = data[offset:]
        }
        ret[i] = &RawMessage{size: uint16(len(buf)), mtype: mtype, data: buf}
        mtype = RawMessageExt
    }
    return ret
}

func (m *RawMessage) Export() []byte {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, m.size)
    binary.Write(buf, binary.LittleEndian, m.mtype)
    binary.Write(buf, binary.LittleEndian, m.data)
    return buf.Bytes()
}

func (m *RawMessage) Extend(rm *RawMessage) error{
    if rm.mtype != RawMessageExt {
        return fmt.Errorf("You can't extend a rawMessage with non Ext type")
    }
    m.data = append(m.data, rm.data...)
    return nil
}

func (m *RawMessage) GetType() uint16 {
    return m.mtype
}

func (m RawMessage) LogInfo() {
    var mtype = ""
    log.Printf("RawMessage info\n")
    switch m.mtype {
    case RawMessageBegin:
        mtype = "BEGIN"
    case RawMessageExt:
        mtype = "EXT"
    default:
        mtype = "Unknown"
    }
    log.Printf("Type %s\n", mtype)
    log.Printf("Data size %d, effective size, %d\n", m.size, len(m.data))
    log.Printf("Content: %x\n\n", m.data)
}
