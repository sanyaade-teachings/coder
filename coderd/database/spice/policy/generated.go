// Package policy code generated. DO NOT EDIT.
package policy

import (
	"context"
	"fmt"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

// String is used to use string literals instead of uuids.
type String string

func (s String) String() string {
	return string(s)
}

type AuthzedObject interface {
	Object() *v1.ObjectReference
}

// PermissionCheck can be read as:
// Can 'subject' do 'permission' on 'object'?
type PermissionCheck struct {
	// Subject has an optional
	Subject    *v1.SubjectReference
	Permission string
	Obj        *v1.ObjectReference
}

// Builder contains all the saved relationships and permission checks during
// function calls that extend from it.
// This means you can use the builder to create a set of relationships to add
// to the graph and/or a set of permission checks to validate.
type Builder struct {
	// Relationships are new graph connections to be formed.
	// This will expand the capability/permissions.
	Relationships []v1.Relationship
	// PermissionChecks are the set of capabilities required.
	PermissionChecks []PermissionCheck
}

func New() *Builder {
	return &Builder{
		Relationships:    make([]v1.Relationship, 0),
		PermissionChecks: make([]PermissionCheck, 0),
	}
}

func (b *Builder) AddRelationship(r v1.Relationship) *Builder {
	b.Relationships = append(b.Relationships, r)
	return b
}

func (b *Builder) CheckPermission(subj AuthzedObject, permission string, on AuthzedObject) *Builder {
	b.PermissionChecks = append(b.PermissionChecks, PermissionCheck{
		Subject: &v1.SubjectReference{
			Object:           subj.Object(),
			OptionalRelation: "",
		},
		Permission: permission,
		Obj:        on.Object(),
	})
	return b
}

// SitePlatform is a custom method to add a standard site-wide platform.
func (b *Builder) SitePlatform() *ObjPlatform {
	return b.Platform(String("site-wide"))
}

type ObjFile struct {
	Obj     *v1.ObjectReference
	Builder *Builder
}

func (b *Builder) File(id fmt.Stringer) *ObjFile {
	o := &ObjFile{
		Obj: &v1.ObjectReference{
			ObjectType: "file",
			ObjectId:   id.String(),
		},
		Builder: b,
	}
	return o
}

func (obj *ObjFile) Object() *v1.ObjectReference {
	return obj.Obj
}

// Template_version schema.zed:220
// Relationship: file:<id>#template_version@template_version:<id>
func (obj *ObjFile) Template_version(subs ...*ObjTemplate_version) *ObjFile {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// CanView schema.zed:222
// Object: file:<id>
// Schema: permission view = template_version->view
func (obj *ObjFile) CanView(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "view", obj.Object()
}

type ObjGroup struct {
	Obj     *v1.ObjectReference
	Builder *Builder
}

func (b *Builder) Group(id fmt.Stringer) *ObjGroup {
	o := &ObjGroup{
		Obj: &v1.ObjectReference{
			ObjectType: "group",
			ObjectId:   id.String(),
		},
		Builder: b,
	}
	return o
}

func (obj *ObjGroup) Object() *v1.ObjectReference {
	return obj.Obj
}

// MemberUser schema.zed:17
// Relationship: group:<id>#member@user:<id>
func (obj *ObjGroup) MemberUser(subs ...*ObjUser) *ObjGroup {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// MemberGroup schema.zed:17
// Relationship: group:<id>#member@group:<id>#member
func (obj *ObjGroup) MemberGroup(subs ...*ObjGroup) *ObjGroup {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// MemberWildcard schema.zed:17
// Relationship: group:<id>#member@user:*
func (obj *ObjGroup) MemberWildcard() *ObjGroup {
	obj.Builder.AddRelationship(v1.Relationship{
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

// CanMembership schema.zed:21
// Object: group:<id>
func (obj *ObjGroup) CanMembership(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "membership", obj.Object()
}

type ObjJob struct {
	Obj     *v1.ObjectReference
	Builder *Builder
}

func (b *Builder) Job(id fmt.Stringer) *ObjJob {
	o := &ObjJob{
		Obj: &v1.ObjectReference{
			ObjectType: "job",
			ObjectId:   id.String(),
		},
		Builder: b,
	}
	return o
}

func (obj *ObjJob) Object() *v1.ObjectReference {
	return obj.Obj
}

// Template_version schema.zed:229
// Relationship: job:<id>#template_version@template_version:<id>
func (obj *ObjJob) Template_version(subs ...*ObjTemplate_version) *ObjJob {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Workspace_build schema.zed:230
// Relationship: job:<id>#workspace_build@workspace_build:<id>
func (obj *ObjJob) Workspace_build(subs ...*ObjWorkspace_build) *ObjJob {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// CanView schema.zed:233
// Object: job:<id>
func (obj *ObjJob) CanView(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "view", obj.Object()
}

type ObjOrganization struct {
	Obj     *v1.ObjectReference
	Builder *Builder
}

func (b *Builder) Organization(id fmt.Stringer) *ObjOrganization {
	o := &ObjOrganization{
		Obj: &v1.ObjectReference{
			ObjectType: "organization",
			ObjectId:   id.String(),
		},
		Builder: b,
	}
	return o
}

func (obj *ObjOrganization) Object() *v1.ObjectReference {
	return obj.Obj
}

// Platform schema.zed:43
// Relationship: organization:<id>#platform@platform:<id>
func (obj *ObjOrganization) Platform(subs ...*ObjPlatform) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// MemberGroup schema.zed:49
// Relationship: organization:<id>#member@group:<id>#membership
func (obj *ObjOrganization) MemberGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// MemberUser schema.zed:49
// Relationship: organization:<id>#member@user:<id>
func (obj *ObjOrganization) MemberUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Default_permissionsGroup schema.zed:53
// Relationship: organization:<id>#default_permissions@group:<id>#membership
func (obj *ObjOrganization) Default_permissionsGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
			Resource: obj.Obj,
			Relation: "default_permissions",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

// Default_permissionsUser schema.zed:53
// Relationship: organization:<id>#default_permissions@user:<id>
func (obj *ObjOrganization) Default_permissionsUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
			Resource: obj.Obj,
			Relation: "default_permissions",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

// Workspace_viewerGroup schema.zed:61
// Relationship: organization:<id>#workspace_viewer@group:<id>#membership
func (obj *ObjOrganization) Workspace_viewerGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Workspace_viewerUser schema.zed:61
// Relationship: organization:<id>#workspace_viewer@user:<id>
func (obj *ObjOrganization) Workspace_viewerUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Workspace_creatorGroup schema.zed:64
// Relationship: organization:<id>#workspace_creator@group:<id>#membership
func (obj *ObjOrganization) Workspace_creatorGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Workspace_creatorUser schema.zed:64
// Relationship: organization:<id>#workspace_creator@user:<id>
func (obj *ObjOrganization) Workspace_creatorUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Workspace_deletorGroup schema.zed:66
// Relationship: organization:<id>#workspace_deletor@group:<id>#membership
func (obj *ObjOrganization) Workspace_deletorGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Workspace_deletorUser schema.zed:66
// Relationship: organization:<id>#workspace_deletor@user:<id>
func (obj *ObjOrganization) Workspace_deletorUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Workspace_editorGroup schema.zed:69
// Relationship: organization:<id>#workspace_editor@group:<id>#membership
func (obj *ObjOrganization) Workspace_editorGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Workspace_editorUser schema.zed:69
// Relationship: organization:<id>#workspace_editor@user:<id>
func (obj *ObjOrganization) Workspace_editorUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Template_viewerGroup schema.zed:77
// Relationship: organization:<id>#template_viewer@group:<id>#membership
func (obj *ObjOrganization) Template_viewerGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Template_viewerUser schema.zed:77
// Relationship: organization:<id>#template_viewer@user:<id>
func (obj *ObjOrganization) Template_viewerUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Template_creatorGroup schema.zed:78
// Relationship: organization:<id>#template_creator@group:<id>#membership
func (obj *ObjOrganization) Template_creatorGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Template_creatorUser schema.zed:78
// Relationship: organization:<id>#template_creator@user:<id>
func (obj *ObjOrganization) Template_creatorUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Template_deletorGroup schema.zed:79
// Relationship: organization:<id>#template_deletor@group:<id>#membership
func (obj *ObjOrganization) Template_deletorGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Template_deletorUser schema.zed:79
// Relationship: organization:<id>#template_deletor@user:<id>
func (obj *ObjOrganization) Template_deletorUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Template_editorGroup schema.zed:80
// Relationship: organization:<id>#template_editor@group:<id>#membership
func (obj *ObjOrganization) Template_editorGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Template_editorUser schema.zed:80
// Relationship: organization:<id>#template_editor@user:<id>
func (obj *ObjOrganization) Template_editorUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Template_permission_managerGroup schema.zed:81
// Relationship: organization:<id>#template_permission_manager@group:<id>#membership
func (obj *ObjOrganization) Template_permission_managerGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Template_permission_managerUser schema.zed:81
// Relationship: organization:<id>#template_permission_manager@user:<id>
func (obj *ObjOrganization) Template_permission_managerUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Template_insights_viewerGroup schema.zed:82
// Relationship: organization:<id>#template_insights_viewer@group:<id>#membership
func (obj *ObjOrganization) Template_insights_viewerGroup(subs ...*ObjGroup) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Template_insights_viewerUser schema.zed:82
// Relationship: organization:<id>#template_insights_viewer@user:<id>
func (obj *ObjOrganization) Template_insights_viewerUser(subs ...*ObjUser) *ObjOrganization {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// CanMembership schema.zed:92
// Object: organization:<id>
func (obj *ObjOrganization) CanMembership(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "membership", obj.Object()
}

// CanCreate_org_member schema.zed:95
// Object: organization:<id>
// Schema: permission create_org_member = platform->create_user
func (obj *ObjOrganization) CanCreate_org_member(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "create_org_member", obj.Object()
}

// CanView_workspaces schema.zed:102
// Object: organization:<id>
func (obj *ObjOrganization) CanView_workspaces(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "view_workspaces", obj.Object()
}

// CanEdit_workspaces schema.zed:103
// Object: organization:<id>
// Schema: permission edit_workspaces = platform->super_admin + workspace_editor
func (obj *ObjOrganization) CanEdit_workspaces(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "edit_workspaces", obj.Object()
}

// CanSelect_workspace_version schema.zed:104
// Object: organization:<id>
// Schema: permission select_workspace_version = platform->super_admin
func (obj *ObjOrganization) CanSelect_workspace_version(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "select_workspace_version", obj.Object()
}

// CanDelete_workspaces schema.zed:105
// Object: organization:<id>
// Schema: permission delete_workspaces = platform->super_admin + workspace_deletor
func (obj *ObjOrganization) CanDelete_workspaces(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "delete_workspaces", obj.Object()
}

// CanCreate_workspace schema.zed:108
// Object: organization:<id>
func (obj *ObjOrganization) CanCreate_workspace(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "create_workspace", obj.Object()
}

// CanView_templates schema.zed:114
// Object: organization:<id>
func (obj *ObjOrganization) CanView_templates(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "view_templates", obj.Object()
}

// CanView_template_insights schema.zed:115
// Object: organization:<id>
// Schema: permission view_template_insights = platform->super_admin + template_insights_viewer
func (obj *ObjOrganization) CanView_template_insights(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "view_template_insights", obj.Object()
}

// CanEdit_templates schema.zed:116
// Object: organization:<id>
// Schema: permission edit_templates = platform->super_admin + template_editor
func (obj *ObjOrganization) CanEdit_templates(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "edit_templates", obj.Object()
}

// CanDelete_templates schema.zed:117
// Object: organization:<id>
// Schema: permission delete_templates = platform->super_admin + template_deletor
func (obj *ObjOrganization) CanDelete_templates(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "delete_templates", obj.Object()
}

// CanManage_template_permissions schema.zed:118
// Object: organization:<id>
// Schema: permission manage_template_permissions = platform->super_admin + template_permission_manager
func (obj *ObjOrganization) CanManage_template_permissions(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "manage_template_permissions", obj.Object()
}

// CanCreate_template schema.zed:120
// Object: organization:<id>
func (obj *ObjOrganization) CanCreate_template(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "create_template", obj.Object()
}

// CanCreate_template_version schema.zed:121
// Object: organization:<id>
// Schema: permission create_template_version = create_template
func (obj *ObjOrganization) CanCreate_template_version(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "create_template_version", obj.Object()
}

// CanCreate_file schema.zed:122
// Object: organization:<id>
// Schema: permission create_file = create_template
func (obj *ObjOrganization) CanCreate_file(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "create_file", obj.Object()
}

type ObjPlatform struct {
	Obj     *v1.ObjectReference
	Builder *Builder
}

func (b *Builder) Platform(id fmt.Stringer) *ObjPlatform {
	o := &ObjPlatform{
		Obj: &v1.ObjectReference{
			ObjectType: "platform",
			ObjectId:   id.String(),
		},
		Builder: b,
	}
	return o
}

func (obj *ObjPlatform) Object() *v1.ObjectReference {
	return obj.Obj
}

// Administrator schema.zed:29
// Relationship: platform:<id>#administrator@user:<id>
func (obj *ObjPlatform) Administrator(subs ...*ObjUser) *ObjPlatform {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// User_adminUser schema.zed:30
// Relationship: platform:<id>#user_admin@user:<id>
func (obj *ObjPlatform) User_adminUser(subs ...*ObjUser) *ObjPlatform {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
			Resource: obj.Obj,
			Relation: "user_admin",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

// User_adminGroup schema.zed:30
// Relationship: platform:<id>#user_admin@group:<id>#membership
func (obj *ObjPlatform) User_adminGroup(subs ...*ObjGroup) *ObjPlatform {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
			Resource: obj.Obj,
			Relation: "user_admin",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

// CanSuper_admin schema.zed:34
// Object: platform:<id>
func (obj *ObjPlatform) CanSuper_admin(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "super_admin", obj.Object()
}

// CanCreate_user schema.zed:35
// Object: platform:<id>
// Schema: permission create_user = user_admin + super_admin
func (obj *ObjPlatform) CanCreate_user(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "create_user", obj.Object()
}

// CanCreate_organization schema.zed:36
// Object: platform:<id>
// Schema: permission create_organization = super_admin
func (obj *ObjPlatform) CanCreate_organization(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "create_organization", obj.Object()
}

type ObjTemplate struct {
	Obj     *v1.ObjectReference
	Builder *Builder
}

func (b *Builder) Template(id fmt.Stringer) *ObjTemplate {
	o := &ObjTemplate{
		Obj: &v1.ObjectReference{
			ObjectType: "template",
			ObjectId:   id.String(),
		},
		Builder: b,
	}
	return o
}

func (obj *ObjTemplate) Object() *v1.ObjectReference {
	return obj.Obj
}

// Organization schema.zed:192
// Relationship: template:<id>#organization@organization:<id>
func (obj *ObjTemplate) Organization(subs ...*ObjOrganization) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// Workspace schema.zed:197
// Relationship: template:<id>#workspace@workspace:<id>
func (obj *ObjTemplate) Workspace(subs ...*ObjWorkspace) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// CanView schema.zed:199
// Object: template:<id>
// Schema: permission view = organization->template_viewer + workspace->view
func (obj *ObjTemplate) CanView(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "view", obj.Object()
}

// CanView_insights schema.zed:200
// Object: template:<id>
// Schema: permission view_insights = organization->view_template_insights
func (obj *ObjTemplate) CanView_insights(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "view_insights", obj.Object()
}

// CanEdit schema.zed:202
// Object: template:<id>
func (obj *ObjTemplate) CanEdit(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "edit", obj.Object()
}

// CanDelete schema.zed:203
// Object: template:<id>
// Schema: permission delete = organization->delete_templates
func (obj *ObjTemplate) CanDelete(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "delete", obj.Object()
}

// CanEdit_pemissions schema.zed:204
// Object: template:<id>
// Schema: permission edit_pemissions = organization->manage_template_permissions
func (obj *ObjTemplate) CanEdit_pemissions(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "edit_pemissions", obj.Object()
}

// CanUse schema.zed:207
// Object: template:<id>
func (obj *ObjTemplate) CanUse(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "use", obj.Object()
}

// CanWorkspace_view schema.zed:210
// Object: template:<id>
func (obj *ObjTemplate) CanWorkspace_view(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "workspace_view", obj.Object()
}

type ObjTemplate_version struct {
	Obj     *v1.ObjectReference
	Builder *Builder
}

func (b *Builder) Template_version(id fmt.Stringer) *ObjTemplate_version {
	o := &ObjTemplate_version{
		Obj: &v1.ObjectReference{
			ObjectType: "template_version",
			ObjectId:   id.String(),
		},
		Builder: b,
	}
	return o
}

func (obj *ObjTemplate_version) Object() *v1.ObjectReference {
	return obj.Obj
}

// Template schema.zed:214
// Relationship: template_version:<id>#template@template:<id>
func (obj *ObjTemplate_version) Template(subs ...*ObjTemplate) *ObjTemplate_version {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// CanView schema.zed:216
// Object: template_version:<id>
// Schema: permission view = template->view
func (obj *ObjTemplate_version) CanView(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "view", obj.Object()
}

type ObjUser struct {
	Obj     *v1.ObjectReference
	Builder *Builder
}

func (b *Builder) User(id fmt.Stringer) *ObjUser {
	o := &ObjUser{
		Obj: &v1.ObjectReference{
			ObjectType: "user",
			ObjectId:   id.String(),
		},
		Builder: b,
	}
	return o
}

func (obj *ObjUser) Object() *v1.ObjectReference {
	return obj.Obj
}

type ObjWorkspace struct {
	Obj     *v1.ObjectReference
	Builder *Builder
}

func (b *Builder) Workspace(id fmt.Stringer) *ObjWorkspace {
	o := &ObjWorkspace{
		Obj: &v1.ObjectReference{
			ObjectType: "workspace",
			ObjectId:   id.String(),
		},
		Builder: b,
	}
	return o
}

func (obj *ObjWorkspace) Object() *v1.ObjectReference {
	return obj.Obj
}

// Organization schema.zed:133
// Relationship: workspace:<id>#organization@organization:<id>
func (obj *ObjWorkspace) Organization(subs ...*ObjOrganization) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// ViewerGroup schema.zed:135
// Relationship: workspace:<id>#viewer@group:<id>#membership
func (obj *ObjWorkspace) ViewerGroup(subs ...*ObjGroup) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// ViewerUser schema.zed:135
// Relationship: workspace:<id>#viewer@user:<id>
func (obj *ObjWorkspace) ViewerUser(subs ...*ObjUser) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// EditorGroup schema.zed:136
// Relationship: workspace:<id>#editor@group:<id>#membership
func (obj *ObjWorkspace) EditorGroup(subs ...*ObjGroup) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// EditorUser schema.zed:136
// Relationship: workspace:<id>#editor@user:<id>
func (obj *ObjWorkspace) EditorUser(subs ...*ObjUser) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// DeletorGroup schema.zed:137
// Relationship: workspace:<id>#deletor@group:<id>#membership
func (obj *ObjWorkspace) DeletorGroup(subs ...*ObjGroup) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// DeletorUser schema.zed:137
// Relationship: workspace:<id>#deletor@user:<id>
func (obj *ObjWorkspace) DeletorUser(subs ...*ObjUser) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// SelectorGroup schema.zed:138
// Relationship: workspace:<id>#selector@group:<id>#membership
func (obj *ObjWorkspace) SelectorGroup(subs ...*ObjGroup) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// SelectorUser schema.zed:138
// Relationship: workspace:<id>#selector@user:<id>
func (obj *ObjWorkspace) SelectorUser(subs ...*ObjUser) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// ConnectorGroup schema.zed:139
// Relationship: workspace:<id>#connector@group:<id>#membership
func (obj *ObjWorkspace) ConnectorGroup(subs ...*ObjGroup) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// ConnectorUser schema.zed:139
// Relationship: workspace:<id>#connector@user:<id>
func (obj *ObjWorkspace) ConnectorUser(subs ...*ObjUser) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// For_userGroup schema.zed:144
// Relationship: workspace:<id>#for_user@group:<id>#membership
func (obj *ObjWorkspace) For_userGroup(subs ...*ObjGroup) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
			Resource: obj.Obj,
			Relation: "for_user",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

// For_userUser schema.zed:144
// Relationship: workspace:<id>#for_user@user:<id>
func (obj *ObjWorkspace) For_userUser(subs ...*ObjUser) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
			Resource: obj.Obj,
			Relation: "for_user",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

// CanWorkspace_owner schema.zed:148
// Object: workspace:<id>
func (obj *ObjWorkspace) CanWorkspace_owner(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "workspace_owner", obj.Object()
}

// CanView schema.zed:152
// Object: workspace:<id>
func (obj *ObjWorkspace) CanView(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "view", obj.Object()
}

// CanEdit schema.zed:158
// Object: workspace:<id>
// Schema: permission edit = organization->edit_workspaces + editor + workspace_owner
func (obj *ObjWorkspace) CanEdit(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "edit", obj.Object()
}

// CanDelete schema.zed:159
// Object: workspace:<id>
// Schema: permission delete = organization->delete_workspaces + deletor + workspace_owner
func (obj *ObjWorkspace) CanDelete(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "delete", obj.Object()
}

// CanSelect_template_version schema.zed:161
// Object: workspace:<id>
func (obj *ObjWorkspace) CanSelect_template_version(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "select_template_version", obj.Object()
}

// CanSsh schema.zed:162
// Object: workspace:<id>
// Schema: permission ssh = connector + workspace_owner
func (obj *ObjWorkspace) CanSsh(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "ssh", obj.Object()
}

type ObjWorkspace_agent struct {
	Obj     *v1.ObjectReference
	Builder *Builder
}

func (b *Builder) Workspace_agent(id fmt.Stringer) *ObjWorkspace_agent {
	o := &ObjWorkspace_agent{
		Obj: &v1.ObjectReference{
			ObjectType: "workspace_agent",
			ObjectId:   id.String(),
		},
		Builder: b,
	}
	return o
}

func (obj *ObjWorkspace_agent) Object() *v1.ObjectReference {
	return obj.Obj
}

// Workspace schema.zed:176
// Relationship: workspace_agent:<id>#workspace@workspace:<id>
func (obj *ObjWorkspace_agent) Workspace(subs ...*ObjWorkspace) *ObjWorkspace_agent {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// CanView schema.zed:178
// Object: workspace_agent:<id>
// Schema: permission view = workspace->view
func (obj *ObjWorkspace_agent) CanView(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "view", obj.Object()
}

type ObjWorkspace_build struct {
	Obj     *v1.ObjectReference
	Builder *Builder
}

func (b *Builder) Workspace_build(id fmt.Stringer) *ObjWorkspace_build {
	o := &ObjWorkspace_build{
		Obj: &v1.ObjectReference{
			ObjectType: "workspace_build",
			ObjectId:   id.String(),
		},
		Builder: b,
	}
	return o
}

func (obj *ObjWorkspace_build) Object() *v1.ObjectReference {
	return obj.Obj
}

// Workspace schema.zed:167
// Relationship: workspace_build:<id>#workspace@workspace:<id>
func (obj *ObjWorkspace_build) Workspace(subs ...*ObjWorkspace) *ObjWorkspace_build {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// CanView schema.zed:172
// Object: workspace_build:<id>
func (obj *ObjWorkspace_build) CanView(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "view", obj.Object()
}

type ObjWorkspace_resources struct {
	Obj     *v1.ObjectReference
	Builder *Builder
}

func (b *Builder) Workspace_resources(id fmt.Stringer) *ObjWorkspace_resources {
	o := &ObjWorkspace_resources{
		Obj: &v1.ObjectReference{
			ObjectType: "workspace_resources",
			ObjectId:   id.String(),
		},
		Builder: b,
	}
	return o
}

func (obj *ObjWorkspace_resources) Object() *v1.ObjectReference {
	return obj.Obj
}

// Workspace schema.zed:182
// Relationship: workspace_resources:<id>#workspace@workspace:<id>
func (obj *ObjWorkspace_resources) Workspace(subs ...*ObjWorkspace) *ObjWorkspace_resources {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// CanView schema.zed:184
// Object: workspace_resources:<id>
// Schema: permission view = workspace->view
func (obj *ObjWorkspace_resources) CanView(ctx context.Context) (context.Context, string, *v1.ObjectReference) {
	return ctx, "view", obj.Object()
}
