package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
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

var (
	clusters []string
	version  bool
)

func init() {
	listCmd.AddCommand(addonsCmd)

	addonsCmd.Flags().BoolVarP(&version, "version", "v", false, "List versions")
}

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
	s := awsClient.NewEksClient()

	c, err := getClusters(s)
	if err != nil {
		fmt.Println("Couldn't get clusters. Error:", err)
	}

	if version != true {
		for _, n := range c {
			fmt.Println("Listing add-ons for cluster:", n)

			i := eks.ListAddonsInput{
				ClusterName: &n,
			}

			a, err := s.ListAddons(context.TODO(), &i)
			if err != nil {
				fmt.Println("Couldn't list add-ons for clusters. Error:", err)
			}

			for _, v := range a.Addons {
				fmt.Println(v)
			}

		}
	} else {
		for _, n := range c {
			fmt.Println("Listing add-ons for cluster:", n)

			i := eks.ListAddonsInput{
				ClusterName: &n,
			}

			a, err := s.ListAddons(context.TODO(), &i)
			if err != nil {
				fmt.Println("Couldn't list add-ons for clusters. Error:", err)
			}

			for _, v := range a.Addons {

				i := eks.DescribeAddonInput{
					ClusterName: &n,
					AddonName:   &v,
				}

				d, err := s.DescribeAddon(context.TODO(), &i)
				if err != nil {
					fmt.Println("Couldn't list add-ons for clusters. Error:", err)
				}

				fmt.Println(aws.ToString(d.Addon.AddonName), "\t...\t", aws.ToString(d.Addon.AddonVersion))

			}

		}
	}

	return err
}
