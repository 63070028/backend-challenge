package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(decodePassword("LLRR=")) //should 210122
	fmt.Println(decodePassword("==RLL")) //should 000210
	fmt.Println(decodePassword("=LLRR")) //should 221012
	fmt.Println(decodePassword("RRL=R")) //should 012001

	fmt.Println(decodePassword("RRR"))   //should 0123
	fmt.Println(decodePassword("LLL"))   //should 3210
	fmt.Println(decodePassword("LLL=L")) //should 432110
	fmt.Println(decodePassword(""))      //should 0

}

func decodePassword(s string) string {

	pass := make([]int, len(s)+1)
	for i, v := range s {
		if v == '=' {
			pass[i+1] = pass[i]
		} else if v == 'L' {
			if pass[i] == 1 {
				pass[i+1] = 0
			} else {
				pass[i+1] = -1
			}
		} else if v == 'R' {
			pass[i+1] = pass[i] + 1
		}

		if pass[i+1] < 0 {
			pass[i+1]++
			for j := i; j >= 0; j-- {
				if s[j] == 'L' {
					if pass[j] > pass[j+1] {
						break
					} else {
						pass[j]++
					}
				} else if s[j] == 'R' {
					if pass[j] < pass[j+1] {
						break
					}
				} else if s[j] == '=' {
					if pass[j] == pass[j+1] {
						break
					} else {
						pass[j] = pass[j+1]
					}
				}

			}
		}
	}

	reuslt := ""
	for _, v := range pass {
		reuslt += strconv.Itoa(v)
	}

	return reuslt
}

/*

	= = R L L

	[]
	= -> 0,0

	[0,0]
	= -> 0,0

	[0,0,0]
	R -> 0,1

	[0,0,0,1]
	L -> 1,0

	[0,0,0,1,0]
	L -> 0,-1

	[0,0,0,1,0,-1]
	last < 0 -> all(pass) +1
	[1,1,1,2,1,0] incorrect
	-1 + 1 = 0

	(L) 0 > 0 ? -> No 0+1 -> 1
	(L) 1 > 1 ? -> No 1+1 -> 2
	(R) 0 < 2 ? -> Yes break

	[0,0,0,2,1,0]

*/

/*

	R R L = R
	pass[0,0,0,0,0,0]

	[]
	R -> 0,1

	[0,1]
	R -> 1,2

	[0,1,2]
	L -> 2,1

	[0,1,2,1] incorrect should -> [0,1,2,0]

*/
