<template>
    <div id="app">
        <router-view/>
    </div>
</template>

<script>
import Vue from 'vue'
import VueRouter from 'vue-router'
import Vuex from 'vuex'
import VueSessionStorage from 'vue-sessionstorage'
import 'es6-promise/auto'

Vue.use(VueRouter);
Vue.use(VueSessionStorage);
Vue.use(Vuex);

export default {
    name: 'app',
    data() {
        return {

        }
    },
    methods: {
        showBefore: function() {
            var t = this.$session.get('admin');
            if (t !== undefined) {
                this.$router.push('/table');
            } else {
                this.$router.push('/auth');
            }
        }
    },
    beforeMount() {
        this.showBefore();
    }
}
</script>

<style lang="scss">
#app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: left;
    color: #2c3e50;
}
</style>
