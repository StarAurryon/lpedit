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

package controller

import "bytes"
import "encoding/binary"

import "github.com/StarAurryon/lpedit/pedal"
import "github.com/StarAurryon/lpedit/message"


func (c *Controller) QueryCurrentPreset() {
    if !c.started { return }
    f := func() {
        m := message.GenPresetQuery(uint16(0xFFFF), uint16(0xFFFF))
        c.writeMessage(m, 0, 0)
    }
    go f()
}

func (c *Controller) QueryAllPresets() {
    if !c.started { return }
    f := func() {
        c.syncPreset = true
        max := pedal.NumberSet * pedal.PresetPerSet
        pb := c.GetPedalBoard()
        for i := 0; i < pedal.NumberSet; i++ {
            pb.LockData()
            pb.SetCurrentSet(uint32(i))
            pb.UnlockData()
            for j := 0; j < pedal.PresetPerSet; j++ {
                pb.LockData()
                pb.SetCurrentPreset(uint32(j))
                pb.UnlockData()
                m := message.GenPresetQuery(uint16(j), uint16(i))
                c.writeMessage(m, 0, 0)
                <- c.presetLoaded
                progress := (((i * pedal.PresetPerSet) + (j + 1)) * 100) / max
                c.notify(nil, pedal.PresetLoadProgress, progress)
            }
        }
        c.syncPreset = false
    }
    go f()
}

func (c *Controller) QueryAllSets() {
    if !c.started { return }
    f := func() {
        pb := c.GetPedalBoard()
        c.syncPreset = true
        for i := 0; i < pedal.NumberSet; i++ {
            pb.LockData()
            pb.SetCurrentSet(uint32(i))
            pb.UnlockData()
            m := message.GenSetQuery(uint32(i))
            c.writeMessage(m, 0, 0)
            <- c.presetLoaded
            progress := ((i + 1) * 100) / pedal.NumberSet
            c.notify(nil, pedal.SetLoadProgress, progress)
        }
        c.syncPreset = false
    }
    go f()
}

func (c *Controller) SetAmpParameterValue(id uint32, pid uint16, value string) error {
    return c.SetPedalBoardItemParameterValue(id*2, pid, value)
}

func (c *Controller) SetPedalParameterValue(id uint32, pid uint16, value string) error {
    return c.SetPedalBoardItemParameterValue(id+4, pid, value)
}

func (c *Controller) SetPedalBoardItemParameterValue(id uint32, pid uint16, value string) error {
    if !c.started { return nil }
    c.pb.LockData()
    pbi := c.pb.GetItem(id);
    if pbi == nil {
        c.pb.UnlockData()
        return nil
    }
    p := pbi.GetParam(pid)
    if p == nil {
        c.pb.UnlockData()
        return nil
    }
    err := p.SetValue(value)
    if err != nil {
        c.pb.UnlockData()
        return err
    }
    switch p2 := p.(type) {
    case *pedal.TempoParam:
        switch p2.GetID() {
        case 0:
            m := message.GenParameterTempoChange(p2)
            c.writeMessage(m, 0, 0)
        case 2:
            m := message.GenParameterTempoChange2(p2)
            c.writeMessage(m, 0, 0)
        }
        binValue := p2.GetBinValue()
        var value float32
        binary.Read(bytes.NewReader(binValue[:]), binary.LittleEndian, &value)
        if value <= 1 {
            m := message.GenParameterChange(p)
            c.writeMessage(m, 0, 0)
        }
    default:
        m := message.GenParameterChange(p)
        go c.writeMessage(m, 0, 0)
    }
    c.pb.UnlockData()
    return nil
}

func (c *Controller) SetPedalActive(id uint32, active bool) {
    c.SetPedalBoardItemActive(id+4, active)
}

func (c *Controller) SetPedalBoardItemActive(id uint32, active bool) {
    f := func() {
        if !c.started { return }
        c.pb.LockData()
        pbi := c.pb.GetItem(id)
        if pbi == nil {
            c.pb.UnlockData()
            return
        }
        pbi.SetActive(active)
        m := message.GenActiveChange(pbi)
        c.pb.UnlockData()
        c.writeMessage(m, 0, 0)
    }
    go f()
}

func (c *Controller) SetAmpType(id uint32, name string) {
    c.SetPedalBoardItemType(id*2, name, "")
}

func (c *Controller) SetPedalType(id uint32, fxType string, fxModel string) {
    c.SetPedalBoardItemType(id+4, fxType, fxModel)
}

func (c *Controller) SetPedalBoardItemType(id uint32, fxType string, fxModel string) {
    if !c.started { return }
    c.pb.LockData()
    pbi := c.pb.GetItem(id)
    if pbi == nil {
        c.pb.UnlockData()
        return
    }
    pbi.SetType2(fxType, fxModel)
    m := message.GenTypeChange(pbi)
    m2 := message.GenPresetQuery(uint16(0xFFFF), uint16(0xFFFF))
    c.pb.UnlockData()
    go c.writeMessage(m, 0, 0)
    //go c.writeMessage(m2, 0x62, 0x71)
    go c.writeMessage(m2, 0, 0)
}
