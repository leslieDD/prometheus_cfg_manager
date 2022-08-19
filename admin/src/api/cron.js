import request from '@/utils/request'

export function getBaseCronApi (info) {
  return request({
    url: '/v1/base/cron',
    method: 'get',
    params: info
  })
}

export function putBaseCronApi (info) {
  return request({
    url: '/v1/base/cron',
    method: 'put',
    data: info
  })
}

export function postBaseCronApi (info) {
  return request({
    url: '/v1/base/cron',
    method: 'post',
    data: info
  })
}

export function deleteBaseCronApi (info) {
  return request({
    url: '/v1/base/cron',
    method: 'delete',
    params: info
  })
}

export function enabledBaseCronApi (info) {
  return request({
    url: '/v1/base/cron/status',
    method: 'put',
    data: info
  })
}

