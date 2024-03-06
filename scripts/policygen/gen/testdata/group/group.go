// Package relationships code generated. DO NOT EDIT.
package policy

import v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"

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
