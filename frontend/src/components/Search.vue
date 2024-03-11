<script setup>

import {computed, onMounted, reactive, ref, watch} from "vue";
import {useStore} from 'vuex'
import {GetSearch,GetMusic} from "../../wailsjs/go/main/App.js";
import {LogError,LogDebug} from "../../wailsjs/runtime/runtime.js";

const store = useStore()

let search_item_load = reactive({})

let page = ref(1)

let search_load = ref(false);

// { "id": "5105986", "platform": "qq", "identifier": "001xd0HI0X9GNq", "title": "一路向北", "name": "周杰伦", "describe": "J III MP3 Player", "type": "", "cover": "https://y.gtimg.cn/music/photo_new/T002R300x300M000002MAeob3zLXwZ.jpg" }
const list = ref([])

onMounted(() => {
  search()
})

watch(computed(()=>store.state.keyword), search)

function getMisic (o){
  search_item_load[o.identifier]=true
  GetMusic(o.platform,o.identifier).then(result => {
    if (result.code===200){
      search_item_load[o.identifier]=false
      LogDebug(JSON.stringify(result.data))
      store.commit('setAudio',result.data)
    }else{
      search_item_load[o.identifier]=false
      LogError(result.message)
    }
  }).catch(err => {
    search_item_load[o.identifier]=false
    LogError(err)
  })
}

function search() {
  let keyword = store.state.keyword
  if (keyword=== "" || search_load.value){
    return
  }
  search_load.value = true

  GetSearch("qq",keyword).then(result => {
    if (result.code===200){
      search_load.value = false
      list.value=result.data
    }else{
      search_load.value = false
      LogError(result.message)
    }
  }).catch(err => {
    search_load.value = false
    LogError(err)
  })
}

</script>

<template>
  <div id="search" style="padding: 0 40px" v-loading="search_load">
    <div style="margin-top:15px;color: #272727;font-size: 18px;font-weight: bold">
      歌曲 >
    </div>
    <div id="content" style="margin-top:15px">
      <el-row :gutter="24">
        <el-col :span="12" v-for="(o, index) in list" :key="o">
          <el-card :body-style="{ padding: '0px' }" v-loading="search_item_load[o.identifier]"
                   style="height: 60px;border: none;" shadow="never" border="false" class="music-card">
            <div style="border-top:1px solid #e6e6e6"></div>
            <div class="music-area" style="width: 40px;height: 40px">
              <img class="music-image" :src=o.cover rel="external nofollow" alt="" @click="getMisic(o)"
                   style="width: 40px;height: 40px;border-radius: 5px"/>
            </div>
            <div class="music-area">
              <div style="padding-left: 10px">
                <span style="font-size: 12px">{{o.title}} - {{o.name}}</span>
              </div>
              <div style="padding-left: 10px">
                <span style="font-size: 12px;color: #808080">{{o.describe}}</span>
              </div>
              <div></div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<style scoped>

.music-card:hover {

}

.music-area{
  display: inline-block;
  vertical-align: middle;
  padding-top: 10px;
  padding-left: 5px;
}
</style>