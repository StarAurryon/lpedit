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

package alsa

// #cgo LDFLAGS: -lasound
/*
#include <fcntl.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <stdarg.h>
#include <stdlib.h>

#include <alsa/asoundlib.h>
#include <alsa/error.h>
#include <alsa/hwdep.h>

*/
import "C"
import "fmt"
import "unsafe"

func ListHWDev() [][]string {
    cardsName := [][]string{}
    var card C.int = -1
    iface := C.CString("pcm")
    defer C.free(unsafe.Pointer(iface))
    var handle *C.snd_ctl_t
    var info *C.snd_ctl_card_info_t
    C.snd_ctl_card_info_malloc(&info)

    for C.snd_card_next(&card) >= 0 && card >= 0 {
        name := C.CString(fmt.Sprintf("hw:%d", card))
        if C.snd_ctl_open(&handle, name, 0) >= 0 {
            if C.snd_ctl_card_info(handle, info) >= 0 {
                cardName := C.GoString(C.snd_ctl_card_info_get_id(info))
                cardFName := C.GoString(C.snd_ctl_card_info_get_name(info))
                cardsName = append(cardsName, []string{
                    fmt.Sprintf("hw:%s", cardName),
                    fmt.Sprintf("%s (hw:%d)", cardFName, card),
                })
            }
            C.snd_ctl_close(handle)
        }
        C.free(unsafe.Pointer(name))
    }
    C.snd_ctl_card_info_free(info)
    return cardsName
}

type Hwdep struct {
    hwdep   *C.snd_hwdep_t
    started bool
}

func (h *Hwdep) Close() error {
    if(h.hwdep == nil) {return nil}
    var err C.int = C.snd_hwdep_close(h.hwdep)
    if(err < 0) {
        return fmt.Errorf(C.GoString(C.snd_strerror(err)))
    }
    h.hwdep = nil
    return nil
}

func (h *Hwdep) IsOpen() bool {
    if h.hwdep == nil { return false }
    return true
}

func (h *Hwdep) Open(dev string) error {
    if(h.hwdep != nil) {fmt.Errorf("Device is already open")}
    c_dev := C.CString(dev)
    defer C.free(unsafe.Pointer(c_dev))
    var err C.int = C.snd_hwdep_open(&h.hwdep, c_dev, C.O_RDWR | C.O_NONBLOCK)
    if(err < 0) {
        return fmt.Errorf(C.GoString(C.snd_strerror(err)))
    }
    return nil
}

func (h *Hwdep) Read(size int) []byte {
    buf := make([]byte, size)
    copied := int(C.snd_hwdep_read(h.hwdep, unsafe.Pointer(&buf[0]),
        C.ulong(size)))
    return buf[:copied]
}
