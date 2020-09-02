package language_test

import (
	"fmt"
	"github.com/enorith/language"
)

func ExampleTranslate() {
	lm := map[string]string{
		"name": "please input of your name",
		"age":  "age between :min and :max",
	}

	lm2 := map[string]string{
		"name": "请输入你的姓名",
		"age":  "年龄需在:min和:max之间",
	}

	language.Register("main", "en", lm)
	language.Register("main", "zh-CN", lm2)

	str, e := language.T("main", "name")
	str2, e2 := language.Translate("main", "name", "zh-CN")
	param := map[string]string{
		"min": "18",
		"max": "58",
	}
	str3, e3 := language.T("main", "age", param)
	str4, e4 := language.Translate("main", "age", "zh-CN", param)

	if e != nil || e2 != nil || e3 != nil || e4 != nil {
		fmt.Println(e, e2, e3, e4)
	}
	fmt.Println(str)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)
	// Output:
	// please input of your name
	// 请输入你的姓名
	// age between 18 and 58
	// 年龄需在18和58之间
}
