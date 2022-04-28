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
                    <li class="inactive">
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
            <div class="container-redirect">
                <div class="heading">
                    <h1>Cari tahu penyakit yang DNA-mu miliki!</h1>
                    <router-link to="/prediksi">
                        <button class="button-redirect">Coba Sekarang!</button>
                    </router-link>
                </div>
            </div>
            <div class="container-body-profile">
                <div class="container-heading">
                    <h1>Perkenalkan,</h1>
                    <p class="nama-kelompok">Istighfar Ikhtiar Tawakal :V</p>
                </div>
                <div class="container-profile">
                    <div class="container-orang">
                        <div class="container-foto">
                            <img src="../assets/foto-fay.png" alt="" class="foto">
                        </div>
                        <div class="container-nama">
                            <h3 class="nama">Fayza Nadia</h3>
                            <h4 class="nim">13520001</h4>
                        </div>
                        <div class="container-biodata">
                            <p class="biodata">
                                Seorang mahasiswa Teknik Informatika yang bobo mulu gws.
                            </p>
                        </div>
                        <div class="instagram">
                            <div class="container-instagram">
                                <a href="https://instagram.com/fayzanadia" target="_blank">
                                    <img src="../assets/logo-ig.png" alt="" class="logo-ig">
                                </a>
                            </div>
                        </div>
                    </div>
                    <div class="container-orang">
                        <div class="container-foto">
                            <img src="../assets/foto-fajar.png" alt="" class="foto">
                        </div>
                        <div class="container-nama">
                            <h3 class="nama">M. Fajar Ramadhan</h3>
                            <h4 class="nim">13520026</h4>
                        </div>
                        <div class="container-biodata">
                            <p class="biodata">
                                Seorang mahasiswa Teknik Informatika yang naik kereta jam 6 pagi aneh bgt.
                            </p>
                        </div>
                        <div class="instagram">
                            <div class="container-instagram">
                                <a href="https://instagram.com/fajarrm" target="_blank">
                                    <img src="../assets/logo-ig.png" alt="" class="logo-ig">
                                </a>
                            </div>
                        </div>
                    </div>
                    <div class="container-orang">
                        <div class="container-foto">
                            <img src="../assets/foto-dimas.png" alt="" class="foto">
                        </div>
                        <div class="container-nama">
                            <h3 class="nama">Dimas Faidh Muzakki</h3>
                            <h4 class="nim">13520156</h4>
                        </div>
                        <div class="container-biodata">
                            <p class="biodata">
                                Seorang mahasiswa Teknik Informatika yang kamar kosnya gede banget gamasuk akal.
                            </p>
                        </div>
                        <div class="instagram">
                            <div class="container-instagram">
                                <a href="https://instagram.com/oohmasdim" target="_blank">
                                    <img src="../assets/logo-ig.png" alt="" class="logo-ig">
                                </a>
                            </div>
                        </div>
                    </div>
                </div>
            </div>      
        </div>
    </div>
</div>
</template>

<script>
export default {
    
}
</script>

<style>
html,
body {
  position: absolute;
  width: 100%;
  height: 100%;
}

body {
  background-image: url("../assets/bg.png");
  background-size: cover; 
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
  transform: translate(0%, 80%);
}

.container-heading p {
  transform: translate(0%, -50%);
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

.container-redirect {
    position: absolute;
    width: 100%;
    height: 100%;
    transform: translate(0%, 50%);
}

.container-redirect h1{
    text-align: left;
    width: 350px;
    font-size: 48px;
    padding-left: 110px;
    transform: translate(0%, -50%);
}

.container-redirect .button-redirect {
    padding: 10px 20px;
    border-radius: 30px;
    font-size: 18px;
    font-weight: 500;
    cursor: pointer;
    transform: translate(-320%, -230%);
    font-family: "Quicksand", helvetica;
    background-color: white;
    color: rgb(94, 31, 172);
    border: 1px solid rgb(94, 31, 172);    
}

.container-redirect .button-redirect:hover{
    background-color: rgb(94, 31, 172);
    color: white;
    border: none;
    transition: 0.5s ease;
}

.container-body-profile {
    transform: translate(0%, 120%);
    padding-bottom: 150px;
}

.container-body-profile .container-profile{
    position: relative;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    transform: translate(0%, 20%);
    gap: 5%;
}

.container-profile .container-orang{
    position: relative;
    width: 20%;
    height: 40%;
}

.container-biodata{
    position: relative;
    width: 80%;
    margin: auto;
}

.nama-kelompok {
    font-size: 24px;
}

.container-nama{
    transform: translate(0%, 20%);
}

.nim{
    transform: translate(0%, -80%);
}

.container-foto .foto{
    position: relative;
    width: 100%;
    margin-bottom: -10%;
}

.instagram .container-instagram .logo-ig{
    position: relative;
    width: 7%;
    margin: 5%; 
}

</style>