# monkey

this project implement the monkey programming language.

- [Writing An Interpreter In Go](https://interpreterbook.com/)
- [Go言語でつくるインタプリタ](https://www.oreilly.co.jp/books/9784873118222/)

---
## トップダウン演算子順位解析(Pratt構文解析)
再帰下降構文解析(Recursive Descent Parsing)

---
## 関数呼び出しの構文解析は、(演算子と考える

以下のようにgreeting関数を用意して発想の切り替え方を説明します。

```
let greeting = fun(name) {
    print("hi!" + name)
}

greeting("john")
```

ふつうプログラミングをするときは、greeting関数を、関数名+(文字列)のように閉じかっこまで含めて考えますが、この構文解析では最後の)右かっこを無視して、演算子と考えます。

```
   greeting ( "john"
   1 + 1
   3 * 4
```
こうすると、加算や乗算のような中置の演算子と同じと考えられます。よって関数呼び出し用の特別な構文解析は必要なく、単純に最も優先順位の高い中置演算子(を実装していくだけになります。

## おすすめのソースコード

[wren](https://github.com/munificent/wren) は二種類の値表現がある。

## IO

|  input  |  process  |  output  |
| ------- | --------- | -------- |
|  text   |   lexer   |  token   |
|  token  |   parser  |  AST     |
|  AST    |   eval    |  Object  |

## Truthy

> a truthly value is a value that is considered `true` 
> when encountered in a Boolean context.

[MDN Web Docs Glossary - Truthy](https://developer.mozilla.org/en-US/docs/Glossary/Truthy)

言語によって真(true)の解釈は異なります。Monkeyの場合は、nullでなく、falseでもないものは真とするので、以下のコードはeverything okay!が表示されます。

```ruby
let x = 10;
if (x) {
    puts("everything okay!");
} else {
    puts("x is too high!");
    shutdownSystem();
}
```