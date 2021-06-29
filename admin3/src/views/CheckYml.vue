<template>
  <div ref="terminalbox" style="height: 100%; background: #002833">
    <div id="terminal" ref="terminal"></div>
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
      term: null, // 保存terminal实例
      websock: null
    };
  },

  created () {

  },

  mounted () {
    let height = this.$refs.terminalbox.clientHeight
    let width = this.$refs.terminalbox.clientWidth
    // console.log(height, width)
    let rows = parseInt(height / 17 + 1, 10);//18是字体高度,根据需要自己修改
    let cols = parseInt(width / 9, 10);
    let term = new Terminal({
      rendererType: "canvas", //渲染类型
      rows: rows, //行数
      cols: cols, //不指定行数，自动回车后光标从下一行开始
      convertEol: true, //启用时，光标将设置为下一行的开头
      scrollback: 50, //终端中的回滚量
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

    // term.on('resize', function (size) {
    //   msg = { operation: "resize", cols: size.cols, rows: size.rows };
    //   _this.webSocketSend.send(JSON.stringify(msg))
    // });

    // 创建terminal实例
    term.open(this.$refs["terminal"])
    // canvas背景全屏
    var fitAddon = new FitAddon()
    term.loadAddon(fitAddon)
    fitAddon.fit()

    function runFakeTerminal (_this) {
      if (term._initialized) {
        return;
      }
      // 初始化
      term._initialized = true;
      // 支持输入与粘贴方法
      term.onData(function (data) {
        _this.webSocketSend(data)
      });

      _this.term = term;
    }
    runFakeTerminal(this)
    this.initWebSocket()
  },

  beforeUnmount () {
    this.websock.close()
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
      // this.webSocketSend(JSON.stringify(actions));
    },
    webSocketOnError () {//连接建立失败重连
      // this.initWebSocket();
    },
    webSocketOnMessage (e) { //数据接收
      // const redata = JSON.parse(e.data);
      this.term.writeln(e.data)
      // console.log(e.data)
    },
    webSocketSend (data) {//数据发送
      this.websock.send(data)
      this.term.writeln(data)
    },
    webSocketClose (e) {  //关闭
      // console.log("closed:", this.websock)
    }
  }
};
</script>
