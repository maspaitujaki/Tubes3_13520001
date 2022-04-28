<script>
export default {
  props: ['tanggal', 'nama', 'penyakit', 'prediksi', 'hasil']
}
</script>

<template>
    <div class="container-test">
        <div class="container-test-box">
            <p>{{tanggal}} - {{nama}} - {{penyakit}} - {{prediksi}} - <span class="bold">{{hasil}}</span></p>
        </div>
    </div>
</template>

<style scoped>
.container-test {
  position: relative;
  margin: auto;
  transform: translate(0%, 130%);
}

.container-test-box {
  margin: auto;
  width: 1000px;
  height: 90px;
  border-radius: 20px;
  background-color: #f2f2f2;
}

.container-test-box p{
  transform: translate(0%, 5%);
  position: relative;
  font-size: 24px;
  text-align: center;
  padding: 25px;
}

.bold {
  font-weight: bold;
  color: rgb(94, 31, 172);
}
</style>