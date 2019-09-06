package main

import (
	"fmt"

	//"io/ioutil"
	"log"
	"bufio"
	"strings"
	//"os/exec"
	//"golang.org/x/crypto/ssh/terminal"
	//"github.com/howeyc/gopass"

	//	"time"
	"github.com/jroimartin/gocui"
)

const delta = 0.2

func main() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen
	//	g.Mouse = true

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	/*	if err := g.SetKeybinding("side1", gocui.KeyEnter, gocui.ModNone, aaa1); err != nil {
			log.Panicln(err)
		}
	*/
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

		v.SetCursor(13, 0)
		fmt.Fprintln(v, "IP Address : ")
		fmt.Fprintln(v, "User Name  : ")
		fmt.Fprintln(v, "Password   : ")
		if err := g.SetKeybinding("side1", gocui.KeyEnter, gocui.ModNone, aaa1); err != nil {
			log.Panicln(err)
		}

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
var passWord string
func aaa1(g *gocui.Gui, v *gocui.View) error {
	var connInfo string
	connInfo = v.Buffer()
	_, y := v.Cursor()
	
	fmt.Fprintln(v, y)
	if y == 0 {
		v.SetCursor(13, y+1)
	} else if y == 1 {
		v.SetCursor(13, y+1)
//		fmt.Fprintln(v, passWord)
//		maskedPassword, _ := gopass.GetPasswdMasked()
//		fmt.Fprintln(v, maskedPassword)
	} else if y == 2 {
//		maskedPassword, _ := gopass.GetPasswdMasked()
//		fmt.Fprintln(v, maskedPassword)
		err := setConnInfo(connInfo)
		if err != nil {
			return err
		}
	}

	return nil
}
var IP_ADDR string
var USERNAME string
var PASSWORD string

func setConnInfo(buf string) error {
	scanner := bufio.NewScanner(strings.NewReader(buf))
	var connInfo []string
	i := 0
	for scanner.Scan() {
		tempStr := scanner.Text()
		//a := strings.Index(tempStr,":")
		//fmt.Fprintln(v,a)
		connInfo[i] = tempStr[strings.Index(tempStr,":")+1: len(tempStr)]
		i=i+1
	}
	_ := printlog(connInfo)
	return nil
}



func aaa2(g *gocui.Gui, v *gocui.View) error {
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
