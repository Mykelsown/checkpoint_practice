# Checkpoint Practice – Go Projects

> These are **practice questions** inspired by your completed projects. They are not the originals — they are reimagined variants to help you prepare for your checkpoint audit.

---

## 1. text-transformer

**Type:** Mandatory | Go

### Objective
Build a command-line tool that reads text from an input file, applies a series of transformations based on special tags embedded in the text, and writes the result to an output file.

```
$ go run . input.txt output.txt
```

### Transformations to implement

- Every instance of `(oct)` should replace the word before it with the **decimal** version of that word (the word will always be a valid octal number).
  - Ex: `"There are 17 (oct) apples"` → `"There are 15 apples"`

- Every instance of `(rev)` should **reverse** the word before it.
  - Ex: `"He said hello (rev) to her"` → `"He said olleh to her"`

- Every instance of `(title)` should convert the word before it to **Title Case**.
  - Ex: `"welcome (title) to the show"` → `"Welcome to the show"`

- Every instance of `(upper)` converts the word before it to **UPPERCASE**.
  - Ex: `"let's go (upper) team!"` → `"let's GO team!"`

- Every instance of `(lower)` converts the word before it to **lowercase**.
  - Ex: `"STOP (lower) right there"` → `"stop right there"`

- For `(upper)`, `(lower)`, and `(title)`, a number can be specified: `(upper, <n>)` transforms the **last n words** accordingly.
  - Ex: `"this is great news (upper, 2)"` → `"this is GREAT NEWS"`

- Punctuation marks `.` `,` `!` `?` `:` `;` must be **immediately attached** to the preceding word, with a single space before the next word.
  - Ex: `"Wait , I forgot something !"` → `"Wait, I forgot something!"`
  - Exception: grouped punctuation like `...` or `!?` should be kept together and treated as a single unit.

- Single quotes `'` must wrap the word(s) between them **without inner spaces**.
  - Ex: `"She called him ' a genius '"` → `"She called him 'a genius'"`

- Every instance of `a` before a word starting with a vowel (`a, e, i, o, u`) or `h` must be changed to `an`.
  - Ex: `"She rode a elephant"` → `"She rode an elephant"`

### Usage Example

```
$ cat input.txt
the hero (title) was brave , and a honest soul . ' truly remarkable '

$ go run . input.txt output.txt

$ cat output.txt
The hero was brave, and an honest soul. 'truly remarkable'
```

---

## 2. glyph-printer

**Type:** Mandatory | Go

### Objective
Build a CLI program that takes a string as an argument and prints it as large ASCII block text using a set of predefined banner template files.

Each character in the banner files spans **8 lines** in height. Characters are separated by a newline `\n`.

Three banner styles are available:
- `classic`
- `bold`
- `sketch`

### Usage Example

```
$ go run . "Hi"
 _   _   _
| | | | (_)
| |_| |  _
|  _  | | |
| | | | | |
|_| |_| |_|




$ go run . "Go\nCode"
 _____
/ ____|
| |  __   ___
| | |_ | / _ \
| |__| || (_) |
 \_____|  \___/



  _____              _
 / ____|            | |
| |       ___    __| |  ___
| |      / _ \  / _` | / _ \
| |____ | (_) || (_| ||  __/
 \_____|  \___/  \__,_| \___|


```

### Requirements
- Handle letters, numbers, spaces, and special characters.
- Handle `\n` in input as a line break.
- If the input is empty or just `\n`, print only empty lines accordingly.

---

## 3. glyph-align

**Type:** Optional | Go

### Objective
Extend `glyph-printer` to support text alignment using the flag `--align=<type>`.

Alignment types:
- `left`
- `right`
- `center`
- `justify`

The output must adapt to the **current terminal width**. Only inputs that fit the terminal will be tested.

Any invalid flag format must print:
```
Usage: go run . [OPTION] [STRING] [BANNER]
Example: go run . --align=right something classic
```

### Usage Example

```
$ go run . --align=center "Hi" classic

              _   _   _
             | | | | (_)
             | |_| |  _
             |  _  | | |
             | | | | | |
             |_| |_| |_|


```

---

## 4. glyph-paint

**Type:** Optional | Go

### Objective
Extend `glyph-printer` to support colorizing part or all of the ASCII output using the flag `--color=<color> [substring]`.

- `<color>` can be any valid color name or code (RGB, hex, ANSI — your choice).
- `[substring]` is the part of the input string to be colored. If omitted, the **entire output** is colored.
- The substring match should apply to **every occurrence** within the string.

Invalid flag format must print:
```
Usage: go run . [OPTION] [STRING]
EX: go run . --color=<color> <substring to be colored> "something"
```

### Usage Example

```
$ go run . --color=blue "Hi" "Hi there"
# "Hi" rendered in blue, "there" in default color
```

---

## 5. glyph-theme

**Type:** Optional | Go

### Objective
Extend `glyph-printer` so that the **banner style is selectable** via a second argument.

```
go run . [STRING] [BANNER]
```

- Supported banners: `classic`, `bold`, `sketch`
- Invalid usage must print:
```
Usage: go run . [STRING] [BANNER]
EX: go run . something classic
```
- The program must still work with just a single `[STRING]` argument (defaulting to `classic`).

### Usage Example

```
$ go run . "hey" bold
$ go run . "hey" sketch
$ go run . "hey"   # defaults to classic
```

---

## 6. glyph-save

**Type:** Optional | Go

### Objective
Extend `glyph-printer` to support writing the ASCII output to a file using the `--output` flag.

```
go run . --output=<fileName.txt> [STRING] [BANNER]
```

- The result must be written to the specified `.txt` file instead of being printed to stdout.
- Invalid flag format must print:
```
Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --output=<fileName.txt> something classic
```

### Usage Example

```
$ go run . --output=result.txt "hey" classic
$ cat -e result.txt
 _
| |__    ___   _   _
|  _ \  / _ \ | | | |
| | | ||  __/ | |_| |
|_| |_| \___|  \__, |
                |___/ $
...
```

---

## 7. glyph-web

**Type:** Mandatory | Go + HTML

### Objective
Create a Go web server that serves a browser-based interface for `glyph-printer`. Users should be able to type a string, select a banner style, and see the ASCII art rendered on the page.

### HTTP Endpoints

- `GET /` — serves the main HTML page
- `POST /render` — receives the string and banner selection, returns the ASCII art

### Main Page Must Include
- A text input field
- A way to select banner style (radio buttons, dropdown, etc.)
- A submit button that sends a `POST` to `/render`
- Display area for the rendered ASCII output

### HTTP Status Codes
- `200 OK` — successful render
- `400 Bad Request` — missing or invalid input
- `404 Not Found` — missing template or banner file
- `500 Internal Server Error` — unhandled errors

### Project Structure
```
project/
├── main.go
├── templates/
│   └── index.html
├── banners/
│   ├── classic.txt
│   ├── bold.txt
│   └── sketch.txt
└── README.md
```

### README.md Must Contain
- Description
- Authors
- Usage (how to run)
- Implementation details / algorithm

---

## 8. glyph-web-style

**Type:** Optional | Go + HTML + CSS

### Objective
Improve the `glyph-web` interface to make it more visually appealing, responsive, and user-friendly.

### Requirements
- Must include a CSS file linked to the HTML template.
- The website must be **consistent**, **responsive** (adapts to screen size), and **interactive** (feedback on actions).
- Text must remain readable regardless of colors chosen.
- Give users visual feedback (e.g., loading indicator, error messages, success states).

### Things to Consider
- Typography and spacing
- Color contrast and accessibility
- Mobile-friendly layout
- Smooth transitions or animations (optional but encouraged)

---

## General Reminders for the Audit

- All Go projects must use only **standard library packages**.
- Code must follow **Go good practices** (proper naming, error handling, package structure).
- It is recommended to write **unit tests** for your functions.
- Your programs must handle **edge cases**: empty strings, invalid flags, missing files, etc.
- For web projects, HTML templates must live in the `templates/` directory at the project root.
