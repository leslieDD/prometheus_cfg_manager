import request from '@/utils/request'

export function uploadMachines (uploadData) {
  return request({
    url: '/v1/upload/machines',
    method: 'post',
    data: uploadData
  })
}