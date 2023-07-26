
# MessagePack

---

- JSON 보다 빠르고 용량이 적은 직렬화 방법
- Json 이나 MessagePack 과 같은 방법을 직렬화라고 합니다.
- 복잡한 객체를 메모리나 파일 등에 기록하기 편하게 연속적인 serial 데이터로 변환하는 것을 이야기합니다.

<img width="403" alt="Screen Shot 2023-07-26 at 11 03 08 AM" src="https://github.com/kjuiop/golang_programming/assets/41246605/c45fd18d-e2ca-45da-bfdf-1cb542b3e118">
- 27 바이트였던 Json 이 18 바이트로 66% 용량이 줄었습니다.
- Json 보다 MessagePack 의 속도가 더 빠릅니다.
- 공식 홈페이지 (https://msgpack.org/)

```go
go get -u github.com/shamaton/msgpack/v2
```

### Sample Code

```java
type Person struct {
	Name string
}

func main() {

	v := Person{Name: "msgpack"}

	d, err := msgpack.Marshal(v)
	if err != nil {
		log.Println("error : ", err.Error())
	}

	spew.Dump(d)

	r := Person{}
	if err := msgpack.Unmarshal(d, &r); err != nil {
		log.Println("error : ", err.Error())
	}

	spew.Dump(r)

}
```

# Marshal Array

---

- struct 형태의 데이터들을 Array 형태로 저장할 때 MessagePack 이 더 좋은 메모리 성능을 보여준다.
- 이러한 메모리 성능은 캐시 데이터를 활용할 때에 극한의 효율을 보여준다.

### Sample Code

```java
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
```

### 결과

```java
## json.Marshal 시 메모리 용량 240
([]uint8) (len=232 cap=240) {
 00000000  5b 7b 22 49 44 22 3a 31  2c 22 4e 61 6d 65 22 3a  |[{"ID":1,"Name":|
 00000010  22 4a 61 6b 65 31 22 2c  22 53 63 6f 72 65 4d 61  |"Jake1","ScoreMa|
 00000020  74 68 22 3a 31 30 30 2c  22 53 63 6f 72 65 45 6e  |th":100,"ScoreEn|
 00000030  67 6c 69 73 68 22 3a 31  30 30 7d 2c 7b 22 49 44  |glish":100},{"ID|
 00000040  22 3a 31 2c 22 4e 61 6d  65 22 3a 22 4a 61 6b 65  |":1,"Name":"Jake|
 00000050  32 22 2c 22 53 63 6f 72  65 4d 61 74 68 22 3a 37  |2","ScoreMath":7|
 00000060  30 2c 22 53 63 6f 72 65  45 6e 67 6c 69 73 68 22  |0,"ScoreEnglish"|
 00000070  3a 39 30 7d 2c 7b 22 49  44 22 3a 31 2c 22 4e 61  |:90},{"ID":1,"Na|
 00000080  6d 65 22 3a 22 4a 61 6b  65 33 22 2c 22 53 63 6f  |me":"Jake3","Sco|
 00000090  72 65 4d 61 74 68 22 3a  38 30 2c 22 53 63 6f 72  |reMath":80,"Scor|
 000000a0  65 45 6e 67 6c 69 73 68  22 3a 31 30 30 7d 2c 7b  |eEnglish":100},{|
 000000b0  22 49 44 22 3a 31 2c 22  4e 61 6d 65 22 3a 22 4a  |"ID":1,"Name":"J|
 000000c0  61 6b 65 34 22 2c 22 53  63 6f 72 65 4d 61 74 68  |ake4","ScoreMath|
 000000d0  22 3a 35 30 2c 22 53 63  6f 72 65 45 6e 67 6c 69  |":50,"ScoreEngli|
 000000e0  73 68 22 3a 34 30 7d 5d                           |sh":40}]|
}

## msgpack.MarshalAsArray 시 메모리 용량 41
([]uint8) (len=41 cap=41) {
 00000000  94 94 01 a5 4a 61 6b 65  31 64 64 94 01 a5 4a 61  |....Jake1dd...Ja|
 00000010  6b 65 32 46 5a 94 01 a5  4a 61 6b 65 33 50 64 94  |ke2FZ...Jake3Pd.|
 00000020  01 a5 4a 61 6b 65 34 32  28                       |..Jake42(|
}
```



