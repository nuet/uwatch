import fetch from '@/utils/request'

export function fetchUserList() {
  return fetch({
    url: '/v1/system_user/getList',
    method: 'get'
  })
}

export function addUser(data) {
  return fetch({
    url: '/v1/system_user',
    method: 'post',
    data
  })
}

export function delUser(id) {
  return fetch({
    url: '/v1/system_user/' + id,
    method: 'delete'
  })
}

export function updateUser(data) {
  return fetch({
    url: '/v1/system_user/' + data.Id,
    method: 'put',
    data
  })
}

export function getUserList() {
  return fetch({
    url: '/v1/system_user/getUserList',
    method: 'get'
  })
}

export function getDepartmentList() {
  return fetch({
    url: '/v1/system_user/getDepartmentList',
    method: 'get'
  })
}

export function getRole() {
  return fetch({
    url: '/v1/system_user/getUserRole',
    method: 'get'
  })
}

export function getTeamList() {
  return fetch({
    url: '/v1/system_user/getTeamList',
    method: 'get'
  })
}
