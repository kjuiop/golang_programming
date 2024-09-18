package main

type Saver interface {
	Save(data []byte) error
}

func SavePerson(person *Person, saver Saver) error {

	if err := person.validate(); err != nil {
		return err
	}

	bytes, err := person.encode()
	if err != nil {
		return err
	}

	return saver.Save(bytes)
}

func main() {
}
