## 클로저 

Go 언어에서 함수는 Closure로서 사용될 수도 있다. 
Closure는 함수 바깥에 있는 변수를 참조하는 함수값을 일컫는데, 이때의 함수는 
바깥의 함수를 마치 함수 안으로 끌어들인 듯이 그 변수를 읽거나 쓸 수 있게 된다. 

Go 언어에서 함수는 일급함수로서 다른 함수로부터 리턴되는 리턴값으로 사용될 수 있다. 

그런데 여기서 이 익명 함수가 그 바깥에 있는 변수 i 를 참조하고 있다. 
익명함수 자체가 로컬 변수로 i 를 갖는 것이 아니기 때문에 외부 변수 i 가 상태를 계속 
유지하는 즉 값을 계속 하나식 증가시키는 기능을 하게 된다. 

예제에서 next := nextValue()   에서 Closure 함수를 next라는 변수에 할당한 후에, 계속 
next()  를 3번 호출하는데 이때마다 Closure 함수 내의 변수 i는 계속 증가된 값을 
가지고 있게 된다. 

```go

package main

import (
    "fmt"
)

func nextValue() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    next := nextValue()
    fmt.Println(next())
    fmt.Println(next())
    fmt.Println(next())
    anotherNext := nextValue()
    fmt.Println(anotherNext())
    fmt.Println(anotherNext())
}
```

