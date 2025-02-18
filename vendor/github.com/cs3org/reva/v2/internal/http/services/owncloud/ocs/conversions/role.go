// Copyright 2018-2021 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

// Package conversions sits between CS3 type definitions and OCS API Responses
package conversions

import (
	"fmt"
	"strings"

	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	"github.com/cs3org/reva/v2/pkg/storage/utils/grants"
)

// Role is a set of ocs permissions and cs3 resource permissions under a common name.
type Role struct {
	Name                   string
	cS3ResourcePermissions *provider.ResourcePermissions
	ocsPermissions         Permissions
}

const (
	// RoleViewer grants non-editor role on a resource.
	RoleViewer = "viewer"
	// RoleSpaceViewer grants non-editor role on a space.
	RoleSpaceViewer = "spaceviewer"
	// RoleEditor grants editor permission on a resource, including folders.
	RoleEditor = "editor"
	// RoleSpaceEditor grants editor permission on a space.
	RoleSpaceEditor = "spaceeditor"
	// RoleFileEditor grants editor permission on a single file.
	RoleFileEditor = "file-editor"
	// RoleUploader grants uploader permission to upload onto a resource.
	RoleUploader = "uploader"
	// RoleManager grants manager permissions on a resource. Semantically equivalent to co-owner.
	RoleManager = "manager"

	// RoleUnknown is used for unknown roles.
	RoleUnknown = "unknown"
	// RoleLegacy provides backwards compatibility.
	RoleLegacy = "legacy"
	// RoleDenied grants no permission at all on a resource
	RoleDenied = "denied"
)

// CS3ResourcePermissions for the role
func (r *Role) CS3ResourcePermissions() *provider.ResourcePermissions {
	return r.cS3ResourcePermissions
}

// OCSPermissions for the role
func (r *Role) OCSPermissions() Permissions {
	return r.ocsPermissions
}

// WebDAVPermissions returns the webdav permissions used in propfinds, eg. "WCKDNVR"
/*
	from https://github.com/owncloud/core/blob/10715e2b1c85fc3855a38d2b1fe4426b5e3efbad/apps/dav/lib/Files/PublicFiles/SharedNodeTrait.php#L196-L215

		$p = '';
		if ($node->isDeletable() && $this->checkSharePermissions(Constants::PERMISSION_DELETE)) {
			$p .= 'D';
		}
		if ($node->isUpdateable() && $this->checkSharePermissions(Constants::PERMISSION_UPDATE)) {
			$p .= 'NV'; // Renameable, Moveable
		}
		if ($node->getType() === \OCP\Files\FileInfo::TYPE_FILE) {
			if ($node->isUpdateable() && $this->checkSharePermissions(Constants::PERMISSION_UPDATE)) {
				$p .= 'W';
			}
		} else {
			if ($node->isCreatable() && $this->checkSharePermissions(Constants::PERMISSION_CREATE)) {
				$p .= 'CK';
			}
		}

*/
// D = delete
// NV = update (renameable moveable)
// W = update (files only)
// CK = create (folders only)
// S = Shared
// R = Shareable
// M = Mounted
// Z = Deniable (NEW)
func (r *Role) WebDAVPermissions(isDir, isShared, isMountpoint, isPublic bool) string {
	var b strings.Builder
	if !isPublic && isShared {
		fmt.Fprintf(&b, "S")
	}
	if r.ocsPermissions.Contain(PermissionShare) {
		fmt.Fprintf(&b, "R")
	}
	if !isPublic && isMountpoint {
		fmt.Fprintf(&b, "M")
	}
	if r.ocsPermissions.Contain(PermissionDelete) {
		fmt.Fprintf(&b, "D") // TODO oc10 shows received shares as deletable
	}
	if r.ocsPermissions.Contain(PermissionWrite) {
		// Single file public link shares cannot be renamed
		if !isPublic || (isPublic && r.cS3ResourcePermissions != nil && r.cS3ResourcePermissions.Move) {
			fmt.Fprintf(&b, "NV")
		}
		if !isDir {
			fmt.Fprintf(&b, "W")
		}
	}
	if isDir && r.ocsPermissions.Contain(PermissionCreate) {
		fmt.Fprintf(&b, "CK")
	}

	if r.CS3ResourcePermissions().DenyGrant {
		fmt.Fprintf(&b, "Z")
	}

	return b.String()
}

// RoleFromName creates a role from the name
func RoleFromName(name string, sharing bool) *Role {
	switch name {
	case RoleDenied:
		return NewDeniedRole()
	case RoleViewer:
		return NewViewerRole(sharing)
	case RoleSpaceViewer:
		return NewSpaceViewerRole()
	case RoleEditor:
		return NewEditorRole(sharing)
	case RoleSpaceEditor:
		return NewSpaceEditorRole()
	case RoleFileEditor:
		return NewFileEditorRole()
	case RoleUploader:
		return NewUploaderRole()
	case RoleManager:
		return NewManagerRole()
	default:
		return NewUnknownRole()
	}
}

// NewUnknownRole creates an unknown role. An Unknown role has no permissions over a cs3 resource nor any ocs endpoint.
func NewUnknownRole() *Role {
	return &Role{
		Name:                   RoleUnknown,
		cS3ResourcePermissions: &provider.ResourcePermissions{},
		ocsPermissions:         PermissionInvalid,
	}
}

// NewDeniedRole creates a fully denied role
func NewDeniedRole() *Role {
	return &Role{
		Name:                   RoleDenied,
		cS3ResourcePermissions: &provider.ResourcePermissions{},
		ocsPermissions:         PermissionsNone,
	}
}

// NewViewerRole creates a viewer role. `sharing` indicates if sharing permission should be added
func NewViewerRole(sharing bool) *Role {
	p := PermissionRead
	if sharing {
		p |= PermissionShare
	}
	return &Role{
		Name: RoleViewer,
		cS3ResourcePermissions: &provider.ResourcePermissions{
			AddGrant:             sharing,
			GetPath:              true,
			GetQuota:             true,
			InitiateFileDownload: true,
			ListContainer:        true,
			ListRecycle:          true,
			Stat:                 true,
		},
		ocsPermissions: p,
	}
}

// NewSpaceViewerRole creates a spaceviewer role
func NewSpaceViewerRole() *Role {
	return &Role{
		Name: RoleSpaceViewer,
		cS3ResourcePermissions: &provider.ResourcePermissions{
			GetPath:              true,
			GetQuota:             true,
			InitiateFileDownload: true,
			ListContainer:        true,
			ListGrants:           true,
			ListRecycle:          true,
			Stat:                 true,
		},
		ocsPermissions: PermissionRead,
	}
}

// NewEditorRole creates an editor role. `sharing` indicates if sharing permission should be added
func NewEditorRole(sharing bool) *Role {
	p := PermissionRead | PermissionCreate | PermissionWrite | PermissionDelete
	if sharing {
		p |= PermissionShare
	}
	return &Role{
		Name: RoleEditor,
		cS3ResourcePermissions: &provider.ResourcePermissions{
			AddGrant:             sharing,
			CreateContainer:      true,
			Delete:               true,
			GetPath:              true,
			GetQuota:             true,
			InitiateFileDownload: true,
			InitiateFileUpload:   true,
			ListContainer:        true,
			ListRecycle:          true,
			Move:                 true,
			RestoreRecycleItem:   true,
			Stat:                 true,
		},
		ocsPermissions: p,
	}
}

// NewSpaceEditorRole creates an editor role
func NewSpaceEditorRole() *Role {
	return &Role{
		Name: RoleSpaceEditor,
		cS3ResourcePermissions: &provider.ResourcePermissions{
			CreateContainer:      true,
			Delete:               true,
			GetPath:              true,
			GetQuota:             true,
			InitiateFileDownload: true,
			InitiateFileUpload:   true,
			ListContainer:        true,
			ListFileVersions:     true,
			ListGrants:           true,
			ListRecycle:          true,
			Move:                 true,
			RestoreFileVersion:   true,
			RestoreRecycleItem:   true,
			Stat:                 true,
		},
		ocsPermissions: PermissionRead | PermissionCreate | PermissionWrite | PermissionDelete,
	}
}

// NewFileEditorRole creates a file-editor role
func NewFileEditorRole() *Role {
	return &Role{
		Name: RoleEditor,
		cS3ResourcePermissions: &provider.ResourcePermissions{
			GetPath:              true,
			GetQuota:             true,
			InitiateFileDownload: true,
			ListGrants:           true,
			ListContainer:        true,
			ListRecycle:          true,
			Stat:                 true,
			InitiateFileUpload:   true,
			RestoreRecycleItem:   true,
		},
		ocsPermissions: PermissionRead | PermissionWrite,
	}
}

// NewUploaderRole creates an uploader role
func NewUploaderRole() *Role {
	return &Role{
		Name: RoleUploader,
		cS3ResourcePermissions: &provider.ResourcePermissions{
			Stat:               true,
			GetPath:            true,
			CreateContainer:    true,
			InitiateFileUpload: true,
		},
		ocsPermissions: PermissionCreate,
	}
}

// NewNoneRole creates a role with no permissions
func NewNoneRole() *Role {
	return &Role{
		Name:                   "none",
		cS3ResourcePermissions: &provider.ResourcePermissions{},
		ocsPermissions:         PermissionInvalid,
	}
}

// NewManagerRole creates an manager role
func NewManagerRole() *Role {
	return &Role{
		Name: RoleManager,
		cS3ResourcePermissions: &provider.ResourcePermissions{
			GetPath:              true,
			GetQuota:             true,
			InitiateFileDownload: true,
			ListGrants:           true,
			ListContainer:        true,
			ListFileVersions:     true,
			ListRecycle:          true,
			Stat:                 true,
			InitiateFileUpload:   true,
			RestoreFileVersion:   true,
			RestoreRecycleItem:   true,
			Move:                 true,
			CreateContainer:      true,
			Delete:               true,
			PurgeRecycle:         true,

			// these permissions only make sense to enforce them in the root of the storage space.
			AddGrant:    true, // managers can add users to the space
			RemoveGrant: true, // managers can remove users from the space
			UpdateGrant: true,
			DenyGrant:   true, // managers can deny access to sub folders
		},
		ocsPermissions: PermissionAll,
	}
}

// RoleFromOCSPermissions tries to map ocs permissions to a role
// TODO: rethink using this. ocs permissions cannot be assigned 1:1 to roles
// NOTE: If resharing=false in the system this function will return SpaceViewerRole instead ViewerRole
func RoleFromOCSPermissions(p Permissions) *Role {
	if p == PermissionInvalid {
		return NewNoneRole()
	}

	if p.Contain(PermissionRead) {
		if p.Contain(PermissionWrite) && p.Contain(PermissionCreate) && p.Contain(PermissionDelete) {
			if p.Contain(PermissionShare) {
				return NewEditorRole(true)
			}

			return NewSpaceEditorRole()
		}
		if p == PermissionRead|PermissionShare {
			return NewViewerRole(true)
		}
		if p == PermissionRead {
			return NewSpaceViewerRole()
		}
	}
	if p == PermissionCreate {
		return NewUploaderRole()
	}
	// legacy
	return NewLegacyRoleFromOCSPermissions(p)
}

// NewLegacyRoleFromOCSPermissions tries to map a legacy combination of ocs permissions to cs3 resource permissions as a legacy role
func NewLegacyRoleFromOCSPermissions(p Permissions) *Role {
	r := &Role{
		Name:                   RoleLegacy, // TODO custom role?
		ocsPermissions:         p,
		cS3ResourcePermissions: &provider.ResourcePermissions{},
	}
	if p.Contain(PermissionRead) {
		r.cS3ResourcePermissions.ListContainer = true
		// r.cS3ResourcePermissions.ListGrants = true
		r.cS3ResourcePermissions.ListRecycle = true
		r.cS3ResourcePermissions.Stat = true
		r.cS3ResourcePermissions.GetPath = true
		r.cS3ResourcePermissions.GetQuota = true
		r.cS3ResourcePermissions.InitiateFileDownload = true
	}
	if p.Contain(PermissionWrite) {
		r.cS3ResourcePermissions.InitiateFileUpload = true
		r.cS3ResourcePermissions.RestoreRecycleItem = true
	}
	if p.Contain(PermissionCreate) {
		r.cS3ResourcePermissions.Stat = true
		r.cS3ResourcePermissions.CreateContainer = true
		// FIXME permissions mismatch: double check ocs create vs update file
		// - if the file exists the ocs api needs to check update permission,
		// - if the file does not exist  the ocs api needs to check update permission
		r.cS3ResourcePermissions.InitiateFileUpload = true
		if p.Contain(PermissionWrite) {
			r.cS3ResourcePermissions.Move = true // TODO move only when create and write?
		}
	}
	if p.Contain(PermissionDelete) {
		r.cS3ResourcePermissions.Delete = true
	}
	if p.Contain(PermissionShare) {
		r.cS3ResourcePermissions.AddGrant = true
		// r.cS3ResourcePermissions.RemoveGrant = true // TODO when are you able to unshare / delete
		// r.cS3ResourcePermissions.UpdateGrant = true
	}
	return r
}

// RoleFromResourcePermissions tries to map cs3 resource permissions to a role
// It needs to know whether this is a link or not, because empty permissions on links mean "INTERNAL LINK"
// while empty permissions on other resources mean "DENIAL". Obviously this is not optimal.
func RoleFromResourcePermissions(rp *provider.ResourcePermissions, islink bool) *Role {
	r := &Role{
		Name:                   RoleUnknown,
		ocsPermissions:         PermissionInvalid,
		cS3ResourcePermissions: rp,
	}
	if rp == nil {
		return r
	}
	if grants.PermissionsEqual(rp, &provider.ResourcePermissions{}) {
		if !islink {
			r.ocsPermissions = PermissionsNone
			r.Name = RoleDenied
		}
		return r
	}
	if rp.ListContainer &&
		rp.ListRecycle &&
		rp.Stat &&
		rp.GetPath &&
		rp.GetQuota &&
		rp.InitiateFileDownload {
		r.ocsPermissions |= PermissionRead
	}
	if rp.InitiateFileUpload &&
		rp.RestoreRecycleItem {
		r.ocsPermissions |= PermissionWrite
	}
	if rp.Stat &&
		rp.CreateContainer &&
		rp.InitiateFileUpload {
		r.ocsPermissions |= PermissionCreate
	}
	if rp.Delete {
		r.ocsPermissions |= PermissionDelete
	}
	if rp.AddGrant {
		r.ocsPermissions |= PermissionShare
	}

	if r.ocsPermissions.Contain(PermissionRead) {
		if r.ocsPermissions.Contain(PermissionWrite) && r.ocsPermissions.Contain(PermissionCreate) && r.ocsPermissions.Contain(PermissionDelete) && r.ocsPermissions.Contain(PermissionShare) {
			r.Name = RoleEditor

			if rp.RemoveGrant {
				r.Name = RoleManager
			}
			return r // editor or manager
		}
		if r.ocsPermissions == PermissionRead|PermissionShare {
			r.Name = RoleViewer
			return r
		}
	}
	if r.ocsPermissions == PermissionCreate {
		r.Name = RoleUploader
		return r
	}
	r.Name = RoleLegacy
	// at this point other ocs permissions may have been mapped.
	// TODO what about even more granular cs3 permissions?, eg. only stat
	return r
}
