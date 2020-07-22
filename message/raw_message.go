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

type rawMessageType uint16
const (
    rawMessageTypeBegin rawMessageType = 1
    rawMessageTypeExt   rawMessageType = 4
)

type RawMessage struct {
    size uint16
    mtype rawMessageType
    data []byte
}

func NewRawMessage(data []byte) RawMessage {
    var m RawMessage
    m.size = binary.LittleEndian.Uint16(data[:2])
    m.mtype = rawMessageType(binary.LittleEndian.Uint16(data[2:4]))
    if len(data) > 4 {
        m.data = data[4:len(data)]
    }
    return m
}

func (m RawMessage) getMessageType() messageType {
    if(m.mtype != rawMessageTypeBegin) {
        return 0
    }
    return messageType(binary.LittleEndian.Uint32(m.data[0:4]))
}

func (m RawMessage) PrintInfo() {
    var mtype = ""
    fmt.Printf("RawMessage info\n")
    switch m.mtype {
    case rawMessageTypeBegin:
        mtype = "BEGIN"
    case rawMessageTypeExt:
        mtype = "EXT"
    default:
        mtype = "Unknown"
    }
    fmt.Printf("Type %s\n", mtype)
    fmt.Printf("Data size %d, effective size, %d\n", m.size, len(m.data))
    fmt.Printf("Content: %x\n\n", m.data)
}
