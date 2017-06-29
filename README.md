# PASH - Punch A Small Hole on PNG

twitter に画像をアップロードする際に汚くなってしまうのを防止するため、
画像に穴を開け透過PNGへ加工するプログラムです。

See also: <https://twitter.com/Fate_uri/status/879572420872384512>

## 使い方

pash.bat もしくは pash.exe に変換したい画像ファイルを D&D してください。
入力画像のフォーマットは PNG と JPEG をサポートしています。

加工後の画像は `{ファイル名}+pash.png` という名前で出力します。
同名のファイルが既に存在した場合は強制的に上書きします。

pash.bat の実行には pash.exe が必要です。
pash.exe は変換に失敗した場合にエラーメッセージを表示しますが、
D&D で実行した場合には一瞬で閉じてしまい確認できません。
pash.bat はそのエラーを確認できるようにしています。

## ビルド方法

    $ go get -u github.com/koron/pash
