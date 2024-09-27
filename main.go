package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type KeyValue struct {
	Key   string
	Value string
}

// 添加列表
func AddList(name string) []*KeyValue {
	result := make([]*KeyValue, 0)
	return result
}

// 在列表中添加值
func AddVal(Key string, Val string) *KeyValue {
	fmt.Println("键值已添加")
	return &KeyValue{
		Key:   Key,
		Value: Val,
	}
}

// 查找值是否存在
func Setnx(name string, list []*KeyValue) (int, string) {
	for _, j := range list {
		if j.Key == name {
			return 0, j.Value
		}
	}
	return 1, ""
}

// 获取指定列表中的值
func Get(name string, list []*KeyValue) (string, error) {
	for _, j := range list {
		if j.Key == name {
			return j.Value, nil
		}
	}
	return "1", errors.New("key not found")
}

// 删除指定列表中的值
func Del(name string, list []*KeyValue) (int, error) {
	for i, j := range list {
		if j.Key == name {
			return i, nil
		}
	}

	return 0, errors.New("key not found")
}

// 遍历输出指定列表中从start到end的键值对
func Lrange(start string, end string, list []*KeyValue) {
	s := false
	for _, j := range list {
		if j.Key == start {
			s = true
		}
		if s {
			fmt.Println(j.Key, ":", j.Value)
		}
		if j.Key == end {
			break
		}
	}
}

// 用于打印菜单的两个函数
func Mainmenu() {
	fmt.Println("-----菜单-----")
	fmt.Println("1.命令行模块")
	fmt.Println("2.使用说明")
	fmt.Println("3.清空命令行")
	fmt.Println("4.退出程序")
}

func CommandLineMenu() {
	fmt.Println("set listname key value:储存值")
	fmt.Println("setnx listname key value:如果键存在 返回 0 如果键不存在 返回 1 并存储值")
	fmt.Println("get listname key:获得对应键所对应的值")
	fmt.Println("del listname key:删除所对应的键和值")
	fmt.Println("lpush listName:添加一个列表")
	fmt.Println("range listName start (起始位置) end(结束位置):获得一个列表 从 start到end 的所有元素")
	fmt.Println("exit:返回菜单")
}

// 用于读取指令
func ReadWords() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')
	command = strings.TrimSpace(command)
	if err != nil {
		return "", err
	}
	return command, nil
}

// 用于分析指令，并处理的函数
func AnalysisZone(mainList map[string][]*KeyValue) error {
	command, err := ReadWords()

	if err != nil {
		return err
	} else {
		switch command {
		case "1":
			fmt.Println("正在启动原神")
			fmt.Println("启动成功")
			val, _ := AnalysisCommand(mainList)
			for val != 0 {
				val, _ = AnalysisCommand(mainList)
			}
			fmt.Println("退出成功")
		case "2":
			CommandLineMenu()
		case "3":
			fmt.Println("没做")
		case "4":
			fmt.Println("正在退出，请稍等")
			err = WriteJson(&mainList)
			if err != nil {
				fmt.Println("保存失败，无法退出")
				fmt.Println("但你可以玩原神")
				fmt.Println("错误捕获：", err)
				return err
			}
			os.Exit(0)
		default:
			fmt.Println("未知指令")
			return errors.New("undefined command")
		}
	}
	return nil
}
func AnalysisCommand(mainList map[string][]*KeyValue) (int, error) {
	command, err := ReadWords()
	if err != nil {
		return 1, err
	} else {
		list := strings.Split(command, " ")
		switch list[0] {
		case "set":
			if len(list) != 4 {
				fmt.Println("输入错误，请重新输入")
				fmt.Println("输入格式应为：set listname key value")
				return 1, errors.New("error input")
			}
			mainList[list[1]] = append(mainList[list[1]], AddVal(list[2], list[3]))
			fmt.Println("添加完成")
		case "setnx":
			if len(list) != 4 {
				fmt.Println("输入错误，请重新输入")
				fmt.Println("输入格式应为：setnx listname key value")
				return 1, errors.New("error input")
			}
			n, key := Setnx(list[2], mainList[list[1]])
			if n == 1 {
				fmt.Println("键不存在")
				mainList[list[1]] = append(mainList[list[1]], AddVal(list[2], list[3]))
			} else {
				if list[3] == key {
					fmt.Println("键值存在且正确")
				} else {
					fmt.Println("键存在，值应为：", key)
				}
			}
		case "get":
			if len(list) != 3 {
				fmt.Println("输入错误，请重新输入")
				fmt.Println("输入格式应为：get listname key")
				return 1, errors.New("error input")
			}
			val, err := Get(list[2], mainList[list[1]])
			if err != nil {
				fmt.Println("出错啦啊啊啊啊啊，我要玩原神，我不要写代码！")
				fmt.Println("错误捕获：", err)
				return 1, err
			}
			fmt.Println("键值为：", val)
		case "del":
			if len(list) != 3 {
				fmt.Println("输入错误，请重新输入")
				fmt.Println("输入格式应为：del listname key")
				return 1, errors.New("error input")
			}
			i, err := Del(list[2], mainList[list[1]])
			if err != nil {
				fmt.Println("出错咯")
				fmt.Println("错误捕获：", err)
			}
			mainList[list[1]] = append(mainList[list[1]][:i], mainList[list[1]][i+1:]...)
			fmt.Println("删除成功")
		case "lpush":
			if len(list) != 2 {
				fmt.Println("输入错误，请重新输入")
				fmt.Println("输入格式应为：lpush listname")
				return 1, errors.New("error input")
			}
			mainList[list[1]] = AddList(list[1])
			fmt.Println("添加成功")
		case "range":
			if len(list) != 4 {
				fmt.Println("输入错误，请重新输入")
				fmt.Println("输入格式应为：range listName start end")
				return 1, errors.New("error input")
			}
			Lrange(list[2], list[3], mainList[list[1]])
		case "exit":
			if len(list) != 1 {
				fmt.Println("输入错误，请重新输入")
				fmt.Println("输入格式应为：exit")
				return 1, errors.New("error input")
			}
			return 0, nil
		default:
			fmt.Println("未知输入")
			return 1, errors.New("undefined input")
		}
	}
	return 1, nil
}

// 去json化
func Readjson(mainList *map[string][]*KeyValue) *map[string][]*KeyValue {
	f, err := os.Open("Data.json")
	//打开并读取文件
	if err != nil {
		fmt.Println("呀，出错咯，可惜没有工作人员:)")
		fmt.Println("错误捕获:", err)
		return nil
	} else {
		defer f.Close()
		read, err := io.ReadAll(f)
		if err != nil {
			fmt.Println("程序错误...错误...")
			fmt.Println("错误捕获：", err)
			return nil
		}
		//读取json文件
		err = json.Unmarshal(read, mainList)
		if err != nil {
			fmt.Println("出错咯，错误捕获: ", err)
			return nil
		}
	}
	return mainList
}

// json化
func WriteJson(mainList *map[string][]*KeyValue) error {
	read, err := json.Marshal(mainList)
	if err != nil {
		fmt.Println("累了，还是玩原神吧")
		fmt.Println("错误捕获", err)
		return err
	}
	err = os.WriteFile("Data.json", read, 0644)
	if err != nil {
		fmt.Println("写入失败，原神启动！")
		return err
	}
	fmt.Println("保存成功")
	return nil
}

func main() {
	mainList := make(map[string][]*KeyValue, 0)
	Readjson(&mainList)
	for {
		Mainmenu()
		AnalysisZone(mainList)
	}
}
