## Go 인터페이스

구조체가 필드들의 집합체라면, inteface는 메서드들의 집합체이다. interface는 타입이 구현해야
하는 메서드 원형들을 정의한다. 하나의 사용자 정의 타입이 inteface를 구현하기 위해서는
단순히 그 인터페이스가 갖는 모든 메서드들을 구현하면 된다. 

```go

package main

import (
    "fmt"
)

type Duck interface {
    Quack()
    Feathers()
}

type Person struct {
    name string
}

func (p Person) Quack() {
    fmt.Println(p.name + " Quack")
}

func (p Person) Feathers() {
    fmt.Println(p.name + " has feathers")
}

func main() {
    var donald Duck
    donald = Person{"Donald"}
    donald.Quack()
    donald.Feathers()
}
```

위 예제에서 Duck 이라는 인터페이스를 정의하고, Person 이라는 구조체가 Duck 인터페이스를
구현하도록 하였다. Duck 인터페이스는 Quack() 과 Feathers() 라는 두 개의 메서드를 갖는데, 
Person 구조체는 이 두 메서드를 구현하였다.

인터페이스 사용

인터페이스를 사용하는 일반적인 예로 함수가 파라미터로 인터페이스를 받아들이는 경우를 들 수 있다. 
함수 파라미터가 interface인 경우, 이는 어떤 타입이든 해당 인터페이스를 구현하기만 하면 모두
입력 파라미터로 사용될 수 있다는 것을 의미한다.

아래 예제에서 showArea() 함수는 shape 인터페이스들을 파라미터로 받아들이고 있는데, 
따라서 Circle 처럼 Shape 인터페이스를 구현한 타입 객체들을 파라미터로 받을 수 있다. 
showArea() 함수 내에서 해당 인터페이스가 가진 메서드 즉 area() 혹윽 perimeter()를
사용할 수 있다. 

```go

package main

import (
    "fmt"
    "math"
)

type Shape interface {
    area() float64
    perimeter() float64
}

type Circle struct {
    x, y, r float64
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
    return 2 * math.Pi * c.r
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

func (r *Rectangle) perimeter() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return 2 * (l + w)
}

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

func main() {
    c := Circle{0, 0, 5}
    r := Rectangle{0, 0, 10, 10}
    fmt.Println(totalArea(&c, &r))
}
```

인터페이스 타입

Go 프로그래밍을 하다보면 흔히 빈 인터페이스를 자주 접하게 되는데, 흔히 인터페이스 타입
으로도 불리운다. 예를 들어 여러 표준 패키지들의 함수 Prototype을 살펴보면, 
아래와 같이 빈 interface가 자주 등장함을 볼 수 있다. 

빈 interface는 interface{}   와 같이 표현된다. 

Empty interface는 메서드를 전혀 갖지 않는 빈 인터페이스로서, Go의 모든 Type은 
적어도 0개의 메서드를 구현하므로 , 흔히 Go에서 모든 Type을 나타내기 위해 빈 인터페이스를
사용한다. 

즉 빈 인터페이스는 어떠한 타입도 담을 수 있는 컨테이너라고 할 수 있다. 

```go

package main

import (
    "fmt"
)

func main() {
    var x interface{}
    x = 1
    x = "Tom"
    x = false
    fmt.Println(x)
}
```

Type Assertion

interface type의 x와 타입 T에 대하여 x.(T)로 표현했을 때, 이는 x가 아니라 nil이 아니며, x는
T 타입에 속한다는 점을 확인 하는 것으로 이러한 표현을 Type Assertion이라 부른다

만약 x가 nil이거나 x의 타입이 T가 아니라면, 런타임 에러가 발생할 것이고, x 가 T 타입인 경우는
T 타입의 x를 리턴한다. 즉, 아래 예제에서 변수 j는 a.(int)로부터 int형 변수 j가 된다. 

```go

package main

import (
    "fmt"
)

func main() {
    var a interface{} = 1
    i := a       // a와 i는 dynamic type, 값은 1
    j := a.(int) // j는 int 타입, 값은 1
    fmt.Println(i)
    fmt.Println(j)
}
```












































