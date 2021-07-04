<template>
  <div class="app">
    <h1>Encrypt your text</h1>
    <p style="white-space: pre-line;"></p>
    <textarea class="app" v-model="plainText" placeholder="insert text"></textarea>
    <button class="app" type="submit" v-on:click="encryptText(plainText)">Encrypt</button>
    <h3 class="app" v-if="key">Cipher key: {{ key }}</h3>
    <h3 class="app" v-if="path"><a v-bind:href="'/ciphertext/'+path">{{ path }}</a></h3>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data: function() {
    return {
      plainText: "",
      key: "",
      path: ""
    }
  },

  methods: {
    encryptText: function(plainText) {
      console.log(plainText)
      axios({
        method: "post",
        url: "http://localhost:5000/api/encrypt_text",
        data: {"text": plainText},
        headers: {"content-type": "application/json"}
      }).then(result => {
        this.key = result.data["key"]
        this.path = result.data["link"]
      }).catch( error => {
        console.error(error)
      })
    }
  }
}
</script>

<style scoped>
.app {
  font-family: "Audiowide", sans-serif;
  -webkit-font-smoothing: antialiased;
}
h3 {
  margin: 40px 0 0;
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
