package interfaces

import (
	"fmt"
	"net"
	"strings"

	"github.com/milosgajdos83/tenus"
)

var (
	iface   []string
	address []string
)

/*
Menambahkan nama interface yang tersedia kedalam array `iface`
, dengan mengecualikan interface `lo` jika ada

Penggunaan hanya memerlukan sebuah variable array string baru, Misalkan :

	`
	package main

	import (
		"fmt"
		iface "github.com/rootshaxor/router-AE/interfaces"
	)

	func main() {
		list_interface := iface.Get_Interface()
		fmt.Println("list_interface")
	}
	`
*/
func Get_Interface() []string {
	Intfs, err := net.Interfaces()
	if err != nil {
		fmt.Println("[Error]:", err.Error())
	}
	for _, dev := range Intfs {
		if dev.Name == "lo" {
			continue
		} else {
			iface = append(iface, dev.Name)
		}
	}
	return iface
}

/*
Menampilkan IP Address pada suatu interface jika tersedia.
Apabila tidak tersedia akan menampilkan

Output :
		Addr   :         <Not Configured!!!>

Jika Dalam suatu interface terdapat hanya 1 IP Address Akan Menampilkan

Output :
		Addr   :         192.168.1.1/24

Jika Interface memiliki lebih dari 1 IP Address maka akan menampilkan

Output :

		Addr   :
                 1 :     192.168.1.2/29
                 2 :     fe80::800:27ff:fe00:0/64
*/
func Extract_Address(interfaces string) {
	dev, err := net.InterfaceByName(interfaces)
	if err != nil {
		fmt.Println("[Error]:", err.Error())
		return
	}

	address, _ := dev.Addrs()
	if len(address) == 0 {
		fmt.Println("Addr   :\t <Not Configured!!!>")
	} else {

		if len(address) > 1 {
			fmt.Println("Addr   :")
			for i, addr := range address {
				fmt.Println("\t\t", i+1, ":\t", addr.String())
			}
		} else {
			fmt.Println("Addr   :\t", address[0])
		}
	}
}

/*
Menampilkan Status dalam sebuah interface  apakah interface tersebut mati ataupun hidup
jika interface hidup maka akan menampilkan

Output : "UP!"

Jika interface mati akan menampilkan

Output : "DOWN!"

Untuk penggunaan hanya memerlukan sebuah variable string baru , misalnya :

	`
	package main

	import (
		"fmt"
		iface "github.com/rootshaxor/router-AE/interfaces"
	)

	func main() {
		stts := iface.Get_StatusInterfaces("eth0")
		fmt.Println("Status Dari eth0 :", stts)
	}
	`

*/
func Get_StatusInterfaces(interfaces string) string {
	dev, err := net.InterfaceByName(interfaces)
	if err != nil {
		return "[Error]: " + err.Error()
	}
	status := strings.Split(dev.Flags.String(), "|")

	if status[0] == "up" {
		return "UP!"
	} else {
		return "DOWN!"
	}
}

func Information_Interfaces(interfaces string) {
	if interfaces == "all" {
		intfs, err := net.Interfaces()
		if err != nil {
			fmt.Println("[Error]:", err.Error())
			return
		}

		for _, dev := range intfs {
			if dev.Name == "lo" {
				continue
			}
			stts := Get_StatusInterfaces(dev.Name)
			fmt.Println("Name   :\t", dev.Name)
			fmt.Println("Mac    :\t", dev.HardwareAddr.String())
			fmt.Println("Status :\t", stts)
			Extract_Address(dev.Name)
			fmt.Println()
		}

	} else {
		dev, err := net.InterfaceByName(interfaces)
		if err != nil {
			fmt.Println("[Error]:", err.Error())
			return
		}
		stts := Get_StatusInterfaces(interfaces)
		fmt.Println("Name   :\t", dev.Name)
		fmt.Println("Mac    :\t", dev.HardwareAddr.String())
		fmt.Println("Status :\t", stts)
		Extract_Address(interfaces)
		fmt.Println()
	}
}

func Add_Address(interfaces, address string) {
	from, err := tenus.NewLinkFrom(interfaces)
	if err != nil {
		fmt.Println("[Error]:", err.Error())
	}
	ip, prefix, err := net.ParseCIDR(address)
	if err != nil {
		fmt.Println("[Error]:", err.Error())
	}

	if err = from.SetLinkIp(ip, prefix); err != nil {
		fmt.Println("[Error]:", err.Error())
	}
}

func Del_Address(interfaces, address string) {
	from, err := tenus.NewLinkFrom(interfaces)
	if err != nil {
		fmt.Println("[Error]:", err.Error())
	}
	ip, prefix, err := net.ParseCIDR(address)
	if err != nil {
		fmt.Println("[Error]:", err.Error())
	}
	if err = from.UnsetLinkIp(ip, prefix); err != nil {
		fmt.Println("[Error]:", err.Error())
	}
}

func UP_Interface(interfaces string) {
	from, err := tenus.NewLinkFrom(interfaces)
	if err != nil {
		fmt.Println("[Error]:", err.Error())
	}

	if err = from.SetLinkUp(); err != nil {
		fmt.Println("[Error]:", err.Error())
	}
}

func DOWN_Interface(interfaces string) {
	from, err := tenus.NewLinkFrom(interfaces)
	if err != nil {
		fmt.Println("[Error]:", err.Error())
	}

	if err = from.SetLinkDown(); err != nil {
		fmt.Println("[Error]:", err.Error())
	}
}

func Set_Default_GW(interfaces string, gw *net.IP) {
	from, err := tenus.NewLinkFrom(interfaces)
	if err != nil {
		fmt.Println("[Error]:", err.Error())
	}

	if err = from.SetLinkDefaultGw(gw); err != nil {
		fmt.Println("[Error]:", err.Error())
	}
}

func Check_Address_Interface(interfaces, addr string) string {
	var fnd string
	dev, err := net.InterfaceByName(interfaces)
	if err != nil {
		fmt.Println("[Error]:", err.Error())
	}

	address, _ := dev.Addrs()
	for _, ip := range address {
		if ip.String() == addr {
			fnd = "found"
			break
		} else {
			fnd = "null"
		}
	}
	return fnd
}

func Check_Address_FOUND(interfaces string) string {
	dev, err := net.InterfaceByName(interfaces)
	if err != nil {
		fmt.Println("[Error]:", err.Error())
	}

	addr, _ := dev.Addrs()
	if len(addr) != 0 {
		return "found"
	} else {
		return "null"
	}
}

func Check_Address_Valid(addr string) error {
	_, _, err := net.ParseCIDR(addr)
	if err != nil {
		return err
	}
	return nil
}

func Jumlah_Address(interfaces string) int {
	dev, err := net.InterfaceByName(interfaces)
	if err != nil {
		fmt.Println("[Error]:", err.Error())
	}

	addr, _ := dev.Addrs()

	return len(addr)
}

func PrintALL_FoundAddress(interfaces string) []string {
	dev, err := net.InterfaceByName(interfaces)
	if err != nil {
		fmt.Println("[Error]:", err.Error())
	}

	addr, _ := dev.Addrs()
	for _, ip := range addr {
		address = append(address, ip.String())
	}

	return address
}
