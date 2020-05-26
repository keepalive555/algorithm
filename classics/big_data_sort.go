// 大文件排序
package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
)

const (
	PartitionSize = 1 * 1024 * 1024 // 1MB，分区大小，单位：字节
)

// 快速排序
func QuickSort(arrs []int, low, high int) {
	if low >= high {
		return
	}
	pivot := arrs[low]
	i := low
	j := high

	for j > i {
		for j > i && arrs[j] >= pivot {
			j--
		}
		arrs[i] = arrs[j]
		for j > i && arrs[i] <= pivot {
			i++
		}
		arrs[j] = arrs[i]
	}

	arrs[i] = pivot
	QuickSort(arrs, low, i-1)
	QuickSort(arrs, i+1, high)

	return
}

func Partition(file string) ([]string, error) {
	// 文件分区，切割处理
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// 文件前缀
	prefix := f.Name()
	// 切割文件
	reader := bufio.NewReader(f)
	// 文件序号
	sequence := 0
	// 文件列表
	files := []string{}
	quit := false
	for !quit {
		// 已读取字节数
		size := 0
		// 分区文件
		partition := fmt.Sprintf("%s.%d", prefix, sequence)
		// 创建分区文件
		pf, err := os.OpenFile(partition, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			return nil, err
		}

		// 读取分区大小的数据
		buffer := make([]int, 0, 1000)
		for size < PartitionSize {
			line, err := reader.ReadBytes('\n')
			if err == io.EOF {
				quit = true
				break
			}
			// 剔除尾部换行符\n
			n, err := strconv.ParseInt(string(line[:len(line)-1]), 10, 64)
			if err != nil {
				fmt.Println("parse int convert error: %+v\n", err)
				continue
			}
			buffer = append(buffer, int(n))
			size += len(line)
		}
		// 内排序
		QuickSort(buffer, 0, len(buffer)-1)
		// 写入分区文件
		for i := 0; i < len(buffer); i++ {
			pf.WriteString(fmt.Sprintf("%d\n", buffer[i]))
		}
		// 保存分区文件
		pf.Close()

		files = append(files, partition)
		sequence++
	}
	return files, nil
}

// 随机生成测试文件
func TestDataGenerator(file string, size int) error {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	writeBytes := 0
	for writeBytes < size {
		data := fmt.Sprintf("%d\n", rand.Intn(0xffffff00))
		writeBytes += len(data)
		_, err := f.WriteString(data)
		if err != nil {
			break
		}
	}
	// 保存文件
	f.Sync()
	return nil
}

// 归并排序
func MergeSort(files []string) (string, error) {
	if len(files) == 0 {
		return "", nil
	}
	if len(files) == 1 {
		return files[0], nil
	}
	// 打开所有文件
	n := len(files)
	fileList := make([]*os.File, n, n)
	for i := 0; i < n; i++ {
		f, err = os.Open(files[i])
		if err != nil {
			return "", err
		}
		defer f.Close()
		fileList[i] = f
	}
	// 利用败者树进行归并排序
}

// --------------------------败者树----------------------------------------
// --------------------------败者树----------------------------------------

func main() {
	file := "./id_list.txt"
	// TestDataGenerator(file, 20*1024*1024)
	files, err := Partition(file)
	fmt.Printf("%+v, %+v\n", files, err)
}
