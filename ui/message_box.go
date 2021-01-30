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

import "github.com/therecipe/qt/core"
import "github.com/therecipe/qt/widgets"


func (l *LPEdit) progress(progress int) {
    if progress != 100 {
        if l.progressWnd == nil {
            l.progressWnd = widgets.NewQProgressDialog2("Progress", "", 0, 100, l, 0)
            l.progressWnd.SetWindowTitle("Progress")
            wndFlags := l.progressWnd.WindowFlags()
            wndFlags ^= core.Qt__CustomizeWindowHint
            wndFlags ^= core.Qt__WindowCloseButtonHint
            wndFlags ^= core.Qt__WindowContextHelpButtonHint
            l.progressWnd.SetWindowFlags(wndFlags)
        }
        l.progressWnd.SetValue(progress)
    } else {
        if l.progressWnd != nil {
            l.progressWnd.Close()
            l.progressWnd.DestroyQObject()
            l.progressWnd = nil
        }
    }
}

func (l * LPEdit) loopError(err string) {
    if l.pbSelector != nil {
        l.pbSelector.SetEnabled(true)
        l.pbSelector.updateButtons()
    }
    mb := widgets.NewQMessageBox(l)
    mb.Critical(l, "An error occured", err, widgets.QMessageBox__Ok, 0)
    l.disconnectSignal()
}
