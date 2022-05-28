# monkey

this project implement the monkey programming language.

- [Writing An Interpreter In Go](https://interpreterbook.com/)
- [Go言語でつくるインタプリタ](https://www.oreilly.co.jp/books/9784873118222/)

## 2.6.2 トップダウン演算子順位解析(Pratt構文解析)

再帰下降構文解析(Recursive Descent Parsing)

### 2.6.4 ASTの準備

疑問点monkey言語で式文のみを実行したらどうなるのか。

```monkey
let x = 5;
x + 10;
```
