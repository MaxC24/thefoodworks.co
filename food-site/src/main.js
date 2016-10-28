import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from './Home.vue'
Vue.use(VueRouter);

const routes = [
	{ path: '/', component: Home }
]

const router = new VueRouter({
	routes
})

const app = new Vue({
	router
}).$mount("#app");
