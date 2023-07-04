package main

import (
	"fmt"
	"github.com/json-iterator/go"
)

func main() {

	data := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	encoded, err := jsoniter.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(encoded))

	// 기본 패키지와 최대한 호환성을 맞춘 방식으로 unMarshal 을 해라
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	decoded := make(map[string]int)
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		panic(err)
	}
	fmt.Println(decoded)
}

/*

func makeConfigJsoniter() {

	var ConfigCompatibleWithStandardLibrary = jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
	}.Froze()
	// 최초 들어올 때에는 json 으로 validation 검사, map_key 를 정렬하는 옵션을 사용할 필요가 있다.
	var ConfigDefault = jsoniter.Config{
		EscapeHTML: true,
	}.Froze()
	// 최초 validation 검사 이후 두번째부터는 이미 검사 이후이기 때문에 validation 과 sort 를 빼줘서 속도를 향상시킨다.

	var ConfigFastest = jsoniter.Config{
		EscapeHTML:                    false,
		MarshalFloatWith6Digits:       true,
		ObjectFieldMustBeSimpleString: true,
	}.Froze()
	// 제일 속도가 빠른 옵션이지만, float 를 6자리까지만 정밀도를 계산하기 때문에 데이터 값이 달라질 수 있다.

}
*/
