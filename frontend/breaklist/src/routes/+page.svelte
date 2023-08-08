<script lang="ts">
	import { onMount } from 'svelte';

	const API_URL = 'http://localhost:3000/';
	let darkMode = true;
	let tasks = [''];
	let state = 0; // 0 -> loading / 1-> done loading

	async function updateTasks() {
		state = 0;
		let response = await fetch(API_URL + 'task');
		let result = await response.json();

		tasks = result.data;
		state = 1;
	}
	onMount(updateTasks);

	function toggleDarkMode() {
		darkMode = !darkMode;

		if (darkMode) window.document.body.classList.add('dark-mode');
		else window.document.body.classList.remove('dark-mode');
	}
</script>

<button id="darkModeBtn" on:click={toggleDarkMode}>
	{#if darkMode}
		<span class="material-icons" style="color:beige;">light_mode</span>
	{:else}
		<span class="material-icons" style="color:black;">dark_mode</span>
	{/if}
</button>

<div class="container">
	<div class="c1">
		<h1>Tasks:</h1>
		{#if state === 1}
			<ol>
				{#each tasks as t}
					<li>{t}</li>
				{/each}
			</ol>
		{:else if state === 0}
			<div style="text-align: center">
				<div class="lds-ripple">
					<div />
					<div />
				</div>
			</div>
		{/if}
	</div>
</div>

<style>
	@import url('https://fonts.googleapis.com/css2?family=Victor+Mono:wght@300;400&display=swap');

	.c1 {
		/* margin-left: 10%; */
		text-align: left;
		font-family: 'Victor Mono', monospace;
	}

	:global(body) {
		background-color: beige;
		color: hsl(215, 68%, 7%);
		transition: background-color 0.3s;
	}
	:global(body.dark-mode) {
		background-color: hsl(215, 68%, 7%);
		color: hsl(0, 0%, 98%);
	}

	.container {
		margin: 5px;
		position: relative;
		height: 100%;
		text-align: center;
	}

	#darkModeBtn {
		overflow: hidden;
		border: none;
		text-align: center;
		text-decoration: none;
		background-color: rgba(0, 0, 0, 0);
		padding: 0;

		position: absolute;
		right: 15px;
		top: 15px;
	}

	@media only screen and (min-width: 768px) {
		#darkModeBtn {
			right: 25px;
			top: 10px;
		}
		.container {
			margin: auto;
			width: 768px;
		}
		:global(body) {
			height: 100%;
			margin: 1%;
			overflow: hidden;
		}
	}

	/* loader */
	.lds-ripple {
		display: inline-block;
		position: relative;
		margin: auto;
		margin-top: 200px;
		width: 80px;
		height: 80px;
	}
	.lds-ripple div {
		position: absolute;
		border: 4px solid rgb(27, 27, 27);
		opacity: 1;
		border-radius: 50%;
		animation: lds-ripple 1s cubic-bezier(0, 0.2, 0.8, 1) infinite;
	}
	.lds-ripple div:nth-child(2) {
		animation-delay: -0.5s;
	}
	@keyframes lds-ripple {
		0% {
			top: 36px;
			left: 36px;
			width: 0;
			height: 0;
			opacity: 0;
		}
		4.9% {
			top: 36px;
			left: 36px;
			width: 0;
			height: 0;
			opacity: 0;
		}
		5% {
			top: 36px;
			left: 36px;
			width: 0;
			height: 0;
			opacity: 1;
		}
		100% {
			top: 0px;
			left: 0px;
			width: 72px;
			height: 72px;
			opacity: 0;
		}
	}

	body.dark-mode .lds-ripple div {
		border: 4px solid rgb(231, 231, 231);
	}
	/* loader */
</style>
