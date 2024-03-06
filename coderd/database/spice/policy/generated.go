// Package relationships code generated. DO NOT EDIT.
package policy

import v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"

type ObjFile struct {
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func File(id string) *ObjFile {
	o := &ObjFile{
		Obj: &v1.ObjectReference{
			ObjectType: "file",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjFile) Type() string {
	return "file"
}

func (obj *ObjFile) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjFile) Template_versionTemplate_version(subs ...*ObjTemplate_version) *ObjFile {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_version",
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
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func Group(id string) *ObjGroup {
	o := &ObjGroup{
		Obj: &v1.ObjectReference{
			ObjectType: "group",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjGroup) Type() string {
	return "group"
}

func (obj *ObjGroup) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjGroup) MemberUser(subs ...*ObjUser) *ObjGroup {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
	obj.Relationships = append(obj.Relationships, v1.Relationship{
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

type ObjJob struct {
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func Job(id string) *ObjJob {
	o := &ObjJob{
		Obj: &v1.ObjectReference{
			ObjectType: "job",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjJob) Type() string {
	return "job"
}

func (obj *ObjJob) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjJob) Template_versionTemplate_version(subs ...*ObjTemplate_version) *ObjJob {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
			Resource: obj.Obj,
			Relation: "template_version",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjJob) Workspace_buildWorkspace_build(subs ...*ObjWorkspace_build) *ObjJob {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
			Resource: obj.Obj,
			Relation: "workspace_build",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

type ObjOrganization struct {
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func Organization(id string) *ObjOrganization {
	o := &ObjOrganization{
		Obj: &v1.ObjectReference{
			ObjectType: "organization",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjOrganization) Type() string {
	return "organization"
}

func (obj *ObjOrganization) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjOrganization) PlatformPlatform(subs ...*ObjPlatform) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) MemberGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) MemberUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Workspace_viewerGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Workspace_viewerUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Workspace_creatorGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Workspace_creatorUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Workspace_deletorGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Workspace_deletorUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Workspace_editorGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Workspace_editorUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Template_viewerGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Template_viewerUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Template_creatorGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Template_creatorUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Template_deletorGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Template_deletorUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Template_editorGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Template_editorUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Template_permission_managerGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Template_permission_managerUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Template_insights_viewerGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

func (obj *ObjOrganization) Template_insights_viewerUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

type ObjPlatform struct {
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func Platform(id string) *ObjPlatform {
	o := &ObjPlatform{
		Obj: &v1.ObjectReference{
			ObjectType: "platform",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjPlatform) Type() string {
	return "platform"
}

func (obj *ObjPlatform) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjPlatform) AdministratorUser(subs ...*ObjUser) *ObjPlatform {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

type ObjTemplate struct {
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func Template(id string) *ObjTemplate {
	o := &ObjTemplate{
		Obj: &v1.ObjectReference{
			ObjectType: "template",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjTemplate) Type() string {
	return "template"
}

func (obj *ObjTemplate) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjTemplate) OrganizationOrganization(subs ...*ObjOrganization) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
			Resource: obj.Obj,
			Relation: "organization",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate) WorkspaceWorkspace(subs ...*ObjWorkspace) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func Template_version(id string) *ObjTemplate_version {
	o := &ObjTemplate_version{
		Obj: &v1.ObjectReference{
			ObjectType: "template_version",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjTemplate_version) Type() string {
	return "template_version"
}

func (obj *ObjTemplate_version) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjTemplate_version) TemplateTemplate(subs ...*ObjTemplate) *ObjTemplate_version {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

type ObjUser struct {
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func User(id string) *ObjUser {
	o := &ObjUser{
		Obj: &v1.ObjectReference{
			ObjectType: "user",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjUser) Type() string {
	return "user"
}

func (obj *ObjUser) Object() *v1.ObjectReference {
	return obj.Obj
}

type ObjWorkspace struct {
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func Workspace(id string) *ObjWorkspace {
	o := &ObjWorkspace{
		Obj: &v1.ObjectReference{
			ObjectType: "workspace",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjWorkspace) Type() string {
	return "workspace"
}

func (obj *ObjWorkspace) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjWorkspace) OrganizationOrganization(subs ...*ObjOrganization) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
			Resource: obj.Obj,
			Relation: "organization",
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
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

type ObjWorkspace_agent struct {
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func Workspace_agent(id string) *ObjWorkspace_agent {
	o := &ObjWorkspace_agent{
		Obj: &v1.ObjectReference{
			ObjectType: "workspace_agent",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjWorkspace_agent) Type() string {
	return "workspace_agent"
}

func (obj *ObjWorkspace_agent) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjWorkspace_agent) WorkspaceWorkspace(subs ...*ObjWorkspace) *ObjWorkspace_agent {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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

type ObjWorkspace_build struct {
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func Workspace_build(id string) *ObjWorkspace_build {
	o := &ObjWorkspace_build{
		Obj: &v1.ObjectReference{
			ObjectType: "workspace_build",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjWorkspace_build) Type() string {
	return "workspace_build"
}

func (obj *ObjWorkspace_build) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjWorkspace_build) WorkspaceWorkspace(subs ...*ObjWorkspace) *ObjWorkspace_build {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func Workspace_resources(id string) *ObjWorkspace_resources {
	o := &ObjWorkspace_resources{
		Obj: &v1.ObjectReference{
			ObjectType: "workspace_resources",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjWorkspace_resources) Type() string {
	return "workspace_resources"
}

func (obj *ObjWorkspace_resources) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjWorkspace_resources) WorkspaceWorkspace(subs ...*ObjWorkspace) *ObjWorkspace_resources {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
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
