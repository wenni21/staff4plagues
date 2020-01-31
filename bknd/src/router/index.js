import Vue from 'vue'
import Router from 'vue-router'
import BigTable from '../components/BigTable.vue'
import Auth from '../components/Auth.vue'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/auth',
            name: 'Auth',
            component: Auth
        },
        {
            path: '/table',
            name: 'BigTable',
            component: BigTable
        }
    ]
})
