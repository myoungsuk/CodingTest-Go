## Go 채널

Go 채널은 그 채널을 통하여 데이터를 주고 받는 통로라 볼 수 있는데, 채널은 make() 함수를 통해
미리 생성되어야 하며, 채널 연산자 <-  를 통해 데이터를 보내고 받는다. 

채널은 흔히 goroutine들 사이 데이타를 주고 받는데 사용되는데, 상대편이 준비될 때까지 채널에서
대기함으로써 별도의 lock을 걸지않고 데이터를 동기화하는데 사용된다. 

아래 예제는 정수형 채널을 생성하고, 한 goroutine에서 그 채널에 123이란 정수데이타ㅡㄹ
보내는 후 이를 다시 메인 루틴에서 채널로부터 123 데이터를 받는 코드이다. 

채널을 생성할 때는 make() 함수에 어떤 타입의 데이터를 채널에서 주고 받을지를 
미리 지정해주어야 한다. 채널로 데이터를 보낼 때는 채널명 <-
데이터와 같이 사용하고, 채널로부터 데이터를 받을 경우는 <-  채널명과 같이 사용한다. 

아래 예제에서 메인 루틴은 마지막에서 채널로부터 데이터를 받고 있는데, 상대편 goroutine 에서
데이터를 전송할때까지 계속 대기하게 된다.

```go

package main

import (
    "fmt"
)

func main() {
    ch := make(chan int)
	
    go func() {
        ch <- 123
    }()
	
    var i int
    i = <-ch
    fmt.Println(i)
}
```

Go 채널은 수신자와 송신자가 서로를 기다리는 속성 때문에, 이를 이용하여 Go루틴이 끝날 때까지
기다리는 기능을 구현할 수 있다. 즉, 익명함수를 사용한 한 Go 루틴에서 어떤 작업이 
실행되고 있을 때, 메인루틴은 <- done 에서 계속 수신하며 대기하고 있게 된다. 

익명함수 Go 루틴에서 작업이 끝난 후, done 채널에 true를 보내면, 수신자 메인루틴은 이를 받고
프로그램을 끝내게 된다. 

```go

package main

import (
    "fmt"
)

func main() {
    done := make(chan bool)
    
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
        done <- true
    }()
    
    <-done
}
```


Go 채널 버퍼링

Go 채널은 2가지의 채널이 있는데, Unbuffered channel 과 Buffered channel 이 있다.
위의 예제에서의 Go 채널은 Unbuffered Channel로서 이 채널에서는 하나의 수신자가 
데이터를 받을 때까지 송신자가 데이터를 보내는 채널에 묶여 있게 된다. 하지만 Buffered channel
을 사용하면 비록 수신자가 받을 준비가 되어 있지 않을 지라도 지정된 버퍼만큼 데이터를 보내고
계속 다른 일을 수행할 수 있다. 

버퍼 채널은 make(chan type, N)함수를 통해 생성되는데, 두번째 파라미터 N에 사용할 버퍼
갯수를 넣는다. 
예를 들어, make(chan int, 10) 은 10개의 정수형을 갖는 버퍼 채널을 만든다.

버퍼 채널을 이용하지 않는 경우, 아래와 같은 코드는 에러를 발생시킨다. 
왜냐하면 메인루틴에서 채널에 1을 보내면서 상대편 수신자를 기다리고 있는데 이 채널을 받는
수신자 Go루틴이 없기 때문이다. 

```go

package main

import (
    "fmt"
)

func main() {
    ch := make(chan int)
    ch <- 1
    fmt.Println(<-ch)
}
```
하지만 아래와 같이 버퍼채널을 사용하면 수신자가 당장 없더라도 최대버퍼 수까지 데이터를 보낼 수 
잇으므로 에러가 발생하지 않는다. 

```go

package main

import (
    "fmt"
)

func main() {
    ch := make(chan int, 1)
    ch <- 1
    fmt.Println(<-ch)
}
```

채널 파라미터 

채널을 함수의 파라미터로 전달할 때, 일반적으로 송수신을 모두 하는 채널을 전달하지만, 특별이 
해당 채널로 송신만 할 것인지 혹은 수신만 할 것인지를 지정할 수도 있다. 송신 파라미터는
(p chan <- int) 와 같이 chan <- 을 사용하고, 수신 파라미터는 (p <- chan int) 와 같이
만약 송신 채널 파라미터에서 수신을 한다거나, 수신 채널에 송신을 하게되면, 에러가 발생한다. 
아래 예제에서 만약 sendChan() 함수 안에서 x := <- ch 를 실행하면 송신전용 채널에
수신을 시도하므로 에러가 발생한다. 

```go

package main

import (
    "fmt"
)

for main() {
    ch : make(chan string, 1)
       sendChan(ch)
         recvChan(ch)
}

func sendChan(ch chan <- string) {
    ch <- "Data"
}

func recvChan(ch <- chan string) {
    data := <- ch
    fmt.Println(data)
}
```

채널을 리턴값으로 사용할 때도 마찬가지로 송신 전용 채널이나 수신 전용 채널을 리턴할 수 있다.

채널 닫기 

채널을 오픈한 후 데이타를 송신한 후 , close() 함수를 사용하여 채널을 닫을 수 있다. 
채널을 닫게 되면 해당 채널로는 더이상 송신을 할 수없지만, 채널이 닫힌 이후에도 수신은 가능하다. 

채널 수신에 사용되는 <- ch 는 두개의 리턴값을 갖게 되는데 첫째는 채널메세지이고 
두번째는 수신이 제대로 되었는가를 나타낸다. 

만약 채널이 닫혔다면, 두번째 리턴값은 false가 되고, 채널이 열려있다면 true가 된다.
    
```go

package main

import (
    "fmt"
)

func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    close(ch)
    
    for i := 0; i < 3; i++ {
        if i == 2 {
            close(ch)
        }
        v, ok := <-ch
        fmt.Println(v, ok)
    }
}
```





























