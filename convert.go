package imago

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// ByteToInt
func (c *iConvert) ByteToInt(data []byte) int {
	var result int32
	bBuf := bytes.NewBuffer(data)
	binary.Read(bBuf, binary.BigEndian, &result)
	return int(result)
}

// ByteToMap
func (c *iConvert) ByteToMap(data []byte) (map[string]string, error) {
	result := map[string]string{}
	err := json.Unmarshal(data, &result)
	return result, err
}

// ByteToInterface
func (c *iConvert) ByteToInterface(data []byte) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	err := json.Unmarshal(data, &result)
	return result, err
}

// IntToByte 数字转字节
func (c *iConvert) IntToByte(data int) []byte {
	result := rune(data)
	bBuf := bytes.NewBuffer([]byte{})
	binary.Write(bBuf, binary.BigEndian, result)
	return bBuf.Bytes()
}

// InterfaceToString
func (c *iConvert) InterfaceToString(data []interface{}) []string {
	result := []string{}
	for _, item := range data {
		result = append(result, item.(string))
	}
	return result
}

// MapToInterface
func (c *iConvert) MapToInterface(data map[string]string) map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range data {
		result[k] = v
	}
	return result
}

// MapToList
func (c *iConvert) MapToList(data []map[string]string, key string) []string {
	result := []string{}
	for _, item := range data {
		result = append(result, item[key])
	}
	return result
}

// MapToURL
func (c *iConvert) MapToURL(data map[string]interface{}) string {
	key := []string{}
	for k := range data {
		key = append(key, k)
	}
	sort.Strings(key)
	result := ""
	for _, k := range key {
		v := fmt.Sprintf("%v", data[k])
		if v != "" {
			result = result + "&" + k + "=" + v
		}
	}
	return result[1:]
}

// StructToURL
func (c *iConvert) StructToURL(data interface{}) string {
	structKey := reflect.TypeOf(data)
	structValue := reflect.ValueOf(data)
	length := structKey.NumField()
	key := []string{}
	for i := 0; i < length; i++ {
		key = append(key, structKey.Field(i).Name)
	}
	sort.Strings(key)
	result := ""
	for _, k := range key {
		vData := structValue.FieldByName(k)
		vType := vData.Type().String()
		vString := ""
		switch vType {
		case "string":
			vString = vData.String()
		case "int":
			vString = strconv.Itoa(int(vData.Int()))
		default:
			vString = fmt.Sprintf("%v", vData)
		}
		if vString != "" {
			result = result + "&" + k + "=" + vString
		}
	}
	return result[1:]
}

// Hash hash
func (c *iConvert) StringToHash(data string) uint32 {
	return crc32.ChecksumIEEE([]byte(data))
}

//RegexParam 正则过滤
func (c *iConvert) StringToRegex(param, pattern string) (result bool, err error) {
	reg, err := regexp.Compile(pattern)
	if err == nil && reg.MatchString(param) == true {
		result = true
	} else {
		result = false
	}
	return
}

//StringToDevice 获取用户设备号
func (c *iConvert) StringToDevice(deviceInput string) (deviceOutput string) {
	deviceOutput = strings.Replace(deviceInput, "<", "", -1)
	deviceOutput = strings.Replace(deviceOutput, ">", "", -1)
	deviceOutput = strings.Replace(deviceOutput, " ", "", -1)
	return
}

// StringToString
func (c *iConvert) StringToString(data string, start, end int) string {
	dataList := []rune(data)
	dataLength := len(data)
	if start < 0 {
		start = 0
	}
	if end > dataLength {
		end = dataLength
	} else if end < 0 {
		end = dataLength + end
	}
	return string(dataList[start:end])
}
