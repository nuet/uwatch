import request from '@/utils/request'

export function getList(query) {
  return request({
    url: 'v1/uw_autofill',
    method: 'get',
    params: query
  })
}
export function UpdateAutoStatus(data) {
  return request({
    url: 'v1/uw_autofill/' + data.Id,
    method: 'put',
    data
  })
}
