package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKSIMetrics(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp("", "ksi-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir) //nolint:errcheck

	// Create a test YAML file
	testYAML := `
ksi1:
  components_in_inventory_count: 100
  unauthorized_components_in_inventory_count: 5

ksi2:
  mttr_critical: 24
  mttr_high: 48
  mttr_medium: 72
  mttr_low: 168

ksi3:
  detected_kev_count: 10
  late_kev_count: 2

ksi4:
  unique_component_configurations_count: 50
  unapproved_configuration_count: 3

ksi5:
  mtti_critical: 1
  mtti_high: 2
  mtti_medium: 4
  mtti_low: 8

ksi6:
  accounts_count: 200
  mfa_accounts_count: 195

ksi7:
  encryption_implementation_count: 30
  fips_encryption_implementation_count: 28

ksi8:
  ipv6_support: true
  customer_configured_fips_140: true
  section_508_compliance: true
  customer_accessible_security_logging: true
  customer_configurable_data_retention: true
`

	testFile := filepath.Join(tmpDir, "test_ksi.yaml")
	err = os.WriteFile(testFile, []byte(testYAML), 0644)
	if err != nil {
		t.Fatalf("failed to write test YAML file: %v", err)
	}

	// Test cases
	tests := []struct {
		name    string
		file    string
		want    *KSIMetrics
		wantErr bool
	}{
		{
			name: "valid yaml file",
			file: testFile,
			want: &KSIMetrics{
				KSI1: KSI1{
					ComponentsInInventoryCount:           100,
					UnauthorizedComponentsInventoryCount: 5,
				},
				KSI2: KSI2{
					MTTRCritical: 24,
					MTTRHigh:     48,
					MTTRMedium:   72,
					MTTRLow:      168,
				},
				KSI3: KSI3{
					DetectedKEVCount: 10,
					LateKEVCount:     2,
				},
				KSI4: KSI4{
					UniqueComponentConfigurationsCount: 50,
					UnapprovedConfigurationCount:       3,
				},
				KSI5: KSI5{
					MTTICritical: 1,
					MTTIHigh:     2,
					MTTIMedium:   4,
					MTTILow:      8,
				},
				KSI6: KSI6{
					AccountsCount:    200,
					MFAAccountsCount: 195,
				},
				KSI7: KSI7{
					EncryptionImplementationCount:     30,
					FIPSEncryptionImplementationCount: 28,
				},
				KSI8: KSI8{
					IPv6Support:                       true,
					CustomerConfiguredFIPS140:         true,
					Section508Compliance:              true,
					CustomerAccessibleSecurityLogging: true,
					CustomerConfigurableDataRetention: true,
				},
			},
			wantErr: false,
		},
		{
			name:    "non-existent file",
			file:    filepath.Join(tmpDir, "nonexistent.yaml"),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty file path",
			file:    "",
			want:    nil,
			wantErr: true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewKSIMetrics(tt.file)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, tt.want, got)

			// Detailed field comparisons
			if got != nil {
				// KSI1
				assert.Equal(t, tt.want.KSI1.ComponentsInInventoryCount, got.KSI1.ComponentsInInventoryCount)
				assert.Equal(t, tt.want.KSI1.UnauthorizedComponentsInventoryCount, got.KSI1.UnauthorizedComponentsInventoryCount)

				// KSI2
				assert.Equal(t, tt.want.KSI2.MTTRCritical, got.KSI2.MTTRCritical)
				assert.Equal(t, tt.want.KSI2.MTTRHigh, got.KSI2.MTTRHigh)
				assert.Equal(t, tt.want.KSI2.MTTRMedium, got.KSI2.MTTRMedium)
				assert.Equal(t, tt.want.KSI2.MTTRLow, got.KSI2.MTTRLow)

				// KSI3
				assert.Equal(t, tt.want.KSI3.DetectedKEVCount, got.KSI3.DetectedKEVCount)
				assert.Equal(t, tt.want.KSI3.LateKEVCount, got.KSI3.LateKEVCount)

				// Continue with other KSI metrics...
			}
		})
	}
}
