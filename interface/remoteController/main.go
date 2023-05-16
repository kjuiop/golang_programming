package main

import (
	"fmt"
	"golang_programming/interface/remoteController/interfaces"
	"golang_programming/interface/remoteController/machine"
)

func main() {

	// tv 를 구현체로 설정
	typeOfController := "pc"
	myRemoteController := getRemoteController(typeOfController)

	result, err := myRemoteController.TurnOnMachine()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}

func getRemoteController(typeOfController string) *interfaces.RemoteController {
	if typeOfController == "tv" {
		return interfaces.RemoteController{}.New(injectTv())
	}
	if typeOfController == "airConditional" {
		return interfaces.RemoteController{}.New(injectAirConditional())
	}

	if typeOfController == "pc" {
		return interfaces.RemoteController{}.New(injectComputer())
	}

	return nil
}

func injectTv() *machine.Tv {
	return machine.Tv{}.New()
}

func injectAirConditional() *machine.AirConditional {
	return machine.AirConditional{}.New()
}

func injectComputer() *machine.Computer {
	return machine.Computer{}.New()
}
