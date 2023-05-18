package provider_test

import (
	"testing"

	"github.com/coder/terraform-provider-coder/provider"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	const (
		legacyVariable     = "Legacy Variable"
		legacyVariableName = "Legacy Variable Name"

		displayName = "Display Name"
	)

	aMap := map[string]interface{}{
		"name":                 "Parameter Name",
		"display_name":         displayName,
		"legacy_variable":      legacyVariable,
		"legacy_variable_name": legacyVariableName,
	}

	var param provider.Parameter
	err := mapstructure.Decode(aMap, &param)
	require.NoError(t, err)
	require.Equal(t, displayName, param.DisplayName)
	require.Equal(t, legacyVariable, param.LegacyVariable)
	require.Equal(t, legacyVariableName, param.LegacyVariableName)
}