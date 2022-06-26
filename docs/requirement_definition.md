# 要件定義
* アマゾンのほしい物リストのような感覚で欲しい車リストを公開したい
* 管理者は車の画像、情報をアマゾンの商品のように登録できる
* 一般ユーザは車一覧から選んで自分のリストに入れることができる
* 一般ユーザーが発行した欲しい車リストのURLを叩いたゲストユーザーはその人の欲しい車リストから車の情報を見ることができる
* 車の詳細画面ではその車の画像やリストに入れている人の数, 情報(年式, メーカー, 駆動方式, 排気量, 大きさ...)をみれる
* 車の詳細画面でコメントできるようにすると楽しそう？
* 認証はauth0を使う
* リストに入れている人の数で人気車ランキングとか良さげ

# db設計
必要なテーブル
* users
* cars
* manufacturers
* countries
* wish_lists
* comments