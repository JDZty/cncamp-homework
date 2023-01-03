// 给定一个字符串数组
// [“I”,“am”,“stupid”,“and”,“weak”]
// 用 for 循环遍历该数组并修改为
// [“I”,“am”,“smart”,“and”,“strong”]
package main

import "fmt"

func main() {
	arr := []string{"I", "am", "stupid", "and", "weak"}

	fmt.Printf("Before modify, arr is %v\n", arr)

	for index, value := range arr {
		if value == "stupid" {
			arr[index] = "smart"
		}
		if value == "weak" {
			arr[index] = "strong"
		}
	}

	fmt.Printf("After modify, arr is %v\n", arr)
}
