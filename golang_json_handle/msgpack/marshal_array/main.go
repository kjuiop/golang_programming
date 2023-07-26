package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/shamaton/msgpack/v2"
	"log"
)

type Student struct {
	ID           int
	Name         string
	ScoreMath    int
	ScoreEnglish int
}

func main() {

	data := []Student{
		{ID: 1, Name: "Jake1", ScoreMath: 100, ScoreEnglish: 100},
		{ID: 1, Name: "Jake2", ScoreMath: 70, ScoreEnglish: 90},
		{ID: 1, Name: "Jake3", ScoreMath: 80, ScoreEnglish: 100},
		{ID: 1, Name: "Jake4", ScoreMath: 50, ScoreEnglish: 40},
	}

	ser1, err := json.Marshal(data)
	if err != nil {
		log.Println("error : ", err.Error())
	}

	spew.Dump(ser1)

	ser2, err := msgpack.MarshalAsArray(data)
	if err != nil {
		log.Println("error : ", err.Error())
	}
	spew.Dump(ser2)

	unSer := make([]Student, 0)
	if err := msgpack.UnmarshalAsArray(ser2, &unSer); err != nil {
		log.Println("error : ", err.Error())
	}
	spew.Dump(unSer)

}
