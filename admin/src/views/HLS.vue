<template>
  <div class="tv-show-board">
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <el-form size="mini" :rules="hlsInfoRule" ref="hlsInfo" :model="hlsInfo" class="demo-form-inline">
          <el-row type="flex" class="row-bg">
            <el-col :span="12">
              <el-form-item label="拉流域名" prop="live_domain">
                <el-input style="width: 300px;" v-model="hlsInfo.live_domain" placeholder="拉流量域名"></el-input>
              </el-form-item>
            </el-col>
            <el-form-item label="流名称" prop="live_name">
              <el-input style="width: 300px;" v-model="hlsInfo.live_name" placeholder="流名称"></el-input>
            </el-form-item>
          </el-row>
          <el-row type="flex" class="row-bg">
            <el-col :span="12">
              <el-form-item label="鉴权密钥" prop="key">
                <el-input style="width: 300px;" v-model="hlsInfo.key" placeholder="鉴权key"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="有效期" prop="unix_time">
                <el-input style="width: 300px;" v-model="hlsInfo.unix_time" placeholder="有效期"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <div class="inline-btn">
                <el-form-item>
                  <el-button type="primary" plain @click="onPlay('hlsInfo')">开始播放</el-button>
                </el-form-item>
              </div>
            </el-col>
          </el-row>
        </el-form>
        <el-form size="mini" :rules="hlsUrlRule" ref="hlsUrl" :model="hlsUrl" class="demo-form-inline">
          <el-row type="flex" class="row-bg">
            <el-col :span="18">
              <el-form-item label="拉流地址" prop="url">
                <el-input style="width: 780px;" v-model="hlsUrl.url" placeholder="拉流地址"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <div class="inline-btn">
                <el-form-item>
                  <el-button type="warning" plain @click="onPlayV2('hlsUrl')">开始播放</el-button>
                </el-form-item>
              </div>
            </el-col>
          </el-row>
        </el-form>
      </div>
      <div class="playVideo-layout">
        <!-- 播放器 -->
        <div id="app-container" class="video-box">
          <video ref="videoElement" :src="videoUrl" id="videoElement" controls muted
            style="width: 100%; height: 100%; object-fit: fill"></video>
        </div>
        <!-- <ck-player :sourceUrl="videoUrl"  @playStatus= "changePlayStatus"></ck-player> -->
      </div>
    </el-card>
  </div>
</template>

<script>

import hlsjs from "hls.js";
import md5 from 'js-md5'
import ckPlayer from '@/views/HLSPlay.vue'

export default {
  name: "HLS",
  components: {
    ckPlayer
  },
  data () {
    return {
      videoUrl: "",
      hlsInfo: {
        live_domain: '',
        live_name: '',
        unix_time: '',
        key: '',
      },
      hlsUrl: {
        url: '',
      },
      hlsInfoRule: {
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
      hlsUrlRule: {
        url: [
          { required: true, message: '请输入鉴权密钥', trigger: ['blue'] }
        ],
      }
    };
  },
  mounted () {
    // this.show();
    this.closeVideo()
  },
  methods: {
    changePlayStatus (status) { //获取子组件的播放状态
      if (status) {
        this.disabledShot = false
      } else {
        this.disabledShot = true
        this.videoUrl = ''
        this.checkedPontsName = ''
      }
    },
    //播放
    onPlay (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          console.log(1, this.hlsInfo)
          this.show(1);
        } else {
          return false
        }
      })
    },
    onPlayV2 (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          console.log(2, this.hlsUrl)
          this.show(2);
        } else {
          return false
        }
      })
    },
    show (sType) {
      // let now = new Date().valueOf()
      if (sType === 1) {
        var hash = md5.create()
        hash.update(this.hlsInfo.live_name)
        hash.update(this.hlsInfo.unix_time)
        hash.update(this.hlsInfo.key)
        var token = hash.hex()
        var videoUrl = 'http://' + this.hlsInfo.live_domain + '/hls/' + this.hlsInfo.live_name + '.m3u8?t=' + this.hlsInfo.unix_time + '&token=' + token  // + "&now=" + now
        this.hlsUrl.url = videoUrl
        this.videoUrl = videoUrl
      } else if (sType === 2) {
        this.videoUrl = this.hlsUrl.url  //  + "&now=" + now
      } else {
        this.$notify({
          title: '错误',
          message: '不支持的操作类型',
          type: 'error',
          duration: 2000,
        });
        return
      }
      console.log(this.videoUrl)
      this.video = this.$refs.videoElement; //定义挂载点
      if (hlsjs.isSupported()) {
        var config = {
          autoStartLoad: true,
          // startPosition : -1,
          // capLevelToPlayerSize: false,
          // debug: false,
          // defaultAudioCodec: undefined,
          // initialLiveManifestSize: 1,
          // maxBufferLength: 30,
          // maxMaxBufferLength: 600,
          // maxBufferSize: 60*1000*1000,
          // maxBufferHole: 0.5,
          // maxSeekHole: 2,
          // lowBufferWatchdogPeriod: 0.5,
          // highBufferWatchdogPeriod: 3,
          // nudgeOffset: 0.1,
          // nudgeMaxRetry : 3,
          // maxFragLookUpTolerance: 0.2,
          // liveSyncDurationCount: 3,
          // liveMaxLatencyDurationCount: 10,
          // enableWorker: true,
          // enableSoftwareAES: true,
          // manifestLoadingTimeOut: 10000,
          // manifestLoadingMaxRetry: 1,
          // manifestLoadingRetryDelay: 500,
          // manifestLoadingMaxRetryTimeout : 64000,
          // startLevel: undefined,
          // levelLoadingTimeOut: 10000,
          // levelLoadingMaxRetry: 4,
          // levelLoadingRetryDelay: 500,
          // levelLoadingMaxRetryTimeout: 64000,
          // fragLoadingTimeOut: 20000,
          // fragLoadingMaxRetry: 6,
          // fragLoadingRetryDelay: 500,
          // fragLoadingMaxRetryTimeout: 64000,
          // startFragPrefetch: false,
          // appendErrorMaxRetry: 3,
          // loader: customLoader,
          // fLoader: customFragmentLoader,
          // pLoader: customPlaylistLoader,
          // xhrSetup: XMLHttpRequestSetupCallback,
          // fetchSetup: FetchSetupCallback,
          // abrController: customAbrController,
          // timelineController: TimelineController,
          // enableCEA708Captions: true,
          // stretchShortVideoTrack: false,
          // forceKeyFrameOnDiscontinuity: true,
          // abrEwmaFastLive: 5.0,
          // abrEwmaSlowLive: 9.0,
          // abrEwmaFastVoD: 4.0,
          // abrEwmaSlowVoD: 15.0,
          // abrEwmaDefaultEstimate: 500000,
          // abrBandWidthFactor: 0.8,
          // abrBandWidthUpFactor: 0.7,
          // minAutoBitrate: 0
        };
        this.hlsjs = new hlsjs(config);
        this.hlsjs.loadSource(this.videoUrl);//设置播放路径
        this.hlsjs.attachMedia(this.video);//解析到video标签上
        // console.log(this.hlsjs);
        this.hlsjs.on(hlsjs.Events.MANIFEST_PARSED, () => {
          this.video.play();
          this.$notify({
            title: '成功',
            message: '开始播放',
            type: 'success',
            duration: 1000,
          });
        });
        this.hlsjs.on(hlsjs.Events.ERROR, (event, data) => {
          // console.log(event, data);
          // // 监听出错事件
          // this.$notify({
          //   title: '错误',
          //   message: '加载失败',
          //   type: 'error'
          // });
          if (data.fatal) {
            switch (data.type) {
              // 网络错误导致的错误，据观察这种类型的错，占比最高
              case Hls.ErrorTypes.NETWORK_ERROR:
                // 根据自己业务（心跳/错误次数 需要加载还是重新请求，这里只给出简单的，加载情况作为示例）
                this.hlsjs.startLoad();
                break;
              case Hls.ErrorTypes.MEDIA_ERROR:
                // 多媒体错误
                this.hlsjs.recoverMediaError();
                break;
              default:
                this.hlsjs.destroy();
                this.$nextTick(() => {
                  // 非网络或多媒体错误，重新获取url
                  this.show(sType);
                })
                break;
            }
          }
        });
      } else {
        this.$notify({
          title: '错误',
          message: '浏览器不支持',
          type: 'error',
          duration: 1000,
        });
      }
    },
    //停止播放
    closeVideo () {
      if (this.hlsjs) {
        this.$refs.videoElement.pause();
        this.video.pause();
        this.hlsjs.destroy();
        this.hlsjs = null;
        this.$emit("closeVideo");
      }
    },
  },
  computed: {},
  watch: {},
  filters: {},
};
</script>

<style scoped>
.tv-show-board {
  position: absolute;
  width: 1024px;
  height: 768px;
  top: calc(50% - 430px);
  left: calc(50% - 512px);
}

.inline-btn {
  float: right;
}
</style>