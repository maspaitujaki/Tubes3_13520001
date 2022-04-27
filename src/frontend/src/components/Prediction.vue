<template>
  <div class="background">
    <div class="top-cont">
      <div class="canvas">
        <div class="container-header">
          <div class="container-logo">
            <img src="../assets/logo-full.svg" alt="" class="logo-full" />
          </div>
          <ul class="navigation-bar">
            <li class="inactive">
              <router-link to="/">Tambah Penyakit</router-link>
            </li>
            <li class="inactive">
              <router-link to="/riwayat">Hasil Prediksi</router-link>
            </li>
            <li class="active">
              <router-link to="/prediksi">Tes DNA</router-link>
            </li>
          </ul>
        </div>
        <div class="container-body">
          <div class="container-heading">
            <h1>Tes DNA</h1>
            <p>Periksa apakah jangan-jangan DNA Anda mengidap penyakit...?</p>
          </div>
          <div class="container-form-prediksi">
            <form>
              <div class="form-nama-user-prediksi">
                <label for="nama-user-prediksi">Nama Pengguna:</label><br />
                <input
                  type="text"
                  id="nama-user-prediksi"
                  name="nama-user-prediksi"
                  placeholder="Nama Pengguna..."
                />
              </div>
              <div class="form-nama-penyakit-prediksi">
                <label for="nama-penyakit-prediksi">Prediksi Penyakit:</label
                ><br />
                <input
                  type="text"
                  id="nama-penyakit-prediksi"
                  name="nama-penyakit-prediksi"
                  placeholder="Prediksi Penyakit..."
                />
              </div>
              <div class="form-sequence-dna">
                <label for="sequence-dna">Sequence DNA:</label><br />
                <input
                  type="file"
                  id="sequence-dna"
                  name="sequence-dna"
                  accept=".txt"
                /><br />
                <div class="text-ketentuan">
                  <p>* File yang diperbolehkan adalah .txt</p>
                  <p>* Tidak ada huruf kecil</p>
                  <p>* Tidak boleh ada huruf selain AGCT</p>
                  <p>* Tidak ada spasi</p>
                </div>
              </div>
              <div class="form-submit-prediksi">
                <button type="button" class="button-submit">Submit!</button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/* HEADER */
html,
body {
  position: absolute;
  width: 100%;
  height: 100%;
}

.background {
  position: relative;
  overflow: hidden;
  height: 650px;
}

.container-header {
  position: absolute;
  width: 100%;
  height: 100%;
}

.container-header .logo-full {
  position: relative;
  transform: translate(-150%, -50%);
  width: 25%;
}

.container-header ul {
  position: relative;
  float: right;
  list-style-type: none;
  transform: translate(0%, -368%);
  margin-right: 3%;
  z-index: 1;
}

.container-header ul li {
  display: inline-block;
  padding: 10px;
}

.container-header ul li a {
  text-decoration: none;
  color: black;
  padding: 5px 20px;
  border: 1px solid transparent;
  border-radius: 30px;
  transition: 0.3s ease;
}

.container-heading {
  transform: translate(0%, 110%);
}

.container-heading p {
  transform: translate(0%, -50%);
}

.container-form-prediksi {
  background-color: rgb(165, 102, 244);
  width: 800px;
  height: 350px;
  border-radius: 8%/20%;
  transform: translate(45%, 28%);
  z-index: 0;
  color: white;
}

ul li a:hover {
  background-color: rgb(165, 102, 244);
  color: white;
}

ul li.active a {
  background-color: white;
  color: rgb(94, 31, 172);
  font-weight: 700;
}

.container-form-prediksi .form-nama-penyakit-prediksi {
  position: relative;
  transform: translate(-25%, 150%);
  font-family: "Quicksand", sans-serif;
}

.container-form-prediksi .form-nama-penyakit-prediksi label {
  position: relative;
}

.container-form-prediksi .form-nama-penyakit-prediksi input {
  position: relative;
  width: 300px;
  height: 30px;
  border-radius: 8px;
  border: 1px solid rgb(165, 102, 244);
  padding: 5px;
  padding-left: 12px;
  transform: translate(0%, 20%);
}

.container-form-prediksi .form-nama-user-prediksi {
  position: relative;
  transform: translate(-25%, 110%);
  font-family: "Quicksand", sans-serif;
}

.container-form-prediksi .form-nama-user-prediksi label {
  position: relative;
}

.container-form-prediksi .form-nama-user-prediksi input {
  position: relative;
  width: 300px;
  height: 30px;
  border-radius: 8px;
  border: 1px solid rgb(165, 102, 244);
  padding: 5px;
  padding-left: 12px;
  transform: translate(0%, 20%);
}

.container-form-prediksi .form-sequence-dna {
  position: relative;
  transform: translate(25%, -34%);
  font-family: "Quicksand", sans-serif;
}

.container-form-prediksi .form-sequence-dna label {
  position: relative;
}

.container-form-prediksi .form-sequence-dna input {
  position: relative;
  width: 300px;
  height: 30px;
  border-radius: 8px;
  border: 1px solid rgb(165, 102, 244);
  padding: 5px;
  transform: translate(0%, 30%);
}

.container-form-prediksi .form-sequence-dna .text-ketentuan {
  position: relative;
  transform: translate(32%, 0%);
  font-family: "Quicksand", sans-serif;
  font-size: 11px;
  text-align: left;
}

.container-form-prediksi .form-submit-prediksi .button-submit {
  position: relative;
  transform: translate(0%, -65%);
  background-color: rgb(255, 255, 255);
  border: none;
  color: rgb(94, 31, 172);
  padding: 10px 20px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  border-radius: 15px;
  font-family: "Quicksand", sans-serif;
  font-weight: 700;
}

.container-form-prediksi .form-submit-prediksi .button-submit:hover {
  position: relative;
  background-color: rgb(240, 240, 240);
  border: none;
  color: rgb(36, 4, 77);
  padding: 10px 20px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  border-radius: 15px;
}
</style>
