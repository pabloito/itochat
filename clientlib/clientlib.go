package lib

import (
	"strings"
)

type CommandType int

const(
    Send CommandType = iota
    Exit
    Invalid
)

type CommandDefinition struct{
	T CommandType
	P string
	S string
}

type Command struct{
    T CommandType
    Str string
    Msg string
}

func CheckCommand(str string) *Command {
    definitions := getCommandDefinitions()
    for _ , definition := range(definitions){
        isCommand, msg := getInnerString(str, definition.P, definition.S)
        if isCommand {
            return &Command{T: definition.T , Str: str, Msg: msg}    
        }	
    }
    return &Command{T: Invalid, Str: str, Msg: ""}
}

func getInnerString(str, prefix, suffix string) (bool, string) {
	correct := strings.HasSuffix(str,suffix) && strings.HasPrefix(str, prefix)
	if !correct{
		return false, ""
	}
	return true, strings.TrimLeft(strings.TrimRight(str,suffix),prefix)
}

func getCommandDefinitions() [2]CommandDefinition{
	definitions := [2]CommandDefinition{{T:Send, P:"Send '", S:"'"},{T:Exit, P:"Exit", S:""}}
	return definitions
}