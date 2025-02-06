package main

//One of them, which
//is called unicode.IsPrint(), can help you to identify the parts of a string that are
//printable using runes. This technique will be illustrated in the Go code of unicode.go,
//which will be presented in two parts. The first part of unicode.go is as follows:
import (
	"fmt"
	s "strings"
	"unicode"
)
//Another handy trick is that if you find yourself using the same function all the time, and
//you want to use something shorter instead, you can assign a variable name to that function
//and use that variable name instead
var f= fmt.Printf

//the second part of the code is as follows 
func main(){
	//see many functions that allow you to play with the case of a
	//string. Additionally, you can see that the strings.EqualFold() function allows you to
	//determine whether two strings are the same in spite of the differences in their letters
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))
	f("%s\n", s.Title("tHis will be A title!"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))
	
	//the third portion of the code 
	//The strings.Count() function counts the number of non-overlapping times the second
	//parameter appears in the string that is given as the first parameter. The
	//strings.HasPrefix() function returns true when the first parameter string begins with
	//the second parameter string, and false otherwise. Similarly, the strings.HasSuffix()
	//function returns true when the first parameter, which is a string, ends with the second
	//parameter, which is also a string, and false otherwise.
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix(("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))
	
	f("Index: %v\n", s.HasIndex("Mihalis", "ha"))
	f("Index: %v\n", s.HasIndex("Mihalis", "Ha"))
	f("Count: %v\n", s.Count("Mihalis", "i"))
	f("Count: %v\n", s.Count("Mihalis", "I"))
	f("Repeat: %s\n", s.Repeat("ab", 5))
	
	f("TrimSpace: %s\n", s.TrimSpace("\t This is line. \n"))
	f("TrimLeft: %s", s.TrimLeft("\t This is a\t . \n", "\n\t"))
	f("TrimRight: %s\n", s.TrimRight("\tThis is a\t line. \n", "\n\t"))
	
	//The fourth segment of the code
	//This code portion contains some pretty advanced and ingenious functions. The first handy
	//function is strings.Split(), which allows you to split the given string according to the
	//desired separator string – the strings.Split() function returns a string slice. Using ""
	//as the second parameter of strings.Split() will allow you to process a string character
	//by character.
	//The strings.Compare() function compares two strings lexicographically, and it may
	//return 0 if the two strings are identical, and -1 or +1 otherwise.
	//Lastly, the strings.Fields() function splits the string parameter using whitespace
	//characters as separators. The whitespace characters are defined in the unicode.IsSpace()
	//function.
 
	f("Compare : %v\n", s.Compare("Mihalis", "MIHALIS"))
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis"))
	f("Compare: %v\n", s.Compare("MIHALIS", "Mihalis"))
	
	f("Fields: %v\n", s.Fields("This is a string!"))
	f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))
	f("%s\n", s.Split("abcd efg", ""))
	
	//The last part of the code 
	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))
	
	lines := [] string{"Line 1", "Line 2", "Line3 "}
	f("Join: %s\n", s.Join(lines, "+++"))
	
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))
	
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction75))
	
}

