# stopwords-iso
stopwords-iso is a go package that removes stop words from a text content

## Example

You can remove stopwords by language
```go
package main
import (
    sw "github.com/toadharvard/stopwords-iso" 
)

func main() {
	stopwordsMapping, _ := sw.NewStopwordsMapping()

	originalString := "This is a sample string with some stopwords."
	language := "en"

	clearedString := stopwordsMapping.ClearStringByLang(originalString, language)
	fmt.Printf("Cleared string: %s\n", clearedString)
}
```

or remove all stopwords from all supported languages

```go
package main
import (
    sw "github.com/toadharvard/stopwords-iso"
)
func main() {
	stopwordsMapping, _ := sw.NewStopwordsMapping()

	originalString := "the book on the table y la pluma es de ella und da Licht ist aus et la porte est ouverte и я it's"

	clearedString := stopwordsMapping.ClearString(originalString)
	fmt.Printf("Cleared string: %s\n", clearedString)
}
```

## Supported languages
This package uses the [stopwords-iso](https://github.com/stopwords-iso/) words pack. All languages supported by stopwords-iso are listed here: https://github.com/stopwords-iso/stopwords-iso?tab=readme-ov-file#credits

## License

Distributed under the MIT license.
See [LICENSE](LICENSE) for more information.
