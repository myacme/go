package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

//func main() {
//	//writeFile()
//	//writeFileByLine()
//	//appendToFile()
//	//getFileInfo()
//	//listFiles()
//	createTempFilesAndDirectories()
//
//}

// 创建文件
func createFile() {
	// 创建文件，如果文件已存在会被截断（清空）
	file, err := os.Create("./file/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // 确保文件关闭
	log.Println("文件创建成功")
}

// 打开文件
func openFile() {
	// 打开文件
	file, err := os.Open("./file/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file) // 确保文件关闭
	log.Println("文件打开成功")
}

// 逐行读取文件
func readFileByLine() {
	file, err := os.Open("./file/test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line = 1
	for scanner.Scan() {
		fmt.Print("第", line, "行: ")
		line++
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// 一次性读取整个文件
func readFile() {
	content, err := os.ReadFile("./file/test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(content))
}

func writeFile() {
	// 方式1：直接写入字符串
	file, err := os.Create("./file/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString("直接写入字符串\n")

	// 方式2：写入字节切片
	data := []byte("写入字节切片\n")
	file.Write(data)

	// 方式3：使用fmt.Fprintf格式化写入
	fmt.Fprintf(file, "格式化写入: %d\n", 123)
}

// 逐行写入文件
func writeFileByLine() {
	file, err := os.Create("./file/test.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, "Hello, World!")
	fmt.Fprintln(writer, "Hello, Go!")
	writer.Flush()
}

// 追加内容到文件
func appendToFile() {
	file, err := os.OpenFile("./file/test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("Appended text\n"); err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	fmt.Println("Text appended successfully!")
}

// 删除文件
func deleteFile() {
	err := os.Remove("./file/test.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Println("File deleted successfully!")
}

// 获取文件信息
func getFileInfo() {
	fileInfo, err := os.Stat("./file/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("文件名:", fileInfo.Name())
	fmt.Println("文件大小:", fileInfo.Size(), "字节")
	fmt.Println("权限:", fileInfo.Mode())
	fmt.Println("最后修改时间:", fileInfo.ModTime())
	fmt.Println("是目录吗:", fileInfo.IsDir())
}

// 检查文件是否存在
func checkFileExist() {
	if _, err := os.Stat("./file/test.txt"); os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("文件存在")
	}
}

// 重命名和移动文件
func renameFile() {
	err := os.Rename("./file/test.txt", "./file/new.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("重命名成功")
}

func copyFile() {
	srcFile, err := os.Open("./file/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create("./file/copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	bytesCopied, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("复制完成，共复制 %d 字节", bytesCopied)
}

// 列出目录下的文件和目录
func listFiles() {
	entries, err := os.ReadDir("./file")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		info, _ := entry.Info()
		fmt.Printf("%-20s %8d %v\n",
			entry.Name(),
			info.Size(),
			info.ModTime().Format("2006-01-02 15:04:05"))
	}
}

// 删除目录
func deleteDirectory() {
	// 删除空目录
	err := os.Remove("emptydir")
	if err != nil {
		log.Fatal(err)
	}

	// 递归删除目录及其内容
	err = os.RemoveAll("./file")
	if err != nil {
		log.Fatal(err)
	}
}

// 创建临时文件和目录
func createTempFilesAndDirectories() {
	// 创建临时文件
	tmpFile, err := os.CreateTemp("", "./file/example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // 清理

	fmt.Println("临时文件:", tmpFile.Name())

	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir) // 清理

	fmt.Println("临时目录:", tmpDir)
}
