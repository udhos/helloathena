package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/athena"
)

const table = "textqldb.textqltable"
const outputBucket = "s3://bucket-name-here/"

func main() {

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		fmt.Printf("config error: %v\n", err)
		return
	}

	cfg.Region = endpoints.UsEast2RegionID

	client := athena.New(cfg)

	query := "select * from " + table

	resultConf := &athena.ResultConfiguration{
		OutputLocation: aws.String(outputBucket),
	}

	params := &athena.StartQueryExecutionInput{
		QueryString:         aws.String(query),
		ResultConfiguration: resultConf,
	}

	req := client.StartQueryExecutionRequest(params)

	resp, err := req.Send(context.TODO())
	if err != nil {
		fmt.Printf("query error: %v\n", err)
		return
	}

	fmt.Println(resp)
}
