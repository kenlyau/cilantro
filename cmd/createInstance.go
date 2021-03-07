package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"time"

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
	vSwitchPrompt := promptui.Prompt{
		Label: "vswitch id",
	}
	vSwitchID, err := vSwitchPrompt.Run()
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
		fmt.Printf("duration failed %v\n", err)
	}

	durationTime, err := time.ParseDuration(duration + "h")
	if err != nil {
		fmt.Printf("duration failed %v\n", err)
	}
	releaseTime := time.Now().Add(durationTime).UTC().Format(time.RFC3339)
	fmt.Println(vSwitchID, securityGroupID, releaseTime)
	_, err = client.CreateInstance("ecs.g5.large", vSwitchID, securityGroupID, releaseTime)
	if err != nil {
		fmt.Printf("create instance failed %v\n", err)
	}
	Select()
}
