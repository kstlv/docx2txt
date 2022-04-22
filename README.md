# docx2txt

Console tool for converting and viewing `.docx` (Microsoft Word) or `.odt` (OpenDocument Format) files.

![docx2txt help image](./help.png)

## Commands

Command | Description
--- | ---
`help` | Displays a list of available commands.
`version` | Displays the program version.
`view <filename>` | Displays text from a docx/odt file.
`convert <filename>` | Convert docx/odt file to text file.

## Usage examples

### View "recipe.docx"

Command (cmd):

```
docx2txt.exe view recipe.docx
```

Command (powershell)

```powershell
.\docx2txt.exe v .\recipe.docx
```

### Convert "notes.odt" to text file

Command (cmd):

```
docx2txt.exe convert notes.odt
```

Command (powershell)

```powershell
.\docx2txt.exe c .\notes.odt
```

## How to Build

Windows:

1. Download and install [golang](https://go.dev/).
2. Go to the source folder and `go build -ldflags "-s -w"`.
