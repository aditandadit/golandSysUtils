package sysUtils

import (
	"fmt"
//	"os"
	"strings"
	"regexp"
	"reflect"
	"flag"
)

var s [3]string = [3]string{"1 b 3 1 a a b", "11 a 1 1 1 1 a a","-1 b 1 -4 a 1"}
type person struct {
	name string
	id int
}
func main() {
	match,_ := regexp.MatchString("adit", "hello adit");
	fmt.Println(match);
	parse, _ := regexp.Compile("[Aa]dit");
	fmt.Println(parse.MatchString("Aadit"));
	counts := make(map[string]int);
	for _,v :=  range s {
		line := strings.Split(v, " ");
		for _,i := range line {
			_, ok := counts[i]
			if ok {
				counts[i]++
			} else {
				counts[i]=0
			}
		}
	}
	for k,v := range counts {
		fmt.Println(k, " ", v);
	}
	parse = regexp.MustCompile("[Bb]")
	for i,_ := range s {
		parse.ReplaceAllString(s[i],"C" );
		fmt.Println(s[i]);
	}
	fmt.Println(" ");
	type t1 int
	type t2 int
	p1 := person{"adit",1};
	p2 := reflect.ValueOf(&p1).Elem();
	fmt.Println(p2.Field(0));

	minusO := flag.Bool("o", false, "o");
	minusC := flag.Int("c",0, "an int" );
	fmt.Println(minusC, minusO);
	for index, val := range flag.Args() {
		fmt.Println(index, " : ", val);
	}

}

func sort(x, y int ) (int, int) {
	if (x>y) {
		return x, y;
	} else {
		return y, x;
	}
}

