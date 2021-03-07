package cmd

import (
	"github.com/kenlyau/cilantro/core"
)

var CommandItems = []string{
	"DescribeInstances",
	"DescribeVpcs",
	"DescribeSecurityGroups",
	"CreateInstance",
	"Exit",
}
var client *core.Client

func Execute(regionID string, accessKeyID string, accessSecret string) {
	client = core.NewClient(regionID, accessKeyID, accessSecret)
	Select()
}
