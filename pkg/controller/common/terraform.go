// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"time"

	"github.com/gardener/gardener-extension-provider-alicloud/pkg/alicloud"
	"github.com/gardener/gardener-extension-provider-alicloud/pkg/imagevector"

	"github.com/gardener/gardener-extensions/pkg/terraformer"
	"github.com/gardener/gardener/pkg/logger"
	"k8s.io/client-go/rest"
)

const (
	TerraformVarAccessKeyID     = "TF_VAR_ACCESS_KEY_ID"
	TerraformVarAccessKeySecret = "TF_VAR_ACCESS_KEY_SECRET"
)

// NewTerraformer creates a new Terraformer.
func NewTerraformer(factory terraformer.Factory, config *rest.Config, purpose, namespace, name string) (terraformer.Terraformer, error) {
	tf, err := factory.NewForConfig(logger.NewLogger("info"), config, purpose, namespace, name, imagevector.TerraformerImage())
	if err != nil {
		return nil, err
	}

	return tf.
<<<<<<< HEAD
=======
		SetVariablesEnvironment(variablesEnvironment).
>>>>>>> implement migrate and restore functionality for infrasturucture
		SetTerminationGracePeriodSeconds(630).
		SetDeadlineCleaning(5 * time.Minute).
		SetDeadlinePod(15 * time.Minute), nil
}

// NewTerraformerWithAuth creates a new Terraformer and initializes it with the credentials.
func NewTerraformerWithAuth(factory terraformer.Factory, config *rest.Config, purpose, namespace, name string, credentials *alicloud.Credentials) (terraformer.Terraformer, error) {
	tf, err := NewTerraformer(factory, config, purpose, namespace, name)
	if err != nil {
		return nil, err
	}

	return tf.SetVariablesEnvironment(TerraformVariablesEnvironmentFromCredentials(credentials)), nil
}

// TerraformVariablesEnvironmentFromCredentials computes the Terraformer variables environment from the
// given ServiceAccount.
func TerraformVariablesEnvironmentFromCredentials(credentials *alicloud.Credentials) map[string]string {
	return map[string]string{
		TerraformVarAccessKeyID:     credentials.AccessKeyID,
		TerraformVarAccessKeySecret: credentials.AccessKeySecret,
	}
}
