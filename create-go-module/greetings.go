package greetings

import "fmt"

// name string -> name เป็น params ส่วน string เป็น type params
// func Hello(...) string -> Hello เป็น function ส่วน string เป็น return type
func Hello(name string) string {
	// := เป็น Operator ทางลัดในการประกาศและ initial ค่าในตัวแปร
	// Sprintf เป็น function ที่ return greeting text ให้ กับ message
	// v = verb ซึ่งเป็น text
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message
}
