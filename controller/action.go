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
import "fmt"

import "github.com/StarAurryon/lpedit/model/pod"
import "github.com/StarAurryon/lpedit/model/pod/message"


func (c *Controller) QueryCurrentPreset() {
    if !c.started { return }
    f := func() {
        m := message.GenPresetQuery(uint16(0xFFFF), uint16(0xFFFF))
        c.writeMessage(m, 0, 0)
    }
    go f()
}

func (c *Controller) QueryAllPresets(async bool) {
    if !c.started { return }
    f := func() {
        c.syncMode = true
        max := pod.NumberSet * pod.PresetPerSet
        pb := c.GetPedalBoard()
        for i := 0; i < pod.NumberSet; i++ {
            pb.LockData()
            pb.SetCurrentSet(uint32(i))
            pb.UnlockData()
            for j := 0; j < pod.PresetPerSet; j++ {
                pb.LockData()
                pb.SetCurrentPreset(uint32(j))
                pb.UnlockData()
                m := message.GenPresetQuery(uint16(j), uint16(i))
                c.writeMessage(m, 0, 0)
                <- c.syncModeChan
                progress := (((i * pod.PresetPerSet) + (j + 1)) * 100) / max
                c.notify(nil, sg.StatusProgress(), progress)
            }
        }
        c.syncMode = false
    }
    if async {
        go f()
    } else {
        f()
    }
}

func (c *Controller) QueryAllSets(async bool) {
    if !c.started { return }
    f := func() {
        pb := c.GetPedalBoard()
        c.syncMode = true
        for i := 0; i < pod.NumberSet; i++ {
            pb.LockData()
            pb.SetCurrentSet(uint32(i))
            pb.UnlockData()
            m := message.GenSetQuery(uint32(i))
            c.writeMessage(m, 0, 0)
            <- c.syncModeChan
            progress := ((i + 1) * 100) / pod.NumberSet
            c.notify(nil, sg.StatusProgress(), progress)
        }
        c.syncMode = false
    }
    if async {
        go f()
    } else {
        f()
    }
}

func (c *Controller) InitPOD() {
    f := func() {
        c.QueryAllSets(false)
        c.QueryAllPresets(false)
        c.notify(nil, sg.StatusInitDone(), nil)
    }
    go f()
}

func (c *Controller) SetDTClass(dtID int, value string) error {
    if !c.started { return nil }
    c.pb.LockData()
    defer c.pb.UnlockData()
    dt := c.pb.GetDT(dtID)
    if dt == nil {
        return fmt.Errorf("DT not found ID:%d", dtID)
    }
    return c.setDTClass(dt, value)
}

func (c *Controller) SetDTClass2(ampID uint32, value string) error {
    if !c.started { return nil }
    c.pb.LockData()
    defer c.pb.UnlockData()
    dt := c.pb.GetDT2(ampID)
    if dt == nil {
        return fmt.Errorf("DT not found AmpID:%d", ampID)
    }
    return c.setDTClass(dt, value)
}

func (c *Controller) setDTClass(dt *pod.DT, value string) error {
    err := dt.SetClass(value)
    if err != nil {
        return err
    }
    m := message.GenDTClassChange(dt)
    go c.writeMessage(m, 0, 0)
    return nil
}

func (c *Controller) SetDTMode(dtID int, value string) error {
    if !c.started { return nil }
    c.pb.LockData()
    defer c.pb.UnlockData()
    dt := c.pb.GetDT(dtID)
    if dt == nil {
        return fmt.Errorf("DT not found ID:%d", dtID)
    }
    return c.setDTMode(dt, value)
}

func (c *Controller) SetDTMode2(ampID uint32, value string) error {
    if !c.started { return nil }
    c.pb.LockData()
    defer c.pb.UnlockData()
    dt := c.pb.GetDT2(ampID)
    if dt == nil {
        return fmt.Errorf("DT not found AmpID:%d", ampID)
    }
    return c.setDTMode(dt, value)
}

func (c *Controller) setDTMode(dt *pod.DT, value string) error {
    err := dt.SetMode(value)
    if err != nil {
        return err
    }
    m := message.GenDTModeChange(dt)
    go c.writeMessage(m, 0, 0)
    return nil
}

func (c *Controller) SetDTTopology(dtID int, value string) error {
    if !c.started { return nil }
    c.pb.LockData()
    defer c.pb.UnlockData()
    dt := c.pb.GetDT(dtID)
    if dt == nil {
        return fmt.Errorf("DT not found ID:%d", dtID)
    }
    return c.setDTTopology(dt, value)
}

func (c *Controller) SetDTTopology2(ampID uint32, value string) error {
    if !c.started { return nil }
    c.pb.LockData()
    defer c.pb.UnlockData()
    dt := c.pb.GetDT2(ampID)
    if dt == nil {
        return fmt.Errorf("DT not found AmpID:%d", ampID)
    }
    return c.setDTTopology(dt, value)
}

func (c *Controller) setDTTopology(dt *pod.DT, value string) error {
    err := dt.SetTopology(value)
    if err != nil {
        return err
    }
    m := message.GenDTTopologyChange(dt)
    go c.writeMessage(m, 0, 0)
    return nil
}

func (c *Controller) SetAmpParameterValue(id uint32, pid uint32, value string) error {
    return c.SetPedalBoardItemParameterValue(id*2, pid, value)
}

func (c *Controller) SetCabParameterValue(id uint32, pid uint32, value string) error {
    if !c.started { return nil }
    c.pb.LockData()
    defer c.pb.UnlockData()
    pbi := c.pb.GetItem((id*2) + 1);
    if pbi == nil {
        return nil
    }
    p := pbi.GetParam(pid)
    if p == nil {
        return nil
    }
    err := p.SetValueCurrent(value)
    if err != nil {
        return err
    }
    m := message.GenParameterCabChange(p)
    go c.writeMessage(m, 0, 0)
    return nil
}

func (c *Controller) SetPedalParameterValue(id uint32, pid uint32, value string) error {
    return c.SetPedalBoardItemParameterValue(id+4, pid, value)
}

func (c *Controller) SetPedalBoardParameterValue(pid uint32, value string) error {
    if !c.started { return nil }
    c.pb.LockData()
    defer c.pb.UnlockData()
    p := c.pb.GetParam(pid)
    if p == nil {
        return nil
    }
    err := p.SetValueCurrent(value)
    if err != nil {
        return err
    }
    m := message.GenParameterPedalBoardChange(p)
    go c.writeMessage(m, 0, 0)
    return nil
}

func (c *Controller) SetPedalBoardItemParameterValue(id uint32, pid uint32, value string) error {
    if !c.started { return nil }
    c.pb.LockData()
    defer c.pb.UnlockData()
    pbi := c.pb.GetItem(id);
    if pbi == nil {
        return nil
    }
    p := pbi.GetParam(pid)
    if p == nil {
        return nil
    }
    err := p.SetValueCurrent(value)
    if err != nil {
        return err
    }
    switch p2 := p.(type) {
    case *pod.TempoParam:
        switch p2.GetID() {
        case 0x3F100000:
            m := message.GenParameterTempoChange(p2)
            c.writeMessage(m, 0, 0)
        case 0x3F100002:
            m := message.GenParameterTempoChange2(p2)
            c.writeMessage(m, 0, 0)
        }
        binValue := p2.GetBinValueCurrent()
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
    return nil
}

func (c *Controller) SetPedalBoardItemParameterValueMin(id uint32, pid uint32, value string) error {
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
    err := p.SetValueMin(value)
    if err != nil {
        c.pb.UnlockData()
        return err
    }
    m := message.GenParameterChangeMin(p)
    go c.writeMessage(m, 0, 0)
    c.pb.UnlockData()
    return nil
}

func (c *Controller) SetPedalBoardItemParameterValueMax(id uint32, pid uint32, value string) error {
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
    err := p.SetValueMax(value)
    if err != nil {
        c.pb.UnlockData()
        return err
    }
    m := message.GenParameterChangeMax(p)
    go c.writeMessage(m, 0, 0)
    c.pb.UnlockData()
    return nil
}

func (c *Controller) SetAmpActive(id uint32, active bool) {
    c.SetPedalBoardItemActive(id*2, active)
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
