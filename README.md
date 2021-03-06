# workerpool
[![License](https://img.shields.io/badge/license-GPL-blue.svg)](https://github.com/lgphone/workerpool/blob/main/LICENSE)

方便的控制go协程并发数,并收集go协程出现的异常


# 安装  Installation
To install this package, you need to setup your Go workspace. The simplest way to install the library is to run:
```shell 
go get github.com/lgphone/workerpool
```

# 示例 Example
```go
package main

import (
	"errors"
	"fmt"
	"github.com/lgphone/workerpool"
	"strconv"
	"time"
)

func main() {
	// set max worker
	workerPool := workerpool.NewWorkerPool(8)
	tasks := []int{4, 3, 6, 8, 2, 1, 10, 31, 62, 51}
	for _, tasks := range tasks {
		taskId := tasks
		// submit task
		workerPool.Submit(func() {
			fmt.Println("start task: ", taskId)
			time.Sleep(time.Millisecond * 1000)
			if taskId == 3 || taskId == 1 {
				panic(errors.New("task err, id: " + strconv.Itoa(taskId)))
			}
			fmt.Println("finished task: ", taskId)
		})
	}
	fmt.Println("还可以做其他事情")
	// wait for all task finished
	if errs := workerPool.Wait(); len(errs) != 0 {
		fmt.Println("has error: ", errs)
	}
	fmt.Println("all done.")
}

```

