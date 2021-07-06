<template>
  <div id="app">
    <ul>
      <li><router-link to='/'><h1>Encrypt text</h1></router-link></li>
      <li><router-link to='/ciphertext'><h1>Get ciphertext</h1></router-link></li>
    </ul>
    <router-view></router-view>
  </div>
</template>

<script>
export default {}
</script>

<style>
#app {
  font-family: "Audiowide", sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #eec49a;
  margin-top: 60px;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}

</style>
