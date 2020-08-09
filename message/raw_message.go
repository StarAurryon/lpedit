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
    RawMessageBegin uint8 = 1
    RawMessageExt   uint8 = 4
    maxDataSize       int = 60
)

type RawMessage struct {
    size  uint8
    ukno0 uint8
    mtype uint8
    ukno1 uint8
    data  []byte
}

func NewRawMessage(data []byte) *RawMessage {
    var m RawMessage
    m.size = data[0]
    m.ukno0 = data[1]
    m.mtype = data[2]
    m.ukno1 = data[3]
    if len(data) > 4 {
        m.data = data[4:len(data)]
    }
    return &m
}

func NewRawMessages(m IMessage, ukno0 uint8, ukno1 uint8) []*RawMessage {
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
        ret[i] = &RawMessage{size: uint8(len(buf)), mtype: mtype, data: buf,
        ukno0: ukno0, ukno1: ukno1}
        mtype = RawMessageExt
    }
    return ret
}

func (m *RawMessage) Export() []byte {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, m.size)
    binary.Write(buf, binary.LittleEndian, m.ukno0)
    binary.Write(buf, binary.LittleEndian, m.mtype)
    binary.Write(buf, binary.LittleEndian, m.ukno1)
    binary.Write(buf, binary.LittleEndian, m.data)
    return buf.Bytes()
}

func (m *RawMessage) Extend(rm *RawMessage) error{
    if rm.mtype != RawMessageExt {
        return fmt.Errorf("You can't extend a rawMessage with non Ext type")
    }
    if len(rm.data) > 0 {
        m.data = append(m.data, rm.data...)
    }
    return nil
}

func (m *RawMessage) GetType() uint8 {
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
