webpackJsonp([18],{"0re4":function(t,e,n){var a=n("D60T");"string"==typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);n("MO7y")("55c8c38f",a,!0)},D60T:function(t,e,n){(t.exports=n("xCkK")(!1)).push([t.i,"\n.mixin-components-container[data-v-7ed4e1bc] {\n  background-color: #f0f2f5;\n  padding: 30px;\n  min-height: calc(100vh - 84px);\n}\n.component-item[data-v-7ed4e1bc]{\n  min-height: 100px;\n}\n",""])},rCis:function(t,e,n){var a=n("rYMm");"string"==typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);n("MO7y")("70ea56d2",a,!0)},rYMm:function(t,e,n){(t.exports=n("xCkK")(!1)).push([t.i,"\n.share-dropdown-menu {\n  width: 250px;\n  position: relative;\n  z-index: 1;\n}\n.share-dropdown-menu-title {\n    width: 100%;\n    display: block;\n    cursor: pointer;\n    background: black;\n    color: white;\n    height: 60px;\n    line-height: 60px;\n    font-size: 20px;\n    text-align: center;\n    z-index: 2;\n    -webkit-transform: translate3d(0, 0, 0);\n            transform: translate3d(0, 0, 0);\n}\n.share-dropdown-menu-wrapper {\n    position: relative;\n}\n.share-dropdown-menu-item {\n    text-align: center;\n    position: absolute;\n    width: 100%;\n    background: #e0e0e0;\n    line-height: 60px;\n    height: 60px;\n    cursor: pointer;\n    font-size: 20px;\n    opacity: 1;\n    -webkit-transition: -webkit-transform 0.28s ease;\n    transition: -webkit-transform 0.28s ease;\n    transition: transform 0.28s ease;\n    transition: transform 0.28s ease, -webkit-transform 0.28s ease;\n}\n.share-dropdown-menu-item:hover {\n      background: black;\n      color: white;\n}\n.share-dropdown-menu-item:nth-of-type(1) {\n      z-index: -1;\n      -webkit-transition-delay: 0.1s;\n              transition-delay: 0.1s;\n      -webkit-transform: translate3d(0, -60px, 0);\n              transform: translate3d(0, -60px, 0);\n}\n.share-dropdown-menu-item:nth-of-type(2) {\n      z-index: -1;\n      -webkit-transition-delay: 0.2s;\n              transition-delay: 0.2s;\n      -webkit-transform: translate3d(0, -60px, 0);\n              transform: translate3d(0, -60px, 0);\n}\n.share-dropdown-menu-item:nth-of-type(3) {\n      z-index: -1;\n      -webkit-transition-delay: 0.3s;\n              transition-delay: 0.3s;\n      -webkit-transform: translate3d(0, -60px, 0);\n              transform: translate3d(0, -60px, 0);\n}\n.share-dropdown-menu-item:nth-of-type(4) {\n      z-index: -1;\n      -webkit-transition-delay: 0.4s;\n              transition-delay: 0.4s;\n      -webkit-transform: translate3d(0, -60px, 0);\n              transform: translate3d(0, -60px, 0);\n}\n.share-dropdown-menu-item:nth-of-type(5) {\n      z-index: -1;\n      -webkit-transition-delay: 0.5s;\n              transition-delay: 0.5s;\n      -webkit-transform: translate3d(0, -60px, 0);\n              transform: translate3d(0, -60px, 0);\n}\n.share-dropdown-menu-item:nth-of-type(6) {\n      z-index: -1;\n      -webkit-transition-delay: 0.6s;\n              transition-delay: 0.6s;\n      -webkit-transform: translate3d(0, -60px, 0);\n              transform: translate3d(0, -60px, 0);\n}\n.share-dropdown-menu.active .share-dropdown-menu-wrapper {\n    z-index: 1;\n}\n.share-dropdown-menu.active .share-dropdown-menu-item:nth-of-type(1) {\n    -webkit-transition-delay: 0.5s;\n            transition-delay: 0.5s;\n    -webkit-transform: translate3d(0, 0px, 0);\n            transform: translate3d(0, 0px, 0);\n}\n.share-dropdown-menu.active .share-dropdown-menu-item:nth-of-type(2) {\n    -webkit-transition-delay: 0.4s;\n            transition-delay: 0.4s;\n    -webkit-transform: translate3d(0, 60px, 0);\n            transform: translate3d(0, 60px, 0);\n}\n.share-dropdown-menu.active .share-dropdown-menu-item:nth-of-type(3) {\n    -webkit-transition-delay: 0.3s;\n            transition-delay: 0.3s;\n    -webkit-transform: translate3d(0, 120px, 0);\n            transform: translate3d(0, 120px, 0);\n}\n.share-dropdown-menu.active .share-dropdown-menu-item:nth-of-type(4) {\n    -webkit-transition-delay: 0.2s;\n            transition-delay: 0.2s;\n    -webkit-transform: translate3d(0, 180px, 0);\n            transform: translate3d(0, 180px, 0);\n}\n.share-dropdown-menu.active .share-dropdown-menu-item:nth-of-type(5) {\n    -webkit-transition-delay: 0.1s;\n            transition-delay: 0.1s;\n    -webkit-transform: translate3d(0, 240px, 0);\n            transform: translate3d(0, 240px, 0);\n}\n.share-dropdown-menu.active .share-dropdown-menu-item:nth-of-type(6) {\n    -webkit-transition-delay: 0s;\n            transition-delay: 0s;\n    -webkit-transform: translate3d(0, 300px, 0);\n            transform: translate3d(0, 300px, 0);\n}\n",""])},ypWw:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var a=n("kCe2"),s=n("+mJe"),r=n("Weyc"),i={props:{items:{type:Array},title:{type:String,default:"vue"}},data:function(){return{isActive:!1}},methods:{clickTitle:function(){this.isActive=!this.isActive}}},o={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"share-dropdown-menu",class:{active:t.isActive}},[n("div",{staticClass:"share-dropdown-menu-wrapper"},[n("span",{staticClass:"share-dropdown-menu-title",on:{click:function(e){if(e.target!==e.currentTarget)return null;t.clickTitle(e)}}},[t._v(t._s(t.title))]),t._v(" "),t._l(t.items,function(e,a){return n("div",{key:a,staticClass:"share-dropdown-menu-item"},[e.href?n("a",{attrs:{href:e.href,target:"_blank"}},[t._v(t._s(e.title))]):n("span",[t._v(t._s(e.title))])])})],2)])},staticRenderFns:[]};var l=n("kgPM")(i,o,!1,function(t){n("rCis")},null,null).exports,d=n("cAgV"),c={name:"componentMixin-demo",components:{PanThumb:a.a,MdInput:s.a,Mallki:r.a,DropdownMenu:l},directives:{waves:d.a},data:function(){return{demo:{title:""},demoRules:{title:[{required:!0,trigger:"change",validator:function(t,e,n){6!==e.length?n(new Error("请输入六个字符")):n()}}]},articleList:[{title:"基础篇",href:"https://segmentfault.com/a/1190000009275424"},{title:"登录权限篇",href:"https://segmentfault.com/a/1190000009506097"},{title:"实战篇",href:"https://segmentfault.com/a/1190000009762198"},{title:"vueAdmin-template 篇",href:"https://segmentfault.com/a/1190000010043013"},{title:"自行封装 component",href:"https://segmentfault.com/a/1190000009090836"},{title:"优雅的使用 icon",href:"https://segmentfault.com/a/https://segmentfault.com/a/1190000012213278"}]}}},p={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"mixin-components-container"},[n("el-row",[n("el-card",{staticClass:"box-card"},[n("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[n("span",[t._v("Buttons")])]),t._v(" "),n("div",{staticStyle:{"margin-bottom":"50px"}},[n("el-col",{staticClass:"text-center",attrs:{span:4}},[n("router-link",{staticClass:"pan-btn blue-btn",attrs:{to:"/components/index"}},[t._v("Components")])],1),t._v(" "),n("el-col",{staticClass:"text-center",attrs:{span:4}},[n("router-link",{staticClass:"pan-btn light-blue-btn",attrs:{to:"/charts/index"}},[t._v("Charts")])],1),t._v(" "),n("el-col",{staticClass:"text-center",attrs:{span:4}},[n("router-link",{staticClass:"pan-btn pink-btn",attrs:{to:"/excel/download"}},[t._v("Excel")])],1),t._v(" "),n("el-col",{staticClass:"text-center",attrs:{span:4}},[n("router-link",{staticClass:"pan-btn green-btn",attrs:{to:"/example/table/complex-table"}},[t._v("Table")])],1),t._v(" "),n("el-col",{staticClass:"text-center",attrs:{span:4}},[n("router-link",{staticClass:"pan-btn tiffany-btn",attrs:{to:"/form/edit-form"}},[t._v("Form")])],1),t._v(" "),n("el-col",{staticClass:"text-center",attrs:{span:4}},[n("router-link",{staticClass:"pan-btn yellow-btn",attrs:{to:"/theme/index"}},[t._v("Theme")])],1)],1)])],1),t._v(" "),n("el-row",{staticStyle:{"margin-top":"50px"},attrs:{gutter:20}},[n("el-col",{attrs:{span:6}},[n("el-card",{staticClass:"box-card"},[n("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[n("span",[t._v("Material Design 的input")])]),t._v(" "),n("div",{staticStyle:{height:"100px"}},[n("el-form",{attrs:{model:t.demo,rules:t.demoRules}},[n("el-form-item",{attrs:{prop:"title"}},[n("md-input",{attrs:{icon:"search",name:"title",placeholder:"输入标题"},model:{value:t.demo.title,callback:function(e){t.$set(t.demo,"title",e)},expression:"demo.title"}},[t._v("标题")])],1)],1)],1)])],1),t._v(" "),n("el-col",{attrs:{span:6}},[n("el-card",{staticClass:"box-card"},[n("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[n("span",[t._v("图片hover效果")])]),t._v(" "),n("div",{staticClass:"component-item"},[n("pan-thumb",{attrs:{width:"100px",height:"100px",image:"https://wpimg.wallstcn.com/577965b9-bb9e-4e02-9f0c-095b41417191"}},[t._v("\n            vue-element-admin\n          ")])],1)])],1),t._v(" "),n("el-col",{attrs:{span:6}},[n("el-card",{staticClass:"box-card"},[n("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[n("span",[t._v("水波纹 waves v-directive")])]),t._v(" "),n("div",{staticClass:"component-item"},[n("el-button",{directives:[{name:"waves",rawName:"v-waves"}],attrs:{type:"primary"}},[t._v("水波纹效果")])],1)])],1),t._v(" "),n("el-col",{attrs:{span:6}},[n("el-card",{staticClass:"box-card"},[n("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[n("span",[t._v("hover text")])]),t._v(" "),n("div",{staticClass:"component-item"},[n("mallki",{attrs:{className:"mallki-text",text:"vue-element-admin"}})],1)])],1)],1),t._v(" "),n("el-row",{staticStyle:{"margin-top":"50px"},attrs:{gutter:20}},[n("el-col",{attrs:{span:6}},[n("el-card",{staticClass:"box-card"},[n("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[n("span",[t._v("Share")])]),t._v(" "),n("div",{staticClass:"component-item",staticStyle:{height:"420px"}},[n("dropdown-menu",{staticStyle:{margin:"0 auto"},attrs:{title:"系列文章",items:t.articleList}})],1)])],1)],1)],1)},staticRenderFns:[]};var m=n("kgPM")(c,p,!1,function(t){n("0re4")},"data-v-7ed4e1bc",null);e.default=m.exports}});