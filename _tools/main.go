package main

import (
	tools "_tools/option-gen"
)

func main() {
	//var modelName string = "ep-20240816142135-k9j5b"
	//zhhelp, err := ai.Chat(modelName, fmt.Sprintf("请你将我接下来提供的 kubectl --help 的帮助说明内容，翻译为清晰易懂的中文，不需要详细说明。以下是帮助说明原文: %s", "Drain node in preparation for maintenance"))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(zhhelp)
	tools.Run()
}
