package main

import (
	"fmt"

	//"io/ioutil"
	"bufio"
	"log"
	"strings"

	//"os/exec"
	//"golang.org/x/crypto/ssh/terminal"
	//"github.com/howeyc/gopass"
	"database/sql"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"

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
	g.Mouse = true

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, aaa2); err != nil {
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

	if v, err := g.SetView("side1", 0, 0, maxX/3, 5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "DB Connection Info"
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

	if v, err := g.SetView("side2", 0, 6, maxX/3, (maxY/4)*2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "aa")
	}

	if _, err := g.SetView("side3", 0, (maxY/4)*2+1, maxX/3, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//fmt.Fprintln(v, "aa")
	}

	if v, err := g.SetView("main", maxX/3+1, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Processlist"
		/*		g.Mouse = true
				if err := g.SetKeybinding("main", gocui.MouseLeft, gocui.ModNone, aaa2); err != nil {
					log.Panicln(err)
				}
		*/ //g.Mouse = true
		fmt.Fprintln(v, "")

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

	//	fmt.Fprintln(v, y)
	if y == 0 {
		v.SetCursor(13, y+1)
	} else if y == 1 {
		v.SetCursor(13, y+1)
		//		fmt.Fprintln(v, passWord)
		//		maskedPassword, _ := gopass.GetPasswdMasked()
		//		fmt.Fprintln(v, maskedPassword)
	} else if y == 2 {
		/*
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("side2")
				if err != nil {
					return err
				}
				v.Clear()
				fmt.Fprintln(v, connInfo)
				return nil
			})
		*/
		//		maskedPassword, _ := gopass.GetPasswdMasked()
		//		fmt.Fprintln(v, maskedPassword)

		err := setConnInfo(g, v, connInfo)
		if err != nil {
			return err
		}

		go displayMain(g, v)
		go getSarInfo(g, v)

	}

	return nil
}

var IP_ADDR string
var USERNAME string
var PASSWORD string

//var connInfo [3]string

func setConnInfo(g *gocui.Gui, v *gocui.View, buf string) error {
	scanner := bufio.NewScanner(strings.NewReader(strings.TrimSuffix(buf, "\n")))
	var connInfo [3]string
	i := 0
	for scanner.Scan() {
		tempStr := scanner.Text()

		connInfo[i] = tempStr[strings.Index(tempStr, ":")+2 : len(tempStr)]

		i = i + 1
	}

	IP_ADDR = connInfo[0]
	USERNAME = connInfo[1]
	PASSWORD = connInfo[2]
	/*
		err := displayMain(g, v)
		if err != nil {
			return err
		}
	*/
	return nil
}

func displayMain(g *gocui.Gui, v *gocui.View) error {
	connString := fmt.Sprintf("%s:%s@tcp(%s:3306)/information_schema", USERNAME, PASSWORD, IP_ADDR)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//processListQuery := "select Command, count(*) from information_schema.processlist group by Command"
	processListQuery := "select user, command from information_schema.processlist"
	processlist, err := db.Query(processListQuery)
	if err != nil {
		return err
	}
	defer processlist.Close()
	var command string
	var cnt string
	var result string
	for processlist.Next() {
		err := processlist.Scan(&command, &cnt)
		if err != nil {
			return err
		}
		result = fmt.Sprintf("%s \n %-20v|  %-20v| ", result, command, cnt)

	}
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("main")
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, result)
		return nil
	})
	return nil
}

func getSarInfo(g *gocui.Gui, v *gocui.View) error {
	//command := fmt.Sprintf("ssh %s sar 1 10", IP_ADDR)
	out, err := exec.Command("ssh", IP_ADDR, "sar", "1", "5").Output()
	output := string(out[:])
	if err != nil {
		return err
	}
	/*	var temp string
		for output.Next() {
			err := output.Scan(&temp)
			if err != nil {
				return err
			}
			result = fmt.Sprintf("%s \n %s", result, temp)

		}
	*/
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("side2")
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, output)
		return nil
	})

	return nil
}

func aaa2(g *gocui.Gui, v *gocui.View) error {
	if err := changeView(g, v); err != nil {
		return err
	}
	//viewName := v.Name()

	return nil
}

func changeView(g *gocui.Gui, v *gocui.View) error {
	viewName := v.Name()
	if _, err := setCurrentViewOnTop(g, viewName); err != nil {
		return err
	}
	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	//	fmt.Fprintln(v, ipAddr)
	/*	g.Update(func(g *gocui.Gui) error {
			v, err := g.View(viewName)
			if err != nil {
				return err
			}
			v.Clear()
			fmt.Fprintln(v, viewName)
			return nil
		})
	*/

	return nil
}

func aaa3(g *gocui.Gui, v *gocui.View) error {
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
