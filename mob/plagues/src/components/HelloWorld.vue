<template>
<v-ons-page>
    <v-ons-toolbar style="background-color: #2D2C2C;">
        <div class="center" style="background-color: #2D2C2C;color: #FFFFFF;">{{ msg }}</div>
    </v-ons-toolbar>
    <div class="hello" style="margin-top:30px;">
        <v-ons-row>
            <v-ons-col width="50%">
                <v-ons-button style="font-size:13px;padding:5px;padding-left:30px;padding-right:30px;" modifier="outline" @click="register">员工信息录入</v-ons-button>
                <span><v-ons-button style="font-size:13px;padding:5px;padding-left:30px;padding-right:30px;" modifier="outline" @click="registerEng">ENGLISH</v-ons-button></span>
            </v-ons-col>
            <v-ons-col width="50%">
                <v-ons-button style="font-size:13px;padding:5px;padding-left:30px;padding-right:30px;" modifier="outline" @click="login">管理员登录</v-ons-button>
            </v-ons-col>
        </v-ons-row>
        <br/><br/>
        <v-ons-row>
            <!-- <img src="../assets/main.jpg" alt="plague" style="width: 100%; height: 530px;"> -->
            <img src="../assets/mainopen.jpg" alt="plague" style="width: 100%;">
        </v-ons-row>  
        <v-ons-modal :visible="loginVisible">
            <v-ons-card style="background-color:#EBEBEB;">
                <div class="center" style="font-size:15px;color:#2D2C2C;padding-top:10px;padding-bottom:10px;">管理员账号</div>
                <div class="content">
                    <br/>
                    <v-ons-row>
                        <v-ons-col width="20%"><span style="color:red;"> * </span><span style="color:#2D2C2C">用户名:</span></v-ons-col>
                        <v-ons-col width="80%"><v-ons-input style="color:#2D2C2C;width:100%;" v-model="username" class="inputs"></v-ons-input></v-ons-col>
                    </v-ons-row>
                    <br/>
                    <v-ons-row>
                        <v-ons-col width="20%"><span style="color:red;"> * </span><span style="color:#2D2C2C">密码:</span></v-ons-col>
                        <v-ons-col width="80%"><v-ons-input type="password" style="color:#2D2C2C;width:100%;" v-model="password" class="inputs"></v-ons-input></v-ons-col>
                    </v-ons-row>
                    <br/>
                    <v-ons-row>
                        <v-ons-col width="20%"><span style="color:red;"> * </span><span style="color:#2D2C2C;">公司:</span></v-ons-col>
                        <v-ons-col width="80%">
                            <v-ons-select style="width:100%;" v-model="selectedItemCompany">
                                <option v-for="item in companies" :value="item.key">
                                    {{ item.name }}
                                </option>
                            </v-ons-select>
                        </v-ons-col>
                    </v-ons-row>
                    <br/>
                    <div style="text-align:center;">
                        <v-ons-row>
                            <v-ons-col width="50%">
                                <v-ons-button style="font-size:13px;background-color:#932323;padding:5px;padding-left:30px;padding-right:30px;" @click="loginActually">进入</v-ons-button>
                            </v-ons-col>
                            <v-ons-col width="50%">
                                <v-ons-button style="font-size:13px;background-color:#932323;padding:5px;padding-left:30px;padding-right:30px;" @click="loginCancel">取消</v-ons-button>
                            </v-ons-col>
                        </v-ons-row>
                        <br/><br/>
                    </div>
                </div>
            </v-ons-card>
        </v-ons-modal>  
    </div>
</v-ons-page>
</template>

<script>
// Webpack CSS import
import 'onsenui/css/onsenui.css';
import 'onsenui/css/onsen-css-components.css';

// JS import
import Vue from 'vue';
import VueOnsen from 'vue-onsenui'; // This imports 'onsenui', so no need to import it separately

import axios from 'axios';
import VueSessionStorage from 'vue-sessionstorage';

import App from '../App';

Vue.use(VueSessionStorage);
Vue.use(VueOnsen); // VueOnsen set here as plugin to VUE. Done automatically if a call to window.Vue exists in the startup code.

const AXIOS = axios.create({
    baseURL: App.BASE_URL
});

const ons = VueOnsen;

const headers = {
    'Content-Type': 'application/json',
    'Auth': 'FF315A2408701FEA'
}

export default {
    name: 'HelloWorld',
    data() {
        return {
            msg: '重大事件员工信息收录系统',
            loginVisible: false,
            selectedItemCompany: '',
            companies: [],
            username: '',
            password: ''
        }
    },
    methods: {
        register() {
            this.$router.push('/myplague');
        },
        registerEng() {
            this.$router.push('/myplagueeng');
        },
        checkSlint() {
            AXIOS.post('/c/d', {}, { headers: headers }).then(
                response => {
                    if (response.data !== undefined && response.data !== 'NaN') {
                        console.log(response.data);
                        this.companies = response.data;
                        this.$session.set('companies', response.data);
                    } else {
                        ons.notification.alert("服务终端 请稍后再试");
                    }
                }
            ).catch(e => { console.log(e); ons.notification.alert("服务终端 请稍后再试"); })
        },
        login() {
            if (this.$session.get('adm') !== undefined) {
                this.$router.push('/admin');
            }
            this.loginVisible = true;
        },
        loginActually() {
            this.loginVisible = false;
            var adminInfo = {
                'username': this.username,
                'password': this.password,
                'company': this.selectedItemCompany
            };
            AXIOS.post('/l', adminInfo, { headers: headers }).then(
                response => {
                    if (response.data !== undefined && response.data !== 'NaN') {
                        console.log(response.data);
                        this.$session.set('admin', 'logged');
                        this.$session.set('adm', response.data);
                        this.$router.push('/admin');
                    }
                }
            ).catch(e => { ons.notification.alert("服务终端 请稍后再试"); })
        },
        loginCancel() {
            this.loginVisible = false;
        }
    },
    beforeMount() {
        // this.$session.set('admin', 'unlogged');
        this.checkSlint();
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h6 {
    font-weight: bolder;
}
.center {
    font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
    text-align: center;
    width: 100%;
    font-size: 13px;
    margin: 0px;
}
.hello {
    font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
    background-color: #EBEBEB;
    color: #2D2C2C;
    width:100%;
    max-width:100%;
    margin: 0px;
    font-size: 12px;
}
.content {
    text-align: center;
    padding: 5px;
    font-size: 12px;
}
</style>
