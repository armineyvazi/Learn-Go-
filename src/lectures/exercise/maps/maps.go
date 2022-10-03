//--Summary:
//  Write a program to display server status.
//
//--Requirements:


//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:


package main

import (
	"fmt"
)

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)
//* Create a function to print server status, including:
func printStatusServer(servers map[string]int){
	//  - Number of servers
	fmt.Println("\nThere are ",len(servers),"servers")
	stats:=make(map[int]int)
	for _,status:=range servers{
		switch status {
		case Online:
			stats[Online]+=1
		case Offline:
			stats[Offline]+=1
		case Maintenance:
			stats[Maintenance]+=1
		case Retired:
			stats[Retired]+=1
		default:
			panic("unhandeled server status")
		}
	}
	fmt.Println(stats[Online],"server online server")
	fmt.Println(stats[Offline],"server online Offline")
	fmt.Println(stats[Maintenance],"server online Maintenance")
	fmt.Println(stats[Retired],"server online Retired")
 
}

func main() {

	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatus:=make(map[string]int)
	for _,server := range servers{
		serverStatus[server]=Online
	}

	//  - display server info
	printStatusServer(serverStatus)
	//  - change `darkstar` to `Retired`
	serverStatus["darkstar"]=Retired
	//  - change `aiur` to `Offline`
	serverStatus["aiur"]=Offline
	//  - display server info
	printStatusServer(serverStatus)
	//  - change all servers to `Maintenance`
	for key,_:=range serverStatus{
		serverStatus[key]=Maintenance
	}
	//  - display server info
	printStatusServer(serverStatus)
}
