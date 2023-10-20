## Go 메서드

Go 구조체처럼 Go 언어는 객체지향 프로그래밍을 고유의 방식으로 지원한다. 타 언어의 OOP의 클래스가
필드와 메서드를 함께 갖는것과 달리 Go 언어에서는 struct가 필드만을 가지며, 메서드는 별도로 
분리되어 정의된다. 

Go 메서든느 특별한 형태의 func 함수이다. 메서드는 함수 정의에서 func 키워드와 함수명 사이에 
그 함수가 어떤 struct 를 위한 메서드인지를 표시하게 된다. 
흔히 reeiver로 불리우는 이 부분은 메서드가 속한 struct 타입과 struct 변수명을 지정하는데

struct 변수명은 함수 내에서 마치 입력 파라미터처럼 사용된다. 예를 들어, 아래 예제는 
rect 라는 struct를 정의하고 있고 area() 라는 메서드를 정의하고 있다. 

func와 area() 사이에 rect 타입의 r이 정의되고 이를 함수 본문에서 사용하고 있다. 
메서드가 선언된 이후에는 
rect 구조체의 객체는 rect.area() 문장처럼 area() 메소드를 struct 객체로부터 직접 호출할 수 있다. 

```go

package main

import (
    "fmt"
)

type rect struct {
    width, height int
}

func (r *rect) area() int {
    return r.width * r.height
}

func main() {
    r := rect{10, 20}
	area := rect.area()
    fmt.Println(r.area())
}
```

Value vs 포인터 receiver

Value receiver는 struct의 데이터를 복자(copy) 하여 전달하며, 포인터 receiver는 
struct의 포인터만을 전달한다. Value receiver의 경우 만약 메서드 내에서 그 struct의 필드값이
변경되더라도 호출자의 데이터는 변경되지 않는 반면, 포인터 receiver는 메서드 내의 필드값
변경이 그대로 호출자에서 반영된다. 

위의 Rect.area() 메서드는 value reciever를 표현한 것으로 만약 area()  메서드 내에서
width나 height 가 변경되더라도 main() 함수의 rect 구조체의 필드값에는 변화가 
없다ㅏ. 하지만 아레와 같이 이를 포인터 receiver로 변경한다면, 메서드 내.r.width++  
필드 변경분이 main() 함수에도 반영되기 때문에 출력값이 11, 220이 된다. 

```go

package main

import (
    "fmt"
)

type rect struct {
    width, height int
}

func (r *rect) area2() int {
    r.width++
    return r.width * r.height
}

func main() {
    r := rect{10, 20}
    area := rect.area2()
    println(rect.width, area)
}
```













































