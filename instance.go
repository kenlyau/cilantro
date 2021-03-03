package main

import (
	"time"

	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/ecs"
)

type ECSInstance struct {
	Client          *ecs.Client
	ID              string
	Region          string
	InstanceType    string
	VSwitchId       string
	ImageId         string
	SecurityGroupId string
	Password        string
}

func (ins *ECSInstance) Create(duration string) error {
	args := &ecs.RunInstanceArgs{}
	if duration != "" {
		now := time.Now()
		d, err := time.ParseDuration(duration)
		if err != nil {
			return err
		}
		releaseTime := now.Add(d).UTC().Format(time.RFC3339)
		args.AutoReleaseTime = releaseTime
	}
	args.RegionId = common.Region(ins.Region)
	args.SecurityGroupId = ins.SecurityGroupId
	args.VSwitchId = ins.VSwitchId
	args.ImageId = ins.ImageId
	args.Password = ins.Password
	args.InstanceType = ins.InstanceType
	args.InstanceChargeType = "PostPaid"
	args.SpotStrategy = "SpotAsPriceGo"
	args.InternetChargeType = "PayByTraffic"
	args.InternetMaxBandwidthOut = 5
	idset, err := ins.Client.RunInstances(args)

	if err == nil {
		ins.ID = idset[0]
	}
	return err
}
