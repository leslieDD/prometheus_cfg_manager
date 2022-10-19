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

export function getAllBaseCronApi (info) {
  return request({
    url: '/v1/base/cron/all',
    method: 'get',
    params: info
  })
}

export function getCronRulesWithSplitPage (queryInfo) {
  return request({
    url: '/v1/cron/rules/split',
    method: 'get',
    params: queryInfo
  })
}

export function postCronRule (jobInfo) {
  return request({
    url: '/v1/cron/rule',
    method: 'post',
    data: jobInfo
  })
}

export function putCronRule (jobInfo) {
  return request({
    url: '/v1/cron/rule',
    method: 'put',
    data: jobInfo
  })
}

export function deleteCronRule (jobID) {
  return request({
    url: '/v1/cron/rule',
    method: 'delete',
    params: { id: jobID }
  })
}

export function enabledCronRule (jInfo) {
  return request({
    url: '/v1/cron/rule/status',
    method: 'put',
    data: jInfo
  })
}

export function batchDeleteCronRules (ids) {
  return request({
    url: '/v1/cron/rules/selection',
    method: 'delete',
    data: ids
  })
}

export function batchEnableCronRules (ids) {
  return request({
    url: '/v1/cron/rules/enable',
    method: 'put',
    data: ids
  })
}

export function batchDisableCronRules (ids) {
  return request({
    url: '/v1/cron/rules/disable',
    method: 'put',
    data: ids
  })
}

export function cronRulesPublish () {
  return request({
    url: '/v1/cron/rules/publish',
    method: 'put'
  })
}

export function getRuleChat (query) {
  return request({
    url: '/v1/cron/rule/image',
    method: 'get',
    params: query,
  })
}

export function sendTestMail () {
  return request({
    url: '/v1/cron/send/testmail',
    method: 'post',
  })
}

export function loadtitle(rule) {
  return request({
    url: '/v1/cron/load/title',
    method: 'post',
    data: rule
  })
}