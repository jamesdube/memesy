<template>
  <img alt="Vue logo" src="../assets/logo.png" />
  <h1>{{ msg }}</h1>
  <p>
    <select class="bg-gray-300 px-3">
      <option>Foo</option>
      <option>Bar</option>
    </select>
  </p>
  <p><input type="text" v-model="memeTop" name="top"></p>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  data() {
    return {
      count: 0,
      memeName: ""
    }
  }
}
</script>
