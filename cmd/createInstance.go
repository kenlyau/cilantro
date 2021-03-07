package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

func ExecuteCreateInstance() {
	numberValidate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid Number")
		}
		return nil
	}
	vpcPrompt := promptui.Prompt{
		Label: "vpc id",
	}
	vpcID, err := vpcPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	securityGroupPrompt := promptui.Prompt{
		Label: "security group id",
	}
	securityGroupID, err := securityGroupPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	durationPrompt := promptui.Prompt{
		Label:    "duration time(hour)",
		Validate: numberValidate,
	}
	duration, err := durationPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	fmt.Println(vpcID, securityGroupID, duration)
	Select()
}
