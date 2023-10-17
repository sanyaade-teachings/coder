package relationships

import v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"

type ObjPlatform struct {
	Obj *v1.ObjectReference
	*Relationships
}

func Platform(id string) *ObjPlatform {
	o := &ObjPlatform{
		Obj: &v1.ObjectReference{
			ObjectType: "platform",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjPlatform) Type() string {
	return "platform"
}

func (obj *ObjPlatform) Administrator(subs ...*ObjUser) *ObjPlatform {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "administrator",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

type ObjTeam struct {
	Obj *v1.ObjectReference
	*Relationships
}

func Team(id string) *ObjTeam {
	o := &ObjTeam{
		Obj: &v1.ObjectReference{
			ObjectType: "team",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjTeam) Type() string {
	return "team"
}

func (obj *ObjTeam) Platform(subs ...*ObjPlatform) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "platform",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) MemberGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "member",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) MemberUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "member",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Workspace_viewerGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace_viewer",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Workspace_viewerUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace_viewer",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Workspace_creatorGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace_creator",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Workspace_creatorUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace_creator",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Workspace_deletorGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace_deletor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Workspace_deletorUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace_deletor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Workspace_version_selectorGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace_version_selector",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Workspace_version_selectorUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace_version_selector",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Dangerous_workspace_connectorGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "dangerous_workspace_connector",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Dangerous_workspace_connectorUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "dangerous_workspace_connector",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Workspace_editorGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace_editor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Workspace_editorUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace_editor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Template_viewerGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_viewer",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Template_viewerUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_viewer",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Template_creatorGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_creator",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Template_creatorUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_creator",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Template_deletorGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_deletor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Template_deletorUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_deletor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Template_editorGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_editor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Template_editorUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_editor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Template_permission_managerGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_permission_manager",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Template_permission_managerUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_permission_manager",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Template_insights_viewerGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_insights_viewer",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Template_insights_viewerUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_insights_viewer",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

type ObjGroup struct {
	Obj *v1.ObjectReference
	*Relationships
}

func Group(id string) *ObjGroup {
	o := &ObjGroup{
		Obj: &v1.ObjectReference{
			ObjectType: "group",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjGroup) Type() string {
	return "group"
}

func (obj *ObjGroup) MemberUser(subs ...*ObjUser) *ObjGroup {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "member",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjGroup) MemberGroup(subs ...*ObjGroup) *ObjGroup {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "member",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "member",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjGroup) MemberWildcard() *ObjGroup {
	obj.AddRelation(v1.Relationship{
		Resource: obj.Obj,
		Relation: "member",
		Subject: &v1.SubjectReference{
			Object: &v1.ObjectReference{
				ObjectType: "user",
				ObjectId:   "*",
			},
			OptionalRelation: "",
		},
		OptionalCaveat: nil,
	})
	return obj
}

type ObjUser struct {
	Obj *v1.ObjectReference
	*Relationships
}

func User(id string) *ObjUser {
	o := &ObjUser{
		Obj: &v1.ObjectReference{
			ObjectType: "user",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjUser) Type() string {
	return "user"
}

type ObjWorkspace struct {
	Obj *v1.ObjectReference
	*Relationships
}

func Workspace(id string) *ObjWorkspace {
	o := &ObjWorkspace{
		Obj: &v1.ObjectReference{
			ObjectType: "workspace",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjWorkspace) ViewBy() string {
	return "workspace"
}

func (obj *ObjWorkspace) Type() string {
	return "workspace"
}

func (obj *ObjWorkspace) Owner(subs ...*ObjTeam) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "owner",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) ViewerGroup(subs ...*ObjGroup) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "viewer",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) ViewerUser(subs ...*ObjUser) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "viewer",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) EditorGroup(subs ...*ObjGroup) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "editor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) EditorUser(subs ...*ObjUser) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "editor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) DeletorGroup(subs ...*ObjGroup) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "deletor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) DeletorUser(subs ...*ObjUser) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "deletor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) SelectorGroup(subs ...*ObjGroup) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "selector",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) SelectorUser(subs ...*ObjUser) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "selector",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) ConnectorGroup(subs ...*ObjGroup) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "connector",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) ConnectorUser(subs ...*ObjUser) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "connector",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

type ObjWorkspace_build struct {
	Obj *v1.ObjectReference
	*Relationships
}

func Workspace_build(id string) *ObjWorkspace_build {
	o := &ObjWorkspace_build{
		Obj: &v1.ObjectReference{
			ObjectType: "workspace_build",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjWorkspace_build) Type() string {
	return "workspace_build"
}

func (obj *ObjWorkspace_build) Workspace(subs ...*ObjWorkspace) *ObjWorkspace_build {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

type ObjWorspace_app struct {
	Obj *v1.ObjectReference
	*Relationships
}

func Worspace_app(id string) *ObjWorspace_app {
	o := &ObjWorspace_app{
		Obj: &v1.ObjectReference{
			ObjectType: "worspace_app",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjWorspace_app) Type() string {
	return "worspace_app"
}

func (obj *ObjWorspace_app) Workspace(subs ...*ObjWorkspace) *ObjWorspace_app {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

type ObjWorkspace_agent struct {
	Obj *v1.ObjectReference
	*Relationships
}

func Workspace_agent(id string) *ObjWorkspace_agent {
	o := &ObjWorkspace_agent{
		Obj: &v1.ObjectReference{
			ObjectType: "workspace_agent",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjWorkspace_agent) Type() string {
	return "workspace_agent"
}

func (obj *ObjWorkspace_agent) Workspace(subs ...*ObjWorkspace) *ObjWorkspace_agent {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

type ObjWorkspace_resources struct {
	Obj *v1.ObjectReference
	*Relationships
}

func Workspace_resources(id string) *ObjWorkspace_resources {
	o := &ObjWorkspace_resources{
		Obj: &v1.ObjectReference{
			ObjectType: "workspace_resources",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjWorkspace_resources) Type() string {
	return "workspace_resources"
}

func (obj *ObjWorkspace_resources) Workspace(subs ...*ObjWorkspace) *ObjWorkspace_resources {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

type ObjTemplate struct {
	Obj *v1.ObjectReference
	*Relationships
}

func Template(id string) *ObjTemplate {
	o := &ObjTemplate{
		Obj: &v1.ObjectReference{
			ObjectType: "template",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjTemplate) Type() string {
	return "template"
}

func (obj *ObjTemplate) Owner(subs ...*ObjTeam) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "owner",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate) Workspace(subs ...*ObjWorkspace) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

type ObjTemplate_version struct {
	Obj *v1.ObjectReference
	*Relationships
}

func Template_version(id string) *ObjTemplate_version {
	o := &ObjTemplate_version{
		Obj: &v1.ObjectReference{
			ObjectType: "template_version",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjTemplate_version) Type() string {
	return "template_version"
}

func (obj *ObjTemplate_version) Template(subs ...*ObjTemplate) *ObjTemplate_version {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "template",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}
