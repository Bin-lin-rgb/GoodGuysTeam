import request from "./request";

export function GetClassification() {
  return request({
    url: `http://localhost:8080/v1/community`,
  });
}

export function GetPostList() {
  return request({
    url: `http://localhost:8080/v1/postlist`,
  });
}

export function GetPostDetails(data:any) {
  return request({
    url: `http://localhost:8080/v1/post/${data.id}`,
  });
}

export function GetPostListWithTime(data: any) {
  return request({
    url: `http://localhost:8080/v1/postlist?latest_time=${data.nextTime}`,
  });
}

export function PostArticle(data: any) {
  return request({
    url: `http://localhost:8080/v1/post`,
    method: "post",
    data: data,
  });
}

export function GetPostListByUserId() {
  return request({
    url: `http://localhost:8080/v1/user/bloglist`,
  });
}