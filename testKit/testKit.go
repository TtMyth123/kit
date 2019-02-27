package testKit

import (
	"encoding/json"
	"io/ioutil"
)

type TestData struct {
	DATA map[string]interface{}
}

func (this *TestData) Reload(fileName string) {
	//读取用户自定义配置
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, this)
	if err != nil {
		panic(err)
	}
}
func (this *TestData) UnmarshalMap(key string, typeData interface{}) error {
	return UnmarshalMap(this.DATA, key, typeData)
}

/**
mapData:Map数据. key:Map数据的key. typeData:要填充的数据
使用方法:

mapData := map[string]interface{};

mapData["A"] = "aaa";

var string v;

UnmarshalMap(mapData,"A", &v);
*/
func UnmarshalMap(mapData map[string]interface{}, key string, typeData interface{}) error {
	data := mapData[key]
	byteData, _ := json.Marshal(data)
	return json.Unmarshal(byteData, typeData)
}

var TestDataOB *TestData

func init() {
	TestDataOB = &TestData{}
	//TestDataOB.Reload("testData.json")
}
