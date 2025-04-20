import request from '@/utils/request'

export function uplink(data) {
  return request({
    url: '/uplink',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getProductinfo
export function getProductinfo(data) {
  return request({
    url: '/getProductinfo',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getProductList
export function getProductList(data) {
  return request({
    url: '/getProductList',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getAllProductInfo
export function getAllProductInfo(data) {
  return request({
    url: '/getAllProductInfo',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getProductHistory
export function getProductHistory(data) {
  return request({
    url: '/getProductHistory',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

