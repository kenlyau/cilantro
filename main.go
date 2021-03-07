package main

import (
	"github.com/kenlyau/cilantro/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	accessKeyID := viper.GetString("ACCESS_KEY_ID")
	accessSecret := viper.GetString("ACCESS_KEY_SECRET")
	regionID := viper.GetString("REGION_ID")
	// client := core.NewClient(regionID, accessKeyID, accessSecret)
	// resp, err := client.DescribeSecurityGroups()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%#v", resp)
	cmd.Execute(regionID, accessKeyID, accessSecret)
}
