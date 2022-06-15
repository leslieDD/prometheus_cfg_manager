<template>
    <div class="ali-player">
        <div class="main">
            <div class="video-center">
                <div v-if="!sourceUrls && !loadingVideo" class="tips">请选择视频源</div>
                <div v-if="waiting" class="tips">获取视频中，请稍等...</div>
                <div v-if="!hls" class ="video" v-loading="loadingVideo" style="background-color: #0c0b0b">
                </div>
                <video v-show ="hls"  id="video" v-loading="loadingVideo"  controls class="video" ref="video"   style="background-color: #0c0b0b"></video>

            </div>
        </div>
    </div>
</template>

<script>
    import Hls from 'hls.js'

    export default {
        name: '',
        components: {},
        props:{
            sourceUrl:{
                type:String,
                default:''
            },
            height:{
                type:String,
                default:'550px'
            }
        },
        data() {
            return {
                hls: null,
                sourceUrls:this.sourceUrl,//如果不赋值，在加载组件时会报错
                loadingVideo:false,
                waiting:false,
                reloadPlayTime:null //当视频流获取超时定时器
            }
        },
        computed: {

        },
        watch: {
            sourceUrl: {
                handler(newVal, oldVal) {
                    if(this.reloadPlayTime) { //重新播放或者播放新视频时，清空定时器
                        console.log('清空定时器...')
                        this.videoPause();
                        clearTimeout(this.reloadPlayTime);
                    }
                    if( newVal && newVal !== oldVal ) {
                        this.waiting = true
                        this.sourceUrls = newVal
                        this.playVideo()
                    }
                },
                // 代表在wacth里声明了firstName这个方法之后立即先去执行handler方法
                immediate: false
            }

        },
        created() {},
        mounted() {
        },
        methods: {
            playVideo() {
                this.$nextTick(()=>[
                    this.loadingVideo = false
                ])
                if (Hls.isSupported()) {
                    // var config = {
                    //   debug: true,
                    //   xhrSetup: function (xhr,url) {
                    //     xhr.withCredentials = true; // do send cookie
                    //     xhr.setRequestHeader("Access-Control-Allow-Headers", "*")
                    //     xhr.setRequestHeader("Access-Control-Allow-Origin","*")
                    //     xhr.setRequestHeader("Access-Control-Allow-Methods", "*")
                    //     xhr.setRequestHeader("Access-Control-Allow-Credentials","true")
                    //   }
                    // };
                    this.hls = new Hls();
                    this.hls.loadSource(this.sourceUrls);
                    this.hls.attachMedia(this.$refs.video);

                    this.hls.on(Hls.Events.MANIFEST_PARSED, (event, data) => {
//                        console.log(event, data);
                        console.log('playing...');
                        this.loadingVideo = false
                        this.waiting = false
                        this.$emit('playStatus', true) // 当点位存在播放地址时，可以截图
                        this.$refs.video.play();
                      /*  playSetTime = setTimeout(()=>{
                            clearTimeout(this.reloadPlayTime);
                        },3000)*/
                        //当正在播放时，取消定时器
                    });
                    let timerArr = []
                    let delay = 5000
                    this.hls.on(Hls.Events.ERROR, (event, data) => {
                        delay += 1000
                        console.log(event, data);
                        if(data.type === "networkError") { //网络故障
                            console.log('网络故障了...')
                            this.reloadPlayTime = setTimeout( ()=> {
                                this.$message.warning('获取视频失败，请重新播放...')
                                this.sourceUrls = ''
                                this.loadingVideo = false
                                this.waiting = false
                                this.disabledShot = true
                                this.$emit('playStatus', false)
                                this.videoPause()
                                // 监听出错事件
                            },delay)
                            timerArr.push(this.reloadPlayTime)
                         
                            if(timerArr.length > 1) {
                               //当视频播放中无法播放时，会多次执行Hls.Events.ERROR，因此需要处理一下定时器触发问题
                                for(let i =1;i< timerArr.length;i++) {
                                    clearTimeout(timerArr[i]);
                                    timerArr.splice(i,1)
                                }
                            }
                            console.log(timerArr)
                        } else  if(data.type === "mediaError") { //推流失败
//                            console.log("播放流断了....");
                        }

                    });
                } else if (this.$refs.video.canPlayType('application/vnd.apple.mpegurl')) {
                  /*  this.$refs.video.src = 'https://video-dev.github.io/streams/x36xhzz/x36xhzz.m3u8';
                    this.$refs.video.addEventListener('loadedmetadata',function() {
                        this.$refs.video.play();
                    });*/
                }
            },
            // 停止流
            videoPause() {
                if (this.hls) {
                    this.$refs.video.pause();
                    this.hls.destroy();
                    this.hls = null;
                }
            }
        },
        beforeDestroy () {
            clearTimeout(this.reloadPlayTime);
            this.hls = null;
        }
    }
</script>

<style lang="scss" scoped>
    .ali-player{
        width: 100%;
        .main {
            box-sizing: border-box;
            color: #FFFFFF;
            .video-center {
                position: relative;
                .name{
                    position: absolute;
                    left: 50%;
                    top: -20px;
                    font-size: 18px;
                    -webkit-transform: translateX(-50%);
                    transform: translateX(-50%);
                }
                .tips{
                    position: absolute;
                    top: 50%;
                    left: 50%;
                    -webkit-transform: translate(-50%, -50%);
                    transform: translate(-50%, -50%);
                    z-index: 9999;
                    opacity: 0.79;
                }
            }
        }
        #video {
            width: 100%;
            height: 500px !important;
            opacity: 0.79;
        }
        .video{
            width: 100%;
            height: 500px !important;
            opacity: 0.79;
        }
    }
</style>

