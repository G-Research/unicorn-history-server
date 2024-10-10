package webservice

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/G-Research/yunikorn-history-server/internal/database/repository"
)

const (
	queryParamSubmissionStartTime          = "submissionStartTime"
	queryParamSubmissionEndTime            = "submissionEndTime"
	queryParamGroups                       = "groups"
	queryParamLimit                        = "limit"
	queryParamOffset                       = "offset"
	queryParamUser                         = "user"
	queryParamTimestampStart               = "timestampStart"
	queryParamTimestampEnd                 = "timestampEnd"
	queryParamClusterID                    = "clusterId"
	queryParamState                        = "state"
	queryParamName                         = "name"
	queryParamLastStateTransitionTimeStart = "lastStateTransitionTimeStart"
	queryParamLastStateTransitionTimeEnd   = "lastStateTransitionTimeEnd"
)

func parsePartitionFilters(r *http.Request) (*repository.PartitionFilters, error) {
	filters := repository.PartitionFilters{}
	lastStateTransitionTimeStart, err := getLastStateTransitionTimeStartQueryParam(r)
	if err != nil {
		return nil, err
	}
	if lastStateTransitionTimeStart != nil {
		filters.LastStateTransitionTimeStart = lastStateTransitionTimeStart
	}
	lastStateTransitionTimeEnd, err := getLastStateTransitionTimeEndQueryParam(r)
	if err != nil {
		return nil, err
	}
	if lastStateTransitionTimeEnd != nil {
		filters.LastStateTransitionTimeEnd = lastStateTransitionTimeEnd
	}
	name := getNameQueryParam(r)
	if name != "" {
		filters.Name = &name

	}
	clusterID := getClusterIDQueryParam(r)
	if clusterID != "" {
		filters.ClusterID = &clusterID
	}
	state := getStateQueryParam(r)
	if state != "" {
		filters.State = &state
	}
	offset, err := getOffsetQueryParam(r)
	if err != nil {
		return nil, err
	}
	if offset != nil {
		filters.Offset = offset
	}
	limit, err := getLimitQueryParam(r)
	if err != nil {
		return nil, err
	}
	if limit != nil {
		filters.Limit = limit
	}
	return &filters, nil
}

func parseApplicationFilters(r *http.Request) (*repository.ApplicationFilters, error) {
	filters := repository.ApplicationFilters{}
	user := getUserQueryParam(r)
	if user != "" {
		filters.User = &user
	}
	submissionStartTime, err := getSubmissionStartTimeQueryParam(r)
	if err != nil {
		return nil, err
	}
	if submissionStartTime != nil {
		filters.SubmissionStartTime = submissionStartTime
	}
	submissionEndTime, err := getSubmissionEndTimeQueryParam(r)
	if err != nil {
		return nil, err
	}
	if submissionEndTime != nil {
		filters.SubmissionEndTime = submissionEndTime
	}
	offset, err := getOffsetQueryParam(r)
	if err != nil {
		return nil, err
	}
	if offset != nil {
		filters.Offset = offset
	}
	limit, err := getLimitQueryParam(r)
	if err != nil {
		return nil, err
	}
	if limit != nil {
		filters.Limit = limit
	}
	groups := getGroupsQueryParam(r)
	if len(groups) > 0 {
		filters.Groups = groups
	}
	return &filters, nil
}

func parseHistoryFilters(r *http.Request) (*repository.HistoryFilters, error) {
	var filters repository.HistoryFilters
	timestampStart, err := getTimestampStartQueryParam(r)
	if err != nil {
		return nil, err
	}
	if timestampStart != nil {
		filters.TimestampStart = timestampStart
	}
	timestampEnd, err := getTimestampEndQueryParam(r)
	if err != nil {
		return nil, err
	}
	if timestampEnd != nil {
		filters.TimestampEnd = timestampEnd
	}
	offset, err := getOffsetQueryParam(r)
	if err != nil {
		return nil, err
	}
	if offset != nil {
		filters.Offset = offset
	}
	limit, err := getLimitQueryParam(r)
	if err != nil {
		return nil, err
	}
	if limit != nil {
		filters.Limit = limit
	}
	return &filters, nil
}

func getClusterIDQueryParam(r *http.Request) string {
	return r.URL.Query().Get(queryParamClusterID)
}
func getStateQueryParam(r *http.Request) string {
	return r.URL.Query().Get(queryParamState)
}
func getNameQueryParam(r *http.Request) string {
	return r.URL.Query().Get(queryParamName)
}

func getUserQueryParam(r *http.Request) string {
	return r.URL.Query().Get(queryParamUser)
}

func getGroupsQueryParam(r *http.Request) []string {
	var groupsSlice []string
	groups := r.URL.Query().Get(queryParamGroups)
	if groups != "" {
		groupsSlice = strings.Split(groups, ",")
	}
	return groupsSlice
}

func getOffsetQueryParam(r *http.Request) (*int, error) {
	offsetStr := r.URL.Query().Get(queryParamOffset)
	if offsetStr == "" {
		return nil, nil
	}

	return toInt(offsetStr)
}

func getLimitQueryParam(r *http.Request) (*int, error) {
	limitStr := r.URL.Query().Get(queryParamLimit)
	if limitStr == "" {
		return nil, nil
	}

	return toInt(limitStr)
}

func toInt(numberString string) (*int, error) {
	offset, err := strconv.Atoi(numberString)
	if err != nil {
		return nil, fmt.Errorf("invalid 'offset' query parameter: %v", err)
	}

	return &offset, nil
}

func getLastStateTransitionTimeStartQueryParam(r *http.Request) (*time.Time, error) {
	startStr := r.URL.Query().Get(queryParamLastStateTransitionTimeStart)
	if startStr == "" {
		return nil, nil
	}

	return toTime(startStr)
}

func getLastStateTransitionTimeEndQueryParam(r *http.Request) (*time.Time, error) {
	endStr := r.URL.Query().Get(queryParamLastStateTransitionTimeEnd)
	if endStr == "" {
		return nil, nil
	}

	return toTime(endStr)
}

func getSubmissionStartTimeQueryParam(r *http.Request) (*time.Time, error) {
	startStr := r.URL.Query().Get(queryParamSubmissionStartTime)
	if startStr == "" {
		return nil, nil
	}

	return toTime(startStr)
}

func getSubmissionEndTimeQueryParam(r *http.Request) (*time.Time, error) {
	endStr := r.URL.Query().Get(queryParamSubmissionEndTime)
	if endStr == "" {
		return nil, nil
	}

	return toTime(endStr)
}

func getTimestampStartQueryParam(r *http.Request) (*time.Time, error) {
	startStr := r.URL.Query().Get(queryParamTimestampStart)
	if startStr == "" {
		return nil, nil
	}

	return toTime(startStr)
}

func getTimestampEndQueryParam(r *http.Request) (*time.Time, error) {
	endStr := r.URL.Query().Get(queryParamTimestampEnd)
	if endStr == "" {
		return nil, nil
	}

	return toTime(endStr)
}

func toTime(millisString string) (*time.Time, error) {
	startMillis, err := strconv.ParseInt(millisString, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid 'timestamp' in query parameter: %v", err)
	}

	// Convert milliseconds since epoch to a time.Time object
	startTime := time.UnixMilli(startMillis)
	return &startTime, nil
}
