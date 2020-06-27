package config_test

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/coffemanfp/shachat/config"
	"github.com/coffemanfp/shachat/utils"
	"gopkg.in/yaml.v3"
)

var configFile string = "../config.yaml"

func TestInvalidConfig(t *testing.T) {
	// Checking that the settings have no value. Because they are not configured yet.
	settings, err := config.GetSettings()
	if err == nil {
		t.Errorf("settings has content, expected empty")
	}

	if valid := settings.Validate(); valid {
		t.Errorf("settings has content, expected empty")
	}
}

func TestConfig(t *testing.T) {

	// Configuring the settings.
	err := config.SetSettingsByFile(configFile)
	if err != nil {
		t.Fatalf("failed to configure settings:\n%s", err)
	}

	// Getting the settings to test.
	settings, err := config.GetSettings()
	if err != nil {
		t.Fatalf("failed to get the settings:\n%s", err)
	}

	// Getting the settings to expected.
	configBytes, err := utils.GetFilebytes(configFile)
	if err != nil {
		t.Fatalf("failed to read the config file:\n%s", err)
	}

	var settingsExpected config.Settings

	switch ext := filepath.Ext(configFile)[1:]; {
	case ext == "json":
		err = json.Unmarshal(configBytes, &settingsExpected)
	case ext == "yaml":
		err = yaml.Unmarshal(configBytes, &settingsExpected)
	}

	if err != nil {
		t.Fatalf("failed to unmarshalling the settings:\n%s", err)
	}

	settingsExpected.Database.Host = fmt.Sprintf(
		"%s:%d",
		settingsExpected.Database.Host,
		settingsExpected.Database.Port,
	)

	if settings != settingsExpected {
		t.Fatalf("settings %v not valid, expected %v", settings, settingsExpected)
	}
}
