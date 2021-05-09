package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)


func user_option() (int,int){
	var lev, num int
	fmt.Println("*** Welcome to factorising practice ***")
	fmt.Println("- Please input the number of questions you would like to pratice [1~20, default is 5]:")
	fmt.Scan(&num)
	if num<1 || num>20 {
		num = 5
	}
	fmt.Println("- Please choose your preferred level:\n" +
		"- Enter 0 to exit\n"+
		"- Enter 1 for junior level, the coefficient of X^2 is always 1\n" +
		"- Enter 2 for senior level, the coefficient of X^2 is between 2~6")
	fmt.Scan(&lev)

	return lev, num
}

// numGen generate a non-zero number between -9~9
//func numGen(bond int, shift int) int{
//	a :=0
//	for a==0 {
//		a= rand.Intn(bond)-shift
//	}
//	return a
//}
func coeffGen(bond int, shift int ) (int,int){
	var a, b int
	a,b=0,0

	for a==0 {
		a= rand.Intn(bond)-shift
	}
	for b==0 {
		b= rand.Intn(bond)-shift
	}

	return a,b
}
func questionGen(index int, a int, b int, coeff1 int, coeff2 int) string{
	var ques string
	if coeff1*coeff2==1{
		ques = "x^2"+signItx(coeff1*b+coeff2*a)+"X"+signIt(a*b)
		fmt.Printf("Question %d:\n   %s\n", index, ques)
	} else if coeff1*b+coeff2*a == 0 {
		ques = strconv.Itoa(coeff1*coeff2)+"x^2"+signIt(a*b)
		fmt.Printf("Question %d:\n   %s\n", index, ques)
	} else {
		ques = strconv.Itoa(coeff1*coeff2)+"x^2"+signItx(coeff1*b+coeff2*a)+"X"+signIt(a*b)
		fmt.Printf("Question %d:\n   %s\n", index, ques)
	}
	return ques
}

func signItx(m int) string{
	var mString string
	mString=strconv.Itoa(m)
	if m>1 {
		mString="+"+mString
	} else if m==1 {
		mString="+"
	} else if m==0 {
		mString=""
	}else if m==-1 {
		mString="-"
	}
	return mString
}
func signIt(m int) string{
	var mString string
	mString=strconv.Itoa(m)
	if m>0 {
		mString="+"+mString
	}
	return mString
}
//func init() {
//	rand.Seed(time.Now().UTC().UnixNano())
//}
func main() {
	var a, b, coeff1, coeff2 int
	var question, answer, solution [20]string

	rand.Seed(time.Now().UnixNano())

	// initialize tabwriter
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

//	start := time.Now()
CHOOSE:
	level, num := user_option()

	for i := 0; i < num; i++{
		switch level {
		case 0:
			fmt.Println("See you next time")
			return
		case 1:
			coeff1, coeff2 = 1, 1
		case 2:
			coeff1, coeff2 = coeffGen(4,-2)
		default:
			fmt.Println("Unknown option. Please choose again. Enter 0 to exit.")
			goto CHOOSE
		}
		a,b = coeffGen(19,9)
		question[i]= questionGen(i+1, a, b, coeff1, coeff2)
		if level==1 {
			solution[i]="(x"+signIt(a)+")"+"(x"+signIt(b)+")"
		} else {
			solution[i]="("+strconv.Itoa(coeff1)+"x"+signIt(a)+")"+"("+strconv.Itoa(coeff2)+"x"+signIt(b)+")"
		}
		fmt.Println("Your answer:")
		fmt.Scan(&answer[i])
	}
	fmt.Println("Well done! Please have of your answers and the solutions:")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "Index", "  Question  ", "   Solution  ", "  Your_answer  ")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "-----", "------------ ", " -------------", " -------------")
	for j := 0; j < num; j++{
		fmt.Fprintf(w, "\n %d\t%s\t%s\t%s\t", j+1, question[j],solution[j], answer[j])
	}
}



