package main

import (
	"fmt"
	"os"
	"runtime"
	"tcp-connect-scanner1/scanner"
	"tcp-connect-scanner1/util"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	if len(os.Args) == 3 {
		ipList := os.Args[1]
		portList := os.Args[2]
		ips, _ := util.GetIpList(ipList)
		ports, _ := util.GetPorts(portList)

		task, _ := scanner.GenerateTask(ips, ports)
		scanner.AssigningTasks(task)
		scanner.PrintResult()
	} else {
		fmt.Printf("%v iplist port\n", os.Args[0])
	}
}
