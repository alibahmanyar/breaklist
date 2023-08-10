<script lang="ts">
	import { onMount } from 'svelte';
    import { fade, blur, slide } from "svelte/transition";

	const API_URL = 'http://localhost:3000/';
	let darkMode = true;
	let tasks = [''];
	let state = 0; // 0 -> loading / 1-> done loading
	let newTask = '', taskToDelete = '';
	let addTaskPopUp = false, delTaskPopUp = false;

	async function updateTasks() {
		state = 0;
		let response = await fetch(API_URL + 'task');
		let result = await response.json();

		tasks = result.data;
		state = 1;
	}
	onMount(updateTasks);
	onMount(() => {
		darkMode = JSON.parse(document.cookie).darkMode;
		if (darkMode) window.document.body.classList.remove('light-mode');
		else window.document.body.classList.add('light-mode');
	});

	function init(el: any){
    	el.focus()
  	}

	function toggleDarkMode() {
		darkMode = !darkMode;

		document.cookie = JSON.stringify({ darkMode: darkMode });

		if (darkMode) window.document.body.classList.remove('light-mode');
		else window.document.body.classList.add('light-mode');
	}

	async function delTask(task: string) {
		console.log(task);

		let response = await fetch(API_URL + 'task', {
			method: 'DELETE',
			body: JSON.stringify({ data: [task] })
		});
		let result = await response.json();

		updateTasks();
	}

	async function addTask() {
		console.log(newTask);
		addTaskPopUp = false;

		let response = await fetch(API_URL + 'task', {
			method: 'POST',
			body: JSON.stringify({ data: [newTask] })
		});
		let result = await response.json();

		updateTasks();

		newTask = '';
	}
</script>

{#if addTaskPopUp}
	<div class="popup" transition:blur>
		<form on:submit={addTask}>
			<div class="vbox font1" id="pp0">
				<div>Add New task:</div>
				<input type="text" style="width: 80%; height: 2vh;" bind:value={newTask} use:init/>


				<div style="display: flex; flex-direction: row;">
					<button class="pp_btn" on:click={() => {addTaskPopUp = false; newTask = '';}} type="reset">Back</button>
					<span style="padding-left: 10px; padding-right: 10px;"></span>
					<button class="pp_btn" type="submit">Add task</button>
				</div>
				
			</div>
		</form>
	</div>
{/if}

{#if delTaskPopUp}
	<div class="popup" transition:blur>
		<form on:submit={() => {delTask(taskToDelete); delTaskPopUp=false;}}>
			<div class="vbox font1" id="pp0">
				<div>Delete {taskToDelete}?</div>
				
				<div style="display: flex; flex-direction: row;">
					<button class="pp_btn" on:click={() => {delTaskPopUp = false;}} type="reset">Back</button>
					<span style="padding-left: 10px; padding-right: 10px;"></span>
					<button class="pp_btn" type="submit">Delete task</button>
				</div>
				
			</div>
		</form>
	</div>
{/if}


<button class="sbtn" id="darkModeBtn" on:click={toggleDarkMode}>
	{#if darkMode}
		<span class="material-icons" style="color:beige;">light_mode</span>
	{:else}
		<span class="material-icons" style="color:black;">dark_mode</span>
	{/if}
</button>

<div class="container font1">
	<h1 id="hd1">
		<div class="hbox0">
			Tasks:
			<button
				class="sbtn"
				on:click={() => {
					addTaskPopUp = true;

				}}
			>
				<span class="material-icons" style="font-size: 2.5rem; padding-top:5px">add</span>
			</button>
		</div>
	</h1>

	<div class="vbox">
		{#if state === 1}
			{#if tasks !== null}
				{#each tasks as t}
					<div class="task hbox0">
						{t}
						<button
							class="sbtn"
							on:click={() => {
								// delTask(t);
								taskToDelete = t;
								delTaskPopUp = true;
							}}
						>
							<span class="material-icons" style="">close</span>
						</button>
					</div>
				{/each}
			{:else}
				No tasks left!
			{/if}
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

	.font1 {
		font-family: 'Victor Mono', monospace;
	}

	.vbox {
		display: flex;
		flex-direction: column;
	}

	.hbox0 {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
	}

	.task {
		padding: 0 5px 0 10px;
		flex-grow: 1;
		height: 50px;
		margin-top: 2px;
		margin-bottom: 2px;
	}

	.task:nth-child(even) {
		transition: background-color 0.3s;
		background-color: hsl(215, 58%, 25%);
	}
	.task:nth-child(odd) {
		transition: background-color 0.3s;
		background-color: hsl(215, 58%, 15%);
	}
	:global(body.light-mode) .task:nth-child(even) {
		background-color: hsl(55, 50%, 85%);
	}
	:global(body.light-mode) .task:nth-child(odd) {
		background-color: hsl(55, 40%, 80%);
	}

	.popup {
		position: fixed;
		top: 0;
		left: 0;
		height: 100%;
		width: 100%;
		background-color: rgba(0, 0, 0, 0.5);
		z-index: 1;
	}
	#pp0 {
		align-self: center;
		margin: auto;
		margin-top: 50%;
		height: fit-content;
		width: 85%;
		padding: 20px;
		align-items: center;

		gap: 10px;

		background-color: rgba(53, 53, 53, 0.5);
	}

	:global(body.light-mode) #pp0 {
		background-color: rgba(255, 255, 255, 0.5);
	}

	

	:global(html, body) {
		height: 100%;
	}

	:global(body) {
		background-color: hsl(215, 68%, 7%);
		color: hsl(0, 0%, 98%);
		transition: background-color 0.3s;
	}
	:global(body.light-mode) {
		background-color: beige;
		color: hsl(215, 68%, 7%);
	}

	.container {
		margin: 5px;
		position: relative;
		height: 100%;
		text-align: left;
	}

	.sbtn {
		overflow: hidden;
		border: none;
		text-align: center;
		text-decoration: none;
		background-color: rgba(0, 0, 0, 0);
		padding: 0;
		color: beige;
	}

	:global(body.light-mode) .sbtn {
		color: hsl(215, 68%, 7%);
	}

	.sbtn:active {
		color: gray;
	}
	:global(body.light-mode) .sbtn:active {
		color: gray;
	}

	.pp_btn {
		margin-top: 20px;
		font-size: 1.2rem;
		border-width: 1px;
		background-color: rgba(0, 0, 0, 0);
		color: hsl(0, 0%, 98%);
		border: 2px solid;
		padding: 5px 40px 5px 40px;
	}
	.pp_btn:active {
		color: gray;
	}

	:global(body.light-mode) .pp_btn {
		color: hsl(215, 68%, 7%);
	}
	:global(body.light-mode) .pp_btn:active {
		color: gray;
	}

	#darkModeBtn {
		position: absolute;
		right: 15px;
		top: 15px;
	}

	#hd1 {
		margin-top: 15%;
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

		#hd1 {
			margin-top: 8%;
		}

		#pp0 {
			margin-top: 20%;
			width: 50%;
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
		border: 4px solid rgb(231, 231, 231);
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

	body.light-mode .lds-ripple div {
		border: 4px solid rgb(27, 27, 27);
	}
	/* loader */
</style>
