package main

import (
	"github.com/coder/coder/v2/coderd/database/spice/policy/playground/relationships"
	"gopkg.in/yaml.v3"

	"github.com/coder/coder/v2/coderd/database/spice/policy"
)

type AssertStruct struct {
	True  []string `yaml:"assertTrue"`
	False []string `yaml:"assertFalse"`
}

type PlaygroundYAML struct {
	Schema        string              `yaml:"schema"`
	Relationships string              `yaml:"relationships"`
	Assertions    AssertStruct        `yaml:"assertions"`
	Validation    map[string][]string `yaml:"validation"`
}

func PlaygroundExport() string {
	relationships.GenerateRelationships()
	y := PlaygroundYAML{
		Schema:        policy.Schema,
		Relationships: relationships.AllRelationsToStrings(),
		Assertions: AssertStruct{
			True:  relationships.AllAssertTrue(),
			False: relationships.AllAssertFalse(),
		},
		Validation: relationships.AllValidations(),
	}
	out, err := yaml.Marshal(y)
	if err != nil {
		panic(err)
	}
	return string(out)
}
