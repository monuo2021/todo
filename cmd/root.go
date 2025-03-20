/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dataFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo CLI应用，管理你的待办事项",
	Long: `todo 将帮助你在更短的时间内完成更多的工作。
它被设计得尽可能简单，以帮助你实现你的目标。支持添加、列出、标记完成任务等功能。`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cfgFile string

func initConfig() {
	viper.SetConfigName(".todo")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("todo")

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Use config file:", viper.ConfigFileUsed())
	} else {
		log.Println("No config file found. Using default settings.")
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cobra.OnInitialize(initConfig)

	home, err := os.UserHomeDir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile.")
	}

	// 安全拼接路径
	defaultPath := filepath.Join(home, ".todos.json")

	// 配置Persistent Flags
	rootCmd.PersistentFlags().StringVarP(
		&dataFile,
		"datafile",
		"d",
		defaultPath,
		"存储待办事项的数据文件路径",
	)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo.yaml)")
}
