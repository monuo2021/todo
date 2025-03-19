/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/monuo2021/todo/include"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有任务",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(dataFile); err == nil {
		items, err := include.LoadItems(dataFile)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		fmt.Println(items)
	} else if os.IsNotExist(err) {
		log.Println("不存在代办事项")
	} else {
		// 其他错误（如权限问题）
		log.Fatalf("文件状态检查失败: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
