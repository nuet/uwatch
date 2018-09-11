import request from '@/utils/request'

export function postAutoFill(data) {
  return request({
    url: '/v1/uw_autofill',
    method: 'post',
    data
  })
}

export function putAutoFill(data) {
  return request({
    url: '/v1/uw_autofill/' + data.Id,
    method: 'put',
    data
  })
}

export function fetchEditAutoFill(id) {
  return request({
    url: '/v1/uw_autofill/' + id,
    method: 'get'
  })
}

export function delAutoFill(id) {
  return request({
    url: '/v1/uw_autofill/' + id,
    method: 'delete'
  })
}

export function detectionAutoFill(id) {
  return request({
    url: '/v1/uw_autofill/detection/' + id,
    method: 'get'
  })
}
