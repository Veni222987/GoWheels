package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

// golang多线程下载器
func main() {
	// 定义命令行参数
	url := flag.String("url", "", "The URL of the file to download")
	numParts := flag.Int("parts", 4, "The number of parts to split the download into")
	flag.Parse()

	// 检查必需的参数
	if *url == "" {
		fmt.Println("Error: URL is required")
		flag.Usage()
		return
	}

	// 获取文件大小
	resp, err := http.Head(*url)
	if err != nil {
		fmt.Println("Error getting file size:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP status", resp.StatusCode)
		return
	}

	fileSize, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		fmt.Println("Error converting Content-Length:", err)
		return
	}

	fmt.Println("File size:", fileSize)

	// 分块下载
	var wg sync.WaitGroup
	partSize := fileSize / *numParts
	for i := 0; i < *numParts; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			start := i * partSize
			end := start + partSize - 1
			if i == *numParts-1 {
				end = fileSize - 1
			}
			err := downloadPart(*url, i, start, end)
			if err != nil {
				fmt.Println("Error downloading part", i, ":", err)
			}
		}(i)
	}

	wg.Wait()

	// 合并文件
	fileName := getFileNameFromURL(*url)
	err = mergeParts(fileName, *numParts)
	if err != nil {
		fmt.Println("Error merging parts:", err)
		return
	}

	fmt.Println("Download completed successfully")
}

// 下载对应的部分
func downloadPart(url string, partNum, start, end int) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent {
		return fmt.Errorf("server did not respond with partial content: %d", resp.StatusCode)
	}

	partFileName := fmt.Sprintf("%s.part%d", getFileNameFromURL(url), partNum)
	partFile, err := os.Create(partFileName)
	if err != nil {
		return err
	}
	defer partFile.Close()

	_, err = io.Copy(partFile, resp.Body)
	return err
}

// 创建新文件并合并下载结果
func mergeParts(fileName string, numParts int) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	for i := 0; i < numParts; i++ {
		partFileName := fmt.Sprintf("%s.part%d", fileName, i)
		partFile, err := os.Open(partFileName)
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, partFile)
		partFile.Close()
		if err != nil {
			return err
		}

		// 删除分块文件
		err = os.Remove(partFileName)
		if err != nil {
			return err
		}
	}

	return nil
}

func getFileNameFromURL(url string) string {
	// 从URL中提取文件名
	tokens := strings.Split(url, "/")
	return tokens[len(tokens)-1]
}
