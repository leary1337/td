// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// ScopeAutosaveSettings represents TL type `scopeAutosaveSettings#5c329f33`.
type ScopeAutosaveSettings struct {
	// True, if photo autosave is enabled
	AutosavePhotos bool
	// True, if video autosave is enabled
	AutosaveVideos bool
	// The maximum size of a video file to be autosaved, in bytes; 512 KB - 4000 MB
	MaxVideoFileSize int64
}

// ScopeAutosaveSettingsTypeID is TL type id of ScopeAutosaveSettings.
const ScopeAutosaveSettingsTypeID = 0x5c329f33

// Ensuring interfaces in compile-time for ScopeAutosaveSettings.
var (
	_ bin.Encoder     = &ScopeAutosaveSettings{}
	_ bin.Decoder     = &ScopeAutosaveSettings{}
	_ bin.BareEncoder = &ScopeAutosaveSettings{}
	_ bin.BareDecoder = &ScopeAutosaveSettings{}
)

func (s *ScopeAutosaveSettings) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.AutosavePhotos == false) {
		return false
	}
	if !(s.AutosaveVideos == false) {
		return false
	}
	if !(s.MaxVideoFileSize == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ScopeAutosaveSettings) String() string {
	if s == nil {
		return "ScopeAutosaveSettings(nil)"
	}
	type Alias ScopeAutosaveSettings
	return fmt.Sprintf("ScopeAutosaveSettings%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ScopeAutosaveSettings) TypeID() uint32 {
	return ScopeAutosaveSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*ScopeAutosaveSettings) TypeName() string {
	return "scopeAutosaveSettings"
}

// TypeInfo returns info about TL type.
func (s *ScopeAutosaveSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "scopeAutosaveSettings",
		ID:   ScopeAutosaveSettingsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AutosavePhotos",
			SchemaName: "autosave_photos",
		},
		{
			Name:       "AutosaveVideos",
			SchemaName: "autosave_videos",
		},
		{
			Name:       "MaxVideoFileSize",
			SchemaName: "max_video_file_size",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *ScopeAutosaveSettings) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode scopeAutosaveSettings#5c329f33 as nil")
	}
	b.PutID(ScopeAutosaveSettingsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *ScopeAutosaveSettings) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode scopeAutosaveSettings#5c329f33 as nil")
	}
	b.PutBool(s.AutosavePhotos)
	b.PutBool(s.AutosaveVideos)
	b.PutInt53(s.MaxVideoFileSize)
	return nil
}

// Decode implements bin.Decoder.
func (s *ScopeAutosaveSettings) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode scopeAutosaveSettings#5c329f33 to nil")
	}
	if err := b.ConsumeID(ScopeAutosaveSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode scopeAutosaveSettings#5c329f33: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *ScopeAutosaveSettings) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode scopeAutosaveSettings#5c329f33 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode scopeAutosaveSettings#5c329f33: field autosave_photos: %w", err)
		}
		s.AutosavePhotos = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode scopeAutosaveSettings#5c329f33: field autosave_videos: %w", err)
		}
		s.AutosaveVideos = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode scopeAutosaveSettings#5c329f33: field max_video_file_size: %w", err)
		}
		s.MaxVideoFileSize = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *ScopeAutosaveSettings) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode scopeAutosaveSettings#5c329f33 as nil")
	}
	b.ObjStart()
	b.PutID("scopeAutosaveSettings")
	b.Comma()
	b.FieldStart("autosave_photos")
	b.PutBool(s.AutosavePhotos)
	b.Comma()
	b.FieldStart("autosave_videos")
	b.PutBool(s.AutosaveVideos)
	b.Comma()
	b.FieldStart("max_video_file_size")
	b.PutInt53(s.MaxVideoFileSize)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *ScopeAutosaveSettings) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode scopeAutosaveSettings#5c329f33 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("scopeAutosaveSettings"); err != nil {
				return fmt.Errorf("unable to decode scopeAutosaveSettings#5c329f33: %w", err)
			}
		case "autosave_photos":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode scopeAutosaveSettings#5c329f33: field autosave_photos: %w", err)
			}
			s.AutosavePhotos = value
		case "autosave_videos":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode scopeAutosaveSettings#5c329f33: field autosave_videos: %w", err)
			}
			s.AutosaveVideos = value
		case "max_video_file_size":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode scopeAutosaveSettings#5c329f33: field max_video_file_size: %w", err)
			}
			s.MaxVideoFileSize = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAutosavePhotos returns value of AutosavePhotos field.
func (s *ScopeAutosaveSettings) GetAutosavePhotos() (value bool) {
	if s == nil {
		return
	}
	return s.AutosavePhotos
}

// GetAutosaveVideos returns value of AutosaveVideos field.
func (s *ScopeAutosaveSettings) GetAutosaveVideos() (value bool) {
	if s == nil {
		return
	}
	return s.AutosaveVideos
}

// GetMaxVideoFileSize returns value of MaxVideoFileSize field.
func (s *ScopeAutosaveSettings) GetMaxVideoFileSize() (value int64) {
	if s == nil {
		return
	}
	return s.MaxVideoFileSize
}
