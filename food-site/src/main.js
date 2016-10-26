import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
Vue.use(VueRouter);

const routes = [
	{ path: '/', component: App },
	{ path: '/mass', component: Vue.component({template: '<div>Hello</div>'}) }
]

const router = new VueRouter({
	routes
})

const app = new Vue({
	router
}).$mount("#app");
