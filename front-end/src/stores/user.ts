import { defineStore } from "pinia";

// const defaultLoginResult = {
//   token: null,
//   user_id: null,
//   user_name: null,
// };

export const useUserStore = defineStore({
  id: "user",
  state: () => {
    return {
      token: "",
    };
  },
  actions: {
    setToken(token: any) {
      this.token = token;
    },
    clearToken() {
      this.token = "";
    }
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
