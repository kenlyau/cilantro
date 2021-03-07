package cmd

import "fmt"

func ExecuteDescribeInstances() {
	resp, err := client.DescribeInstances()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	Select()
}
