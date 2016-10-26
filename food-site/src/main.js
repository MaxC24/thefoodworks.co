import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from './Home.vue'
import App from './App.vue'
Vue.use(VueRouter);

const routes = [
	{ path: '/', component: App },
	{ path: '/fw', component: Home }
]

const router = new VueRouter({
	routes
})

const app = new Vue({
	router
}).$mount("#app");
