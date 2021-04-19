package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	fmt.Print("输入路径：")
	var path string
	if _, err := fmt.Scan(&path); err != nil {
		log.Fatalln(err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	result := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		listName := getListName(line)
		if _, ok := result[listName]; ok {
			result[listName] = append(result[listName], line)
		} else {
			result[listName] = make([]string, 0)
			result[listName] = append(result[listName], line)
		}
		count++
	}
	fmt.Printf("任务总数：%d\n",count)
	for k, v := range result {
		fmt.Printf("类别：%s 任务数：%d 占比：%.2f%%\n", k, len(v), 100*(float64(len(v))/float64(count)))
	}
}

func getListName(source string) string {
	re := regexp.MustCompile(`\<(.+)\>`)
	if listNames := re.FindStringSubmatch(source); len(listNames) > 0 {
		return listNames[0]
	}
	return "empty"
}
