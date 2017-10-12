package conversion

import (
	"testing"
	"fmt"
)

func TestTimeNow(t *testing.T) {
	to := TimeNow()
	fmt.Println("the time is ",to)
}

func TestTimeYear(t *testing.T) {
	to := TimeYear()
	fmt.Println("the time is ",to)
}

func TestTimeHour(t *testing.T) {
	to := TimeHour()
	fmt.Println("the time is ",to)
}

func TestUint32ToBytes(t *testing.T) {
	var from uint32 = 10
	to := Uint32ToBytes(from)
	fmt.Printf("TestUint32ToBytes from values is %v, to values is %v \n",from,to)
}

func TestBytesToUint32(t *testing.T) {
	from := make([]byte, 4)
	to := BytesToUint32(from)
	fmt.Printf("TestBytesToUint32 from values is %v, to values is %v \n",from,to)
}

func TestStringToInt(t *testing.T) {
	var from string = "88"
	to,err := StringToInt(from)
	fmt.Printf("TestStringToInt from values is %v, to values is %v ,err is %v \n",from,to,err)
}

func TestStringToUint(t *testing.T) {
	var from string = "88"
	to,err := StringToUint(from)
	fmt.Printf("TestStringToUint from values is %v, to values is %v ,err is %v \n",from,to,err)
}
func TestStringToUint8(t *testing.T) {
	var from string = "88"
	to,err := StringToUint8(from)
	fmt.Printf("TestStringToUint8 from values is %v, to values is %v ,err is %v \n",from,to,err)
}

func TestStringToUint32(t *testing.T) {
	var from string = "88"
	to,err := StringToUint32(from)
	fmt.Printf("TestStringToUint32 from values is %v, to values is %v ,err is %v \n",from,to,err)
}

func TestStringToUint64(t *testing.T) {
	var from string = "88"
	to,err := StringToUint64(from)
	fmt.Printf("TestStringToUint64 from values is %v, to values is %v ,err is %v \n",from,to,err)
}

func TestIntToString(t *testing.T) {
	var from int = 99
	to := IntToString(from)
	fmt.Printf("TestIntToString from values is %v, to values is %v  \n",from,to)
}

func TestInt64ToString(t *testing.T) {
	var from int64 = 99
	to := Int64ToString(from)
	fmt.Printf("TestInt64ToString from values is %v, to values is %v  \n",from,to)
}

func TestUint32ToString(t *testing.T) {
	var from uint32 = 66
	to := Uint32ToString(from)
	fmt.Printf("TestUint32ToString from values is %v, to values is %v  \n",from,to)
}

func TestUint64ToString(t *testing.T) {
	var from uint64 = 66
	to := Uint64ToString(from)
	fmt.Printf("TestUint64ToString from values is %v, to values is %v  \n",from,to)
}

func TestFloat64ToString(t *testing.T) {
	var from float64 = 66.78
	to := Float64ToString(from)
	fmt.Printf("TestFloat64ToString from values is %v, to values is %v  \n",from,to)
}

func TestStringToFloat(t *testing.T) {
	var from string = "11.12"
	to,err := StringToFloat(from)
	fmt.Printf("TestStringToFloat from values is %v, to values is %v, err is %v \n",from,to,err)
}

func TestStringToFloat64(t *testing.T) {
	var from string = "11.1314"
	to,err := StringToFloat64(from)
	fmt.Printf("TestStringToFloat64 from values is %v, to values is %v, err is %v \n",from,to,err)
}

func TestToString(t *testing.T) {
	var from interface {} = 98
	to := ToString(from)
	fmt.Printf("TestToString from values is %v, to values is %v  \n",from,to)
}

func TestToStringSlice(t *testing.T){
	var from []interface{}
	slice := []int{1,2,3,4}
	from = append(from,slice)
	to,err := ToStringSlice(from)
	fmt.Printf("TestToStringSlice from values is %v, to values is %v, err is %v \n",from,to,err)
}

func TestToBool(t *testing.T) {
	var from interface{} = "t"
	to,err := ToBool(from)
	fmt.Printf("TestToBool from values is %v, to values is %v, err is %v \n",from,to,err)
}

func TestToFloat32(t *testing.T) {
	var from interface{} = 35
	to,err := ToFloat32(from)
	fmt.Printf("TestToFloat32 from values is %v, to values is %v, err is %v \n",from,to,err)
}

func TestToFloat64(t *testing.T) {
	var from interface{} = 35.00
	to,err := ToFloat64(from)
	fmt.Printf("TestToFloat64 from values is %v, to values is %v, err is %v \n",from,to,err)
}

func TestToInt8(t *testing.T) {
	var from interface{} = 35.00
	to,err := ToInt8(from)
	fmt.Printf("TestToInt8 from values is %v, to values is %v, err is %v \n",from,to,err)
}

func TestToInt16(t *testing.T) {
	var from interface{} = 35.00
	to,err := ToInt16(from)
	fmt.Printf("TestToInt16 from values is %v, to values is %v, err is %v \n",from,to,err)
}

func TestToInt(t *testing.T) {
	var from interface{} = 35.00
	to,err := ToInt(from)
	fmt.Printf("TestToInt from values is %v, to values is %v, err is %v \n",from,to,err)
}

func TestToInt32(t *testing.T) {
	var from interface{} = 35.00
	to,err := ToInt32(from)
	fmt.Printf("TestToInt32 from values is %v, to values is %v, err is %v \n",from,to,err)
}

func TestToInt64(t *testing.T) {
	var from interface{} = 35.00
	to,err := ToInt64(from)
	fmt.Printf("TestToInt64 from values is %v, to values is %v, err is %v \n",from,to,err)
}

func TestToInt64Slice(t *testing.T) {
	var from []interface{}
	for i := 0;i < 10;i++{
		from = append(from,i)
	}
	to,err := ToInt64Slice(from)
	fmt.Printf("TestToInt64Slice from values is %v, to values is %v, err is %v \n",from,to,err)
}

func TestToUint32Slice(t *testing.T) {
	var from []interface{}
	for i := 0;i < 10;i++{
		from = append(from,i)
	}
	to,err := ToUint32Slice(from)
	fmt.Printf("TestToUint32Slice from values is %v, to values is %v, err is %v \n",from,to,err)
}

func TestToUint64Slice(t *testing.T) {
	var from []interface{}
	for i := 0;i < 10;i++{
		from = append(from,i)
	}
	to,err := ToUint64Slice(from)
	fmt.Printf("TestToUint64Slice from values is %v, to values is %v, err is %v \n",from,to,err)
}



