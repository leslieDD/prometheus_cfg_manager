import request from '@/utils/request'

export function clearSystemLog () {
  return request({
    url: '/v1/system/log/clear',
    method: 'post'
  })
}

export function getSystemLog (queryInfo) {
  return request({
    url: '/v1/system/log',
    method: 'get',
    params: queryInfo
  })
}

export function getSystemReocdeSetting () {
  return request({
    url: '/v1/system/log/setting',
    method: 'get'
  })
}

export function putSystemReocdeSetting (settingInfo) {
  return request({
    url: '/v1/system/log/setting',
    method: 'put',
    data: settingInfo
  })
}