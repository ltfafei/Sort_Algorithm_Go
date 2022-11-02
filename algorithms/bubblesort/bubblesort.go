//-*- coding: UTF-8 -*-
// Author: afei00123

package bubblesort

func Bubblesort(values []int) {
	//设置一个标志位
	flag := true

	for i := 0; i < len(values) - 1; i ++ {
		flag = true
		for j := 0; j < len(values) - i - 1; j ++ {
			//如果索引j的值大于索引j+1的值的话，那将他们位置进行交换
			if values[j] > values[j + 1] {
				values[j], values[j + 1] = values[j + 1], values[j]
				flag = false
			}
		}
		//排序完成则跳出循环
		if flag == true {
			break
		}
	}
}