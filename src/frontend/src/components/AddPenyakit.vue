<template>
  <div class="background">
    <div class="top-cont">
      <div class="canvas">
        <div class="container-header">
          <div class="container-logo">
            <router-link to="/">
                <img src="../assets/logo-full.svg" alt="" class="logo-full" />
            </router-link>
          </div>
          <ul class="navigation-bar">
            <li class="active">
              <router-link to="/penyakit">Tambah Penyakit</router-link>
            </li>
            <li class="inactive">
              <router-link to="/riwayat">Hasil Prediksi</router-link>
            </li>
            <li class="inactive">
              <router-link to="/prediksi">Tes DNA</router-link>
            </li>
          </ul>
        </div>
        <div class="container-body">
          <div class="container-heading">
            <h1>Tambah Penyakit</h1>
            <p>Tambahkan penyakit baru ke dalam sistem kami!</p>
          </div>
          <div class="container-form">
            <form>
              <div class="form-nama-penyakit">
                <label for="nama-penyakit">Nama Penyakit:</label><br />
                <input
                  type="text"
                  id="nama-penyakit"
                  name="nama-penyakit"
                  placeholder="Penyakit..."
                  ref="namaPenyakit"
                />
              </div>
              <div class="form-sequence-dna">
                <label for="sequence-dna">Sequence DNA:</label><br />
                <input
                  type="file"
                  @change="upload"
                  id="sequence-dna"
                  name="sequence-dna"
                  accept=".txt"
                  ref="doc"
                /><br />
                <div class="text-ketentuan">
                  <p v-bind:class="{'green': !txtFlag, 'red': txtFlag}">* File yang diperbolehkan adalah .txt</p>
                  <p v-bind:class="{'green': !capitalFlag, 'red': capitalFlag}">* Tidak ada huruf kecil</p>
                  <p v-bind:class="{'green': !acgtFlag, 'red': acgtFlag}">* Tidak boleh ada huruf selain AGCT</p>
                  <p v-bind:class="{'green': !spaceFlag, 'red': spaceFlag}">* Tidak ada spasi</p>
                </div>
              </div>
              <div class="form-submit">
                <button @click="submitPenyakit" type="button" class="button-submit">Submit!</button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: "AddPenyakit",
  data() {
    return {
      acgtFlag : true,
      txtFlag : true,
      capitalFlag : true,
      spaceFlag : true,
      file : null,
      content : null,
    };
  },
  methods: {
    resetFlags() {
      this.txtFlag = true;
      this.acgtFlag = true;
      this.capitalFlag = true;
      this.spaceFlag = true;
    },
    async submitPenyakit(e){
      const namaPenyakit = this.$refs.namaPenyakit.value;
      if (this.txtFlag || this.capitalFlag || this.acgtFlag || this.spaceFlag || this.content === "") {
        alert("Mohon periksa kembali file anda!");
      } else if (namaPenyakit === "") {
        alert("Mohon isi nama penyakit!");
      // } else if penyakitnya udah ada
      } else {
        const formData = {
          nama : namaPenyakit,
          rantai : this.content
        }
        console.log(formData)
        await axios({
          method: "post",
          url: "https://dna-matching-api.herokuapp.com/penyakit/create",
          data: formData,
          headers: { "Content-Type": "application/json" },
        })
          .then(function (response) {
            //handle success
            alert("Penambahan penyakit berhasil!");
          })
          .catch(function (response) {
            //handle error
            if (response.status === 400) {
              alert("Data penyakit tersebut sudah ada!");
            }else{
              alert("Penambahan penyakit gagal!");
            }
          });
          this.$refs.namaPenyakit.value = "";
          this.$refs.doc.value = null;
          this.resetFlags();

      }

      },
    readUploadedFileAsText(inputFile){
      const temporaryFileReader = new FileReader();
      
      if(inputFile.name.includes(".txt")){
        this.txtFlag = false;
      }

      return new Promise((resolve, reject) => {
        temporaryFileReader.onerror = () => {
          temporaryFileReader.abort();
          reject(new DOMException("Problem parsing input file."));
        };

        temporaryFileReader.onload = () => {
          resolve(temporaryFileReader.result);
        };
        temporaryFileReader.readAsText(inputFile);
      });
    },
    async upload(){
      this.file = this.$refs.doc.files[0];
  
      try {
        this.content = await this.readUploadedFileAsText(this.file)  
      } catch (e) {
        alert(e)
      }
      const rg_atgc = /[^AGCT]/g;
      const rg_capital = /[^A-Z]/g;
      const rg_space = /\s/g;
      this.acgtFlag = rg_atgc.test(this.content);
      this.capitalFlag = rg_capital.test(this.content);
      this.spaceFlag = rg_space.test(this.content);
      console.log(this.content);
    },
  },
};
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

.green {
  color: #B0FD96  ;
}

.red {
  color: #fff;
}

.background {
  position: relative;
  height: 100%;
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

.container-form {
  background-color: rgb(165, 102, 244);
  width: 800px;
  height: 350px;
  border-radius: 8%/20%;
  transform: translate(0%, 28%);
  z-index: 0;
  color: white;
  margin: auto;
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

.container-form .form-nama-penyakit {
  position: relative;
  transform: translate(-25%, 110%);
  font-family: "Quicksand", sans-serif;
}

.container-form .form-nama-penyakit label {
  position: relative;
}

.container-form .form-nama-penyakit input {
  position: relative;
  width: 300px;
  height: 30px;
  border-radius: 8px;
  border: 1px solid rgb(165, 102, 244);
  padding: 5px;
  padding-left: 12px;
  transform: translate(0%, 20%);
}

.container-form .form-sequence-dna {
  position: relative;
  transform: translate(25%, 5%);
  font-family: "Quicksand", sans-serif;
}

.container-form .form-sequence-dna label {
  position: relative;
}

.container-form .form-sequence-dna input {
  position: relative;
  width: 300px;
  height: 30px;
  border-radius: 8px;
  border: 1px solid rgb(165, 102, 244);
  padding: 5px;
  transform: translate(0%, 30%);
}

.container-form .form-sequence-dna .text-ketentuan {
  position: relative;
  transform: translate(32%, 0%);
  font-family: "Quicksand", sans-serif;
  font-size: 11px;
  text-align: left;
}

.container-form .form-submit .button-submit {
  position: relative;
  transform: translate(0%, 90%);
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

.container-form .form-submit .button-submit:hover {
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
