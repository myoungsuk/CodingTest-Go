### http.Handle() 사용

HandleFunc() 메서드와 비슷하게, HTTP Handler를 정의하는 또 다른 방식으로 http.Handle()
메서드를 사용할 수 있다. http.Handle() 메서드는 첫번째 파라미터로 URL을 받아들이고, 
두번째 파라미터로 http.Handler 인터페이스를 갖는 객체를 받아들인다. 

http.Handler 인터페이스 다응과 같은 메서드 하나를 같은 이터페이스이다. ServeHTTP() 메서드는
HTTP Response에 데이타를 쓰기 위한 Writer와 HTTP Request 입력 데이터를 파라미터로 갖는다. 

```go

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

```

아래 예제는 http.Handler 인터페이스를 갖는 testHandler 라는 struct를 정의하고
이 struct의 메서드 ServeHTTP()를 구현하였다.
    
```go
    package main

    import (
        "fmt"
        "net/http"
    )

    func main() {
        http.Handle("/hello", testHandler{})
        http.ListenAndServe(":5000", nil)
    }

    type testHandler struct{
http.Handler
}

func(h *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    str:= "your request path is " + req.URL.path
    w.Write([]byte(str))
}


```



































