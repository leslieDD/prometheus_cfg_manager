import request from '@/utils/request'

export function getPriv (groupInfo) {
  return request({
    url: '/v1/manager/user/priv',
    method: 'get',
    params: groupInfo
  })
}

export function savePriv (groupInfo, privInfo) {
  return request({
    url: '/v1/manager/user/priv',
    method: 'put',
    params: groupInfo,
    data: privInfo
  })
}
