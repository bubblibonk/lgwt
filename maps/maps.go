package main

import (
	"fmt"
)
type DictionaryErr string;
var ErrNotFound = DictionaryErr("word not found")
var ErrWordExists = DictionaryErr("word exists")

type Dictionary map[string]string;
func (e DictionaryErr) Error() string{
    return string(e);
}
func (d Dictionary) Search(word string) (string,error) {
    definition,ok:=d[word]
    if !ok{
        return "",ErrNotFound
    }
    return definition,nil
}

func (d Dictionary) Add(word,definition string) error{
    _,err:=d.Search(word)
    switch err{
    case ErrNotFound:
        d[word] = definition
    case ErrWordExists:
        return nil
    default:
        return err
    
    }    
    return nil
}
func main(){
    d:=Dictionary{"Test":"This is just a word"}
    d.Add("bonk","go to jail")
    fmt.Println(d.Search("Test"))
    fmt.Println(d)
}   

