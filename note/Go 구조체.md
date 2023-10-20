## 구조체

Go 에서 struct 는 Custom Data Type 을 표현하는데 사용되는데 , Go 의 struct는 필드들의 집합체이며
필드들의 컨테이너이다. Go 에서 struct는 필드 데이터만을 가지며, (행위를 표현하는) 메서드를
갖지 않는다. 

Go에는 전통적인 OOP 언어가 가지는 클래스, 객체, 상속 개념이 없다. 전통적인 OOP 클래스
는 Go 언어에서 Custom 타입을 정의하는 struct 으로 표현되는데, 전통적인 OOP 클래스가 
필드와 메서드를 함께 갖는 것과 달리 Go 언어의 struct는 필드만을 가지며, 메서드는 별도로 
분리하여 정의한다. 

Struct 선언

struct 를 정의하기 위해서는 Custom Type 을 정의하는데 사용하는 type 문을 사용한다. 

예를 들어 name 과 age필드를 갖는 person이라는 struct를 정의하기 위해서는 아래와 rkxdms 
type 문을 사용할 수 있다. 

```go

package main

import (
    "fmt"
)

type person struct {
    name string
    age  int
}

func main() {
    p := person{}
    p.name = "Lee"
    p.age = 10
    fmt.Println(p)
}
```

위 예제에서 person struct 를 정의하기 위해서는 type 문을 사용하였다. type 문은 Go 언어에서
Custom Type 을 정의하는데 사용되는데, 위 예제에서는 person 이라는 Custom Type 을 정의하였다.
person struct 는 name 과 age 필드를 갖는데, 이때 필드명은 소문자로 시작하는데, 이는 해당
필드가 private 임을 의미한다.

Struct 객체 생성 
선언된 struct 타입으로부터 객체를 생성하는 방법은 몇 가지 방법들이 있다. 위의 예제처럼 
person{}  을 사용하여 빈 person 객체를 먼저 할당하고, 나중에 그 필드값을 채워넣는 방법이
있다. struct 필드를 엑세스하기 위해서는 .(dot)을 사용한다. 

struct 객체를 생성할 때, 초기값을 함께 할당하는 방법도 있다. 즉, 아래 첫번째 예처럼, 
struct 필드값을 순서적으로 {} 괄호안에 넣을 수 있으며, 두번째 예처럼 순서에 상관없이 
필드명을 지정하고 (named field) 그 값을 넣을 수도 있다. 특히 두번째 예처럼 필드명을 
지정하는 경우, 만약 일부 필드가 생략된 경우 생략된 필드들은 Zero value 를 갖는다.

```go

package main

import (
    "fmt"
)

type person struct {
    name string
    age  int
}

func main() {
var p1 person
p1 = person{"Bob", 20}
p2 := person{name: "Sean", age: 50}
p3 := person{name: "Anna"}

fmt.Println(p1)

fmt.Println(p2)

fmt.Println(p3)
}
```

또 다른 객체 생성 방법으로 Go 내장함수 new()를 쓸 수 있다. new()를 사용하면 모든 필드를
Zero value로 초기화하고 person 객체의 포인터(*person)을 리턴한다. 

객체 포인터인 경우에도 필드 액세스시 .을 사용하는데 이때 포인터는 자동으로 Dereference 된다. 

```go

package main

import (
    "fmt"
)

type person struct {
    name string
    age  int
}

func main() {
    p4 := new(person)
    p4.name = "Lee"
    p4.age = 10
    fmt.Println(p4)
}
```

Go 에서 struct는 기본적으로 mutable 개체로서 필드값이 변화할 경우 해당 개체 메모리에서
직접 변경된다. 하지만, struct 개체를 다른 함수의 파라미터로 넘긴다면, Pass by Value
에 따라 객체를 복사해서 전달하게 된다. 그래서 만약 pass by reference로 struct를 
전달하고자 한다면, strucct의 포인터를 전달해야 한다. 

생성자(constructor)함수

때로 구조체(struct)의 필드가 사용 전에 초기화되어야 하는 경우가 있다. 예를 들어 struct의 
필드가 map 타입인 경우 map을 사전에 미리 초기화해 놓으면, 외부 struct 사용자가 
매번 map을 초기화해야한다는것을 기억할 필요가 없다. 

이러한 목적을 위해 생성자 함수를 사용할 수 있다. 생성자 함수는 
struct를 리턴하는 함수로서 그 함수 본문에서 필요한 필드를 초기화한다. 

아래 예제에서 생성자 함수 newDict()는 dict라는 struct의 map 필드를 초기화한 후 
그 struct 포인터를 리턴하고 있다. 

```go

package main

import (
    "fmt"
)

type dict struct {
    data map[int]string
}

func newDict() *dict {
    d := dict{}
    d.data = map[int]string{}
    return &d
}

func main() {
    dict := newDict()
    dict.data[1] = "A"
    fmt.Println(dict.data[1])
}
```









































