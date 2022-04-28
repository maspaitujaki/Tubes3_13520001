<script>
export default {
  props: ['nomor', 'id', 'nama', 'tanggal', 'penyakit', 'prediksi', 'hasil'],
  methods: {
      parseDate(tanggal) {
        const rg_tanggal = /\d{4}\-(0?[1-9]|1[012])\-(0?[1-9]|[12][0-9]|3[01])*/g;
        return tanggal.match(rg_tanggal)[0];
      }
  }
}
</script>

<template>
    <div class="container-result">
        <div class="container-result-box">
            <p><span class="bold">{{nomor}}.</span> {{this.parseDate(tanggal)}} - {{nama}} - {{penyakit}} - {{prediksi}} - {{hasil}}</p>
        </div>
    </div>
</template>

<style scoped>
.container-result {
  position: relative;
  margin: auto;
  transform: translate(0%, 130%);
}

.container-result-box {
  margin: auto;
  width: 1000px;
  height: 90px;
  border-radius: 20px;
  background-color: #f2f2f2;
}

.container-result-box p{
  transform: translate(0%, 5%);
  position: relative;
  font-size: 24px;
  text-align: left;
  padding: 25px;
}

.bold {
  font-weight: bold;
  color: rgb(94, 31, 172);
}
</style>