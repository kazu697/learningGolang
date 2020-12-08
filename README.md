# learningGolang

## Golangのインストール(windows)
下記のサイトからmsiファイルを実行

https://golang.org/dl/

## 実行方法
go run main.go

## moduleを作成する
go mod init 

## 自作packageのインポート
1. GOPATHのパス配下にsrcフォルダを作成する
2. srcフォルダ配下にプロジェクトを作成
3. プロジェクトにpackage化したいものをフォルダにする

## testingモジュール
1. {モジュール名}_test.goというファイルを同じディレクトリに作成する
2. 下記のコマンドでテストを実行する

    go test ./...