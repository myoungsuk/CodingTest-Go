## Go 에러

Go는 내장 타입으로 error 라는 interface 타입을 갖는다. Go 에러는 이 error 인터페이스를
통해서 주고 받게 되는데, 이 interface는 다음과 같은 하나의 메서드를 갖는다. 
개발자는 이 인터페이스를구현하는 커스텀 에러 타입을 만들 수 있다. 

```go

type error interface {
    Error() string
}
```

Go 에러처리

Go 함수가 결과와 에러를 함께 리턴한다면, 이 에러가 nil 인지를 체크해서 에러가 없는지를
체크할 수 있다. 

```go

package main

import (
    "fmt"
)

func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("can't divide by 0")
    }
    return a / b, nil
}

func main() {
    c, err := divide(9, 3)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(c)
}
```

Go 에러처리 2

Go 에러처리 1에서는 에러가 발생하면 에러를 출력하고 프로그램을 종료했지만, 에러가 발생하면
에러를 출력하고 계속 프로그램을 실행하고 싶다면 다음과 같이 하면 된다. 

```go  

package main

import (
    "fmt"
)

func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("can't divide by 0")
    }
    return a / b, nil
}

func main() {
    c, err := divide(9, 0)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(c)
    }
}
```
