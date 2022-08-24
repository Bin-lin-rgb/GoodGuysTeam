import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
// ElementPlus
import ElementPlus from "unplugin-element-plus/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";
// ArcoResolver
import { ArcoResolver } from "unplugin-vue-components/resolvers";
// import { createStyleImportPlugin } from "vite-plugin-style-import";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    AutoImport({
      resolvers: [ElementPlusResolver(), ArcoResolver()],
    }),
    Components({
      resolvers: [
        ElementPlusResolver(),
        ArcoResolver({
          sideEffect: true,
        }),
      ],
    }),
    // createStyleImportPlugin({
    //   libs: [
    //     {
    //       libraryName: "@arco-design/web-vue",
    //       esModule: true,
    //       resolveStyle: (name) => {
    //         // css
    //         return `@arco-design/web-vue/es/${name}/style/css.js`;
    //         // less
    //         // return `@arco-design/web-vue/es/${name}/style/index.js`;
    //       },
    //     },
    //   ],
    // }),
    ElementPlus({
      // options
    }),
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  // server: {
  //   proxy: {
  //     "/api": {
  //       target: "http://jsonplaceholder.typicode.com",
  //       changeOrigin: true,
  //       rewrite: (path) => path.replace(/^\/api/, ""),
  //     },
  //   },
  // },
});
