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

```
let x = 10;
if (x) {
    puts("everything okay!");
} else {
    puts("x is too high!");
    shutdownSystem();
}
```

## 3.7 return 文

### `3.7 return 文`の evalBlockStatementの動きの説明が難しかったので整理

```
if (10 > 1) {
    return 10
}
```
上記ソースコードは、条件文が(10 > 1)でtrueなので、ブロック文{に入り、object.ReturnValue{Value:10}を返す。

次にブロック文がネストされている状態を考えてみます。以下のようなコードはプログラミング言語で一般的にある構造です。

```
if (10 > 1) {
    if (5 > 2) {
        return 10
    }

    return 1
}
```

(10 > 1)のIf式が真なので、ブロック文{
を評価します。

```
if (10 > 1) { <-このブロックを評価。
    if (5 > 2) {
        return 10
    }

    return 1
}
```

ネストのIf式(5 > 2)は真なので、そのブロック文を評価します。

```
if (10 > 1) { 
    if (5 > 2) {<-このブロックを評価。
        return 10
    }

    return 1
}
```

ブロック文には他の式や分はなく、Return文１つしかないためreturn文は、object.ReturnValue{Value:10}と評価されます。**そして、if式の戻り値となります。**

以下の状態です。

```
if (10 > 1) {
    object.ReturnValue{Value:10}

    return 1
}
```

以下、Page.148で実装したevalBlockStatementでは、評価した値がobject.ReturnValueなら、その値を戻り値としてevalBlockStatementのを終了します。

```
func evalBlockStatement(block *ast.BlockStatement) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
        // ここの評価は、単純な文の場合もあるが、
        // If式とその後のブロック文つまり、
        // 今まで説明してきたネストのIf文を評価する場合
        // もある。その時は、ネスト文のobject.ReturnValue
        // が返ってきて、このresult変数に代入される。
		result = Eval(statement)

		// ここは、単純にそのブロック文のreturnを返すためだけではない。
        //
		if result != nil && result.Type() == object.RETURN_VALUE_OBJ {
			return result
		}
	}

	return result
}
```
