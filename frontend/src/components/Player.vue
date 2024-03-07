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
    "lrc": "http://localhost:51730/yyxz",
    // "lrc":"[ti:烟雨行舟0] [ar:司南0] [al:烟雨行舟0] [by:0] [offset:00] [00:00.000]烟雨行舟 - 司南 [00:09.470]词：Bye【清禾音社】 [00:18.940]曲：Bye【清禾音社】 [00:28.410]编曲：Bye【清禾音社】 [00:37.880]烟雨入江南 [00:40.180]山水如墨染 [00:42.900]宛若丹青未干 [00:45.370]提笔然 点欲穿 [00:48.220]行舟临秀川 [00:50.790]画鷁推清澜 [00:53.990]缱绻怡然 [00:58.910]天色沉靛蓝 [01:01.570]波光似锦缎 [01:04.260]缀着零星白帆 [01:06.740]诗情满 千卷难 [01:09.580]渔舟齐桅杆 [01:12.250]鸳鸯凭舷栏 [01:15.300]琴瑟相伴 [01:21.050]一叶轻船 [01:23.460]一双桨悠懒 [01:26.180]一绵江风微拂素罗衫 [01:31.840]渔火点点聚散 [01:34.230]欸乃声声浅淡 [01:37.040]天近晚 [01:38.410]炊烟袅飘沿斑 [01:42.370]一叶轻船 [01:44.820]一双人倚揽 [01:47.550]一曲烟雨行舟太缓慢 [01:53.040]执手相看 [01:55.510]把酒当歌言欢 [01:58.640]红尘路漫漫 [02:02.050]愿今生与你共览 [02:51.300]暮霭渐褪暗 [02:52.970] [02:53.560]胭脂余味淡 [02:55.560] [02:56.340]我支着竹绢伞 [02:58.700]你闲摆花团扇 [03:01.590]浪儿晃曳慢 [03:03.700] [03:04.280]夜垂云流缓 [03:07.490]且吟且谈 [03:10.260] [03:13.040]一叶轻船 [03:15.500]一双桨悠懒 [03:18.300]一绵江风微拂素罗衫 [03:22.330] [03:23.680]渔火点点聚散 [03:25.720] [03:26.330]欸乃声声浅淡 [03:29.000]天近晚 [03:30.680]炊烟袅飘沿斑 [03:33.470] [03:34.290]一叶轻船 [03:36.690]一双人倚揽 [03:39.020] [03:39.580]一曲烟雨行舟太缓慢 [03:43.970] [03:44.980]执手相看 [03:47.340]把酒当歌言欢 [03:49.980] [03:50.610]红尘路漫漫 [03:54.010]欲今生与你共览 [03:56.310]一叶轻船 [03:58.020]一双月对半 [04:00.890]一帘清梦幽幽醉阑珊 [04:05.130] [04:06.180]移舟靠岸 [04:08.280] [04:09.010]案前惟剩空盏 [04:11.270] [04:11.970]莫怨良辰短 [04:15.290]曲终了韵意未完",
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
    <APlayer theme="#FE668E" ref="aplayer" style="margin: 0;padding: 0" :audio="audio"/>
  </div>
</template>

<style scoped>
</style>