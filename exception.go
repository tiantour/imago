package imago

import (
	"log"

	"github.com/kataras/iris"
	"github.com/labstack/echo"
)

// OK 真假提示
func (e *iexception) OK(ok bool, message string) {
	if !ok {
		panic(message)
	}
}

// Err 错误提示
func (e *iexception) Err(err error, message string) {
	if err != nil {
		if message == "" {
			message = err.Error()
		}
		panic(message)
	}
}

//  Catch 抛出异常
func (e *iexception) Catch(ctx interface{}) (exception func()) {
	return func() {
		// 捕获异常
		if r := recover(); r != nil {
			log.Println(r) // 1.写入日志
			switch ctx.(type) {
			case echo.Context:
				ctx.(echo.Context).JSON(200, H{"status": "404", "message": r}) // 2.返回错误
			case *iris.Context:
				ctx.(*iris.Context).JSON(200, H{"status": "404", "message": r}) // 2.返回错误
			}
		}
		// 捕获异常，不做处理
		// recover()
	}
}
