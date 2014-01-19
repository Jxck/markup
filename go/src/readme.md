# TCP Fast Open

HTTPは、その下層にあたるトランスポートレイヤーのプロトコルとして、通常TCPを使用します。
したがって、TCPのレイヤで速度が改善することは、そのままWebの高速化につながる可能性があるといえます。

GoogleはWebを速くするための活動として、TCPのようなプロトコルレイヤの改善にも取り組んでいます。
今回はその中の一つ、TCP Fast Openを取り上げます。

[TCP Fast Open](http://tools.ietf.org/html/draft-ietf-tcpm-fastopen-05)

検証環境等は最下部に記載します.


# 3 Way Handshake

TCPは、「正確、確実にデータを届ける」ことを重視した設計になっています。
特に接続確立時には双方の状態をきちんと確認しながら実施するための _3 Way Handshake_ (以下3WH)という方式を採用しています。

例えば、クライアントがサーバに対してHTTPリクエストを送信したい場合は、以下のような流れになります。

![3 Way Handshake](https://cacoo.com/diagrams/0bnHpQSj05mRkpdz-BCDBB.png)


```go
import (
  "hoge"
)

func hoge() {
  log.Println("hoge")
}
```
