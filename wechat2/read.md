微信小程序无法在前端直接获取用户的手机号，只能获取到aes加密后的手机号信息和一个code。将加密后的手机信息和code传到我们自己写的服务就可以解密了。

解密需要两个步骤：

1.使用code从微信API获取session key。

直接使用以下参数，对api发起get请求

https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=APPSECRET&js_code=CODE&grant_type=authorization_code
属性	描述
appid
微信小程序的appid,从微信小程序平台获取，每个小程序都不一样
secret
微信小程序的app secret，从微信小程序平台获取，每个小程序都不一样
js_code
小程序前端获取用户手机号信息时的code
grant_type
固定值 authorization_code

2.使用session key解密加密的手机信息。