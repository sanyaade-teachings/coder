// Package relationships code generated. DO NOT EDIT.
package relationships

import v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"

type ObjResource struct {
	Obj           *v1.ObjectReference
	Relationships []v1.Relationship
}

func Resource(id string) *ObjResource {
	o := &ObjResource{
		Obj: &v1.ObjectReference{
			ObjectType: "resource",
			ObjectId:   id,
		},
		Relationships: []v1.Relationship{},
	}
	return o
}

func (obj *ObjResource) Type() string {
	return "resource"
}

func (obj *ObjResource) Object() *v1.ObjectReference {
	return obj.Obj
}

func (obj *ObjResource) WriterUser(subs ...*ObjUser) *ObjResource {
	for i := range subs {
		sub := subs[i]
		obj.Relationships = append(obj.Relationships, v1.Relationship{
			Resource: obj.Obj,
			Relation: "writer",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

func (obj *ObjResource) ViewerUser(subs ...*ObjUser) *ObjResource {
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
