package gen

import (
	_ "embed"
	"fmt"
	"go/format"
	"sort"
	"strings"
	"text/template"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	core "github.com/authzed/spicedb/pkg/proto/core/v1"
	"github.com/authzed/spicedb/pkg/schemadsl/compiler"
	"golang.org/x/xerrors"
)

//go:embed relationships.tmpl
var templateText string

type GenerateOptions struct {
	Package string
}

func Generate(schema string, opts GenerateOptions) (string, error) {
	if opts.Package == "" {
		opts.Package = "policy"
	}

	var prefix string
	compiled, err := compiler.Compile(compiler.InputSchema{
		Source:       "policy.zed",
		SchemaString: schema,
	}, compiler.ObjectTypePrefix(prefix))
	if err != nil {
		return "", xerrors.Errorf("compile schema: %w", err)
	}

	tpl, err := template.New("relationships.tmpl").Funcs(template.FuncMap{
		"capitalize": capitalize,
	}).Parse(templateText)
	if err != nil {
		return "", xerrors.Errorf("parse template: %w", err)
	}

	var output strings.Builder
	_, _ = output.WriteString(`// Package relationships code generated. DO NOT EDIT.`)
	_, _ = output.WriteString("\n")
	_, _ = output.WriteString(fmt.Sprintf(`package %s`, opts.Package))
	_, _ = output.WriteString("\n")
	_, _ = output.WriteString(`import (
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"

	"fmt"
)`)
	_, _ = output.WriteString("\n")

	// sort order for consistency
	sort.Slice(compiled.ObjectDefinitions, func(i, j int) bool {
		return compiled.ObjectDefinitions[i].Name < compiled.ObjectDefinitions[j].Name
	})
	// TODO: Handle caveats
	sort.Slice(compiled.CaveatDefinitions, func(i, j int) bool {
		return compiled.CaveatDefinitions[i].Name < compiled.CaveatDefinitions[j].Name
	})

	for _, obj := range compiled.ObjectDefinitions {
		d := newDef(obj)
		err := tpl.Execute(&output, d)
		if err != nil {
			return "", xerrors.Errorf("[%q] execute template: %w", obj.Name, err)
		}
		output.WriteString("\n")
	}

	formatted, err := format.Source([]byte(output.String()))
	if err != nil {
		return "", xerrors.Errorf("format source: %w", err)
	}
	return string(formatted), nil
}

// objectDefinition is the type used for the template generator.
type objectDefinition struct {
	// The core type
	*core.NamespaceDefinition

	DirectRelations []objectDirectRelation
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

	// Each relation is a "relation" line in the schema. Example:
	// 	relation member: group#membership | user
	for _, r := range obj.Relation {
		if r.UsersetRewrite != nil {
			// This is a permission.
			continue
		}

		// Each "AllowedDirectRelations" is a "subject" on the right hand side
		// of a relation. So in the example above, the subjects are:
		//	- group#membership
		//	- user
		multipleSubjects := make([]objectDirectRelation, 0)
		for _, d := range r.TypeInformation.AllowedDirectRelations {
			optRel := ""
			// TODO: I cannot recall what this is
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

			multipleSubjects = append(multipleSubjects, objectDirectRelation{
				RelationName: r.Name,
				FunctionName: r.Name,
				Subject: v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: d.Namespace,
						// This ObjectID will be overwritten, so just use a placeholder.
						ObjectId: "<id>",
					},
					OptionalRelation: optRel,
				},
			})
		}

		// If we have more than 1 subject, we need to suffix the name of the function with the
		// object type. Otherwise, we have 2 functions with the same name.
		//
		// Example:
		//	- group#member -- func MemberGroup
		//	- user  	   -- func MemberUser
		//
		// If we only have 1 subject, we do not need the suffix.
		// If someone could find a typesafe argument pattern with generics, that would be great.
		if len(multipleSubjects) > 1 {
			for i := range multipleSubjects {
				multipleSubjects[i].FunctionName += capitalize(multipleSubjects[i].Subject.Object.ObjectType)
			}
		}

		rels = append(rels, multipleSubjects...)
	}
	d.DirectRelations = rels
	return d
}

func capitalize(name string) string {
	return strings.ToUpper(string(name[0])) + name[1:]
}
