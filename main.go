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

package main

import "lpedit/alsa"
import "lpedit/message"
import "lpedit/pedal"
import "fmt"

func main() {
  var hwdep alsa.Hwdep
  dev := "hw:PODHD500X"
  err := hwdep.Open(dev)
  if(err != nil) {
      fmt.Printf("Could not open device %s: %s\n", dev, err)
      return
  }

  pb := pedal.NewPedalBoard()

  for {
      buf := hwdep.Read(1000)
      rm := message.NewRawMessage(buf)
      //rm.PrintInfo()
      err, m := message.NewMessage(rm)
      if err != nil {
          fmt.Println(err)
          continue
      }
      for !m.Ready() {
          buf := hwdep.Read(1000)
          rm := message.NewRawMessage(buf)
          //rm.PrintInfo()
          err = m.Extend(rm)
          if err != nil {
              fmt.Println(err)
              break
          }
      }
      err = m.Parse(pb)
      if err != nil {
          fmt.Println(err)
      }
      pb.PrintInfo()
  }

  err = hwdep.Close()
  if(err != nil) {
      fmt.Printf("Could not close device %s: %s\n", dev, err)
      return
  }
}
