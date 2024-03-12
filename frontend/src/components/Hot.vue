<script setup>
import {onMounted, reactive,ref} from 'vue';
import {useStore} from 'vuex'
import {GetHotContent} from "../../wailsjs/go/main/App.js";
import {LogError} from "../../wailsjs/runtime/runtime.js";

import {Swiper, SwiperSlide} from 'swiper/vue';
import {Navigation,Pagination, Scrollbar,A11y} from "swiper";


import 'swiper/css';
import 'swiper/css/navigation';
import 'swiper/css/pagination';
import 'swiper/css/scrollbar';

const modules = ref([Navigation, Pagination, Scrollbar, A11y])

let store = useStore()

let audio_load = reactive({})

const list = reactive({
  //{ID: Platform:bili Identifier:BV1YK421x7Ue Title:上头！小时候听这首歌飙车飚了两万里地！ Name:真栗 Describe:《rage your dream》——动画《头文字D》片尾曲原唱：m.o.v.e翻唱：真栗混音：酸奶茉 Type: Cover:http://i0.hdslb.com/bfs/archive/062085e053107479047cdd2c41be77d212feefbc.jpg}
  bili: [],
  kugou: [],
  qq: [],
  netease: [],
})

function getAudio(o) {
  // audio_load[o.bvid] = true
  // GetBiliAudio(o.bvid).then(result => {
  //   if (result.code === 200) {
  //     store.commit('setAudio', result.data)
  //     audio_load[o.bvid] = false
  //   } else {
  //     audio_load[o.bvid] = false
  //     LogError(result.message)
  //   }
  // }).catch(err => {
  //   audio_load[o.bvid] = false
  //   LogError(err)
  // })
}

onMounted(() => {
  getRankingData()
})

function getRankingData() {
  GetHotContent("bili").then(result => {
    if (result.code === 200) {
      list.bili = result.data.slice(0, 10)
      list.kugou = result.data.slice(10, 20)
      list.qq = result.data.slice(20, 30)
      list.netease = result.data.slice(30, 40)
    } else {
      LogError(result.message)
    }
  }).catch(err => {
    LogError(err)
  })
}

</script>

<template>
  <div id="hot" style="padding: 30px 40px">
    <div style="font-size: 32px;color: #272727;font-weight: bold">热门音乐</div>
    <div style="margin-top: 25px">
      <div style="font-size: 18px;font-weight: bold;">哔哩哔哩</div>
      <div class="swiper swiper-bili">
        <swiper class="swiper-wrapper" :modules="modules" :slides-per-group="3" :slides-per-view="3" :space-between="20" navigation>
          <swiper-slide v-for="(o, index) in list.bili" :key="index" class="swiper-bili-slide">
            <el-card :body-style="{ padding: '0px' }">
              <img :src=o.cover class="image bili-image" alt=""/>
              <div style="height: 18px;width:253px;">
                <span style="font-size: 12px;font-weight: bold;width:253px;white-space: normal;">{{ o.title }}</span>
              </div>
            </el-card>
          </swiper-slide>
        </swiper>
      </div>

      <div class="platform">酷狗音乐</div>
      <div class="swiper swiper-kugou">
        <swiper class="swiper-wrapper" :modules="modules"  :slides-per-group="3" :slides-per-view="3" :space-between="20" navigation>
          <swiper-slide v-for="(o, index) in list.kugou" class="swiper-kugou" >
            <el-card :body-style="{ padding: '0px' }">
              <img :src=o.cover class="image kugou-image" alt=""/>
              <div style="height: 18px;width:253px;">
                <span style="font-size: 12px;font-weight: bold;width:253px;white-space: normal;">{{o.title}}</span>
              </div>
            </el-card>
          </swiper-slide>
        </swiper>
      </div>

      <div class="platform">QQ音乐</div>
      <div class="swiper swiper-qq">
        <swiper class="swiper-wrapper" :modules="modules"  :slides-per-group="3" :slides-per-view="3" :space-between="20" navigation>
          <swiper-slide v-for="(o, index) in list.qq" class="swiper-kugou">
            <el-card :body-style="{ padding: '0px' }">
              <img :src=o.cover class="image kugou-image" alt=""/>
              <div style="height: 18px;width:253px;">
                <span style="font-size: 12px;font-weight: bold;width:253px;white-space: normal;">{{o.title}}</span>
              </div>
            </el-card>
          </swiper-slide>
        </swiper>
      </div>

      <div class="platform">网易云音乐</div>
      <div class="swiper swiper-cloud">
        <swiper class="swiper-wrapper" :modules="modules"  :slides-per-group="3" :slides-per-view="3" :space-between="20" navigation>
          <swiper-slide v-for="(o, index) in list.netease" class="swiper-cloud-slide">
            <el-card :body-style="{ padding: '0px' }">
              <img :src=o.cover class="image cloud-image" alt=""/>
              <div style="height: 18px;width:253px;">
                <span style="font-size: 12px;font-weight: bold;width:253px;white-space: normal;">{{o.title}}</span>
              </div>
            </el-card>
          </swiper-slide>
        </swiper>
      </div>
    </div>
  </div>
</template>

<style scoped>

.image {
  width: 253px;
  height: 142px;
}

.platform {
  margin-top: 60px;
  font-size: 18px;
  font-weight: bold;
  color: #272727
}

.swiper {
  margin-top: 10px;
}

</style>
