# Where 
> 一个用于国内地区检索的库

## Install 

```shell
go get go get -u github.com/jiajun-c/where@v0.0.2
```

## Usage

输入地区的名字，可以是省，市，区的名字，返回一个地区信息数组

```go
package main

import (
	"fmt"

	"github.com/jiajun-c/where"
)

func main() {
	ans, _ := where.Where("武汉")
	fmt.Println(ans)
}
// output
// [{湖北省 武汉市  City}]
```

## TODO
- [ ] 智能识别，如湖北武汉医院，解析出湖北，武汉，City
- [ ] 输出所请求的省的全部市，市的全部区

