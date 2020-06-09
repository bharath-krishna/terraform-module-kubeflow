package test

import (
	"strings"
	"testing"
	"time"

	//"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func getDefaultTerraformOptions(t *testing.T) (string, *terraform.Options, error) {

	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, "..", ".")

	random_id := strings.ToLower(random.UniqueId())

	terraformOptions := &terraform.Options{
		TerraformDir:       tempTestFolder,
		Vars:               map[string]interface{}{},
		MaxRetries:         5,
		TimeBetweenRetries: 5 * time.Minute,
		NoColor:            true,
		Logger:             logger.TestingT,
	}

	terraformOptions.Vars["install_istio"] = false

	return random_id, terraformOptions, nil
}

func TestApplyAndDestroyWithDefaultValues(t *testing.T) {
	t.Parallel()

	_, options, err := getDefaultTerraformOptions(t)
	assert.NoError(t, err)

	options.Vars["cert_manager_namespace"] = "cert-manager-etjllj"
	options.Vars["istio_operator_namespace"] = "istio-operator-etjllj"
	options.Vars["istio_namespace"] = "istio-system-etjllj"

	/*	k8sOptions := k8s.NewKubectlOptions("", "", "default")
		k8s.CreateNamespace(t, k8sOptions, namespace)
		// website::tag::5::Make sure to delete the namespace at the end of the test
		defer k8s.DeleteNamespace(t, k8sOptions, namespace)
	*/
	defer terraform.Destroy(t, options)
	_, err = terraform.InitAndApplyE(t, options)
	assert.NoError(t, err)
}
