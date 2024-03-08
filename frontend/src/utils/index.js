import {useStore} from 'vuex'

const store = useStore()

function PouloLog(msg) {
    let now = Date.now()
    const log = {
        message: msg,
        version: now.toString(),
    }
    store.commit('addLog',log)
}

export default PouloLog