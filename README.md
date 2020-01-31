如何部署:

需要環境 google golang, MongoDB

mongoDB: https://docs.mongodb.com/manual/installation/
golang:  https://golang.org/dl/

配置:
在plague目錄下的 plague/PlaConf.go 有數據庫連接配置
MongoDB默認database: plague_of_corona
MongoDB默認collection:
everyone_info （員工錄入信息表）
admins_info   （管理員表）
companies_info（公司名稱表）
plague_info   （默認同程信息表）

admins_info結構:
{
    "username" : "admin",
    "password" : "admin123",
    "company" : "mycompany"
}
companies_info結構:
{
    "name" : "My Awesome Company",
    "key" : "mycompany"
}
admins_info的company字段要與companies_info的key字段一致

admin_info需要添加最高級管理員賬號
{
    "username" : "bigadmin",
    "password" : "big123",
    "company" : "devops"
}

後端命令運行: 
go build 
nohup ./plagueOfCorona &


前端H5頁面mob
mob:
可修改App.vue中的BASE_URL為服務端IP

前端web頁面bknd
bknd:
可修改Auth.vue中的BASE_URL為服務端IP

vuejs框架需要安裝:
mob： OnsenUI+vue
bknd： bootstrap 
nodejs
npm install -g @vue/cli

前端打包發佈:
npm install
npm run build
將目錄下的dist包發到服務器上使用nginx運行