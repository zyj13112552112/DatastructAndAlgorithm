package 困难

import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
	st := []rune{}
	for _,v:=range s{
		if v!=' '{
			st = append(st, v)
		}
	}//删掉空格
	fmt.Println(string(st))
	num := []int{}
	for i:=0;i< len(st);i++ {
		if st[i]=='-'{
			i++
			if st[i] == '('{
				l:=1
				i++
				now := i
				for ; l!=0; i++{
					if st[i] == '(' {
						l++
					}
					if st[i] == ')' {
						l--
					}
				}
				i-=1
				num = append(num, -1*calculate(string(st[now:i])) )
			}else{
				t,_ := strconv.Atoi(string(st[i]))

				num = append(num,-1*t)

			}

			continue
		}
		if st[i]>='0'&&st[i]<='9'{
			t,_ := strconv.Atoi(string(st[i]))
			if i==0 {
				num = append(num, t)
			}else{
				if st[i-1]>='0'&&st[i-1]<='9'{
					if num[len(num)-1]<0{
						num[len(num)-1] = num[len(num)-1]*10 - t
					}else{
						num[len(num)-1] = num[len(num)-1]*10 + t
					}
				}else{
					num = append(num, t)
				}
			}
		}
	}
	all := 0
	for _,v := range num{

		all+=v
	}
	return all
}