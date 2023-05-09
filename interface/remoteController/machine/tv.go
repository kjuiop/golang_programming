package machine

type Tv struct{}

func (Tv) New() *Tv {
	return &Tv{}
}

func (*Tv) TurnOn() (string, error) {
	return "TV를 킵니다.", nil
}
