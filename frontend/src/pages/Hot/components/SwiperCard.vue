<script setup lang="ts">
import { ref } from "vue";

import { Navigation, Pagination, Scrollbar, A11y } from "swiper";

import "swiper/css";
import "swiper/css/navigation";
import "swiper/css/pagination";
import "swiper/css/scrollbar";

const modules = ref([Navigation, Pagination, Scrollbar, A11y]);

const props = defineProps<{
  dataSource: any[];
}>();
</script>

<template>
  <swiper
    class="swiper-wrapper"
    :modules="modules"
    :slides-per-group="3"
    :slides-per-view="3"
    :space-between="20"
    navigation
  >
    <swiper-slide
      v-for="(o, index) in dataSource"
      :key="index"
      class="swiper-bili-slide"
    >
      <el-card :body-style="{ padding: '0px' }">
        <img :src="o.pic" class="image bili-image" alt="" />
        <div
          style="
            background-color: rgba(84, 73, 73, 0.27);
            transform: translate(0, -22px);
          "
        >
          <span style="font-size: 12px; color: #ffffff">{{ o.view }}</span>
          <span style="font-size: 12px; color: #ffffff; margin-left: 15px">{{
            o.reply
          }}</span>
          <span style="font-size: 12px; color: #ffffff; margin-left: 15px">{{
            o.like
          }}</span>
        </div>
        <div style="height: 18px; width: 253px; transform: translate(0, -20px)">
          <span
            style="
              font-size: 12px;
              font-weight: bold;
              width: 253px;
              white-space: normal;
            "
            >{{ o.title }}</span
          >
        </div>
      </el-card>
    </swiper-slide>
  </swiper>
</template>

<style scoped>
.image {
  width: 253px;
  height: 142px;
}

.platform {
  margin-top: 20px;
}

.swiper-container {
  margin-top: 12px;
}
</style>
