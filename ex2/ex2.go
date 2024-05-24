package main

import "fmt"

func main() {
	fmt.Println(decodePassword("LLRR=")) //should 210122
	fmt.Println(decodePassword("==RLL")) //should 000210
	fmt.Println(decodePassword("=LLRR")) //should 221012
	fmt.Println(decodePassword("RRL=R")) //should 012001
}

func decodePassword(s string) string {
	// ans := ""
	// password := [len(s)+1]int
	// for i := 0; i<len(s); i++{
	// 	if(s[i] == 'L'){
			
	// 	}
	// }

	return ""
}


//R
//0,1

//RR
//0,1,2

//RRR
//0,1,2,3


//L
//1,0

//LL
//2,1,0

//LR
//L -> p[0] = p[1]+1
//R -> p[1] = p[2]-1
//p[2] 0
// 0 -> -1 ->0
// 1 0 1

//LRR
//L -> p[0] = p[1]+1
//R -> p[1] = p[2]-1
//R -> p[2] = p[3]-1
//p[3] 0
// 0 -> -1 -> -2 -> -1
// 2, 1, 0, 1
// 1 0 1 2

//=R
//001

//R=
//011

//==RL
//0 -> 1 -> 0 -> 0 -> 0
//00010

//==RR
//0 -> -1 -> -2 -> -2 -> -2
//00012

