# docx2txt

Console tool for converting and viewing `.docx` (Microsoft Word) or `.odt` (OpenDocument Format) files.

## Commands

Command | Description
--- | ---
`help` | Displays a list of available commands.
`version` | Displays the program version.
`view <filename>` | Displays text from a docx/odt file.
`convert <filename>` | Convert docx/odt file to text file.

## Usage examples

View `recipe.docx`:

```
docx2txt.exe view recipe.docx
```

Convert `notes.odt` to text file:

```
docx2txt.exe convert notes.odt
```

## How to Build

Windows:

1. Download and install [golang](https://go.dev/)
2. Go to the source folder and `go build -ldflags "-s -w"`
