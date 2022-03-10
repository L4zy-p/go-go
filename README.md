# ลืมไปแล้วอ่ะ ~

## มาเริ่มกันใหม่ start

สร้าง mod มาก่อนเลย mode นี่ทำหน้าที่คล้าย package.json คือเก็บ dependencies ที่ใช้

```
go mod init [name]
```

สร้าง file main.go แล้วลองโค้ด


```
package main // นี่คือกรุ๊ป function

// เป็น package ที่จะใช้งาน
import (
	"fmt"

	"rsc.io/quote"
)

// เป็น func default ตัวแรกในการรัน package main
func main() {
	fmt.Println("Hello i can remember")
	fmt.Println(quote.Go())
}
```

run ดู

```
go run main.go
```

ตอนแรกเรา error จะไม่เจอ package rsc.io/quote
ให้เราลอง run คำสั่ง

rsc.io/quote เป็น external package ซึ่งเราต้องติดตั้งก่อนเสมอ เหมือน npm install ...

```
go mod tidy
```

จากนั้นจะทำการ install package ที่เราใช้ให้ auto

## สร้าง go module

module ก็คือส่วนที่ทำงานแยกกัน หรือ สามารถนำมาใช้ร่วมกันได้

```
go mod init example.com/greetings
```

การใช้คำสั่งนี้จะเป็นการ publish ขึ้นไปข้างบน เราสามารถ download ลงมาจาก Go tools ได้

สร้าง file greeting.go แล้วโค้ด

```
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
```



