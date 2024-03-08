<script setup>

import {computed, onMounted, reactive, ref, watch} from "vue";
import {useStore} from 'vuex'
// import {GetBiliAudio, GetBiliSearch} from "../../wailsjs/go/main/App.js";
import {LogError} from "../../wailsjs/runtime/runtime.js";

const store = useStore()

const audio_load = reactive({})

let page = ref(1)

let search_load = ref(false);

const list = ref([])

onMounted(() => {
  search(page)
})

watch(computed(()=>store.state.keyword), search)

function getAudio (o){
  // audio_load[o.bvid]=true
  // GetBiliAudio(o.bvid).then(result => {
  //   if (result.code===200){
  //     store.commit('setAudio',result.data)
  //     audio_load[o.bvid]=false
  //   }else{
  //     audio_load[o.bvid]=false
  //     LogError(result.message)
  //   }
  // }).catch(err => {
  //   audio_load[o.bvid]=false
  //   LogError(err)
  // })
}

function search() {
  // let keyword = store.state.keyword
  // if (keyword=== "" || search_load.value){
  //   return
  // }
  // search_load.value = true
  // GetBiliSearch(keyword,page.value).then(result => {
  //   if (result.code===200){
  //     list.value=result.data
  //     search_load.value = false
  //   }else{
  //     LogError(result.message)
  //     search_load.value = false
  //   }
  // }).catch(err => {
  //   LogError(err)
  //   search_load.value = false
  // })
}

</script>

<template>
<div>
  <div id="content" style="margin-top:15px;padding: 30px 40px 0 40px" v-loading="search_load">
    <el-row :gutter="20">
      <el-col :span="8" v-for="(o, index) in list" :key="o" style="">
        <el-card :body-style="{ padding: '0px' }" @click="getAudio(o)" v-loading="audio_load[o.bvid]">
          <img :src=o.pic rel="external nofollow" class="image" alt=""/>
          <div style="background-color: rgba(84,73,73,0.27);transform: translate(0, -18px);">
            <span style="font-size: 12px;color: #ffffff">{{o.view}}</span>
            <span style="font-size: 12px;color: #ffffff;margin-left: 5px">{{o.reply}}</span>
            <span style="font-size: 12px;color: #ffffff;margin-left: 5px">{{o.like}}</span>
          </div>
          <div style="height: 18px;transform: translate(0, -18px);">
            <span style="font-size: 12px;font-weight: bold">{{o.title}}</span>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</div>
</template>

<style scoped>

.image {
  width: 100%;
  height: 145px;
  display: block;
}

</style>