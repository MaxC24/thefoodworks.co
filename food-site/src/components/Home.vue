<template>
	<div id="container">
	<div>
		<div id="hero-img">
			<!-- <button class="btn">Blog</button>
			<button class="btn">Contact</button>
 -->		</div>
		<div class="par">
			<p>
				One of the greatest problems of our modern times is: <strong>The Paradox Of Choices</strong>.
				If you have ever ordered delivery food you know how difficult it is to choose from the multitude of options available. Restaurants menus are way too long and it's very hard to know how the food is going to look like or whether is going to be any good.
				The advent of Grubhub and Seamless made the food delivery option faster and more accessible, but, because of the added choices and long menus to scroll through, the research for the optimal dish became harder and time consuming.
			</p>
			<div class="img"><img src="http://tamaramccleary.com/wp-content/uploads/2015/01/choices.jpg" /></div>
		</div>
		<div class="par">
			<div class="img" id="mobimg-one" ><img src="http://www.siriuscom.com/wp-content/uploads/2014/11/Solutions.jpg" /></div>
			<p>
				With our product we solve the problem of investing too much time looking for the best option by making the experience more engaging , interactive and especially more social. The user will be able to explore more choices at a faster pace, and at the same time reducing "the stress" and "fear of the unknown" related to choosing a dish never ordered before.
				We cannot completely eliminate <strong>The Paradox Of Choices</strong> but at least we can make it less disrupting in our everyday life.
			<!-- 	Shortly we are planning to showcase a demo of our application, so t
				We are doing this by showing you a stream of pictures of dishes that people have ordered nearby. You will be able to click on the image and see the details of the dish and if you like it, you will be able to order it.  We believe that by showing you the dish, you will not have to look through endless menus of restaurants.  Furthermore, with a service like Yelp, the resturant gets a good grade if a single dish there was good.  But with this service, the dish or a particular resturant gets a rating.  This will make sure the resturant is not over glorified for a particular dish. -->
			</p>
			<div class="img" id="mobimg-two"><img src="http://www.siriuscom.com/wp-content/uploads/2014/11/Solutions.jpg" /></div>
		</div>
	</div>
	<foot />
	</div>
</template>

<script>
	import foot from './Footer.vue'
	export default {
		// name: 'home'g
		components: {
			foot
		}
	}
</script>


<style>
	*{
	  box-sizing: border-box;
	}
	#container {
		margin-bottom: 0;
	}
	#app {
	  font-family: 'Avenir', Helvetica, Arial, sans-serif;
	  -webkit-font-smoothing: antialiased;
	  -moz-osx-font-smoothing: grayscale;
	  text-align: center;
	  color: #2c3e50;
	}
	body {
		margin: 0%;
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
		padding-bottom: 3vw;
	}

	#mobimg-two {
		display: none;
	}

	#mobimg-one img{
		width: 20vw;
		padding-bottom: 0vw;
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
		#hero-img {
			display: inline-block;
			width: 90vw;
			height: 60vw;
		}
		.par p {
			font-size: 1.2vw;
			width: 40vw;
		}
		.par .img {
			width: 40vw;
		}
		.par .img img {
			width: 30vw;
			padding-bottom: 0vw;
		}
		#mobimg-one img {
			width: 20vw;
		}
	}
</style>