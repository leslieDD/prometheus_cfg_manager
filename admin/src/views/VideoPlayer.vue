<template>
    <div>
        <video ref="videoPlayer" id="videoplayer" class="video-js"></video>
    </div>
</template>

<script>
export default {
    name: "VideoPlayer",
    props: {
        options: {
            type: Object,
            default () {
                return {};
            }
        }
    },
    data () {
        return {
            player: null
        };
    },
    created () {
        this.$nextTick( () => {
            this.player = this.$VideoJs(this.$refs.videoPlayer, this.options);
            const player = this.$VideoJs('videoplayer_html5_api');
            player.ready(function () {
                let _this = this;
                _this.playbackRate( parseFloat(4));
            });
        });
    },
    mounted () {

    },
    beforeDestroy () {
        if (this.player) {
            this.player.dispose();
        }
    }
};
</script>
