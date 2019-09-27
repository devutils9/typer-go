package game

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strings"

	"github.com/shilangyu/typer-go/settings"
	"github.com/shilangyu/typer-go/utils"
)

// ChooseText randomly chooses a text from the dataset
func ChooseText() (string, error) {
	if _, err := os.Stat(settings.I.TextsPath); os.IsNotExist(err) {
		return "", errors.New("Didnt find typer texts, make sure your path is correct")
	}

	bytes, err := ioutil.ReadFile(path.Join(utils.Root(), "texts.txt"))
	if err != nil {
		return "", nil
	}
	content := string(bytes)
	texts := strings.Split(content, "\n")
	texts = texts[:len(texts)-1]

	return texts[rand.Intn(len(texts))], nil
}

// MessageType defines available commands
type MessageType string

const (
	newPlayer MessageType = "new-player"
)

// Parser returns message type and message itself from a connection
func Parser(data string) (MessageType, string) {
	s := strings.Split(data[:len(data)-1], ":")
	return MessageType(s[0]), strings.Join(s[1:], ":")
}

// Compose creates a data string for connections
func Compose(t MessageType, msg string) []byte {
	return []byte(string(t) + ":" + msg + "\n")
}
