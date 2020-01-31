<template>
<div style="min-height: 100% !important;height: 100%;">
    <v-ons-page>
        <v-ons-toolbar style="background-color: #2D2C2C;">
            <div class="center" style="background-color: #2D2C2C;color: #FFFFFF;">{{title}}</div>
        </v-ons-toolbar>
        <div id="m1">
            <div style="text-align:left;padding:5px;">{{titleMin}}</div>
            <div class="content">
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Staff Name (Real Name):</span></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="staffName" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Staff Number:</span></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="staffNumber" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Cellphone Number:</span></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="cellphone" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:green;"> * </span><span style="color:#2D2C2C">Email:</span></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="email" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Company Name:</span></v-ons-col></v-ons-row>
                <v-ons-row>
                    <v-ons-col width="100%">
                        <v-ons-select style="width: 60%;padding:2px;" v-model="companyName">
                            <option v-for="item in companies" :value="item.key">
                                {{ item.name }}
                            </option>
                        </v-ons-select>
                    </v-ons-col>
                </v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Department:</span></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="department" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Recent health Condition Desc (e.g.: mild cough/fever/healthy..):</span></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="healthStatus" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Current Location:</span></v-ons-col></v-ons-row>
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

                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Any infected cases in your area:</span></v-ons-col></v-ons-row>
                <v-ons-row>
                    <v-ons-col width="100%"><v-ons-checkbox id="checkedPlagueArea" style="padding:5px;" :value="checkedYesPlagueValue" v-model="checkedYesPlague">Yes, there are.&nbsp;&nbsp;</v-ons-checkbox></v-ons-col>
                </v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:yellow;"> * </span><span style="color:#2D2C2C">Infected area description (if there is) :</span></v-ons-col></v-ons-row>
                <v-ons-card style="background-color:#B9B9B9;">
                    <div class="center" style="font-style: italic;font-size:12px;">{{titleAreaDesc}}:</div>
                    <div class="content">
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Date of discover:</span></v-ons-col></v-ons-row>
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
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Number of infected (approx./ball park):</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="20%"><v-ons-input float v-model="infectionApproximity" style="background-color:#B9B9B9;" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><span style="color:green;"> * </span><span style="color:#2D2C2C">Description(optional e.g.: area quarantined):</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="infectionDesc" style="background-color:#B9B9B9;" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                    </div>
                </v-ons-card>

                <!-- pass plague origin -->

                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Recently crossed or went to pandemic parameters {{originalPlagueLocName}}:</span></v-ons-col></v-ons-row>
                <v-ons-row>
                    <v-ons-col width="100%"><v-ons-checkbox style="padding:5px;" :value="plaguePassValue" v-model="plaguePass">Yes, I've been there lately.&nbsp;&nbsp;</v-ons-checkbox></v-ons-col>
                </v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:yellow;"> * </span><span style="color:#2D2C2C">If you have been to the outbreak zone fill in:</span></v-ons-col></v-ons-row>
                <v-ons-card style="background-color:#B9B9B9;">
                    <div class="center" style="font-style: italic;font-size:12px;">{{titlePassPlague}}:</div>
                    <div class="content">
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Date of trip:</span></v-ons-col></v-ons-row>
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
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Name of the area:</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><v-ons-input float v-model="passWhere" style="background-color:#B9B9B9;" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Reason for being there:</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><v-ons-input style="background-color:#B9B9B9;" float v-model="passWhat" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Train No. / Flight No.:</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><v-ons-input style="background-color:#B9B9B9;" float v-model="passHow" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                    </div>
                </v-ons-card>

                <!-- infection contact -->

                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Had contact with infected:</span></v-ons-col></v-ons-row>
                <v-ons-row>
                    <v-ons-col width="100%"><v-ons-checkbox style="padding:5px;" :value="plagueContactValue" v-model="plagueContact">Yes, Unfortunately I have.&nbsp;&nbsp;</v-ons-checkbox></v-ons-col>
                </v-ons-row>
                <v-ons-row><v-ons-col width="100%"><span style="color:yellow;"> * </span><span style="color:#2D2C2C">Details (if you have contacted with anyone infected):</span></v-ons-col></v-ons-row>
                <v-ons-card style="background-color:#B9B9B9;">
                    <div class="center" style="font-style: italic;font-size:12px;">{{titleInfectedContact}}:</div>
                    <div class="content">
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Date:</span></v-ons-col></v-ons-row>
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
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Location (city / area):</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><v-ons-input style="background-color:#B9B9B9;" float v-model="plagueContactPlace" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Is physical contact:</span></v-ons-col></v-ons-row>
                        <v-ons-row>
                            <v-ons-col width="100%" style="color:#FFFFFF;"><v-ons-checkbox style="padding:5px;" value="plagueContactValue" v-model="plagueContactClose">Yes, Unfortunately.&nbsp;&nbsp;</v-ons-checkbox></v-ons-col>
                        </v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><span style="color:green;"> * </span><span style="color:#2D2C2C">Description (optional):</span></v-ons-col></v-ons-row>
                        <v-ons-row><v-ons-col width="100%"><v-ons-input style="background-color:#B9B9B9;color:#2D2C2C;" float v-model="plagueContactDesc" class="inputs"></v-ons-input></v-ons-col></v-ons-row>
                    </div>
                </v-ons-card>
                <v-ons-row><v-ons-col width="100%"><span style="color:red;"> * </span><span style="color:#2D2C2C">Planned return to work date:</span></v-ons-col></v-ons-row>
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
                            <v-ons-button style="font-size:13px;background-color:#932323;padding:5px;padding-left:30px;padding-right:30px;" @click="submitRecord">submit record</v-ons-button>
                        </v-ons-col>
                        <v-ons-col width="50%">
                            <v-ons-button style="font-size:13px;background-color:#932323;padding:5px;padding-left:30px;padding-right:30px;" @click="resetSubmit">reset submit</v-ons-button>
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
    name: 'MyPlagueEng',
    data () {
        return {
            originalPlagueLocName: 'Wuhan (or within parameters of Hubei province)',
            title: 'Staff Temporary Information Collection System for Pandemic Event',
            titleMin: 'Staff Information Collection',
            titleAreaDesc: 'Close by area report',
            titlePassPlague: 'Details on crossing pandemic parameters',
            titleInfectedContact: 'Details on contact with infected individual',
            titleUsage: 'click on "reset submit" first if you need to submit updated info',
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
                { text: 'name of province', value: 'default' }
            ],
            selectedItemProvince: 'default',
            cities: [
                { text: 'name of city', value: 'default' }
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
                ons.notification.alert('your last report has been saved, for re-submit click reset submit first.');
                return;
            }
            if (this.staffName.trim() === '') {
                ons.notification.alert('please fill in staff name');
                return;
            }
            if (this.staffNumber.trim() === '') {
                ons.notification.alert('please fill in staff number');
                return;
            }
            if (this.cellphone.trim() === '') {
                ons.notification.alert('please fill in cellphone number');
                return;
            }
            if (this.companyName.trim() === '') {
                ons.notification.alert('please fill in company name');
                return;
            }
            if (this.department.trim() === '') {
                ons.notification.alert('please fill in department name');
                return;
            }
            if (this.healthStatus.trim() === '') {
                ons.notification.alert('please fill in recent health condition');
                return;
            }
            if (this.selectedItemProvince === 'default' || this.selectedItemCity === 'default') {
                ons.notification.alert('please select your current location');
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
                    ons.notification.alert('please fill in the approximate number of infected in your area');
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
                    ons.notification.alert('please fill in name of pandemic area you have been to');
                    return;
                }
                if (this.passWhat.trim() === '') {
                    ons.notification.alert('please fill in the reason for crossing pandemic area');
                    return;
                }
                if (this.passHow.trim() === '') {
                    ons.notification.alert('please fill in the flight/train no.');
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
                    ons.notification.alert('please fill in the location that you had contact with infected.');
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
                        ons.notification.alert("your report has been saved. thanks for your support.");
                        this.$session.set('done', 'true');
                    }
                }
            ).catch(e => { ons.notification.alert("sorry, server error. try later."); })
        },
        resetSubmit() {
            this.$session.set('done', 'false');
            ons.notification.alert("submit refreshed, please click submit your new report.");
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