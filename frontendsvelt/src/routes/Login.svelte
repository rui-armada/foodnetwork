<script>
    import { onMount } from 'svelte';
    import axios from 'axios';
  
    let email = '';
    let password = '';
    let name = '';
    let isRegister = false;
  
    async function handleSubmit(event) {
      event.preventDefault();
      try {
        const route = isRegister ? 'http://localhost:8080/api/v1/register' : 'http://localhost:8080/api/v1/login';
        const response = await axios.post(route, { email, password, name });
  
        localStorage.setItem('token', response.data.token);
        window.location.href = '/dashboard';
      } catch (error) {
        alert('error');
      }
    }
  </script>
  
  <style>
    /* Add your styles here */
  </style>
  
  <div>
    <h2>{isRegister ? 'Register' : 'Login'}</h2>
    <form on:submit={handleSubmit}>
      {#if isRegister}
      <div>
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label>Name:</label>
        <input type="text" bind:value={name} placeholder="Enter your name" />
      </div>
      {/if}
      <div>
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label>Email:</label>
        <input type="email" bind:value={email} placeholder="Enter your email" />
      </div>
      <div>
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label>Password:</label>
        <input type="password" bind:value={password} placeholder="Enter your password" />
      </div>
      <button type="submit">{isRegister ? 'Register' : 'Login'}</button>
      <p>
        {isRegister ? 'Already have an account?' : "Don't have an account yet?"}
        <button type="button" on:click={() => (isRegister = !isRegister)}>{isRegister ? 'Login' : 'Register'}</button>
      </p>
    </form>
  </div>
  