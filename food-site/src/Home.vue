<template>
	<div id="hero-img">
		<button class="btn">Blog</button>
		<button class="btn">Contact</button>
	</div>
</template>

<script>
	export default {
		// name: 'home'
	}
</script>


<style>
	#hero-img {
		width: 100vw;
		height: 65vw;
		background-image: url('./assets/hero.jpg');
		background-size: 100%;
	}
	.btn {
		float: right;
		margin: 5px;
	}
</style>