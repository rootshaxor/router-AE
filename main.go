package main

import (
	iface "github.com/rootshaxor/router-AE/interfaces"
	"fmt"
	"math/rand"
	"time"

	"github.com/abiosoft/ishell"
)

var (
	app   = ishell.New()
	r     *rand.Rand
	intfs = iface.Get_Interface()
)

const (
	NAME    = "router-AE"
	VERSION = "0.0.1"
	STATUS  = "alpha"
)

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	app.AutoHelp(true)
	app.SetPrompt("[ router-AE ] >> ")
	app.DeleteCmd("exit")
	app.IgnoreCase(true)
	app.AddCmd(&ishell.Cmd{
		Name: "exit",
		Help: "Exit From Router!",
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 {
				c.Print("[Warning]: Exit From Router [y/n] ")
				ext := c.ReadLine()
				switch ext {
				case "YES", "Y", "yes", "y":
					app.Close()
				case "NO", "N", "no", "n":
					c.Println("[Info]: Exit Canceled!")
				default:
					c.Println("[Error]: No Options (", ext, ")")
				}
			} else {
				ext := c.Args[0]
				switch ext {
				case "YES", "Y", "yes", "y", "-y", "--yes":
					app.Close()
				case "NO", "N", "no", "n", "-n", "--no":
					c.Println("[Info]: Exit Canceled!")
				default:
					c.Println("[Error]: No Options (", ext, ")")
				}
			}
		},
	})

	app.AddCmd(&ishell.Cmd{
		Name: "banner",
		Help: "Print Banner Router-AE",
		Func: func(c *ishell.Context) {
			Print_Banner()
		},
	})
}

func Print_Banner() {
	var satu, dua, tiga, empat, lima string

	satu = `
	                    _
	.__   _|_ _ .___/\ |_
	|(_)|_||_(/_|  /--\|_

	`
	dua = `
                                    ___,   ___
                                   /   |  / (_)
 ,_    __        _|_  _   ,_      |    |  \__
/  |  /  \_|   |  |  |/  /  |-----|    |  /
   |_/\__/  \_/|_/|_/|__/   |_/    \__/\_/\___/



	`
	tiga = `
                     __                  ___    ______
   _________  __  __/ /____  _____      /   |  / ____/
  / ___/ __ \/ / / / __/ _ \/ ___/_____/ /| | / __/
 / /  / /_/ / /_/ / /_/  __/ /  /_____/ ___ |/ /___
/_/   \____/\__,_/\__/\___/_/        /_/  |_/_____/

	`
	empat = `
                 |                  \    ____|
  __| _ \  |   | __|  _ \  __|     _ \   __|
 |   (   | |   | |    __/ |_____| ___ \  |
_|  \___/ \__,_|\__|\___|_|     _/    _\_____|

	`
	lima = `
                 _                   _    _____
 _ __ ___  _   _| |_ ___ _ __       / \  | ____|
| '__/ _ \| | | | __/ _ \ '__|____ / _ \ |  _|
| | | (_) | |_| | ||  __/ | |_____/ ___ \| |___
|_|  \___/ \__,_|\__\___|_|      /_/   \_\_____|


	`

	strlen := 1
	const num = "12345"
	res := make([]byte, strlen)

	for i := range res {
		res[i] = num[r.Intn(len(num))]
	}

	hasil := string(res)
	switch hasil {
	case "1":
		fmt.Println(satu)
	case "2":
		fmt.Println(dua)
	case "3":
		fmt.Println(tiga)
	case "4":
		fmt.Println(empat)
	case "5":
		fmt.Println(lima)
	}

}

func Menu_IPAddress() *ishell.Cmd {

	m_ipaddr := &ishell.Cmd{
		Name: "ip",
		Help: "IPAddress menu",
	}

	m_ipaddr.AddCmd(&ishell.Cmd{
		Name: "print",
		Help: "Print Address on Interfaces",
		Completer: func([]string) []string {
			return intfs
		},
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 {
				iface.Information_Interfaces("all")
			} else {
				iface.Information_Interfaces(c.Args[0])
			}
		},
	})

	m_ipaddr.AddCmd(&ishell.Cmd{
		Name: "add",
		Help: "Add Address On Interfaces [ipv4/ipv6]",
		Completer: func([]string) []string {
			return intfs
		},
		Func: func(c *ishell.Context) {
			if len(c.Args) < 2 {
				c.Println("[Warning]: you must input IP ADDRESS")
			} else {
				addr := iface.Check_Address_Interface(c.Args[0], c.Args[1])
				switch addr {
				case "found":
					c.Println("[Warning]: Address alredy found in ", c.Args[0])
				case "null":
					if err := iface.Check_Address_Valid(c.Args[1]); err != nil {
						c.Println("[Error]:", err.Error())
						return
					}
					iface.Add_Address(c.Args[0], c.Args[1])
				}
			}
		},
	})

	m_ipaddr.AddCmd(&ishell.Cmd{
		Name: "del",
		Help: "Delete Address On Interfaces",
		Completer: func([]string) []string {
			return intfs
		},
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 {
				c.Println("[Warning]: Please Select some interface!")
				return
			} else {
				s := iface.Check_Address_FOUND(c.Args[0])
				switch s {
				case "found":
					addr := iface.PrintALL_FoundAddress(c.Args[0])
					pilih := c.Checklist(addr, "Please Select Address to Delete ? : ", nil)
					out := func() (b []string) {
						for _, i := range pilih {
							b = append(b, addr[i])
						}
						return
					}
					for _, i := range out() {
						iface.Del_Address(c.Args[0], i)
					}
					c.Println("[info]: Delete Success !")
				case "null":
					c.Println("[Error]: No Address Found in", c.Args[0])
				}

			}
		},
	})

	return m_ipaddr
}

func main() {
	Print_Banner()

	app.AddCmd(Menu_IPAddress())
	app.Run()
}
