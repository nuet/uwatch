import request from '@/utils/request'

export function getTaskList(query) {
  return request({
    url: 'v1/uw_task',
    method: 'get',
    params: query
  })
}

export function getTask(id) {
  return request({
    url: 'v1/uw_task/' + id,
    method: 'get'
  })
}

export function getRecord(param) {
  return request({
    url: 'v1/uw_task/getRecords',
    method: 'get',
    params: param
  })
}
