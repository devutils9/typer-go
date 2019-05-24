package ui

import (
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/jroimartin/gocui"
	"github.com/shilangyu/typeracer-go/utils"
	"github.com/shilangyu/typeracer-go/widgets"
)

// CreateWelcome creates welcome screen widgets
func CreateWelcome(g *gocui.Gui) error {
	w, h := g.Size()

	// main menu views
	sign := widgets.NewText("welcome-sign", figure.NewFigure("typeracer", "", false).String(), false, true, w/2, h/5)

	infoItems := utils.Center([]string{
		"Single player mode - test your typing skills offline!",
		"Multi player mode - battle against other typers",
		"Settings - change app settings",
		"Exit - exit the app",
	})
	info := widgets.NewText("welcome-menu-info", infoItems[0], true, true, w/2, 3*h/4)

	menuItems := []string{"single player", "multi player", "settings", "exit"}
	menu := widgets.NewMenu("welcome-main-menu", utils.Center(menuItems), w/2, h/2, true, true, func(i int) {
		g.Update(info.ChangeText(infoItems[i]))
	}, func(i int) {
		switch i {
		case 3:
			g.Close()
			os.Exit(0)
		default:

		}
	})

	g.SetManager(sign, menu, info)

	err := menu.Init(g)
	if err != nil {
		return err
	}

	screen = append(screen, menu, info, sign)
	return nil
}
