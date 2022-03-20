package fetchers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/eks"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type EKSProvider struct {
	client *eks.Client
}

func NewEksProvider(cfg aws.Config) *EKSProvider {
	svc := eks.NewFromConfig(cfg)
	return &EKSProvider{
		client: svc,
	}
}

func (provider EKSProvider) DescribeCluster(ctx context.Context, clusterName string) (*eks.DescribeClusterOutput, error) {
	input := &eks.DescribeClusterInput{
		Name: &clusterName,
	}
	response, err := provider.client.DescribeCluster(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to describe cluster %s from eks , error - %w", clusterName, err)
	}

	return response, err
}
