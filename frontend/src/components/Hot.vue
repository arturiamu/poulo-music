<script setup>
import {reactive, onMounted, onActivated} from 'vue';
import {useStore} from 'vuex'
import {GetBiliMusicRanking,GetBiliAudio} from "../../wailsjs/go/main/App.js";
import {LogError} from "../../wailsjs/runtime/runtime.js";

import { Swiper, SwiperSlide, useSwiper } from 'swiper/vue';
import Navigation  from '../../node_modules/swiper/modules/navigation/navigation.js';
import Pagination  from '../../node_modules/swiper/modules/pagination/pagination.js';
import Scrollbar  from '../../node_modules/swiper/modules/scrollbar/scrollbar.js';
import A11y  from '../../node_modules/swiper/modules/a11y/a11y.js';

import 'swiper/css';
import 'swiper/css/navigation';
import 'swiper/css/pagination';
import 'swiper/css/scrollbar';

const modules= [Navigation, Pagination, Scrollbar, A11y]

let store = useStore()

let audio_load = reactive({})

const list = reactive({
  bili:[],
  kugou:[],
  qq:[],
  netease: [],
})

function getAudio (o){
  audio_load[o.bvid]=true
  GetBiliAudio(o.bvid).then(result => {
    if (result.code===200){
      store.commit('setAudio',result.data)
      audio_load[o.bvid]=false
    }else{
      audio_load[o.bvid]=false
      LogError(result.message)
    }
  }).catch(err => {
    audio_load[o.bvid]=false
    LogError(err)
  })
}

// onActivated(() => {
//   LogError("Hot onActivated")
// })

onMounted(() => {
  getRankingData()
})

function getRankingData() {
  GetBiliMusicRanking().then(result => {
    if (result.code===200){
      list.bili=result.data.slice(0,10)
      list.kugou=result.data.slice(10,20)
      list.qq=result.data.slice(20,30)
      list.netease=result.data.slice(30,40)
    }else{
      LogError(result.message)
    }
  }).catch(err => {
    LogError(err)
  })
}

</script>

<template>
  <div id="hot" style="padding: 30px 40px 0 40px">
    <div style="font-size: 35px;color: #272727">热门音乐</div>
    <div class="platform"  style="font-size: 18px;color: #272727">哔哩哔哩</div>
    <div class="swiper swiper-bili">
      <swiper class="swiper-wrapper" :slides-per-group="3" :slides-per-view="3" :space-between="20" :modules="modules" navigation>
        <swiper-slide v-for="(o, index) in list.bili" class="swiper-bili-slide">
          <el-card :body-style="{ padding: '0px' }">
            <img :src=o.pic class="image bili-image" alt=""/>
              <div style="background-color: rgba(84,73,73,0.27);transform: translate(0, -22px);">
                <span style="font-size: 12px;color: #ffffff">{{o.view}}</span>
                <span style="font-size: 12px;color: #ffffff;margin-left: 15px">{{o.reply}}</span>
                <span style="font-size: 12px;color: #ffffff;margin-left: 15px">{{o.like}}</span>
              </div>
              <div style="height: 18px;width:253px;transform: translate(0, -20px);">
                <span style="font-size: 12px;font-weight: bold;width:253px;white-space: normal;">{{o.title}}</span>
              </div>
          </el-card>
        </swiper-slide>
      </swiper>
    </div>

    <div class="platform" style="font-size: 18px;color: #272727">酷狗音乐</div>
    <div class="swiper swiper-kugou">
      <swiper class="swiper-wrapper"  :slides-per-group="3" :slides-per-view="3" :space-between="20">
        <swiper-slide v-for="(o, index) in list.kugou" class="swiper-kugou">
          <el-card :body-style="{ padding: '0px' }">
            <img :src=o.pic class="image kugou-image" alt=""/>
            <div style="background-color: rgba(84,73,73,0.27);transform: translate(0, -22px);">
              <span style="font-size: 12px;color: #ffffff">{{o.view}}</span>
              <span style="font-size: 12px;color: #ffffff;margin-left: 15px">{{o.reply}}</span>
              <span style="font-size: 12px;color: #ffffff;margin-left: 15px">{{o.like}}</span>
            </div>
            <div style="height: 18px;width:253px;transform: translate(0, -20px);">
              <span style="font-size: 12px;font-weight: bold;width:253px;white-space: normal;">{{o.title}}</span>
            </div>
          </el-card>
        </swiper-slide>
      </swiper>
    </div>

    <div class="platform" style="font-size: 18px;color: #272727">QQ音乐</div>
    <div class="swiper swiper-qq">
      <swiper class="swiper-wrapper"  :slides-per-group="3" :slides-per-view="3" :space-between="20">
        <swiper-slide v-for="(o, index) in list.qq" class="swiper-kugou">
          <el-card :body-style="{ padding: '0px' }">
            <img :src=o.pic class="image kugou-image" alt=""/>
            <div style="background-color: rgba(84,73,73,0.27);transform: translate(0, -22px);">
              <span style="font-size: 12px;color: #ffffff">{{o.view}}</span>
              <span style="font-size: 12px;color: #ffffff;margin-left: 15px">{{o.reply}}</span>
              <span style="font-size: 12px;color: #ffffff;margin-left: 15px">{{o.like}}</span>
            </div>
            <div style="height: 18px;width:253px;transform: translate(0, -20px);">
              <span style="font-size: 12px;font-weight: bold;width:253px;white-space: normal;">{{o.title}}</span>
            </div>
          </el-card>
        </swiper-slide>
      </swiper>
    </div>

    <div class="platform" style="font-size: 18px;color: #272727">网易云音乐</div>
    <div class="swiper swiper-cloud">
      <swiper class="swiper-wrapper"  :slides-per-group="3" :slides-per-view="3" :space-between="20">
        <swiper-slide v-for="(o, index) in list.netease" class="swiper-cloud-slide">
          <el-card :body-style="{ padding: '0px' }">
            <img :src=o.pic class="image cloud-image" alt=""/>
            <div style="background-color: rgba(84,73,73,0.27);transform: translate(0, -22px);">
              <span style="font-size: 12px;color: #ffffff">{{o.view}}</span>
              <span style="font-size: 12px;color: #ffffff;margin-left: 15px">{{o.reply}}</span>
              <span style="font-size: 12px;color: #ffffff;margin-left: 15px">{{o.like}}</span>
            </div>
            <div style="height: 18px;width:253px;transform: translate(0, -20px);">
              <span style="font-size: 12px;font-weight: bold;width:253px;white-space: normal;">{{o.title}}</span>
            </div>
          </el-card>
        </swiper-slide>
      </swiper>
    </div>
  </div>
</template>

<style scoped>

.image {
  width: 253px;
  height: 142px;
}

.platform{
  margin-top: 20px;
}

.swiper-container{
  margin-top: 12px;
}

</style>
