package relationships

import (
	"fmt"
	"sort"
	"strings"

	core "github.com/authzed/spicedb/pkg/proto/core/v1"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/spicedb/pkg/tuple"
)

func NewRelationships() *Relationships {
	return &Relationships{
		Relations:   []v1.Relationship{},
		True:        []v1.Relationship{},
		False:       []v1.Relationship{},
		Validations: []v1.Relationship{},
	}
}

type Relationships struct {
	Relations   []v1.Relationship
	True        []v1.Relationship
	False       []v1.Relationship
	Validations []v1.Relationship
}

func (r *Relationships) AddValidation(relationship v1.Relationship) {
	r.Validations = append(r.Validations, relationship)
}

// AddRelation adds the graph relation for the playground.
func (r *Relationships) AddRelation(relationship v1.Relationship) {
	r.Relations = append(r.Relations, relationship)
}

func (r *Relationships) AssertTrue(relationship v1.Relationship) {
	r.True = append(r.True, relationship)
}

func (r *Relationships) AssertFalse(relationship v1.Relationship) {
	r.False = append(r.False, relationship)
}

func (r Relationships) AllRelations() []v1.Relationship {
	return r.Relations
}

func (r Relationships) AllFalse() []v1.Relationship {
	return r.False
}

func (r Relationships) AllTrue() []v1.Relationship {
	return r.True
}

func (r Relationships) AllValidations() []v1.Relationship {
	return r.Validations
}

type ObjectWithRelationships interface {
	AllRelations() []v1.Relationship
	AllTrue() []v1.Relationship
	AllFalse() []v1.Relationship
	AllValidations() []v1.Relationship
	Type() string
	Object() *v1.ObjectReference
}

var allObjects []ObjectWithRelationships

func AllAssertTrue() []string {
	all := make([]string, 0)
	for _, o := range allObjects {
		for _, t := range o.AllTrue() {
			rStr, err := tuple.StringRelationship(&t)
			if err != nil {
				panic(err)
			}
			all = append(all, rStr)
		}
	}
	return all
}

func AllValidations() map[string][]string {
	all := make(map[string][]string, 0)
	for _, o := range allObjects {
		for _, t := range o.AllValidations() {
			rStr := tuple.StringONR(&core.ObjectAndRelation{
				Namespace: t.Resource.ObjectType,
				ObjectId:  t.Resource.ObjectId,
				Relation:  t.Relation,
			})

			all[rStr] = []string{}
		}
	}
	return all
}

func AllAssertFalse() []string {
	all := make([]string, 0)
	for _, o := range allObjects {
		for _, t := range o.AllFalse() {
			rStr, err := tuple.StringRelationship(&t)
			if err != nil {
				panic(err)
			}
			all = append(all, rStr)
		}
	}
	return all
}

func AllRelationsToStrings() string {
	// group all the objects
	buckets := make(map[string][]ObjectWithRelationships)
	bucketKeys := make([]string, 0)
	for _, o := range allObjects {
		o := o
		if _, ok := buckets[o.Type()]; !ok {
			bucketKeys = append(bucketKeys, o.Type())
		}
		buckets[o.Type()] = append(buckets[o.Type()], o)
	}

	var allStrings []string

	sort.Strings(bucketKeys)
	for _, bucketKey := range bucketKeys {
		bucket := buckets[bucketKey]
		allStrings = append(allStrings, "// All "+bucket[0].Type()+"s")
		for _, o := range bucket {
			for _, r := range o.AllRelations() {
				r := r
				rStr, err := tuple.StringRelationship(&r)
				if err != nil {
					panic(err)
				}
				allStrings = append(allStrings, rStr)
			}
		}
	}
	return strings.Join(allStrings, "\n")
}

func WorkspaceWithDeps(id string, team *ObjTeam, template *ObjTemplate) *ObjWorkspace {
	// Building a workspace means the team needs access to the template + provisioner
	template.CanUseBy(team) // This should be a perm check

	workspace := Workspace(id).Owner(team)
	build := Workspace_build(fmt.Sprintf("%s/build", id)).
		Workspace(workspace)
	agent := Workspace_agent(fmt.Sprintf("%s/agent", id)).
		Workspace(workspace)
	app := Worspace_app(fmt.Sprintf("%s/app", id)).
		Workspace(workspace)
	resources := Workspace_resources(fmt.Sprintf("%s/resources", id)).
		Workspace(workspace)

	// Add the template + provisioner relations
	template.Workspace(workspace)

	var _, _, _, _ = build, agent, app, resources
	return workspace
}

func (obj *ObjTemplate) Version(id string) *ObjTemplate_version {
	// This "/" syntax is not required. We usually use uuids, but strings
	// are easier to read, and this helps us intuitively see relations.
	return Template_version(fmt.Sprintf("%s/%s", obj.Obj.ObjectId, id)).
		Template(obj)
}
