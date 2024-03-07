<script setup>
import {onMounted,ref,watch,computed} from "vue";
import {useStore} from 'vuex'
import APlayer from "@worstone/vue-aplayer";

const aplayer = ref(null);

const audio = ref([
  {
    "name": "夜曲",
    "artist": "周杰伦",
    "url": "https://ws.stream.qqmusic.qq.com/M5000024jrso28p8VA.mp3?fromtag=4&guid=wwwilingkucom&trace=0f6a2f7f8b7447f3&uin=0&vkey=8A48373F1E80532F09437949A122A00EA2011DBC5F02DE57CBAB5B00C95D7BE15BA0BF8239F69F4064644B612BACBA85397D89A511A379AA",
    "cover": "https://y.gtimg.cn/music/photo_new/T002R300x300M000002MAeob3zLXwZ.jpg",
    "lrc": "http://www.ilingku.com/api/qqmusic/v2?act=lrcgc&cid=001xd0HI0X9GNq",
  },
  {
    "name": "烟雨行舟",
    "artist": "司南",
    "url": "https://ws.stream.qqmusic.qq.com/M500001eeAlf2xkjmp.mp3?fromtag=4&guid=wwwilingkucom&trace=679da360a245a328&uin=0&vkey=F429A61694368B254B5248A90029FB70133AA9F87156148B4D897DDAC17227EFB6253C688972E694D622123F799492CC44AC43A54FDFB457",
    "cover": "https://y.gtimg.cn/music/photo_new/T002R300x300M000002EC9eB3Cl0qa.jpg",
    "lrc": "https://www.ilingku.com/api/qqmusic/v2?act=lrcgc&cid=003953y22qBMUT"
  },
  {
    "name": "Legends Never Die",
    "artist": "Against the Current",
    "url": "https://api.i-meto.com/meting/api?server=netease&type=url&id=506196018&auth=c5c21102c12896557ecf3fd43a415c050fe757a6",
    "cover": "https://api.i-meto.com/meting/api?server=netease&type=pic&id=109951163918904060&auth=ab3ada1788834c2c30fc14be2320219dbeaebc12",
    "lrc": "https://api.i-meto.com/meting/api?server=netease&type=lrc&id=506196018&auth=9038409d31ec7f40ea7888b77f0839051a2859a5"
  }
]);

let store = useStore()
function audioChange(){
  aplayer.value.clearList()
  aplayer.value.addList(store.state.audio)
}

watch(computed(()=>store.state.audio), audioChange)

onMounted(() => {

})

</script>

<template>
  <div id="aplayer">
    <APlayer theme="#FE668E" ref="aplayer" style="margin: 0;padding: 0" :audio="audio" lrcType="0"/>
  </div>
</template>

<style scoped>
</style>