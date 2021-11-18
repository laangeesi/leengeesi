package main //This delcares that below functions belong to this main package

import (
	"net" //We import the "net" package
	"fmt" //We import the "fmt" package
)


func main() {                       //func main excused below actions
   iplist, _ := net.Interfaces()
   for _, ipinfo := range iplist { //for _, delcares the loop to list all interfaces 
      fmt.Println("Index :", ipinfo.Index) //func "fmt,Prinlin" 
      fmt.Println("Name  :", ipinfo.Name)
      fmt.Println("HWaddr:", ipinfo.HardwareAddr)
      fmt.Println("MTU   :", ipinfo.MTU)
      fmt.Println("Flags :", ipinfo.Flags)
      addrs,_ := ipinfo.Addrs()
      for _, ipaddr := range addrs {
         fmt.Println("Addr  :", ipaddr) 
      }
      fmt.Println()    //calls the fmt.Println within the func main and prints 
   }

}




