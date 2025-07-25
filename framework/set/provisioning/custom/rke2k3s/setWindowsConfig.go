package rke2k3s

import (
	"os"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/rancher/tfp-automation/config"
	"github.com/rancher/tfp-automation/framework/set/provisioning/custom/nullresource"
	"github.com/sirupsen/logrus"
)

// SetCustomRKE2Windows is a function that will set the custom RKE2 cluster configurations in the main.tf file.
func SetCustomRKE2Windows(terraformConfig *config.TerraformConfig, terratestConfig *config.TerratestConfig, configMap []map[string]any,
	newFile *hclwrite.File, rootBody *hclwrite.Body, file *os.File) (*hclwrite.File, *os.File, error) {
	nullresource.CustomWindowsNullResource(rootBody, terraformConfig, terraformConfig.ResourcePrefix)
	rootBody.AppendNewline()

	_, err := file.Write(newFile.Bytes())
	if err != nil {
		logrus.Infof("Failed to write custom Windows RKE2 configurations to main.tf file. Error: %v", err)
		return nil, nil, err
	}

	return newFile, file, nil
}
