## Go - Map

Go 언어에서는 Map이라는 자료형을 제공한다. Map은 Key와 Value로 이루어진 자료형이다.

Go 언어는 Map 타입을 내장하고 있는데, "map[Key타입]Value타입" 과 같이 선언한다.

```go

package main

import (
    "fmt"
)

func main() {
    var idMap map[int]string
    idMap = make(map[int]string)
    idMap[901] = "Apple"
    idMap[134] = "Grape"
    idMap[777] = "Tomato"

    str := idMap[134]
    fmt.Println(str)

    noData := idMap[999]
    fmt.Println(noData)

    delete(idMap, 777)
    fmt.Println(idMap)
}
```

Map은 make() 함수의 첫번째 파라미터로 map 키워드와 키타입 dmf wlwjdgksmsep, 
이때의 make() 함수는 해시테이블 자료구조를 메모리에 생성하고 그 메모리를 가리키는 map value로 
리턴한다. (map value)는 내부적으로 runtime.hmap구조체를 가리키는 포인터이다. 

map은 make() 함수를 써서 초기화할 수도 있지만, 리터럴을 사용해 초기화할 수도 있다. 
    
```go

package main

import (
    "fmt"
)

func main() {
    fruits := map[string]int{
        "apple":  1500,
        "banana": 3000,
        "grape":  2000,
    }
    fmt.Println(fruits["grape"])
}
```

map 사용 
처음 map이 make() 함수에 의해 초기화 되었을 때는 , 아무데이터가 없는 상태이다. 
이때 새로운 데이타를 추가하기 위해서는 "map변수[키] = 값" 과 같이 해당 키에 그 값을 할당하면
된다. 예를 들면, 아래 예제에서 키 901에 apple을 할당하면 새 해시 키-값 쌍이 추가된다. 

만약 901의 값이 이미 존재했다면 추가대신 값만 갱신한다. 

```go

package main

import (
    "fmt"
)

func main() {
    var m (map[int]string)
    //추가 혹은 갱신
    m[901] = "Apple"
    m[134] = "Grape"
    m[777] = "Tomato"

// 키에 대한 값 읽기
    str := m[134]
    fmt.Println(str)

    noData := m[999] // 값이 없으면 nil 혹은 zero 리턴
    fmt.Println(noData)

    // 삭제
    delete(m, 777)
}
```














































