## Go 에서 Heap 사용하기 

go의 "container/heap" 패키지는 heap.inteface를 구혀하는 모든 Type 에 대해 minheap 기능을 제고
제공한다. MinHeap은 트리로서 ㄱ장 작은 값이 Root에 오게 되며 ,또한 각 노드는 그 노드의 서브트리보다
작은 값을 갖는 구조를 갖는다. 

이러한 MinHeap에서 모든 요소를 루트로부터 차례로 Pop하게 되면, 가장 작은 값붜어 올림차순으로
소트된 값을 얻게 된다. 

heap 패키지는 heap.Interface를 구현하는 모든 타입에 대해 minheap 기능을 제공한다.

Len() int Element 요소의 수
Less(i, j int) bool i번째 요소가 j번째 요소보다 작은지 여부
Swap(i, j int) i번째 요소와 j번째 요소를 바꾼다.
Push(x interface{}) 요소를 추가한다.
Pop() interface{} 요소를 제거하고 리턴한다.

heap.Interface를 구현하는 타입은 다음과 같은 메서드를 구현해야 한다.

```go
package main

import (
    "container/heap"
    "fmt"
)

type IntegerHeap []int

func (iheap IntegerHeap) Len() int {
    return len(iheap)
}

func (iheap IntegerHeap) Less(i, j int) bool {
    return iheap[i] < iheap[j]
}

func (iheap IntegerHeap) Swap(i, j int) {
    iheap[i], iheap[j] = iheap[j], iheap[i]
}

func (iheap *IntegerHeap) Push(heapintf interface{}) {
    *iheap = append(*iheap, heapintf.(int))
}

func (iheap *IntegerHeap) Pop() interface{} {
    var n int
    var x1 int
    var previous IntegerHeap = *iheap
    n = len(previous)
    x1 = previous[n-1]
    *iheap = previous[0 : n-1]
    return x1
}

func main() {
    var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
    heap.Init(intHeap)
    heap.Push(intHeap, 2)
    fmt.Printf("minimum: %d\n", (*intHeap)[0])
    for intHeap.Len() > 0 {
        fmt.Printf("%d \n", heap.Pop(intHeap))
    }
}
```



