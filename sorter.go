//-*- coding: UTF-8 -*-
// Author: afei00123

package main

import (
	"Sort_Algorithm_Go/algorithms/bubblesort"
	"Sort_Algorithm_Go/algorithms/qsort"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

//flag模块进行传参
var infile *string = flag.String("i", "data\\unsorted.txt", "Please input need sorting files")
var outfile = flag.String("o", "data\\sorted.txt", "Please input file path to save.")
var algorithm = flag.String("a", "qsort", "Please Sort algorihm. qsort or bubblesort.")

//从文件中读取数据
func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Faild to open the input file ", infile)
	}
	defer file.Close()
	//bufio提供了按行读取文件功能
	br := bufio.NewReader(file)
	//Go语言提供的内置函数make()可以用于灵活地创建数组切片
	values = make([]int, 0)
	//异常处理
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			//文件读取完成会报EOF
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected!")
			return
		}
		//字符数组转换为字符串
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		//append()函数可以从尾部增加元素，从而生成一个新的数组切片。
		values = append(values, value)
	}
	fmt.Println(values)
	return
}

//写入保存数据
func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create output file ", outfile)
		return err
	}
	defer file.Close()

	//按行for循环写入文件
	for _, value := range values {
		//Itoa常量计数器，初始值为零，遇到const重置为零
		str := strconv.Itoa(value)
		_, _ = file.WriteString(str + "\n")
	}
	return nil
}

func main()  {
	flag.Parse()

	if infile != nil {
		fmt.Println("	Please enter -h to view the help information.\n")
		fmt.Println("	USEG: -i c:\\unsort.txt -o c:\\sort.txt -a qsort")
		fmt.Println("	infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm, "\n")
	}
	values, err := readValues(*infile)

	//接收参数选择对应的算法进行排序
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
			case "qsort":
				qsort.QuickSort(values)
			case "bubblesort":
				bubblesort.Bubblesort(values)
			default:
				fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported!")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs ", t2.Sub(t1), " to complete.")
		err := writeValues(values, *outfile)
		if err != nil {
			return
		}
	} else {
		fmt.Println(err)
	}
}