package svc

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"

	"github.com/CiscoM31/godata"
	libregraph "github.com/owncloud/libre-graph-api-go"
	"github.com/owncloud/ocis/v2/services/graph/pkg/service/v0/errorcode"

	revactx "github.com/cs3org/reva/v2/pkg/ctx"
	"github.com/cs3org/reva/v2/pkg/events"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const memberTypeUsers = "users"

// GetGroups implements the Service interface.
func (g Graph) GetGroups(w http.ResponseWriter, r *http.Request) {
	logger := g.logger.SubloggerWithRequestID(r.Context())
	logger.Info().Interface("query", r.URL.Query()).Msg("calling get groups")
	sanitizedPath := strings.TrimPrefix(r.URL.Path, "/graph/v1.0/")
	odataReq, err := godata.ParseRequest(r.Context(), sanitizedPath, r.URL.Query())
	if err != nil {
		logger.Debug().Err(err).Interface("query", r.URL.Query()).Msg("could not get groups: query error")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, err.Error())
		return
	}

	groups, err := g.identityBackend.GetGroups(r.Context(), r.URL.Query())
	if err != nil {
		logger.Debug().Err(err).Msg("could not get groups: backend error")
		var errcode errorcode.Error
		if errors.As(err, &errcode) {
			errcode.Render(w, r)
		} else {
			errorcode.GeneralException.Render(w, r, http.StatusInternalServerError, err.Error())
		}
		return
	}

	groups, err = sortGroups(odataReq, groups)
	if err != nil {
		logger.Debug().Err(err).Interface("query", r.URL.Query()).Msg("cannot get groups: could not sort groups according to query")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, err.Error())
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, &ListResponse{Value: groups})
}

// PostGroup implements the Service interface.
func (g Graph) PostGroup(w http.ResponseWriter, r *http.Request) {
	logger := g.logger.SubloggerWithRequestID(r.Context())
	logger.Info().Msg("calling post group")
	grp := libregraph.NewGroup()
	err := StrictJSONUnmarshal(r.Body, grp)
	if err != nil {
		logger.Debug().Err(err).Interface("body", r.Body).Msg("could not create group: invalid request body")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, fmt.Sprintf("invalid request body: %s", err.Error()))
		return
	}

	if _, ok := grp.GetDisplayNameOk(); !ok {
		logger.Debug().Err(err).Interface("group", grp).Msg("could not create group: missing required attribute")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "Missing Required Attribute")
		return
	}

	// Disallow user-supplied IDs. It's supposed to be readonly. We're either
	// generating them in the backend ourselves or rely on the Backend's
	// storage (e.g. LDAP) to provide a unique ID.
	if _, ok := grp.GetIdOk(); ok {
		logger.Debug().Msg("could not create group: id is a read-only attribute")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "group id is a read-only attribute")
		return
	}

	if grp, err = g.identityBackend.CreateGroup(r.Context(), *grp); err != nil {
		var errcode errorcode.Error
		if errors.As(err, &errcode) {
			errcode.Render(w, r)
		} else {
			logger.Debug().Interface("group", grp).Msg("could not create group: backend error")
			errorcode.GeneralException.Render(w, r, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if grp != nil && grp.Id != nil {
		e := events.GroupCreated{
			GroupID: grp.GetId(),
		}
		if currentUser, ok := revactx.ContextGetUser(r.Context()); ok {
			e.Executant = currentUser.GetId()
		}
		g.publishEvent(e)
	}
	render.Status(r, http.StatusOK) // FIXME 201 should return 201 created
	render.JSON(w, r, grp)
}

// PatchGroup implements the Service interface.
func (g Graph) PatchGroup(w http.ResponseWriter, r *http.Request) {
	logger := g.logger.SubloggerWithRequestID(r.Context())
	logger.Info().Msg("calling patch group")
	groupID := chi.URLParam(r, "groupID")
	groupID, err := url.PathUnescape(groupID)
	if err != nil {
		logger.Debug().Str("id", groupID).Msg("could not change group: unescaping group id failed")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "unescaping group id failed")
		return
	}

	if groupID == "" {
		logger.Debug().Msg("could not change group: missing group id")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "missing group id")
		return
	}
	changes := libregraph.NewGroup()
	err = StrictJSONUnmarshal(r.Body, changes)
	if err != nil {
		logger.Debug().Err(err).Interface("body", r.Body).Msg("could not change group: invalid request body")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, fmt.Sprintf("invalid request body: %s", err.Error()))
		return
	}

	if reflect.ValueOf(*changes).IsZero() {
		logger.Debug().Interface("body", r.Body).Msg("ignoring empyt request body")
		render.Status(r, http.StatusNoContent)
		render.NoContent(w, r)
		return
	}

	if changes.HasDisplayName() {
		groupName := changes.GetDisplayName()
		err = g.identityBackend.UpdateGroupName(r.Context(), groupID, groupName)
	}

	if memberRefs, ok := changes.GetMembersodataBindOk(); ok {
		// The spec defines a limit of 20 members maxium per Request
		if len(memberRefs) > g.config.API.GroupMembersPatchLimit {
			logger.Debug().
				Int("number", len(memberRefs)).
				Int("limit", g.config.API.GroupMembersPatchLimit).
				Msg("could not add group members, exceeded members limit")
			errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest,
				fmt.Sprintf("Request is limited to %d members", g.config.API.GroupMembersPatchLimit))
			return
		}
		memberIDs := make([]string, 0, len(memberRefs))
		for _, memberRef := range memberRefs {
			memberType, id, err := g.parseMemberRef(memberRef)
			if err != nil {
				logger.Debug().
					Str("memberref", memberRef).
					Msg("could not change group: Error parsing member@odata.bind values")
				errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "Error parsing member@odata.bind values")
				return
			}
			logger.Debug().Str("membertype", memberType).Str("memberid", id).Msg("add group member")
			// The MS Graph spec allows "directoryObject", "user", "group" and "organizational Contact"
			// we restrict this to users for now. Might add Groups as members later
			if memberType != memberTypeUsers {
				logger.Debug().
					Str("type", memberType).
					Msg("could not change group: could not add member, only user type is allowed")
				errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "Only user are allowed as group members")
				return
			}
			memberIDs = append(memberIDs, id)
		}
		err = g.identityBackend.AddMembersToGroup(r.Context(), groupID, memberIDs)
	}

	if err != nil {
		logger.Debug().Err(err).Msg("could not change group: backend could not add members")
		var errcode errorcode.Error
		if errors.As(err, &errcode) {
			errcode.Render(w, r)
		} else {
			errorcode.GeneralException.Render(w, r, http.StatusInternalServerError, err.Error())
		}
		return
	}
	render.Status(r, http.StatusNoContent) // TODO StatusNoContent when prefer=minimal is used, otherwise OK and the resource in the body
	render.NoContent(w, r)
}

// GetGroup implements the Service interface.
func (g Graph) GetGroup(w http.ResponseWriter, r *http.Request) {
	logger := g.logger.SubloggerWithRequestID(r.Context())
	logger.Info().Msg("calling get group")
	groupID := chi.URLParam(r, "groupID")
	groupID, err := url.PathUnescape(groupID)
	if err != nil {
		logger.Debug().Str("id", groupID).Msg("could not get group: unescaping group id failed")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "unescaping group id failed")
	}

	if groupID == "" {
		logger.Debug().Msg("could not get group: missing group id")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "missing group id")
		return
	}

	logger.Debug().
		Str("id", groupID).
		Interface("query", r.URL.Query()).
		Msg("calling get group on backend")
	group, err := g.identityBackend.GetGroup(r.Context(), groupID, r.URL.Query())
	if err != nil {
		logger.Debug().Err(err).Msg("could not get group: backend error")
		var errcode errorcode.Error
		if errors.As(err, &errcode) {
			errcode.Render(w, r)
		} else {
			errorcode.GeneralException.Render(w, r, http.StatusInternalServerError, err.Error())
		}
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, group)
}

// DeleteGroup implements the Service interface.
func (g Graph) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	logger := g.logger.SubloggerWithRequestID(r.Context())
	logger.Info().Msg("calling delete group")
	groupID := chi.URLParam(r, "groupID")
	groupID, err := url.PathUnescape(groupID)
	if err != nil {
		logger.Debug().Err(err).Str("id", groupID).Msg("could not delete group: unescaping group id failed")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "unescaping group id failed")
		return
	}

	if groupID == "" {
		logger.Debug().Msg("could not delete group: missing group id")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "missing group id")
		return
	}

	logger.Debug().Str("id", groupID).Msg("calling delete group on backend")
	err = g.identityBackend.DeleteGroup(r.Context(), groupID)

	if err != nil {
		logger.Debug().Err(err).Msg("could not delete group: backend error")
		var errcode errorcode.Error
		if errors.As(err, &errcode) {
			errcode.Render(w, r)
		} else {
			errorcode.GeneralException.Render(w, r, http.StatusInternalServerError, err.Error())
		}
		return
	}

	e := events.GroupDeleted{
		GroupID: groupID,
	}
	if currentUser, ok := revactx.ContextGetUser(r.Context()); ok {
		e.Executant = currentUser.GetId()
	}
	g.publishEvent(e)
	render.Status(r, http.StatusNoContent)
	render.NoContent(w, r)
}

// GetGroupMembers implements the Service interface.
func (g Graph) GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	logger := g.logger.SubloggerWithRequestID(r.Context())
	logger.Info().Msg("calling get group members")
	sanitizedPath := strings.TrimPrefix(r.URL.Path, "/graph/v1.0/")
	groupID := chi.URLParam(r, "groupID")
	groupID, err := url.PathUnescape(groupID)
	if err != nil {
		logger.Debug().Str("id", groupID).Msg("could not get group members: unescaping group id failed")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "unescaping group id failed")
		return
	}

	if groupID == "" {
		logger.Debug().Msg("could not get group members: missing group id")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "missing group id")
		return
	}

	odataReq, err := godata.ParseRequest(r.Context(), sanitizedPath, r.URL.Query())
	if err != nil {
		logger.Debug().Err(err).Interface("query", r.URL.Query()).Msg("could not get users: query error")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, err.Error())
		return
	}

	logger.Debug().Str("id", groupID).Msg("calling get group members on backend")
	members, err := g.identityBackend.GetGroupMembers(r.Context(), groupID, odataReq)
	if err != nil {
		logger.Debug().Err(err).Msg("could not get group members: backend error")
		var errcode errorcode.Error
		if errors.As(err, &errcode) {
			errcode.Render(w, r)
		} else {
			errorcode.GeneralException.Render(w, r, http.StatusInternalServerError, err.Error())
		}
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, members)
}

// PostGroupMember implements the Service interface.
func (g Graph) PostGroupMember(w http.ResponseWriter, r *http.Request) {
	logger := g.logger.SubloggerWithRequestID(r.Context())
	logger.Info().Msg("Calling post group member")

	groupID := chi.URLParam(r, "groupID")
	groupID, err := url.PathUnescape(groupID)
	if err != nil {
		logger.Debug().
			Err(err).
			Str("id", groupID).
			Msg("could not add member to group: unescaping group id failed")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "unescaping group id failed")
		return
	}

	if groupID == "" {
		logger.Debug().Msg("could not add group member: missing group id")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "missing group id")
		return
	}
	memberRef := libregraph.NewMemberReference()
	err = StrictJSONUnmarshal(r.Body, memberRef)
	if err != nil {
		logger.Debug().
			Err(err).
			Interface("body", r.Body).
			Msg("could not add group member: invalid request body")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, fmt.Sprintf("invalid request body: %s", err.Error()))
		return
	}
	memberRefURL, ok := memberRef.GetOdataIdOk()
	if !ok {
		logger.Debug().Msg("could not add group member: @odata.id reference is missing")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "@odata.id reference is missing")
		return
	}
	memberType, id, err := g.parseMemberRef(*memberRefURL)
	if err != nil {
		logger.Debug().Err(err).Msg("could not add group member: error parsing @odata.id url")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "Error parsing @odata.id url")
		return
	}
	// The MS Graph spec allows "directoryObject", "user", "group" and "organizational Contact"
	// we restrict this to users for now. Might add Groups as members later
	if memberType != memberTypeUsers {
		logger.Debug().Str("type", memberType).Msg("could not add group member: Only users are allowed as group members")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "Only users are allowed as group members")
		return
	}

	logger.Debug().Str("memberType", memberType).Str("id", id).Msg("calling add member on backend")
	err = g.identityBackend.AddMembersToGroup(r.Context(), groupID, []string{id})

	if err != nil {
		logger.Debug().Err(err).Msg("could not add group member: backend error")
		var errcode errorcode.Error
		if errors.As(err, &errcode) {
			errcode.Render(w, r)
		} else {
			errorcode.GeneralException.Render(w, r, http.StatusInternalServerError, err.Error())
		}
		return
	}

	e := events.GroupMemberAdded{
		GroupID: groupID,
		UserID:  id,
	}
	if currentUser, ok := revactx.ContextGetUser(r.Context()); ok {
		e.Executant = currentUser.GetId()
	}
	g.publishEvent(e)
	render.Status(r, http.StatusNoContent)
	render.NoContent(w, r)
}

// DeleteGroupMember implements the Service interface.
func (g Graph) DeleteGroupMember(w http.ResponseWriter, r *http.Request) {
	logger := g.logger.SubloggerWithRequestID(r.Context())
	logger.Info().Msg("calling delete group member")

	groupID := chi.URLParam(r, "groupID")
	groupID, err := url.PathUnescape(groupID)
	if err != nil {
		logger.Debug().Err(err).Str("id", groupID).Msg("could not delete group member: unescaping group id failed")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "unescaping group id failed")
		return
	}

	if groupID == "" {
		logger.Debug().Msg("could not delete group member: missing group id")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "missing group id")
		return
	}

	memberID := chi.URLParam(r, "memberID")
	memberID, err = url.PathUnescape(memberID)
	if err != nil {
		logger.Debug().Err(err).Str("id", memberID).Msg("could not delete group member: unescaping member id failed")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "unescaping member id failed")
		return
	}

	if memberID == "" {
		logger.Debug().Msg("could not delete group member: missing member id")
		errorcode.InvalidRequest.Render(w, r, http.StatusBadRequest, "missing member id")
		return
	}
	logger.Debug().Str("groupID", groupID).Str("memberID", memberID).Msg("calling delete member on backend")
	err = g.identityBackend.RemoveMemberFromGroup(r.Context(), groupID, memberID)

	if err != nil {
		logger.Debug().Err(err).Msg("could not delete group member: backend error")
		var errcode errorcode.Error
		if errors.As(err, &errcode) {
			errcode.Render(w, r)
		} else {
			errorcode.GeneralException.Render(w, r, http.StatusInternalServerError, err.Error())
		}
		return
	}
	e := events.GroupMemberRemoved{
		GroupID: groupID,
		UserID:  memberID,
	}
	if currentUser, ok := revactx.ContextGetUser(r.Context()); ok {
		e.Executant = currentUser.GetId()
	}
	g.publishEvent(e)
	render.Status(r, http.StatusNoContent)
	render.NoContent(w, r)
}

func sortGroups(req *godata.GoDataRequest, groups []*libregraph.Group) ([]*libregraph.Group, error) {
	if req.Query.OrderBy == nil || len(req.Query.OrderBy.OrderByItems) != 1 {
		return groups, nil
	}
	var less func(i, j int) bool

	switch req.Query.OrderBy.OrderByItems[0].Field.Value {
	case displayNameAttr:
		less = func(i, j int) bool {
			return strings.ToLower(groups[i].GetDisplayName()) < strings.ToLower(groups[j].GetDisplayName())
		}
	default:
		return nil, fmt.Errorf("we do not support <%s> as a order parameter", req.Query.OrderBy.OrderByItems[0].Field.Value)
	}

	if req.Query.OrderBy.OrderByItems[0].Order == _sortDescending {
		sort.Slice(groups, reverse(less))
	} else {
		sort.Slice(groups, less)
	}

	return groups, nil
}
