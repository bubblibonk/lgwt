package iteration

import "strings"

func Repeat(letter string)string{
    var repeated strings.Builder
    for i:=0;i<5;i++{
        repeated.WriteString(letter);
    }
    return repeated.String();
}
