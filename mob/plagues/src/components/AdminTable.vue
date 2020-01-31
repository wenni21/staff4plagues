<template>
    <v-ons-page>
        <v-ons-toolbar style="background-color: #2D2C2C;">
            <div class="center" style="background-color: #2D2C2C;color: #FFFFFF;font-size:12px;">{{ msg }}</div>
            <div class="right">
                <span><v-ons-toolbar-button icon="ion-navicon, material: md-save" @click="excel"></v-ons-toolbar-button></span>
            </div>
        </v-ons-toolbar>
        <div id="m1">
            <v-ons-list style="height: 100%;background-color: #EBEBEB;">
                <v-ons-list-item modifier="longdivider" v-for="item in mainList" :key="item.index" @click="seeOne(item)">
                    <span><v-ons-icon size="20px" icon="md-face"></v-ons-icon></span>&nbsp;&nbsp;
                    {{item.recordName}}
                </v-ons-list-item>
            </v-ons-list>
            <v-ons-modal :visible="oneModal">
                <div style="background-color: #B9B9B9; color: #000000; text-align: left; padding: 10px; height: 70%; overflow-y: auto;">
                    <v-ons-row><v-ons-col width="100%"><span> * </span>&nbsp;&nbsp;<span style="color:#535353;text-decoration: underline;">员工姓名:</span></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-model="one.staffName" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><span> * </span>&nbsp;&nbsp;<span style="color:#535353;text-decoration: underline;">员工工号:</span></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-model="one.staffNumber" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><span> * </span>&nbsp;&nbsp;<span style="color:#535353;text-decoration: underline;">手机号:</span></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-model="one.cellphone" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><span> * </span>&nbsp;&nbsp;<span style="color:#535353;text-decoration: underline;">Email:</span></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-model="one.email" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><span> * </span>&nbsp;&nbsp;<span style="color:#535353;text-decoration: underline;">公司:</span></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-model="one.companyName" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><span> * </span>&nbsp;&nbsp;<span style="color:#535353;text-decoration: underline;">部门:</span></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-model="one.department" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><span> * </span>&nbsp;&nbsp;<span style="color:#535353;text-decoration: underline;">现在所在城市:</span></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-model="one.currentLocation" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><span> * </span>&nbsp;&nbsp;<span style="color:#535353;text-decoration: underline;">现在健康状况:</span></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-model="one.healthStatus" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><span> * </span>&nbsp;&nbsp;<span style="color:#535353;text-decoration: underline;">是否与感染者有过接触:</span></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-if="one.hasContactWithInfected === 1" v-model="yes" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-if="one.hasContactWithInfected !== 1" v-model="no" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><span> * </span>&nbsp;&nbsp;<span style="color:#535353;text-decoration: underline;">与感染者接触详情:</span></v-ons-col></v-ons-row>
                    <v-ons-card style="background-color:#EBEBEB;">
                        <div class="content">
                            <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;">接触日期:</span></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><v-ons-input v-model="one.infectedContactDetails.time" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;">接触地点:</span></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="one.infectedContactDetails.place" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;">是否近距离接触:</span></v-ons-col></v-ons-row>
                            <v-ons-row>
                                <v-ons-row><v-ons-col width="100%"><v-ons-input v-if="one.infectedContactDetails.isCloseContact === 1" v-model="yes" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                                <v-ons-row><v-ons-col width="100%"><v-ons-input v-if="one.infectedContactDetails.isCloseContact !== 1" v-model="no" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                            </v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;">情况描述:</span></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="one.infectedContactDetails.desc" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                        </div>
                    </v-ons-card>
                    <v-ons-row><v-ons-col width="100%"><span> * </span>&nbsp;&nbsp;<span style="color:#535353;text-decoration: underline;">是否最近旅行经过疫情发源地:</span></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-if="one.hasPassedOutbreakZone === 1" v-model="yes" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-if="one.hasPassedOutbreakZone !== 1" v-model="no" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;">最近旅行经过疫情发源地详情描述:</span></v-ons-col></v-ons-row>
                    <v-ons-card style="background-color:#EBEBEB;">
                        <div class="content">
                            <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;">旅行经过日期:</span></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><v-ons-input v-model="one.outbreakZoneTrip.when" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;">经过地区名称描述:</span></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="one.outbreakZoneTrip.where" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;">经过所述地区原因:</span></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="one.outbreakZoneTrip.what" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;">车次/航班号:</span></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="one.outbreakZoneTrip.how" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                        </div>
                    </v-ons-card>
                    <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;">所在地区周围是否有疫情:</span></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-if="one.currentLocationStatus === 1" v-model="yes" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input v-if="one.currentLocationStatus !== 1" v-model="no" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                    <v-ons-card style="background-color:#EBEBEB;">
                        <div class="content">
                            <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;">出现日期:</span></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><v-ons-input v-model="one.currentLocationDetail.appearTime" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;">感染者数量(大概数量):</span></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="20%"><v-ons-input float v-model="one.currentLocationDetail.approximity" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;">大概情况描述:</span></v-ons-col></v-ons-row>
                            <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="one.currentLocationDetail.desc" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>
                        </div>
                    </v-ons-card>
                    <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#535353;text-decoration: underline;;">所在地区复工政策(复工日期):</span></v-ons-col></v-ons-row>
                    <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="one.localWorkPolicy" style="width:100%;"></v-ons-input></v-ons-col></v-ons-row>

                    <div style="text-align:center;margin-top:20px;margin-bottom:20px;">
                        <v-ons-row>
                            <v-ons-col width="100%">
                                <v-ons-button style="font-size:13px;background-color:#932323;padding:5px;padding-left:30px;padding-right:30px;" @click="oneModal = false">关闭</v-ons-button>
                            </v-ons-col>
                        </v-ons-row>
                    </div>
                </div>
            </v-ons-modal>
        </div>
        <div class="center" style="background-color: #2D2C2C;padding-top:10px;padding-bottom:15px;color: #FFFFFF;">{{titleFoot}}</div>
    </v-ons-page>
</template>

<script>
// Webpack CSS import
import 'onsenui/css/onsenui.css';
import 'onsenui/css/onsen-css-components.css';

// JS import
import Vue from 'vue';
import VueOnsen from 'vue-onsenui'; // This imports 'onsenui', so no need to import it separately

import VueSessionStorage from 'vue-sessionstorage';
import axios from 'axios';

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
    name: 'AdminTable',
    data () {
        return {
            msg: '有可能被感染的人员(经过感染区域或周围有感染者)',
            yes: '是',
            no: '否',
            mainList: [],
            admin: {},
            titleFoot: App.FOOT,
            one: {
                index: '',
                recordTime: 0,
                recordName: '',
                staffName: '',
                staffNumber: '',
                cellphone: '',
                email: '',
                companyName: '',
                department: '',
                healthStatus: '',
                currentLocation: '',
                currentLatLon: '',
                currentLocationStatus: 0,
                currentLocationDetail: { appearTime: '', approximity: 0, desc: '' },
                hasPassedOutbreakZone: 0,
                outbreakZoneTrip: { when: '', where: '', how: '', what: '' },
                hasContactWithInfected: 0,
                infectedContactDetails: { place: '', time: '', isCloseContact: 0, desc: '' },
                localWorkPolicy: ''
            },
            oneModal: false,
            searchName: '',
            deleteConf: false,
            page: 1
        }
    },
    methods: {
        excel() {
            
        },
        seeOne(oneSolo) {
            this.one = oneSolo;
            this.oneModal = true;
        },
        finishSeeOne() {
            this.oneModal = false;
        },
        searchOne() {
            AXIOS.post('/fa', {}, { headers: headers }).then(
                response => {
                    if (response.data !== undefined && response.data !== 'NaN' && response.data.length > 0) {
                        console.log(response.data);
                        this.mainList = response.data;
                    } else {
                        this.mainList = [];
                    }
                }
            ).catch(e => { ons.notification.alert("服务终端 请稍后再试"); this.page = 1; })
        },
        showBefore() {
            this.admin = this.$session.get('adm');
            AXIOS.post('/fa', {}, { headers: headers }).then(
                response => {
                    if (response.data !== undefined && response.data !== 'NaN') {
                        console.log(response.data);
                        this.mainList = response.data;
                    } else {
                        this.mainList = [];
                    }
                }
            ).catch(e => { ons.notification.alert("服务终端 请稍后再试"); })
        }
    },
    beforeMount() {
        this.showBefore();
    }
}
</script>

<style scoped>
.center {
    font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
    text-align: center;
    width: 100%;
    font-size: 13px;
    margin: 0px;
}
#m1 {
    font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
    width:100%;
    max-width:100%;
    margin: 0px;
    height: 95%;
    font-size: 12px;
}
</style>