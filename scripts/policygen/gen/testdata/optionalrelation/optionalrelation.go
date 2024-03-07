// Package policy code generated. DO NOT EDIT.
package policy

import (
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

// Folder optionalrelation.zed:10
// Relationship: file:<id>#folder@folder:<id>
func (obj *ObjFile) Folder(subs ...*ObjFolder) *ObjFile {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
			Resource: obj.Obj,
			Relation: "folder",
			Subject: &v1.SubjectReference{
				Object:           sub.Obj,
				OptionalRelation: "",
			},
			OptionalCaveat: nil,
		})
	}
	return obj
}

// CanReadBy optionalrelation.zed:12
// Object: file:<id>
// Schema: permission read = folder->read
func (obj *ObjFile) CanReadBy(subs ...AuthzedObject) *ObjFile {
	for i := range subs {
		sub := subs[i]
		obj.Builder.CheckPermission(sub, "read", obj)
	}
	return obj
}

type ObjFolder struct {
	Obj     *v1.ObjectReference
	Builder *Builder
}

func (b *Builder) Folder(id fmt.Stringer) *ObjFolder {
	o := &ObjFolder{
		Obj: &v1.ObjectReference{
			ObjectType: "folder",
			ObjectId:   id.String(),
		},
		Builder: b,
	}
	return o
}

func (obj *ObjFolder) Object() *v1.ObjectReference {
	return obj.Obj
}

// Owner optionalrelation.zed:4
// Relationship: folder:<id>#owner@user:<id>
func (obj *ObjFolder) Owner(subs ...*ObjUser) *ObjFolder {
	for i := range subs {
		sub := subs[i]
		obj.Builder.AddRelationship(v1.Relationship{
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

// CanReadBy optionalrelation.zed:6
// Object: folder:<id>
// Schema: permission read = owner
func (obj *ObjFolder) CanReadBy(subs ...AuthzedObject) *ObjFolder {
	for i := range subs {
		sub := subs[i]
		obj.Builder.CheckPermission(sub, "read", obj)
	}
	return obj
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
