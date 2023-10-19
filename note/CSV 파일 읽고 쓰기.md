## CSV 파일 읽기

Go에서 CSV 파일을 처리하기 위해 "endoing/csv" 패키지를 사용한다. 

CSV 파일을 읽기 위해선 먼저 csv.NewReader()를 실행하여 csv 파일 포맷을 읽을 수 있는 Reader
를 생성하고, 이 리더로부터 Read() 혹은 ReadAll() 메서드를 호출하여 데이터를 읽어들인다. 

Read() 메서드는 한번에 한줄씩 읽어들이고, ReadAll() 메서드는 한번에 모든 라인을 읽어들인다. 

```go
package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "bufio"
)

func main() {
    file, err := os.Open("test.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    reader := csv.NewReader(bufio.NewReader(file))
    for {
        line, err := reader.Read()
        if err != nil {
            break
        }
        fmt.Println(line[0], line[1], line[2])
    }
}
```

# CSV 파일 쓰기

CSV 파일을 쓰기 위해서는 csv.NewWriter()를 실행하여 csv 파일 포맷을 쓸 수 있는 Writer를 생성하고,
이 Writer로부터 Write() 메서드를 호출하여 데이터를 쓴다. 

```go
package main

import (
    "encoding/csv"
    "os"
    "bufio"
)

func main() {
    file, err := os.Create("test.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    writer := csv.NewWriter(bufio.NewWriter(file))
    writer.Write([]string{"1", "김철수", "서울시 강남구"})
    writer.Write([]string{"2", "이영희", "서울시 강북구"})
    writer.Flush()
}
```


