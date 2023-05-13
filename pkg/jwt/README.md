# 如何使用JWT

## 1. JWT是什么

**JWT（JSON Web Token）**是一种基于 JSON 的开放标准，用于在网络应用间传递声明。JWT被设计为可安全地将用户身份验证和授权数据作为 JSON 对象在各个应用程序之间传递。

JWT 主要由三部分组成：`Header`，`Payload` 和` Signature`。

Header 包含了两部分信息：令牌的类型（即 JWT）和所使用的算法，通常采用的算法是 HMAC SHA256 或 RSA。

**Payload（载荷）包含了要传递的信息，可以包括用户 ID、用户角色、过期时间等信息。**

Signature（签名）是使用私钥对 Header 和 Payload 进行签名生成的，用来验证消息确实是由发送方发出的，以及在传输过程中没有被篡改过。

JWT 的优点包括：

1. 无状态：JWT 本身就包含了用户信息，**不需要再去查询数据库或者其他的存储设备**。
2. 安全性：由于JWT包含了签名，所以一旦JWT被篡改，接收方就能够检测到。
3. 便捷性：JWT 的格式是轻量级的，容易传输，可以通过 URL、POST 参数或者在HTTP header中发送。

使用 JWT 的过程可以分为以下几个步骤：

1. 在服务器端生成一个 JWT，包括 Header、Payload 和 Signature。
2. 在需要验证用户身份的请求中，将 JWT 添加到 HTTP header 中。
3. 服务器收到请求后，从 HTTP header 中提取 JWT，并对其进行验证，包括签名验证、Payload 中的信息验证等。
4. 如果验证成功，服务器返回请求所需的数据。
5. 如果验证失败，则拒绝请求。

JWT 是一种非常流行的身份验证和授权机制，被广泛应用于各种互联网应用中。

## 2. GO生成JWT

```go
package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"myWeb/config"
	"time"
)

type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成 token

func GenerateToken(id int64 ,username string) (string, error) {
	// 定义 token 的过期时间
	expireTime := time.Now().Add(config.ExpireDuration).Unix()

	// 创建一个自定义的 Claims
	claims := &Claims{
        ID:id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "pogf",
		},
	}

	// 使用 JWT 签名算法生成 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 将 token 进行加盐加密
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

```

详细分析:

```go
type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
```

该结构体用于定义 JWT 的 payload，也就是 JWT 载荷中存储的信息。其中，`ID` 和 `Username` 用于标识用户身份，`jwt.StandardClaims` 则是 JWT 的标准声明，包含了 JWT 的一些基本信息，比如过期时间、签发时间等。在创建 JWT 时，我们需要将这个结构体传入，以便生成 JWT 的 payload。在验证 JWT 时，我们也需要解析出这个结构体，以便获取 JWT 中存储的用户身份信息。

```go
expireTime := time.Now().Add(config.ExpireDuration).Unix()
```

这行代码的作用是根据配置文件中的过期时间（`config.ExpireDuration`）计算出当前时间加上过期时间后的 Unix 时间戳（`expireTime`），以便在生成 JWT 时设置过期时间。

`time.Now()` 函数返回当前的本地时间（time.Time 类型），`Add()` 方法则用于对时间进行加减，参数是 `time.Duration` 类型，表示时间间隔。因此，`time.Now().Add(config.ExpireDuration)` 计算出的是过期时间的时间点。接着，调用 `Unix()` 方法将时间点转换为对应的 Unix 时间戳（int64 类型），用于设置 JWT 的过期时间。

```go
	claims := &Claims{
        ID:id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "pogf",
		},
	}
```

这段代码是在创建一个包含用户身份信息的JWT token，其中`Claims`结构体定义了JWT token中所包含的信息，包括`ID`和`Username`，以及`jwt.StandardClaims`结构体中的标准信息，如`ExpiresAt`（过期时间），`IssuedAt`（签发时间）和`Issuer`（签发者）等。

在上述代码中，`ID`和`Username`分别代表用户的唯一标识和用户名，这些信息将被编码到JWT token中。`StandardClaims`中的信息则用于控制JWT token的生命周期和有效性。

`ExpiresAt`表示JWT token的过期时间，超过这个时间后，JWT token将失效，无法再被使用。`IssuedAt`表示JWT token的签发时间，`Issuer`表示JWT token的签发者，这些信息可以帮助验证JWT token的合法性。

```go
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
```

它首先创建了一个 JWT 对象 `token`，其中包含了指定的签名方法 `jwt.SigningMethodHS256` 和声明 `claims`。

JWT 中有三个部分：头部（Header）、载荷（Payload）和签名（Signature）。头部指定了所使用的签名算法，例如 `HS256` 表示使用 HMAC-SHA256 算法进行签名。载荷是 JWT 中用来携带实际信息的部分，其中可以自定义各种声明信息，例如上面代码中的 `claims` 变量。签名是将头部和载荷进行加密得到的一串字符串，用于验证 JWT 是否合法。

因此，这段代码就是创建了一个带有指定声明的 JWT 对象，并指定了使用 HS256 算法进行签名。

```go
	// 将 token 进行加盐加密
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
```

加盐加密的目的是为了增加数据的安全性，防止攻击者使用彩虹表等方式对加密后的数据进行破解。加盐是指在原始数据的基础上加上一些随机的字符串或数字，使得同样的原始数据加盐后的结果是不同的。加密是指将原始数据和盐一起通过某种加密算法进行加密，得到一串密文。密文是不能被破解的，只能通过使用同样的盐和加密算法对原始数据进行加密后得到相同的密文来验证数据的真实性。

在JWT中，加盐加密可以有效地防止攻击者伪造token，从而保障应用的安全性。具体地，JWT在生成token时会使用指定的秘钥对payload进行签名，这个秘钥就是加盐的一部分。只有使用相同的秘钥才能对payload进行解密并验证token的真实性。因此，只有知道秘钥的人才能生成有效的token，从而有效地保护了应用的安全性。

### 单元测试

```go
func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(1, "wyf")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(token)
}

```

```shell
$ go test -v -run  TestGenerateToken
=== RUN   TestGenerateToken
    jwt_test.go:10: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ3eWYiLCJleHAiOjE2ODM4NzU4NjcsImlhdCI6MTY4Mzc4OTQ2NywiaXNzIjoicG9nZiJ9.rXzf1fzidsfwe4HRBt7JN_NxAxUceD0HxdpCQsbPuFc
--- PASS: TestGenerateToken (0.00s)
PASS
ok      myWeb/util      0.058s

```

## 3. GO解析JWT

### 相关函数介绍

```go
func ParseToken(tokenString string) (*Claims, error) {
	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorClaimsInvalid)
	}
}

```

```go
func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {
	return new(Parser).ParseWithClaims(tokenString, claims, keyFunc)
}
```

`jwt.ParseWithClaims()`函数用于解析并验证JWT token，并提取出其中的Claims数据结构。该函数接受三个参数：

- `tokenString`: 要解析的JWT token字符串。
- `claims`: 一个结构体指针，用于接收从JWT token中解析出来的Claims数据。
- `keyFunc`: 一个回调函数，用于验证JWT token的签名。在函数内部，可以对JWT token的签名进行验证，比如检查签名是否正确、是否过期等等。该函数需要返回一个interface{}类型的值，表示用于验证签名的密钥。如果验证成功，应该返回一个非空的密钥；如果验证失败，应该返回一个nil值和相应的错误信息。

在上述代码中，`&Claims{}`作为第二个参数传递给了`jwt.ParseWithClaims()`函数，这表示从JWT token中解析出的Claims数据将被解码为该结构体类型。第三个参数是一个匿名函数，用于验证JWT token的签名，`config.JwtKey`被作为用于验证签名的密钥传递。如果JWT token验证通过，`jwt.ParseWithClaims()`函数将返回一个`*jwt.Token`类型的指针，其中包含了解码后的Claims数据，以及一些其他的元数据。如果验证失败，函数将返回相应的错误信息。

### 单元测试

```go
func TestParseToken(t *testing.T) {
	c, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ3eWYiLCJleHAiOjE2ODM4NzU4NjcsImlhdCI6MTY4Mzc4OTQ2NywiaXNzIjoicG9nZiJ9.rXzf1fzidsfwe4HRBt7JN_NxAxUceD0HxdpCQsbPuFc")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}
```

```shell
$ go test -v -run  TestParseToken
=== RUN   TestParseToken
    jwt_test.go:18: {1 wyf { 1683875867  1683789467 pogf 0 }}
--- PASS: TestParseToken (0.00s)
PASS
ok      myWeb/util      0.796s
```

## 4. 总结


JWT（JSON Web Token）是一种用于身份验证的开放标准，它通过在用户和服务器之间传递被加密的 JSON 对象来安全地传输信息。JWT 由三个部分组成：头部、载荷和签名。其中，头部包含加密算法和 token 类型等信息，载荷包含存储在 token 中的用户信息，签名用于验证 token 的真实性和完整性。

使用 JWT 时需要注意以下几点：

1. 避免在 JWT 中存储敏感信息，比如密码、银行卡号等。
2. 需要使用安全的算法来签名和加密 JWT，比如使用 HMAC 或 RSA 加密。
3. JWT 中包含的用户信息不应该过多，只保留必要的信息即可。
4. 在使用 JWT 进行身份验证时，需要防止令牌被盗用。可以采用一些技术手段，比如限制令牌的有效期、限制令牌的使用次数等。
5. 在生成 JWT 时需要注意加盐加密，以提高安全性。
6. 在解析 JWT 时需要对 JWT 进行校验，以保证 JWT 的真实性和完整性。
7. 最后，需要注意遵循安全的开发实践，保障应用的安全性。