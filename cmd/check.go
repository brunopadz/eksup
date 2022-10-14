package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/spf13/cobra"
	"log"
	"ueks/pkg/provider/aws"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	RunE:  test,
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func test(cmd *cobra.Command, _ []string) error {

	c, err := awsClient.NewClient("us-east-1")
	if err != nil {
		log.Fatalln(err)
	}
	n := eks.NewFromConfig(c)

	i := &eks.ListClustersInput{
		Include:    []string{"all"},
		MaxResults: nil,
		NextToken:  nil,
	}

	l, err := n.ListClusters(context.TODO(), i)
	if err != nil {
		log.Fatalln(err)
	}

	for k, v := range l.Clusters {
		fmt.Println(k)
		fmt.Println(v)
	}

	return err
}
