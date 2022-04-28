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
            <li class="active">
              <router-link to="/riwayat">Hasil Prediksi</router-link>
            </li>
            <li class="inactive">
              <router-link to="/prediksi">Tes DNA</router-link>
            </li>
          </ul>
        </div>
        <div class="container-body">
          <div class="container-heading">
            <h1>Hasil Prediksi</h1>
            <p>Periksa hasil seluruh tes yang telah diajukan!</p>
          </div>
          <div class="container-form">
            <form>
              <div class="form-search">
                <input
                  type="text"
                  id="query-search"
                  name="query-search"
                  ref="querySearch"
                  placeholder="<YYYY-MM-DD><spasi><nama_penyakit>..."
                />
              </div>
              <div class="form-submit-riwayat">
                <button @click="submit" type="button" class="button-submit">Search!</button>
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
  data(){
    return{
      tanggalFlag : false,
      penyakitFlag : false,
      spaceFlag : false
    };
  },
  methods:{
    async submit(e){
      let query = this.$refs.querySearch.value;
      let tanggal;
      let penyakit;
      console.log(query)
      // regex
      const rg_tanggal = /\d{4}\-(0?[1-9]|1[012])\-(0?[1-9]|[12][0-9]|3[01])*/g;
      const rg_penyakit = /[a-zA-Z0-9]*/g;
      const rg_space = /\s/g;

      // cek tanggal
      this.tanggalFlag = rg_tanggal.test(query);
      if (this.tanggalFlag){
        tanggal = query.match(rg_tanggal)[0];
        query = query.replace(rg_tanggal, '');
      }
      
      //ilangin spasi
      query = query.replace(rg_space, '');
      
      // ambil penyakit
      this.penyakitFlag = rg_penyakit.test(query);
      if (this.penyakitFlag){
        penyakit = query.match(rg_penyakit)[0];
        if (penyakit === ""){
          this.penyakitFlag = false;
        }
      }

      console.log(this.tanggalFlag)
      console.log(tanggal)
      console.log(this.penyakitFlag)
      console.log(penyakit)


      if (this.tanggalFlag && this.penyakitFlag){
        await axios(
          {
            method: 'get',
            url: 'https://dna-matching-api.herokuapp.com/pemeriksaan/get/whenwhat?penyakit=' + penyakit + '&tanggal=' + tanggal,
            headers: {
              'Content-Type': 'application/json'
            }
          }
        ).then(
          (response) => {
            console.log(response.data);
          }
        ).catch(
          (error) => {
            
          }
        );
      }
      else if (this.tanggalFlag && !this.penyakitFlag){
        await axios(
          {
            method: 'get',
            url: 'https://dna-matching-api.herokuapp.com/pemeriksaan/get/when?tanggal=' + tanggal,
            headers: {
              'Content-Type': 'application/json'
            }
          }
        ).then(
          (response) => {
            console.log(response.data);
          }
        ).catch(
          (error) => {
            
          }
        );
      }
      else if (!this.tanggalFlag && this.penyakitFlag){
        await axios(
          {
            method: 'get',
            url: 'https://dna-matching-api.herokuapp.com/pemeriksaan/get/what?penyakit=' + penyakit,
            headers: {
              'Content-Type': 'application/json'
            }
          }
        ).then(
          (response) => {
            console.log(response.data);
          }
        ).catch(
          (error) => {
            
          }
        );
      }
      else{
        await axios(
          {
            method: 'get',
            url: 'https://dna-matching-api.herokuapp.com/pemeriksaan/get',
            headers: {
              'Content-Type': 'application/json'
            }
          }
        ).then(
          (response) => {
            console.log(response.data);
          }
        ).catch(
          (error) => {
            
          }
        );
      }

    }
  }
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

.container-form {
  background-color: rgb(165, 102, 244);
  width: 800px;
  height: 80px;
  border-radius: 2.5%/25%;
  transform: translate(45%, 130%);
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

.container-form .form-search {
  position: relative;
  transform: translate(-25%, 110%);
  font-family: "Quicksand", sans-serif;
}

.container-form .form-search label {
  position: relative;
}

.container-form .form-search input {
  position: relative;
  width: 620px;
  height: 30px;
  border-radius: 8px;
  border: 1px solid rgb(165, 102, 244);
  padding: 5px;
  padding-left: 12px;
  transform: translate(22%, -60%);
}

.container-form .form-submit-riwayat .button-submit {
  position: relative;
  transform: translate(335%, -50%);
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

.container-form .form-submit-riwayat .button-submit:hover {
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
