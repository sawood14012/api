// +build !ignore_autogenerated

/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.BannedUser":                schema_pkg_apis_toolchain_v1alpha1_BannedUser(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.BannedUserSpec":            schema_pkg_apis_toolchain_v1alpha1_BannedUserSpec(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecord":          schema_pkg_apis_toolchain_v1alpha1_MasterUserRecord(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecordSpec":      schema_pkg_apis_toolchain_v1alpha1_MasterUserRecordSpec(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecordStatus":    schema_pkg_apis_toolchain_v1alpha1_MasterUserRecordStatus(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSet":             schema_pkg_apis_toolchain_v1alpha1_NSTemplateSet(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec":         schema_pkg_apis_toolchain_v1alpha1_NSTemplateSetSpec(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetStatus":       schema_pkg_apis_toolchain_v1alpha1_NSTemplateSetStatus(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTier":            schema_pkg_apis_toolchain_v1alpha1_NSTemplateTier(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierSpec":        schema_pkg_apis_toolchain_v1alpha1_NSTemplateTierSpec(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierStatus":      schema_pkg_apis_toolchain_v1alpha1_NSTemplateTierStatus(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.RegistrationService":       schema_pkg_apis_toolchain_v1alpha1_RegistrationService(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.RegistrationServiceSpec":   schema_pkg_apis_toolchain_v1alpha1_RegistrationServiceSpec(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.RegistrationServiceStatus": schema_pkg_apis_toolchain_v1alpha1_RegistrationServiceStatus(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccount":               schema_pkg_apis_toolchain_v1alpha1_UserAccount(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountSpec":           schema_pkg_apis_toolchain_v1alpha1_UserAccountSpec(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountSpecBase":       schema_pkg_apis_toolchain_v1alpha1_UserAccountSpecBase(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountSpecEmbedded":   schema_pkg_apis_toolchain_v1alpha1_UserAccountSpecEmbedded(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountStatus":         schema_pkg_apis_toolchain_v1alpha1_UserAccountStatus(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserSignup":                schema_pkg_apis_toolchain_v1alpha1_UserSignup(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserSignupSpec":            schema_pkg_apis_toolchain_v1alpha1_UserSignupSpec(ref),
		"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserSignupStatus":          schema_pkg_apis_toolchain_v1alpha1_UserSignupStatus(ref),
	}
}

func schema_pkg_apis_toolchain_v1alpha1_BannedUser(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BannedUser is used to maintain a list of banned e-mail addresses",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.BannedUserSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.BannedUserSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_BannedUserSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BannedUserSpec defines the desired state of BannedUser",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"email": {
						SchemaProps: spec.SchemaProps{
							Description: "The e-mail address of the account that has been banned",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"email"},
			},
		},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_MasterUserRecord(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MasterUserRecord keeps all information about user, user accounts and namespaces provisioned in CodeReady Toolchain",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecordSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecordStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecordSpec", "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.MasterUserRecordStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_MasterUserRecordSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MasterUserRecordSpec defines the desired state of MasterUserRecord",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"userID": {
						SchemaProps: spec.SchemaProps{
							Description: "UserID is the user ID from RHD Identity Provider token (“sub” claim)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"disabled": {
						SchemaProps: spec.SchemaProps{
							Description: "If set to true then the corresponding user should not be able to login (but the underlying UserAccounts still exists) \"false\" is assumed by default",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"banned": {
						SchemaProps: spec.SchemaProps{
							Description: "If set to true then the corresponding user has been banned from logging in and accessing their resources",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"deprovisioned": {
						SchemaProps: spec.SchemaProps{
							Description: "If set to true then the corresponding UserAccount should be deleted \"false\" is assumed by default",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"userAccounts": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "The list of user accounts in the member clusters which belong to this MasterUserRecord",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountEmbedded"),
									},
								},
							},
						},
					},
				},
				Required: []string{"userID"},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountEmbedded"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_MasterUserRecordStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MasterUserRecordStatus defines the observed state of MasterUserRecord",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type":       "",
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is an array of current Master User Record conditions Supported condition types: Provisioning, UserAccountNotReady and Ready",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"),
									},
								},
							},
						},
					},
					"userAccounts": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "The status of user accounts in the member clusters which belong to this MasterUserRecord",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountStatusEmbedded"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition", "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountStatusEmbedded"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_NSTemplateSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSTemplateSet defines user environment via templates that are used for namespace provisioning",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec", "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_NSTemplateSetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSTemplateSetSpec defines the desired state of NSTemplateSet",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"tierName": {
						SchemaProps: spec.SchemaProps{
							Description: "The name of the tier represented by this template set",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"namespaces": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "The namespace templates",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetNamespace"),
									},
								},
							},
						},
					},
				},
				Required: []string{"tierName", "namespaces"},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetNamespace"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_NSTemplateSetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSTemplateSetStatus defines the observed state of NSTemplateSet",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type":       "",
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is an array of current NSTemplateSet conditions Supported condition types: ConditionReady",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_NSTemplateTier(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSTemplateTier configures user environment via templates used for namespaces the user has access to",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierSpec", "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_NSTemplateTierSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSTemplateTierSpec defines the desired state of NSTemplateTier",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"namespaces": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "The namespace templates",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierNamespace"),
									},
								},
							},
						},
					},
				},
				Required: []string{"namespaces"},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateTierNamespace"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_NSTemplateTierStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSTemplateTierStatus defines the observed state of NSTemplateTier",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type":       "",
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is an array of current NSTemplateTier conditions Supported condition types: ConditionReady",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_RegistrationService(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RegistrationService configures the registration service deployment",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.RegistrationServiceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.RegistrationServiceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.RegistrationServiceSpec", "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.RegistrationServiceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_RegistrationServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RegistrationServiceSpec defines the desired state of RegistrationService",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"envVars": {
						SchemaProps: spec.SchemaProps{
							Description: "The environment variables are supposed to be set to registration service deployment template",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_RegistrationServiceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RegistrationServiceStatus defines the observed state of RegistrationService",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type":       "",
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is an array of current Registration Service deployment conditions Supported condition reasons: Deploying, and Deployed",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserAccount(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserAccount keeps all information about user provisioned in the cluster",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountSpec", "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserAccountStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserAccountSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserAccountSpec defines the desired state of UserAccount",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"userID": {
						SchemaProps: spec.SchemaProps{
							Description: "UserID is the user ID from RHD Identity Provider token (“sub” claim) Is to be used to create Identity and UserIdentityMapping resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"disabled": {
						SchemaProps: spec.SchemaProps{
							Description: "If set to true then the corresponding user should not be able to login \"false\" is assumed by default",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"nsLimit": {
						SchemaProps: spec.SchemaProps{
							Description: "The namespace limit name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nsTemplateSet": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace template set",
							Ref:         ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec"),
						},
					},
				},
				Required: []string{"userID", "nsLimit", "nsTemplateSet"},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserAccountSpecBase(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserAccountSpecBase defines the common fields between UserAccountSpec and UserAccountSpecEmbedded",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nsLimit": {
						SchemaProps: spec.SchemaProps{
							Description: "The namespace limit name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nsTemplateSet": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace template set",
							Ref:         ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec"),
						},
					},
				},
				Required: []string{"nsLimit", "nsTemplateSet"},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserAccountSpecEmbedded(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserAccountSpecEmbedded defines the desired state of UserAccount",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nsLimit": {
						SchemaProps: spec.SchemaProps{
							Description: "The namespace limit name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nsTemplateSet": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace template set",
							Ref:         ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec"),
						},
					},
				},
				Required: []string{"nsLimit", "nsTemplateSet"},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.NSTemplateSetSpec"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserAccountStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserAccountStatus defines the observed state of UserAccount",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type":       "",
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is an array of current User Account conditions Supported condition types: ConditionReady",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserSignup(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserSignup registers a user in the CodeReady Toolchain",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserSignupSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserSignupStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserSignupSpec", "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.UserSignupStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserSignupSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserSignupSpec defines the desired state of UserSignup",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"targetCluster": {
						SchemaProps: spec.SchemaProps{
							Description: "The cluster in which the user is provisioned in If not set then the target cluster will be picked automatically",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"approved": {
						SchemaProps: spec.SchemaProps{
							Description: "If Approved set to 'true' then the user has been manually approved If not set then the user is subject of auto-approval (if enabled)",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"deactivated": {
						SchemaProps: spec.SchemaProps{
							Description: "Deactivated is used to deactivate the user.  If not set, then by default the user is active",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"username": {
						SchemaProps: spec.SchemaProps{
							Description: "The user's username, obtained from the identity provider.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"username"},
			},
		},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_UserSignupStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserSignupStatus defines the observed state of UserSignup",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type":       "",
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is an array of current UserSignup conditions Supported condition types: PendingApproval, Provisioning, Complete",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"),
									},
								},
							},
						},
					},
					"compliantUsername": {
						SchemaProps: spec.SchemaProps{
							Description: "CompliantUsername is used to store the transformed, DNS-1123 compliant username",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"},
	}
}
