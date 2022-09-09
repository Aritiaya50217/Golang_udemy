package report

import (
	"fmt"
	"reflect"
	"strings"
)

// Text : generate report string
func Text(v interface{}) string {
	// Type
	// Value
	val := reflect.ValueOf(v)
	var sb strings.Builder
	text(&sb, &val)
	return sb.String()
}

func text(sb *strings.Builder, v *reflect.Value) {
	switch v.Kind() {
	// กรณี type เดิมเป็น struct
	case reflect.Struct:
		structName := v.Type().Name()
		// ถ้าไม่กำหนด ชื่อของ struct ให้ตั้งเป็น anonymous
		if len(structName) == 0 {
			structName = "anonymous"
		}
		sb.WriteString(structName + " ( ")
		// loop เพื่อนับ field in struct
		for i := 0; i < v.NumField(); i++ {
			val := v.Field(i)
			if i > 0 {
				sb.WriteString(" , ")
			}
			// แสดงชื่อ field เป็นภาษาไทย
			displayName := v.Type().Field(i).Name
			// report (TagName in struct) => `report:"ชื่อ"`
			tagName, ok := v.Type().Field(i).Tag.Lookup("report")
			if ok {
				// multi tag
				tags := strings.Split(tagName, ",")
				if len(tags) == 2 {
					displayName = tags[0]
					// `report:"ชื่อ,uppercase"` ** ห้ามเว้นวรรคคั่น
					// tag[0] = ชื่อ , tag[1] = uppercase
					if tags[1] == "uppercase" {
						tmp := v.Field(i).String()
						upper := strings.ToUpper(tmp)
						val = reflect.ValueOf(upper)
					}
				}
			}
			sb.WriteString(displayName + " : ")
			// val = v.Field(i)
			text(sb, &val)
		}
		sb.WriteString(" )")
	// กรณี type เดิมเป็น string
	case reflect.String:
		sb.WriteString(v.String())
	// type int
	case reflect.Int:
		sb.WriteString(fmt.Sprintf("%d", v.Int()))
	}
}
