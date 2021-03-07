package cmd

import "fmt"

func ExecuteDescribeSecurityGroups() {
	resp, err := client.DescribeSecurityGroups()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	Select()
}
