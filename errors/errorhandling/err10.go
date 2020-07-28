package main
import (e
	"fmt"
	"errors"
)
func Divide (a int,b int)(int,error) {
	if b==0 {
		return 0,errors.New("cant divide by 0")
	}elsec{
		return a/b,nil
	}
	}
	func main() {
		if result,err := Divide(4, 0); err != nil {
			fmt.Println("Error occured: ", err)
		} else {
			fmt.Println("4/0 is", result)
		}
	}