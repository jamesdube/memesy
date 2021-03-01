<template>
	<div>
		<!-- start -->
		<div class="bg-gray-100 h-screen p-16">
			<div class="container mx-auto lg:w-1/3 /mt-40 ">
				<div class="flex justify-center mb-16">
					<img class="h-10 w-auto sm:h-16" src="https://tailwindui.com/img/logos/workflow-mark-indigo-600.svg" alt="">
				</div>
				<div class="shadow border border-gray-200 bg-white rounded-md">
					<form @submit.prevent="createMeme">
						<div class="p-6 border-b border-gray-200 flex justify-center">
							<img v-bind:src="meme" class="h-64 w-64 rounded-md">
						</div>
						<div class="p-6 border-b border-gray-200 /bg-red-200">
							<div class="w-full md:w-1/3 /px-3 mb-6 md:mb-0 /bg-blue-200">
								<label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-city">
									Direction
								</label>
								<div class="relative">
									<select v-model="form.meme" :class="loadingClass" class="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded-md leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-state">
										<option v-for="meme in memes">{{ meme.Name }}</option>
									</select>
									<div class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
										<svg class="fill-current h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"/></svg>
									</div>
								</div>
							</div>
							<div class="flex flex-wrap mt-6 -mx-3 mb-2 /bg-pink-200">
								<div class="w-full md:w-1/2 px-3 mb-6 md:mb-0">
									<label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="grid-city">
										Top text
									</label>
									<input v-model="form.top" :class="loadingClass" class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded-md py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-city" type="text" placeholder="A goofy">
								</div>
								<div class="w-full md:w-1/2 px-3 mb-6 md:mb-0">
									<label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="bottom">
										Bottom Text
									</label>
									<input v-model="form.bottom" :class="loadingClass" class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded-md py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="bottom" type="text" placeholder="ball">
								</div>
							</div>
						</div>
						<div class="p-6 /mt-6 flex justify-end">
							<svg v-if="working" xmlns="http://www.w3.org/2000/svg" class="w-10 h-10 fill-current text-purple-700" viewBox="0 0 50 50"><path d="M43.935 25.145c0-10.318-8.364-18.683-18.683-18.683-10.318 0-18.683 8.365-18.683 18.683h4.068c0-8.071 6.543-14.615 14.615-14.615s14.615 6.543 14.615 14.615h4.068z"><animateTransform attributeType="xml" attributeName="transform" type="rotate" from="0 25 25" to="360 25 25" dur="0.6s" repeatCount="indefinite"/></path></svg>
							<button role="button" :class="loadingClass" class="bg-purple-700 rounded-md ml-4 px-8 py-2 text-white font-bold focus:border-red-400">
								Create
							</button>
						</div>
					</form>
				</div>
			</div>
		</div>
		<!-- end -->
	</div>
	
</template>

<script>

const BE = import.meta.env.VITE_MEMESY_BACKEND;

export default {

	name: 'App',
	data(){
		return {
			working: false,
			meme: "https://picsum.photos/200",
			memes : [],
			form: {
				meme: "",
				top: "",
				bottom: ""
			}
		}
	},
	created: function () {
	    // `this` points to the vm instance
	    this.getMemes()
	},
	methods:{
		getMemes(){
			fetch(BE)
			.then(response => response.json())
			.then(data => {
				console.log(data)
				this.memes = data.data
			}
			);
		},
		createMeme(){
			if(!this.working){
				this.working = true
			}
			const data = { Name: this.form.meme, TopText: this.form.top };
	  		//let that = this
	  		fetch('http://127.0.0.1:3000/memes/create',{
	  			method: "POST",
	  			headers: {
	  				'Content-Type': 'application/json',
	  			},
	  			body: JSON.stringify(data),
	  		})
	  		.then(response => response.json())
	  		.then(data => {
	  			this.working = false
	  			console.log(data.Name)
	  			this.meme = data.Image
	  		})
	  		.catch((error) => {
	  			this.working = false
	  			console.error('Error:', error);
	  		});
	  	}
  	},
  	computed: {
		loadingClass() {
			return {
				'opacity-50' : this.working, 
				'cursor-not-allowed': this.working
			}
		}
	}
};
</script>