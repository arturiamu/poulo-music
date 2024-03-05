import {useStore} from 'vuex'

const store = useStore()
function Log() {
    const log = {
        message: "this is a test log",
        version: "12345",
    }
    store.commit('addAudio', log)
}