## 배열

배열의 선언은 다음과 같이 할 수 있다.

```go

package main

import (
    "fmt"
)

func main() {
    var a [3]int
    a[0] = 1
    a[1] = 2
    a[2] = 3
    fmt.Println(a[1])
}
```

배열의 선언과 동시에 초기화를 할 수도 있다.

```go

package main

import (
    "fmt"
)

func main() {
    a := [3]int{1, 2, 3}
    fmt.Println(a[1])
	
	var a3 = [...]int{1, 2, 3} //배열크기 자동으로
}

```

다차원 배열

```go

package main

import (
    "fmt"
)

func main() {
    a := [3][2]int{
        {1, 2},
        {3, 4},
        {5, 6},
    }
    fmt.Println(a[2][1])
}
```

