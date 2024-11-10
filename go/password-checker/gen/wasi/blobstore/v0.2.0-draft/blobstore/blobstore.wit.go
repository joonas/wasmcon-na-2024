// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package blobstore represents the imported interface "wasi:blobstore/blobstore@0.2.0-draft".
//
// wasi-cloud Blobstore service definition
package blobstore

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/protochron/http-password-checker-go/gen/wasi/blobstore/v0.2.0-draft/container"
	"github.com/protochron/http-password-checker-go/gen/wasi/blobstore/v0.2.0-draft/types"
)

// Container represents the imported type alias "wasi:blobstore/blobstore@0.2.0-draft#container".
//
// See [container.Container] for more information.
type Container = container.Container

// Error represents the type alias "wasi:blobstore/blobstore@0.2.0-draft#error".
//
// See [types.Error] for more information.
type Error = types.Error

// ContainerName represents the type alias "wasi:blobstore/blobstore@0.2.0-draft#container-name".
//
// See [types.ContainerName] for more information.
type ContainerName = types.ContainerName

// ObjectID represents the type alias "wasi:blobstore/blobstore@0.2.0-draft#object-id".
//
// See [types.ObjectID] for more information.
type ObjectID = types.ObjectID

// CreateContainer represents the imported function "create-container".
//
// creates a new empty container
//
//	create-container: func(name: container-name) -> result<container, error>
//
//go:nosplit
func CreateContainer(name ContainerName) (result cm.Result[string, Container, Error]) {
	name0, name1 := cm.LowerString(name)
	wasmimport_CreateContainer((*uint8)(name0), (uint32)(name1), &result)
	return
}

// GetContainer represents the imported function "get-container".
//
// retrieves a container by name
//
//	get-container: func(name: container-name) -> result<container, error>
//
//go:nosplit
func GetContainer(name ContainerName) (result cm.Result[string, Container, Error]) {
	name0, name1 := cm.LowerString(name)
	wasmimport_GetContainer((*uint8)(name0), (uint32)(name1), &result)
	return
}

// DeleteContainer represents the imported function "delete-container".
//
// deletes a container and all objects within it
//
//	delete-container: func(name: container-name) -> result<_, error>
//
//go:nosplit
func DeleteContainer(name ContainerName) (result cm.Result[Error, struct{}, Error]) {
	name0, name1 := cm.LowerString(name)
	wasmimport_DeleteContainer((*uint8)(name0), (uint32)(name1), &result)
	return
}

// ContainerExists represents the imported function "container-exists".
//
// returns true if the container exists
//
//	container-exists: func(name: container-name) -> result<bool, error>
//
//go:nosplit
func ContainerExists(name ContainerName) (result cm.Result[string, bool, Error]) {
	name0, name1 := cm.LowerString(name)
	wasmimport_ContainerExists((*uint8)(name0), (uint32)(name1), &result)
	return
}

// CopyObject represents the imported function "copy-object".
//
// copies (duplicates) an object, to the same or a different container.
// returns an error if the target container does not exist.
// overwrites destination object if it already existed.
//
//	copy-object: func(src: object-id, dest: object-id) -> result<_, error>
//
//go:nosplit
func CopyObject(src ObjectID, dest ObjectID) (result cm.Result[Error, struct{}, Error]) {
	src0, src1, src2, src3 := lower_ObjectID(src)
	dest0, dest1, dest2, dest3 := lower_ObjectID(dest)
	wasmimport_CopyObject((*uint8)(src0), (uint32)(src1), (*uint8)(src2), (uint32)(src3), (*uint8)(dest0), (uint32)(dest1), (*uint8)(dest2), (uint32)(dest3), &result)
	return
}

// MoveObject represents the imported function "move-object".
//
// moves or renames an object, to the same or a different container
// returns an error if the destination container does not exist.
// overwrites destination object if it already existed.
//
//	move-object: func(src: object-id, dest: object-id) -> result<_, error>
//
//go:nosplit
func MoveObject(src ObjectID, dest ObjectID) (result cm.Result[Error, struct{}, Error]) {
	src0, src1, src2, src3 := lower_ObjectID(src)
	dest0, dest1, dest2, dest3 := lower_ObjectID(dest)
	wasmimport_MoveObject((*uint8)(src0), (uint32)(src1), (*uint8)(src2), (uint32)(src3), (*uint8)(dest0), (uint32)(dest1), (*uint8)(dest2), (uint32)(dest3), &result)
	return
}