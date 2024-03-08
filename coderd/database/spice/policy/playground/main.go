package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	core "github.com/authzed/spicedb/pkg/proto/core/v1"
	"github.com/authzed/spicedb/pkg/tuple"
	"gopkg.in/yaml.v3"

	"github.com/coder/coder/v2/coderd/database/spice/policy"
	"github.com/coder/coder/v2/coderd/database/spice/policy/playground/relationships"
)

func main() {
	if len(os.Args) <= 1 {
		usage()
		return
	}

	switch os.Args[1] {
	case "export":
		fmt.Println(PlaygroundExport())
	default:
		usage()
	}
}

func usage() {
	fmt.Println("Usage: policycmd [command]")
	fmt.Println("Commands:")
	fmt.Println("  export")
	// fmt.Println("  import")
	//fmt.Println("  playground")
	//fmt.Println("  run")
	//fmt.Println("  test")
}

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
		Relationships: RelationshipsToCSV(relationships.Playground.Builder.Relationships),
		Assertions: AssertStruct{
			True:  RelationshipsToStrings(relationships.Playground.True),
			False: RelationshipsToStrings(relationships.Playground.False),
		},
		Validation: ExportValidations(relationships.Playground.Validations),
	}
	out, err := yaml.Marshal(y)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func RelationshipsToStrings(rels []v1.Relationship) []string {
	allStrings := make([]string, 0)
	for _, r := range rels {
		rStr, err := tuple.StringRelationship(&r)
		if err != nil {
			panic(err)
		}
		allStrings = append(allStrings, rStr)
	}
	return allStrings
}

func GroupRelations(rels []v1.Relationship) map[string][]v1.Relationship {
	buckets := make(map[string][]v1.Relationship)

	for _, rel := range rels {
		rel := rel
		if _, ok := buckets[rel.Resource.ObjectType]; !ok {
			buckets[rel.Resource.ObjectType] = []v1.Relationship{}
		}
		buckets[rel.Resource.ObjectType] = append(buckets[rel.Resource.ObjectType], rel)
	}

	for k := range buckets {
		sort.Slice(buckets[k], func(i, j int) bool {
			// Do best effort to keep same things grouped within the section
			return buckets[k][i].Resource.ObjectId < buckets[k][j].Resource.ObjectId
		})
	}
	return buckets
}

func RelationshipsToCSV(rels []v1.Relationship) string {
	buckets := GroupRelations(rels)

	// group all the objects
	var allStrings []string
	for bucketKey := range buckets {
		bucket := buckets[bucketKey]
		allStrings = append(allStrings, fmt.Sprintf("// %ss with  %d relations", bucketKey, len(bucket)))
		sort.Slice(bucket, func(i, j int) bool {
			// Do best effort to keep same things grouped within the section
			return bucket[i].Resource.ObjectId < bucket[j].Resource.ObjectId
		})

		allStrings = append(allStrings, RelationshipsToStrings(bucket)...)
	}
	return strings.Join(allStrings, "\n")
}

func ExportValidations(validate []v1.Relationship) map[string][]string {
	all := make(map[string][]string, 0)

	buckets := GroupRelations(validate)
	for k := range buckets {
		bucket := buckets[k]
		for _, r := range bucket {
			rStr := tuple.StringONR(&core.ObjectAndRelation{
				Namespace: r.Resource.ObjectType,
				ObjectId:  r.Resource.ObjectId,
				Relation:  r.Relation,
			})
			// The playground will populate this array when you click "Regenerate"
			all[rStr] = []string{}
		}
	}

	return all
}
