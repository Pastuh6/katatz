package main

import ("fmt"
		"strconv"
		"strings"
	)

func proverka(n string) int {
	Rim := [...]string {"I", "V", "X", "L", "C"}
	Arab := [...]string {"1","2","3","4","5","6","7","8","9"}
	for a := range Rim {
		if n[0:1] == Rim[a] {
			return 2
		} 
	}//////////////////////////////////////////////////////////////Proverka
	for a := range Arab {
		if n[0:1] == Arab[a] {
			return 1
		}
	}
	uslovie ()
	return 0
}
func rome(r string) int {
	rimsk := map[string]int{"I": 1, "V": 5, "X": 10}
	arabika := 0
	for i := range r {
		if i <len(r)-1 && rimsk[r[i:i+1]] < rimsk[r[i+1:i+2]]  {//////////////na arabsk
			arabika -= rimsk[r[i:i+1]]
		} else {
			arabika += rimsk[r[i:i+1]]
		}
	}
	return arabika
}
func arabs(a int) string {
	rimk := map[int]string{100: "C", 90: "XC", 50: "L", 40:"XL", 10:"X", 9:"IX", 5:"V", 4:"IV", 1:"I"}//////////////na rimsk
	mk := [...]int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	ro := ""
	o := 0
	for o < 9 {
		if a - mk[o] >= 0 {
			a -= mk[o]
			ro += rimk[mk[o]]
		} else {
			o++
		}
	}
	return ro
}
func uslovie () string {
	return fmt.Sprintf("Ошибка")
}
func plus(lups string) string {
	var x,y [2]int
	nums := strings.Split(lups, "+")
	for qoq, san := range nums {
		switch proverka(san) {
		case 1: 
			v,err := strconv.Atoi(san)////////////////////plus operatcia
			if err != nil{
			}
			x[qoq]=v
			y[qoq]=proverka(san)
		case 2:
			x[qoq]=rome(san)
			y[qoq]=proverka(san)
		}
	}
	if x[0] > 10 || x[0] < 1 || x[1] > 10 || x[1] < 1 || y[0] != y[1] {
		return uslovie()
	} else if y[0] == 1 {
		return fmt.Sprintf(" %d + %d = %d", x[0], x[1], x[0]+x[1])
	} else {
		return fmt.Sprintf(" %s + %s = %s", arabs(x[0]), arabs(x[1]), arabs(x[0]+x[1]))
	}
}
func minus(sunim string) string {
	var x,y [2]int
	nums := strings.Split(sunim, "-")
	for qoq, san := range nums {
		switch proverka(san) {
		case 1: 
			v,err := strconv.Atoi(san)//////////////minus operacia
			if err != nil{
			}
			x[qoq]=v
			y[qoq]=proverka(san)
		case 2:
			x[qoq]=rome(san)
			y[qoq]=proverka(san)
		}
	}
	if x[0] > 10 || x[0] < 1 || x[1] > 10 || x[1] < 1 || y[0] != y[1] {
		return uslovie()
	} else if y[0] == 1 {
		return fmt.Sprintf(" %d - %d = %d", x[0], x[1], x[0]-x[1])
	} else {
		return fmt.Sprintf(" %s - %s = %s", arabs(x[0]), arabs(x[1]), arabs(x[0]-x[1]))
	}
}
func umnojit(umn string) string {
	var x,y [2]int
	nums := strings.Split(umn, "*")
	for qoq, san := range nums {
		switch proverka(san) {
		case 1: 
			v,err := strconv.Atoi(san)//////////////umojenie operacia
			if err != nil{
			}
			x[qoq]=v
			y[qoq]=proverka(san)
		case 2:
			x[qoq]=rome(san)
			y[qoq]=proverka(san)
		}
	}
	if x[0] > 10 || x[0] < 1 || x[1] > 10 || x[1] < 1 || y[0] != y[1] {
		return uslovie()
	} else if y[0] == 1 {
		return fmt.Sprintf(" %d * %d = %d", x[0], x[1], x[0]*x[1])
	} else {
		return fmt.Sprintf(" %s * %s = %s", arabs(x[0]), arabs(x[1]), arabs(x[0]*x[1]))
	}
}
func razdel(del string) string {
	var x,y [2]int
	nums := strings.Split(del, "/")
	for qoq, san := range nums {
		switch proverka(san) {
		case 1: 
			v,err := strconv.Atoi(san)///////////////delenei operazia
			if err != nil{
			}
			x[qoq]=v
			y[qoq]=proverka(san)
		case 2:
			x[qoq]=rome(san)
			y[qoq]=proverka(san)
		}
	}
	if x[0] > 10 || x[0] < 1 || x[1] > 10 || x[1] < 1 || y[0] != y[1] {
		return uslovie()
	} else if y[0] == 1 {
		return fmt.Sprintf(" %d + %d = %d", x[0], x[1], x[0]/x[1])
	} else {
		return fmt.Sprintf(" %s + %s = %s", arabs(x[0]), arabs(x[1]), arabs(x[0]/x[1]))
	}
}
func main() {
	fmt.Println("Введите значения:")
	var e string
	fmt.Scanf("%s", &e)
	if e[0] == '-' || proverka(e) == 0 {
		print(uslovie())
		return
	}
	for i := range e {
		switch e[i] {
		case '+':
			fmt.Println(plus(e))
		case '-':
			fmt.Println(minus(e))
		case '*':
			fmt.Println(umnojit(e))
		case '/':
			fmt.Println(razdel(e))
		default:
			continue
		}
	}
}
