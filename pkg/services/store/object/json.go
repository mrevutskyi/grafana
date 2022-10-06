package object

import (
	"encoding/json"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() { //nolint:gochecknoinits
	jsoniter.RegisterTypeEncoder("object.ObjectSearchResult", &searchResultCodec{})
	jsoniter.RegisterTypeEncoder("object.WriteObjectResponse", &writeResponseCodec{})
	jsoniter.RegisterTypeEncoder("object.RawObject", &rawObjectCodec{})
	jsoniter.RegisterTypeEncoder("object.ReadObjectResponse", &readResponseCodec{})
}

// Unlike the standard JSON marshal, this will write bytes as JSON when it can
type rawObjectCodec struct{}

func (obj *RawObject) MarshalJSON() ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(obj)
}

func (codec *rawObjectCodec) IsEmpty(ptr unsafe.Pointer) bool {
	f := (*RawObject)(ptr)
	return f.UID == "" && f.Body == nil
}

func (codec *rawObjectCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*RawObject)(ptr)
	stream.WriteObjectStart()
	stream.WriteObjectField("UID")
	stream.WriteString(obj.UID)

	if obj.Kind != "" {
		stream.WriteMore()
		stream.WriteObjectField("kind")
		stream.WriteString(obj.Kind)
	}
	if obj.Version != "" {
		stream.WriteMore()
		stream.WriteObjectField("version")
		stream.WriteString(obj.Version)
	}
	if obj.Updated > 0 {
		stream.WriteMore()
		stream.WriteObjectField("updated")
		stream.WriteInt64(obj.Updated)
	}
	if obj.UpdatedBy != "" {
		stream.WriteMore()
		stream.WriteObjectField("updatedBy")
		stream.WriteString(obj.UpdatedBy)
	}
	if obj.Created > 0 {
		stream.WriteMore()
		stream.WriteObjectField("created")
		stream.WriteInt64(obj.Created)
	}
	if obj.CreatedBy != "" {
		stream.WriteMore()
		stream.WriteObjectField("createdBy")
		stream.WriteString(obj.CreatedBy)
	}
	if obj.Body != nil {
		stream.WriteMore()
		if json.Valid(obj.Body) {
			stream.WriteObjectField("body")
			stream.WriteRaw(string(obj.Body)) // works for strings
		} else {
			stream.WriteObjectField("body_base64")
			stream.WriteVal(obj.Body) // works for strings
		}
	}
	if obj.ETag != "" {
		stream.WriteMore()
		stream.WriteObjectField("etag")
		stream.WriteString(obj.ETag)
	}
	if obj.Size > 0 {
		stream.WriteMore()
		stream.WriteObjectField("size")
		stream.WriteInt64(obj.Size)
	}

	if obj.Sync != nil {
		stream.WriteMore()
		stream.WriteObjectField("sync")
		stream.WriteVal(obj.Sync)
	}

	stream.WriteObjectEnd()
}

// Unlike the standard JSON marshal, this will write bytes as JSON when it can
type readResponseCodec struct{}

func (obj *ReadObjectResponse) MarshalJSON() ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(obj)
}

func (codec *readResponseCodec) IsEmpty(ptr unsafe.Pointer) bool {
	f := (*ReadObjectResponse)(ptr)
	return f == nil
}

func (codec *readResponseCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ReadObjectResponse)(ptr)
	stream.WriteObjectStart()
	stream.WriteObjectField("object")
	stream.WriteVal(obj.Object)

	if len(obj.SummaryJson) > 0 {
		stream.WriteMore()
		stream.WriteObjectField("summary")
		stream.WriteRaw(string(obj.SummaryJson))
	}

	stream.WriteObjectEnd()
}

// Unlike the standard JSON marshal, this will write bytes as JSON when it can
type searchResultCodec struct{}

func (obj *ObjectSearchResult) MarshalJSON() ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(obj)
}

func (codec *searchResultCodec) IsEmpty(ptr unsafe.Pointer) bool {
	f := (*ObjectSearchResult)(ptr)
	return f.UID == "" && f.Body == nil
}

func (codec *searchResultCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ObjectSearchResult)(ptr)
	stream.WriteObjectStart()
	stream.WriteObjectField("UID")
	stream.WriteString(obj.UID)

	if obj.Kind != "" {
		stream.WriteMore()
		stream.WriteObjectField("kind")
		stream.WriteString(obj.Kind)
	}
	if obj.Name != "" {
		stream.WriteMore()
		stream.WriteObjectField("name")
		stream.WriteString(obj.Name)
	}
	if obj.Description != "" {
		stream.WriteMore()
		stream.WriteObjectField("description")
		stream.WriteString(obj.Description)
	}
	if obj.Updated > 0 {
		stream.WriteMore()
		stream.WriteObjectField("updated")
		stream.WriteInt64(obj.Updated)
	}
	if obj.UpdatedBy != "" {
		stream.WriteMore()
		stream.WriteObjectField("updatedBy")
		stream.WriteVal(obj.UpdatedBy)
	}
	if obj.Body != nil {
		stream.WriteMore()
		if json.Valid(obj.Body) {
			stream.WriteObjectField("body")
			stream.WriteRaw(string(obj.Body)) // works for strings
		} else {
			stream.WriteObjectField("body_base64")
			stream.WriteVal(obj.Body) // works for strings
		}
	}
	if obj.Labels != nil {
		stream.WriteMore()
		stream.WriteObjectField("labels")
		stream.WriteVal(obj.Labels)
	}
	if obj.ErrorJson != nil {
		stream.WriteMore()
		stream.WriteObjectField("error")
		stream.WriteRaw(string(obj.ErrorJson))
	}
	if obj.FieldsJson != nil {
		stream.WriteMore()
		stream.WriteObjectField("fields")
		stream.WriteRaw(string(obj.FieldsJson))
	}

	stream.WriteObjectEnd()
}

// Unlike the standard JSON marshal, this will write bytes as JSON when it can
type writeResponseCodec struct{}

func (obj *WriteObjectResponse) MarshalJSON() ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(obj)
}

func (codec *writeResponseCodec) IsEmpty(ptr unsafe.Pointer) bool {
	f := (*WriteObjectResponse)(ptr)
	return f == nil
}

func (codec *writeResponseCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*WriteObjectResponse)(ptr)
	stream.WriteObjectStart()
	stream.WriteObjectField("status")
	stream.WriteString(obj.Status.String())

	if obj.Error != nil {
		stream.WriteMore()
		stream.WriteObjectField("error")
		stream.WriteVal(obj.Error)
	}
	if obj.Object != nil {
		stream.WriteMore()
		stream.WriteObjectField("object")
		stream.WriteVal(obj.Object)
	}
	if len(obj.SummaryJson) > 0 {
		stream.WriteMore()
		stream.WriteObjectField("summary")
		stream.WriteRaw(string(obj.SummaryJson))
	}
	stream.WriteObjectEnd()
}
