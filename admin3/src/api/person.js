import request from '@/utils/request'

export function loadTxt () {
  return request({
    url: '/v1/txt/programer/say',
    method: 'get'
  })
}
