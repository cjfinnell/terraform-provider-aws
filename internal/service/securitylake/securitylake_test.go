// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitylake_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccSecurityLake_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"AWSLogSource": {
			"basic":       testAccAWSLogSource_basic,
			"disappears":  testAccAWSLogSource_disappears,
			"multiRegion": testAccAWSLogSource_multiRegion,
		},
		"DataLake": {
			"basic":           testAccDataLake_basic,
			"disappears":      testAccDataLake_disappears,
			"tags":            testAccDataLake_tags,
			"lifecycle":       testAccDataLake_lifeCycle,
			"lifecycleUpdate": testAccDataLake_lifeCycleUpdate,
			"replication":     testAccDataLake_replication,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
