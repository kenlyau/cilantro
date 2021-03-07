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
