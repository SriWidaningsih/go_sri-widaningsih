CREATE DATABASE Tiket_Kereta;
USE Tiket_Kereta;
CREATE TABLE identitas (
  id_identitas INT AUTO_INCREMENT PRIMARY KEY,
  nama VARCHAR(255) NOT NULL,
  tempat_lahir VARCHAR(255) NOT NULL,
  tgl_lahir DATE NOT NULL,
  alamat VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  no_telp VARCHAR(20) NOT NULL
);

CREATE TABLE pelanggan (
  id_pelanggan INT AUTO_INCREMENT PRIMARY KEY,
  id_identitas INT NOT NULL,
  FOREIGN KEY (id_identitas) REFERENCES identitas(id_identitas)
);

CREATE TABLE pembayaran (
  id_pembayaran INT AUTO_INCREMENT PRIMARY KEY,
  id_pelanggan INT NOT NULL,
  kode_jadwal VARCHAR(10) NOT NULL,
  jumlah_bayar INT NOT NULL,
  waktu_bayar DATETIME NOT NULL,
  FOREIGN KEY (id_pelanggan) REFERENCES pelanggan(id_pelanggan),
  FOREIGN KEY (kode_jadwal) REFERENCES jadwal(kode_jadwal)
);

CREATE TABLE tiket (
  id_tiket INT AUTO_INCREMENT PRIMARY KEY,
  id_pembayaran INT NOT NULL,
  id_stasiun_asal INT NOT NULL,
  id_stasiun_tujuan INT NOT NULL,
  kode_jadwal VARCHAR(10) NOT NULL,
  jumlah_tiket INT NOT NULL,
  FOREIGN KEY (id_pembayaran) REFERENCES pembayaran(id_pembayaran),
  FOREIGN KEY (kode_jadwal) REFERENCES jadwal(kode_jadwal),
  FOREIGN KEY (id_stasiun_asal) REFERENCES stasiun(id_stasiun),
  FOREIGN KEY (id_stasiun_tujuan) REFERENCES stasiun(id_stasiun)
);

CREATE TABLE jadwal (
  kode_jadwal VARCHAR(10) PRIMARY KEY,
  id_stasiun_asal INT NOT NULL,
  id_stasiun_tujuan INT NOT NULL,
  waktu_keberangkatan DATETIME NOT NULL,
  waktu_tiba DATETIME NOT NULL,
  harga INT NOT NULL,
  FOREIGN KEY (id_stasiun_asal) REFERENCES stasiun(id_stasiun),
  FOREIGN KEY (id_stasiun_tujuan) REFERENCES stasiun(id_stasiun)
);

CREATE TABLE stasiun (
  id_stasiun INT AUTO_INCREMENT PRIMARY KEY,
  nama_stasiun VARCHAR(255) NOT NULL,
  alamat_stasiun VARCHAR(255) NOT NULL,
  no_telp_stasiun VARCHAR(20) NOT NULL
);
ALTER TABLE stasiun ADD id_stasiun_asal INT NOT NULL AFTER id_stasiun;
ALTER TABLE stasiun ADD id_stasiun_tujuan INT NOT NULL AFTER id_stasiun_asal;


CREATE TABLE admin (
  id_admin INT AUTO_INCREMENT PRIMARY KEY,
  id_identitas INT NOT NULL,
  FOREIGN KEY (id_identitas) REFERENCES identitas(id_identitas)
);


