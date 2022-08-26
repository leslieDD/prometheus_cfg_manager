<template>
  <div class="tv-show-board">
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <el-form size="mini" :rules="rtmpInfoRule" ref="rtmpInfo" :model="rtmpInfo" class="demo-form-inline">
          <el-row type="flex" class="row-bg">
            <el-col :span="12">
              <el-form-item label="拉流域名" prop="live_domain">
                <el-input style="width: 300px;" v-model="rtmpInfo.live_domain" placeholder="拉流量域名"></el-input>
              </el-form-item>
            </el-col>
            <el-form-item label="流名称" prop="live_name">
              <el-input style="width: 300px;" v-model="rtmpInfo.live_name" placeholder="流名称"></el-input>
            </el-form-item>
          </el-row>
          <el-row type="flex" class="row-bg">
            <el-col :span="12">
              <el-form-item label="鉴权密钥" prop="key">
                <el-input style="width: 300px;" v-model="rtmpInfo.key" placeholder="鉴权key"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="有效期" prop="unix_time">
                <el-input style="width: 300px;" v-model="rtmpInfo.unix_time" placeholder="有效期"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <div class="inline-btn">
                <el-form-item>
                  <el-button type="primary" plain @click="onPlay('rtmpInfo')">开始播放</el-button>
                </el-form-item>
              </div>
            </el-col>
          </el-row>
        </el-form>
        <el-form size="mini" :rules="rtmpUrlRule" ref="rtmpUrl" :model="rtmpUrl" class="demo-form-inline">
          <el-row type="flex" class="row-bg">
            <el-col :span="18">
              <el-form-item label="拉流地址" prop="url">
                <el-input style="width: 780px;" v-model="rtmpUrl.url" placeholder="拉流地址"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <div class="inline-btn">
                <el-form-item>
                  <el-button type="warning" plain @click="onPlayV2('rtmpUrl')">开始播放</el-button>
                </el-form-item>
              </div>
            </el-col>
          </el-row>
        </el-form>
      </div>
      <div class="playVideo-layout">
        <video ref="videoPlayer" id='videoplayer' class="video-js vjs-default-skin vjs-big-play-centered" />
      </div>
    </el-card>
  </div>
</template>

<script>

import videojs from 'video.js'
import 'video.js/dist/video-js.css'
// import 'videojs-flash'

import md5 from 'js-md5'

export default {
  name: "HlsVideo",
  components: {
  },
  data () {
    return {
      videoUrl: "",
      rtmpInfo: {
        live_domain: '',
        live_name: '',
        unix_time: '',
        key: '',
      },
      rtmpUrl: {
        url: '',
      },
      videoPlayer: null,
      setInterval: null,
      rtmpInfoRule: {
        live_domain: [
          { required: true, message: '请输入拉流域名', trigger: ['blue'] }
        ],
        live_name: [
          { required: true, message: '请输入流名称', trigger: ['blue'] }
        ],
        unix_time: [
          { required: true, message: '请输入时效长（单位秒）', trigger: ['blue'] }
        ],
        key: [
          { required: true, message: '请输入鉴权密钥', trigger: ['blue'] }
        ],
      },
      rtmpUrlRule: {
        url: [
          { required: true, message: '请输入鉴权密钥', trigger: ['blue'] }
        ],
      },
      options: {
        autoplay: false,
        muted: false,
        loop: false,
        preload: 'auto',
        controls: true,
        language: 'zh-CN',
        aspectRatio: '16:9',
        fluid: true,
        // sources: [
        //     {
        //         src: "",
        //         type: "video/mp4"
        //     }
        // ],
        notSupportedMessage: '此视频暂无法播放，请稍后再试',
        controlBar: {
          remainingTimeDisplay: false, //
          currentTimeDisplay: true, // 当前时间
          timeDivider: true, // 时间分割线
          durationDisplay: true, // 总时间
          progressControl: true, // 进度条

          customControlSpacer: true, //
          fullscreenToggle: true, // 全屏按钮
          volumePanel: false
        }
      },
    };
  },
  mounted () {
  },
  methods: {
    //播放
    onPlay (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.show(1);
        } else {
          return false
        }
      })
    },
    onPlayV2 (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.show(2);
        } else {
          return false
        }
      })
    },
    show (sType) {
      // let now = new Date().valueOf()
      debugger
      if (sType === 1) {
        var unixTimestamp = (parseInt(Date.now() / 1000) + parseInt(this.rtmpInfo.unix_time)).toString()
        var hash = md5.create()
        hash.update(this.rtmpInfo.live_name)
        hash.update(unixTimestamp)
        hash.update(this.rtmpInfo.key)
        var token = hash.hex()
        var videoUrl = 'http://' + this.rtmpInfo.live_domain + '/hls/' + this.rtmpInfo.live_name + '.m3u8?t=' + unixTimestamp + '&token=' + token
        this.rtmpUrl.url = videoUrl
        this.videoUrl = videoUrl
      } else if (sType === 2) {
        this.videoUrl = this.rtmpUrl.url  //  + "&now=" + now
      } else {
        this.$notify({
          title: '错误',
          message: '不支持的操作类型',
          type: 'error',
          duration: 2000,
        });
        return
      }
      this.videoPlayer = videojs(document.querySelector('#videoplayer'), {
        autoplay: true, // 是否自动播放
        controls: true, // 是否显示控件
      })
      this.videoPlayer.src({
        src: this.videoUrl,
        type: 'application/x-mpegURL',
      })
      // this.videoPlayer.play()
      // const player = videojs('videoplayer_html5_api');
      // player.ready(function () {
      //     let _this = this;
      //     _this.playbackRate( parseFloat(4));
      // });
    },
    destroy () {
      if (this.videoPlayer !== null) {
        this.videoPlayer.dispose() // dispose()会直接删除Dom元素
      }
    }
  },
  computed: {},
  watch: {},
  filters: {},
};
</script>

<style scoped>
.tv-show-board {
  /* position: absolute; */
  width: 1050px;
  /* height: 868px; */
  margin: 0 auto;
  /* left: 50%;
  top: 50%;
  width:1024px;
  height:768px;
  margin-left:-1024px;
  margin-top:-768px; */
  /* width: 1024px;
  height: 768px;
  top: calc(50% - 430px);
  left: calc(50% - 512px);   */
}

.inline-btn {
  float: right;
}

#videoplayer {
  width: 1024px;
  height: 500px
}
</style>