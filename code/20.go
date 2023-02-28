package 中等

import (
	//"fmt"
	"strings"
)
//有穷状态自动机
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if s=="-." {return false} //-0.0
    status := 1
	for _,v:=range s{
		switch status {
		case 0:return false
		case 1:
			if v == '+' || v == '-'  {
				status = 2
			}else if v == '.'{
				status = 11
			}else if Diagital(v) {
				status = 3
			}else {
				status = 0
			}
		case 2:
			if Diagital(v) {
				status = 3
			}else if v == '.'{
			    status = 7
			} else {
				status = 0
			}
		case 3:
			if Diagital(v) {
				status = 3
			}else if v == 'E' || v == 'e' {
				status = 4
			}else if v == '.'{
				status = 7
			}else {
				status = 0
			}
		case 4:
			if v == '+' ||v == '-' {
				status = 5
			}else if Diagital(v){
				status = 6
			}else {
				status = 0
			}
		case 5:
			if Diagital(v) {
				status = 6
			}else {
				status = 0
			}
		case 6:
			if Diagital(v) {
				status = 6
			}else {
				status = 0
			}
		case 7:
			if Diagital(v) {
				status = 8
			}else if v=='E'||v=='e'{
				status = 9
			}else {
				status = 0
			}
		case 8:
			if Diagital(v) {
				status = 8
			}else if v == 'E' || v == 'e' {
				status = 9
			}else {
				status = 0
			}
		case 9:
			if v =='+'||v=='-'{
				status = 10
			}else if Diagital(v){
				status = 12
			}else {
				status = 0
			}
		case 10:
			if Diagital(v) {
				status = 12
			}else {
				status = 0
			}
		case 11:
			if Diagital(v){
				status = 8
			}else {
				status = 0
			}
		case 12:
			if Diagital(v) {
				status = 12
			}else {
				status = 0
			}
		}
	}
	if status ==3 || status == 6 || status==7 ||status==8 ||status == 12{
		return true
	}
	return false
}
func Diagital(v int32)bool{
	if v >='0' && v <= '9' {
		return true
	}
	return false
}