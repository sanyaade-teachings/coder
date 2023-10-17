package relationships

import (
	"fmt"
	"sort"
	"strings"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/spicedb/pkg/tuple"
)

func NewRelationships() *Relationships {
	return &Relationships{
		Relations: []v1.Relationship{},
	}
}

type Relationships struct {
	Relations []v1.Relationship
}

func (r *Relationships) AddRelation(relationship v1.Relationship) {
	r.Relations = append(r.Relations, relationship)
}

//func (r *Relationships)

func (r Relationships) AllRelations() []v1.Relationship {
	return r.Relations
}

type ObjectWithRelationships interface {
	AllRelations() []v1.Relationship
	Type() string
}

var allObjects []ObjectWithRelationships

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

func WorkspaceWithDeps(id string) *ObjWorkspace {
	workspace := Workspace(id)
	build := Workspace_build(fmt.Sprintf("%s/build", id)).
		Workspace(workspace)
	agent := Workspace_agent(fmt.Sprintf("%s/agent", id)).
		Workspace(workspace)
	app := Worspace_app(fmt.Sprintf("%s/app", id)).
		Workspace(workspace)
	resources := Workspace_resources(fmt.Sprintf("%s/resources", id)).
		Workspace(workspace)

	var _, _, _, _ = build, agent, app, resources
	return workspace
}

func (obj *ObjTemplate) Version(id string) *ObjTemplate_version {
	// This "/" syntax is not required. We usually use uuids, but strings
	// are easier to read, and this helps us intuitively see relations.
	return Template_version(fmt.Sprintf("%s/%s", obj.Obj.ObjectId, id)).
		Template(obj)
}
