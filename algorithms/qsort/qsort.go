//-*- coding: UTF-8 -*-
// Author: afei00123

package qsort

//定义一个切片变量，存放输入值，再定义两个变量left和right进行比较操作
func quickSort(values []int, left, right int)  {
	//临时变量用于存放交换的值
	temp := values[left]
	//设置分界值
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j --
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= temp && i <= p {
			i ++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	//根据分界值进行左右两边排序
	if p - left > 1 {
		quickSort(values, left, p - 1)
	}
	if right -p  > 1 {
		quickSort(values, p + 1, right)
	}
}

//定义一个方法进行传参
func QuickSort(values []int) {
	//初始化left值为0
	quickSort(values, 0, len(values) - 1)
}