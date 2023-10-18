// Package relationships code generated. DO NOT EDIT.
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

func (obj *ObjPlatform) Object() *v1.ObjectReference {
	return obj.Obj
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

func (obj *ObjPlatform) ValidateSuper_admin() *ObjPlatform {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "super_admin",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjPlatform) CanSuper_adminBy(subs ...ObjectWithRelationships) *ObjPlatform {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "super_admin",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjPlatform) CannotSuper_adminBy(subs ...ObjectWithRelationships) *ObjPlatform {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "super_admin",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
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

func (obj *ObjTeam) Object() *v1.ObjectReference {
	return obj.Obj
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

func (obj *ObjTeam) Parent(subs ...*ObjTeam) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "parent",
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

func (obj *ObjTeam) Provisioner_viewerGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "provisioner_viewer",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Provisioner_viewerUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "provisioner_viewer",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Provisioner_creatorGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "provisioner_creator",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Provisioner_creatorUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "provisioner_creator",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Provisioner_deletorGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "provisioner_deletor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Provisioner_deletorUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "provisioner_deletor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Provisioner_editorGroup(subs ...*ObjGroup) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "provisioner_editor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "membership",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) Provisioner_editorUser(subs ...*ObjUser) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
			Resource: obj.Obj,
			Relation: "provisioner_editor",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) ValidateDirect_membership() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "direct_membership",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateView_workspaces() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view_workspaces",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateEdit_workspaces() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "edit_workspaces",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateSelect_workspace_version() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "select_workspace_version",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateDelete_workspaces() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "delete_workspaces",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateConnect_workspaces() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "connect_workspaces",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateCreate_workspace() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "create_workspace",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateView_templates() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view_templates",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateView_template_insights() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view_template_insights",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateEdit_templates() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "edit_templates",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateDelete_templates() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "delete_templates",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateManage_template_permissions() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "manage_template_permissions",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateCreate_template() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "create_template",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateCreate_template_version() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "create_template_version",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateCreate_file() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "create_file",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateView_provisioners() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view_provisioners",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateEdit_provisioners() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "edit_provisioners",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateDelete_provisioners() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "delete_provisioners",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) ValidateCreate_provisioners() *ObjTeam {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "create_provisioners",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTeam) CanDirect_membershipBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "direct_membership",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotDirect_membershipBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "direct_membership",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanView_workspacesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view_workspaces",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotView_workspacesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view_workspaces",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanEdit_workspacesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "edit_workspaces",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotEdit_workspacesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "edit_workspaces",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanSelect_workspace_versionBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "select_workspace_version",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotSelect_workspace_versionBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "select_workspace_version",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanDelete_workspacesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "delete_workspaces",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotDelete_workspacesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "delete_workspaces",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanConnect_workspacesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "connect_workspaces",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotConnect_workspacesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "connect_workspaces",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanCreate_workspaceBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "create_workspace",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotCreate_workspaceBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "create_workspace",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanView_templatesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view_templates",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotView_templatesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view_templates",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanView_template_insightsBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view_template_insights",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotView_template_insightsBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view_template_insights",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanEdit_templatesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "edit_templates",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotEdit_templatesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "edit_templates",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanDelete_templatesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "delete_templates",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotDelete_templatesBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "delete_templates",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanManage_template_permissionsBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "manage_template_permissions",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotManage_template_permissionsBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "manage_template_permissions",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanCreate_templateBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "create_template",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotCreate_templateBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "create_template",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanCreate_template_versionBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "create_template_version",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotCreate_template_versionBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "create_template_version",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanCreate_fileBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "create_file",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotCreate_fileBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "create_file",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanView_provisionersBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view_provisioners",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotView_provisionersBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view_provisioners",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanEdit_provisionersBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "edit_provisioners",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotEdit_provisionersBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "edit_provisioners",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanDelete_provisionersBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "delete_provisioners",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotDelete_provisionersBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "delete_provisioners",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CanCreate_provisionersBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "create_provisioners",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTeam) CannotCreate_provisionersBy(subs ...ObjectWithRelationships) *ObjTeam {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "create_provisioners",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
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

func (obj *ObjGroup) Object() *v1.ObjectReference {
	return obj.Obj
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

func (obj *ObjGroup) ValidateMembership() *ObjGroup {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "membership",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjGroup) CanMembershipBy(subs ...ObjectWithRelationships) *ObjGroup {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "membership",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjGroup) CannotMembershipBy(subs ...ObjectWithRelationships) *ObjGroup {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "membership",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
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

func (obj *ObjUser) Object() *v1.ObjectReference {
	return obj.Obj
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

func (obj *ObjWorkspace) Type() string {
	return "workspace"
}

func (obj *ObjWorkspace) Object() *v1.ObjectReference {
	return obj.Obj
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

func (obj *ObjWorkspace) ValidateView() *ObjWorkspace {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjWorkspace) ValidateEdit() *ObjWorkspace {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "edit",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjWorkspace) ValidateDelete() *ObjWorkspace {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "delete",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjWorkspace) ValidateSelect_template_version() *ObjWorkspace {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "select_template_version",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjWorkspace) ValidateSsh() *ObjWorkspace {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "ssh",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjWorkspace) CanViewBy(subs ...ObjectWithRelationships) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) CannotViewBy(subs ...ObjectWithRelationships) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) CanEditBy(subs ...ObjectWithRelationships) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "edit",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) CannotEditBy(subs ...ObjectWithRelationships) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "edit",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) CanDeleteBy(subs ...ObjectWithRelationships) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "delete",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) CannotDeleteBy(subs ...ObjectWithRelationships) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "delete",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) CanSelect_template_versionBy(subs ...ObjectWithRelationships) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "select_template_version",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) CannotSelect_template_versionBy(subs ...ObjectWithRelationships) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "select_template_version",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) CanSshBy(subs ...ObjectWithRelationships) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "ssh",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace) CannotSshBy(subs ...ObjectWithRelationships) *ObjWorkspace {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "ssh",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
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

func (obj *ObjWorkspace_build) Object() *v1.ObjectReference {
	return obj.Obj
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

func (obj *ObjWorkspace_build) ValidateView() *ObjWorkspace_build {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjWorkspace_build) CanViewBy(subs ...ObjectWithRelationships) *ObjWorkspace_build {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace_build) CannotViewBy(subs ...ObjectWithRelationships) *ObjWorkspace_build {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
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

func (obj *ObjWorspace_app) Object() *v1.ObjectReference {
	return obj.Obj
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

func (obj *ObjWorspace_app) ValidateView() *ObjWorspace_app {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjWorspace_app) ValidateConnect() *ObjWorspace_app {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "connect",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjWorspace_app) CanViewBy(subs ...ObjectWithRelationships) *ObjWorspace_app {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorspace_app) CannotViewBy(subs ...ObjectWithRelationships) *ObjWorspace_app {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorspace_app) CanConnectBy(subs ...ObjectWithRelationships) *ObjWorspace_app {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "connect",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorspace_app) CannotConnectBy(subs ...ObjectWithRelationships) *ObjWorspace_app {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "connect",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
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

func (obj *ObjWorkspace_agent) Object() *v1.ObjectReference {
	return obj.Obj
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

func (obj *ObjWorkspace_agent) ValidateView() *ObjWorkspace_agent {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjWorkspace_agent) CanViewBy(subs ...ObjectWithRelationships) *ObjWorkspace_agent {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace_agent) CannotViewBy(subs ...ObjectWithRelationships) *ObjWorkspace_agent {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
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

func (obj *ObjWorkspace_resources) Object() *v1.ObjectReference {
	return obj.Obj
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

func (obj *ObjWorkspace_resources) ValidateView() *ObjWorkspace_resources {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjWorkspace_resources) CanViewBy(subs ...ObjectWithRelationships) *ObjWorkspace_resources {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjWorkspace_resources) CannotViewBy(subs ...ObjectWithRelationships) *ObjWorkspace_resources {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
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

func (obj *ObjTemplate) Object() *v1.ObjectReference {
	return obj.Obj
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

func (obj *ObjTemplate) ValidateView() *ObjTemplate {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTemplate) ValidateView_insights() *ObjTemplate {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view_insights",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTemplate) ValidateEdit() *ObjTemplate {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "edit",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTemplate) ValidateDelete() *ObjTemplate {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "delete",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTemplate) ValidateEdit_pemissions() *ObjTemplate {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "edit_pemissions",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTemplate) ValidateUse() *ObjTemplate {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "use",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTemplate) CanViewBy(subs ...ObjectWithRelationships) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate) CannotViewBy(subs ...ObjectWithRelationships) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate) CanView_insightsBy(subs ...ObjectWithRelationships) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view_insights",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate) CannotView_insightsBy(subs ...ObjectWithRelationships) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view_insights",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate) CanEditBy(subs ...ObjectWithRelationships) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "edit",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate) CannotEditBy(subs ...ObjectWithRelationships) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "edit",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate) CanDeleteBy(subs ...ObjectWithRelationships) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "delete",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate) CannotDeleteBy(subs ...ObjectWithRelationships) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "delete",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate) CanEdit_pemissionsBy(subs ...ObjectWithRelationships) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "edit_pemissions",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate) CannotEdit_pemissionsBy(subs ...ObjectWithRelationships) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "edit_pemissions",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate) CanUseBy(subs ...ObjectWithRelationships) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "use",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate) CannotUseBy(subs ...ObjectWithRelationships) *ObjTemplate {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "use",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
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

func (obj *ObjTemplate_version) Object() *v1.ObjectReference {
	return obj.Obj
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

func (obj *ObjTemplate_version) ValidateView() *ObjTemplate_version {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjTemplate_version) CanViewBy(subs ...ObjectWithRelationships) *ObjTemplate_version {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjTemplate_version) CannotViewBy(subs ...ObjectWithRelationships) *ObjTemplate_version {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

type ObjFile struct {
	Obj *v1.ObjectReference
	*Relationships
}

func File(id string) *ObjFile {
	o := &ObjFile{
		Obj: &v1.ObjectReference{
			ObjectType: "file",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjFile) Type() string {
	return "file"
}

func (obj *ObjFile) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjFile) Template_version(subs ...*ObjTemplate_version) *ObjFile {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
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

func (obj *ObjFile) ValidateView() *ObjFile {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjFile) CanViewBy(subs ...ObjectWithRelationships) *ObjFile {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjFile) CannotViewBy(subs ...ObjectWithRelationships) *ObjFile {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

type ObjProvisioner struct {
	Obj *v1.ObjectReference
	*Relationships
}

func Provisioner(id string) *ObjProvisioner {
	o := &ObjProvisioner{
		Obj: &v1.ObjectReference{
			ObjectType: "provisioner",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjProvisioner) Type() string {
	return "provisioner"
}

func (obj *ObjProvisioner) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjProvisioner) Owner(subs ...*ObjTeam) *ObjProvisioner {
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

func (obj *ObjProvisioner) ValidateView() *ObjProvisioner {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjProvisioner) ValidateUse() *ObjProvisioner {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "use",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjProvisioner) CanViewBy(subs ...ObjectWithRelationships) *ObjProvisioner {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjProvisioner) CannotViewBy(subs ...ObjectWithRelationships) *ObjProvisioner {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjProvisioner) CanUseBy(subs ...ObjectWithRelationships) *ObjProvisioner {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "use",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjProvisioner) CannotUseBy(subs ...ObjectWithRelationships) *ObjProvisioner {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "use",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

type ObjJob struct {
	Obj *v1.ObjectReference
	*Relationships
}

func Job(id string) *ObjJob {
	o := &ObjJob{
		Obj: &v1.ObjectReference{
			ObjectType: "job",
			ObjectId:   id,
		},
		Relationships: NewRelationships(),
	}
	allObjects = append(allObjects, o)
	return o
}

func (obj *ObjJob) Type() string {
	return "job"
}

func (obj *ObjJob) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjJob) Template_version(subs ...*ObjTemplate_version) *ObjJob {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
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

func (obj *ObjJob) Workspace_build(subs ...*ObjWorkspace_build) *ObjJob {
	for i := range subs {
		sub := subs[i]
		obj.AddRelation(v1.Relationship{
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

func (obj *ObjJob) ValidateView() *ObjJob {
	obj.AddValidation(v1.Relationship{
		Resource:       obj.Obj,
		Relation:       "view",
		OptionalCaveat: nil,
	})
	return obj
}

func (obj *ObjJob) CanViewBy(subs ...ObjectWithRelationships) *ObjJob {
	for i := range subs {
		sub := subs[i]
		obj.AssertTrue(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjJob) CannotViewBy(subs ...ObjectWithRelationships) *ObjJob {
	for i := range subs {
		sub := subs[i]
		obj.AssertFalse(v1.Relationship{
			Resource: obj.Obj,
			Relation: "view",
			Subject: &v1.SubjectReference{
				Object:           sub.Object(),
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

