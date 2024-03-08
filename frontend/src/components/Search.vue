<script setup>
import { computed, onMounted, reactive, ref, watch } from "vue";
import { useStore } from "vuex";
import { GetSearch, GetMusic } from "@wails/go/main/App.js";
import { LogError, LogDebug } from "@wails/runtime/runtime.js";

const store = useStore();

let search_item_load = reactive({});

let page = ref(1);

let search_load = ref(false);

// { "id": "5105986", "platform": "qq", "identifier": "001xd0HI0X9GNq", "title": "一路向北", "name": "周杰伦", "describe": "J III MP3 Player", "type": "", "cover": "https://y.gtimg.cn/music/photo_new/T002R300x300M000002MAeob3zLXwZ.jpg" }
const list = ref([]);

onMounted(() => {
  search();
});

watch(
  computed(() => store.state.keyword),
  search
);

function getMisic(o) {
  search_item_load[o.identifier] = true;
  GetMusic(o.platform, o.identifier)
    .then((result) => {
      if (result.code === 200) {
        search_item_load[o.identifier] = false;
        LogDebug(JSON.stringify(result.data));
        store.commit("setAudio", result.data);
      } else {
        search_item_load[o.identifier] = false;
        LogError(result.message);
      }
    })
    .catch((err) => {
      search_item_load[o.identifier] = false;
      LogError(err);
    });
}

function search() {
  let keyword = store.state.keyword;
  if (keyword === "" || search_load.value) {
    return;
  }
  search_load.value = true;

  GetSearch("qq", keyword)
    .then((result) => {
      if (result.code === 200) {
        search_load.value = false;
        list.value = result.data;
      } else {
        search_load.value = false;
        LogError(result.message);
      }
    })
    .catch((err) => {
      search_load.value = false;
      LogError(err);
    });
}
</script>

<template>
  <div>
    <div
      id="content"
      style="margin-top: 15px; padding: 30px 40px 0 40px"
      v-loading="search_load"
    >
      <el-row :gutter="20">
        <el-col :span="8" v-for="(o, index) in list" :key="o" style="">
          <el-card
            :body-style="{ padding: '0px' }"
            @click="getMisic(o)"
            v-loading="search_item_load[o.identifier]"
          >
            <img :src="o.cover" rel="external nofollow" class="image" alt="" />
            <div style="height: 18px; transform: translate(0, -18px)">
              <span style="font-size: 12px; font-weight: bold">{{
                o.title
              }}</span>
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
