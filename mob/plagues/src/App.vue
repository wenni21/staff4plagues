<template>
    <div id="app">
        <router-view/>
    </div>
</template>

<script>
const BASE_URL = 'http://192.168.1.105:9528/plague';
export default {
    name: 'App',
    BASE_URL: BASE_URL,
    FOOT: 'Developed by a Developer in 2020',
}
</script>

<style>
#app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
}
</style>
