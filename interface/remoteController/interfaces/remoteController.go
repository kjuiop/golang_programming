package interfaces

type RemoteController struct {
	machineImpl Machine
}

func (RemoteController) New(machine Machine) *RemoteController {
	return &RemoteController{machine}
}

func (remoteController *RemoteController) TurnOnMachine() (string, error) {
	return remoteController.machineImpl.TurnOn()
}

type Machine interface {
	TurnOn() (string, error)
}
