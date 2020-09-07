//common virtual machine detection techniques

package components

import (
	"fmt"
	"net"
    "strings"
    ps "github.com/mitchellh/go-ps"
    "github.com/lxn/win"
    
)

//check for VMWARE by the machines MAC address
func checkVmwareMAC()bool{


    addrs, err := net.InterfaceAddrs()

    if err != nil {
        fmt.Println(err)
    }

    var currentIP, currentNetworkHardwareName string

    for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {

                currentIP = ipnet.IP.String()
            }
        }
    }
    interfaces, _ := net.Interfaces()
    for _, interf := range interfaces {

        if addrs, err := interf.Addrs(); err == nil {
            for _, addr := range addrs {
               
                if strings.Contains(addr.String(), currentIP) {
                    currentNetworkHardwareName = interf.Name
                }
            }
        }
    }

    netInterface, err := net.InterfaceByName(currentNetworkHardwareName)

    if err != nil {
        fmt.Println(err)
    }
    macAddress := netInterface.HardwareAddr


    if err != nil {
        fmt.Println("No able to parse MAC address : ", err)
    }
	if (strings.Contains(macAddress.String(), "00:05:69") || strings.Contains(macAddress.String(), "00:1c:14") ||
			strings.Contains(macAddress.String(), "00:0c:29") || strings.Contains(macAddress.String(), "00:50:56")) { 
				return true
		}
	return false
}

//check for important background tools
func checkBackgroundProcs() bool { 
    processList, err := ps.Processes()
    
    if err != nil {
        fmt.Println("Error")
    }
    for index := range processList {
        var process ps.Process
        process = processList[index]
        if (strings.Contains(process.Executable(), "vmwareservices.exe") || strings.Contains(process.Executable(), "vmwaretray.exe")) {
            return true
        }
    }
        return false
    
}

//check for vm screen resolution 
func checkWindowRes() bool { 
    width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
    height := int(win.GetSystemMetrics(win.SM_CYSCREEN))
    if (width == 800 && height == 600 || width == 1024 && height == 768) { 
        return true
    }
    return false
}
