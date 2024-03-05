<script setup>
import {onMounted,ref,watch,computed} from "vue";
import {useStore} from 'vuex'
import APlayer from "@worstone/vue-aplayer";

const aplayer = ref(null);

const audio = ref([
  {
    "title": "烟雨行舟（原唱：伦桑）",
    "author": "司南",
    "url": "https://api.i-meto.com/meting/api?server=netease&type=url&id=1301884692&auth=898d1269a732530c578b63624fe8be0bdfc14205",
    "pic": "https://api.i-meto.com/meting/api?server=netease&type=pic&id=109951167056907210&auth=a4976ca0e04dc7a5592a706e1bd70c93acd1591f",
    "lrc": "http://api.i-meto.com/meting/api?server=netease&type=lrc&id=1301884692&auth=425c4730fe0caf12cf4c0eabacc3398097c8ba49"
    // "lrc": "[00:00.00] 作词 : Bye[00:01.00] 作曲 : Bye[00:21.95]编曲：Bye[00:23.58]混音：MR鱼[00:25.24]母带：MR曾经[00:27.09]封面：小草清清[00:38.07]烟雨入江南[00:40.33]山水如墨染[00:43.02]宛若丹青未干[00:44.93]提笔然 点欲穿[00:48.59]行舟临秀川[00:50.89]画鹢推清澜[00:54.06]缱绻怡然[00:58.96]天色沉靛蓝[01:01.64]波光似锦缎[01:04.31]缀着零星白帆[01:06.79]诗情满 千卷难[01:09.79]渔舟齐桅杆[01:12.36]鸳鸯凭舷栏[01:15.37]琴瑟相伴[01:21.05]一叶轻船[01:23.41]一双桨悠懒[01:26.28]一绵江风微拂素罗衫[01:31.95]渔火点点聚散[01:34.38]欸乃声声浅淡[01:37.14]天近晚[01:38.74]炊烟袅飘沿斑[01:42.41]一叶轻船[01:44.73]一双人倚揽[01:47.70]一曲烟雨行舟太缓慢[01:53.16]执手相看[01:55.59]把酒当歌言欢[01:58.73]红尘路漫漫[02:02.05]愿今生与你共览[02:51.49]暮霭渐褪暗[02:53.70]胭脂余味淡[02:56.30]我支着竹绢伞[02:58.75]你闲摆花团扇[03:01.82]浪儿晃曳慢[03:04.30]夜垂云流缓[03:07.54]且吟且谈[03:13.08]一叶轻船[03:15.52]一双桨悠懒[03:18.40]一绵江风微拂素罗衫[03:23.88]渔火点点聚散[03:26.38]欸乃声声浅淡[03:29.10]天近晚[03:30.79]炊烟袅飘沿斑[03:34.57]一叶轻船[03:36.78]一双人倚揽[03:39.69]一雨烟雨行舟太缓慢[03:45.01]执手相看[03:47.48]把酒当歌言欢[03:50.72]红尘路漫漫[03:54.18]欲今生与你共览[03:57.64]一叶轻船[03:58.32]一双月对半[04:01.13]一帘清梦幽幽醉阑珊[04:06.21]移舟靠岸[04:09.10]案前惟剩空盏[04:12.14]莫怨良辰短[04:15.47]曲终了韵意未完[04:20.47]*第一次制作，感谢支持。"
  },
  {
    "title": "Legends Never Die",
    "author": "Against the Current",
    "url": "https://api.i-meto.com/meting/api?server=netease&type=url&id=506196018&auth=c5c21102c12896557ecf3fd43a415c050fe757a6",
    "pic": "https://api.i-meto.com/meting/api?server=netease&type=pic&id=109951163918904060&auth=ab3ada1788834c2c30fc14be2320219dbeaebc12",
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
<!--    <APlayer theme="#FE668E" ref="aplayer" style="margin: 0;padding: 0"/>-->
    <APlayer theme="#FE668E" ref="aplayer" style="margin: 0;padding: 0" :audio="audio"/>
  </div>
</template>

<style scoped>
</style>