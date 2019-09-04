package main

import (
	"fmt"
	"log"

	//	"time"
	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen
	g.Mouse = true
	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("side1", gocui.KeyEnter, gocui.ModNone, aaa1); err != nil {
		log.Panicln(err)
	}
	///	aaa(g)
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("side1", 0, 0, maxX/3, maxY/4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "side1"
		v.Editable = true
		v.Wrap = true
		v.SetCursor(11, 0)
		fmt.Fprintln(v, "IP Adress: ")
		fmt.Fprintln(v, "User Name(default dba): ")
		if _, err = setCurrentViewOnTop(g, "side1"); err != nil {
			return err
		}

	}

	if v, err := g.SetView("side2", 0, maxY/4+1, maxX/3, (maxY/4)*2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "aa")
	}

	if v, err := g.SetView("side3", 0, (maxY/4)*2+1, maxX/3, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "aa")
	}

	if v, err := g.SetView("main", maxX/3+1, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "aa")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

/*
func aaa(g *gocui.Gui){
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("main")
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, "hello world")
		return nil
	})
}
*/
func aaa(g *gocui.Gui, v *gocui.View) error {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("main")
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, "hello world")
		return nil
	})

	return nil
}

func aaa1(g *gocui.Gui, v *gocui.View) error {
	var ipAddr string
	ipAddr = v.Buffer()
	v.SetCursor(24, 1)

	//	fmt.Fprintln(v, ipAddr)
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("main")
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, ipAddr)
		return nil
	})

	return nil
}
