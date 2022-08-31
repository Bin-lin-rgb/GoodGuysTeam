import request from "./request";

export function login(data: any) {
  return request({
    url: `http://localhost:8080/v1/user/login?username=${data.username}&password=${data.password}`,
    method: "post",
    // data,
  });
}
export function register(data: any) {
  return request({
    url: `http://localhost:8080/v1/user/register?username=${data.username}&password=${data.password}`,
    method: "post",
  });
}

export function getUserInfo() {
  return request({
    url: `http://localhost:8080/v1/user`,
  });
}

export function updateProfile(data: any) {
  return request({
    url: `http://localhost:8080/v1/user/profile`,
    method: "put",
    data: data,
  });
}
