package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// 定义根命令 `quest`
// Use: "quest"代表的是名称
var rootCmd = &cobra.Command{
	Use:   "quest",
	Short: "Turn your todos into RPG quests",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		// os返回0是正常退出，非1是异常退出
		os.Exit(1)
	}
}
