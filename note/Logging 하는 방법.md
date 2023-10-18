### 로그 패키지

Go 에서 로그를 사용하기 위하여 Go의 표준 패키지 "log" 를 이용할 수 있다. 가장 단순한
예제로 아래 예와 같이 log 패키지의 Print*() 같은 출력 함수들을 직접 사용하면, 디폴트 출력
Stdout에 날짜/시간 와 함께 해당 로그 내용을 화면에 출력된다. 
만약 굳이 날짜/시간을 없애고자 한다면, log.SetFlags(0) 을 지정한다. 

```go
package main

import (
    "log"
)

func main() {
    log.Println("Hello World")
    log.SetFlags(0)
    log.Println("Hello World")
}

/* 로그출력 내용
2018/01/01 00:00:00 Hello World
 */
```

log 패키지는 여러가지 로거들을 지원하기 위해 "Logger라는 "라는 타입을 제공하고 있다. 

새로운 로거를 만들기 위해 log.New() 함수를 사용하게 된다. log.New()는 3개의 파라미터를
받아들이는데, 첫번째는 io.Writer 인터페이스를 지원하는 타입으로 표준콘솔출력(os.Stdout)
표준에러, 파일포인터 혹은 io.Writer를 지원하는 모든 타겟이 사용될 수 있다. 

두번째 파라미터는 로그의 prefix를 지정하는데, 보통 로그의 prefix는 로그의 종류를 나타내는데
사용된다. 세번째 파라미터는 로그의 flag를 지정하는데, 보통 로그의 flag는 로그의 출력 형식을
지정하는데 사용된다.

아래 예제는 log.New() 함수를 사용하여 새로운 로거를 만들고, 이 로거를 사용하여 로그를 출력하는
예제이다. 

```go
package main

import (
    "log"
    "os"
)

var myLogger *log.Logger

func main() {
    myLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    myLogger.Println("Hello World")
}

func run() {
    myLogger.Println("Hello World")
}

/* 로그출력 내용
INFO: 2018/01/01 00:00:00 10_18.go:14: Hello World
INFO: 2018/01/01 00:00:00 10_18.go:18: Hello World
 */
```

## 파일 로그

보통 많은 경우 로그파일에 로그를 출력하기 원하는데, 파일에 로깅하기 위해서느 먼저 로그파일을
오픈하고 파일 포인터를 log.New()의 첫번째 파라미터에 넣어 주면 된다. 
로그파일을 오픈할 때는 일반적으로 Write Only,Append Only 모드로 오픈한다. 

아래 예제는 로그파일을 오픈하고, 이 파일 포인터를 log.New()의 첫번째 파라미터에 넣어주는 예제이다. 

```go
package main

import (
    "log"
    "os"
)

var myLogger *log.Logger

func main() {
    fpLog, err := os.OpenFile("logfile.txt, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        panic(err)
    }
    defer fpLog.Close()

    myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    myLogger.Println("Hello World")
}

/* 로그출력 내용
INFO: 2018/01/01 00:00:00 10_18.go:14: Hello World
 */
``` 






































