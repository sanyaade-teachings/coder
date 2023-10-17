package main

import (
	_ "embed"
	"fmt"
	"go/format"
	"strings"
	"text/template"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	core "github.com/authzed/spicedb/pkg/proto/core/v1"

	"github.com/authzed/spicedb/pkg/schemadsl/compiler"
	"github.com/coder/coder/v2/coderd/database/spice/policy"
)

//go:embed relationships.tmpl
var templateText string

func main() {
	fmt.Println(Generate())
}

func capitalize(name string) string {
	return strings.ToUpper(string(name[0])) + name[1:]
}

func Generate() string {
	var prefix string
	compiled, err := compiler.Compile(compiler.InputSchema{
		Source:       "policy.zed",
		SchemaString: policy.Schema,
	}, &prefix)
	if err != nil {
		panic(err)
	}

	tpl := template.New("zanzobjects").Funcs(template.FuncMap{
		"capitalize": capitalize,
		"unique":     uniquePermissions,
	})

	tpl, err = tpl.Parse(templateText)
	if err != nil {
		panic(err)
	}

	var output strings.Builder
	output.WriteString(`// Code generated. DO NOT EDIT.`)
	output.WriteString("\n")
	output.WriteString(`package relationships`)
	output.WriteString("\n")
	output.WriteString(`import v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"`)
	output.WriteString("\n")

	for _, obj := range compiled.ObjectDefinitions {
		d := newDef(obj)
		var _ = d
		err := tpl.Execute(&output, d)
		if err != nil {
			panic(err)
		}
		output.WriteString("\n")
	}

	formatted, err := format.Source([]byte(output.String()))
	if err != nil {
		fmt.Println(output.String())
		panic(err)
	}
	return string(formatted)
}

type objectDefinition struct {
	// The core type
	*core.NamespaceDefinition

	Permissions     []objectPermission
	DirectRelations []objectDirectRelation
}

type objectPermission struct {
	Permission   string
	FunctionName string
}

type objectDirectRelation struct {
	RelationName string
	FunctionName string
	Subject      v1.SubjectReference
}

func newDef(obj *core.NamespaceDefinition) objectDefinition {
	d := objectDefinition{
		NamespaceDefinition: obj,
	}
	rels := make([]objectDirectRelation, 0)
	perms := make([]objectPermission, 0)

	for _, r := range obj.Relation {
		if r.UsersetRewrite != nil {
			// This is a permission.
			perms = append(perms, objectPermission{
				Permission:   r.Name,
				FunctionName: capitalize(r.Name),
			})
			continue
		}

		dedup := 0
		multipleSubjects := make([]objectDirectRelation, 0)
		// For the "relation" we should write a helper function to create
		// this relationship between two objects.
		for _, d := range r.TypeInformation.AllowedDirectRelations {
			optRel := ""
			if d.GetRelation() != "..." {
				optRel = d.GetRelation()
			}

			if d.GetPublicWildcard() != nil {
				multipleSubjects = append(multipleSubjects, objectDirectRelation{
					RelationName: r.Name,
					FunctionName: r.Name,
					Subject: v1.SubjectReference{
						Object: &v1.ObjectReference{
							ObjectType: d.Namespace,
							ObjectId:   "*",
						},
						OptionalRelation: optRel,
					},
				})
				continue
			}

			dedup++
			multipleSubjects = append(multipleSubjects, objectDirectRelation{
				RelationName: r.Name,
				FunctionName: r.Name,
				Subject: v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: d.Namespace,
						ObjectId:   "<id>",
					},
					OptionalRelation: optRel,
				},
			})
		}

		if dedup > 1 {
			for i := range multipleSubjects {
				// Remove method name conflicts
				multipleSubjects[i].FunctionName += capitalize(multipleSubjects[i].Subject.Object.ObjectType)
			}
		}
		rels = append(rels, multipleSubjects...)
	}
	d.DirectRelations = rels
	d.Permissions = perms
	return d
}

func uniquePermissions(perms []objectPermission) []objectPermission {
	seen := make(map[string]struct{})
	out := make([]objectPermission, 0)
	for _, perm := range perms {
		perm := perm
		if _, ok := seen[perm.Permission]; ok {
			continue
		}
		seen[perm.Permission] = struct{}{}
		out = append(out, perm)
	}
	return out
}
