package gen

import (
	_ "embed"
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

func Generate(schema string) (string, error) {
	var prefix string
	compiled, err := compiler.Compile(compiler.InputSchema{
		Source:       "policy.zed",
		SchemaString: schema,
	}, compiler.ObjectTypePrefix(prefix))
	if err != nil {
		return "", xerrors.Errorf("compile schema: %w", err)
	}

	tpl, err := template.New("spice_objects").Funcs(template.FuncMap{
		"capitalize": capitalize,
	}).Parse(templateText)
	if err != nil {
		return "", xerrors.Errorf("parse template: %w", err)
	}

	var output strings.Builder
	_, _ = output.WriteString(`// Package relationships code generated. DO NOT EDIT.`)
	_, _ = output.WriteString("\n")
	_, _ = output.WriteString(`package relationships`)
	_, _ = output.WriteString("\n")
	_, _ = output.WriteString(`import v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"`)
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

	for _, r := range obj.Relation {
		if r.UsersetRewrite != nil {
			// This is a permission.
			continue
		}

		// multipleSubjects is relations
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

		for i := range multipleSubjects {
			// TODO: Check for any method name conflicts
			multipleSubjects[i].FunctionName += capitalize(multipleSubjects[i].Subject.Object.ObjectType)
		}
		rels = append(rels, multipleSubjects...)
	}
	d.DirectRelations = rels
	return d
}

func capitalize(name string) string {
	return strings.ToUpper(string(name[0])) + name[1:]
}
