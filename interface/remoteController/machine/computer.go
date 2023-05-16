package machine

type Computer struct{}

func (Computer) New() *Computer {
	return &Computer{}
}

func (*Computer) TurnOn() (string, error) {
	return "컴퓨터를 킵니다.", nil
}
