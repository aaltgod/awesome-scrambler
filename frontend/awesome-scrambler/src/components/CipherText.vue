<template>
  <div class="cipher-text">
    <h2 class="cipher-text">Key<br></h2>
    <textarea class="cipher-text" v-model="key" placeholder="insert your key"></textarea>
    <button class="cipher-text" type="submit" v-on:click="getCipherText(key)">Get ciphertext<br></button>
    <h2 v-if="cipherText">Your ciphertext<br><br>
      <textarea class="cipher-text"  > {{ cipherText }} </textarea>
    </h2>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data: function() {
    return {
      cipherText: "",
      key: "",
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
        data: {"key": this.key},
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
<style>
.cipher-text {
  font-family: "Audiowide", sans-serif;
  -webkit-font-smoothing: antialiased;
}

textarea {
  width: 490px;
  height: 213px;
  font-family: 'Audiowide', sans-serif;
}
</style>