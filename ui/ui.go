package ui

import "github.com/shilangyu/typeracer-go/widgets"
import	"github.com/jroimartin/gocui"
 

// screen is a collection of current views
var screen []widgets.Widget


func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}