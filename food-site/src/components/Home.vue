<template>
	<div>
	<div>
		<div id="hero-img">
			<!-- <button class="btn">Blog</button>
			<button class="btn">Contact</button>
 -->		</div>
		<div class="par">
			<p>
				We live in a world of too many choices.
				Making a decision about the food you would like to order
				becomes a difficult decision with all the choices avaiable.  
				The menus for resturants are all very long with too many decisions to make.  You might have some doubts, you don't know what the food is going to look like. The advent of Grubhub and Seamless make getting to these resturant menus easier online, but the added choices make finding the perfect dish even harder.  
				After not being able to make a decision after a while, you start to eat your "goto" dish for your meal.  However, there could be many great dishes around you that are being ordered.  
			</p>
			<div class="img"><img src="http://tamaramccleary.com/wp-content/uploads/2015/01/choices.jpg" /></div>
		</div>
		<div class="par">
			<div class="img" id="mobimg-one" ><img src="../assets/hero.jpg" /></div>
			<p>
				We are solving the problem of find the right dish for your meal.  
				We are doing this by showing you a stream of pictures of dishes that people have ordered nearby. You will be able to click on the image and see the details of the dish and if you like it, you will be able to order it.  We believe that by showing you the dish, you will not have to look through endless menus of resturants.  Furthermore, with a service like Yelp, the resturant gets a good grade if a single dish there was good.  But with this service, the dish or a particular resturant gets a rating.  This will make sure the resturant is not over glorified for a particular dish.
			</p>
			<div class="img" id="mobimg-two"><img src="../assets/hero.jpg" /></div>
		</div>
	</div>
	<Footer />
	</div>
</template>

<script>
	import Footer from './Footer.vue'
	export default {
		// name: 'home'g
		components: {
			Footer
		}
	}
</script>


<style>
	*{
	  box-sizing: border-box;
	}
	#app {
	  font-family: 'Avenir', Helvetica, Arial, sans-serif;
	  -webkit-font-smoothing: antialiased;
	  -moz-osx-font-smoothing: grayscale;
	  text-align: center;
	  color: #2c3e50;
	}
	body {
		margin: 0;
	}
	#hero-img {
		width: 100vw;
		height: 65vw;
		background-image: url('../assets/hero.jpg');
		background-size: 100%;
		margin-bottom: 2vw;
	}
	button:focus {
		outline:none;
	}
	.btn {
		color: white;
		background-color: rgba(0, 0, 0, 0);
		float: right;
		margin: 10px;
		border: 3px solid white;
		border-radius: 5px;
		width: 10vw;
		height: 3vw;
		font-size: 1.5vw
	}
	.btn:hover {
		color:black;
		background-color: white;
		border: none;
	}
	.par {
		width: 100vw;
		margin: 0%;
	}
	.par p {
		width: 50vw;
		padding:2% 5%;
		display: inline-block;
		font-size: 1.5vw;
	}
	.par .img {
		width:	40%;
		display: inline-block;
		background-size: 100%;
	}
	.par .img img {
		width: 30vw;
		padding-bottom: 5vw;
	}

	#mobimg-two {
		display: none;
	}
	
	@media (max-width: 620px) {
		.par p {
			font-size: 3vw;
			width: 100vw;
		}
		.par .img {
			width: 100vw;
		}
		.par .img img{
			width: 60vw;
		}
		#mobimg-two {
			display: block;
		}
		#mobimg-one {
			display: none;
		}
	}

	@media (min-width: 1500px) {
		.par p {
			font-size: 1.2vw;
			width: 40vw;
		}
		.par .img {
			width: 40vw;
		}
		.par .img img {
			width: 30vw;
			padding-bottom: 3vw;
		}
	}
</style>