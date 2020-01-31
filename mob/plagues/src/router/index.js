import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import MyPlague from '@/components/MyPlague'
import Admins from '@/components/Admins'
import AdminTable from '@/components/AdminTable'
import MyPlagueEng from '@/components/MyPlagueEng'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            name: 'HelloWorld',
            component: HelloWorld
        },
        {
            path: '/myplague/',
            name: 'MyPlague',
            component: MyPlague
        },
        {
            path: '/admin',
            name: 'Admin',
            component: Admins
        },
        {
          path: '/adminTable',
          name: 'AdminTable',
          component: AdminTable
        },
        {
          path: '/myplagueeng/',
          name: 'MyPlagueEng',
          component: MyPlagueEng
        },
    ]
})
