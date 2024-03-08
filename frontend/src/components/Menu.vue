<script setup>
import { useStore } from "vuex";
import { onMounted, watch, ref, computed } from "vue";
import { useRouter } from "vue-router";
import {
  Fire,
  Like,
  Tool,
  MusicList,
  Music,
  List,
  Entertainment,
  History,
  Creative,
  RecordDisc,
} from "@icon-park/vue-next";
import { LogError } from "@wails/runtime/runtime.js";

let store = useStore();

let keyword = ref("");

const router = useRouter();
function logChange() {}

watch(
  computed(() => store.state.log),
  logChange
);

const menus = [
  {
    name: "",
    label: "私人订制",
    children: [
      { name: "hot", icon: Fire, router: "/hot", label: "热门音乐" },
      {
        name: "recommend",
        icon: Creative,
        router: "/recommend",
        label: "个性推荐",
      },
    ],
  },
  {
    name: "",
    label: "资料库",
    children: [
      { name: "recent", icon: History, router: "/recent", label: "最近添加" },
      { name: "artist", icon: Entertainment, router: "/artist", label: "艺人" },
      { name: "album", icon: RecordDisc, router: "/album", label: "专辑" },
      { name: "song", icon: Music, router: "/song", label: "歌曲" },
    ],
  },
  {
    name: "",
    label: "播放列表",
    children: [
      {
        name: "playlist",
        icon: MusicList,
        router: "/playlist",
        label: "所有播放列表",
      },
      {
        name: "i_like",
        icon: Like,
        router: "/play_list/i_like",
        label: "我喜欢",
      },
      {
        name: "test_list",
        icon: List,
        router: "/play_list/test",
        label: "测试列表",
      },
    ],
  },
  {
    name: "",
    label: "系统",
    children: [
      { name: "setting", icon: Tool, router: "/setting", label: "系统设置" },
    ],
  },
];

onMounted(() => {
  document.getElementById("hot").style.backgroundColor = "#d1ccc9";
});

function clearMenuStyle() {
  for (let i in menus) {
    for (let j in menus[i].children) {
      document.getElementById(menus[i].children[j].name).style.backgroundColor =
        "";
    }
  }
}

function chooseMenu(x) {
  LogError(x);
  clearMenuStyle();
  router.push({
    path: x.router,
  });
  document.getElementById(x.name).style.backgroundColor = "#d1ccc9";
}

function search() {
  if (keyword.value === "") {
    return;
  }
  store.commit("setKeyword", keyword.value);
  clearMenuStyle();
  if (router.currentRoute.value.name !== "search") {
    router.push({
      path: "/search",
    });
  }
}
</script>

<template>
  <div id="menu">
    <div id="search-input">
      <el-input
        placeholder="输入关键词"
        v-model="keyword"
        prefix-icon="Search"
        size="small"
        @keyup.enter.native="search()"
        border="false"
        style="color: #fe668e; margin-bottom: 20px"
      >
      </el-input>
    </div>

    <div v-for="(o, index) in menus">
      <div
        style="
          color: #a6a19f;
          font-size: 10px;
          font-weight: bold;
          padding-left: 5px;
        "
      >
        {{ o.label }}
      </div>
      <div v-for="(c, index) in o.children">
        <div
          :id="c.name"
          style="padding-left: 15px; border-radius: 5px"
          @click="chooseMenu(c)"
        >
          <el-button type="text" :icon="c.icon" style="color: #ff2a40">
          </el-button>
          <el-button
            type="text"
            style="color: #4f4d4d; font-size: 12px; margin-left: 3px"
            >{{ c.label }}
          </el-button>
        </div>
      </div>
    </div>
  </div>
  <div
    id="message"
    :style="{
      'border-top': store.state.log.message ? '1px solid #cfcbc9' : '',
    }"
  >
    {{ store.state.log.message }}
  </div>
</template>

<style scoped>
#menu {
  padding: 0 10px;
}

#message {
  display: grid;
  place-items: center;
  position: absolute;
  top: 645px;
  left: 5px;
  height: 35px;
  width: 178px;
  color: #5c5a5a;
  font-size: 13px;
  overflow: hidden;
  white-space: nowrap;
}

.el-input {
  --el-input-focus-border-color: #dbd4d3;
}

:deep(.el-input__wrapper) {
  border: 3px solid #da8984;
  background-color: #dad6d5 !important;
}
</style>
