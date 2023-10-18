### Static 파일 핸들러

HTML, 이미지, JavaScript 파일 등의 정적 파일이 요청되었을 때 처리하는 핸들러입니다.

```go

package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func main() {
	http.Handle("/", new(staticHandler))

	http.ListenAndServe(":5000", nil)
}

type staticHandler struct {
	http.Handler
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	localPath := "public" + req.URL.Path
	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	contentType := getContentType(localPath)
	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)

	switch ext {
	case ".html":
		contentType = "text/html"
	case ".js":
		contentType = "text/javascript"
	case ".css":
		contentType = "text/css"
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	default:
		contentType = "text/plain"
	}
	return contentType
}

```

http.FileServer를 사용한 간단한 핸들러

특정 폴더 안에 있는 정적 파일들을 웹 서버에서 클라이언트로 그대로 전달하기 위해 
내장 기능인 http.FileServer()를 사용할 수 있다. http.FileServer()는 해당 
디렉토리 내의 모든 정적 리소스를 1대 1로 매핑하여 그대로 전달하는 일을 하는데, 위의
ServeHTTP() 핸들러처럼 세밀한 제어를 할 수는 없다. 

아래 예제 코드는 위의 예제처럼 wwwroot 폴더 안의 정적 파일들을 클라이언트에 
전달한다. 만약 시작 위치를 /static 으로 변경하려면 , Handle()의 첫번째 파라미터를 
/static 으로 지정하면 된다. 

```go
package main

import (
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("wwwroot")))
	//http.Handle("/static", http.FileServer(http.Dir("wwwroot")))
    http.ListenAndServe(":5000", nil)
}
```









































