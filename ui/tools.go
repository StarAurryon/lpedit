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

package ui

import "github.com/StarAurryon/qt/widgets"

func AddWidget(lay widgets.QLayout_ITF, widget widgets.QWidget_ITF, x int, y int,) {
     switch _lay := lay.(type) {
     case *widgets.QHBoxLayout:
         _lay.AddWidget(widget, 0, 0)
     case *widgets.QVBoxLayout:
         _lay.AddWidget(widget, 0, 0)
     case *widgets.QGridLayout:
         _lay.AddWidget2(widget, y, x, 0)
     case *widgets.QLayout:
         _lay.AddWidget(widget)
    }
 }

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
