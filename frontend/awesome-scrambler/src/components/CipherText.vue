<template>
  <div class="cipher-text">
    <router-link to='/'><h2>Home</h2></router-link>
    <a v-if="cipherText">Your cipher text<br><br>{{ cipherText }}</a>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data: function() {
    return {
      cipherText: "",
    }
  },

  mounted: function() {
    this.getCipherText()
  },

  methods: {
    getCipherText: function() {
      axios({
        method: "post",
        url: "http://localhost:5000/api/get_cipher_text",
        data: {"link": this.$route.params.path},
        headers: {"content-type": "application/json"}
      }).then( result => {
        this.cipherText = result.data["cipher_text"]
      }).catch( error => {
        console.error(error)
      })
    }
  }
}
</script>
