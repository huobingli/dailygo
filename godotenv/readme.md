## repo 
github.com/joho/godotenv

## 简介
防止一台机器中运行多个项目，环境变量冲突，godotenv库从.env文件中读取配置，然后存储到程序的环境变量中。

## 默认载入.env文件
``` go
// 默认load .env文件
func filenamesOrDefault(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{".env"}
	}
	return filenames
}

```