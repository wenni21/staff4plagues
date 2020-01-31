<template>
    <div id="app">
        <div id="authBasics">
          <h6 style="font-style: italic;color: #1C1C1C;font-size: 30px;">重大事件员工信息收录系统 v1.0</h6><br/>
          <b-form @submit="onSubmit" @reset="onReset" v-if="show" style="width:350px;text-align:left;">
              <b-form-group label="用户名:" label-for="username">
                  <b-form-input id="username" type="text" v-model="form.username" required placeholder="输入您的用户名"></b-form-input>
              </b-form-group>
              <b-form-group label="密码:" label-for="password">
                  <b-form-input id="password" type="password" v-model="form.password" required placeholder="输入您的密码"></b-form-input>
              </b-form-group>
              <br/><br/>
              <b-button type="submit" variant="primary">提交</b-button>
              <b-button type="reset" variant="danger">取消</b-button>
            </b-form>
            <b-modal id="modal-messenger" size="md" v-model="mainMessanger" centered :title="messTitle">
                <div><b-alert variant="info" show>{{messageDisplay}}</b-alert></div>
            </b-modal>
        </div>
    </div>
</template>

<script>
import 'es6-promise/auto'
import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.use(VueRouter);
import Vuex from 'vuex'
Vue.use(Vuex);
import axios from 'axios'
import VueSessionStorage from 'vue-sessionstorage'
Vue.use(VueSessionStorage);

const BASE_URL = 'http://127.0.0.1:9528/plague'  //prod
const AXIOS = axios.create({
    baseURL: BASE_URL  //prod
})
const headers = {
    'Content-Type': 'application/json',
    'Auth': 'FF315A2408701FEA'
}
export default {
    data () {
        return {
            BASE_URL: BASE_URL,
            form: {
                username: '',
                password: '',
            },
            show: true,
            userInfo: '',
            mainMessanger: false,
            messageDisplay: '',
            messTitle: ''
        }
    },
    methods: {
        onSubmit (evt) {
            evt.preventDefault();
            var params = {
                'userName': this.form.username,
                'password': this.form.password,
                'company': ''
            };
            AXIOS.post('/l', params, { headers: headers }).then(response => {
                if (response.data !== undefined) {
                    console.log(response.data);
                    this.$session.set('admin', response.data);
                    this.$router.push('/table/');
                } else {
                    console.log(JSON.stringify(response.data));
                    this.messTitle = '认证失败';
                    this.messageDisplay = '您的用户名密码认证失败';
                    this.mainMessanger = true;
                }
            }).catch(e => { console.log(e); })
        },
        onReset (evt) {
            evt.preventDefault();
            this.form.username = '';
            this.form.password = '';
            this.show = false;
            this.$nextTick(() => { this.show = true });
        }
    }
}
</script>
<style lang="scss">
h5, h6 {
    font-weight: bold;
    margin-top: 20px;
    margin-bottom: 10px;
}
#authBasics {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: left;
    color: #2c3e50;
    margin: 0 auto;
    margin-top:10%;
    width: 500px;
    height: 600px;
}
</style>
