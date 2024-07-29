// Package fsquota provides functions for working with filesystem quotas
package fsquota

import "os/user"

// User Functions

// UserQuotasSupported checks if quotas are supported on a given path
func UserQuotasSupported(path string) (supported bool, err error) {
	return userQuotasSupported(path)
}

// SetUserQuota configures a user's quota
func SetUserQuota(path string, user *user.User, limits Limits) (info *Info, err error) {
	return setUserQuota(path, user, &limits)
}

// GetUserInfo retrieves a user's quota information
func GetUserInfo(path string, user *user.User) (info *Info, err error) {
	return getUserInfo(path, user)
}

// GetUserReport retrieves a report of all user quotas present at the given path
func GetUserReport(path string) (report *Report, err error) {
	return getUserReport(path)
}

// Group functions

// GroupQuotasSupported checks if group quotas are supported on a given path
func GroupQuotasSupported(path string) (supported bool, err error) {
	return groupQuotasSupported(path)
}

// SetGroupQuota configures a group's quota
func SetGroupQuota(path string, group *user.Group, limits Limits) (info *Info, err error) {
	return setGroupQuota(path, group, &limits)
}

// GetGroupInfo retrieves a group's quota information
func GetGroupInfo(path string, group *user.Group) (info *Info, err error) {
	return getGroupInfo(path, group)
}

// GetGroupReport retrieves a report of all group quotas present at the given path
func GetGroupReport(path string) (report *Report, err error) {
	return getGroupReport(path)
}

// Project Functions

// SetProjectQuota configures a group's quota
func SetProjectQuota(path string, project *Project, limits Limits) (info *Info, err error) {
	return setProjectQuota(path, project, &limits)
}

// GetProjectInfo retrieves a projects's quota information
func GetProjectInfo(path string, project *Project) (info *Info, err error) {
	return getProjectInfo(path, project)
}

// GetProjectReport retrieves a report of all group quotas present at the given path
func GetProjectReport(path string) (report *Report, err error) {
	return getProjectReport(path)
}

// ProjectQuotasSupported checks if group quotas are supported on a given path
func ProjectQuotasSupported(path string) (supported bool, err error) {
	return projectQuotasSupported(path)
}
