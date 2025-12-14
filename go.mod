module snippetbox.yang.net

go 1.25.4

// 告诉Go，在运行 go run, go test 等的时候，应该使用哪个版本的包
// indirect 代表软件包还没有出现在代码库之中的 import 语句之中
require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
)
