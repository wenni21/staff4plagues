<template>
<div id="app" style="margin-top:0px;min-height:100%;height:100%;width:100%;">
    <div id="basics" >
        <b-navbar toggleable="lg" type="dark" variant="dark">
            <b-btn variant="dark" fluid align="center" style="font-size:18px;color:#FFF;font-style:italic;">{{msg}}</b-btn>
            <b-collapse is-nav id="nav_collapse">
                <b-navbar-nav class="ml-auto">
                    <b-nav-form>
                        <b-form-input class="mr-sm-2" type="text" placeholder="Search" v-model="searchWord" />
                        <b-button class="my-2 my-sm-0" @click=searcher()>查找/刷新</b-button>
                    </b-nav-form>
                </b-navbar-nav>
            </b-collapse>
        </b-navbar>
        <b-container fluid>
            <b-row>
                <b-col cols="12" lg="12" style="padding:0px;background-color:#424242;">
                    <div style="height:800px;overflow-y:auto;">
                    <b-table striped hover dark :items="mainList" :fields="mainTableFields">
                        <template slot="detailsCheck" slot-scope="row">
                            <b-button :id="row.index" variant="secondary" style="font-size:14px;" size="sm"
                                    :press-sync="mainList.index" v-b-modal.pplDetails @click="getLiveData(row.item)">详情</b-button>
                        </template>
                    </b-table>
                    <b-pagination size="md" align="right" variant="dark" :total-rows="totalNum" v-model="currentPage" :per-page="100" @input="postFunc"></b-pagination>
                    </div>
                    <b-modal id="pplDetails" header-bg-variant="dark" header-text-variant="light" 
                             footer-bg-variant="dark" footer-text-variant="light" size="lg" centered :title="titleOfDetails">
                        <div style="height:600px;overflow-y:auto;width:100%;">
                            <b-card no-body>
                                <b-tabs card>
                                    <b-tab v-for="item in onePerson" :title="titleOfDetails" active>
                                        <b-form-group horizontal label="标签: " label-class="text-sm-right" label-for="recName">
                                            <b-form-input id="recName" readonly :placeholder="one.recordName"></b-form-input>
                                        </b-form-group>
                                        <b-form-group horizontal :label="currentLocationStatusLabel" label-class="text-sm-right" label-for="currentLocationStatus">
                                            <b-form-input id="currentLocationStatus" readonly :placeholder="one.currentLocationStatus"></b-form-input>
                                        </b-form-group>
                                        <b-table striped hover dark :items="currentLocationDetailTable" :fields="currentLocationDetailFields"></b-table>

                                        <b-form-group horizontal :label="hasPassedOutbreakZoneLabel" label-class="text-sm-right" label-for="hasPassedOutbreakZone">
                                            <b-form-input id="hasPassedOutbreakZone" readonly :placeholder="one.hasPassedOutbreakZone"></b-form-input>
                                        </b-form-group>
                                        <b-table striped hover dark :items="outbreakZoneTripTable" :fields="outbreakZoneTripFields"></b-table>

                                        <b-form-group horizontal :label="hasContactWithInfectedLabel" label-class="text-sm-right" label-for="hasContactWithInfected">
                                            <b-form-input id="hasContactWithInfected" readonly :placeholder="one.hasContactWithInfected"></b-form-input>
                                        </b-form-group>
                                        <b-table striped hover dark :items="infectedContactDetailsTable" :fields="infectedContactDetailsFields"></b-table>
                                    </b-tab>
                                </b-tabs>
                            </b-card>
                        </div>
                    </b-modal>
                </b-col>
            </b-row>
            <b-row>
                <b-col style="width:100%;padding:0px;">
                    <b-navbar type="dark" variant="dark">
                        <!-- <b-btn variant="dark" fluid align="center" style="font-size:10px;color:#FFF;font-style:italic;">Devloped by KION APAC IT - ITCS</b-btn> -->
                        <b-btn variant="dark" fluid align="center" style="font-size:10px;color:#FFF;font-style:italic;">{{footTag}}</b-btn>
                    </b-navbar>
                </b-col>
            </b-row>
        </b-container>
        <b-alert v-model="alertControl" variant="secondary" dismissible>
            {{messageAlert}}
        </b-alert>
    </div>
</div>
</template>

<script>
import axios from 'axios'
import Vue from 'vue'
import VueSessionStorage from 'vue-sessionstorage'
import Auth from './Auth.vue'
import router from '../router'
import VueRouter from 'vue-router'

Vue.use(VueSessionStorage);
Vue.use(VueRouter);

const sleep = (milliseconds) => { return new Promise(resolve => setTimeout(resolve, milliseconds)) };

const AXIOS = axios.create({
    baseURL: Auth.BASE_URL  //prod
})
const headers = {
    'Content-Type': 'application/json',
    'Auth': 'FF315A2408701FEA'
}

export default {
    name: 'BigTable',
    focusRef (ref) {
        this.$nextTick(() => {
            this.$nextTick(() => { (ref.$el || ref).focus() });
        });
    },
    data() {
        let self = this; //for $nextTick instance
        return {
            msg: '重大事件员工信息收录系统', 
            footTag: 'Developed by a developer in Jan/2020',
            yes: '是',
            no: '否',
            token: {
                username: '',
                password: '',
                company: ''
            },
            titleOfDetails: '人员详情',
            searchWord: '',
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
                currentLocationStatus: '', //number
                currentLocationDetail: { appearTime: '', approximity: '', desc: '' },
                hasPassedOutbreakZone: '', //number
                outbreakZoneTrip: { when: '', where: '', how: '', what: '' },
                hasContactWithInfected: '', //number
                infectedContactDetails: { place: '', time: '', isCloseContact: '', desc: '' }, //isCloseContact
                localWorkPolicy: ''
            },
            onePerson: [ this.one ],
            totalNum: 0,
            currentPage: 1,
            mainList: [],
            mainTableFields: [
                { label: ' ', key: 'detailsCheck', sortable: false },
                { label: '员工姓名', key: 'staffName', sortable: false },
                { label: '员工号', key: 'staffNumber', sortable: false },
                { label: '手机号', key: 'cellphone', sortable: false },
                { label: '邮箱', key: 'email', sortable: false },
                { label: '公司名称', key: 'companyName', sortable: false },
                { label: '部门名称', key: 'department', sortable: false },
                { label: '健康状况', key: 'healthStatus', sortable: false },
                { label: '当前所在地', key: 'currentLocation', sortable: false },
                { label: '同程感染', key: 'currentLatLon', sortable: false },
                { label: '复工日期', key: 'localWorkPolicy', sortable: false }
            ],
            currentLocationStatusLabel: '所在地区周围是否有疫情',
            currentLocationDetailFields: [
                { label: '出现日期', key: 'appearTime', sortable: false },
                { label: '感染者大概数量', key: 'approximity', sortable: false },
                { label: '大概情况', key: 'desc', sortable: false }
            ],
            currentLocationDetailTable: [],
            hasPassedOutbreakZoneLabel: '是否最近经过疫情发源地',
            outbreakZoneTripFields: [
                { label: '旅行经过日期', key: 'when', sortable: false },
                { label: '经过地区名称描述', key: 'where', sortable: false },
                { label: '经过所述地区原因', key: 'what', sortable: false },
                { label: '车次/航班号', key: 'how', sortable: false }
            ],
            outbreakZoneTripTable: [],
            hasContactWithInfectedLabel: '是否与感染者有过接触',
            infectedContactDetailsFields: [
                { label: '接触日期', key: 'time', sortable: false },
                { label: '接触地点', key: 'place', sortable: false },
                { label: '是否近距离接触', key: 'isCloseContact', sortable: false },
                { label: '情况描述', key: 'desc', sortable: false }
            ],
            infectedContactDetailsTable: [],
            alertControl: false,
            messageAlert: ''
        }
    },
    methods: {
        searcher() {
            this.getData();
            this.searchWord = '';
        },
        getData() {
            var filter = {
                'staffName': this.searchWord,
                'staffNumber': this.searchWord,
                'department': this.searchWord,
                'company': this.token.company,
                'cellphone': this.searchWord
            };
            AXIOS.post('/fb/' + this.currentPage, filter, { headers: headers }).then(
                response => {
                    if (response.data !== undefined && response.data !== 'NaN' && response.data.length > 0) {
                        console.log(response.data);
                        this.mainList = response.data;
                    } else {
                        this.mainList = [];
                        this.currentPage = 1;
                    }
                }
            ).catch(e => { console.log(e); this.currentPage = 1; })
        },
        postFunc() {
            AXIOS.post('/cnt', {}, { headers: headers }).then(
                response => {
                    if (response.data !== undefined) {
                        console.log(response.data);
                        this.totalNum = response.data.count;
                        getData();
                    } 
                }
            ).catch(e => { console.log(e); })
        },
        getLiveData(item) {
            this.one.index = item.index;
            this.one.staffName = item.index;
            this.one.staffNumber = item.staffNumber;
            this.one.recordTime = item.recordTime;
            this.one.recordName = item.recordName;
            this.one.cellphone = item.cellphone;
            this.one.email = item.email;
            this.one.companyName = item.companyName;
            this.one.department = item.department;
            this.one.healthStatus = item.healthStatus;
            this.one.currentLocation = item.currentLocation;
            this.one.currentLatLon = item.currentLatLon;
            if (item.currentLocationStatus === 1) {
                this.one.currentLocationStatus = this.yes;
            } else {
                this.one.currentLocationStatus = this.no;
            }
            this.one.currentLocationDetail.appearTime = item.currentLocationDetail.appearTime;
            this.one.currentLocationDetail.approximity = item.currentLocationDetail.approximity + '个人';
            this.one.currentLocationDetail.desc = item.currentLocationDetail.desc;
            this.currentLocationDetailTable = [];
            this.currentLocationDetailTable.push(this.one.currentLocationDetail);
            if (item.hasPassedOutbreakZone === 1) {
                this.one.hasPassedOutbreakZone = this.yes;
            } else {
                this.one.hasPassedOutbreakZone = this.no;
            }
            this.one.outbreakZoneTrip = item.outbreakZoneTrip;
            this.outbreakZoneTripTable = [];
            this.outbreakZoneTripTable.push(this.one.outbreakZoneTrip);
            if (item.hasContactWithInfected === 1) {
                this.one.hasContactWithInfected = this.yes;
            } else {
                this.one.hasContactWithInfected = this.no;
            }
            this.one.infectedContactDetails.place = item.infectedContactDetails.place;
            this.one.infectedContactDetails.time = item.infectedContactDetails.time;
            if (item.infectedContactDetails.isCloseContact === 1) {
                this.one.infectedContactDetails.isCloseContact = this.yes;
            } else {
                this.one.infectedContactDetails.isCloseContact = this.no;
            }
            this.one.infectedContactDetails.desc = item.infectedContactDetails.desc;
            this.infectedContactDetailsTable = [];
            this.infectedContactDetailsTable.push(this.one.infectedContactDetails);
            this.one.localWorkPolicy = item.localWorkPolicy;
        },
        loadUpAdmin() {
            console.log(this.$session.get('admin'));
            if (this.$session.get('admin') !== undefined) {
                this.token = this.$session.get('admin');
            }
        }
    },
    beforeMount() {
        console.log(this.$session.get('admin'));
        this.loadUpAdmin();
        this.getData();
    }
}
</script>

<style lang="scss">
#basics {
    font-family: Verdana, Arial, Helvetica, sans-serif;
    font-size: 14px;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: left;
    color: #2c3e50;
}
</style>
