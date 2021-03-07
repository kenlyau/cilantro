package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func Select() {
	prompt := promptui.Select{
		Label: "Select Command",
		Items: CommandItems,
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err)
	}
	switch result {
	case "DescribeInstances":
		ExecuteDescribeInstances()
	case "DescribeVpcs":
		ExecuteDescribeVpcs()
	case "DescribeSecurityGroups":
		ExecuteDescribeSecurityGroups()
	case "CreateInstance":
		ExecuteCreateInstance()
	default:
		fmt.Println("ok")
	}
}
