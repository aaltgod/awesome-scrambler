<template>
  <div class="app">
    <h1>Encrypt your text</h1>
    <p style="white-space: pre-line;"></p>
    <textarea class="app" v-model="plainText" placeholder="insert text"></textarea>
    <button class="app" type="submit" v-on:click="encryptText(plainText)">Encrypt</button>
    <h3 class="app" v-if="key">Cipher key: {{ key }}</h3>
    <h3 class="app" v-if="path">Your link with the cipher text: <a v-bind:href="'/ciphertext/'+path">{{ path }}</a><br></h3>
    <h4>You can send a message to ggfgde8@gmail.com with the subject: Encrypt</h4>
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
      if (this.plainText.length === 0) {
        return
      }

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
