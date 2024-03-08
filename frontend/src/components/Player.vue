<script setup>
import {onMounted,ref,watch,computed} from "vue";
import {useStore} from 'vuex'
import APlayer from "@worstone/vue-aplayer";

const aplayer = ref(null);

const audio = ref([]);

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