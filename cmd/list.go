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
	"github.com/spf13/viper"
)

var (
	doneOpt bool
	allOpt  bool
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
	if _, err := os.Stat(viper.GetString("dataFile")); err == nil {
		items, err = include.LoadItems(viper.GetString("dataFile"))
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
		// allOpt 为 true 时，显示所有任务；allOpt 为 false 时，接下来判断 doneOpt
		// doneOpt 为 true 时，只显示已完成的任务；doneOpt 为 false 时，显示未完成的任务
		if allOpt || item.Done == doneOpt {
			line := item.Label() + "\t" + item.PrettyDone() + "\t" + item.PrettyP() + "\t" + item.Text + "\t"
			fmt.Fprintln(w, line)
		}
	}

	// 应用格式规则并刷新输出到终端
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all Todos")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
