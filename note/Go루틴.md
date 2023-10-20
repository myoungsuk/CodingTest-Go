## Go루틴

goroutine은 Go 런타임이 관리하는 Lightweight 논리적 쓰레드이다. 
Go 에서 go 키워드를 사용하여 함수를 호출하면, 런타임시 새로운 goroutine을 실행한다. 

goroutine은 비동기적으로 함수루틴을 실행하므로, 여러 코드를 동시에 실행하는데 사용된다. 

goroutine은 OS쓰레드보다 훨씬 가볍게 비동기 Concurrent 처리를 구현하기 위해 만든 것으로
기본적으로 Go 런타임이 자체 관리한다. Go 런타임 상에서 관리되는 작업단위인 여러 
goroutine들은 종종 하나의 OS 쓰레드 1개로도 실행되곤 한다. 즉 Go루틴들은 OS쓰레드와 
1대 1로 대응되지 않고, 
Multiplexing으로 훨씬 적은 OS 쓰레드를 사용한다. 메모리 측면에서도 OS 쓰레드가 
1메가바이트의 스택을 갖는 반면, goroutine은 이보다 훨씬 작은 몇 킬로바이트의 스택을 갖는다. 

Go 런타임은 Go루틴을 관리하면서 Go 채널을 통해 Go루틴 간의 통신을 쉽게 할 수 있도록 하였다. 

아래 예제에서 main 함수를 보면, 먼저 say() 라는 함수를 동기적으로 호출하고, 다음으로 동일한
say() 함수를ㄹ 비동기적으로 3번 호출하고 있다. 첫번째 동기적 호출은 say() 함수가 완전히 
끝났을 때 다음 문장으로 이동하고, 다음 3개의 go say() 비동기 호출은 별도의 Go루틴들에서 동작
하면서, 메인루틴은 계속 다음 문장을 실행한다. 여기서 goroutine들은 그 실행순서가 
일정하지 않으므로 프로그램 실행시 마다 다른 출력 결과를 나타낼 수 있다. 

```go

package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 3; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    say("sync")
    go say("async1")
    go say("async2")
    go say("async3")
    time.Sleep(1000 * time.Millisecond)
}
```

익명함수 Go루틴

익명함수를 사용하여 Go루틴을 실행할 수도 있다. 

```go

package main

import (
    "fmt"
    "time"
)

func main() {
    go func() {
        fmt.Println("Hello")
    }()
    time.Sleep(1000 * time.Millisecond)
}
```

익명함수 Go루틴 2

익명함수 Go루틴은 함수의 파라미터로도 사용할 수 있다. 

```go

package main

import (
    "fmt"
    "time"
)

func main() {
    go func(msg string) {
        fmt.Println(msg)
    }("Hello")
    time.Sleep(1000 * time.Millisecond)
}
```


Go는 디폴트로 1개의 cpu를 사용한다. 즉 여러개의 Go 루틴을 만들더라도, 1개의 cpu에서 
작업을 시분할하여 처리한다. 만약 머신이 복수개의 cpu를 가진 경우, go 프로그램을 다중
cpu에서 병렬처리 하게 할 수 있는데, 병렬처리를 위해서는 아래와 같이 
runtime.GOMAXPROCS 함수를 호출하여야 한다. 









































