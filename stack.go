
package main

import "fmt"

//type student interface{
//	getName() string
//}
//
//type studentData struct {
//	name string
//}
//func (data studentData)getName() string{
//	return data.name
//}
//func main() {
// var s student
// var userdata string
// fmt.Scan(&userdata)
// s = studentData{userdata}
// fmt.Println(s.getName())
// fmt.Println(s)
//////}
type stack struct{
	stackArr []int
}
type stackMethods interface{
	push(a int) bool
	pop()
	top() int
	isEmpty() bool
}

func (s *stack)push(a int) bool{
	s.stackArr = append(s.stackArr, a)
	return true
}
func (s *stack)pop(){
	if len(s.stackArr) == 0{
		return
	}
	s.stackArr = s.stackArr[:len(s.stackArr) - 1]
}
func (s *stack)top() int{
	if len(s.stackArr) == 0{
		return -1
	}
	return s.stackArr[len(s.stackArr) - 1]
}

func (s *stack)isEmpty() bool{
	if len(s.stackArr) == 0{
		return true
	}
	return false
}
func main(){
	var st stackMethods
	st = &stack{}
	st.push(13)
	st.push(14)
	st.push(15)
	for{
		if st.isEmpty(){
			break
		}
		fmt.Println(st.top())
		st.pop()
	}
}



