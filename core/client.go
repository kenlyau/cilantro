package core

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

type Client struct {
	regionID  string
	ecsClient *ecs.Client
}

func NewClient(regionID string, accessKeyID string, accessSecret string) *Client {
	ecsClient, err := ecs.NewClientWithAccessKey(regionID, accessKeyID, accessSecret)
	if err != nil {
		panic(err)
	}
	client := &Client{
		regionID:  regionID,
		ecsClient: ecsClient,
	}
	return client
}

func (c *Client) DescribeInstances() (*ecs.DescribeInstancesResponse, error) {
	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"
	request.PageSize = "100"
	response, err := c.ecsClient.DescribeInstances(request)
	return response, err
}

func (c *Client) DescribeVpcs() (*ecs.DescribeVpcsResponse, error) {
	request := ecs.CreateDescribeVpcsRequest()
	request.Scheme = "https"
	response, err := c.ecsClient.DescribeVpcs(request)
	return response, err
}

func (c *Client) DescribeSecurityGroups() (*ecs.DescribeSecurityGroupsResponse, error) {
	request := ecs.CreateDescribeSecurityGroupsRequest()
	request.Scheme = "https"
	request.PageSize = "100"
	response, err := c.ecsClient.DescribeSecurityGroups(request)
	return response, err
}

func (c *Client) CreateInstance(instanceType string, vSwitchID string, securityGroupID string, releaseTime string) (*ecs.CreateInstanceResponse, error) {
	request := ecs.CreateCreateInstanceRequest()
	request.Scheme = "https"
	request.VSwitchId = vSwitchID
	request.SecurityGroupId = securityGroupID
	request.InstanceType = instanceType
	request.DataDisk = &[]ecs.CreateInstanceDataDisk{{}}
	request.Arn = &[]ecs.CreateInstanceArn{{}}
	request.Tag = &[]ecs.CreateInstanceTag{{}}
	request.InstanceChargeType = "PostPaid"
	request.SpotStrategy = "SpotAsPriceGo"
	request.ImageId = "debian_10_7_x64_20G_alibase_20210128.vhd"

	response, err := c.ecsClient.CreateInstance(request)
	if err != nil {
		if releaseTime != "" {
			//自动释放
			_, err := c.AutoReleaseTime(response.InstanceId, releaseTime)
			if err != nil {
				return nil, err
			}
		}
	}

	return response, err
}

func (c *Client) AutoReleaseTime(instanceID string, releaseTime string) (*ecs.ModifyInstanceAutoReleaseTimeResponse, error) {
	request := ecs.CreateModifyInstanceAutoReleaseTimeRequest()
	request.Scheme = "https"
	request.InstanceId = instanceID
	request.AutoReleaseTime = releaseTime
	response, err := c.ecsClient.ModifyInstanceAutoReleaseTime(request)
	return response, err
}
