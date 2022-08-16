import axios from "axios";

//1. 创建axios对象
const service = axios.create();

//2. 请求拦截器
service.interceptors.request.use(
  (config) => {
    // let loginResult = JSON.parse(localStorage.getItem("loginResult")); // 解析从localStorage里拿出的loginResult
    // if (loginResult) {
    //   const token = loginResult.token; // 取出accessToken
    //   config.headers.Authorization = `Bearer ${token}`; // 将accessToken放入到请求头里
    // }
    return config;
  },
  (error) => {
    Promise.reject(error);
  }
);

//3. 响应拦截器
service.interceptors.response.use(
  (response) => {
    if (response.status === 200) {
      return Promise.resolve(response.data);
    } else {
      return Promise.reject(response.data);
    }
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default service;
