# Behavioral » Interpreter

## Elements

 - Interpreter: the one who interpret the language;
 - Language: the whole pattern we need to be interpret;

## Description

The interpreter design pattern is used to solve business cases where it's
useful to have a language to perform common operations. The most famous
interpreter we can talk about is probably SQL.

## Implementation

The implementation is trivial. In the following examples I'll show a very
simple usage of interpreter. In this case we use interpreter to calculate some
mathematical operations. For example the sum and the multiplication operations.

```go
func TestSumOperator(t *testing.T) {
	sentence := "5 + 3"
	i := interpreter{}
	_ = i.of(sentence)
	if i.exec() != 8 {
		t.Error([]string{
			"La somma di 5 con 3",
			"non dovrebbe essere",
			strconv.Itoa(i.exec()),
		})
	}
}
```

As you can see, given a sentence like `5 + 3` parsed, it is possible to exec
the string and calculate the result.

```go
func TestMulOperator(t *testing.T) {
	sentence := "5 * 3"
	i := interpreter{}
	_ = i.of(sentence)
	if i.exec() != 15 {
		t.Error([]string{
			"Multiplication between 5 and 3",
			"shouldnt be",
			strconv.Itoa(i.exec()),
		})
	}
}
```
