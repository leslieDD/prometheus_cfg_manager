import request from '@/utils/request'

export function getUserGroup (queryInfo) {
  return request({
    url: '/v1/manager/groups',
    method: 'get',
    params: queryInfo
  })
}

export function postUserGroup (groupInfo) {
  return request({
    url: '/v1/manager/group',
    method: 'post',
    data: groupInfo
  })
}

export function putUserGroup (groupInfo) {
  return request({
    url: '/v1/manager/group',
    method: 'put',
    data: groupInfo
  })
}

export function deleteUserGroup (groupInfo) {
  return request({
    url: '/v1/manager/group',
    method: 'delete',
    params: groupInfo
  })
}

export function enabledUserGroup (info) {
  return request({
    url: '/v1/manager/group/status',
    method: 'put',
    data: info
  })
}
