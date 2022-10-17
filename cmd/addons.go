package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/spf13/cobra"
	"log"
	awsClient "ueks/pkg/provider/aws"
)

var addonsCmd = &cobra.Command{
	Use:     "addons",
	Aliases: []string{"a"},
	Short:   "Check current add-ons version and which can be upgraded",
	RunE:    listAddons,
}

func init() {
	listCmd.AddCommand(addonsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addonsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addonsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var clusters []string

func getClusters(clt *eks.Client) ([]string, error) {
	i := &eks.ListClustersInput{
		Include:    []string{"all"},
		MaxResults: nil,
		NextToken:  nil,
	}

	l, err := clt.ListClusters(context.TODO(), i)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range l.Clusters {
		clusters = append(clusters, v)
	}

	return clusters, err
}

func listAddons(cmd *cobra.Command, args []string) error {
	//cfg, err := awsClient.NewClient()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//clt := eks.NewFromConfig(cfg)

	clt := awsClient.NewEksClient()

	c, err := getClusters(clt)
	if err != nil {
		fmt.Println("eita")
	}

	for _, v := range c {
		fmt.Println(v)
	}

	//for _, v := range c {
	//	fmt.Println("Listing add-ons for", v)
	//
	//	i := eks.ListAddonsInput{
	//		ClusterName: &v,
	//	}
	//
	//	for _, a := range i {
	//
	//	}
	//
	//}

	return err
}
