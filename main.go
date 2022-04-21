package main

import (
	"fmt"
	"strings"
	"regexp"
	"archive/zip"
	"io"
	"os"
)

const programName = "docx2txt"
const programVersion = "1.0"

func main() {
	if(len(os.Args) <= 1){
		printMessage("Please run this tool with correct arguments. Help will be displayed now.")
		printHelp()
	} else {
		command := os.Args[1]

		if(command == "view" || command == "v"){
			if(len(os.Args) >= 3){
				getText(false)
			} else {
				printGoodbyeMessage("To view, you must provide the path to the file.")
			}
		} else if (command == "convert" || command == "c"){
			if(len(os.Args) >= 3){
				getText(true)
			} else {
				printGoodbyeMessage("To convert, you must provide the path to the file.")
			}
		} else if (command == "help" || command == "h" || command == "?"){
			printHelp()
		} else if (command == "version"){
			printGoodbyeMessage(programVersion)
		} else {
			printGoodbyeMessage("Unknown command.")
		}
	}
}

func getText(saveTextFile bool){
	fileName := os.Args[2]

	const docx = ".docx"
	const odt = ".odt"
	const txt = ".txt"

	var text string
	var xmlFile string
	var fileExtension string

	reg := regexp.MustCompile("</w:t>")

	if(strings.HasSuffix(fileName, docx)){
		fileExtension = docx
		xmlFile = "word/document.xml"
	} else if(strings.HasSuffix(fileName, odt)){
		fileExtension = odt
		xmlFile = "content.xml"
		reg = regexp.MustCompile("</text:.?")
	} else {
		printGoodbyeMessage("File not supported. This tool only works with \"" + docx + "\" or \"" + odt + "\".")
	}

	buf := new(strings.Builder)

	r, err := zip.OpenReader(fileName)
	if err != nil {
		printGoodbyeMessage("File does not exist.")
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name != xmlFile {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			printGoodbyeMessage("Failed to open file.")
		}
		_, err = io.Copy(buf, rc)
		if err != nil {
			printGoodbyeMessage("Failed to get data from file.")
		}

		text = buf.String()

		rc.Close()
		break
	}

	if(len(text) == 0){
		printGoodbyeMessage("Something went wrong.")
	}
	
	split := reg.Split(text, -1)
	text = ""

	reg = regexp.MustCompile(".*>")
	
	for i := range split {
		split2 := reg.Split(split[i], -1)
		text += split2[1] + "\n"
	}

	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&quot;", "\"")

	text = strings.TrimSpace(text)

	if(saveTextFile){
		fileTxt := strings.TrimSuffix(fileName, fileExtension) + txt

		textFile, createErr := os.Create(fileTxt)
		if createErr != nil {
			printGoodbyeMessage("Failed to create text file.")
		}
	
		defer textFile.Close()
	
		_, writeErr := textFile.WriteString(text)
		if writeErr != nil {
			printGoodbyeMessage("Failed to write data to text file.")
		}

		printGoodbyeMessage("Saved to text file: \"" + fileTxt + "\".")
	} else {
		printMessage("Text from \"" + fileName + "\":\n\n" + text)
	}
}

func printMessage(message string){
	fmt.Println("[" + programName + "] " + message)
}

func printGoodbyeMessage(message string){
	printMessage(message)
	os.Exit(1)
}

func printHelp(){
	fmt.Println("     _                 ____  _        _   ")
	fmt.Println("  __| | ___   _____  _|___ \\| |___  _| |_ ")
	fmt.Println(" / _` |/ _ \\ / __\\ \\/ / __) | __\\ \\/ / __|")
	fmt.Println("| (_| | (_) | (__ >  < / __/| |_ >  <| |_ ")
	fmt.Println(" \\__,_|\\___/ \\___/_/\\_\\_____|\\__/_/\\_\\__|")
	fmt.Println()
	fmt.Println(programName + " ver. " + programVersion + " by Danil Kostylev")
	fmt.Println()
	fmt.Println("Usage: " + programName + ".exe command\n\nCommands:")
	fmt.Println()
	fmt.Println("+--------------------+----------------------------------------+")
	fmt.Println("| Command            | Description                            |")
	fmt.Println("+--------------------+----------------------------------------+")
	fmt.Println("| help               | Displays a list of available commands. |")
	fmt.Println("| version            | Displays the program version.          |")
	fmt.Println("| view <filename>    | Displays text from a docx/odt file.    |")
	fmt.Println("| convert <filename> | Convert docx/odt file to text file.    |")
	fmt.Println("+--------------------+----------------------------------------+")
	os.Exit(1)
}
