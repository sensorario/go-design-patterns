# Chain of responsibility

## Elements

 - Handler - an interface for requests handling
 - RequestHandler - handles requests it is responsible for
 - Client - sends commands to the first object in the chain that may handle the command

## Description

It consists of a source of command objects and a series of processing objects.
Each processing object contains logic that defines the types of command objects
that it can handle. The rest are passed to the next processing object in the
chain.

Is an object oriented version of the if ... else if ... else if ....... else
... endif idiom, with the benefit that the condition–action blocks can be
dynamically rearranged and reconfigured at runtime.

## Implementation

In this example the program aims to find right translation for a particular
word. Languages available are English, French and Spanish. The chain is
represented by the list of translator. Each translator knows who's next
translator. Provide Next Translator in list. Expose a method to tell if it
knows or not a word. Finally, … provide translation of a known word.

Here the interface of each translator: EnglishTranslator, SpanishTranslator and
FrenchTranslator.

```go
type TranslatorRing interface {
	Next(TranslatorRing)
	GetNext() TranslatorRing
	KnowsWord(s string) bool
	TranslationOf(s string) string
}
```

Here an example of concrete translator. We have a struct that known who's next.
Assign next translator. Provide next translator and translated word. As
example, EnglishTranslator could be the following.

```go
type EnglishTranslator struct {
	next TranslatorRing
}

func (fr *EnglishTranslator) Next(r TranslatorRing) {
	fr.next = r
}

func (fr *EnglishTranslator) GetNext() TranslatorRing {
	return fr.next
}

func (fr *EnglishTranslator) KnowsWord(s string) bool {
	dict := NewEnglishDict()
	if _, ok := dict[s]; ok {
		return true
	}
	return false
}

func (fr *EnglishTranslator) TranslationOf(s string) string {
	dict := NewEnglishDict()
	if val, ok := dict[s]; ok {
		return val
	}
	panic("Oops!")
}

func NewEnglishDict() map[string]string {
	dict := map[string]string{}
	dict["topo"] = "mouse"
	dict["cocomero"] = "watermelon"
	return dict
}
```

Finally, we have a client that works with the entire chain.

```go
type TranslatorChain struct {
	Start TranslatorRing
}

func (h *TranslatorChain) Translate(s string) string {
	r := h.Start
	for {
		if r.KnowsWord(s) {
			return r.TranslationOf(s)
		}
		if r.GetNext() != nil {
			r = r.GetNext()
		} else {
			log.Fatal("No translation found")
		}
	}
}

func (h *TranslatorChain) CountRings() int {
	numOfRings := 0
	if h.Start != nil {
		numOfRings++
	}
	r := h.Start
	for r.GetNext() != nil {
		r = r.GetNext()
		numOfRings++
	}
	return numOfRings

}
```

A concrete example could be the following. Only spanish dictionary known
  translation of "ciao". Only french translator knows translation of word
  "casa". But what if both english and spanish vocabulary knowns word
  "cocomero"? Because of the EnglishTranslator have more priority in the chain,
  .. `chain.Translate("cocomero")` will return english word "watermelon" and
  not spanish "sandía".

```go
func main() {
	chain := TranslatorChain{
		Start: &EnglishTranslator{
			&FrenchTranslator{
				&SpanishTranslator{},
			},
		},
	}

	fmt.Println("vim-go")

	fmt.Println(strconv.Itoa(chain.CountRings())) // 3
	fmt.Println(chain.Translate("ciao"))          // spanish: hola
	fmt.Println(chain.Translate("casa"))          // french: maison
	fmt.Println(chain.Translate("topo"))          // english mouse
	fmt.Println(chain.Translate("cocomero"))          // english mouse
}
```
