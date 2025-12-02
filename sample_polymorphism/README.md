# Polymorphism Variants

このディレクトリでは 3 つの多態パターンを比較しています。すべて `go run ./sample_polymorphism` で動作確認できます。

## Version 1 (interface + 値レシーバ)
- 概要: `Notifier1` interface を直接スライスで扱い、実装は値レシーバ。
- Pros: シンプルで idiomatic、mock も楽。
- Cons: ポインタで状態を持たせたい場合に値コピーを意識する必要がある。

## Version 2 (channel interface 埋め込み)
- 概要: `Notifier2` がフィールドに `channel2` interface を保持し、コンストラクタで実装を隠蔽。
- Pros: 呼び出し側は統一された `Notify` のみを使い、内部実装差し替えが容易。
- Cons: ラッパ構造体とインタフェースの二重抽象でやや冗長、nil チェックが必要。

## Version 3 (タグ付き struct 分岐)
- 概要: タグ `notifierType` でメール/SMS を切り替え、バックエンドをポインタで保持。
- Pros: 型スイッチ不要でタグに従って明示的に分岐でき、追加チャネルを列挙型で管理しやすい。
- Cons: タグとフィールドの整合性が手動管理になるため、メンテでミスが入りやすい。実装追加時はタグ定数と switch を忘れないよう注意。
