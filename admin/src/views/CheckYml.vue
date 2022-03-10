<template>
  <div style="height: 100%">
    <el-descriptions class="margin-top" :column="2" size="mini" border>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-switch-button"></i>
          执行重启Prometheus服务
        </template>
        <div style="width: 250px">
          <el-button type="warning" @click="restart" icon="el-icon-sunny" size="small"
            >重 启</el-button
          >
        </div>
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-first-aid-kit"></i>
          执行测试配置文件（prometheus.yml）
        </template>
        <div style="width: 250px">
          <el-button type="primary" @click="check" icon="el-icon-hot-water" size="small"
            >测 试</el-button
          >
        </div>
      </el-descriptions-item>
    </el-descriptions>
    <div ref="terminalbox" style="height: 100%; background: #002833">
      <div id="terminal" ref="terminal"></div>
    </div>
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
    let rows = parseInt(height / 18, 10);//18是字体高度,根据需要自己修改
    let cols = parseInt(width / 9, 10);
    let term = new Terminal({
      rendererType: "canvas", //渲染类型
      rows: rows, //行数
      cols: cols, //不指定行数，自动回车后光标从下一行开始
      convertEol: true, //启用时，光标将设置为下一行的开头
      scrollback: 50, //终端中的回滚量
      disableStdin: false, //是否应禁用输入。
      // cursorStyle: "bar", //光标样式 underline
      cursorBlink: true, //光标闪烁
      fontSize: 16,
      theme: {
        foreground: '#ffffff', // 字体
        background: '#1b212f', // 背景色
        cursor: '#c16f93', // 设置光标
        selection: 'rgba(255, 255, 255, 0.3)',
        black: '#000000',
        brightBlack: '#808080',
        red: '#ce2f2b',
        brightRed: '#f44a47',
        green: '#00b976',
        brightGreen: '#05d289',
        yellow: '#e0d500',
        brightYellow: '#f4f628',
        magenta: '#bd37bc',
        brightMagenta: '#d86cd8',
        blue: '#1d6fca',
        brightBlue: '#358bed',
        cyan: '#00a8cf',
        brightCyan: '#19b8dd',
        white: '#e5e5e5',
        brightWhite: '#ffffff'
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
      let wsuri = `${protocol}${document.location.host}/v1/ws`
      if (!this.$store.getters.token) {
        this.$notify({
          title: '失败',
          message: 'token验证失败，需要重新登录！',
          type: 'error'
        });
        this.$router.push({ name: 'login' })
        return false
      }
      this.websock = new WebSocket(wsuri, [this.$store.getters.token])
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
      this.term.write(e.data + "\n")
      // console.log(e.data)
    },
    webSocketSend (data) {//数据发送
      this.websock.send(data)
      this.term.write(data + "\n")
    },
    webSocketClose (e) {  //关闭
      // console.log("closed:", this.websock)
    },
    restart () {
      this.websock.send("restart")
      this.term.write('send command: \x1B[1;3;31mrestart\x1B[0m\n')
    },
    check () {
      this.websock.send("check")
      this.term.write('send command: \x1B[1;3;31mcheck\x1B[0m\n')
    }
  }
};
</script>
<style scoped>
.control-btn {
  margin-bottom: 5px;
  display: flex;
  justify-content: space-around;
  flex-wrap: nowrap;
}
</style>