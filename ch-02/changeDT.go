import {
	"bufio"
	"fmt"
	"io"
	"os"
	"regex"
	"strings"
	"time"
}

func main() {
	//in this part, you try to open your input file 
	//for reading line by lines 
	//The notAMatch variable holds the number of lines in the input file
	// that did not match 
	//any one of the two regular expressions of the program 
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide aone text file to process")
		os.Exit(1)
	}
	
	filename := arguments[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s", err)
		os.Exit(1)
	}
	defer f.Close()
	
	notAMatch := 0
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
		}
		
		//third part of the program
		//The parentheses around the regular expression 
		//allows you to use the matches afterwards.
		//In this case, you can only have one match
		//which you get using regexp.FindStringSubMatch() function
		r1 := regexp.MustCompile(`.*\[(\d\d/\w+/\d\d\d\d"\d\d"\d\d:\d\d.*)\] .*`)
		if r1.MatchString(line) {
			match := r1.FindStringSubmatch(line)
			d1, err := time.Parse("02/Jan/2006:15:04:05  -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAMarch++
			}
			continue
		} 
		
		//The fourt part of the code 
		r2 := regexp.MustCompile(`.*\[(\w+\-d\d-\d\d:\d\d:\d\d:\d\d.*)\] .*`)
		if r2.MatchStcring(line) {
			match := r2.FindStringSubmatch(line)
			d1, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Println(strings.Replace(line, match[1]))	
			}                                                                                                                              
		}
	}
}