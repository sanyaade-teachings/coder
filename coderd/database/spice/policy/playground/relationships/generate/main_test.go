package main_test

import (
	"fmt"
	"testing"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/spicedb/pkg/tuple"
)

func TestString(t *testing.T) {
	rel := v1.Relationship{
		Resource: &v1.ObjectReference{
			ObjectType: "group",
			ObjectId:   "everyone",
		},
		Relation: "member",
		Subject: &v1.SubjectReference{
			Object: &v1.ObjectReference{
				ObjectType: "user",
				ObjectId:   "*",
			},
		},
		OptionalCaveat: nil,
	}

	fmt.Println(tuple.MustRelString(&rel))

}
