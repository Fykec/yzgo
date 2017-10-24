# Youzan Golang SDK


## 通过App Id 和 App Secret来调用 
``` go
client := YZClient{appID: "<AppID>", appSecret: "<AppSecret>"}
	ret := client.Invoke("kdt.shop.basic.get", "1.0.0", "GET", map[string]string{}, map[string]string{})
	println(ret)
```

## 通过OAuth 的access token 来调用 通过appId和appSecret获取accessToken
```bash
curl -X POST https://open.youzan.com/oauth/token -H 'content-type: application/x-www-form-urlencoded' -d 'client_id=testclient&client_secret=testclientsecret&grant_type=silent&kdt_id=88888'
```

```go
    client := YZClient{isOAuth: true, accessToken: "<accessToken>"}
	ret := client.Invoke("kdt.shop.basic.get", "1.0.0", "GET", map[string]string{}, map[string]string{})
	println(ret)
```

## Go Test 

### 在shell中设置 YZAccessToken

```
export YZAccessToken=<Your Access Token>
```

### Test

```go
go test
```