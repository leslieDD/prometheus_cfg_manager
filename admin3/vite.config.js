import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
import ViteComponents from 'vite-plugin-components'

ViteComponents({
  // relative paths to the directory to search for components.
  dirs: ['src/components'],

  // valid file extensions for components.
  extensions: ['vue'],
  // search for subdirectories
  deep: true,

  // generate `components.d.ts` global declrations, 
  // also accepts a path for custom filename
  globalComponentsDeclaration: false,

  // Allow subdirectories as namespace prefix for components.
  directoryAsNamespace: false,
  // Subdirectory paths for ignoring namespace prefixes
  // works when `directoryAsNamespace: true`
  globalNamespaces: [],
})

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
    /*
      路径别名
      若为文件系统路径必须是绝对路径的形式，否则将以别名原样呈现，不会解析为文件系统路径路径 
    */
    //'@': process.cwd()+'/src'
    //'@':path.resolve('src')
    //'@':path.resolve(__dirname, 'src')
    '@': path.resolve(__dirname, './src'),
    'comps': path.resolve(__dirname, 'src/components')
    },
  },
  plugins: [vue(), ViteComponents()],

  /*
    Project root directory/项目根目录（index.html所在位置），可以是绝对路径，也可以是相对于本配置文件的路径。
    default：process.cwd() 返回 Node.js 进程的当前工作目录
  */
  //root:process.cwd(),

  /*
    Default: /
    Base public path (应用基础请求路径) when served in development or production. Valid values include:
    
    Absolute URL pathname, e.g. /foo/
    Full URL, e.g. https://foo.com/
    Empty string or ./ (for embedded deployment)
  */
  // base: '/admin/',

  /*
    Directory to serve as plain static assets. 
    Files in this directory are served at / during dev and copied to the root of outDir during build, and are always served or copied as-is without transform. 
    The value can be either an absolute file system path or a path relative to project root.
    静态资源目录，开发模式下会自动放到 / 下，生产模式下会自动放到 outDir 根路径下。
    该目录可以配置为文件系统绝对目录或者相对于项目根目录的相对路径
  */
  publicDir:'public',

  /*
    Default: 'development' for serve, 'production' for build
    Specifying this in config will override the default mode for both serve and build. This value can also be overridden via the command line --mode option.
  */
  //mode:'',

  //vite开发服务器配置
  server: {
    host: 'localhost',
    port: 3000,
    open: true,
    strictPort: false,
    https: false,

    // 反向代理
    proxy: {
      '/v1': {
        target: 'http://127.0.0.1:8200',
        changeOrigin: true
        // rewrite: (path) => path.replace(/^\/admin/, '')
      }
    },
    //proxy: {
    //   // string shorthand
    //   '/foo': 'http://localhost:4567/foo',
    //   // with options
    //   '/api': {
    //     target: 'http://jsonplaceholder.typicode.com',
    //     changeOrigin: true,
    //     rewrite: (path) => path.replace(/^\/api/, '')
    //   },
    //   // with RegEx
    //   '^/fallback/.*': {
    //     target: 'http://jsonplaceholder.typicode.com',
    //     changeOrigin: true,
    //     rewrite: (path) => path.replace(/^\/fallback/, '')
    //   }
    // }
  },
  //生产模式打包配置
  build:{
    outDir: 'dist',//Specify the output directory (relative to project root).
  }
})