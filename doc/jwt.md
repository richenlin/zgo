# JWT

## 说明
[7519规范](https://tools.ietf.org/html/rfc7519)  
[JWT网站](https://jwt.io/)  

### JWT
Json web token (JWT), 是为了在网络应用环境间传递声明而执行的一种基于JSON的开放标准（(RFC 7519).该token被设计为紧凑且安全的，特别适用于分布式站点的单点登录（SSO）场景。JWT的声明一般被用来在身份提供者和服务提供者间传递被认证的用户身份信息，以便于从资源服务器获取资源，也可以增加一些额外的其它业务逻辑所必须的声明信息，该token也可直接被用于认证，也可被加密。

### 传统的session认证
http协议本身是一种无状态的协议，而这就意味着如果用户向我们的应用提供了用户名和密码来进行用户认证，那么下一次请求时，用户还要再一次进行用户认证才行，因为根据http协议，我们并不能知道是哪个用户发出的请求，所以为了让我们的应用能识别是哪个用户发出的请求，我们只能在服务器存储一份用户登录的信息，这份登录信息会在响应时传递给浏览器，告诉其保存为cookie,以便下次请求时发送给我们的应用，这样我们的应用就能识别请求来自哪个用户了,这就是传统的基于session认证。  

但是这种基于session的认证使应用本身很难得到扩展，随着不同客户端用户的增加，独立的服务器已无法承载更多的用户，而这时候基于session认证应用的问题就会暴露出来.  

### 基于session认证所显露的问题
Session: 每个用户经过我们的应用认证之后，我们的应用都要在服务端做一次记录，以方便用户下次请求的鉴别，通常而言session都是保存在内存中，而随着认证用户的增多，服务端的开销会明显增大。  

扩展性: 用户认证之后，服务端做认证记录，如果认证的记录被保存在内存中的话，这意味着用户下次请求还必须要请求在这台服务器上,这样才能拿到授权的资源，这样在分布式的应用上，相应的限制了负载均衡器的能力。这也意味着限制了应用的扩展能力。  

CSRF: 因为是基于cookie来进行用户识别的, cookie如果被截获，用户就会很容易受  

### 基于token的鉴权机制
基于token的鉴权机制类似于http协议也是无状态的，它不需要在服务端去保留用户的认证信息或者会话信息。这就意味着基于token认证机制的应用不需要去考虑用户在哪一台服务器登录了，这就为应用的扩展提供了便利。

流程上是这样的：

用户使用用户名密码来请求服务器
服务器进行验证用户的信息
服务器通过验证发送给用户一个token
客户端存储token，并在每次请求时附送上这个token值
服务端验证token值，并返回数据
这个token必须要在每次请求时传递给服务端，它应该保存在请求头里， 另外，服务端要支持CORS(跨来源资源共享)策略，一般我们在服务端这么做就可以了Access-Control-Allow-Origin: *。

## 内容

### JWT
JWT是由三段信息构成的，将这三段信息文本用.链接一起就构成了Jwt字符串。就像这样:
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```
解码
```json
{
  "alg": "HS256",
  "typ": "JWT"
},
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022
}
```
```
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  
your-256-bit-secret

)
```

### JWT的构成

第一部分我们称它为头部（header),第二部分我们称其为载荷（payload, 类似于飞机上承载的物品)，第三部分是签证（signature).
#### header
jwt的头部承载两部分信息：

声明类型，这里是jwt
声明加密的算法 通常直接使用 HMAC SHA256
完整的头部就像下面这样的JSON：
```json
{
  "alg": "HS256",
  "typ": "JWT",
}
```
然后将头部进行base64加密（该加密是可以对称解密的),构成了第一部分.
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
```
#### playload
载荷就是存放有效信息的地方。这些有效信息包含三个部分

标准中注册的声明
公共的声明
私有的声明
标准中注册的声明 (建议但不强制使用) ：

iss: jwt签发者
sub: jwt所面向的用户
aud: 接收jwt的一方
exp: jwt的过期时间，这个过期时间必须要大于签发时间
nbf: 定义在什么时间之前，该jwt都是不可用的.
iat: jwt的签发时间
jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。

公共的声明 ：
公共的声明可以添加任何的信息，一般添加用户的相关信息或其他业务需要的必要信息.但不建议添加敏感信息，因为该部分在客户端可解密.

私有的声明 ：
私有声明是提供者和消费者所共同定义的声明，一般不建议存放敏感信息，因为base64是对称解密的，意味着该部分信息可以归类为明文信息。

定义一个payload:
```json
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022
}
```

然后将其进行base64加密，得到Jwt的第二部分。
```
eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ
```

#### signature
jwt的第三部分是一个签证信息，这个签证信息由三部分组成：

header (base64后的)
payload (base64后的)
secret
这个部分需要base64加密后的header和base64加密后的payload使用.连接组成的字符串，然后通过header中声明的加密方式进行加盐secret组合加密，然后就构成了jwt的第三部分。

``` javascript
var encodedString = base64UrlEncode(header) + '.' + base64UrlEncode(payload);

var signature = HMACSHA256(encodedString, 'secret'); // SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

将这三部分用.连接成一个完整的字符串,构成了最终的jwt:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

注意：secret是保存在服务器端的，jwt的签发生成也是在服务器端的，secret就是用来进行jwt的签发和jwt的验证，所以，它就是你服务端的私钥，在任何场景都不应该流露出去。一旦客户端得知这个secret, 那就意味着客户端是可以自我签发jwt了。

## 应用

一般是在请求头里加入Authorization，并加上Bearer标注：
```js
fetch('api/user/1', {
  headers: {
    'Authorization': 'Bearer ' + token
  }
})
```

服务端会验证token，如果验证通过就会返回相应的资源。整个流程就是这样的:
```
  Browser(浏览器)                                       Server(服务器)
     │    1.Post /users/login with username and password   │
     │  ───────────────────────────────────────────────>   │
     │    3.Return the JWT to the Browser                  │  2.creates a JWT and a secret
     │  <───────────────────────────────────────────────   │
     │    4.Sends the JWT on the Authorization Header      │
     │  ───────────────────────────────────────────────>   │
     │    6.Sends response to the client                   │  5. Check JWT signature, Get user infomation from the JWT
     │  <───────────────────────────────────────────────   │
```

### 优点
因为json的通用性，所以JWT是可以进行跨语言支持的，像JAVA,JavaScript,NodeJS,PHP等很多语言都可以使用。  
因为有了payload部分，所以JWT可以在自身存储一些其他业务逻辑所必要的非敏感信息。  
便于传输，jwt的构成非常简单，字节占用很小，所以它是非常便于传输的。  
它不需要在服务端保存会话信息, 所以它易于应用的扩展  

### 安全相关
不应该在jwt的payload部分存放敏感信息，因为该部分是客户端可解密的部分。  
保护好secret私钥，该私钥非常重要。  
令牌一经签发后,理论上是不能销毁的,这是因为验证令牌的操作在应用服务器上,而不是验证服务器,但是我们可以通过共享缓存,比如redis来标记jti来回避这个问题.


#### 攻击面

##### 敏感信息泄露
JWT中的header 和 payload 虽然看起来不可读，但实际上都只经过简单编码， 开发者可能误将敏感信息存储在里面。使用上述工具可以方便地解码JWT中前两部分的信息。

##### 指定算法为none
上面提到算法 none 是JWT规范中强制要求实现的，但有些实现JWT的库直接将使用none 算法的token视为已经过校验。 这样攻击者就可以设置alg 为none ，使signature 部分为空，然后构造包含任意payload 的JWT来欺骗服务端。

##### 将签名算法从非对称类型改为对称类型
使用非对称加密算法（主要基于RSA、ECDSA，如S256）分发JWT的过程是使用私钥（private）加密生成JWT，使用公钥（public）解密验证。
使用对称加密算法（主要基于HMAC，如HS256）分发JWT的过程是使用同一个密钥（secret）生成和验证JWT。

如果服务端期待收到的算法类型为RS256，然后以RS256和public去验证JWT，而实际上收到的算法类型是HS256， 那么服务端就可能尝试把public当作secret，然后用HS256算法解密验证JWT。

由于RS256的public人人都可获得，攻击者可以预先以public为密钥，用HS256算法伪造包含任意payload 的JWT，从而成功通过服务端的验证。

##### 爆破密钥
WT的安全性依赖于密钥的保密性，任何拥有密钥的人都可以构造任何内容的合法token。

当一个JSON Web Token 被分发出去，如果密钥不够强壮就存在被爆破的风险，而且整个爆破过程可以离线进行。

已经有人写了一些工具，推荐如下：

jwtbrute
Sjord’ python script
John the Ripper

##### 伪造密钥
JWT采用header 中的kid 字段关联校验算法的密钥，这个密钥可能是对称加密的密钥，也可能是非对称加密的公钥。 如果能够猜测kid 和 密钥的关联性，攻击者就可能修改kid 来欺骗服务端，使其校验时使用攻击者可控的密钥， 于是攻击者就可以伪造任意内容的可通过校验的JWT。

##### 安全建议
验证函数应忽略JWT中的algo 字段，预先就明确JWT使用的算法，如果需要使用多种算法，可以在header 中使用表示”key ID” 的kid 字段，查询每个kid 对应的算法。 JWT/JWS 标准应该移除 header 中的algo 字段。JWT的许多安全缺陷都来自于开发者依赖这一客户端可控的字段。 开发者应升级相应库到最新版本，因为旧版本可能存在致命缺陷。

不应该在jwt的payload部分存放敏感信息，因为该部分是客户端可解密的部分。
保护好secret私钥，该私钥非常重要。
请使用https的安全通道进行传输
验证函数应忽略JWT中的algo 字段，预先就明确JWT使用的算法，如果需要使用多种算法， 可以在header中使用表示”key ID”的kid字段，查询每个kid对应的算法。
JWT/JWS 标准应该移除 header 中的algo 字段。JWT的许多安全缺陷都来自于开发者依赖这一客户端可控的字段。
开发者应升级相应库到最新版本，因为旧版本可能存在致命缺陷。

