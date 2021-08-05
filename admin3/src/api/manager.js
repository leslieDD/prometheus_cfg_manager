import request from '@/utils/request'

export function getManagerGroup (queryInfo) {
  return request({
    url: '/v1/manager/groups',
    method: 'get',
    params: queryInfo
  })
}

export function getManangerGroupEnabled () {
  return request({
    url: '/v1/manager/groups/enabled',
    method: 'get'
  })
}

export function postManagerGroup (groupInfo) {
  return request({
    url: '/v1/manager/group',
    method: 'post',
    data: groupInfo
  })
}

export function putManagerGroup (groupInfo) {
  return request({
    url: '/v1/manager/group',
    method: 'put',
    data: groupInfo
  })
}

export function deleteManagerGroup (groupInfo) {
  return request({
    url: '/v1/manager/group',
    method: 'delete',
    params: groupInfo
  })
}

export function enabledManagerGroup (info) {
  return request({
    url: '/v1/manager/group/status',
    method: 'put',
    data: info
  })
}

// 
export function getManagerUser (queryInfo) {
  return request({
    url: '/v1/manager/users',
    method: 'get',
    params: queryInfo
  })
}

export function postManagerUser (userInfo) {
  return request({
    url: '/v1/manager/user',
    method: 'post',
    data: userInfo
  })
}

export function putManagerUser (userInfo) {
  return request({
    url: '/v1/manager/user',
    method: 'put',
    data: userInfo
  })
}

export function deleteManagerUser (userInfo) {
  return request({
    url: '/v1/manager/user',
    method: 'delete',
    params: userInfo
  })
}

export function enabledManagerUser (info) {
  return request({
    url: '/v1/manager/user/status',
    method: 'put',
    data: info
  })
}

export function getAllUserList () {
  return request({
    url: '/v1/manager/users/list',
    method: 'get'
    // data: info
  })
}

export function updateGroupMember (groupInfo, memberInfo) {
  return request({
    url: '/v1/manager/group/member',
    method: 'put',
    data: memberInfo,
    params: groupInfo
  })
}

export function getAllUserForGroup(groupInfo){
  return request({
    url: '/v1/manager/group/member',
    method: 'get',
    params: groupInfo
  })
}

export function getGroupList(){
  return request({
    url: '/v1/manager/groups/list',
    method: 'get'
  })
}

export function getAdminSetting(){
  return request({
    url: '/v1/manager/setting',
    method: 'get'
  })
}

export function putAdminSetting(paramsInfo){
  return request({
    url: '/v1/manager/setting',
    method: 'put',
    data: paramsInfo
  })
}