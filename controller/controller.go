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

import "fmt"
import "log"
import "sync"
import "time"

import "lpedit/alsa"
import "lpedit/message"
import "lpedit/pedal"

type Controller struct {
    pb      *pedal.PedalBoard
    dev     string
    rLDM    chan int
    pM      chan int
    rWG     sync.WaitGroup
    rCh     chan *message.RawMessage
    rStart  bool
    hwdep   alsa.Hwdep
    started bool
    notif   func(error, pedal.ChangeType, interface{})
}

func NewController() *Controller {
    c := &Controller{pb: pedal.NewPedalBoard(), started: false}
    return c
}

func (c *Controller) GetPedalType() map[string][]string {
    return pedal.GetPedalType()
}

func (c *Controller) ListDevices() [][]string {
    return alsa.ListHWDev()
}

func (c *Controller) notify(err error, ntype pedal.ChangeType, obj interface{}) {
    n := func(err error, ntype pedal.ChangeType, obj interface{}) {
        if err != nil {
            log.Println(err)
        }
        if c.notif != nil {
            c.notif(err, ntype, obj)
        }
    }
    go n(err, ntype, obj)
}

func (c *Controller) IsStarted() bool {
    return c.started
}

func (c *Controller) SetNotify(n func(error, pedal.ChangeType, interface{})) {
    c.notif = n
}

func (c *Controller) Start(dev string) {
    c.rLDM = make(chan int)
    c.pM = make(chan int)
    c.rCh = make(chan *message.RawMessage, 100)
    c.notify(nil, pedal.NormalStart, nil)
    go c.readRawMessage(dev)
    go c.processRawMessage()
    go c.monitor()
}

func (c *Controller) monitor(){
    time.Sleep(100 * time.Millisecond)
    c.rWG.Wait()
    if !c.hwdep.IsOpen() {
        c.Stop()
    }
}

func (c *Controller) genMessage(rm *message.RawMessage) {
    err, m := message.NewMessage(*rm)
    if err != nil {
        log.Println(err)
        return
    }
    m.LogInfo()

    c.pb.LockData()
    err, ct, obj := m.Parse(c.pb)
    c.notify(err, ct, obj)
    c.pb.UnlockData()
}

func (c *Controller) processRawMessage() {
    c.rWG.Add(1)
    defer c.rWG.Done()
    var m *message.RawMessage

    for {
        select {
        case <-c.pM:
            return
        case <-time.After(10 * time.Millisecond):
            if m != nil {
                go c.genMessage(m)
                m = nil
            }
        case _m := <-c.rCh:
            switch _m.GetType() {
            case message.RawMessageBegin:
                if m != nil {
                    go c.genMessage(m)
                }
                m = _m
            case message.RawMessageExt:
                m.Extend(_m)
            }
        }
    }
}

func (c *Controller) readRawMessage(dev string) {
    c.rWG.Add(1)
    defer c.rWG.Done()
    c.dev = dev
    if err := c.hwdep.Open(dev); err != nil {
        c.notify(fmt.Errorf("Could not open device %s: %s\n", dev, err),
            pedal.ErrorStop, nil)
        c.pM <- 0
        return
    }

    c.started = true

    for {
        select {
        case <-c.rLDM:
            return
        default:
            buf := c.hwdep.Read(1000)
            if len(buf) == 0 {
                time.Sleep(100 * time.Millisecond)
                continue
            }
            c.rCh <- message.NewRawMessage(buf)
        }
    }
}

func (c *Controller) Stop() {
    c.started = false
    c.pM <- 0
    c.rLDM <- 0
    c.rWG.Wait()
    close(c.rLDM)
    close(c.pM)
    close(c.rCh)
    if c.hwdep.IsOpen() {
        if err := c.hwdep.Close(); err != nil {
            c.notify(fmt.Errorf("Could not close device %s: %s\n", c.dev, err),
                pedal.ErrorStop, nil)
        } else {
            c.notify(nil, pedal.NormalStop, nil)
        }
    }
}
