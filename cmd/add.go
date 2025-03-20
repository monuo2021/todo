/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/monuo2021/todo/include"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [任务名称]",
	Short: "添加新任务",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	var items []include.Item

	// 检查文件是否存在
	if _, err := os.Stat(viper.GetString("dataFile")); err == nil {
		// 文件存在时加载已有的代办事项
		items, err = include.LoadItems(viper.GetString("dataFile"))
		if err != nil {
			log.Printf("数据加载失败（已忽略加载操作）: %v", err)
		}
	} else if os.IsNotExist(err) {
		// 文件不存在时初始化空列表
		items = []include.Item{}
	} else {
		// 其他错误（如权限问题）
		log.Fatalf("文件状态检查失败: %v", err)
	}

	// 获取新添加的代办事项
	for _, arg := range args {
		item := include.Item{Text: arg}
		item.SetPriority(priority)
		items = append(items, item)
	}

	// 保存新的代办事项
	err := include.SaveItems(viper.GetString("dataFile"), items)
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "优先级:1,2,3")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
