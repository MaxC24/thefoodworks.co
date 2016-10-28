<template>
	<div>
		<div id="hero-img">
			<button class="btn">Blog</button>
			<button class="btn">Contact</button>
		</div>
		<div class="par">
			<p>
				Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
			</p>
			<div class="img"><img src="./assets/hero.jpg" /></div>
		</div>
		<div class="par">
			<div class="img"><img src="./assets/hero.jpg" /></div>
			<p>
				Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
			</p>
		</div>
	</div>
</template>

<script>
	export default {
		// name: 'home'g
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
		background-image: url('./assets/hero.jpg');
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
		padding:5%;
		display: inline-block;
		font-size: 2vw;
	}
	.par .img {
		width:	40%;
		display: inline-block;
		background-size: 100%;
	}
	.par .img img {
		width: 30vw;
		padding-bottom: 10vw;
	}
</style>