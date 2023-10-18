# SMTP 이메일 보내기

이메일을 발송하기 위한 한 방법으로 SMTP 메일을 이용할 수 있다. Go에서 SMTP를 이용해
이메일을 발송하기 위해서는 내장 패키지인 net/smtp 패키지를 사용할 수 있다. smtp의 SendMail
함수는 SMTP 메일을 보내는 간단한 기능을 제공하는데, 이 함수는 다음과 같은 5개의 파라미터를 갖는다.

01. SMTP 서버 주소 포트번호
2. smtp.Auth 인터페이스를 갖는 타입인데, SMTP 서버에 인증하기 위한 정보를 제공한다.
3. 보내는 사람의 이메일 주소
4. 받는 사람의 이메일 주소 수신자는 복수일 수 있으므로 string 슬라이스를 사용한다.
5. 메시지를 적는다. 메시지는 RFC 822 이메일 스타일이어야 하는데, 즉, 메시지 헤더가 먼저 오고,
6. 빈칸 한 줄이 그다음으로, 마지막으로 메시지 Body를 적는다.
7. 메시지 헤더는 From, To, Subject 등을 적는다.
8. 메시지의 모든 라인은 CRLF로 끝나야 한다.

아래 예제는 smtp.SendMail() 함수를 사용하여 이메일을 발송하는 예제이다.

```go
package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	auth := smtp.PlainAuth("", ")
	to := []string{"
		from := "sender@live.com"

		headerSubject := "subject : 테스트\r\n"
		headerBlank := "\r\n"
		body := "This is a test email from Go"
		msg := []byte(headerTo + headerSubject + headerBlank + body)

		err := smtp.SendMail("smtp.live.com:587", auth, from, to, msg)
		if err != nil{
		panic(err)
	}
	}
```

