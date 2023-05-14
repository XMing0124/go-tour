package cmd

import (
	"github.com/go-programming-tour-book/tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderLineToUpperCamelCase
	ModeUnderLineToLowerCamelCase
	ModeCamelCaseToUnderLine
)

var desc = strings.Join([]string{
	"改子命令支持各种格式单词转换，模式如下:",
	"1:全部单词转换为大写",
	"2:全部单词转换为小写",
	"3:下划线转首字母大写驼峰",
	"4:下划线转首字母小写驼峰",
	"5:驼峰转下划线",
}, "\n")

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeCamelCaseToUnderLine:
			content = word.CamelCaseToUnderline(str)
		case ModeUnderLineToLowerCamelCase:
			content = word.UnderLineToLowerCamelCase(str)
		case ModeUnderLineToUpperCamelCase:
			content = word.UnderlineToUpperCamelCase(str)
		default:
			log.Fatal("暂不支持该格式转换,请执行help word查看帮助文档")
		}
		log.Printf("输出结果: %s", content)
	},
}

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(wordCmd)
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的格式")
}

// Exeute command
func Exeute() error {
	return rootCmd.Execute()
}
