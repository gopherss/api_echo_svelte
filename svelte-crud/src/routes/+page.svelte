<script>
	import Time from 'svelte-time';
	let listUsers = new Array();
	let user = {
		ID: Number(),
		name: String(),
		age: Number(),
		created: String()
	};
	let statusUpdate = false;

	const cleanForm = () => {
		user.name = String();
		user.age = Number();
	};

	const getListUsers = async () => {
		let response = await fetch('http://localhost:8080');
		let data = await response.json();
		listUsers = data;
		cleanForm();
	};

	getListUsers();

	const saveNewUsers = async () => {
		await fetch('http://localhost:8080/save', {
			method: 'POST',
			body: JSON.stringify({
				name: user.name,
				age: Number(user.age)
			}),
			headers: {
				'Content-Type': 'application/json'
			}
		});
		getListUsers();
	};

	const deleteUsers = async (ID = 0) => {
		let status = confirm('are you sure to delete this user?');

		if (status) {
			await fetch('http://localhost:8080/delete/' + String(ID), {
				method: 'DELETE'
			});

			getListUsers();
		} else {
			alert('Ok Carnal :)');
		}
	};

	const updatedUsers = async (ID = 0) => {
		await fetch('http://localhost:8080/update/' + String(ID), {
			method: 'PUT',
			body: JSON.stringify({
				name: user.name,
				age: Number(user.age)
			}),
			headers: {
				'Content-Type': 'application/json'
			}
		});

		getListUsers();
		statusUpdate = false;
	};

	const editUser = (userEdited = user) => {
		user = userEdited;
		statusUpdate = true;
	};
</script>

<nav class="navbar navbar-expand-lg navbar-dark bg-dark">
	<div class="container-fluid">
		<a class="navbar-brand" href="/">Wuelcome</a>

		<div class="collapse navbar-collapse" id="navbarColor02">
			<ul class="navbar-nav me-auto">
				<li class="nav-item dropdown">
					<a
						class="nav-link dropdown-toggle"
						data-bs-toggle="dropdown"
						href="/"
						role="button"
						aria-haspopup="true"
						aria-expanded="false">Technologies</a
					>
					<div class="dropdown-menu">
						<a
							target="_blank"
							rel="noreferrer"
							class="dropdown-item"
							href="https://echo.labstack.com/">Echo Framework</a
						>
						<a target="_blank" rel="noreferrer" class="dropdown-item" href="https://svelte.dev/"
							>Svelte JS</a
						>
						<div class="dropdown-divider" />
						<a target="_blank" rel="noreferrer" class="dropdown-item" href="https://go.dev/"
							>Golang</a
						>
						<a
							target="_blank"
							rel="noreferrer"
							class="dropdown-item"
							href="https://developer.mozilla.org/es/docs/Web/JavaScript">JavaScript</a
						>
					</div>
				</li>
			</ul>
		</div>
	</div>
</nav>

<div class="container p-4">
	<div class="row">
		<div class="col-md-6">
			<form class="bg-dark p-2">
				<div class="mb-3 row">
					<label for="nameid" class="col-sm-2 col-form-label text-white">Name</label>
					<div class="col-sm-10">
						<input type="text" bind:value={user.name} class="form-control" id="nameid" />
					</div>
				</div>
				<div class="mb-3 row">
					<label for="nameid" class="col-sm-2 col-form-label text-white">Age</label>
					<div class="col-sm-10">
						<input type="text" bind:value={user.age} class="form-control text-center" id="nameid" />
					</div>
				</div>

				<div class="d-grid gap-2">
					{#if !statusUpdate}
						<button type="button" class="btn btn-outline-info btn-lg" on:click={saveNewUsers}>
							save
						</button>
					{:else}
						<button
							type="button"
							class="btn btn-outline-info btn-lg"
							on:click={() => updatedUsers(user.ID)}
						>
							Edit
						</button>
					{/if}
				</div>
			</form>
		</div>
		<div class="col-md-6">
			{#if listUsers.length === 0}
				<div class="alert alert-dark" role="alert">you must fill out the form!</div>
			{:else if listUsers.length > 0}
				{#each listUsers as user}
					<div class="card mt-2">
						<div class="row">
							<div class="col-md-4">
								<img src="imgs/not-found.png" class="img-fluid " alt="" />
							</div>
							<div class="col-md-8">
								<div class="card-body">
									<h5 class="card-title">
										<strong>Name: {user.name}</strong>
										<br />
										<span>Age: {user.age}</span>
									</h5>
									<p class="card-text">
										Time: <Time timestamp={user.created} format="D/MM/YYYY - hh:mm" />
									</p>
									<button class="btn btn-outline-warning " on:click={() => deleteUsers(user.ID)}>
										Delete
									</button>
									<button class="btn btn-outline-success" on:click={() => editUser(user)}>
										Update
									</button>
								</div>
							</div>
						</div>
					</div>
				{/each}
			{/if}
		</div>
	</div>
</div>
