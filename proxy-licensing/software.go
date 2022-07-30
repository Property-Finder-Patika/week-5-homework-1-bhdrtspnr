package main

import "fmt"

type Licence interface {
	CheckLicence()
}

type Software struct {
	name         string
	licenceCount int
	quota        int
}

func (s *Software) CheckLicence() bool {
	if s.licenceCount >= s.quota {
		return false
	}
	return true
}

func (s *Software) AddUser() {
	if s.CheckLicence() {
		fmt.Printf("Software %s's quota is %d, active users: %d, you can use the software!", s.name, s.quota, s.licenceCount)
		s.licenceCount++
	}
	fmt.Println("Quota is full!")

}

type SoftwareProxy struct {
	software *Software
}

func NewSoftwareProxy(softwareIndex int) *SoftwareProxy {
	return &SoftwareProxy{&softwareList[softwareIndex]}
}

var softwareList = []Software{
	{name: "HP", quota: 50, licenceCount: 10},
	{name: "ALM", quota: 50, licenceCount: 20},
	{name: "QC", quota: 50, licenceCount: 30},
	{name: "TESTO", quota: 50, licenceCount: 49},
}
