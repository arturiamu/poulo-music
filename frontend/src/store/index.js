import { createStore } from 'vuex'

const store = createStore({
    state () {
        return {
            keyword:"",
            audio:{},
            log:{
                message:"this is a test log",
                version:"12345",
            },
        }
    },
    mutations: {
        setKeyword (state,keyword) {
            state.keyword = keyword
        },
        setAudio (state,audio) {
            state.audio = audio
        },
        addAudio (state,log) {
            state.log = log
        },
        removeLog (state,version) {
            if (state.log.version === version) {
                state.log = {}
            }
        }
    }
})
export default store
