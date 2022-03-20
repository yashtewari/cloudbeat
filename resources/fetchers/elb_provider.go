package fetchers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
)

type ELBProvider struct {
	client *elasticloadbalancing.Client
}

func NewELBProvider(cfg aws.Config) *ELBProvider {
	svc := elasticloadbalancing.NewFromConfig(cfg)
	return &ELBProvider{
		client: svc,
	}
}

// DescribeLoadBalancer method will return up to 400 results
// If we will ever want to increase this number, DescribeLoadBalancers support paginated requests
func (provider ELBProvider) DescribeLoadBalancer(ctx context.Context, balancersNames []string) ([]types.LoadBalancerDescription, error) {
	input := &elasticloadbalancing.DescribeLoadBalancersInput{
		LoadBalancerNames: balancersNames,
	}

	response, err := provider.client.DescribeLoadBalancers(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to describe load balancers %s from elb, error - %w", balancersNames, err)
	}

	return response.LoadBalancerDescriptions, err
}
