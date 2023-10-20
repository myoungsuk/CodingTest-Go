## for문

Go는 반복문에 for 하나 밖에 없다. 

```go
package main

import (
    "fmt"
)

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
}
```

for문은 위와 같이 사용하며, for문의 조건식은 괄호로 묶지 않는다.

for 문 - 조건식만 쓰는 for 루프
while 루프와 같이 쓰이도록 한다. for 루프는 n이 100을 넘는지를 체크하는 조건식만을 사용하고 있다
```go
package main

import (
    "fmt"
)

func main() {
    i := 0
    for i < 10 {
        fmt.Println(i)
        i++
    }
}
```

for 문 - 무한 루프

```go

package main

import (
    "fmt"
)

func main() {
    for {
        println("Infinite loop")
    }
}
```

for range문

for range 문은 컬렉션으로부터 한 요소 씩 가져와 차례로 for 블럭의 문장들을 실행한다. 
이는 다른 언어의 foreach와 비슷한 용법이다. 

for range 문은 "for 인덱스, 요소값 := range 컬렉션" 같이 for 루프를 구성하는데, 
range 키워드 다음의 컬렉션으로부터 하나씩 요소를 리턴해서 그 요소의 위치인덱스와 값을 
for 키워드 다음의 2갸의 변수에 각각 할당한다. 

```go
package main

import (
    "fmt"
)

func main() {
    a := []int{1, 2, 3, 4, 5}
    for i, v := range a {
        fmt.Println(i, v)
    }
}
```

break, continue goto 문

경우에 따라 for 루프내에서 즉시 빠져나올 필요가 있는데, 이때 break 문을 사용한다. 
만약 for 루프의 중간에서 나머지 문장들을 실행하지 않고 for 루프 시작부분으로 바로 가려면
continue 문을 사용한다. 

그리고 기타 임의의 문장으로 이동하기 위해 goto 문을 사용할 수 있다. 

```go

package main

import (
    "fmt"
)

func main() {
   var a = 1
   for a < 15 {
	   if a == 5 {
		   a += a
		   continue // for루프 시작으로
	   }
	   a++
	   if a > 10 {
		   break //루프 빠져나옴
	   }
   }
   if a == 11 {
       goto END //goto 사용예
   }
   println(a)
   
   END:
       println("End")
   
}
```
break 문은 보통 단독으로 사용되지만, 경우에 따라 "break 레이블" 과 같이 사용하여 
지정된 레이블로 이동할 수도 있다. 
break의 "레이블" 은 보통 현재의 for 루프를 바로위에 적게 되는데, 이러한 "break"레이블
은 현재의 루프를 빠져나와 지정된 레이블로 이동하고, 
break 문의 직속 for 루프 전체의 다음 문장을 실행하게 한다. 

아래 예제는 언뜻 보기에 무한루프를 돌 것 같지만 실제로는 OK를 출력하고 프로그램을 정상 종료한다. 
이는 break L1 문이 for 루프를 빠져나와 L1 레이블로 이동한 후 , break가 있는 현재
for 루프를 건너뛰고 다음 문장인 println() 으로 이동하기 떄문이다. 

```go

package main

import (
    "fmt"
)

func main() {
    i := 0

L1:
    for {
        if i == 0 {
            break L1
        }
    }
    fmt.Println("OK")
}
```






























