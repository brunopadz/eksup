package cmd

import (
	"context"
	"eksup/pkg/getver"
	awsClient "eksup/pkg/provider/aws"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/spf13/cobra"
)

var clustersCmd = &cobra.Command{
	Use:     "clusters",
	Aliases: []string{"a"},
	Short:   "List current add-ons versions and which can be upgraded",
	RunE:    listClusters,
}

func init() {
	listCmd.AddCommand(clustersCmd)

	clustersCmd.Flags().BoolVarP(&version, "check", "c", false, "Check for updates")
}

func listClusters(cmd *cobra.Command, args []string) error {
	s := awsClient.NewEksClient()

	c, err := awsClient.GetClusters(s)
	if err != nil {
		fmt.Println("Couldn't get clusters. Error:", err)
	}

	fmt.Println("List of clusters:")

	if version != true {
		for _, cluster := range c {
			i := eks.DescribeClusterInput{
				Name: &cluster,
			}

			a, err := s.DescribeCluster(context.TODO(), &i)
			if err != nil {
				fmt.Println("error")
			}

			fmt.Println(aws.ToString(a.Cluster.Name))
		}
	} else {
		versions, err := getver.GetVersion()
		if err != nil {
			fmt.Println("Couldn't get versions. Error:", err)
		}

		for _, cluster := range c {
			i := eks.DescribeClusterInput{
				Name: &cluster,
			}

			a, err := s.DescribeCluster(context.TODO(), &i)
			if err != nil {
				fmt.Println("error")
			}

			if aws.ToString(a.Cluster.Version) != versions[0] {
				fmt.Println(aws.ToString(a.Cluster.Name), "is running version:", styleRed.Render(aws.ToString(a.Cluster.Version)), "and can be upgraded to version:", styleBlue.Render(versions[0]))
			} else {
				fmt.Println(aws.ToString(a.Cluster.Name), "is running version:", styleGreen.Render(aws.ToString(a.Cluster.Version)), "and is up to date")
			}
		}
	}

	return err
}
