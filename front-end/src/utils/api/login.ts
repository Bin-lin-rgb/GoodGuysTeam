import request from "./request";

export function login(data: object) {
  return request({
    url: "http://localhost:9090/add",
    method: "post",
    data,
  });
}
