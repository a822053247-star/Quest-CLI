package main

import "quest/cmd"

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	cmd.Execute()
	//fmt.Println("Quest CLI")
	//q := quest.Quest{
	//	ID:        1,
	//	Title:     "Read paper",
	//	XP:        30,
	//	Boss:      false,
	//	Completed: false,
	//	CreatedAt: time.Now(),
	//}
	//quests := []quest.Quest{q}
	//err := storage.SaveQuests(quests)
	//if err != nil {
	//	fmt.Println("保存失败:", err)
	//	return
	//}
	//fmt.Println("保存成功")
	//loadedQuests, err := storage.LoadQuests()
	//if err != nil {
	//	fmt.Println("加载失败:", err)
	//	return
	//}
	//fmt.Println("加载成功", loadedQuests)
}
