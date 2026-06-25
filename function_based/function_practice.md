# Function Practice – Core Logic Drills

> These are isolated function challenges based on the core logic used across your projects.
> No full program needed — just write and test each function individually.
> All solutions should be written in **Go**.

---

## Section 1 – Number Conversion

---

### 1.1 `HexToDecimal`

Write a function that takes a string representing a **hexadecimal number** and returns its decimal equivalent as an integer. Do not use `strconv.ParseInt` or any conversion library — implement the logic manually.

```go
func HexToDecimal(hex string) (int, error)
```

**Examples:**
```
HexToDecimal("1E")   → 30, nil
HexToDecimal("FF")   → 255, nil
HexToDecimal("0")    → 0, nil
HexToDecimal("XYZ")  → 0, error
```

**Constraints:**
- Input may be upper or lowercase (`"1e"` and `"1E"` are both valid).
- Return an error for invalid hex strings.

---

### 1.2 `BinaryToDecimal`

Write a function that takes a string representing a **binary number** and returns its decimal equivalent as an integer. Implement the conversion logic manually without using parsing libraries.

```go
func BinaryToDecimal(bin string) (int, error)
```

**Examples:**
```
BinaryToDecimal("1010")  → 10, nil
BinaryToDecimal("1")     → 1, nil
BinaryToDecimal("0")     → 0, nil
BinaryToDecimal("102")   → 0, error
```

**Constraints:**
- Return an error if the string contains characters other than `0` and `1`.

---

### 1.3 `OctalToDecimal`

Write a function that takes a string representing an **octal number** and returns its decimal equivalent as an integer. Implement manually.

```go
func OctalToDecimal(oct string) (int, error)
```

**Examples:**
```
OctalToDecimal("17")   → 15, nil
OctalToDecimal("100")  → 64, nil
OctalToDecimal("89")   → 0, error
```

---

### 1.4 `DecimalToBase`

Write a function that converts a **decimal integer** to a string representation in any base from 2 to 16.

```go
func DecimalToBase(n int, base int) (string, error)
```

**Examples:**
```
DecimalToBase(30, 16)  → "1E", nil
DecimalToBase(10, 2)   → "1010", nil
DecimalToBase(15, 8)   → "17", nil
DecimalToBase(10, 1)   → "", error  (invalid base)
```

---

## Section 2 – String Case Manipulation

---

### 2.1 `ToUppercase`

Write a function that converts a string to **all uppercase** without using `strings.ToUpper`.

```go
func ToUppercase(s string) string
```

**Examples:**
```
ToUppercase("hello")   → "HELLO"
ToUppercase("Go Lang") → "GO LANG"
ToUppercase("123!abc") → "123!ABC"
```

---

### 2.2 `ToLowercase`

Write a function that converts a string to **all lowercase** without using `strings.ToLower`.

```go
func ToLowercase(s string) string
```

**Examples:**
```
ToLowercase("HELLO")   → "hello"
ToLowercase("Go Lang") → "go lang"
```

---

### 2.3 `Capitalize`

Write a function that **capitalizes the first letter** of a string and lowercases the rest.

```go
func Capitalize(s string) string
```

**Examples:**
```
Capitalize("hello")   → "Hello"
Capitalize("WORLD")   → "World"
Capitalize("gO")      → "Go"
Capitalize("")        → ""
```

---

### 2.4 `CapitalizeN`

Write a function that takes a slice of words and a number `n`, and returns a new slice where the **last n words** are capitalized.

```go
func CapitalizeN(words []string, n int) []string
```

**Examples:**
```
CapitalizeN(["it", "was", "the", "best", "of", "times"], 3)
→ ["it", "was", "the", "Best", "Of", "Times"]

CapitalizeN(["go", "is", "fun"], 10)
→ ["Go", "Is", "Fun"]   // n > len: apply to all
```

---

### 2.5 `ReverseString`

Write a function that reverses a string **character by character**.

```go
func ReverseString(s string) string
```

**Examples:**
```
ReverseString("hello")   → "olleh"
ReverseString("Go!")     → "!oG"
ReverseString("")        → ""
ReverseString("a")       → "a"
```

---

### 2.6 `IsPalindrome`

Write a function that returns `true` if a string reads the same forwards and backwards (case-insensitive, ignoring spaces).

```go
func IsPalindrome(s string) bool
```

**Examples:**
```
IsPalindrome("racecar")      → true
IsPalindrome("A man a plan") → true
IsPalindrome("hello")        → false
IsPalindrome("")             → true
```

---

## Section 3 – Punctuation & Spacing

---

### 3.1 `FixPunctuation`

Write a function that ensures the punctuation marks `. , ! ? : ;` are **directly attached to the preceding word** and separated from the following word by a single space.

```go
func FixPunctuation(s string) string
```

**Examples:**
```
FixPunctuation("Hello , world !")         → "Hello, world!"
FixPunctuation("Wait . Come back .")      → "Wait. Come back."
FixPunctuation("Yes ; no ; maybe")        → "Yes; no; maybe"
```

**Constraint:** Groups like `...` or `!?` must be kept together as a single unit and treated the same way.

---

### 3.2 `FixQuotes`

Write a function that finds pairs of single quotes `'` in a string and removes any spaces between the quotes and the words they wrap.

```go
func FixQuotes(s string) string
```

**Examples:**
```
FixQuotes("She is ' awesome '")
→ "She is 'awesome'"

FixQuotes("As he said : ' I will return '")
→ "As he said : 'I will return'"

FixQuotes("' wow '")
→ "'wow'"
```

---

### 3.3 `FixArticles`

Write a function that scans a string and replaces `a` with `an` when the next word starts with a vowel (`a, e, i, o, u`) or the letter `h`.

```go
func FixArticles(s string) string
```

**Examples:**
```
FixArticles("She rode a elephant")     → "She rode an elephant"
FixArticles("He is a honest man")      → "He is an honest man"
FixArticles("I have a cat")            → "I have a cat"
FixArticles("It was a unique chance")  → "It was a unique chance"
```

**Note:** Only standalone lowercase `a` should be replaced.

---

### 3.4 `NormalizeSpaces`

Write a function that removes **extra whitespace** between words, trimming leading and trailing spaces, and ensuring only a single space exists between each word.

```go
func NormalizeSpaces(s string) string
```

**Examples:**
```
NormalizeSpaces("  hello   world  ")   → "hello world"
NormalizeSpaces("go   is   great")     → "go is great"
NormalizeSpaces("")                    → ""
```

---

## Section 4 – ASCII & Character Logic

---

### 4.1 `IsVowel`

Write a function that returns `true` if a given `rune` is a vowel (a, e, i, o, u) — case-insensitive.

```go
func IsVowel(c rune) bool
```

**Examples:**
```
IsVowel('a')  → true
IsVowel('E')  → true
IsVowel('b')  → false
IsVowel('U')  → true
```

---

### 4.2 `IsPrintableASCII`

Write a function that returns `true` if all characters in a string are **printable ASCII** characters (codes 32–126).

```go
func IsPrintableASCII(s string) bool
```

**Examples:**
```
IsPrintableASCII("Hello!")   → true
IsPrintableASCII("hi\nthere")→ false  (\n is not printable)
IsPrintableASCII("")         → true
IsPrintableASCII("café")     → false  (é is non-ASCII)
```

---

### 4.3 `CharFrequency`

Write a function that returns a map of each character in a string to how many times it appears.

```go
func CharFrequency(s string) map[rune]int
```

**Examples:**
```
CharFrequency("hello")
→ map['h':1, 'e':1, 'l':2, 'o':1]

CharFrequency("")
→ map[]
```

---

### 4.4 `PadLeft` / `PadRight` / `PadCenter`

Write three functions that pad a string to a given total width using a specified fill character.

```go
func PadLeft(s string, width int, fill rune) string
func PadRight(s string, width int, fill rune) string
func PadCenter(s string, width int, fill rune) string
```

**Examples:**
```
PadLeft("hi", 6, ' ')    → "    hi"
PadRight("hi", 6, ' ')   → "hi    "
PadCenter("hi", 6, ' ')  → "  hi  "
PadCenter("hi", 7, '-')  → "--hi---"  // odd padding: extra on the right
```

> These are directly related to text alignment in terminal output.

---

## Section 5 – Slice & Word Utilities

---

### 5.1 `SplitWords`

Write a function that splits a string into a slice of words by whitespace, without using `strings.Split` or `strings.Fields`.

```go
func SplitWords(s string) []string
```

**Examples:**
```
SplitWords("hello world")      → ["hello", "world"]
SplitWords("  go  is  great ") → ["go", "is", "great"]
SplitWords("")                 → []
```

---

### 5.2 `JoinWords`

Write a function that joins a slice of strings into a single string with a given separator.

```go
func JoinWords(words []string, sep string) string
```

**Examples:**
```
JoinWords(["hello", "world"], " ")  → "hello world"
JoinWords(["a", "b", "c"], ", ")    → "a, b, c"
JoinWords([], " ")                  → ""
```

---

### 5.3 `CountOccurrences`

Write a function that counts how many times a **substring** appears in a string (non-overlapping).

```go
func CountOccurrences(s, sub string) int
```

**Examples:**
```
CountOccurrences("a king kitten have kit", "kit")  → 2
CountOccurrences("hello", "ll")                    → 1
CountOccurrences("aaa", "aa")                      → 1  // non-overlapping
CountOccurrences("hello", "xyz")                   → 0
```

---

### 5.4 `ReplaceNth`

Write a function that replaces only the **nth occurrence** of a substring in a string (1-indexed).

```go
func ReplaceNth(s, old, new string, n int) string
```

**Examples:**
```
ReplaceNth("cat cat cat", "cat", "dog", 2)  → "cat dog cat"
ReplaceNth("aaa", "a", "b", 1)              → "baa"
ReplaceNth("hello", "x", "y", 1)            → "hello"  // no match
```

---

### 5.5 `UniqueWords`

Write a function that returns a slice of words from the input string with **duplicates removed**, preserving order of first appearance.

```go
func UniqueWords(s string) []string
```

**Examples:**
```
UniqueWords("the cat sat on the mat")
→ ["the", "cat", "sat", "on", "mat"]

UniqueWords("go go go")
→ ["go"]
```

---

## Section 6 – File I/O Utilities

---

### 6.1 `ReadFileToString`

Write a function that reads the entire content of a file and returns it as a string.

```go
func ReadFileToString(filename string) (string, error)
```

**Behaviour:**
- Return an error if the file does not exist or cannot be opened.
- Return the full content as a string if successful.

---

### 6.2 `WriteStringToFile`

Write a function that writes a string to a file, **creating the file if it doesn't exist** and overwriting it if it does.

```go
func WriteStringToFile(filename, content string) error
```

**Behaviour:**
- Create the file with appropriate permissions (`0644`).
- Return an error if writing fails.

---

### 6.3 `FileExists`

Write a function that returns `true` if a file exists at the given path, `false` otherwise.

```go
func FileExists(path string) bool
```

**Examples:**
```
FileExists("sample.txt")   → true  (if file is present)
FileExists("missing.txt")  → false
```

---

### 6.4 `GetFileExtension`

Write a function that extracts the **file extension** from a filename string (without using `filepath` package).

```go
func GetFileExtension(filename string) string
```

**Examples:**
```
GetFileExtension("result.txt")     → ".txt"
GetFileExtension("archive.tar.gz") → ".gz"
GetFileExtension("README")         → ""
GetFileExtension(".hidden")        → ""
```

---

## Section 7 – Flag Parsing

---

### 7.1 `ParseFlag`

Write a function that parses a CLI-style flag string of the format `--key=value` and returns the key and value separately.

```go
func ParseFlag(arg string) (key, value string, err error)
```

**Examples:**
```
ParseFlag("--output=result.txt")  → "output", "result.txt", nil
ParseFlag("--align=center")       → "align", "center", nil
ParseFlag("--color=red")          → "color", "red", nil
ParseFlag("output=result.txt")    → "", "", error  // missing --
ParseFlag("--badformat")          → "", "", error  // missing =
```

---

### 7.2 `HasFlag`

Write a function that checks if a given flag key (e.g. `"output"`) exists in a slice of command-line arguments.

```go
func HasFlag(args []string, flagKey string) bool
```

**Examples:**
```
HasFlag(["--output=file.txt", "hello", "standard"], "output")  → true
HasFlag(["hello", "standard"], "align")                        → false
```

---

### 7.3 `ExtractFlagValue`

Write a function that finds a flag by key in a slice of args and returns its value. Returns an empty string and an error if not found.

```go
func ExtractFlagValue(args []string, flagKey string) (string, error)
```

**Examples:**
```
ExtractFlagValue(["--color=blue", "hello"], "color")   → "blue", nil
ExtractFlagValue(["hello", "standard"], "color")       → "", error
```

---

## Bonus Challenges

### B.1
Write a function `WrapText(s string, lineWidth int) string` that wraps a string so no line exceeds `lineWidth` characters, breaking only at word boundaries.

### B.2
Write a function `CenterInTerminal(s string, termWidth int) string` that returns the string padded with spaces so it appears centered within a terminal of given width.

### B.3
Write a function `ApplyTransformN(words []string, n int, fn func(string) string) []string` that applies a transformation function to the **last n** elements of a words slice and returns the full slice.

### B.4
Write a function `ParseColorCode(input string) (r, g, b int, err error)` that parses a color string in either `rgb(r,g,b)` format or hex format `#RRGGBB` and returns the RGB components.

---

## Tips for Practice

- Write each function in its own `.go` file with a matching `_test.go` file.
- Use **table-driven tests** — define a slice of `{input, expected}` structs and loop through them.
- Try implementing without standard library helpers first, then compare with the library version.
- Focus on **edge cases**: empty strings, zero values, `n` larger than slice length, invalid inputs.
