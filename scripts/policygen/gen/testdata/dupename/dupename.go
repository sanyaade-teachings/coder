// Package relationships code generated. DO NOT EDIT.
package policy

import v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"

type ObjOther struct {
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func Other(id string) *ObjOther {
	o := &ObjOther{
		Obj: &v1.ObjectReference{
			ObjectType: "other",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjOther) Type() string {
	return "other"
}

type ObjPerson struct {
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func Person(id string) *ObjPerson {
	o := &ObjPerson{
		Obj: &v1.ObjectReference{
			ObjectType: "person",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjPerson) Type() string {
	return "person"
}

func (obj *ObjPerson) UserUser(subs ...*ObjUser) *ObjPerson {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
			Resource: obj.Obj,
			Relation: "user",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjPerson) UserOther(subs ...*ObjOther) *ObjPerson {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
			Resource: obj.Obj,
			Relation: "user",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjPerson) TestUser(subs ...*ObjUser) *ObjPerson {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
			Resource: obj.Obj,
			Relation: "test",
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
