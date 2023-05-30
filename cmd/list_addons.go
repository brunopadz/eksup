package cmd

import (
	"context"
	awsClient "eksup/pkg/provider/aws"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/spf13/cobra"
	"os"
)

var addonsCmd = &cobra.Command{
	Use:     "addons",
	Aliases: []string{"a"},
	Short:   "List current add-ons versions and which can be upgraded",
	RunE:    listAddons,
}

var (
	version bool
)

func init() {
	listCmd.AddCommand(addonsCmd)

	addonsCmd.Flags().BoolVarP(&version, "check", "c", false, "Check for updates")
}

func listAddons(cmd *cobra.Command, args []string) error {
	s := awsClient.NewEksClient()

	c, err := awsClient.GetClusters(s)
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
			clusterInput := eks.DescribeClusterInput{
				Name: &n,
			}

			clusterGet, err := s.DescribeCluster(context.TODO(), &clusterInput)
			if err != nil {
				fmt.Println("Couldn't describe current cluster")
			}

			fmt.Println("Listing add-ons for cluster:", n)

			addonsInput := eks.ListAddonsInput{
				ClusterName: &n,
			}

			addonsList, err := s.ListAddons(context.TODO(), &addonsInput)
			if err != nil {
				fmt.Println("Couldn't list add-ons for clusters. Error:", err)
			}

			for _, v := range addonsList.Addons {
				i := eks.DescribeAddonInput{
					ClusterName: &n,
					AddonName:   &v,
				}

				d, err := s.DescribeAddon(context.TODO(), &i)
				if err != nil {
					fmt.Println("Couldn't list add-ons for clusters. Error:", err)
				}

				iu := eks.DescribeAddonVersionsInput{
					AddonName:         &v,
					KubernetesVersion: clusterGet.Cluster.Version,
					MaxResults:        nil,
					NextToken:         nil,
				}
				u, err := s.DescribeAddonVersions(context.TODO(), &iu)
				if err != nil {
					fmt.Println("Error while trying to find versions. Error:", err)
					os.Exit(1)
				}

				for _, w := range u.Addons {
					for k1, v1 := range w.AddonVersions {
						if k1 == 0 {
							if aws.ToString(d.Addon.AddonVersion) == aws.ToString(v1.AddonVersion) {
								fmt.Println(aws.ToString(w.AddonName), "\t...\t", aws.ToString(d.Addon.AddonVersion), "Already on latest version.")
							} else {
								fmt.Println(aws.ToString(w.AddonName), "\t...\t", aws.ToString(d.Addon.AddonVersion), "↑", aws.ToString(v1.AddonVersion))
							}
						}
					}
				}
			}
		}
	}

	return err
}
