package main
import (
"fmt")

type student struct{
rollno int
name string
}
func main(){
 var s1 Student = Student{101,"vishnu"}
 var s2 Student  = Student{name:"priyan"}
 fmt.Println(s1)
 fmt.Println(s1.name)
 fmt.Println(s2)
 fmt.Printf("Name of the student with 101 roll no is %v\n",s1.name)
}
