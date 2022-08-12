import request from "./request";

export function login(data: any) {
  return request({
    url: "http://localhost:8080/user/login?username=${data.username}&password=${data.password}",
    method: "post",
    data,
  });
}
export function register(data: any) {
  return request({
    url: `http://localhost:8080/user/register?username=${data.username}&password=${data.password}`,
    method: "post",
    // data,
  });
}
