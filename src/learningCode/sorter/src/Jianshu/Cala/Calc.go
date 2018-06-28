/*
对行命令s输入框架cli的使用
*/
package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func main() {

	var help = func() { //使用函数作为类型定义
		fmt.Println("Usage for calc tool.")
		fmt.Println("=================================")
		fmt.Println("add 1 2, return 3")
		fmt.Println("sub 1 2, return -1")
		fmt.Println("mul 1 2, return 2")
		fmt.Println("sqrt 2, return 1.4142135623730951")
	}

	app := cli.NewApp()
	app.Name = "calc with Jeremy learning"
	app.Usage = "calc tool operate by Jeremy"
	app.Version = "0.1.0"
	/*app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
	return nil
	}*/

	//
	//定义三个参数 operation,numberone,numbertwo
	//

	app.Flags = []cli.Flag{ //一个cli.Flag的切片
		cli.StringFlag{
			Name:  "operation,o",
			Value: "add",
			Usage: "calc operation",
		},
		cli.Float64Flag{
			Name:  "numberone,n1",
			Value: 0,
			Usage: "number one for operation",
		},
		cli.Float64Flag{
			Name:  "numbertwo,n2",
			Value: 0,
			Usage: "number two for operation",
		},
	}

	//重写Action方法
	app.Action = func(c *cli.Context) error {
		operation := c.String("operation")
		numberone := c.Float64("numberone")
		numbertwo := c.Float64("numbertwo")

		if operation == "add" {
			rt := numberone + numbertwo
			fmt.Println("Result", rt)

		} else if operation == "sub" {
			rt := numberone - numbertwo
			fmt.Println("Result", rt)

		} else if operation == "mul" {
			rt := numberone * numbertwo
			fmt.Println("Result", rt)
		} else if operation == "sqrt" {
			rt := math.Sqrt(numberone)
			fmt.Println("Result", rt)
		} else {
			help() //函数类型变量的使用
		}
		return nil
	}
	app.Run(os.Args)
}

func Request(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return " ", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}
