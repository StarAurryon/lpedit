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
    pb         *pedal.PedalBoard
    dev        string
    hwdep      alsa.Hwdep
    notifyCB   func(error, pedal.ChangeType, interface{}) // notifyCallBack
    //Status
    started    bool
    //Threads
    waitGroup  sync.WaitGroup
    stopRRM    chan int //readRawMessage
    stopPRM    chan int //processRawMessage
    stopWRM    chan int //writeRawMessage
    readQueue  chan *message.RawMessage
    writeMux   sync.Mutex
    writeQueue chan *message.RawMessage
}

func NewController() *Controller {
    c := &Controller{pb: pedal.NewPedalBoard(), started: false}
    return c
}

func (c *Controller) GetPedalBoard() *pedal.PedalBoard {
    return c.pb
}

func (c *Controller) GetPedalType() map[string][]string {
    return pedal.GetPedalType()
}

func (c *Controller) ListDevices() [][]string {
    return alsa.ListHWDev()
}

func (c *Controller) IsStarted() bool {
    return c.started
}

func (c *Controller) SetNotify(n func(error, pedal.ChangeType, interface{})) {
    c.notifyCB = n
}

func (c *Controller) Start(dev string) {
    c.stopRRM = make(chan int, 10)
    c.stopPRM = make(chan int, 10)
    c.stopWRM = make(chan int, 10)
    c.readQueue = make(chan *message.RawMessage, 100)
    c.writeQueue = make(chan *message.RawMessage, 100)

    if err := c.hwdep.Open(dev); err != nil {
        c.notify(fmt.Errorf("Could not open device %s: %s\n", dev, err),
            pedal.ErrorStop, nil)
        c.signalStop()
        return
    }

    c.started = true
    c.notify(nil, pedal.NormalStart, nil)
    go c.readRawMessage()
    go c.processRawMessage()
    go c.writeRawMessage()
    go c.monitor()
}

func (c *Controller) Stop() {
    c.signalStop()
    c.waitGroup.Wait()
    close(c.stopRRM)
    close(c.stopPRM)
    close(c.stopWRM)
    close(c.readQueue)
    close(c.writeQueue)
    c.started = false
    if c.hwdep.IsOpen() {
        if err := c.hwdep.Close(); err != nil {
            c.notify(fmt.Errorf("Could not close device %s: %s\n", c.dev, err),
                pedal.ErrorStop, nil)
        } else {
            c.notify(nil, pedal.NormalStop, nil)
        }
    }
}

func (c *Controller) notify(err error, ntype pedal.ChangeType, obj interface{}) {
    n := func(err error, ntype pedal.ChangeType, obj interface{}) {
        if err != nil {
            log.Println(err)
        }
        if c.notifyCB != nil {
            c.notifyCB(err, ntype, obj)
        }
    }
    go n(err, ntype, obj)
}

func (c *Controller) monitor(){
    time.Sleep(100 * time.Millisecond)
    c.waitGroup.Wait()
    if !c.hwdep.IsOpen() {
        c.Stop()
    }
}

func (c *Controller) parseMessage(rm *message.RawMessage) {
    err, m := message.NewMessage(*rm)
    if err != nil {
        log.Println(err)
        return
    }
    m.LogInfo()
    c.pb.LockData()
    err, ct, obj := m.Parse(c.pb)
    c.pb.UnlockData()
    c.notify(err, ct, obj)
}

func (c *Controller) processRawMessage() {
    c.waitGroup.Add(1)
    defer c.waitGroup.Done()
    var m *message.RawMessage

    for {
        select {
        case <-c.stopPRM:
            return
        case <-time.After(10 * time.Millisecond):
            if m != nil {
                go c.parseMessage(m)
                m = nil
            }
        case _m := <-c.readQueue:
            switch _m.GetType() {
            case message.RawMessageBegin:
                if m != nil {
                    go c.parseMessage(m)
                }
                m = _m
            case message.RawMessageExt:
                m.Extend(_m)
            }
        }
    }
}

func (c *Controller) readRawMessage() {
    c.waitGroup.Add(1)
    defer c.waitGroup.Done()
    for {
        select {
        case <-c.stopRRM:
            return
        default:
            buf := c.hwdep.Read(1000)
            if len(buf) == 0 {
                time.Sleep(100 * time.Millisecond)
                continue
            }
            c.readQueue <- message.NewRawMessage(buf)
        }
    }
}

func (c* Controller) writeMessage(m message.IMessage, ukno0 uint8, ukno1 uint8) {
    rms := message.NewRawMessages(m, ukno0, ukno1)
    c.writeMux.Lock()
    for _, rm := range rms {
        c.writeQueue <- rm
    }
    c.writeMux.Unlock()
}

func (c* Controller) writeRawMessage() {
    c.waitGroup.Add(1)
    defer c.waitGroup.Done()
    for {
        select {
        case <-c.stopWRM:
            return
        case m := <- c.writeQueue:
            log.Println("Writing raw message")
            m.LogInfo()
            c.hwdep.Write(m.Export())
        }
    }
}

func (c *Controller) signalStop() {
    c.stopPRM <- 0
    c.stopRRM <- 0
    c.stopWRM <- 0
}
