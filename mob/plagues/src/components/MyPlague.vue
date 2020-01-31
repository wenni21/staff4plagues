<template>
<div style="min-height: 100% !important;height: 100%;">
    <v-ons-page>
        <v-ons-toolbar style="background-color: #2D2C2C;">
            <div class="center" style="background-color: #2D2C2C;color: #FFFFFF;">{{title}}</div>
        </v-ons-toolbar>
        <div id="m1">
            <div style="text-align:left;padding:5px;">{{titleMin}}</div>
            <div class="content">
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">员工姓名(您的真实姓名):</span></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="staffName" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">员工工号:</span></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="staffNumber" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">手机号码:</span></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="cellphone" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:green;"> * </span><span style="color:#2D2C2C">电子邮箱:</span></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="email" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">公司名称:</span></v-ons-col></v-ons-row>
                <v-ons-row>
                    <v-ons-col width="100%">
                        <v-ons-select style="width: 60%;padding:2px;" v-model="companyName">
                            <option v-for="item in companies" :value="item.key">
                                {{ item.name }}
                            </option>
                        </v-ons-select>
                    </v-ons-col>
                </v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">部门名称:</span></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="department" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">最近健康状况描述(如:有时咳嗽/未发烧/健康):</span></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="healthStatus" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">员工目前所在地:</span></v-ons-col></v-ons-row>
                <v-ons-row>
                    <v-ons-col width="50%">
                        <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;" v-model="selectedItemProvince" @change="citiesNow($event)">
                            <option v-for="item in provinces" :value="item.value">
                                {{ item.text }}
                            </option>
                        </v-ons-select>
                    </v-ons-col>
                    <v-ons-col width="50%">
                        <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;" v-model="selectedItemCity">
                            <option v-for="item in cities" :value="item.value">
                                {{ item.text }}
                            </option>
                        </v-ons-select>
                    </v-ons-col>
                </v-ons-row><br/>

                <!-- plague details -->

                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">所在地区周围是否有疫情:</span></v-ons-col></v-ons-row>
                <v-ons-row>
                    <v-ons-col width="100%"><v-ons-checkbox id="checkedPlagueArea" style="padding:5px;" :value="checkedYesPlagueValue" v-model="checkedYesPlague">所在地区周围有疫情&nbsp;&nbsp;</v-ons-checkbox></v-ons-col>
                </v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:yellow;"> * </span><span style="color:#2D2C2C">所在地区周围疫情描述（如果有疫情请详细说明）:</span></v-ons-col></v-ons-row>
                <v-ons-card style="background-color:#B9B9B9;">
                    <div class="center" style="font-style: italic;font-size:12px;">{{titleAreaDesc}}:</div>
                    <div class="content">
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">出现日期:</span></v-ons-col></v-ons-row>
                        <v-ons-row>
                            <v-ons-col width="20%">
                                <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;margin:2px;" v-model="selectedItemAppearYear">
                                    <option v-for="item in passYear" :value="item.value">
                                        {{ item.text }}
                                    </option>
                                </v-ons-select>
                            </v-ons-col>
                            <v-ons-col width="20%">
                                <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;margin:2px;" v-model="selectedItemAppearMonth">
                                    <option v-for="item in passMonth" :value="item.value">
                                        {{ item.text }}
                                    </option>
                                </v-ons-select>
                            </v-ons-col>
                            <v-ons-col width="20%">
                                <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;margin:2px;" v-model="selectedItemAppearDay">
                                    <option v-for="item in passDay" :value="item.value">
                                        {{ item.text }}
                                    </option>
                                </v-ons-select>
                            </v-ons-col>
                        </v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">感染者数量(大概数量):</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="20%"><v-ons-input float v-model="infectionApproximity" style="background-color:#B9B9B9;" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><span style="color:green;"> * </span><span style="color:#2D2C2C">大概情况描述(可选填 如:小区已设隔离检查):</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="infectionDesc" style="background-color:#B9B9B9;" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                    </div>
                </v-ons-card>

                <!-- pass plague origin -->

                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">是否最近旅行经过疫情发源地 {{originalPlagueLocName}}:</span></v-ons-col></v-ons-row>
                <v-ons-row>
                    <v-ons-col width="100%"><v-ons-checkbox style="padding:5px;" :value="plaguePassValue" v-model="plaguePass">最近旅行经过疫情发源地&nbsp;&nbsp;</v-ons-checkbox></v-ons-col>
                </v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:yellow;"> * </span><span style="color:#2D2C2C">最近旅行经过疫情发源地 详情描述（如果最近经过则填写）:</span></v-ons-col></v-ons-row>
                <v-ons-card style="background-color:#B9B9B9;">
                    <div class="center" style="font-style: italic;font-size:12px;">{{titlePassPlague}}:</div>
                    <div class="content">
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">旅行经过日期:</span></v-ons-col></v-ons-row>
                        <v-ons-row>
                            <v-ons-col width="20%">
                                <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;margin:2px;" v-model="selectedItemPassYear">
                                    <option v-for="item in passYear" :value="item.value">
                                        {{ item.text }}
                                    </option>
                                </v-ons-select>
                            </v-ons-col>
                            <v-ons-col width="20%">
                                <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;margin:2px;" v-model="selectedItemPassMonth">
                                    <option v-for="item in passMonth" :value="item.value">
                                        {{ item.text }}
                                    </option>
                                </v-ons-select>
                            </v-ons-col>
                            <v-ons-col width="20%">
                                <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;margin:2px;" v-model="selectedItemPassDay">
                                    <option v-for="item in passDay" :value="item.value">
                                        {{ item.text }}
                                    </option>
                                </v-ons-select>
                            </v-ons-col>
                        </v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">经过地区名称描述:</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="passWhere" style="background-color:#B9B9B9;" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">经过所述地区原因:</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><v-ons-input style="background-color:#B9B9B9;" float v-model="passWhat" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">车次/航班号:</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><v-ons-input style="background-color:#B9B9B9;" float v-model="passHow" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                    </div>
                </v-ons-card>

                <!-- infection contact -->

                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">是否最近与感染者接触:</span></v-ons-col></v-ons-row>
                <v-ons-row>
                    <v-ons-col width="100%"><v-ons-checkbox style="padding:5px;" :value="plagueContactValue" v-model="plagueContact">最近与感染者接触&nbsp;&nbsp;</v-ons-checkbox></v-ons-col>
                </v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:yellow;"> * </span><span style="color:#2D2C2C">与感染者接触详情（如有接触则填写）:</span></v-ons-col></v-ons-row>
                <v-ons-card style="background-color:#B9B9B9;">
                    <div class="center" style="font-style: italic;font-size:12px;">{{titleInfectedContact}}:</div>
                    <div class="content">
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">接触日期:</span></v-ons-col></v-ons-row>
                        <v-ons-row>
                            <v-ons-col width="20%">
                                <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;margin:2px;" v-model="selectedItemContactYear">
                                    <option v-for="item in passYear" :value="item.value">
                                        {{ item.text }}
                                    </option>
                                </v-ons-select>
                            </v-ons-col>
                            <v-ons-col width="20%">
                                <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;margin:2px;" v-model="selectedItemContactMonth">
                                    <option v-for="item in passMonth" :value="item.value">
                                        {{ item.text }}
                                    </option>
                                </v-ons-select>
                            </v-ons-col>
                            <v-ons-col width="20%">
                                <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;margin:2px;" v-model="selectedItemContactDay">
                                    <option v-for="item in passDay" :value="item.value">
                                        {{ item.text }}
                                    </option>
                                </v-ons-select>
                            </v-ons-col>
                        </v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">接触地点:</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><v-ons-input style="background-color:#B9B9B9;" float v-model="plagueContactPlace" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">是否近距离接触:</span></v-ons-col></v-ons-row>
                        <v-ons-row>
                            <v-ons-col width="100%" style="color:#FFFFFF;"><v-ons-checkbox style="padding:5px;" value="plagueContactValue" v-model="plagueContactClose">与感染者近距离接触&nbsp;&nbsp;</v-ons-checkbox></v-ons-col>
                        </v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><span style="color:green;"> * </span><span style="color:#2D2C2C">情况描述(可选):</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><v-ons-input style="background-color:#B9B9B9;color:#2D2C2C;" float v-model="plagueContactDesc" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                    </div>
                </v-ons-card>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">所在地区复工政策(复工日期):</span></v-ons-col></v-ons-row>
                <v-ons-row>
                    <v-ons-col width="20%">
                        <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;margin:2px;" v-model="selectedItemReturnYear" disabled>
                            <option v-for="item in passYear" :value="item.value">
                                {{ item.text }}
                            </option>
                        </v-ons-select>
                    </v-ons-col>    
                    <v-ons-col width="20%">
                        <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;margin:2px;" v-model="selectedItemReturnMonth">
                            <option v-for="item in passMonth" :value="item.value">
                                {{ item.text }}
                            </option>
                        </v-ons-select>
                    </v-ons-col>
                    <v-ons-col width="20%">
                        <v-ons-select style="width: 90%;background-color: #FFFFFF;padding:2px;margin:2px;" v-model="selectedItemReturnDay">
                            <option v-for="item in passDay" :value="item.value">
                                {{ item.text }}
                            </option>
                        </v-ons-select>
                    </v-ons-col>
                </v-ons-row>
                <br/><br/>
                <v-ons-row><v-ons-col width="100%"><span> * </span><span style="color:#2D2C2C">{{titleUsage}}</span></v-ons-col></v-ons-row>
                <br/><br/>
                <div style="text-align:center;">
                    <v-ons-row>
                        <v-ons-col width="50%">
                            <v-ons-button style="font-size:13px;background-color:#932323;padding:5px;padding-left:30px;padding-right:30px;" @click="submitRecord">提交记录</v-ons-button>
                        </v-ons-col>
                        <v-ons-col width="50%">
                            <v-ons-button style="font-size:13px;background-color:#932323;padding:5px;padding-left:30px;padding-right:30px;" @click="resetSubmit">提交重置</v-ons-button>
                        </v-ons-col>
                    </v-ons-row>
                    <br/><br/>
                </div>
            </div>
        <div class="center" style="background-color: #2D2C2C;padding-top:10px;padding-bottom:10px;color: #FFFFFF;">{{titleFoot}}</div>
        </div>
    </v-ons-page>
</div>
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

const ons = VueOnsen;

const AXIOS = axios.create({
    baseURL: App.BASE_URL
})

const GET_ALL_PROVINCE_URL = 'https://restapi.amap.com/v3/config/district?key=418be98113e39497de778af65096f011'
var GET_ALL_CITIES_URL = 'https://restapi.amap.com/v3/config/district?keywords=#&subdistrict=1&key=418be98113e39497de778af65096f011';

const headers = {
    'Content-Type': 'application/json',
    'Auth': 'FF315A2408701FEA'
}

export default {
    name: 'MyPlague',
    data () {
        return {
            originalPlagueLocName: '武汉（或湖北其他地区）',
            title: '重大事件员工信息收录系统',
            titleMin: '员工信息录入 (重新提交請先點擊 提交重置 再點擊 提交记录)',
            titleAreaDesc: '所在地区附近情况',
            titlePassPlague: '最近经过疫情始发地详情',
            titleInfectedContact: '与感染者接触详情',
            titleUsage: '如需重新提交请先点击"提交重置"后再次提交',
            titleFoot: App.FOOT,
            passYear: [ 
                { text: '2019', value: '2019' },
                { text: '2020', value: '2020' }
            ],
            passMonth: [ 
                { text: '01', value: '01' },
                { text: '02', value: '02' },
                { text: '03', value: '03' },
                { text: '04', value: '04' },
                { text: '05', value: '05' },
                { text: '06', value: '06' },
                { text: '07', value: '07' },
                { text: '08', value: '08' },
                { text: '09', value: '09' },
                { text: '10', value: '10' },
                { text: '11', value: '11' },
                { text: '12', value: '12' },
            ],
            passDay: [ 
                { text: '01', value: '01' },
                { text: '02', value: '02' },
                { text: '03', value: '03' },
                { text: '04', value: '04' },
                { text: '05', value: '05' },
                { text: '06', value: '06' },
                { text: '07', value: '07' },
                { text: '08', value: '08' },
                { text: '09', value: '09' },
                { text: '10', value: '10' },
                { text: '11', value: '11' },
                { text: '12', value: '12' },
                { text: '13', value: '13' },
                { text: '14', value: '14' },
                { text: '15', value: '15' },
                { text: '16', value: '16' },
                { text: '17', value: '17' },
                { text: '18', value: '18' },
                { text: '19', value: '19' },
                { text: '20', value: '20' },
                { text: '21', value: '21' },
                { text: '22', value: '22' },
                { text: '23', value: '23' },
                { text: '24', value: '24' },
                { text: '25', value: '25' },
                { text: '26', value: '26' },
                { text: '27', value: '27' },
                { text: '28', value: '28' },
                { text: '29', value: '29' },
                { text: '30', value: '30' },
                { text: '31', value: '31' },
            ],
            staffName: '',
            staffNumber: '',
            cellphone: '',
            email: '',
            companyName: '',
            department: '',
            healthStatus: '',
            currentLocation: '',
            provinces: [
                { text: '请选择您当前所在省', value: 'default' }
            ],
            selectedItemProvince: 'default',
            cities: [
                { text: '请选择您当前所在市', value: 'default' }
            ],
            selectedItemCity: 'default',
            checkedYesPlague: ['yesPlagueArea'],
            checkedYesPlagueValue: 'yesPlagueArea',
            selectedItemAppearYear: '2020',
            selectedItemAppearMonth: '01',
            selectedItemAppearDay: '01',
            infectionAppearTime: '',
            infectionApproximity: '0',
            infectionDesc: '',
            plaguePass: ['plaguePass'],
            plaguePassValue: 'plaguePass',
            selectedItemPassYear: '2020',
            selectedItemPassMonth: '01',
            selectedItemPassDay: '01',
            passWhen: '',
            passWhere: '',
            passWhat: '',
            passHow: '',
            plagueContact: ['plagueContacted'],
            plagueContactValue: 'plagueContacted',
            selectedItemContactYear: '2020',
            selectedItemContactMonth: '01',
            selectedItemContactDay: '01',
            plagueContactTime: '',
            plagueContactPlace: '',
            plagueContactClose: ['plagueCloseContact'],
            plagueContactCloseValue: 'plagueCloseContact',
            plagueContactDesc: '',
            selectedItemReturnYear: '2020',
            selectedItemReturnMonth: '02',
            selectedItemReturnDay: '10',
            localWorkPolicy: '',
            EveryoneInfo: {
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
            companies: []
        }
    },
    methods: {
        submitRecord() {
            if (this.$session.get('done') !== undefined && this.$session.get('done') === 'true') {
                ons.notification.alert('您的资料已提交 如需要重新提交请先点击 “提交重置”');
                return;
            }
            if (this.staffName.trim() === '') {
                ons.notification.alert('请填写员工姓名');
                return;
            }
            if (this.staffNumber.trim() === '') {
                ons.notification.alert('请填写员工工号');
                return;
            }
            if (this.cellphone.trim() === '') {
                ons.notification.alert('请填写员工手机号');
                return;
            }
            if (this.companyName.trim() === '') {
                ons.notification.alert('请填写公司名称');
                return;
            }
            if (this.department.trim() === '') {
                ons.notification.alert('请填写部门名称');
                return;
            }
            if (this.healthStatus.trim() === '') {
                ons.notification.alert('请填写健康状况');
                return;
            }
            if (this.selectedItemProvince === 'default' || this.selectedItemCity === 'default') {
                ons.notification.alert('请选择当前所在省市');
                return;
            }
            this.localWorkPolicy = '2020-' + this.selectedItemReturnMonth + '-' + this.selectedItemReturnDay
            this.currentLocation = this.selectedItemProvince + ',' + this.selectedItemCity;
            var currentLocationStatus = 0;
            var localDetailTime = '';
            var localDetailApprox = 0;
            var localDetailDesc = '';
            if (this.checkedYesPlague[0] !== undefined) {
                this.infectionAppearTime = this.selectedItemAppearYear + '-' + this.selectedItemAppearMonth + '-' + this.selectedItemAppearDay;
                localDetailTime = this.infectionAppearTime;
                if (Number(this.infectionApproximity.trim()) <= 0) {
                    ons.notification.alert('请填写周围疫情感染者大概数量');
                    return;
                }
                localDetailApprox = Number(this.infectionApproximity.trim());
                localDetailDesc = this.infectionDesc.trim();
                currentLocationStatus = 1;
            }
            var hasPassedOutbreakZone = 0;
            var passingWhen = '';
            var passingWhere = '';
            var passingWhat = '';
            var passingHow = '';
            if (this.plaguePass[0] !== undefined) {
                if (this.passWhere.trim() === '') {
                    ons.notification.alert('请填写经过地区名称');
                    return;
                }
                if (this.passWhat.trim() === '') {
                    ons.notification.alert('请填写经过所述地区原因');
                    return;
                }
                if (this.passHow.trim() === '') {
                    ons.notification.alert('请填写经过地区出行方式');
                    return;
                }
                this.passWhen = this.selectedItemPassYear + '-' + this.selectedItemPassMonth + '-' + this.selectedItemPassDay;
                passingWhen = this.passWhen;
                passingWhere = this.passWhere.trim();
                passingWhat = this.passWhat.trim();
                passingHow = this.passHow.trim();
                hasPassedOutbreakZone = 1;
            }
            var hasContactWithInfected = 0;
            var contactTime = '';
            var contactPlace = '';
            var contactClosely = 0;
            var contactDesc = '';
            if (this.plagueContact[0] !== undefined) {
                if (this.plagueContactPlace.trim() === '') {
                    ons.notification.alert('请填写与感染者接触地点');
                    return;
                }
                this.plagueContactTime = this.selectedItemContactYear + '-' + this.selectedItemContactMonth + "-" + this.selectedItemContactDay;
                contactTime = this.plagueContactTime;
                contactPlace = this.plagueContactPlace.trim();
                if (this.plagueContactClose[0] !== undefined) {
                    contactClosely = 1;
                }
                contactDesc = this.plagueContactDesc;
                hasContactWithInfected = 1;
            }
            this.EveryoneInfo.staffName = this.staffName.trim();
            this.EveryoneInfo.staffNumber = this.staffNumber.trim();
            this.EveryoneInfo.cellphone = this.cellphone.trim();
            this.EveryoneInfo.email = this.email.trim();
            this.EveryoneInfo.companyName = this.companyName.trim();
            this.EveryoneInfo.department = this.department.trim();
            this.EveryoneInfo.healthStatus = this.healthStatus.trim();
            this.EveryoneInfo.currentLocation = this.currentLocation;
            this.EveryoneInfo.currentLocationStatus = currentLocationStatus;
            this.EveryoneInfo.currentLocationDetail.appearTime = localDetailTime;
            this.EveryoneInfo.currentLocationDetail.approximity = localDetailApprox;
            this.EveryoneInfo.currentLocationDetail.desc = localDetailDesc;
            this.EveryoneInfo.hasPassedOutbreakZone = hasPassedOutbreakZone;
            this.EveryoneInfo.outbreakZoneTrip.when = passingWhen;
            this.EveryoneInfo.outbreakZoneTrip.where = passingWhere;
            this.EveryoneInfo.outbreakZoneTrip.what = passingWhat;
            this.EveryoneInfo.outbreakZoneTrip.how = passingHow;
            this.EveryoneInfo.hasContactWithInfected = hasContactWithInfected;
            this.EveryoneInfo.infectedContactDetails.time = contactTime;
            this.EveryoneInfo.infectedContactDetails.place = contactPlace;
            this.EveryoneInfo.infectedContactDetails.isCloseContact = contactClosely;
            this.EveryoneInfo.infectedContactDetails.desc = contactDesc;
            this.EveryoneInfo.localWorkPolicy = this.localWorkPolicy.trim();
            AXIOS.post('/inputs/e', this.EveryoneInfo, { headers: headers }).then(
                response => {
                    if (response.data !== undefined && response.data === 's') {
                        ons.notification.alert("您的资料提交成功 谢谢支持");
                        this.$session.set('done', 'true');
                    }
                }
            ).catch(e => { ons.notification.alert("对不起 提交失败 请稍后重试"); })
        },
        resetSubmit() {
            this.$session.set('done', 'false');
            ons.notification.alert("已刷新 您可以再次提交");
        },
        citiesNow(event) {
            var a = axios.create();
            console.log(event.target.value);
            var url = GET_ALL_CITIES_URL.replace('#', this.selectedItemProvince);
            a.get(url).then(response => {
                console.log(response.data);
                if (response.data !== undefined) {
                    var feedbk = response.data;
                    if (feedbk.status === '1') {
                        var districts = feedbk.districts[0].districts;
                        districts.forEach(e => {
                            this.cities.push({text: e.name, value: e.name});
                        });
                        console.log(this.cities);
                    }
                }
            })
        }
    },
    beforeMount() {
            this.companies = this.$session.get('companies');
            var a = axios.create();
            a.get(GET_ALL_PROVINCE_URL).then(response => {
                console.log(response.data);
                if (response.data !== undefined) {
                    var feedbk = response.data;
                    if (feedbk.status === '1') {
                        var districts = feedbk.districts[0].districts;
                        districts.forEach(e => {
                            this.provinces.push({text: e.name, value: e.name});
                        });
                        console.log(this.provinces);
                    }
                }
            })
        }
}

</script>

<style>
.center {
    font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
    text-align: center;
    width: 100%;
    font-size: 13px;
    margin: 0px;
}
.content {
    text-align: left;
    padding: 5px;
    font-size: 12px;
}
.inputs {
    background-color: #EBEBEB;
    color: #000000;
    padding: 2px;
    width: 97%;
    margin: 2px;
    font-size: 15px;
}
#m1 {
    font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
    background-color: #EBEBEB;
    color: #2D2C2C;
    width:100%;
    max-width:100%;
    margin: 0px;
    font-size: 12px;
}
</style>