import { defineStore } from "pinia";

export const useUserStore = defineStore({
  id: "user",
  state: () => {
    return {
      token: "",
      username: "",
      user_id: "",
    };
  },
  actions: {
    setToken(token: any) {
      this.token = token;
    },
    clearToken() {
      this.token = "";
      this.user_id = "";
      this.username = "";
    },
    setUsername(name: any) {
      this.username = name;
    },
    setUserId(id: any) {
      this.user_id = id;
    },
  },
  // 开启数据缓存
  persist: {
    enabled: true,
    strategies: [
      {
        storage: localStorage,
        paths: ["token"],
      },
    ],
  },
});
