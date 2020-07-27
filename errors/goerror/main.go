package main
import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/util"
	"os"
)
func main() {
	in,err := os.Open(*fileIn)
	if err != nil {
		if os.IsNotExist(err)
		fmt.Println("....")
		fmt.Println("file open error", err)
		os.exit(1)
	}
	data,err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Println("data read error" , err)
		os.Exit(1)
	}
	var v interface{}
	if err := json.Unmarshal(data,&v); err != nil {
		fmt.Println("json error", err)
		os.Exit(1)
	}
	out , err := os.Create(*fileOut)
	if err != nil {
		fmt.Println("file create error")
		os.Exit(1)
	}
	if _,err = io.Copy(out,bytes.NewReader(data)); err != nil {
		fmt.Println("data write error",err)
		os.Exit(1)
	}
	if err := out.Close(); err != nil {
		fmt.Println("warn", err)
	}
	if err := out.Close(); err != nil {
		fmt.Println("warn", err)
	

}
var (
	fileIn =  flag.String("in"," " ,"specify input file path")
	fileOut = flag.String("out" "","specify output file path")
)
func init(){
	flag.Parse()
}