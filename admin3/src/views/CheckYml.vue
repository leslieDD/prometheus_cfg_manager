<template>
  <div style="height: 100%; background: #002833">
    <div id="terminal" ref="terminal"></div>
    //terminal容器
  </div>
</template>

<script>
// 引入xterm，请注意这里和3.x版本的引入路径不一样
import { Terminal } from "xterm";
import "xterm/css/xterm.css";
import "xterm/lib/xterm.js";
import { FitAddon } from "xterm-addon-fit";

export default {
  name: "CheckYml",
  data () {
    return {
      shellWs: "",
      term: "", // 保存terminal实例
      rows: 40,
      cols: 100,
      urlParam: {},
      websock: null
    };
  },

  created () {

  },

  mounted () {
    let _this = this;
    // 获取容器宽高/字号大小，定义行数和列数
    this.rows = document.querySelector("#terminal").offsetHeight / 16 - 6
    this.cols = document.querySelector("#terminal").offsetWidth / 14

    let term = new Terminal({
      rendererType: "canvas", //渲染类型
      rows: parseInt(this.rows), //行数
      cols: parseInt(this.cols), // 不指定行数，自动回车后光标从下一行开始
      convertEol: true, //启用时，光标将设置为下一行的开头
      //   scrollback: 50, //终端中的回滚量
      disableStdin: false, //是否应禁用输入。
      cursorStyle: "underline", //光标样式
      cursorBlink: true, //光标闪烁
      theme: {
        foreground: "#7e9192", //字体
        background: "#002833", //背景色
        cursor: "help", //设置光标
        lineHeight: 16
      }
    });

    // 创建terminal实例
    term.open(this.$refs["terminal"])

    // 换行并输入起始符“$”
    term.prompt = () => {
      term.write("\r\n$ ")
    };
    term.prompt()

    // // canvas背景全屏
    var fitAddon = new FitAddon()
    term.loadAddon(fitAddon)
    fitAddon.fit()

    window.addEventListener("resize", resizeScreen)

    // 内容全屏显示
    function resizeScreen () {
      // 不传size

      try {
        fitAddon.fit()

        // 窗口大小改变时触发xterm的resize方法，向后端发送行列数，格式由后端决定
        // 这里不使用size默认参数，因为改变窗口大小只会改变size中的列数而不能改变行数，所以这里不使用size.clos,而直接使用获取我们根据窗口大小计算出来的行列数
        term.onResize(() => {
          _this.onSend({ Op: "resize", Cols: term.cols, Rows: term.rows })
        });
      } catch (e) {
        console.log("e", e.message)
      }
    }

    function runFakeTerminal (_this) {
      if (term._initialized) {
        return;
      }
      // 初始化
      term._initialized = true;

      term.writeln("Welcome to use Superman. ");
      term.writeln(
        `This is Web Terminal of pod\x1B[1;3;31m ${_this.urlParam.podName
        }\x1B[0m in namespace\x1B[1;3;31m ${_this.urlParam.namespace}\x1B[0m`
      );

      term.prompt();

      // / **
      //     *添加事件监听器，用于按下键时的事件。事件值包含
      //     *将在data事件以及DOM事件中发送的字符串
      //     *触发了它。
      //     * @返回一个IDisposable停止监听。
      //  * /
      //   / ** 更新：xterm 4.x（新增）
      //  *为数据事件触发时添加事件侦听器。发生这种情况
      //  *用户输入或粘贴到终端时的示例。事件值
      //  *是`string`结果的结果，在典型的设置中，应该通过
      //  *到支持pty。
      //  * @返回一个IDisposable停止监听。
      //  * /
      // 支持输入与粘贴方法
      term.onData(function (key) {
        let order = {
          Data: key,
          Op: "stdin"
        };
        _this.webSocketSend(order);
        // 为解决窗体resize方法才会向后端发送列数和行数，所以页面加载时也要触发此方法
        _this.webSocketSend({
          Op: "resize",
          Cols: parseInt(term.cols),
          Rows: parseInt(term.rows)
        });
      });

      _this.term = term;
    }
    runFakeTerminal(_this)
    this.initWebSocket()
  },

  beforeUnmount () {
    this.webSocketClose()
  },

  methods: {
    initWebSocket () { //初始化weosocket
      const protocol = document.location.protocol === 'https:' ? 'wss://' : 'ws://'
      // let wsuri = protocol + document.location.host + '/ws/'
      let wsuri = protocol + '127.0.0.1:8200/v1/ws'
      this.websock = new WebSocket(wsuri)
      this.websock.onmessage = this.webSocketOnMessage
      this.websock.onopen = this.webSocketOnOpen
      this.websock.onerror = this.webSocketOnError
      this.websock.onclose = this.webSocketClose
    },
    webSocketOnOpen () { //连接建立之后执行send方法发送数据
      let actions = { "test": "12345" };
      this.webSocketSend(JSON.stringify(actions));
    },
    webSocketOnError () {//连接建立失败重连
      this.initWebSocket();
    },
    webSocketOnMessage (e) { //数据接收
      const redata = JSON.parse(e.data);
      console.log(redata)
    },
    webSocketSend (Data) {//数据发送
      this.websock.send(Data);
    },
    webSocketClose (e) {  //关闭
      this.websock.close()
      this.websock.onclose()
    },
    //删除左右两端的空格
    trim (str) {
      // str.trim()
      return str.replace(/(^\s*)|(\s*$)/g, "");
    }
  }
};
</script>
