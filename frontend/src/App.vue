<script setup>
import Menu from './components/Menu.vue'
import Player from './components/Player.vue'
import { RouterView } from 'vue-router'
import {useRouter} from 'vue-router'
import {onMounted} from "vue";
import {useStore} from "vuex";

const router = useRouter()
const store = useStore()

onMounted(() => {
  router.push({
    path:"/"
  })
})

</script>

<template>
  <div style="height: 680px;width: 980px">
    <div id="menu-main">
      <Menu/>
    </div>
    <div id="content-wrapper">
      <div id="player-main">
        <Player/>
      </div>
      <div id="vacuum">
        <div style="color: #c9c9c9;padding-left: 40px;padding-top:10px;font-size: 13px" v-show="router.currentRoute.value.name === 'search'">
          正在显示<span style="color: #323232">"{{store.state.keyword}}"</span>的搜索结果
        </div>
      </div>
      <div id="router-view-main">
        <RouterView v-slot="{ Component }">
          <KeepAlive>
            <component :is="Component" :key="$route.name" v-if="$route.meta.keep"></component>
          </KeepAlive>
        </RouterView>
      </div>
    </div>
  </div>
</template>


<style>

html, body {
  font-family: Arial, Helvetica, sans-serif;
}

body {
  padding: 0;
  margin: 0;
  overflow: hidden;
  height: 100%;
}

#menu-main{
  background-color: #e8e2e0;
  padding-top:60px;
  width: 198px;
  height: 100%;
  vertical-align: top;
  display: inline-block;
  --wails-draggable:drag;
  border-right: 2px solid #c3bebd;
}

#player-main{
  --wails-draggable:drag;
  width: 780px;
  position: relative;
  z-index: 2;
  border-bottom: 1px solid #fcfcfc;
}

#content-wrapper{
  display: inline-block;
}

#vacuum{
  height: 35px;
  width:780px;
  background-color: rgba(255, 255, 255, 0.95);
  position: relative;
  z-index: 2;
  backdrop-filter: blur(5px);
}

#router-view-main{
  width:780px;
  height:560px;
  overflow: scroll;
  position: relative;
  margin-top: -125px;
  z-index: 1;
  padding-top: 125px;
}

</style>
