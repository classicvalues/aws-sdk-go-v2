// Code generated by smithy-go-codegen DO NOT EDIT.

package endpoints

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/endpoints"
	"regexp"
)

// Options is the endpoint resolver configuration options
type Options struct {
	DisableHTTPS bool
}

// Resolver Greengrass endpoint resolver
type Resolver struct {
	partitions endpoints.Partitions
}

// ResolveEndpoint resolves the service endpoint for the given region and options
func (r *Resolver) ResolveEndpoint(region string, options Options) (endpoint aws.Endpoint, err error) {
	if len(region) == 0 {
		return endpoint, &aws.MissingRegionError{}
	}

	opt := endpoints.Options{
		DisableHTTPS: options.DisableHTTPS,
	}
	return r.partitions.ResolveEndpoint(region, opt)
}

// New returns a new Resolver
func New() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var partitionRegexp = struct {
	Aws      *regexp.Regexp
	AwsCn    *regexp.Regexp
	AwsIso   *regexp.Regexp
	AwsIsoB  *regexp.Regexp
	AwsUsGov *regexp.Regexp
}{

	Aws:      regexp.MustCompile("^(us|eu|ap|sa|ca|me|af)\\-\\w+\\-\\d+$"),
	AwsCn:    regexp.MustCompile("^cn\\-\\w+\\-\\d+$"),
	AwsIso:   regexp.MustCompile("^us\\-iso\\-\\w+\\-\\d+$"),
	AwsIsoB:  regexp.MustCompile("^us\\-isob\\-\\w+\\-\\d+$"),
	AwsUsGov: regexp.MustCompile("^us\\-gov\\-\\w+\\-\\d+$"),
}

var defaultPartitions = endpoints.Partitions{
	{
		ID: "aws",
		Defaults: endpoints.Endpoint{
			Hostname:          "greengrass.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.Aws,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"ap-northeast-1": endpoints.Endpoint{},
			"ap-northeast-2": endpoints.Endpoint{},
			"ap-south-1":     endpoints.Endpoint{},
			"ap-southeast-1": endpoints.Endpoint{},
			"ap-southeast-2": endpoints.Endpoint{},
			"eu-central-1":   endpoints.Endpoint{},
			"eu-west-1":      endpoints.Endpoint{},
			"eu-west-2":      endpoints.Endpoint{},
			"us-east-1":      endpoints.Endpoint{},
			"us-east-2":      endpoints.Endpoint{},
			"us-west-2":      endpoints.Endpoint{},
		},
	},
	{
		ID: "aws-cn",
		Defaults: endpoints.Endpoint{
			Hostname:          "greengrass.{region}.amazonaws.com.cn",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.AwsCn,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"cn-north-1": endpoints.Endpoint{},
		},
	},
	{
		ID: "aws-iso",
		Defaults: endpoints.Endpoint{
			Hostname:          "greengrass.{region}.c2s.ic.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.AwsIso,
		IsRegionalized: true,
	},
	{
		ID: "aws-iso-b",
		Defaults: endpoints.Endpoint{
			Hostname:          "greengrass.{region}.sc2s.sgov.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.AwsIsoB,
		IsRegionalized: true,
	},
	{
		ID: "aws-us-gov",
		Defaults: endpoints.Endpoint{
			Hostname:          "greengrass.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.AwsUsGov,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"dataplane-us-gov-east-1": endpoints.Endpoint{
				Hostname: "greengrass-ats.iot.us-gov-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-east-1",
				},
			},
			"dataplane-us-gov-west-1": endpoints.Endpoint{
				Hostname: "greengrass-ats.iot.us-gov-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-west-1",
				},
			},
			"fips-us-gov-east-1": endpoints.Endpoint{
				Hostname: "greengrass-fips.us-gov-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-east-1",
				},
			},
			"us-gov-east-1": endpoints.Endpoint{
				Hostname: "greengrass.us-gov-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-east-1",
				},
			},
			"us-gov-west-1": endpoints.Endpoint{
				Hostname: "greengrass.us-gov-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-west-1",
				},
			},
		},
	},
}
