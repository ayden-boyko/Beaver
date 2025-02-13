package tests

import (
	"testing"

	"Beaver/pkg"
)

const test_file = "./test.json"

var testBeaver *pkg.Beaver

func init() {
	testBeaver, _ = pkg.NewBeaver(test_file)
}

// test beaver creation
func TestBeaverCreated(t *testing.T) {
	if testBeaver == nil {
		t.Errorf("Beaver not created")
	}
}

func TestBeaverConfigured(t *testing.T) {
	if testBeaver.GetLevel() != "info" {
		t.Errorf("Beaver not configured, level")
	}
	if testBeaver.GetFilePath() != test_file {
		t.Errorf("Beaver not configured, filepath")
	}
}

func TestBeaverCreatedWithYAML(t *testing.T) {
	testYAMLPath := "test_yaml.yaml"
	tbeaver, err := pkg.NewBeaverFromFile(testYAMLPath)
	if err != nil {
		t.Errorf("Error creating beaver: %v", err)
	}
	if tbeaver == nil {
		t.Errorf("Beaver not created")
	}
}

func TestBeaverCreatedWithJSON(t *testing.T) {
	testJSONPath := "test_json.json"
	tbeaver, err := pkg.NewBeaverFromFile(testJSONPath)
	if err != nil {
		t.Errorf("Error creating beaver: %v", err)
	}
	if tbeaver == nil {
		t.Errorf("Beaver not created")
	}
}

// test general log
func TestBeaverLog(t *testing.T) {
	t.Errorf("TestBeaverLog")
}

// test warn log
func TestBeaverWarn(t *testing.T) {
	t.Errorf("TestBeaverWarn")
}

// test error log
func TestBeaverError(t *testing.T) {
	t.Errorf("TestBeaverError")
}
