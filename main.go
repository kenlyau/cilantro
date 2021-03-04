package main

import (
	"fmt"

	"github.com/denverdino/aliyungo/ecs"
	"github.com/spf13/viper"
)

var ACCESS_KEY_ID string
var ACCESS_KEY_SECRET string
var INSTANCE_TYPE string
var IMAGE_ID string
var SECURITY_GROUP_ID string
var V_SWITCH_ID string
var PASSWORD string
var REGION_ID string

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	ACCESS_KEY_ID = viper.GetString("ACCESS_KEY_ID")
	ACCESS_KEY_SECRET = viper.GetString("ACCESS_KEY_SECRET")
	INSTANCE_TYPE = viper.GetString("INSTANCE_TYPE")
	IMAGE_ID = viper.GetString("IMAGE_ID")
	SECURITY_GROUP_ID = viper.GetString("SECURITY_GROUP_ID")
	V_SWITCH_ID = viper.GetString("V_SWITCH_ID")
	PASSWORD = viper.GetString("PASSWORD")
	REGION_ID = viper.GetString("REGION_ID")

	client := ecs.NewClient(ACCESS_KEY_ID, ACCESS_KEY_SECRET)

	ecsInstance := &ECSInstance{
		Client:          client,
		Region:          REGION_ID,
		InstanceType:    INSTANCE_TYPE,
		VSwitchId:       V_SWITCH_ID,
		ImageId:         IMAGE_ID,
		SecurityGroupId: SECURITY_GROUP_ID,
		Password:        PASSWORD,
	}

	err = ecsInstance.Create("1h")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("create ecs instance success, [%s] \n", ecsInstance.ID)

}
