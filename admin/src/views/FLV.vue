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
        <video id="videoElement" class="centeredVideo" controls width="1024" height="428">
          Your browser is too old which doesn't support HTML5 video.
        </video>
      </div>
    </el-card>
  </div>
</template>

<script>

import flvjs from 'flv.js'
import md5 from 'js-md5'
import ckPlayer from '@/views/HLSPlay.vue'

export default {
  name: "FLV",
  components: {
    ckPlayer
  },
  data() {
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
      flvPlayer: null,
      setInterval: null,
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
  mounted() {
  },
  methods: {
    //播放
    onPlay(formName){
      this.$refs[formName].validate((valid) => {
        if (valid) {
          console.log(1, this.hlsInfo)
          this.show(1);
        } else {
          return false
        }
      })
    },
    onPlayV2(formName){
      this.$refs[formName].validate((valid) => {
        if (valid) {
          console.log(2, this.hlsUrl)
          this.show(2);
        } else {
          return false
        }
      })
    },
    show(sType) {
      // let now = new Date().valueOf()
      if (sType===1) {
        var unixTimestamp = (parseInt(Date.now() / 1000) + parseInt(this.hlsInfo.unix_time)).toString()
        var hash = md5.create()
        hash.update(this.hlsInfo.live_name)
        hash.update(unixTimestamp)
        hash.update(this.hlsInfo.key)
        var token = hash.hex()
        var videoUrl = 'http://' + this.hlsInfo.live_domain + '/live/flv?app=live&t=' + unixTimestamp + '&token=' + token + '&stream=' + this.hlsInfo.live_name
        this.hlsUrl.url = videoUrl
        this.videoUrl = videoUrl
      } else if (sType === 2) {
        this.videoUrl = this.hlsUrl.url  //  + "&now=" + now
      } else {
        this.$notify({
          title: '错误',
          message: '不支持的操作类型',
          type: 'error'
        });
        return
      }
      this.flv_destroy()
      var vElement = document.getElementById('videoElement');
      if (flvjs.isSupported()) {
          this.flvPlayer = flvjs.createPlayer({
              type: 'flv',
              enableWorker: true,     //浏览器端开启flv.js的worker,多进程运行flv.js
              isLive: true,           //直播模式
              hasAudio: true,        //关闭音频             
              hasVideo: true,
              stashInitialSize: 128,  
              enableStashBuffer: false, //播放flv时，设置是否启用播放缓存，只在直播起作用。
              url: this.videoUrl,
          });
          this.flvPlayer.attachMediaElement(vElement)
          this.flvPlayer.load() //加载
      }
      // this.setInterval = setInterval(function () {
      //     vElement.playbackRate = 1
      //     console.log("时延校正判断");
      //     if (!vElement.buffered.length) {
      //         return
      //     }
      //     var end = vElement.buffered.end(0)
      //     var diff = end - vElement.currentTime
      //     console.log(diff)
      //     if (5 <= diff && diff <=60) {
      //         console.log("二倍速")
      //         vElement.playbackRate = 2
      //     }
      //     if (diff > 60) {
      //         console.log("跳帧")
      //         vElement.currentTime = end
      //     }
      // }, 2500)
    },
    flv_start() {
      this.flvPlayer.play()
    },
    flv_pause() {
      this.flvPlayer.pause()
    },
    flv_destroy() {
      // if (this.setInterval !== null) {
      //   clearInterval(this.setInterval)
      //   this.setInterval = null
      // }
      if (this.flvPlayer === null) {
        return
      }
      this.flvPlayer.pause()
      this.flvPlayer.unload()
      this.flvPlayer.detachMediaElement()
      this.flvPlayer.destroy()
      this.flvPlayer = null
    },
    flv_seekto() {
      this.flvPlayer.currentTime = parseFloat(document.getElementsByName('seekpoint')[0].value)
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
  margin:0 auto;
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
</style>