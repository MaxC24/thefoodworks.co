<template>
	<div id="footer">
	<a href="mailto:massimo.crapanzano@gmail.com"><button class="btn btn-ft">contact</button></a>
	</div>
</template>

<script>
	export default {
		// name: footer
	}
</script>

<style>
	.btn-ft {
		margin:6vw 2vw 0vw 0vw;
	}
	#footer {
		height: 10vw;
		background-image: url('../assets/hero.jpg');
		background-position: 0px 15vw;
		background-size: 100%;
		margin-top: 3vw;
	}
	@media (min-width: 1800px) {
		#footer {
			display: inline-block;
			width: 90vw;
			vertical-align: top;
		} 
	}

	@media (max-width: 720px) {
		.btn-ft {
			font-size: 3vw;
			width: 30vw;
			height: 7vw;
			margin:1.5vw 2vw 0vw 0vw;
		}
	}
</style>