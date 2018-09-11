import request from '@/utils/request'

export function fetchAlarmData() {
  return request({
    url: '/v1/uw_monitor/alarm',
    method: 'get'
  })
}

export function fetchHgGraphData() {
  return request({
    url: '/v1/uw_monitor/hgGraph',
    method: 'get'
  })
}

