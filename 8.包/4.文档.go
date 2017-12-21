Golang 文档 ：扩展工具 godoc 能自动提取注释生成帮助文档。


• 仅和成员相邻 (中间没有空行) 的注释被当做帮助信息。 
• 相邻行会合并成同一段落，用空行分隔段落。
• 缩进表示格式化文本，比如示例代码。
• 自动转换 URL 为链接。
• 自动合并多个源码文件中的 package 文档。 
• 无法显式 package main 中的成员文档。


Package

• 建议用专门的 doc.go 保存 package 帮助信息。
• 包文档第一整句 (中英文句号结束) 被当做 packages 列表说明。 


Example

只要 Example 测试函数名称符合以下规范即可。
              格式                              示例
--------+--------------------------------+---------------------------------
package  Example,    Example_suffix          Example_test
func     ExampleF,   ExampleF_suffix         ExampleHello
type     ExampleT,   ExampleT_suffix         ExampleUser, ExampleUser_copy
method   ExampleT_M, ExampleT_M_suffix       ExampleUser_ToString

说明:使用 suffix 作为示例名称，其首字母必须小写。如果文件中仅有一个 Example 函数，且调用了该文件中的其他成员，那么示例会显示整个文件内容，而不仅仅是测试函数自己。


Bug

非测试源码文件中以 BUG(author) 开始的注释，会在帮助文档 Bugs 节点中显示。

// BUG(yuhen): memory leak.