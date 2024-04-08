package e2e

import (
	"context"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/openshift/rosa/tests/ci/labels"
	"github.com/openshift/rosa/tests/utils/common"
	"github.com/openshift/rosa/tests/utils/config"
	"github.com/openshift/rosa/tests/utils/exec/rosacli"
)

var _ = Describe("Network verifier",
	labels.Day2,
	labels.FeatureNetworkVerifier,
	func() {
		defer GinkgoRecover()

		var (
			clusterID      string
			rosaClient     *rosacli.Client
			networkService rosacli.NetworkVerifierService
			clusterService rosacli.ClusterService
		)

		BeforeEach(func() {
			By("Get the cluster")
			clusterID = config.GetClusterID()
			Expect(clusterID).ToNot(Equal(""), "ClusterID is required. Please export CLUSTER_ID")

			By("Init the client")
			rosaClient = rosacli.NewClient()
			networkService = rosaClient.NetworkVerifier
			clusterService = rosaClient.Cluster
		})

		//OCP-64917 - [OCM-152] Verify network via the rosa cli
		It("can verify network - [id:64917]",
			labels.High,
			labels.MigrationToVerify,
			labels.Exclude,
			func() {
				By("Get cluster description")
				output, err := clusterService.DescribeCluster(clusterID)
				Expect(err).To(BeNil())
				clusterDetail, err := clusterService.ReflectClusterDescription(output)
				Expect(err).To(BeNil())

				By("Check if non BYO VPC cluster")
				isBYOVPC, err := clusterService.IsBYOVPCCluster(clusterID)
				Expect(err).To(BeNil())
				if !isBYOVPC {
					Skip("It does't support the verification for non byo vpc cluster - cannot run this test")
				}

				By("Run network verifier vith clusterID")
				output, err = networkService.CreateNetworkVerifierWithCluster(clusterID)
				Expect(err).ToNot(HaveOccurred())
				Expect(output.String()).To(ContainSubstring("Run the following command to wait for verification to all subnets to complete:\n" + "rosa verify network --watch --status-only"))

				By("Get the cluster subnets")
				var subnetsNetworkInfo string
				for _, networkLine := range clusterDetail.Network {
					if value, containsKey := networkLine["Subnets"]; containsKey {
						subnetsNetworkInfo = value
						break
					}
				}
				subnets := strings.Replace(subnetsNetworkInfo, " ", "", -1)
				region := clusterDetail.Region
				installerRoleArn := clusterDetail.STSRoleArn

				By("Check the network verifier status")
				err = wait.PollUntilContextTimeout(context.Background(), 20*time.Second, 200*time.Second, false, func(context.Context) (bool, error) {
					output, err = networkService.GetNetworkVerifierStatus(
						"--region", region,
						"--subnet-ids", subnets,
					)
					if strings.Contains(output.String(), "pending") {
						return false, err
					}
					return true, err
				})
				common.AssertWaitPollNoErr(err, "Network verification result are not ready after 200")

				output, err = networkService.GetNetworkVerifierStatus(
					"--region", region,
					"--subnet-ids", subnets,
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(output.String()).ToNot(ContainSubstring("failed"))

				By("Check the network verifier with tags attributes")
				output, err = networkService.CreateNetworkVerifierWithCluster(clusterID,
					"--tags", "t1:v1")
				Expect(err).ToNot(HaveOccurred())

				By("Check the network verifier status")
				err = wait.PollUntilContextTimeout(context.Background(), 20*time.Second, 200*time.Second, false, func(context.Context) (bool, error) {
					output, err = networkService.GetNetworkVerifierStatus(
						"--region", region,
						"--subnet-ids", subnets,
					)
					if strings.Contains(output.String(), "pending") {
						return false, err
					}
					return true, err
				})
				common.AssertWaitPollNoErr(err, "Network verification result are not ready after 200")

				output, err = networkService.GetNetworkVerifierStatus(
					"--region", region,
					"--subnet-ids", subnets,
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(output.String()).ToNot(ContainSubstring("failed"))

				By("Run network verifier vith subnet id")
				if installerRoleArn == "" {
					Skip("It does't support the verification with subnets for non STS cluster - cannot run this test")
				}
				output, err = networkService.CreateNetworkVerifierWithSubnets(
					"--region", region,
					"--subnet-ids", subnets,
					"--role-arn", installerRoleArn,
					"--tags", "t2:v2",
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(output.String()).To(ContainSubstring("Run the following command to wait for verification to all subnets to complete:\n" + "rosa verify network --watch --status-only"))
				Expect(output.String()).To(ContainSubstring("pending"))

				By("Check the network verifier with hosted-cp attributes")
				output, err = networkService.CreateNetworkVerifierWithSubnets(
					"--region", region,
					"--subnet-ids", subnets,
					"--role-arn", installerRoleArn,
					"--hosted-cp",
				)
				Expect(err).ToNot(HaveOccurred())

			})

	})