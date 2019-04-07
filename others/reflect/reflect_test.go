package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int8
}

func (p *Person) UpdateAge(age int8) {
	p.Age = age
}

func (p *Person) PrintProfile() {
	fmt.Printf("A person Named: %s, Age: %d\n", p.Name, p.Age)
}

func TestReflect(t *testing.T) {
	p := &Person{"George", 30}

	// 获取struct字段
	fieldName, ok := reflect.TypeOf(*p).FieldByName("Name")
	if ! ok {
		t.Error("field Name is not exists.")
		return
	}

	t.Log(fieldName)

	// 调用方法
	reflect.ValueOf(p).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(int8(18))})

	p.PrintProfile()

	checkType(p)
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	b := map[int]string{
		1: "a",
		2: "b",
		4: "c",
	}

	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{3, 2, 1}

	t.Log(reflect.DeepEqual(s1, s2))
}

// 根据反射机制给不同结构体内需要的字段进行填充
func TestZeusConfig(t *testing.T) {
	configs := map[string]interface{}{"ApiUrl": "http://api.ltwen.com"}

	type Customer struct {
		Name   string
		ApiUrl string
	}

	type VIP struct {
		AccountName string
		Level       int8
		ApiUrl      string
	}

	c := &Customer{Name: "Bob"}

	if err := fillConfig(c, configs); err != nil {
		t.Error(err)
	}

	t.Log(c)

	v := new(VIP)
	v.AccountName = "George"
	if err := fillConfig(v, configs); err != nil {
		t.Error(err)
	}

	t.Log(v)
}

func fillConfig(oriStruct interface{}, configs map[string]interface{}) error {
	// 判断类型是否为指针结构体
	if reflect.TypeOf(oriStruct).Kind() != reflect.Ptr {
		return errors.New("结构体必须为指针结构体")
	}

	// 遍历map配置，如果存在字段相同、值类型也相同，则进行赋值
	for configName, configValue := range configs {

		fieldStruct, isExists := reflect.TypeOf(oriStruct).Elem().FieldByName(configName)

		if ! isExists {
			continue
		}

		configValueType := reflect.TypeOf(configValue)
		if fieldStruct.Type != configValueType {
			continue
		}

		reflect.ValueOf(oriStruct).Elem().FieldByName(configName).Set(reflect.ValueOf(configValue))
	}

	return nil
}

func checkType(v interface{}) {
	t := reflect.TypeOf(v)

	switch tp := t.Kind(); tp {
	case reflect.Int32, reflect.Int64:
		fmt.Println("int32 or int64")
	case reflect.Float32, reflect.Float64:
		fmt.Println("float32 or float 64")
	default:
		fmt.Println("unknow type:", tp)
	}
}
