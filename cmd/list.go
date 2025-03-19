/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

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
	var items []include.Item

	// 检查文件是否存在
	if _, err := os.Stat(dataFile); err == nil {
		items, err = include.LoadItems(dataFile)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	} else if os.IsNotExist(err) {
		log.Println("不存在代办事项")
	} else {
		// 其他错误（如权限问题）
		log.Fatalf("文件状态检查失败: %v", err)
	}

	sort.Sort(include.ByPri(items))

	// 创建 tabwriter 实例，配置对齐参数
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	// 遍历 items 切片中的每个待办事项
	for _, item := range items {
		line := item.Label() + "\t" + item.PrettyP() + "\t" + item.Text + "\t"
		// 写入缓冲区（非立即输出）
		fmt.Fprintln(w, line)
	}

	// 应用格式规则并刷新输出到终端
	w.Flush()
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
