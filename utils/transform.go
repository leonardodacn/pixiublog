package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// StructToMap transform struct to map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

// StructsToMap transform structs to maps
func StructsToMap(objs []interface{}) []map[string]interface{} {
	var datas = make([]map[string]interface{}, len(objs))
	for k, v := range objs {
		var data = make(map[string]interface{})
		keys := reflect.TypeOf(v)
		values := reflect.ValueOf(v)
		for i := 0; i < keys.NumField(); i++ {
			data[keys.Field(i).Name] = values.Field(i).Interface()
		}
		datas[k] = data
	}
	return datas
}

func StructToJson(obj interface{}) string {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("StructToJson err: ", err)
	}
	return string(jsonBytes)
}

func JsonToStruct(jsonStr string, data interface{}) {
	json.Unmarshal([]byte(jsonStr), &data)
}

func JsonToMap(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMap err: ", err)
	}
	return mapResult
}

// StructToJsonThenMap transform to json then transform to map, it can use json tag to control the result
func StructToJsonThenMap(obj interface{}) map[string]interface{} {
	jsonStr := StructToJson(obj)
	data := JsonToMap(jsonStr)
	return data
}
