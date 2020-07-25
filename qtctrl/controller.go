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

package qtctrl

import "github.com/therecipe/qt/core"

import "lpedit/controller"

type Controller struct {
    core.QObject
    controller.Controller
    _ func() `constructor:"init"`
    _ func() `signal:"loop"`
    _ func(string) `signal:"loopError"`
    _ func() `signal:modelUpdated`
}

func (c *Controller) init() {
    c.Controller = *controller.NewController()
    c.SetNotify(c.notif)
}

func (c *Controller) notif(err error, n controller.NotificationType) {
    switch n {
    case controller.NormalStart:
        go c.Loop()
    case controller.NormalStop:
        go c.Loop()
    case controller.ErrorStop:
        go c.LoopError(err.Error())
    case controller.MessageProcessed:
        go c.ModelUpdated()
    }
}
