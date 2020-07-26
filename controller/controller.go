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

type NotificationType int

const (
    ErrorProcess NotificationType = iota
    ErrorStop
    NormalStop
    NormalStart
    MessageProcessed
)

type Controller struct {
    pb      *pedal.PedalBoard
    pbMux   sync.Mutex
    dev     string
    rLDM    chan int
    pM      chan int
    rWG     sync.WaitGroup
    rCh     chan *message.RawMessage
    rStart  bool
    hwdep   alsa.Hwdep
    started bool
    notif   func(error, NotificationType)
}

func NewController() *Controller {
    c := &Controller{pb: pedal.NewPedalBoard(), started: false}
    return c
}

func (c *Controller) GetPedal(id int) *pedal.Pedal {
    if id < 0 || id > 7 {
        return nil
    }
    p, _ := c.pb.GetItem(uint32(id + 4)).(*pedal.Pedal)
    return p
}

func (c *Controller) GetPedalType() map[string][]string {
    return pedal.GetPedalType()
}

func (c *Controller) ListDevices() [][]string {
    return alsa.ListHWDev()
}

func (c *Controller) LockData(){
    c.pbMux.Lock()
}

func (c *Controller) UnlockData(){
    c.pbMux.Unlock()
}

func (c *Controller) notify(err error, ntype NotificationType) {
    n := func(err error, ntype NotificationType) {
        if err != nil {
            log.Println(err)
        }
        if c.notif != nil {
            c.notif(err, ntype)
        }
    }
    go n(err, ntype)
}

func (c *Controller) IsStarted() bool {
    return c.started
}

func (c *Controller) SetNotify(n func(error, NotificationType)) {
    c.notif = n
}

func (c *Controller) Start(dev string) {
    c.rLDM = make(chan int)
    c.pM = make(chan int)
    c.rCh = make(chan *message.RawMessage, 100)
    c.notify(nil, NormalStart)
    go c.readDevMsg(dev)
    go c.processMsg()
    go c.monitor()
}

func (c *Controller) monitor(){
    time.Sleep(100 * time.Millisecond)
    c.rWG.Wait()
    if !c.hwdep.IsOpen() {
        c.Stop()
    }
}

func (c *Controller) processMsg() {
    c.rWG.Add(1)
    defer c.rWG.Done()
    var m *message.Message
    var err error

    for {
        select {
        case <-c.pM:
            fmt.Println("Exiting startProcess")
            return
        case rm := <-c.rCh:
            if m == nil {
                err, m = message.NewMessage(*rm)
                if err != nil {
                    m = nil
                    c.notify(err, ErrorProcess)
                    continue
                }
            } else {
                err = m.Extend(*rm)
                if err != nil {
                    m = nil
                    c.notify(err, ErrorProcess)
                    continue
                }
            }
            if m.Ready() {
                c.LockData()
                err = m.Parse(c.pb)
                c.UnlockData()
                if err != nil {
                    c.notify(err, ErrorProcess)
                } else {
                    c.notify(nil, MessageProcessed)
                }
                m = nil
            }
        }
    }
}

func (c *Controller) readDevMsg(dev string) {
    c.rWG.Add(1)
    defer c.rWG.Done()
    c.dev = dev
    if err := c.hwdep.Open(dev); err != nil {
        c.notify(fmt.Errorf("Could not open device %s: %s\n", dev, err),
            ErrorStop)
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
                ErrorStop)
        } else {
            c.notify(nil, NormalStop)
        }
    }
}
