package main

import "fmt"

func main() {

	for {
		addUserMenu()
	}

}

func addUserMenu() {
	var softwareInput int
	fmt.Printf("Current software list: %v\n", softwareList)
	fmt.Println("Please enter the index of the software you'd like to add a user:")
	fmt.Scanln(&softwareInput)

	softwareProxy := NewSoftwareProxy(softwareInput)
	fmt.Printf("Softwareproxy name:%s \n Softwareproxy quota: %d \n Softwareproxy licence count: %d \n ", softwareProxy.software.name, softwareProxy.software.quota, softwareProxy.software.licenceCount)
	softwareProxy.software.AddUser()
	fmt.Printf("UPDATED: \n Softwareproxy name:%s \n Softwareproxy quota: %d \n Softwareproxy licence count: %d \n ", softwareProxy.software.name, softwareProxy.software.quota, softwareProxy.software.licenceCount)
}
