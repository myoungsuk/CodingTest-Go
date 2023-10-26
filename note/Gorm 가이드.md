### Gorm 가이드

- 완전한 기능을 가진 ORM
- Associations
- Hooks
- Preload, Joins 을 통한 데이터 가져오기 
- Transactions, Nestied Transactions, Save Point, RollbackTo to Saved Point
- Context, Prepared Statement 모드 , DryRun 모드 
- Batch insert, FindBatches, Find/Create with Map, CRUD with SQL Expr and Context VAlueer
- SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, Named Argument, 
- SubQuery
- Composite Primary Key, Indexes, Constratins
- Auto Migration, Custom Table Name, Custom Column Name, Custom Table Option
- Logger

확장 가능하고 유연한 플러그인 API

빠르게 시작하기 
    
```go

package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Product struct {
    gorm.Model
    Code string
    Price uint
}

func main() {
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
if err != nil {
    panic("failed to connect database")
}

// 테이블 자동 생성
db.AutoMigrate(&Product{})

//생성
db.Create(&Product{Code: "D42", Price :100})

//읽기
var product Product
db.First(&product, 1)
db.First(&product, "code = ?", "D42")

//수정 - product의 price를 200으로 
db.Model (&product).Update("Price", 200)

//수정 - 여러개의 필드를 수정하기
db.Model(&product).Updates(Product{Price : 200, Code: "F42"}) // price만 업데이트
db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

//삭제
db.Delete(&product, 1)
}
```

모델 선언 
모델은 기본 Go 유형, 포인터/별칭 또는 Scanner 및 Valuer 인터페이스를 구현하는 사용자 정의 유형이 
있는 일반 구조체이다. 

예를 들면
type User struct {
ID uint
Name string
Email *string
Age uint8
Birthday *time.Time
MemberNumber sql.NullString
ActivatedAt sql.NullTime
CreatedAt time.Time
UpdatedAt time.Time
}

Gorm은 구성보다 규칙을 선호한다. 기본적으로 Gorm은 ID를 기본 키로 사용하고, 
복수형 snake_cases를 구조체 이름을 테이블 이름으로, snake_case를 열 이름으로, 
새성/업데이트 시간을 추적하기 위해 CreatedAt, UpdatedAt을 사용한다. =

```go

type Model struct {
Id uint `gorm:"primaryKey"`
CreatedAt time.Time
UpdatedAt time.Time
DeletedAt gorm.DeletedAt `gorm:"index"`
}

```

고급

필드 수준 권한

내보낸 필드는 Gorm으로 CRUD를 수행할 때 모든 권한을 가지며, Gorm에서는 태그를 사용하여
필드 수준 권한을 읽기 전용, 쓰기 전용, 만들기 전용, 업데이트 전용 또는 무시로 
설정할 수 있다. 

Gorm 마이그레이터를 사용하여 테이블을 만들 때 무시된 필드는 생성되지 않는다. 

```go
type User strut {
Name string `gorm:"<-:create"` // 읽기/생성 허용
Name string `gorm:"<-:update"` // 읽기/업데이트 허용


Embedded Struct 

type User struct {
gorm.Model
Name string
}

type User struct {
Id uint `gorm:"primaryKey"`
CreatedAt time.Time
UpdatedAt time.Time
DeletedAt gorm.DeletedAt `gorm:"index"`
Name string
}

```

for a nomal struct field, you can embed it with the tag embedded, for example

```go
type Author struct {
Name string
Email string
}

type Blog struct {
ID int
Author Author `gorm:"embedded"`
Upvoteds int32
}

type Blog struct {
ID int 64
Name string
Email string
Upvoteds int32
}
```



















































