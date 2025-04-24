# fw-sample
fw-sample は、シンプルな HTTP サーバーフレームワークのサンプル実装です。このフレームワークは、ルーティング機能を提供し、GET および POST リクエストを処理することができます。

# 特徴
- 簡単なルーティング: URL パスと HTTP メソッドに基づいてハンドラを登録し、リクエストを処理します。
- HTTP インターフェース準拠: http.Handler インターフェースを実装しており、標準の http.ListenAndServe に直接渡すことができます。
- 拡張可能: 必要に応じて機能を追加してカスタマイズ可能です。

# 使用方法
1. フレームワークの初期化
New 関数を使用して App インスタンスを作成します。
```go
app := fwsample.New()
```

1. ルートの登録
Router の Get または Post メソッドを使用してルートを登録します。
```go
app.Router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
})
```

1. サーバーの起動
Run メソッドを使用してサーバーを起動します。
```go
if err := app.Run(":8080"); err != nil {
    log.Fatal(err)
}
```


# サンプルコード
以下は、簡単なサンプルアプリケーションのコードです。

```go
package main

import (
    "log"
    "net/http"
    "fwsample"
)

func main() {
    app := fwsample.New()

    // ルートの登録
    app.Router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    app.Router.Post("/echo", func(w http.ResponseWriter, r *http.Request) {
        body := make([]byte, r.ContentLength)
        r.Body.Read(body)
        w.Write(body)
    })

    // サーバーの起動
    log.Println("Starting server on :8080")
    if err := app.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
```


