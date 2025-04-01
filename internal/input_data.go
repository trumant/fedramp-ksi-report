package internal

import (
	"os"

	"gopkg.in/yaml.v3"
)

type KSIMetrics struct {
	CSOName string `yaml:"cso_name"`
	KSI1    KSI1   `yaml:"ksi1"`
	KSI2    KSI2   `yaml:"ksi2"`
	KSI3    KSI3   `yaml:"ksi3"`
	KSI4    KSI4   `yaml:"ksi4"`
	KSI5    KSI5   `yaml:"ksi5"`
	KSI6    KSI6   `yaml:"ksi6"`
	KSI7    KSI7   `yaml:"ksi7"`
	KSI8    KSI8   `yaml:"ksi8"`
}

// KSI1 represents the key security indicator metrics for inventory coverage
type KSI1 struct {
	ComponentsInInventoryCount           uint `yaml:"components_in_inventory_count"`
	UnauthorizedComponentsInventoryCount uint `yaml:"unauthorized_components_in_inventory_count"`
}

func (c *KSI1) Coverage() float64 {
	if c.ComponentsInInventoryCount == 0 {
		return 0
	}
	return float64(c.ComponentsInInventoryCount-c.UnauthorizedComponentsInventoryCount) / float64(c.ComponentsInInventoryCount)
}

// KSI2 represents the key security indicator metrics for vulnerability remediation time by severity
type KSI2 struct {
	// MTTR (Mean Time to Remediate) values in hours
	MTTRCritical int `yaml:"mttr_critical"`
	MTTRHigh     int `yaml:"mttr_high"`
	MTTRMedium   int `yaml:"mttr_medium"`
	MTTRLow      int `yaml:"mttr_low"`
}

// KSI3 represents the key security indicator metrics for CISA KEV findings
type KSI3 struct {
	DetectedKEVCount int `yaml:"detected_kev_count"`
	LateKEVCount     int `yaml:"late_kev_count"`
}

// LatePercentage calculates the percentage of late KEV findings
func (c *KSI3) RemediatedOnTimePercentage() float64 {
	if c.DetectedKEVCount == 0 {
		return 0
	}
	return float64(c.DetectedKEVCount) / float64(c.DetectedKEVCount+c.LateKEVCount)
}

type KSI4 struct {
	UniqueComponentConfigurationsCount int `yaml:"unique_component_configurations_count"`
	UnapprovedConfigurationCount       int `yaml:"unapproved_configuration_count"`
}

func (c *KSI4) SecureConfigurationPercentage() float64 {
	if c.UniqueComponentConfigurationsCount == 0 {
		return 0
	}
	return float64(c.UniqueComponentConfigurationsCount) / float64(c.UniqueComponentConfigurationsCount+c.UnapprovedConfigurationCount)
}

type KSI5 struct {
	MTTICritical int `yaml:"mtti_critical"`
	MTTIHigh     int `yaml:"mtti_high"`
	MTTIMedium   int `yaml:"mtti_medium"`
	MTTILow      int `yaml:"mtti_low"`
}

type KSI6 struct {
	AccountsCount    int `yaml:"accounts_count"`
	MFAAccountsCount int `yaml:"mfa_accounts_count"`
}

func (c *KSI6) PrivilegedAccessCompliancePercentage() float64 {
	if c.AccountsCount == 0 {
		return 0
	}
	return float64(c.MFAAccountsCount) / float64(c.AccountsCount)
}

type KSI7 struct {
	EncryptionImplementationCount     int `yaml:"encryption_implementation_count"`
	FIPSEncryptionImplementationCount int `yaml:"fips_encryption_implementation_count"`
}

func (c *KSI7) CryptographicModuleCompliancePercentage() float64 {
	if c.EncryptionImplementationCount == 0 {
		return 0
	}
	return float64(c.FIPSEncryptionImplementationCount) / float64(c.EncryptionImplementationCount)
}

type KSI8 struct {
	IPv6Support                       bool `yaml:"ipv6_support"`
	CustomerConfiguredFIPS140         bool `yaml:"customer_configured_fips_140"`
	Section508Compliance              bool `yaml:"section_508_compliance"`
	CustomerAccessibleSecurityLogging bool `yaml:"customer_accessible_security_logging"`
	CustomerConfigurableDataRetention bool `yaml:"customer_configurable_data_retention"`
}

func NewKSIMetrics(filepath string) (*KSIMetrics, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var metrics KSIMetrics
	err = yaml.Unmarshal(data, &metrics)
	if err != nil {
		return nil, err
	}
	return &metrics, nil
}
