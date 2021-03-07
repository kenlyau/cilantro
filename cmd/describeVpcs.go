package cmd

import "fmt"

func ExecuteDescribeVpcs() {
	resp, err := client.DescribeVpcs()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	Select()
}
