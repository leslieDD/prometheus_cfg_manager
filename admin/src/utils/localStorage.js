const storage_key = 'userInfo'
const job_search_type = 'job_search_type'
const ip_search_type = 'ip_search_type'

export function getStorageUserInfo() {
  const getLocalData = localStorage.getItem(storage_key)
  const jsonObj = JSON.parse(getLocalData)
  return jsonObj
}

export function saveStorageUserInfo(jsonData) {
  const str_jsonData = JSON.stringify(jsonData)
  localStorage.setItem(storage_key, str_jsonData)
}

export function removeStorageUserInfo() {
  localStorage.removeItem(storage_key)
}

export function getStorageToken() {
  const getLocalData = localStorage.getItem(storage_key)
  if (!getLocalData) {
    return null
  }
  const jsonObj = JSON.parse(getLocalData)
  if (!jsonObj.session) {
    return null
  }
  return jsonObj.session.token
}

export function getStorageVal(key) {
  return localStorage.getItem(key)
}

export function setStorageVal(key, value) {
  localStorage.setItem(key, value)
}